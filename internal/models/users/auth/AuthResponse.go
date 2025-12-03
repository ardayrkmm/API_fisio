package auth

import (
	userModel "api_fisioterapi/internal/models/users"
)
type AuthResponse struct {
	Message string `json:"message"`
	User   	userModel.PublicUser  `json:"user"`
	Token   string `json:"token"`
}