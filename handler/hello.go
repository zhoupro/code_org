package handler

import (
	"code_org/handler/dto"
	"context"
)

func Hello(ctx context.Context, req dto.ReqHello) (dto.ResHello, error) {
	ret := new(dto.ResHello)
	ret.Echo = req.Name
	return *ret, nil
}
