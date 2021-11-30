package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"primitive/internal/svc"
)

type IntLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntLogic {
	return IntLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntLogic) Int() (resp int, err error) {
	// todo: add your logic here and delete this line

	return
}
