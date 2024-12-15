package helper

import "github.com/go-playground/validator/v10"

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func ApiResponse(Message string, Code int, Status string, Data interface{}) Response {
	return Response{
		Meta: Meta{
			Message: Message,
			Code:    Code,
			Status:  Status,
		},
		Data: Data,
	}
}

func FormatError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}