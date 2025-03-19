package model

type MenuInput struct {
	Menu_name         string  `json:"menu_name"`
	Menu_icon         string  `json:"menu_icon"`
	Menu_url          string  `json:"menu_url"`
	Menu_parent_id    string  `json:"menu_parent_id"`
	Menu_level        int     `json:"menu_level"`
	Menu_Number_order float64 `json:"menu_number_order"`
	Menu_group_name   string  `json:"menu_group_name"`
}
type MenuOutput struct {
	Id                string       `json:"id"`
	Menu_name         string       `json:"menu_name"`
	Menu_icon         string       `json:"menu_icon"`
	Menu_url          string       `json:"menu_url"`
	Menu_parent_id    string       `json:"menu_parent_id"`
	Menu_level        int          `json:"menu_level"`
	Menu_Number_order float64      `json:"menu_number_order"`
	Menu_group_name   string       `json:"menu_group_name"`
	Children          []MenuOutput `json:"children"`
}
