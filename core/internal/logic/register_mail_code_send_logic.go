package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/core/pkg"
	"cloud_disk/core/utils"
	"context"
	"errors"
	"fmt"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterMailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterMailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterMailCodeSendLogic {
	return &RegisterMailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterMailCodeSendLogic) RegisterMailCodeSend(req *types.McsRequest) (resp *types.McsResponse, err error) {
	//判断该邮箱是否注册过
	resp = &types.McsResponse{}
	cnt, err := models.XormEngine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	fmt.Println("---------------------------------------", err)
	if err != nil {
		err = errors.New("未获取该邮箱验证码")
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("该邮箱已被注册")
		return nil, err
	}
	//生成随机验证码并发送
	code := utils.RandomCodeGenerate()
	err = utils.MailCodeSend(req.Email, code)
	if err != nil {
		err = errors.New("验证码发送失败")
		return nil, err
	}
	err = models.RedisEngine.Set(l.ctx, req.Email, code, pkg.EmailVerificationCodeTimeOut).Err()
	if err != nil {
		err = errors.New("验证码存储失败")
		return nil, err
	}
	resp.Code = code
	return
}
