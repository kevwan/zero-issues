package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"primitive/internal/svc"
)

type StringLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStringLogic(ctx context.Context, svcCtx *svc.ServiceContext) StringLogic {
	return StringLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StringLogic) String() (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
