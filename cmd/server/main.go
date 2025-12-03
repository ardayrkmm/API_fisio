package main

import (
	"log"

	"api_fisioterapi/internal/config"
)

func main() {
    config.InitDB()

    log.Println("Server berjalan...")
}
