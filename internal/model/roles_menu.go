package model

import (
	"time"
)

type RolesMenu struct {
	Id         string    `json:"id" example:"rm-123"`
	Menu_id    string    `json:"menu_id" example:"menu-123"`
	Role_id    string    `json:"role_id" example:"role-123"`
	ListMethod []string  `json:"list_method" example:"['GET','POST']"`
	Is_deleted bool      `json:"is_deleted" example:"false"`
	Created_at time.Time `json:"created_at" example:"2025-03-24T12:00:00Z"`
}
