package services

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandom4Digit() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}