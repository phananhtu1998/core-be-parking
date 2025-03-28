package model

import "time"

type RoleAccount struct {
	Account_id string `json:"account_id" binding:"required" example:"b16fdfbc-da22-4ab2-b045-614677e536d6"`
	Role_id    string `json:"role_id" binding:"required" example:"351ec27d-2134-4993-b120-b72a950f5574"`
	License_id string `json:"license_id" example:"xxx-yyyy-zzz"`
}
type RoleAccountOutput struct {
	Id string `json:"id"`
	RoleAccount
	Create_at time.Time `json:"create_at"`
	Update_at string    `json:"update_at"`
}
