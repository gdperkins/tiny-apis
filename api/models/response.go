package models

// ErrorResponse is the common error response use for all non success responses
type ErrorResponse struct {
	Message string      `json:"message"`
	Errors  []ErrorItem `json:"errors,omitempty"`
}

// ErrorItem represents a single error that occurred during the
// processing of a API request.
type ErrorItem struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// NewError returns a error response with a single initialized error item
func NewError(message string, code int, description string) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Errors: []ErrorItem{
			{
				Code:        code,
				Description: description,
			},
		},
	}
}
