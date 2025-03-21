package model

type RoleInput struct {
	Code             string `json:"code"`
	Role_name        string `json:"role_name"`
	Role_left_value  int    `json:"role_left_value"`
	Role_right_value int    `json:"role_right_value"`
	Role_max_number  int    `json:"role_max_number"`
	Is_licensed      bool   `json:"is_licensed"`
	Created_by       string `json:"created_by"`
	Is_deleted       bool   `json:"is_deleted"`
	Created_at       string `json:"created_at"`
	Updated_by       string `json:"updated_by"`
}

type RoleOutput struct {
	Id               string `json:"id"`
	Code             string `json:"code"`
	Role_name        string `json:"role_name"`
	Role_left_value  int    `json:"role_left_value"`
	Role_right_value int    `json:"role_right_value"`
	Role_max_number  int    `json:"role_max_number"`
	Is_licensed      bool   `json:"is_licensed"`
	Created_by       string `json:"created_by"`
	Is_deleted       bool   `json:"is_deleted"`
	Created_at       string `json:"created_at"`
	Updated_by       string `json:"updated_by"`
}
