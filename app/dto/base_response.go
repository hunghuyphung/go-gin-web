package dto

type BaseResponse[T any] struct {
	Code string
	Desc string
	Data T
}
