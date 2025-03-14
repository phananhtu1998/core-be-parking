package model

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
