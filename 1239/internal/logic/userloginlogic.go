package logic

import (
	"context"

	"issue/internal/svc"
	"issue/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.LoginReq) (*types.LoginReply, error) {
	// todo: add your logic here and delete this line

	return &types.LoginReply{}, nil
}
