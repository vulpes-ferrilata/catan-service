package app_errors

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBusinessRuleError(message string, params ...string) AppError {
	return &bussinessRuleError{
		message: message,
		params:  params,
	}
}

type bussinessRuleError struct {
	message string
	params  []string
}

func (b bussinessRuleError) Error() string {
	return fmt.Sprintf("business rule violation: %s", b.message)
}

func (b bussinessRuleError) Status(translator ut.Translator) *status.Status {
	translatedParams := make([]string, 0)

	for _, param := range b.params {
		if translatedParam, err := translator.T(param); err == nil {
			translatedParams = append(translatedParams, translatedParam)
		} else {
			translatedParams = append(translatedParams, param)
		}
	}

	detail, err := translator.T(b.message, translatedParams...)
	if err != nil {
		detail = b.message
	}

	stt := status.New(codes.FailedPrecondition, detail)

	return stt
}
