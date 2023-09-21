package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	// Set Content-Type:application/json; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	// Retrieve the custom stCode code if it's an *AppError
	if e, ok := err.(*AppError); ok {
		code = e.StatusCode
	}
	// Retrieve the custom stCode code if it's an error of framework *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		err = DefaultErrWrap(code, e.Message, e)
	}
	// Return status-code with error message
	return c.Status(code).JSON(err)
}

type AppError struct {
	StatusCode int      `json:"status_code" bson:"status_code" xml:"status_code"`
	Err        string   `json:"error" bson:"error" xml:"error"`
	Message    string   `json:"message" bson:"message" xml:"message"`
	Cause      []string `json:"cause" bson:"cause" xml:"cause"`
}

func NewApiError(statusCode int, error string, message string, cause []string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Err:        error,
		Message:    message,
		Cause:      cause,
	}
}

func DefaultErrWrap(code int, message string, err error) *AppError {
	return NewApiError(
		code,
		getStringStatusCode(code),
		message,
		getCauses(err),
	)
}

func getStringStatusCode(code int) string {
	return strings.ReplaceAll(strings.ToLower(http.StatusText(code)), " ", "_")
}

func getCauses(err error) []string {
	causes := make([]string, 0)
	if err != nil && err.Error() != "" {
		s := err.Error()
		causes = strings.Split(s, "Cause:")
	}
	return causes
}

func (ae *AppError) MarshalJSON() ([]byte, error) { return json.Marshal(*ae) }

func (ae *AppError) Error() string {
	errCode := getStringStatusCode(ae.StatusCode)
	//ex.
	// 404_not_found: page not found, cause: ["does not exist in the system", "an internal error occurred"].
	return fmt.Sprintf("%v_%s: %s, cause: %v", errCode, ae.Err, ae.Message, ae.Cause)
}
