package services

import (
	"fmt"
	"time"

	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/constants"
	"github.com/amirmalekian/Mashinchi/src/data/cache"
	"github.com/amirmalekian/Mashinchi/src/pkg/logging"
	"github.com/amirmalekian/Mashinchi/src/pkg/service_errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redisClient: redis}
}

func (s *OtpService) SetOtp(ctx *gin.Context, phoneNumber, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, phoneNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[OtpDto](ctx, s.redisClient, key)

	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	err = cache.Set(ctx, s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}

	return nil
}

func (s *OtpService) ValidateOtp(ctx *gin.Context, phoneNumber, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, phoneNumber)
	res, err := cache.Get[OtpDto](ctx, s.redisClient, key)
	if err != nil {
		return err
	} else if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpInValid}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(ctx, s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
