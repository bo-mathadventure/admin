package encoder

import (
	"github.com/bo-mathadventure/admin/ent"
	"github.com/gofiber/fiber/v2"
)

func ParseUser(user *ent.User) fiber.Map {
	data := fiber.Map{
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	}

	return data
}
