package smsotpverifier

import (
	context "context"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type server struct {
	logger     *zap.Logger
	smsAdapter Adapter
	db         *gorm.DB
}

func (s *server) SendOTP(ctx context.Context, req *SendOTPRequest) (*SendOTPResponse, error) {
	otpRequest, err := newOTPRequest(req)
	if err != nil {
		return nil, err
	}
	err = s.smsAdapter.SendSMS(ctx, req.ToNumber,
		fmt.Sprintf("%s: \nOTP: %s", req.Message, otpRequest.OTP),
		req.FromNumber)
	if err != nil {
		return nil, err
	}
	err = s.db.Save(otpRequest).Error
	return &SendOTPResponse{
		Uuid: otpRequest.ID,
	}, err
}

func (s *server) VerifyOTP(ctx context.Context, req *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	otpRequest := OTPRequest{}
	err := s.db.Where("id=?", req.GetUuid).Find(&otpRequest).Error
	if err != nil {
		return nil, err
	}

	err = otpRequest.Validate(req.Otp)
	er := s.db.Model(&otpRequest).Update("attempts", otpRequest.Attempts).Error
	if er != nil {
		s.logger.Error("Error updating the attempt count", zap.Error(err), zap.String("otp", otpRequest.String()))
	}
	if err != nil {
		return nil, err
	}
	err = s.db.Delete(otpRequest).Error
	if err != nil {
		return nil, err
	}
	return &VerifyOTPResponse{
		AttemptCount: otpRequest.Attempts,
	}, nil
}
