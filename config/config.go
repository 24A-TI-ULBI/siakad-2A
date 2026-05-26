package config

import (
	"backend/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

var IPPort string
var Net string
var FiberConfig fiber.Config

// InitConfig loads environment variables from .env file
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	IPPort, Net = helper.GetAddress()
	
	// FiberConfig adalah konfigurasi Fiber mengikuti pola boilerplate gocroot
	FiberConfig = fiber.Config{
		Prefork:       false, // set true di production untuk multi-core
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "SIAKAD",
		AppName:       "Portal Informasi Akademik Kampus",
		Network:       Net,
	}
}
