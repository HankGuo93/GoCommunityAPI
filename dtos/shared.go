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

func CreateErrorDto(key string, err error) ErrorDto {
	return ErrorDto{
		BaseDto: BaseDto{
			Success:  false,
			Messages: []string{fmt.Sprintf("%s -> %v", key, err.Error())},
		},
		Errors: map[string]interface{}{"Error": err},
	}
}

func CreateErrorDtoWithMessage(message string) ErrorDto {
	return ErrorDto{
		BaseDto: BaseDto{
			Success:  false,
			Messages: []string{message},
		},
	}
}
