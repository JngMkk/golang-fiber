package handlers

import "github.com/gofiber/fiber/v2"

type HTTPError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewHTTPResp(c *fiber.Ctx, status int, data interface{}) error {
	if err, ok := data.(error); ok {
		switch status {
		case 403:
			return c.Status(status).JSON(newAccessForbidden())
		case 404:
			return c.Status(status).JSON(newNotFound())
		case 409:
			return c.Status(status).JSON(newResourceConflict())
		default:
			return c.Status(status).JSON(newError(err))
		}
	} else {
		return c.Status(status).JSON(data)
	}
}

func newError(err error) HTTPError {
	e := HTTPError{}
	e.Errors = make(map[string]interface{})

	switch v := err.(type) {
	default:
		e.Errors["body"] = v.Error()
	}

	return e
}

func newAccessForbidden() HTTPError {
	return generateErrorBody("Access Forbidden")
}

func newNotFound() HTTPError {
	return generateErrorBody("Resource Not Found")
}

func newResourceConflict() HTTPError {
	return generateErrorBody("Resource Conflict")
}

func generateErrorBody(msg string) HTTPError {
	e := HTTPError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = msg
	return e
}
