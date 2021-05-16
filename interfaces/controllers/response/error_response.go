package response

type ErrorResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(_code uint, _message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    _code,
		Message: _message,
	}
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
