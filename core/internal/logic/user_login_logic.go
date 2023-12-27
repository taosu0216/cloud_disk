package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/core/utils"
	"context"
	"errors"
	"fmt"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	user := &models.UserBasic{}
	ifExist, err := models.XormEngine.Where("name = ? AND password = ?", req.Name, utils.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	fmt.Println(user.Id, user.Name, user.Identity)
	if !ifExist {
		return nil, errors.New("用户名或密码错误")
	}

	token := utils.GenerateToken(user.Id, user.Name, user.Identity)
	if token == "" {
		return nil, errors.New("生成token失败")
	}
	resp = &types.LoginResponse{
		Token: token,
	}
	return
}
