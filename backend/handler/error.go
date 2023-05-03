package handler

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func HandleError(c *fiber.Ctx, errorCode string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "success": false, "status": fiber.StatusBadRequest, "message": errorCode})
}

func HandleErrorCode(c *fiber.Ctx, statsuCode int, errorCode string) error {
	return c.Status(statsuCode).JSON(fiber.Map{"error": true, "success": false, "status": statsuCode, "message": errorCode})
}

func HandleBodyParseError(c *fiber.Ctx, extra error) error {
	log.WithError(extra).Warn("failed to parse body request")
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "success": false, "status": fiber.StatusBadRequest, "message": "ERR_INVALID_REQUEST", "extra": extra.Error()})
}

func HandleSuccess(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"error": false, "status": fiber.StatusOK, "success": true})
}

func HandleInternalError(c *fiber.Ctx, err error) error {
	log.WithError(err).Error("internal server error")
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "success": false, "status": fiber.StatusInternalServerError, "message": err.Error()})
}

func HandleInvalidPermissions(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": true, "success": false, "status": fiber.StatusForbidden, "message": "ERR_INVALID_PERMISSIONS"})
}

func HandleInvalidLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "success": false, "status": fiber.StatusUnauthorized, "message": "ERR_INVALID_CREDENTIALS"})
}
func HandleInsufficentData(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "success": false, "status": fiber.StatusNotFound, "message": "ERR_INSUFFICIENT_DATA"})
}
