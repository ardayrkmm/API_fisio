package auth

type RegisterRequest struct {
	Email        string `json:"email" binding:"required, min=3, max=1000"`
	Nama         string `json:"nama" binding:"required, min=3, max=1000"`
	NomerTelepon int    `json:"no_telepon" binding:"required,number, min=13, max=15"`
	Password     string `json:"password" binding:"required"`
}