package impl

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-api/global"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/crypto"
	"go-backend-api/pkg/response"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type sUser struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewUserImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) *sUser {
	return &sUser{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sUser) CreateUser(ctx context.Context, in *model.AccountInput) (codeResult int, out model.AccountOutput, err error) {
	// Khởi tạo transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeMenuErrror, out, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var committed bool
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()
	// Kiểm tra user và email có tồn tại hay chưa
	accountFound, err := s.r.CheckAccountBaseExists(ctx, in.Email)
	if err != nil {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, err
	}
	if accountFound > 0 {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, fmt.Errorf("Email has already registered")
	}
	accountFound, err = s.r.CheckAccountBaseExists(ctx, in.UserName)
	if err != nil {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, err
	}
	if accountFound > 0 {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, fmt.Errorf("Username has already registered")
	}
	accountBase := database.Account{}
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, model.AccountOutput{}, err
	}
	accountBase.Password = crypto.HashPassword(global.Config.JWT.PASSWORD, userSalt, global.Config.JWT.SECRET_KEY)
	rand.Seed(time.Now().UnixNano())
	newUUID := uuid.New().String()
	_, err = s.r.InsertAccount(ctx, database.InsertAccountParams{
		ID:       newUUID,
		Number:   rand.Int31(),
		Name:     in.Name,
		Username: in.UserName,
		Email:    in.Email,
		Password: accountBase.Password,
		Salt:     userSalt,
		Status:   in.Status,
		Images:   in.Images,
	})
	if err != nil {
		log.Printf("Lỗi khi chèn tài khoản: %v", err)
		return response.ErrCodeParamInvalid, model.AccountOutput{}, err
	}
	// thêm vào bảng role account
	err = s.r.CreateRoleAccount(ctx, database.CreateRoleAccountParams{
		ID:        newUUID,
		AccountID: newUUID,
		RoleID:    in.RoleId,
	})
	if err != nil {
		log.Printf("Lỗi khi chèn tài khoản vào bảng role account: %v", err)
		return response.ErrCodeParamInvalid, model.AccountOutput{}, err
	}
	accountOutput := model.AccountOutput{
		ID:       newUUID,
		Name:     in.Name,
		Email:    in.Email,
		UserName: in.UserName,
		Status:   in.Status,
		Images:   in.Images,
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeMenuErrror, out, err
	}
	committed = true
	return response.ErrCodeSucces, accountOutput, err
}
