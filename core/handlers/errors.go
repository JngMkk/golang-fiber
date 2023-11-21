package handlers

type HTTPError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(err error) HTTPError {
	e := HTTPError{}
	e.Errors = make(map[string]interface{})

	switch v := err.(type) {
	default:
		e.Errors["body"] = v.Error()
	}

	return e
}

func NewAccessForbidden() HTTPError {
	e := HTTPError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "Access Forbidden"
	return e
}

func NewNotFound() HTTPError {
	e := HTTPError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "Resource Not Found"
	return e
}
