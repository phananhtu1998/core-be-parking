package repo

import (
	"fmt"
	"go-backend-api/global"
	"time"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

// AddOTP implements IUserAuthRepository.
func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	// panic("unimplemented")
	key := fmt.Sprintf("user:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
