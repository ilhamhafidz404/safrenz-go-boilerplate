package utils

import "github.com/gofiber/fiber/v2"

// SFR-01-001 : SUCCESS
// SFR-01-002 : Empty from Database
// SFR-01-003 : ERROR
// SFR-01-004 : Middleware Block / Unauthorize
// SFR-01-005 : Cant Connect to third-party

type ResponseFormat struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, httpStatusCode int, customCode string, message string, data interface{}) error {
	return c.Status(httpStatusCode).JSON(ResponseFormat{
		StatusCode: customCode,
		Message:    message,
		Data:       data,
	})
}

func ErrorResponse(c *fiber.Ctx, httpStatusCode int, customCode string, message string) error {
	return c.Status(httpStatusCode).JSON(ResponseFormat{
		StatusCode: customCode,
		Message:    message,
		Data:       nil,
	})
}
