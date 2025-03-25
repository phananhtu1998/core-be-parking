package impl

import (
	"context"
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

type sAccount struct {
	r *database.Queries
}

func NewAccountImpl(r *database.Queries) *sAccount {
	return &sAccount{r: r}
}

// Tạo tài khoản mới
func (s *sAccount) CreateAccount(ctx context.Context, in *model.AccountInput) (codeResult int, out model.AccountOutput, err error) {
	// TODO: check Email
	accountFound, err := s.r.CheckAccountBaseExists(ctx, in.Email)
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
	accountBase.Password = crypto.HashPassword(in.Password, userSalt, global.Config.JWT.SECRET_KEY)
	rand.Seed(time.Now().UnixNano())
	newUUID := uuid.New().String()
	_, err = s.r.InsertAccount(ctx, database.InsertAccountParams{
		ID:       newUUID,
		Number:   rand.Int31(),
		Name:     in.Name,
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
	accountOutput := model.AccountOutput{
		ID:     newUUID,
		Name:   in.Name,
		Email:  in.Email,
		Status: in.Status,
		Images: in.Images,
	}
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
		ID:     accountItem.ID,
		Name:   accountItem.Name,
		Email:  accountItem.Email,
		Status: accountItem.Status,
		Images: accountItem.Images,
	}, nil
}

// Cập nhật tài khoản
func (s *sAccount) UpdateAccount(ctx context.Context, in *model.AccountInput, id string) (codeResult int, out model.AccountOutput, err error) {
	err = s.r.EditAccountById(ctx, database.EditAccountByIdParams{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
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
	accountOutput := model.AccountOutput{
		ID:     updatedAccount.ID,
		Name:   updatedAccount.Name,
		Email:  updatedAccount.Email,
		Status: updatedAccount.Status,
		Images: updatedAccount.Images,
	}
	return response.ErrCodeSucces, accountOutput, nil
}

// Xóa tài khoản
func (s *sAccount) DeleteAccount(ctx context.Context, id string) (codeResult int, err error) {
	// TODO: Thêm logic xóa tài khoản
	err = s.r.DeleteAccountById(ctx, id)
	if err != nil {
		return response.ErrInvalidOTP, err
	}
	return response.ErrCodeSucces, err
}

// Lấy danh sách tất cả tài khoản
func (s *sAccount) GetAllAccount(ctx context.Context) (codeResult int, out []model.AccountOutput, err error) {
	lst, err := s.r.GetAllAccounts(ctx)
	if err != nil {
		return response.ErrCodeAuthFailed, nil, err
	}
	for _, item := range lst {
		out = append(out, model.AccountOutput{
			ID:     item.ID,
			Name:   item.Name,
			Email:  item.Email,
			Status: item.Status,
			Images: item.Images,
		})
	}

	log.Println("Successfully fetched accounts:", len(out))
	return response.ErrCodeSucces, out, nil
}
