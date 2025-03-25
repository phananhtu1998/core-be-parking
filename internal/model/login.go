package model

type LoginInput struct {
	Email    string `json:"email" binding:"required" example:"phananhtu1998@gmail.com"`
	Password string `json:"password" binding:"required" example:"123"`
}

type LoginOutput struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshToken"`
}

type GetCacheToken struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type ChangePasswordInput struct {
	Password string `json:"password"`
}

type GetCacheTokenForChangePassword struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Number int64  `json:"number"`
}
