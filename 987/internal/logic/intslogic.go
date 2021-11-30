package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"primitive/internal/svc"
)

type IntsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntsLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntsLogic {
	return IntsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntsLogic) Ints() (resp []int, err error) {
	// todo: add your logic here and delete this line

	return
}
