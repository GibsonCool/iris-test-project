package model

type BaseResponse struct {
	Status  int
	Success string
	Message string
	Type    string
	Data    interface{}
}
