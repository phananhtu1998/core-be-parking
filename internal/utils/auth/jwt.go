package auth

import (
	"context"
	"errors"
	"fmt"
	"go-backend-api/global"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/cache"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

var ctx = context.Background()

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
func CreateTokenNoExpiration(dateStart, dateEnd string) (string, error) {
	// Create custom claims
	claims := jwt.MapClaims{
		"jti":       uuid.New().String(),
		"iss":       "parkingdevgo",
		"datestart": dateStart,
		"dateend":   dateEnd,
	}

	// Create token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
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
	_, err := global.Rdb.Get(ctx, redisKey).Result()

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

// CheckTokenRevoked kiểm tra xem token có bị thu hồi do đổi mật khẩu không
func CheckTokenRevoked(subtoken string, issuedAt int64) (bool, error) {
	ctx := context.Background()
	var infoUser model.GetCacheToken

	// Lấy thông tin user từ cache
	if err := cache.GetCache(ctx, subtoken, &infoUser); err != nil {
		log.Println("Lỗi khi lấy cache token:", err)
		// Nếu không có cache, giả sử token chưa bị thu hồi
		return false, nil
	}

	// Tạo key Redis để kiểm tra thời gian đổi mật khẩu
	invalidationKey := fmt.Sprintf("TOKEN_IAT_AVAILABLE_%s", infoUser.ID)
	changepasswordTimestampStr, err := global.Rdb.Get(ctx, invalidationKey).Result()

	if err == redis.Nil {
		// Không có dữ liệu trong Redis => Token chưa bị thu hồi
		return false, nil
	} else if err != nil {
		// Lỗi khi truy vấn Redis
		log.Println("Lỗi khi lấy giá trị từ Redis:", err)
		return false, err
	}

	// Chuyển đổi timestamp từ string sang int64
	changepasswordTimestamp, err := strconv.ParseInt(changepasswordTimestampStr, 10, 64)
	if err != nil {
		log.Println("Lỗi khi parse timestamp từ Redis:", err)
		return false, fmt.Errorf("lỗi khi parse timestamp từ Redis: %v", err)
	}

	// Kiểm tra nếu token đã bị thu hồi do thay đổi mật khẩu
	if issuedAt < changepasswordTimestamp {
		log.Println("Token revoked due to password change")
		return true, nil
	}

	return false, nil
}

func ParseJwtTokenPayload(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.API_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
