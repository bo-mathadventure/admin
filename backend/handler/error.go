package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"reflect"
)

// APIResponse definition for API responses
type APIResponse struct {
	Error            bool             `json:"error"`
	Success          bool             `json:"success"`
	StatusCode       int              `json:"status" example:"200"`
	Message          string           `json:"message,omitempty" example:"ERR_INVALID_PERMISSIONS"`
	ExtraData        string           `json:"extra,omitempty" example:"user don't have permissions"`
	ValidationErrors []*ErrorResponse `json:"validation,omitempty"`
}

// HandleError handles gerneral errors with HTTP 400
func HandleError(c *fiber.Ctx, errorCode string) error {
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusBadRequest,
		Message:    errorCode,
	})
}

// HandleErrorCode template for errors with custom statusCode
func HandleErrorCode(c *fiber.Ctx, statusCode int, errorCode string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: statusCode,
		Message:    errorCode,
	})
}

// HandleBodyParseError handles errors during body parsing with HTTP 400
func HandleBodyParseError(c *fiber.Ctx, extra error) error {
	log.WithError(extra).Warn("failed to parse body request")
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusBadRequest,
		Message:    "ERR_INVALID_REQUEST",
		ExtraData:  extra.Error(),
	})
}

// HandleInternalError handles all kind of interal errors with HTTP 500
func HandleInternalError(c *fiber.Ctx, err error) error {
	log.WithError(err).Error("internal server error")
	return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusInternalServerError,
		Message:    err.Error(),
	})
}

// HandleInvalidPermissions handles invalid permissions with HTTP 403
func HandleInvalidPermissions(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusForbidden,
		Message:    "ERR_INVALID_PERMISSIONS",
	})
}

// HandleInvalidLogin handles invalid credentials with HTTP 401
func HandleInvalidLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusUnauthorized,
		Message:    "ERR_INVALID_CREDENTIALS",
	})
}

// HandleInsufficientData handles insufficient data from workadventure with HTTP 404
func HandleInsufficientData(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusNotFound,
		Message:    "ERR_INSUFFICIENT_DATA",
	})
}

// HandleInvalidID handles invalid id errors with HTTP 400
func HandleInvalidID(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusBadRequest,
		Message:    "ERR_INVALID_ID",
	})
}

// HandleNotFound handles HTTP 404
func HandleNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusNotFound,
		Message:    "ERR_NOTFOUND",
	})
}

// ErrorResponse validation error fields
type ErrorResponse struct {
	FailedField string `json:"field" example:"validUntil"`
	Tag         string `json:"tag" example:"datetime"`
}

// Validate the validator instance. can be used to append custom validation on init
var Validate = validator.New()

// ValidateStruct verifies structs with validation Tags using the Validate validator
func ValidateStruct(c *fiber.Ctx, this interface{}) (bool, error) {
	var errors []*ErrorResponse
	err := Validate.Struct(this)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			failedFieldName := err.StructField()
			if field, success := reflect.TypeOf(this).Elem().FieldByName(err.StructField()); success {
				if jsonName, ok := field.Tag.Lookup("json"); ok {
					failedFieldName = jsonName
				}
			}
			errors = append(errors, &ErrorResponse{
				FailedField: failedFieldName,
				Tag:         err.Tag(),
			})
		}
	} else {
		return true, nil
	}
	return false, c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Error:            true,
		Success:          false,
		StatusCode:       fiber.StatusBadRequest,
		Message:          "ERR_VALIDATION",
		ValidationErrors: errors,
	})
}

// HandleSuccess handles all HTTP 200 with no data
func HandleSuccess(c *fiber.Ctx) error {
	return c.JSON(APIResponse{
		Error:      false,
		Success:    true,
		StatusCode: fiber.StatusOK,
	})
}
