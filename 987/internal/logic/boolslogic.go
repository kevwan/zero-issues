package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"primitive/internal/svc"
)

type BoolsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBoolsLogic(ctx context.Context, svcCtx *svc.ServiceContext) BoolsLogic {
	return BoolsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BoolsLogic) Bools() (resp []bool, err error) {
	// todo: add your logic here and delete this line

	return
}
