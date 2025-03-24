package model

import (
	"encoding/json"
	"time"
)

type RolesMenu struct {
	Id         string          `json:"id"`
	Menu_id    string          `json:"menu_id"`
	Role_id    string          `json:"role_id"`
	ListMethod json.RawMessage `json:"list_method"`
	Is_deleted bool            `json:"is_deleted"`
	Created_at time.Time       `json:"created_at"`
}
