package repo

import (
	"go-backend-api/global"
	"go-backend-api/internal/database"
)

// INTERFACE VERSION

type IUserRepository interface {
	GetUserByEmai(email string) bool
}

type userReponsitory struct {
	sqlc *database.Queries
}

// GetUserByEmai implements IUserRepository.

func (up *userReponsitory) GetUserByEmai(email string) bool {
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	// user, err := up.sqlc.GetUserByEmaiSQLC(ctx, email)
	// if err != nil {
	// 	return false
	// }
	return true
}

func NewUserRepository() IUserRepository {
	return &userReponsitory{
		sqlc: database.New(global.Mdbc),
	}
}
