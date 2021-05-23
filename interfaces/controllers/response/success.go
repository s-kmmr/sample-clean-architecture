package response

type Success struct {
	Msg string `json:"message"`
}

func NewSuccess(msg string) Success {
	return Success{
		Msg: msg,
	}
}
