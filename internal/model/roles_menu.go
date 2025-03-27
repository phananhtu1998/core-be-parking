package model

type RolesMenu struct {
	Menu_id    string   `json:"menu_id" binding:"required" example:"menu-123"`
	Role_id    string   `json:"role_id" binding:"required" example:"role-123"`
	ListMethod []string `json:"list_method" binding:"required" example:"['GET','POST']"`
}

type RoleMenuOutput struct {
	Id string `json:"id" example:"rm-123"`
	RolesMenu
	Menu_name       string `json:"menu_name" example:"Menu 1"`
	Menu_url        string `json:"menu_url" example:"/menu1"`
	Menu_icon       string `json:"menu_icon" example:"fa-solid fa-house"`
	Menu_group_name string `json:"menu_group_name" example:"Group 1"`
	Role_code       string `json:"role_code" example:"role-123"`
	Role_name       string `json:"role_name" example:"Role 1"`
}
