package options

type DocumentOptions struct {
	ClaimID      string `query:"claim_id"`
	DocumentType string `query:"document_type"`
	Version      int64  `query:"version"`
	Status       string `query:"status"`
}
