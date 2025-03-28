package model

type License struct {
	DateStart string `json:"datestart" binding:"required" example:"2025-03-28 11:22:30"`
	DateEnd   string `json:"dateend" binding:"required" example:"2025-03-30 11:22:30"`
}

type LicenseOutput struct {
	License string `json:"license"`
}
