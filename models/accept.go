package models

type AcceptResponse struct {
	ID                  string `json:"id"`
	SkipClientNotify    bool   `json:"skip_client_notify"`
	Status              string `json:"status"`
	UserRequestRevision string `json:"user_request_revision"`
	Version             int64  `json:"version"`
}

type AcceptRequest struct {
	Version int64 `json:"version"`
}
