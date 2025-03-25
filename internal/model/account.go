package model

type AccountInput struct {
	Name     string `json:"name" binding:"required" example:"Admin"`
	Email    string `json:"email" binding:"required" example:"admin@gmail.com"`
	Password string `json:"password" binding:"required" example:"hello@1234"`
	Status   bool   `json:"status"`
	Images   string `json:"images" example:"/upload/images/phananhtu.jpg"`
}
type AccountOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status bool   `json:"status"`
	Images string `json:"images"`
}
