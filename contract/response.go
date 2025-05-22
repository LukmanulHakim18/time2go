package contract

import (
	"context"
	"fmt"
	"net/http"

	cError "github.com/LukmanulHakim18/core/error"
	cLang "github.com/LukmanulHakim18/core/lang"
)

func GetDefaultResponse(ctx context.Context, messageEn, MessageId string) *DefaultResponse {
	message := cLang.CustomTextByLanguage(ctx, messageEn, MessageId)
	return &DefaultResponse{
		Code:    "t2g-2000",
		Message: message,
	}
}
func ErrorField(field string) *cError.Error {
	return &cError.Error{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  "t2g-4000",
		LocalizedMessage: cError.Message{
			English:   fmt.Sprintf("error field %s", field),
			Indonesia: fmt.Sprintf("error pada field %s", field),
		},
	}
}
