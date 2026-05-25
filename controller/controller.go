package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
)

// Homepage handler
func Homepage(c *fiber.Ctx) error {
	if c.Accepts("html") != "" {
		return c.SendFile("./frontend/index.html")
	}

	info := model.AppInfo{
		Name:    "Portal Informasi Akademik Kampus",
		Version: "1.0.0",
		Status:  "Server is running",
	}
	return helper.SuccessResponse(c, info)
}

// IPServer handler — mengambil public IP via icanhazip.com (pola boilerplate gocroot)
func IPServer(c *fiber.Ctx) error {
	ip := helper.GetIPaddress()
	return helper.SuccessResponse(c, fiber.Map{
		"ip_address": ip,
	})
}
