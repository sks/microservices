package httputil

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		IdleConnTimeout:       10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// Do Make the HTTP call with custom client
func Do(req *http.Request) (*http.Response, error) {
	return httpClient.Do(req)
}

// GetResponse get response after closing all body
func GetResponse(req *http.Request) ([]byte, int, error) {
	resp, err := Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}
