package model

type BaseResponse struct {
	ErrMsg int
	Msg    string
	Data   map[string]interface{}
}
