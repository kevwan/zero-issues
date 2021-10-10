package logic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"user/service/user/api/internal/svc"
	"user/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLogic {
	return UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req types.LoginRequest) (*types.LoginResponse, error) {

	var mLoginResp types.LoginResponse

	err := l.svcCtx.Mysql.Table("user").Where("name=?", req.Name).First(&mLoginResp).Error
	if err != nil {
		return nil, err
	}

	//登录成功分发token
	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	secret := l.svcCtx.Config.Auth.AccessSecret

	token, err := l.getJwtToken(secret, now, expire, mLoginResp)
	if err != nil {
		return nil, errors.New("token生成出错：" + err.Error())
	}

	return &types.LoginResponse{
		Name:  "登录成功",
		Token: token,
	}, nil
}

func (l *UserLogic) getJwtToken(secretKey string, iat, seconds int64, user types.LoginResponse) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = user.Id
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
