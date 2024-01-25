package models

type ApplyChangesRequest struct {
	Changes           []Change `json:"changes"`
	LastKnownRevision string   `json:"last_known_revision"`
}

type ApplyChangesResponse struct {
}

type Change struct {
	Comment string `json:"comment"`
}
