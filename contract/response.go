package contract

import (
	"context"
	"fmt"
	"net/http"

	cError "github.com/LukmanulHakim18/core/error"
	cLang "github.com/LukmanulHakim18/core/lang"
	"github.com/LukmanulHakim18/time2go/config"
)

func GetDefaultResponse(ctx context.Context, messageEn, MessageId string) *DefaultResponse {
	message := cLang.CustomTextByLanguage(ctx, messageEn, MessageId)
	return &DefaultResponse{
		Code:    "t2g-2000",
		Message: message,
	}
}

type ErrorConst string

func getErrorCode(code int) string {
	return fmt.Sprintf("%s-%d", config.GetConfig("app_code").GetString(), code)
}

func BuildError(httpStatusCode, errorCode int, messageId, messageEn string) *cError.Error {
	return &cError.Error{
		StatusCode: httpStatusCode,
		ErrorCode:  getErrorCode(errorCode),
		LocalizedMessage: cError.Message{
			English:   messageId,
			Indonesia: messageEn,
		},
	}
}

func ErrorField(field string) *cError.Error {
	return &cError.Error{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  getErrorCode(4000),
		LocalizedMessage: cError.Message{
			English:   fmt.Sprintf("error field %s", field),
			Indonesia: fmt.Sprintf("error pada field %s", field),
		},
	}
}
