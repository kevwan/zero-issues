package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"primitive/internal/svc"
)

type BoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) BoolLogic {
	return BoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BoolLogic) Bool() (resp bool, err error) {
	// todo: add your logic here and delete this line

	return
}
