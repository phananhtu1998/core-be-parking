package utils

import (
	"database/sql"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"sort"
)

// ProcessMenuUpdates xử lý logic cập nhật menu và trả về danh sách các tham số cập nhật
func ProcessMenuUpdates(menuUpdates []model.MenuInput, allMenus []database.GetAllMenusRow) ([]database.UpdateSingleMenuParams, error) {
	if len(menuUpdates) == 0 {
		return nil, fmt.Errorf("no menu updates provided")
	}

	// Tạo map để lưu thông tin menu hiện tại
	menuMap := make(map[string]database.GetAllMenusRow)
	for _, m := range allMenus {
		menuMap[m.ID] = m
	}

	// Danh sách các menu cần cập nhật
	updateParamsList := []database.UpdateSingleMenuParams{}

	for _, menuInput := range menuUpdates {
		currentMenu, exists := menuMap[menuInput.Id]
		if !exists {
			return nil, fmt.Errorf("menu with ID %s not found", menuInput.Id)
		}

		updateParams := database.UpdateSingleMenuParams{
			MenuName:        menuInput.Menu_name,
			MenuIcon:        menuInput.Menu_icon,
			MenuUrl:         menuInput.Menu_url,
			MenuParentID:    sql.NullString{String: menuInput.Menu_parent_id, Valid: menuInput.Menu_parent_id != ""},
			MenuLevel:       int32(menuInput.Menu_level),
			MenuNumberOrder: int32(menuInput.Menu_Number_order),
			MenuGroupName:   menuInput.Menu_group_name,
			ID:              menuInput.Id,
		}
		updateParamsList = append(updateParamsList, updateParams)

		if int32(menuInput.Menu_Number_order) != currentMenu.MenuNumberOrder {
			siblingMenus := getSiblingMenus(allMenus, currentMenu)
			updateParamsList = append(updateParamsList, processSiblingMenus(siblingMenus, currentMenu, menuInput)...)
		}
	}

	return updateParamsList, nil
}

// getSiblingMenus lấy danh sách các menu cùng cấp
func getSiblingMenus(allMenus []database.GetAllMenusRow, currentMenu database.GetAllMenusRow) []database.GetAllMenusRow {
	siblingMenus := []database.GetAllMenusRow{}
	for _, m := range allMenus {
		parentID := m.MenuParentID.String
		currentParentID := currentMenu.MenuParentID.String
		if m.MenuLevel == currentMenu.MenuLevel &&
			parentID == currentParentID &&
			m.ID != currentMenu.ID {
			siblingMenus = append(siblingMenus, m)
		}
	}

	sort.Slice(siblingMenus, func(i, j int) bool {
		return siblingMenus[i].MenuNumberOrder < siblingMenus[j].MenuNumberOrder
	})

	return siblingMenus
}

// processSiblingMenus xử lý cập nhật thứ tự cho các menu cùng cấp
func processSiblingMenus(siblingMenus []database.GetAllMenusRow, currentMenu database.GetAllMenusRow, menuInput model.MenuInput) []database.UpdateSingleMenuParams {
	var updateParamsList []database.UpdateSingleMenuParams
	newOrderNumber := int32(menuInput.Menu_Number_order)
	oldOrderNumber := currentMenu.MenuNumberOrder

	for _, m := range siblingMenus {
		var newOrder int32
		if newOrderNumber > oldOrderNumber {
			if m.MenuNumberOrder > oldOrderNumber && m.MenuNumberOrder <= newOrderNumber {
				newOrder = m.MenuNumberOrder - 1
				updateParamsList = append(updateParamsList, createUpdateParams(m, newOrder))
			}
		} else if newOrderNumber < oldOrderNumber {
			if m.MenuNumberOrder >= newOrderNumber && m.MenuNumberOrder < oldOrderNumber {
				newOrder = m.MenuNumberOrder + 1
				updateParamsList = append(updateParamsList, createUpdateParams(m, newOrder))
			}
		}
	}

	return updateParamsList
}

// createUpdateParams tạo tham số cập nhật cho menu
func createUpdateParams(menu database.GetAllMenusRow, newOrder int32) database.UpdateSingleMenuParams {
	return database.UpdateSingleMenuParams{
		MenuName:        menu.MenuName,
		MenuIcon:        menu.MenuIcon,
		MenuUrl:         menu.MenuUrl,
		MenuParentID:    menu.MenuParentID,
		MenuLevel:       menu.MenuLevel,
		MenuNumberOrder: newOrder,
		MenuGroupName:   menu.MenuGroupName,
		ID:              menu.ID,
	}
}
