package dtos

import (
	"errors"
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

type BaseDto struct {
	Success  bool     `json:"success"`
	Messages []string `json:"message"`
}

type ErrorDto struct {
	BaseDto
	Errors map[string]interface{} `json:"errors"`
}

func CreateBadRequestErrorDto(err error) ErrorDto {
	res := ErrorDto{}
	res.Errors = make(map[string]interface{})
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errs := err.(validator.ValidationErrors)
		res.Messages = make([]string, len(errs))
		count := 0
		for _, v := range errs {
			if v.ActualTag == "required" {
				var message = fmt.Sprintf("%v is required", v.Field)
				res.Errors[v.Field] = message
				res.Messages[count] = message
			} else {
				var message = fmt.Sprintf("%v has to be %v", v.Field, v.ActualTag)
				res.Errors[v.Field] = message
				res.Messages = append(res.Messages, message)
			}
			count++
		}
		return res
	}
	res.Errors["Error"] = err.Error()
	return res
}

func CreateDetailedErrorDto(key string, err error) map[string]interface{} {
	return map[string]interface{}{
		"success":  false,
		"messages": []string{fmt.Sprintf("%s -> %v", key, err.Error())},
		"errors":   err,
	}
}

func CreateErrorDtoWithMessage(message string) map[string]interface{} {
	return map[string]interface{}{
		"success":       false,
		"full_messages": []string{message},
	}
}
