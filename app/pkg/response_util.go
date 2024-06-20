package pkg

import (
	"go-gin-web/app/constant"
	"go-gin-web/app/dto"
)

func BuildResponse[T any](code string, status constant.ResponseStatus, data T) dto.BaseResponse[T] {
	return dto.BaseResponse[T]{
		Code: code,
		Desc: status.ResponseStatus(),
		Data: data,
	}
}
