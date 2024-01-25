package models

type ApplyChangesResultRequest struct {
	Changes []ApplyChangeResult `json:"changes"`
}

type ApplyChangesResultResponse struct {
}

type ApplyChangeResult struct {
	Error  ChangeError `json:"error"`
	ID     int64       `json:"id"`
	Kind   string      `json:"kind"`
	Status Status      `json:"status"`
}

type ChangeError struct {
	Code    string `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}
