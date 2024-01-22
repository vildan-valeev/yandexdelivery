package yandexlogistic

type CreateOptions struct {
	RequestID string `query:"request_id"`
}

type AcceptOptions struct {
	ClaimID string `query:"claim_id"`
}

type InfoOptions struct {
	ClaimID string `query:"claim_id"`
}

type CancelOptions struct {
	ClaimID string `query:"claim_id"`
}
