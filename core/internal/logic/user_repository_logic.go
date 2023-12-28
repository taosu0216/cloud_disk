package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/core/utils"
	"context"
	"errors"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositoryLogic {
	return &UserRepositoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositoryLogic) UserRepository(req *types.UrsReq, userIdentity string) (resp *types.UrsResp, err error) {
	ur := new(models.UserRepository)
	ur.Identity = utils.GenerateUUID()
	ur.UserIdentity = userIdentity
	ur.ParentId = req.ParentId
	ur.Name = req.Name
	ur.Ext = req.Ext
	ur.RepositoryIdentity = req.RepositoryIdentity
	ur.CreatedAt = time.Now()
	_, err = models.XormEngine.Insert(ur)
	if err != nil {
		return nil, errors.New("用户存储池存入失败")
	}
	return
}
