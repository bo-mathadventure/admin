package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type APIResponse struct {
	Error            bool             `json:"error"`
	Success          bool             `json:"success"`
	StatusCode       int              `json:"status" example:"200"`
	Message          string           `json:"message,omitempty" example:"ERR_INVALID_PERMISSIONS"`
	ExtraData        string           `json:"extra,omitempty" example:"user don't have permissions"`
	ValidationErrors []*ErrorResponse `json:"validation,omitempty"`
}

func HandleError(c *fiber.Ctx, errorCode string) error {
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusBadRequest,
		Message:    errorCode,
	})
}

func HandleErrorCode(c *fiber.Ctx, statusCode int, errorCode string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: statusCode,
		Message:    errorCode,
	})
}

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

func HandleInternalError(c *fiber.Ctx, err error) error {
	log.WithError(err).Error("internal server error")
	return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusInternalServerError,
		Message:    err.Error(),
	})
}

func HandleInvalidPermissions(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusForbidden,
		Message:    "ERR_INVALID_PERMISSIONS",
	})
}

func HandleInvalidLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusUnauthorized,
		Message:    "ERR_INVALID_CREDENTIALS",
	})
}
func HandleInsufficientData(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusNotFound,
		Message:    "ERR_INSUFFICIENT_DATA",
	})
}
func HandleInvalidID(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusBadRequest,
		Message:    "ERR_INVALID_ID",
	})
}

func HandleNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(APIResponse{
		Error:      true,
		Success:    false,
		StatusCode: fiber.StatusNotFound,
		Message:    "ERR_NOTFOUND",
	})
}

type ErrorResponse struct {
	FailedField string `json:"field" example:"validUntil"`
	Tag         string `json:"tag" example:"datetime"`
}

var Validate = validator.New()

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

func HandleSuccess(c *fiber.Ctx) error {
	return c.JSON(APIResponse{
		Error:      false,
		Success:    true,
		StatusCode: fiber.StatusOK,
	})
}
