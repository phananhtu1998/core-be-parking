package model

type AccountInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
	Images   string `json:"images"`
}
type AccountOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status bool   `json:"status"`
	Images string `json:"images"`
}
