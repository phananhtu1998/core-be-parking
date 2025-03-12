package impl

import (
	"context"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/pkg/response"
	"log"
)

type sAccount struct {
	r *database.Queries
}

func NewAccountImpl(r *database.Queries) *sAccount {
	return &sAccount{r: r}
}

// Tạo tài khoản mới
func (s *sAccount) CreateAccount(ctx context.Context, in *model.AccountInput) (codeResult int, err error) {
	// TODO: Thêm logic tạo tài khoản
	return response.ErrCodeSucces, nil
}

// Lấy thông tin tài khoản theo ID
func (s *sAccount) GetAccountById(ctx context.Context, id string) (codeResult int, out model.AccountOutput, err error) {
	log.Println("CALL service Get Account By Id ...")
	accountItem, err := s.r.GetAccountById(ctx, id)
	if err != nil {
		return response.ErrCodeOtpNotExists, out, err
	}
	return response.ErrCodeSucces, model.AccountOutput{
		ID:     accountItem.ID,
		Name:   accountItem.Name,
		Email:  accountItem.Email,
		Status: accountItem.Status,
		Images: accountItem.Images,
	}, nil
}

// Cập nhật tài khoản
func (s *sAccount) UpdateAccount(ctx context.Context, in *model.AccountInput) (codeResult int, out model.AccountOutput, err error) {
	//err = s.r.EditAccountById(ctx, in)
	return response.ErrCodeSucces, model.AccountOutput{}, nil
}

// Xóa tài khoản
func (s *sAccount) DeleteAccount(ctx context.Context, id string) (codeResult int, err error) {
	// TODO: Thêm logic xóa tài khoản
	return response.ErrCodeSucces, nil
}

// Lấy danh sách tất cả tài khoản
func (s *sAccount) GetAllAccount(ctx context.Context) (codeResult int, out []model.AccountOutput, err error) {
	log.Println("Starting GetAllAccount...") // Log để kiểm tra

	lst, err := s.r.GetAllAccounts(ctx)
	if err != nil {
		log.Println("Error fetching accounts:", err)
		return response.ErrCodeAuthFailed, nil, err
	}

	var accounts []model.AccountOutput
	for _, item := range lst {
		accounts = append(accounts, model.AccountOutput{
			ID:     item.ID,
			Name:   item.Name,
			Email:  item.Email,
			Status: item.Status,
			Images: item.Images,
		})
	}

	log.Println("Successfully fetched accounts:", len(accounts))
	return response.ErrCodeSucces, accounts, nil
}
