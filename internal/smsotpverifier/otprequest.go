package smsotpverifier

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/sks/microservices/internal/fnhelper"
	"github.com/sks/microservices/internal/models"
)

var (
	// ErrMaximumAttemptsExceeded max attempts exceeded for OTP verification
	ErrMaximumAttemptsExceeded = errors.New("sms.verify.otp.maximum.attempts.exceeded")
	// ErrOTPExpired denotes that the otp has expired
	ErrOTPExpired = errors.New("sms.verify.otp.expired")
	// ErrInvalidOTP thrown in case of invalid OTP
	ErrInvalidOTP = errors.New("sms.otp.verify.invalid.otp.error")
)

// OTPRequest details on the OTP request
type OTPRequest struct {
	models.Base
	ID              string `gorm:"id"`
	OTP             string `gorm:"otp"`
	Attempts        uint32 `gorm:"attempts"`
	MaxAttempts     uint32 `gorm:"maxattempts"`
	ValidForMinutes uint32 `gorm:"validforminutes"`
}

func (o OTPRequest) String() string {
	return fmt.Sprintf("%s [%d/%d]", o.ID, o.Attempts, o.MaxAttempts)
}

// newOTPRequest create new OTP Request
func newOTPRequest(req *SendOTPRequest) (*OTPRequest, error) {
	var otp string
	err := fnhelper.Retry(func() error {
		var err error
		otp, err = GenerateRandomNumber(6)
		return err
	}, 3, time.Nanosecond)
	if err != nil {
		return nil, err
	}
	return &OTPRequest{
		ID:              uuid.New().String(),
		Base:            models.NewBase("admin"),
		Attempts:        0,
		MaxAttempts:     getValOrDefault(req.MaxAttempts, 3),
		ValidForMinutes: getValOrDefault(req.ValidForMinutes, 3),
		OTP:             otp,
	}, nil
}

func getValOrDefault(val, defaultVal uint32) uint32 {
	if val != 0 {
		return val
	}
	return defaultVal
}

// Validate validate if the given otp is a match
func (o *OTPRequest) Validate(otp string) error {
	err := o.HasExpired()
	if err != nil {
		return err
	}
	o.Attempts++
	if o.OTP != otp {
		return ErrInvalidOTP
	}
	return nil
}

// HasExpired denotes if the otp has expired
func (o OTPRequest) HasExpired() error {
	if o.Attempts >= o.MaxAttempts {
		return ErrMaximumAttemptsExceeded
	}
	if o.CreatedOn.Add(time.Duration(o.ValidForMinutes) * time.Minute).Before(time.Now()) {
		return ErrOTPExpired
	}
	return nil
}

// GenerateRandomNumber generate random number of max digits
func GenerateRandomNumber(max int) (string, error) {
	b := make([]byte, max)
	_, err := io.ReadAtLeast(rand.Reader, b, max)
	if err != nil {
		return "", err
	}
	for i := range b {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

var table = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
