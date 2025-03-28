package model

type LoginInput struct {
	Username string `json:"username" binding:"required" example:"phananhtu1998"`
	Password string `json:"password" binding:"required" example:"123"`
}

type LoginOutput struct {
	ID           string `json:"id"`
	UserName     string `json:"username"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshToken"`
}

type GetCacheToken struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

type ChangePasswordInput struct {
	OldPassword     string `json:"oldpassword"`
	NewPassword     string `json:"newpassword"`
	ConfirmPassword string `json:"confirmpassword"`
}

type GetCacheTokenForChangePassword struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Number   int64  `json:"number"`
}
