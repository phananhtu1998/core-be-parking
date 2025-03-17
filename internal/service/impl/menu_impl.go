package impl

import (
	"context"
	"database/sql"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/pkg/response"

	"github.com/google/uuid"
)

type sMenu struct {
	r *database.Queries
}

func NewMenuImpl(r *database.Queries) *sMenu {
	return &sMenu{r: r}
}

func (s *sMenu) CreateMenu(ctx context.Context, in *model.MenuInput) (int, model.MenuOutput, error) {
	// Kiểm tra url có tồn tại trong DB
	if urlFound, err := s.r.CountMenuByURL(ctx, in.Menu_url); err != nil {
		return response.ErrCodeMenuErrror, model.MenuOutput{}, err
	} else if urlFound > 0 {
		return response.ErrCodeMenuHasExists, model.MenuOutput{}, nil
	}

	// Tạo ID mới cho menu
	newUUID := uuid.New().String()

	// Thêm menu vào DB
	if _, err := s.r.InsertMenu(ctx, database.InsertMenuParams{
		ID:              newUUID,
		MenuName:        in.Menu_name,
		MenuIcon:        in.Menu_icon,
		MenuUrl:         in.Menu_url,
		MenuParentID:    sql.NullString{String: in.Menu_parent_id, Valid: true},
		MenuLevel:       int32(in.Menu_level),
		MenuNumberOrder: in.Menu_Number_order,
		MenuGroupName:   in.Menu_group_name,
	}); err != nil {
		return response.ErrCodeMenuErrror, model.MenuOutput{}, err
	}

	// Tạo output từ input
	output := model.MenuOutput{
		Id:                newUUID, // Sử dụng ID mới
		Menu_name:         in.Menu_name,
		Menu_icon:         in.Menu_icon,
		Menu_url:          in.Menu_url,
		Menu_parent_id:    in.Menu_parent_id,
		Menu_level:        in.Menu_level,
		Menu_Number_order: in.Menu_Number_order,
		Menu_group_name:   in.Menu_group_name,
	}

	// Trả về kết quả thành công
	return response.ErrCodeSucces, output, nil
}
