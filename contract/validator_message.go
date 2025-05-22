package contract

import "context"

func (m *Event) Validate(ctx context.Context) error {
	if m.GetClientName() == "" {
		return ErrorField("client_name").BuildError(ctx)
	}
	return nil
}
