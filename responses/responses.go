package responses

// APIResponse is implemented by all the APIResponse* types.
type APIResponse interface {
	// Base returns the object of type APIResponseBase contained in each implemented type.
	Base() APIResponseBase
}

type APIResponseBase struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

// Base returns the APIResponseBase itself.
func (a APIResponseBase) Base() APIResponseBase {
	return a
}

type StatusCode int
