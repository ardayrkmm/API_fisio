package authCon

import (
	ds "api_fisioterapi/internal/config"
	middlewares "api_fisioterapi/internal/middleware"
	userModel "api_fisioterapi/internal/models/users"
	dss "api_fisioterapi/internal/services"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateEmailVerificationToken(userID, email string) (string, string, error) {
    // 1. Generate angka 4 digit
    rand.Seed(time.Now().UnixNano())
    otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

    // 2. Masukkan OTP ke dalam JWT Claims
    expiration := time.Now().Add(15 * time.Minute) // Berlaku 15 menit
    claims := jwt.MapClaims{
        "user_id":  userID,
        "otp":      otpCode, // Simpan OTP di sini
        "purpose":  "email_verification",
        "exp":      expiration.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secret := os.Getenv("JWT_SECRET")
    signedToken, err := token.SignedString([]byte(secret))

    return signedToken, otpCode, err
}
func SendVerificationToken(c *gin.Context) {
    userID, _ := middlewares.GetUserIDFromContext(c)

    var user userModel.User
    ds.DB.First(&user, "id_user = ?", userID)

    // --- PERBAIKAN DI SINI ---
    // Generate 4 angka acak
    rand.Seed(time.Now().UnixNano())
    otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

    // Update kode di database agar yang lama tidak berlaku
    ds.DB.Model(&user).Update("verifikasi_status", otpCode)

    // Kirim angka OTP ke email
    err := dss.SendVerificationEmail(user.Email, otpCode)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to send email"})
        return
    }

    c.JSON(200, gin.H{
        "message": "Kode verifikasi 4 digit baru telah dikirim ke " + user.Email,
    })
}
type VerifyRequest struct {
    OTP   string `json:"otp" binding:"required"`   // Angka 4 digit dari email
    Token string `json:"token" binding:"required"` // JWT dari response register
}

func VerifyEmail(c *gin.Context) {
    var req VerifyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Data tidak lengkap"})
        return
    }

    // 1. Bongkar JWT-nya
    secret := os.Getenv("JWT_SECRET")
    token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil || !token.Valid {
        c.JSON(400, gin.H{"error": "Sesi verifikasi habis atau tidak valid"})
        return
    }

    claims := token.Claims.(jwt.MapClaims)

    // 2. COCOKKAN: OTP dari user vs OTP dari dalam JWT
    otpInToken, ok := claims["otp"].(string)
	if !ok {
    c.JSON(400, gin.H{"error": "Format token salah (OTP tidak ditemukan)"})
    return
	}

// 2. GUNAKAN VARIABELNYA DI SINI (Agar tidak error "declared and not used")
	if req.OTP != otpInToken {
    c.JSON(400, gin.H{"error": "Kode OTP salah!"})
    return
	}

    // 3. Jika cocok, ambil userID dan update status di DB jadi 1
    userID := claims["user_id"].(string)
    ds.DB.Model(&userModel.User{}).Where("id_user = ?", userID).Update("verifikasi_status", 1)

    c.JSON(200, gin.H{"message": "Verifikasi Berhasil!"})
}