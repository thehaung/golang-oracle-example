package domain

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	RawMessage   string `json:"raw_message"`
	ErrorCode    int    `json:"error_code"`
}
