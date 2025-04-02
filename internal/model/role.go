package model

import "time"

type Role struct {
	Id               string    `json:"id"`
	Code             string    `json:"code"`
	Role_name        string    `json:"role_name"`
	Role_left_value  int       `json:"role_left_value"`
	Role_right_value int       `json:"role_right_value"`
	Role_max_number  int       `json:"role_max_number"`
	Is_licensed      bool      `json:"is_licensed"`
	Created_by       string    `json:"created_by"`
	Is_deleted       bool      `json:"is_deleted"`
	Created_at       time.Time `json:"created_at"`
	Updated_by       string    `json:"updated_by"`
}

type RoleHierarchyOutput struct {
	Id        string                `json:"id"`
	Code      string                `json:"code"`
	Role_name string                `json:"role_name"`
	Children  []RoleHierarchyOutput `json:"children"`
}

type RolePermission struct {
	Id              string `json:"id"`
	Role_name       string `json:"role_name"`
	Menu_group_name string `json:"menu_group_name"`
	Method          string `json:"method"`
}
