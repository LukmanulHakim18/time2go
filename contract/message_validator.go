package contract

import (
	"context"
	"net/http"
	"time"
)

func (m *Event) Validate(ctx context.Context) error {
	if m.GetClientName() == "" {
		return ErrorField("client_name").BuildError(ctx)
	}
	if m.GetEventName() == "" {
		return ErrorField("event_name").BuildError(ctx)
	}
	if m.GetEventId() == "" {
		return ErrorField("event_id").BuildError(ctx)
	}
	if m.GetScheduleAt() == "" {
		return ErrorField("schedule_at").BuildError(ctx)
	}
	if _, err := time.Parse(time.RFC3339, m.GetScheduleAt()); err != nil {
		return BuildError(http.StatusBadRequest, 4000,
			"format waktu schedule_at tidak sesuai dengan format RFC 3339.",
			"schedule_at time format is not valid; it must follow RFC 3339.")
	}
	if err := m.GetRequestConfig().Validate(ctx); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestConfig) Validate(ctx context.Context) error {
	if m.GetMethod() == "" {
		return ErrorField("method").BuildError(ctx)
	}
	if m.GetUrl() == "" {
		return ErrorField("url").BuildError(ctx)
	}

	return nil
}
