package model

import "time"

type HTTPRequestConfig struct {
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]string `json:"query_params"`
	Body        []byte            `json:"body"` // disimpan sebagai []byte
	Timeout     time.Duration     `json:"timeout"`
	Auth        *BasicAuthConfig  `json:"auth,omitempty"`
}

type BasicAuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
