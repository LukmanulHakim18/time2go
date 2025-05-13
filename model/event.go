package model

import "time"

type Event struct {
	ClientName    string            `json:"client_name"`
	EventName     string            `json:"event_name"`
	EventID       string            `json:"event_id"`
	ScheduleAt    time.Time         `json:"schedule_at"`
	Status        string            `json:"status"`         // waiting, running, success, failed
	LastError     string            `json:"last_error"`     // jika gagal, simpan error terakhir
	RequestConfig HTTPRequestConfig `json:"request_config"` // konfigurasi http
	RetryPolicy   RetryPolicy       `json:"retry_policy"`   // strategi retry
}

type RetryPolicyType string

type RetryPolicy struct {
	Type        RetryPolicyType `json:"type"`         // e.g. "fixed", "exponential",
	RetryCount  int             `json:"retry_count"`  // sudah berapa kali dicoba
	MaxAttempts int             `json:"max_attempts"` // batas maksimal retry
}
