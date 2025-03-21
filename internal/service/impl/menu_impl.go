package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/repo"
	"go-backend-api/internal/utils"
	"go-backend-api/pkg/response"
	"log"

	"github.com/google/uuid"
)

type sMenu struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewMenuImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) *sMenu {
	return &sMenu{
		r:   r,
		qTx: qTx,
		db:  db,
	}
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
		MenuParentID:    sql.NullString{String: in.Menu_parent_id, Valid: false},
		MenuLevel:       int32(in.Menu_level),
		MenuNumberOrder: int32(in.Menu_Number_order),
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
func (s *sMenu) GetAllMenu(ctx context.Context) (int, []model.MenuOutput, error) {
	// Lấy danh sách menu từ database
	lstMenu, err := s.r.GetAllMenus(ctx)
	if err != nil {
		return response.ErrCodeMenuErrror, nil, err
	}
	// Nhóm menu theo parent_id và lấy danh sách menu gốc
	// Convert lstMenu to []model.MenuOutput
	var menuOutputs []model.MenuOutput
	for _, menu := range lstMenu {
		menuOutputs = append(menuOutputs, model.MenuOutput{
			Id:                menu.ID,
			Menu_name:         menu.MenuName,
			Menu_icon:         menu.MenuIcon,
			Menu_url:          menu.MenuUrl,
			Menu_parent_id:    menu.MenuParentID.String,
			Menu_level:        int(menu.MenuLevel),
			Menu_Number_order: int(menu.MenuNumberOrder),
			Menu_group_name:   menu.MenuGroupName,
		})
	}

	menuMap, rootMenus := repo.GroupMenusByParent(menuOutputs)

	// Xây dựng cây menu từ danh sách đã nhóm
	finalMenu := repo.BuildMenuTree(rootMenus, menuMap)

	return response.ErrCodeSucces, finalMenu, nil
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
		Menu_Number_order: int(menubyid.MenuNumberOrder),
		Menu_parent_id:    menubyid.MenuParentID.String,
		Menu_level:        int(menubyid.MenuLevel),
		Children:          children,
	}

	return response.ErrCodeSucces, out, err
}

func (s *sMenu) EditMenuById(ctx context.Context, menuUpdates []model.MenuInput) (int, model.MenuOutput, error) {
	// Bắt đầu transaction mới
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeMenuErrror, model.MenuOutput{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var committed bool
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	// Lấy thông tin tất cả menu (sử dụng transaction)
	allMenus, err := s.r.GetAllMenus(ctx)
	if err != nil {
		return response.ErrCodeMenuErrror, model.MenuOutput{}, fmt.Errorf("failed to get all menus: %w", err)
	}

	// Xử lý logic cập nhật menu
	updateParamsList, err := utils.ProcessMenuUpdates(menuUpdates, allMenus)
	if err != nil {
		return response.ErrCodeMenuErrror, model.MenuOutput{}, err
	}

	// Thực hiện các truy vấn cập nhật trong transaction
	for _, updateParams := range updateParamsList {
		err := s.r.UpdateSingleMenu(ctx, updateParams)
		if err != nil {
			log.Printf("Lỗi cập nhật menu ID %s: %v", updateParams.ID, err)
			return response.ErrCodeMenuErrror, model.MenuOutput{}, fmt.Errorf("failed to update menu ID %s: %w", updateParams.ID, err)
		}
	}

	// Commit transaction sau khi cập nhật thành công
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeMenuErrror, model.MenuOutput{}, fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true

	// Lấy thông tin menu cuối cùng được cập nhật
	var lastUpdatedMenu model.MenuOutput
	if len(menuUpdates) > 0 {
		lastMenuInput := menuUpdates[len(menuUpdates)-1]
		updatedMenu, err := s.r.GetMenuById(ctx, lastMenuInput.Id)
		if err != nil {
			return response.ErrCodeMenuErrror, model.MenuOutput{}, fmt.Errorf("failed to get updated menu: %w", err)
		}

		lastUpdatedMenu = model.MenuOutput{
			Id:                updatedMenu.ID,
			Menu_name:         updatedMenu.MenuName,
			Menu_icon:         updatedMenu.MenuIcon,
			Menu_url:          updatedMenu.MenuUrl,
			Menu_parent_id:    updatedMenu.MenuParentID.String,
			Menu_level:        int(updatedMenu.MenuLevel),
			Menu_Number_order: int(updatedMenu.MenuNumberOrder),
			Menu_group_name:   updatedMenu.MenuGroupName,
		}
	}

	return response.ErrCodeSucces, lastUpdatedMenu, nil
}
