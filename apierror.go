package yandexlogistic

import "fmt"

// APIError относится к типу Error. Не использовать для маршаллинга запросов.
type APIError struct {
	code    Code
	message string
}

// ErrorCode returns the error code received from the Telegram API.
func (a *APIError) ErrorCode() string {
	return string(a.code)
}

// Description returns the error description received from the Telegram API.
func (a *APIError) Description() string {
	return a.message
}

// Error returns the error string.
func (a *APIError) Error() string {
	return fmt.Sprintf("API error: %s %s", a.ErrorCode(), a.Description())
}

// Статусы из YandexApi
type Code string

const (
	NotFound    Code = "not_found"
	OldVersionc Code = "old_version"
)
