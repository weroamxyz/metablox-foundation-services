package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidAuth
	CodeNeedLogin
	CodeError
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:     "success",
	CodeInvalidAuth: "Invalid auth",
	CodeNeedLogin:   "Please login first",
	CodeError:       "error",
}

func (c ResCode) Msg() string {
	return codeMsgMap[c]
}
