package impl

import (
	"context"
	"database/sql"
	"encoding/json"
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
func (s *sMenu) GetAllMenu(ctx context.Context) (codeResult int, out []model.MenuOutput, err error) {
	lstMenu, err := s.r.GetAllMenus(ctx)
	if err != nil {
		return response.ErrCodeMenuErrror, nil, err
	}
	var children []model.MenuOutput
	for _, item := range lstMenu {
		if data, ok := item.Children.([]byte); ok {
			if err := json.Unmarshal(data, &children); err != nil {
				children = []model.MenuOutput{}
			}
		}
		out = append(out, model.MenuOutput{
			Id:                item.ID,
			Menu_name:         item.MenuName,
			Menu_icon:         item.MenuIcon,
			Menu_url:          item.MenuUrl,
			Menu_Number_order: item.MenuNumberOrder,
			Menu_parent_id:    item.MenuParentID.String,
			Menu_level:        int(item.MenuLevel),
			Children:          children,
		})
	}
	return response.ErrCodeSucces, out, err
}

func (s *sMenu) GetMenuById(ctx context.Context, id string) (codeResult int, out model.MenuOutput, err error) {
	// Lấy menu từ repository
	menubyid, err := s.r.GetMenuById(ctx, id)
	if err != nil {
		return response.ErrCodeMenuErrror, model.MenuOutput{}, err
	}
	var children []model.MenuOutput
	if menubyid.Children != nil {
		if data, ok := menubyid.Children.([]byte); ok {
			if err := json.Unmarshal(data, &children); err != nil {
				children = []model.MenuOutput{}
			}
		}
	}
	out = model.MenuOutput{
		Id:                menubyid.ID,
		Menu_name:         menubyid.MenuName,
		Menu_icon:         menubyid.MenuIcon,
		Menu_url:          menubyid.MenuUrl,
		Menu_Number_order: menubyid.MenuNumberOrder,
		Menu_parent_id:    menubyid.MenuParentID.String,
		Menu_level:        int(menubyid.MenuLevel),
		Children:          children,
	}

	return response.ErrCodeSucces, out, err
}

func (s *sMenu) EditMenuById(ctx context.Context, in *model.MenuInput, id string) (codeResult int, out model.MenuOutput, err error) {
	// Thực hiện cập nhật menu
	err = s.r.EditMenuById(ctx, database.EditMenuByIdParams{
		ID:           id,
		MenuName:     in.Menu_name,
		MenuIcon:     in.Menu_icon,
		MenuUrl:      in.Menu_url,
		MenuParentID: sql.NullString{String: in.Menu_parent_id, Valid: in.Menu_parent_id != ""},
	})

	if err != nil {
		return response.ErrCodeMenuErrror, out, err
	}

	// Lấy lại dữ liệu menu sau khi cập nhật
	menu, err := s.r.GetMenuById(ctx, id)
	if err != nil {
		return response.ErrCodeMenuErrror, out, err
	}

	// Gán dữ liệu vào output
	out = model.MenuOutput{
		Id:                menu.ID,
		Menu_name:         menu.MenuName,
		Menu_icon:         menu.MenuIcon,
		Menu_url:          menu.MenuUrl,
		Menu_parent_id:    menu.MenuParentID.String,
		Menu_Number_order: menu.MenuNumberOrder,
		Menu_level:        int(menu.MenuLevel),
	}

	return response.ErrCodeSucces, out, err
}
