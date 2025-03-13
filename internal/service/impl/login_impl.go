package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"go-backend-api/global"
	consts "go-backend-api/internal/const"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/internal/utils/cache"
	"go-backend-api/internal/utils/crypto"
	"go-backend-api/pkg/response"
	"log"
	"time"
)

type sLogin struct {
	r *database.Queries
}

func NewLoginImpl(r *database.Queries) *sLogin {
	return &sLogin{r: r}
}

func (s *sLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	accountBase, err := s.r.GetOneAccountInfoAdmin(ctx, in.Email)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	out.ID = accountBase.ID
	out.Email = accountBase.Email
	log.Println("matching pass: ", crypto.MatchingPassword(accountBase.Password, in.Password, accountBase.Salt))
	if !crypto.MatchingPassword(accountBase.Password, in.Password, accountBase.Salt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}
	subToken := utils.GenerateCliTokenUUID(int(accountBase.Number))
	log.Println("subtoken:", subToken)
	infoAccount, err := s.r.GetAccountById(ctx, accountBase.ID)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("lỗi ở phần lấy thông tin tài khoản")
	}
	infoAccountJson, err := json.Marshal(infoAccount)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}
	err = global.Rdb.Set(ctx, subToken, infoAccountJson, time.Duration(consts.REFRESH_TOKEN)*time.Hour).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("lỗi ở phần set vào redis")
	}
	out.Token, err = auth.CreateToken(subToken)
	out.RefreshToken, err = auth.CreateRefreshToken(subToken)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("lỗi ở phần tạo token: %v", err)
	}
	return response.ErrCodeSucces, out, err
}
func (s *sLogin) Logout(ctx context.Context) (codeResult int, err error) {
	// Lấy subjectUUID từ context
	subjectUUID := ctx.Value("subjectUUID")
	if subjectUUID == nil {
		return response.ErrCodeAuthFailed, fmt.Errorf("subjectUUID not found in context")
	}
	// Lấy thông tin user từ cache
	var infoUser model.GetCacheToken
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, err
	}
	log.Println("User info from cache:", infoUser.ID)
	return response.ErrCodeSucces, err
}
