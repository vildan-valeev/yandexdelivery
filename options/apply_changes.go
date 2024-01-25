package options

type ApplyChangesOptions struct {
	ClaimID   string `query:"claim_id"`
	RequestID string `query:"request_id"`
}
