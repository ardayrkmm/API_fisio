package middleware

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string   `json:"id_user"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}