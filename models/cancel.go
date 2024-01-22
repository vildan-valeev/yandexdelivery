package models

type ResponseCancel struct {
	ID                  string `json:"id,omitempty"`
	SkipClientNotify    bool   `json:"skip_client_notify,omitempty"`
	Status              string `json:"status,omitempty"`
	UserRequestRevision string `json:"user_request_revision,omitempty"`
	Version             int64  `json:"version,omitempty"`
}
