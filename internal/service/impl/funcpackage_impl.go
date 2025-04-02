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
	"time"

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

func (s *sFuncpackage) CreateFuncPackage(ctx context.Context, in *model.Role) (codeResult int, out model.Role, err error) {
	var leftValue, rightValue int32
	newID := uuid.New().String()
	// Bắt đầu transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()
	subjectUUID := ctx.Value("subjectUUID")
	println("subjectUUID account: ", subjectUUID)
	var infoUser model.GetCacheToken
	var role_id string
	// Lấy Id tài khoản đang đăng nhập từ context
	if err := cache.GetCache(ctx, subjectUUID.(string), &infoUser); err != nil {
		return 0, out, err
	}
	// Nếu created_by trống, đây là node gốc
	if in.Created_by == "" {
		// Lấy giá trị right lớn nhất
		maxRightValue, err := s.r.GetMaxRightValue(ctx)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get max right value: %w", err)
		}

		// Đặt node mới là root - Sửa phần này
		maxRightValueInt64 := maxRightValue.(int64)
		leftValue = int32(maxRightValueInt64) + 1
		rightValue = int32(maxRightValueInt64) + 2
	} else {
		// Lấy thông tin của node cha
		RoleId, err := s.r.GetOneRoleAccountByAccountId(ctx, infoUser.ID)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("role id not found: %w", err)
		}
		role_id = RoleId.RoleID
		parentRole, err := s.r.GetParentRoleInfo(ctx, RoleId.RoleID)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("parent role not found: %w", err)
		}
		// Cập nhật right values
		err = s.r.UpdateRightValuesForInsert(ctx, parentRole.RoleRightValue)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to update right values: %w", err)
		}

		// Cập nhật left values
		err = s.r.UpdateLeftValuesForInsert(ctx, parentRole.RoleRightValue)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to update left values: %w", err)
		}
		// Lấy giá trị role max number của tài khoản hiện tại
		rolemaxnumber, err := s.r.GetRoleById(ctx, RoleId.RoleID)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get role max number: %w", err)
		}
		log.Println("rolemaxnumber: ", rolemaxnumber.RoleMaxNumber)
		// Lấy tổng giá trị mà được phép tạo của tài khoản hiện tại
		log.Println("infoUser: ", infoUser.ID)
		summaxnumber, err := s.r.GetTotalAccounts(ctx, infoUser.ID)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get total accounts: %w", err)
		}
		if summaxnumber.(int64) < int64(rolemaxnumber.RoleMaxNumber) {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get total accounts: %w", err)
		}
		// log.Println("summaxnumber: ", summaxnumber)
		// Đặt giá trị cho node mới
		leftValue = parentRole.RoleRightValue
		rightValue = parentRole.RoleRightValue + 1
	}

	// Tạo role mới
	_, err = s.r.CreateRole(ctx, database.CreateRoleParams{
		ID:             newID,
		Code:           in.Code,
		RoleName:       in.Role_name,
		RoleLeftValue:  leftValue,
		RoleRightValue: rightValue,
		RoleMaxNumber:  int32(in.Role_max_number),
		CreatedBy:      role_id,
	})
	if err != nil {
		return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to create role: %w", err)
	}
	// Commit transaction
	if err = tx.Commit(); err != nil {
		return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return response.ErrCodeSucces, model.Role{
		Id:               newID,
		Code:             in.Code,
		Role_name:        in.Role_name,
		Role_left_value:  int(leftValue),
		Role_right_value: int(rightValue),
		Role_max_number:  in.Role_max_number,
		Is_licensed:      in.Is_licensed,
		Created_by:       in.Created_by,
		Created_at:       time.Now(),
	}, nil
}
