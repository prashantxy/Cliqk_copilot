package models

type AdRequest struct {
	Product  string `json:"product" binding:"required"`
	Audience string `json:"audience" binding:"required"`
	Platform string `json:"platform" binding:"required"`
	Tone     string `json:"tone" binding:"required"`
}
