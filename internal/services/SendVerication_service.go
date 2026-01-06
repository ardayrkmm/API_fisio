package services

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(to string, code string) error {
	m := gomail.NewMessage()

	// Mengambil data dari environment variables
	senderEmail := os.Getenv("SMTP_EMAIL")
	senderPass := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortRaw := os.Getenv("SMTP_PORT")

	// Konversi port dari string ke int
	smtpPort, _ := strconv.Atoi(smtpPortRaw)

	m.SetHeader("From", senderEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Kode Verifikasi Akun")

	m.SetBody("text/html", fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; text-align: center;">
			<h2>Verifikasi Akun</h2>
			<p>Kode verifikasi Anda adalah:</p>
			<h1 style="letter-spacing:5px; color: #2c3e50;">%s</h1>
			<p>Kode ini berlaku untuk waktu terbatas. Jangan bagikan kepada siapa pun.</p>
		</div>
	`, code))

	// Menggunakan variabel dari env untuk dialer
	d := gomail.NewDialer(
		smtpHost,
		smtpPort,
		senderEmail,
		senderPass,
	)

	return d.DialAndSend(m)
}