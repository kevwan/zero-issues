package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"primitive/internal/svc"
)

type StringsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStringsLogic(ctx context.Context, svcCtx *svc.ServiceContext) StringsLogic {
	return StringsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StringsLogic) Strings() (resp []string, err error) {
	// todo: add your logic here and delete this line

	return
}
