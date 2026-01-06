package services

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
rand.Seed(time.Now().UnixNano())
    otp := 1000 + rand.Intn(9000)
    return strconv.Itoa(otp)
}
