package logic

import (
	"cloud_disk/core/models"
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.InfoResponse{}
	user := &models.UserBasic{}
	ifExist, err := models.XormEngine.Where("identity = ?", req.Identity).Get(user)
	if err != nil {
		return nil, err
	}
	if !ifExist {
		return nil, errors.New("用户不存在")
	}
	resp.Name = user.Name
	resp.Email = user.Email
	return
}
