package initialize

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-backend-api/internal/database"
	"log"
)

func InsertData(ctx context.Context, queries *database.Queries) {
	accountID := "872ab326-b40b-4fb7-b28b-c5f8157fea7c"
	_, err := queries.GetAccountById(ctx, accountID)
	if err == sql.ErrNoRows {
		_, err = queries.InsertAccountData(ctx)

		if err == nil {
			fmt.Println("Inserted Account")
		}
	} else {
		fmt.Println("Account already exists")
	}
	roleID := "eb6d9850-2b77-47fc-ae2a-0a0ba9842280"
	_, err = queries.GetRoleById(ctx, roleID)
	if err == sql.ErrNoRows {
		_, err = queries.CreateRoleData(ctx)
		if err == nil {
			fmt.Println("Inserted Role")
		}
	} else {
		fmt.Println("Role already exists")
	}
	// Kiểm tra và insert License
	licenseID := "3375f96b-dcc5-492a-ab49-cb3b0af401a1"
	_, err = queries.GetLicenseById(ctx, licenseID)
	if err == sql.ErrNoRows {
		_, err = queries.CreateLicenseData(ctx)
		if err == nil {
			fmt.Println("Inserted License")
		}
	} else {
		fmt.Println("License already exists")
	}
	// Kiểm tra và insert RoleAccount
	roleAccountID := "369f54a1-300a-4ded-9ab0-b37e71cdc3e9"
	_, err = queries.GetRoleAccountById(ctx, roleAccountID)
	if err == sql.ErrNoRows {
		err = queries.CreateRoleAccountData(ctx)
		if err == nil {
			fmt.Println("Inserted RoleAccount")
		}
	} else {
		fmt.Println("RoleAccount already exists")
	}
	//Tạo role menu
	roleMenus := []struct {
		ID         string
		MenuID     string
		RoleID     string
		ListMethod []string
	}{
		{"369f54a1-300a-4ded-9ab0-b37e71cdc3e8", "efc78d7f-6dbb-4957-89d6-03d38be4869c", "eb6d9850-2b77-47fc-ae2a-0a0ba9842280", []string{"GET", "POST", "PUT", "DELETE"}},
		{"369f54a1-300a-4ded-9ab0-b37e71cdc3e7", "7c97cf69-9bc9-4c4f-8099-7078e728992d", "eb6d9850-2b77-47fc-ae2a-0a0ba9842280", []string{"GET", "POST", "PUT", "DELETE"}},
		{"369f54a1-300a-4ded-9ab0-b37e71cdc3e6", "c60dc41b-e454-4df2-887a-4a8359d45404", "eb6d9850-2b77-47fc-ae2a-0a0ba9842280", []string{"GET", "POST", "PUT", "DELETE"}},
	}

	for _, roleMenu := range roleMenus {
		// Truy vấn dựa trên ID thay vì RoleID
		ls, err := queries.GetRolesMenuByID(ctx, roleMenu.ID)
		log.Println("ls", ls)

		if err != nil {
			fmt.Println("Error fetching RoleMenu:", err)
			continue
		}

		if len(ls) == 0 { // Nếu danh sách rỗng thì mới insert
			listMethodJSON, jsonErr := json.Marshal(roleMenu.ListMethod)
			if jsonErr != nil {
				fmt.Println("Error marshaling ListMethod:", jsonErr)
				continue
			}

			err = queries.CreateRolesMenu(ctx, database.CreateRolesMenuParams{
				ID:         roleMenu.ID,
				MenuID:     roleMenu.MenuID,
				RoleID:     roleMenu.RoleID,
				ListMethod: listMethodJSON,
			})

			if err == nil {
				fmt.Println("Inserted RoleMenu:", roleMenu.ID)
			} else {
				fmt.Println("Error inserting RoleMenu:", err)
			}
		} else {
			fmt.Println("RoleMenu already exists:", roleMenu.ID)
		}
	}

	// Tạo menu
	menus := []struct {
		ID           string
		Name         string
		Icon         string
		Url          string
		ParentID     string
		Level        int
		Number_order int
		Group_name   string
	}{
		{"efc78d7f-6dbb-4957-89d6-03d38be4869c", "Tài khoản", "database", "tai-khoan", "", 1, 1, "/account"},
		{"7c97cf69-9bc9-4c4f-8099-7078e728992d", "Gói chức năng", "database", "goi-chuc-nang", "", 1, 2, "/funcpackage"},
		{"c60dc41b-e454-4df2-887a-4a8359d45404", "Chức năng", "database", "chuc-nang", "", 1, 3, "/menu"},
	}
	for _, menu := range menus {
		_, err = queries.GetMenuById(ctx, menu.ID)
		if err == sql.ErrNoRows {
			_, err = queries.InsertMenu(ctx, database.InsertMenuParams{
				ID:              menu.ID,
				MenuName:        menu.Name,
				MenuIcon:        menu.Icon,
				MenuUrl:         menu.Url,
				MenuParentID:    sql.NullString{String: menu.ParentID, Valid: menu.ParentID != ""},
				MenuLevel:       int32(menu.Level),
				MenuNumberOrder: int32(menu.Number_order),
				MenuGroupName:   menu.Group_name,
			})
			if err == nil {
				fmt.Println("Inserted Menu:", menu.Name)
			} else {
				fmt.Println("Error inserting menu:", menu.Name, err)
			}
		} else {
			fmt.Println("Menu already exists:", menu.Name)
		}
	}
}
