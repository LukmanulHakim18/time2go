package model

import (
	"time"

	"github.com/LukmanulHakim18/time2go/util"
)

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

const (
	RETRY_POLICY_TYPE_FIXED       RetryPolicyType = "FIXED"
	RETRY_POLICY_TYPE_EXPONENTIAL RetryPolicyType = "EXPONENTIAL"
)

type RetryPolicy struct {
	Type         RetryPolicyType `json:"type"`          // e.g. "fixed", "exponential",
	RetryCount   int             `json:"retry_count"`   // sudah berapa kali dicoba
	MaxAttempts  int             `json:"max_attempts"`  // batas maksimal retry
	AttemptCount int             `json:"attempt_count"` // counter
}

// func event

func (e *Event) GetIndexKey() string {
	return util.CreateEventKey(util.KEY_TYPE_INDEX, e.ClientName, e.EventName, e.EventID)
}
func (e *Event) GetTriggerKey() string {
	return util.CreateEventKey(util.KEY_TYPE_TRIGGER, e.ClientName, e.EventName, e.EventID)
}
func (e *Event) GetDataKey() string {
	return util.CreateEventKey(util.KEY_TYPE_DATA, e.ClientName, e.EventName, e.EventID)
}
func (e *Event) GetLockKey() string {
	return util.CreateEventKey(util.KEY_TYPE_LOCK, e.ClientName, e.EventName, e.EventID)
}
