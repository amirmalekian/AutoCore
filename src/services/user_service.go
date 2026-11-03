package services

import (
	"github.com/amirmalekian/Mashinchi/src/api/dto"
	"github.com/amirmalekian/Mashinchi/src/common"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/data/db"
	"github.com/amirmalekian/Mashinchi/src/pkg/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	logger     *logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDB()
	logger := logging.NewLogger(cfg)
	return &UserService{
		logger:     &logger,
		cfg:        cfg,
		database:   database,
		otpService: NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(ctx *gin.Context, req *dto.GetOtpRequest) error {
	otp := common.GenerateOTP()
	err := s.otpService.SetOtp(ctx, req.PhoneNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
