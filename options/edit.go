package options

type EditOptions struct {
	ClaimID string `query:"claim_id"`
	Version int64  `query:"version"`
}
