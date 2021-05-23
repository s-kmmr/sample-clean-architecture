package model

type ApiError struct {
	code uint
	err  error
}

type ErrCode uint

const (
	Unknown ErrCode = iota
	ConnectionFailed
	MemberNotFound
	FailCreateMember
)

func NewApiError(_code uint, _err error) *ApiError {
	return &ApiError{
		code: _code,
		err:  _err,
	}
}
func (a *ApiError) Code() uint {
	return a.code
}
func (a *ApiError) Err() error {
	return a.err
}

func (a *ApiError) Error() string {
	return a.err.Error()
}
