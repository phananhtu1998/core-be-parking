package model

type License struct {
	RoleId    string `json:"role_id" binding:"required" example:"2b796313-1134-44b3-b527-2c27d41a1624"`
	DateStart string `json:"datestart" binding:"required" example:"2025-03-28 11:22:30"`
	DateEnd   string `json:"dateend" binding:"required" example:"2025-03-30 11:22:30"`
}

type LicenseOutput struct {
	License string `json:"license"`
}
