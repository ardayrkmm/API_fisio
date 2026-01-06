package auth

type RegisterRequest struct {
	Nama         string `json:"nama" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	NomerTelepon string `json:"no_telepon" binding:"required"`
}
