package models

// AuthRequest representa el JSON de entrada
type AuthRequest struct {
	Token string `json:"token" example:""`
}
