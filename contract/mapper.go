package contract

import (
	"time"

	"github.com/LukmanulHakim18/time2go/model"
)

func FromProtoEvent(in *Event) model.Event {
	scheduleAt, _ := time.Parse(time.RFC3339, in.ScheduleAt)
	return model.Event{
		ClientName:    in.GetClientName(),
		EventName:     in.GetEventName(),
		EventID:       in.GetEventId(),
		ScheduleAt:    scheduleAt,
		Status:        in.GetStatus(),
		LastError:     in.GetLastError(),
		RequestConfig: FromProtoRequestConfig(in.GetRequestConfig()),
		RetryPolicy:   FromProtoRetryPolicy(in.GetRetryPolicy()),
	}
}
func FromProtoRequestAuth(in *BasicAuthConfig) *model.BasicAuthConfig {
	if in == nil {
		return nil
	}
	return &model.BasicAuthConfig{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
	}
}

func FromProtoRequestConfig(in *HTTPRequestConfig) model.HTTPRequestConfig {
	timeout, _ := time.ParseDuration(in.Timeout)
	return model.HTTPRequestConfig{
		Method:      in.GetMethod(),
		URL:         in.GetUrl(),
		Headers:     in.GetHeaders(),
		QueryParams: in.GetQueryParams(),
		Body:        in.GetBody(),
		Timeout:     timeout,
		Auth:        FromProtoRequestAuth(in.GetAuth()),
	}
}

func FromProtoRetryPolicy(in *RetryPolicy) model.RetryPolicy {
	return model.RetryPolicy{
		Type:         RetryTypeMapper(in.GetType()),
		RetryCount:   int(in.GetRetryCount()),
		MaxAttempts:  int(in.GetMaxAttempts()),
		AttemptCount: int(in.GetAttemptCount()),
	}
}

func RetryTypeMapper(in RetryPolicyType) model.RetryPolicyType {
	if in == RetryPolicyType_FIXED {
		return model.RETRY_POLICY_TYPE_FIXED
	} else {
		return model.RETRY_POLICY_TYPE_EXPONENTIAL
	}
}

func ToProtoEvent(in model.Event) *Event {
	return &Event{
		ClientName:    in.ClientName,
		EventName:     in.EventName,
		EventId:       in.EventID,
		ScheduleAt:    in.ScheduleAt.Format(time.RFC3339),
		Status:        in.Status,
		LastError:     in.LastError,
		RequestConfig: ToProtoRequestConfig(in.RequestConfig),
		RetryPolicy:   ToProtoRetryPolicy(in.RetryPolicy),
	}
}

func ToProtoRequestAuth(in *model.BasicAuthConfig) *BasicAuthConfig {
	if in == nil {
		return nil
	}
	return &BasicAuthConfig{
		Username: in.Username,
		Password: in.Password,
	}
}

func ToProtoRequestConfig(in model.HTTPRequestConfig) *HTTPRequestConfig {
	return &HTTPRequestConfig{
		Method:      in.Method,
		Url:         in.URL,
		Headers:     in.Headers,
		QueryParams: in.QueryParams,
		Body:        in.Body,
		Timeout:     in.Timeout.String(),
		Auth:        ToProtoRequestAuth(in.Auth),
	}
}

func ToProtoRetryPolicy(in model.RetryPolicy) *RetryPolicy {
	return &RetryPolicy{
		Type:         ToProtoRetryType(in.Type),
		RetryCount:   int32(in.RetryCount),
		MaxAttempts:  int32(in.MaxAttempts),
		AttemptCount: int32(in.AttemptCount),
	}
}

func ToProtoRetryType(in model.RetryPolicyType) RetryPolicyType {
	switch in {
	case model.RETRY_POLICY_TYPE_FIXED:
		return RetryPolicyType_FIXED
	case model.RETRY_POLICY_TYPE_EXPONENTIAL:
		return RetryPolicyType_EXPONENTIAL
	default:
		return RetryPolicyType_RETRY_POLICY_TYPE_UNSPECIFIED
	}
}
