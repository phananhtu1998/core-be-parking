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
		MenuParentID:    sql.NullString{String: in.Menu_parent_id, Valid: true},
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

func (s *sMenu) EditMenuById(ctx context.Context, menuUpdates []model.MenuInput) (int, []model.MenuOutput, error) {
	// Bắt đầu transaction mới
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeMenuErrror, nil, fmt.Errorf("failed to begin transaction: %w", err)
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
		return response.ErrCodeMenuErrror, nil, fmt.Errorf("failed to get all menus: %w", err)
	}

	// Xử lý logic cập nhật menu
	updateParamsList, err := utils.ProcessMenuUpdates(menuUpdates, allMenus)
	if err != nil {
		return response.ErrCodeMenuErrror, nil, err
	}

	// Thực hiện các truy vấn cập nhật trong transaction
	for _, updateParams := range updateParamsList {
		err := s.r.UpdateSingleMenu(ctx, updateParams)
		if err != nil {
			log.Printf("Lỗi cập nhật menu ID %s: %v", updateParams.ID, err)
			return response.ErrCodeMenuErrror, nil, fmt.Errorf("failed to update menu ID %s: %w", updateParams.ID, err)
		}
	}

	// Commit transaction sau khi cập nhật thành công
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeMenuErrror, nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true

	// Lấy lại tất cả menu sau khi cập nhật
	allMenus, err = s.r.GetAllMenus(ctx)
	if err != nil {
		return response.ErrCodeMenuErrror, nil, fmt.Errorf("failed to get all menus: %w", err)
	}

	// Convert to MenuOutput format
	var menuOutputs []model.MenuOutput
	for _, menu := range allMenus {
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

	// Xây dựng cây menu
	menuMap, rootMenus := repo.GroupMenusByParent(menuOutputs)
	finalMenus := repo.BuildMenuTree(rootMenus, menuMap)

	return response.ErrCodeSucces, finalMenus, nil
}

func (s *sMenu) DeleteMenu(ctx context.Context, id string) (int, error) {
	// Bắt đầu transaction mới
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

	// Lấy thông tin menu cần xóa
	_, err = s.r.GetMenuById(ctx, id)
	if err != nil {
		return response.ErrCodeMenuErrror, fmt.Errorf("menu not found: %w", err)
	}

	// Lấy tất cả menu để kiểm tra menu cha-con
	allMenus, err := s.r.GetAllMenus(ctx)
	if err != nil {
		return response.ErrCodeMenuErrror, fmt.Errorf("failed to get all menus: %w", err)
	}

	// Tạo map để lưu thông tin menu hiện tại
	menuMap := make(map[string]database.GetAllMenusRow)
	for _, m := range allMenus {
		menuMap[m.ID] = m
	}

	// Kiểm tra xem menu có phải là menu cha không
	isParent := false
	for _, m := range allMenus {
		if m.MenuParentID.String == id {
			isParent = true
			break
		}
	}

	if isParent {
		// Nếu là menu cha, xóa tất cả menu con
		for _, m := range allMenus {
			if m.MenuParentID.String == id {
				// Xóa menu con
				err := s.r.DeleteMenu(ctx, m.ID)
				if err != nil {
					return response.ErrCodeMenuErrror, fmt.Errorf("failed to delete child menu %s: %w", m.ID, err)
				}
			}
		}
		// Xóa menu cha
		err := s.r.DeleteMenu(ctx, id)
		if err != nil {
			return response.ErrCodeMenuErrror, fmt.Errorf("failed to delete parent menu: %w", err)
		}
	} else {
		// Nếu là menu con, chỉ cập nhật Is_deleted
		err := s.r.UpdateMenuDeleted(ctx, id)
		if err != nil {
			return response.ErrCodeMenuErrror, fmt.Errorf("failed to update menu deleted status: %w", err)
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeMenuErrror, fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true

	return response.ErrCodeSucces, nil
}
