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
		log.Println("ACCOUNTID: ", infoUser.ID)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("role id not found: %w", err)
		}
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
		// kiểm tra số lượng max number CỦA CHA
		maxNumberParents, err := s.r.GetRoleById(ctx, RoleId.RoleID)
		log.Println("ROLEID: ", RoleId.ID)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get max number parents: %w", err)
		}
		// kiểm tra số lượng đã tạo và số lượng giới hạn của role
		totalAccounts, err := s.r.GetTotalAccounts(ctx, maxNumberParents.CreatedBy)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get total accounts: %w", err)
		}
		// Chuyển đổi RoleMaxNumber từ string sang int

		// Đặt giá trị cho node mới
		leftValue = parentRole.RoleRightValue
		rightValue = parentRole.RoleRightValue + 1
		log.Printf("Type of totalAccounts.Totalaccount: %T, Value: %v", totalAccounts.Totalaccount, totalAccounts.Totalaccount)
		if totalAccounts.Totalaccount.(int64) < int64(maxNumberParents.RoleMaxNumber) {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("Func package maximum account number: %w", err)
		}
	}

	// Tạo role mới
	_, err = s.r.CreateRole(ctx, database.CreateRoleParams{
		ID:             newID,
		Code:           in.Code,
		RoleName:       in.Role_name,
		RoleLeftValue:  leftValue,
		RoleRightValue: rightValue,
		RoleMaxNumber:  int32(in.Role_max_number),
		CreatedBy:      in.Created_by,
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
