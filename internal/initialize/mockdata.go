package initialize

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-api/internal/database"
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
}
