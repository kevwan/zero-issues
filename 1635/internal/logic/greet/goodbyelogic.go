package greet

import (
	"context"

	"1635/internal/svc"
	"1635/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodByeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodByeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodByeLogic {
	return &GoodByeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodByeLogic) GoodBye(req types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
