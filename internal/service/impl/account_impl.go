package impl

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-api/global"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/cache"
	"go-backend-api/internal/utils/crypto"
	"go-backend-api/pkg/response"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type sAccount struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewAccountImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) *sAccount {
	return &sAccount{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

// Tạo tài khoản mới
func (s *sAccount) CreateAccount(ctx context.Context, in *model.AccountInput) (codeResult int, out model.AccountOutput, err error) {
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
	subjectUUID := ctx.Value("subjectUUID")
	println("subjectUUID account: ", subjectUUID)
	var infoUser model.GetCacheToken
	// Lấy Id tài khoản đang đăng nhập từ context
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, out, err
	}
	// TODO: check Email and username
	accountFound, err := s.r.CheckAccountBaseExists(ctx, database.CheckAccountBaseExistsParams{
		Email:    in.Email,
		Username: in.UserName,
	})
	if err != nil {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, err
	}
	if accountFound > 0 {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, fmt.Errorf("Email has already registered")
	}
	// TODO: hash Password
	accountBase := database.Account{}
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, model.AccountOutput{}, err
	}
	// Kiểm tra role_max_number
	countRoleMaxNumber, err := s.r.GetRoleById(ctx, in.RoleId)
	if err != nil {
		return response.ErrCodeRoleNotFound, model.AccountOutput{}, err
	}
	//Kiểm tra số lượng tài khoản theo role count(account_id) theo role_id
	countAccountId, err := s.r.CheckCountRoleId(ctx, in.RoleId)
	if err != nil {
		return response.ErrCodeRoleNotFound, model.AccountOutput{}, err
	}
	// Kiểm tra tài khoản đã được
	if int64(countRoleMaxNumber.RoleMaxNumber) < (countAccountId + 1) {
		return response.ErrCodeRoleAccountMaxNumber, model.AccountOutput{}, fmt.Errorf("Số lượng tài khoản tạo đã vượt quá số lượng quy định")
	}
	//
	accountBase.Password = crypto.HashPassword(global.Config.JWT.PASSWORD, userSalt, global.Config.JWT.SECRET_KEY)
	rand.Seed(time.Now().UnixNano())
	newUUID := uuid.New().String()
	_, err = s.r.InsertAccount(ctx, database.InsertAccountParams{
		ID:        newUUID,
		Number:    rand.Int31(),
		Name:      in.Name,
		Username:  in.UserName,
		Email:     in.Email,
		Password:  accountBase.Password,
		Salt:      userSalt,
		Status:    in.Status,
		Images:    in.Images,
		CreatedBy: infoUser.ID,
	})
	if err != nil {
		log.Printf("Lỗi khi chèn tài khoản: %v", err)
		return response.ErrCodeParamInvalid, model.AccountOutput{}, err
	}

	// nếu license tồn tại thì thêm vào bảng role account
	// thêm vào bảng role account
	err = s.r.CreateRoleAccount(ctx, database.CreateRoleAccountParams{
		ID:        newUUID,
		AccountID: newUUID,
		RoleID:    in.RoleId,
		CreatedBy: sql.NullString{infoUser.ID, true},
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

// Lấy thông tin tài khoản theo ID
func (s *sAccount) GetAccountById(ctx context.Context, id string) (codeResult int, out model.AccountOutput, err error) {
	log.Println("CALL service Get Account By Id ...")
	accountItem, err := s.r.GetAccountById(ctx, id)
	if err != nil {
		return response.ErrCodeOtpNotExists, out, err
	}
	// truyền password đã hash trong db,password input, salt trong db
	checkpass := crypto.MatchingPassword("fca10a2c4d80b0151fd49bf277ee1447d2d67f2ddf0b0066a174833fc92f4f7f", "123", "e404cc8042ede7884b7d9464ad262221")
	//log.Println("hashpass: ", hashpass)
	log.Println("checkpass: ", checkpass)
	return response.ErrCodeSucces, model.AccountOutput{
		ID:       accountItem.ID,
		UserName: accountItem.Username,
		Name:     accountItem.Name,
		Email:    accountItem.Email,
		Status:   accountItem.Status,
		Images:   accountItem.Images,
	}, nil
}

// Cập nhật tài khoản
func (s *sAccount) UpdateAccount(ctx context.Context, in *model.AccountInput, id string) (codeResult int, out model.AccountOutput, err error) {
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
	// Kiểm tra username and mail
	accountFound, err := s.r.CheckAccountBaseExists(ctx, database.CheckAccountBaseExistsParams{
		Email:    in.Email,
		Username: in.UserName,
	})
	if err != nil {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, err
	}
	if accountFound > 0 {
		return response.ErrCodeUserHasExists, model.AccountOutput{}, fmt.Errorf("Username or Email has already registered")
	}
	err = s.r.EditAccountById(ctx, database.EditAccountByIdParams{
		Name:     in.Name,
		Username: in.UserName,
		Email:    in.Email,
		Status:   in.Status,
		Images:   in.Images,
		ID:       id,
	})

	if err != nil {
		return response.ErrInvalidOTP, model.AccountOutput{}, err
	}
	updatedAccount, err := s.r.GetAccountById(ctx, id)
	if err != nil {
		return response.ErrCodeParamInvalid, model.AccountOutput{}, err
	}
	// Update account funcpackage
	err = s.r.UpdateRoleAccountByAccountId(ctx, database.UpdateRoleAccountByAccountIdParams{
		RoleID:    in.RoleId,
		AccountID: id,
	})

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeMenuErrror, out, err
	}
	committed = true
	accountOutput := model.AccountOutput{
		ID:       updatedAccount.ID,
		Name:     updatedAccount.Name,
		UserName: updatedAccount.Username,
		Email:    updatedAccount.Email,
		Status:   updatedAccount.Status,
		Images:   updatedAccount.Images,
	}
	return response.ErrCodeSucces, accountOutput, nil
}

// Xóa tài khoản
func (s *sAccount) DeleteAccount(ctx context.Context, id string) (codeResult int, err error) {
	// Khởi tạo transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeMenuErrror, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var committed bool
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()
	// Update is_deleted trong bảng account
	err = s.r.DeleteAccountById(ctx, id)
	if err != nil {
		return response.ErrInvalidOTP, err
	}
	//Update is_deleted trong bảng role_account
	err = s.r.DeleteRoleAccountByAccountId(ctx, id)
	if err != nil {
		return response.ErrInvalidOTP, err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeMenuErrror, err
	}
	committed = true
	return response.ErrCodeSucces, err
}

// Lấy danh sách tất cả tài khoản
func (s *sAccount) GetAllAccount(ctx context.Context) (codeResult int, out []model.AccountOutput, err error) {
	subjectUUID := ctx.Value("subjectUUID")
	println("subjectUUID account: ", subjectUUID)
	var infoUser model.GetCacheToken
	// Lấy Id tài khoản đang đăng nhập từ context
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, out, err
	}
	// Lây
	lst, err := s.r.GetAllAccountByCreatedBy(ctx, infoUser.ID)
	if err != nil {
		return response.ErrCodeAuthFailed, nil, err
	}
	for _, item := range lst {
		out = append(out, model.AccountOutput{
			ID:       item.ID,
			Name:     item.Name,
			UserName: item.Username,
			Email:    item.Email,
			Status:   item.Status,
			Images:   item.Images,
		})
	}

	log.Println("Successfully fetched accounts:", len(out))
	return response.ErrCodeSucces, out, nil
}
