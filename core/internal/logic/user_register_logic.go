package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/core/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserRegisterResponse{}
	code, err := models.RedisEngine.Get(l.ctx, req.Email).Result()
	//判断验证码是否读取成功
	if err != nil {
		return nil, err
	}
	//判断验证码是否一致
	if code != req.VerifyCode {
		return nil, errors.New("验证码错误")
	}
	//判断用户名是否重复
	cnt, err := models.XormEngine.Where("username = ?", req.Username).Count(new(models.UserBasic))
	if cnt > 0 {
		return nil, errors.New("用户名已被注册")
	}
	//存入mysql
	user := &models.UserBasic{
		Identity:  utils.GenerateUUID(),
		Name:      req.Username,
		Password:  utils.Md5(req.Password),
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
	n, err := models.XormEngine.Insert(user)
	if err != nil {
		return nil, err
	}
	fmt.Println("insert into", n, "rows")
	return
}
