package impl

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/internal/utils/cache"
	"go-backend-api/pkg/response"
	"log"

	"github.com/google/uuid"
)

type sFuncpackage struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewFuncpackageImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) service.Ifuncpackage {
	return &sFuncpackage{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sFuncpackage) CreateFuncPackage(ctx context.Context, in *model.FuncpackageInput) (codeResult int, out model.FuncpackageOutput, err error) {
	var leftValue, rightValue int32
	newID := uuid.New().String()
	// Bắt đầu transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()
	subjectUUID := ctx.Value("subjectUUID")
	println("subjectUUID account: ", subjectUUID)
	var infoUser model.GetCacheToken
	// Lấy Id tài khoản đang đăng nhập từ context
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, out, err
	}
	// Nếu created_by trống, đây là node gốc

	// Lấy thông tin của node cha
	RoleId, err := s.r.GetOneRoleAccountByAccountId(ctx, infoUser.ID)
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("role id not found: %w", err)
	}
	parentRole, err := s.r.GetParentRoleInfo(ctx, RoleId.RoleID)
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("parent role not found: %w", err)
	}

	// Lấy giá trị role max number của tài khoản hiện tại
	rolemaxnumber, err := s.r.GetRoleById(ctx, RoleId.RoleID)
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("failed to get role max number: %w", err)
	}
	log.Println("RoleId: ", RoleId.RoleID)
	log.Println("rolemaxnumber: ", rolemaxnumber.RoleMaxNumber)
	// Lấy tổng giá trị mà được phép tạo của tài khoản hiện tại
	log.Println("infoUser: ", infoUser.ID)
	summaxnumber, err := s.r.GetTotalAccounts(ctx, database.GetTotalAccountsParams{
		CreatedBy: infoUser.ID,
		Column2:   infoUser.ID,
		Column3:   infoUser.ID,
	})
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("Lỗi khi lấy tổng số tài khoản được phép tạo: %w", err)
	}
	summaxnumberInt64 := summaxnumber.(int64)
	log.Println("Tổng:", summaxnumberInt64+1+int64(in.Role_max_number))
	accountCreated, err := s.r.GetAccountCreated(ctx, infoUser.ID)
	log.Println("accountCreated: ", accountCreated)
	if int64(rolemaxnumber.RoleMaxNumber) <= (summaxnumberInt64 + int64(in.Role_max_number) + accountCreated) {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("Số lượng tài khoản của gói này đã vượt quá số lượng cho phép")
	}

	// Cập nhật right values
	err = s.r.UpdateRightValuesForInsert(ctx, parentRole.RoleRightValue)
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("failed to update right values: %w", err)
	}

	// Cập nhật left values
	err = s.r.UpdateLeftValuesForInsert(ctx, parentRole.RoleRightValue)
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("failed to update left values: %w", err)
	}

	leftValue = parentRole.RoleRightValue
	rightValue = parentRole.RoleRightValue + 1

	// Tạo role mới
	_, err = s.r.CreateRole(ctx, database.CreateRoleParams{
		ID:             newID,
		Code:           in.Code,
		RoleName:       in.Role_name,
		RoleLeftValue:  leftValue,
		RoleRightValue: rightValue,
		RoleMaxNumber:  int32(in.Role_max_number),
		LicenseID:      "",
		CreatedBy:      infoUser.ID,
	})
	if err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("failed to create role: %w", err)
	}
	// Commit transaction
	if err = tx.Commit(); err != nil {
		return response.ErrCodeRoleError, model.FuncpackageOutput{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return response.ErrCodeSucces, model.FuncpackageOutput{
		Id: newID,
		FuncpackageInput: model.FuncpackageInput{
			Code:            in.Code,
			Role_name:       in.Role_name,
			Role_max_number: in.Role_max_number,
		},
	}, nil
}
func (s *sFuncpackage) GetAlFuncPackageByCreatedBy(ctx context.Context) (codeResult int, out []model.FuncpackageOutput, err error) {
	subjectUUID := ctx.Value("subjectUUID")
	var infoUser model.GetCacheToken
	// Lấy Id tài khoản đang đăng nhập từ context
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, out, err
	}
	println("infoUser.ID: ", infoUser.ID)
	lstFuncpackage, err := s.r.GetAllFuncPackageByCreatedBy(ctx, infoUser.ID)
	if err != nil {
		return response.ErrCodeRoleError, out, fmt.Errorf("Lỗi khi lấy danh sách gói chức năng")
	}
	for _, item := range lstFuncpackage {
		out = append(out, model.FuncpackageOutput{
			Id: item.ID,
			FuncpackageInput: model.FuncpackageInput{
				Code:            item.Code,
				Role_name:       item.RoleName,
				Role_max_number: int(item.RoleMaxNumber),
			},
		})
	}
	return response.ErrCodeSucces, out, nil
}
func (s *sFuncpackage) DeleteFuncPackage(ctx context.Context, id string) (codeResult int, err error) {
	// Bắt đầu transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeRoleError, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()
	// Kiểm tra có gói con hay không nếu có không được xóa
	// Kiểm tra có tài khoản đang sử dụng hay không

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return response.ErrCodeRoleError, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return response.ErrCodeSucces, nil
}
