package health

import (
	"context"
	"net/http"
	"runtime"

	"github.com/sks/microservices/internal/httputil"
	"go.uber.org/zap"
)

var availableEndpoints = map[string]string{
	"/health": "Runs the health checks against various registered modules",
	"/memory": "Prints the current memory usage",
	"/":       "Shows this help page",
}

// Reporter an interface
type Reporter interface {
	AddModule(name string, function func(context.Context) (interface{}, error))
}

func newReporter(logger *zap.Logger) *reporter {
	reporter := &reporter{
		logger:    logger.Named("health"),
		functions: make(map[string]func(context.Context) (interface{}, error)),
		server:    http.NewServeMux(),
	}
	reporter.server.HandleFunc("/health", reporter.health)
	reporter.server.HandleFunc("/memory", reporter.memory)
	reporter.server.HandleFunc("/", reporter.index)
	return reporter
}

type reporter struct {
	logger    *zap.Logger
	server    *http.ServeMux
	functions map[string]func(context.Context) (interface{}, error)
}

func (r *reporter) AddModule(name string, function func(context.Context) (interface{}, error)) {
	r.functions[name] = function
}

func (r *reporter) memory(w http.ResponseWriter, req *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	httputil.WriteJSONResponse(w, m)
}
func (r *reporter) index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	}
	httputil.WriteJSONResponse(w, availableEndpoints)
}

type healthReport struct {
	Passed  map[string]interface{} `json:"passed"`
	Errored map[string]string      `json:"errored"`
}

func (r *reporter) health(w http.ResponseWriter, req *http.Request) {
	response := healthReport{
		Passed:  make(map[string]interface{}),
		Errored: make(map[string]string),
	}
	r.logger.Info("Need to check health checks", zap.Int("count", len(r.functions)))
	for i := range r.functions {
		val, err := r.functions[i](req.Context())
		if err != nil {
			r.logger.Warn("Health check failed",
				zap.String("service", i), zap.Error(err))
			response.Errored[i] = err.Error()
			continue
		}
		r.logger.Debug("completed health check", zap.String("service", i))
		response.Passed[i] = val
	}
	httputil.WriteJSONResponse(w, response)
}
