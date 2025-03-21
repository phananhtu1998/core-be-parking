package impl

import (
	"context"
	"encoding/json"
	"errors"
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

	"github.com/google/uuid"
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
	if accountBase.Status == false {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("tài khoản đã bị khóa")
	}
	subToken := utils.GenerateCliTokenUUID(int(accountBase.Number))
	log.Println("subtoken:", subToken)
	infoAccount, err := s.r.GetAccountById(ctx, accountBase.ID)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("lỗi ở phần lấy thông tin tài khoản")
	}
	infoAccountJson, err := json.Marshal(infoAccount)
	// check bảng keytoken có tồn tại hay chưa
	// Nếu chưa thì insert
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}
	err = global.Rdb.Set(ctx, subToken, infoAccountJson, time.Duration(consts.REFRESH_TOKEN)*time.Hour).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("lỗi ở phần set vào redis")
	}
	out.AccessToken, err = auth.CreateToken(subToken)
	out.RefreshToken, err = auth.CreateRefreshToken(subToken)
	getAccountKT, err := s.r.CountByAccount(ctx, accountBase.ID)
	if getAccountKT > 0 {
		err := s.r.UpdateRefreshTokenAndUsedTokens(ctx, database.UpdateRefreshTokenAndUsedTokensParams{
			AccountID:    accountBase.ID,
			RefreshToken: out.RefreshToken,
		})
		if err != nil {
			return response.ErrInvalidToken, out, fmt.Errorf("lỗi update key: %v", err)
		}
	} else {
		err := s.r.InsertKey(ctx, database.InsertKeyParams{
			ID:           uuid.NewString(),
			AccountID:    accountBase.ID,
			RefreshToken: out.RefreshToken,
		})
		if err != nil {
			return response.ErrInvalidToken, out, fmt.Errorf("lỗi insert key: %v", err)
		}
	}
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
	// đưa vào danh sách blacklist
	redisKey := fmt.Sprintf("TOKEN_BLACK_LIST_%s", subjectUUID)
	err = global.Rdb.Set(ctx, redisKey, 1, 0).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, fmt.Errorf("lỗi ở phần set vào redis")
	}
	//
	log.Println("User info from cache:", infoUser.ID)
	err = s.r.DeleteKey(ctx, infoUser.ID)
	return response.ErrCodeSucces, err
}
func (s *sLogin) RefreshTokens(ctx context.Context) (codeResult int, out model.LoginOutput, err error) {
	// Lấy lấy RefreshTokens để kiểm tra
	refresToken := ctx.Value("refreshToken")
	if refresToken == nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// Ép kiểu sang string
	refreshTokenStr, ok := refresToken.(string)
	log.Println("RefreshToken kiểm tra:", refreshTokenStr)
	if !ok {
		return response.ErrCodeAuthFailed, out, errors.New("invalid refresh token format")
	}
	// kiểm tra refresh trong db
	countRefreshToken, err := s.r.CountRefreshTokenByAccount(ctx, refreshTokenStr)
	if countRefreshToken == 0 {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Tài khoản chưa được đăng ký hoặc đã đăng nhập ở nơi khác vui lòng đăng nhập lại")
	}
	log.Println("countRefreshToken: ", countRefreshToken)
	// lấy Id của account
	subjectUUID := ctx.Value("subjectUUID")
	if subjectUUID == nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// Lấy thông tin user từ cache
	var infoUser model.GetCacheToken
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, out, err
	}
	// Kiểm tra trong db coi có sử dụng chưa
	getRefreshTokenUsed, err := s.r.CountByTokenAndAccount(ctx, database.CountByTokenAndAccountParams{
		AccountID:    infoUser.ID,
		JSONCONTAINS: fmt.Sprintf("[\"%s\"]", refreshTokenStr),
	})
	if getRefreshTokenUsed > 0 {
		err := s.r.DeleteKey(ctx, infoUser.ID)
		return response.ErrCodeAuthFailed, out, err
	}
	accountBase, err := s.r.GetOneAccountInfoAdmin(ctx, infoUser.Email)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Lỗi lấy thông tin tài khoản")
	}
	subToken := utils.GenerateCliTokenUUID(int(accountBase.Number))
	infoAccount, err := s.r.GetAccountById(ctx, accountBase.ID)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("lỗi ở phần lấy thông tin tài khoản")
	}
	infoAccountJson, err := json.Marshal(infoAccount)
	// check bảng keytoken có tồn tại hay chưa
	// Nếu chưa thì insert
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}
	err = global.Rdb.Set(ctx, subToken, infoAccountJson, time.Duration(consts.REFRESH_TOKEN)*time.Hour).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("lỗi ở phần set vào redis")
	}
	out.ID = accountBase.ID
	out.Email = accountBase.Email
	out.AccessToken, err = auth.CreateToken(subToken)
	out.RefreshToken, err = auth.CreateRefreshToken(subToken)
	err = s.r.UpdateRefreshTokenAndUsedTokens(ctx, database.UpdateRefreshTokenAndUsedTokensParams{
		AccountID:       accountBase.ID,
		RefreshToken:    out.RefreshToken,
		JSONARRAY:       out.RefreshToken, // Đảm bảo kiểu string
		JSONARRAYAPPEND: refreshTokenStr,  // Ép kiểu đúng khi truyền vào
	})
	if err != nil {
		return response.ErrInvalidToken, out, fmt.Errorf("lỗi update key: %v", err)

	}
	return codeResult, out, err
}

// ChangePassword
func (s *sLogin) ChangePassword(ctx context.Context, in *model.ChangePasswordInput) (codeResult int, out model.LoginOutput, err error) {
	subjectUUID := ctx.Value("subjectUUID")
	log.Println("subjectUUID:", subjectUUID)
	if subjectUUID == nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("subjectUUID not found in context")
	}
	// Lấy thông tin user từ cache
	var infoUser model.GetCacheTokenForChangePassword
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, out, err
	}
	log.Println("in.Password", in.Password)
	// lưu thông tin password vào db
	outUser, err := s.r.GetAccountById(ctx, infoUser.ID)
	Password := crypto.HashPassword(in.Password, outUser.Salt, global.Config.JWT.SECRET_KEY)
	err = s.r.ChangPasswordById(ctx, database.ChangPasswordByIdParams{
		Password: Password,
		ID:       infoUser.ID,
	})
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Lỗi update password: %v", err)
	}
	// set giá trị vào redis
	invalidationKey := fmt.Sprintf("TOKEN_IAT_AVAILABLE_%s", infoUser.ID)
	changpasswordDate := time.Now()
	err = global.Rdb.Set(ctx, invalidationKey, changpasswordDate.Unix(), time.Duration(consts.REFRESH_TOKEN)*time.Hour).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Lỗi khi set redis")
	}
	// set lại thông tin cho output
	out.ID = infoUser.ID
	out.Email = infoUser.Email
	subToken := utils.GenerateCliTokenUUID(int(infoUser.Number))
	out.AccessToken, err = auth.CreateToken(subToken)
	out.RefreshToken, err = auth.CreateRefreshToken(subToken)
	log.Println("RefreshToken")
	// kiểm tra và cập nhật keytoken
	err = s.r.UpdateRefreshToken(ctx, database.UpdateRefreshTokenParams{
		AccountID:    infoUser.ID,
		RefreshToken: out.RefreshToken,
	})
	if err != nil {
		return response.ErrInvalidToken, out, fmt.Errorf("lỗi update key: %v", err)

	}
	return response.ErrCodeSucces, out, err
}
