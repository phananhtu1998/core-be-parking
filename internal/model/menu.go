package model

import "database/sql"

type MenuInput struct {
	Id                string `json:"id"`
	Menu_name         string `json:"menu_name" binding:"required" example:"Dashboard"`
	Menu_icon         string `json:"menu_icon" binding:"required" example:"database"`
	Menu_url          string `json:"menu_url" binding:"required" example:"/dashboard"`
	Menu_parent_id    string `json:"menu_parent_id"`
	Menu_level        int    `json:"menu_level" binding:"required" example:"1"`
	Menu_Number_order int    `json:"menu_number_order" binding:"required" example:"1"`
	Menu_group_name   string `json:"menu_group_name" binding:"required" example:"/api/dashboard"`
}
type MenuOutput struct {
	Id                string       `json:"id"`
	Menu_name         string       `json:"menu_name"`
	Menu_icon         string       `json:"menu_icon"`
	Menu_url          string       `json:"menu_url"`
	Menu_parent_id    string       `json:"menu_parent_id"`
	Menu_level        int          `json:"menu_level"`
	Menu_Number_order int          `json:"menu_number_order"`
	Menu_group_name   string       `json:"menu_group_name"`
	Children          []MenuOutput `json:"children"`
}

type MenuOutputFuncpackage struct {
	Id                string `json:"id"`
	STT               string `json:"stt"`
	Menu_name         string `json:"menu_name"`
	Menu_icon         string `json:"menu_icon"`
	Menu_url          string `json:"menu_url"`
	Menu_parent_id    string `json:"menu_parent_id"`
	Menu_level        int    `json:"menu_level"`
	Menu_Number_order int    `json:"menu_number_order"`
	Menu_group_name   string `json:"menu_group_name"`
}

type MenuSwag struct {
	ID              string         `json:"id"`
	MenuName        string         `json:"menu_name"`
	MenuIcon        string         `json:"menu_icon"`
	MenuUrl         string         `json:"menu_url"`
	MenuParentID    sql.NullString `json:"menu_parent_id"` // NULLABLE để tránh lỗi khi không có parent
	MenuLevel       int            `json:"menu_level"`
	MenuNumberOrder int            `json:"menu_number_order"`
	MenuGroupName   string         `json:"menu_group_name"`
	IsDeleted       bool           `json:"is_deleted"`
}
