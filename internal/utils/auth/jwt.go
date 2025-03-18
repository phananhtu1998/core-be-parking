package auth

import (
	"context"
	"fmt"
	"go-backend-api/global"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}

func CreateToken(uuidToken string) (string, error) {
	// 1. Set time expiration
	timeEx := global.Config.JWT.ACCESS_TOKEN
	if timeEx == "" {
		timeEx = "72h"
	}
	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}
	now := time.Now()
	expiresAt := now.Add(expiration)
	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "parkingdevgo",
			Subject:   uuidToken,
		},
	})
}
func CreateRefreshToken(uuidToken string) (string, error) {
	// 1. Set time expiration
	timeEx := global.Config.JWT.REFRESH_TOKEN
	if timeEx == "" {
		timeEx = "168h"
	}
	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}
	now := time.Now()
	expiresAt := now.Add(expiration)
	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "parkingdevgo",
			Subject:   uuidToken,
		},
	})
}
func ParseJwtTokenSubject(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.API_SECRET_KEY), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// validate token
func VerifyTokenSubject(token string) (*jwt.StandardClaims, error) {
	claims, err := ParseJwtTokenSubject(token)
	if err != nil {
		return nil, err
	}
	if err = claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}

// CheckBlacklist kiểm tra xem token có trong danh sách đen không
func CheckBlacklist(key string) bool {
	// Tạo key Redis đúng định dạng
	redisKey := fmt.Sprintf("TOKEN_BLACK_LIST_%s", key)

	// Kiểm tra key trong Redis
	_, err := global.Rdb.Get(context.Background(), redisKey).Result()

	// Nếu không có lỗi => key tồn tại (token bị blacklist)
	if err == nil {
		return true
	}
	// Nếu lỗi là "key not found" => token chưa bị blacklist
	if err == redis.Nil {
		return false
	}
	// Nếu có lỗi khác => log lỗi (tuỳ chỉnh nếu cần)
	fmt.Println("Lỗi khi kiểm tra Redis:", err)
	return false
}
