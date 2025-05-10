package contract

import (
	"context"

	cLang "github.com/LukmanulHakim18/core/lang"
)

func GetDefaultResponse(ctx context.Context, messageEn, MessageId string) *DefaultResponse {
	message := cLang.CustomTextByLanguage(ctx, messageEn, MessageId)
	return &DefaultResponse{
		Code:    "t2g-2000",
		Message: message,
	}
}
