package logic

import (
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UserFileNameModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameModifyLogic {
	return &UserFileNameModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameModifyLogic) UserFileNameModify(req *types.UfwReq, userIdentity string) (resp *types.UfwResp, err error) {
	data := new(models.UserRepository)
	data.UpdatedAt = time.Now()
	data.Name = req.Name
	fmt.Println(req.Identity, userIdentity)
	cnt, err := models.XormEngine.Table("user_repository").Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if cnt == 0 {
		return nil, errors.New("数据库修改失败")
	}
	if err != nil {
		return nil, err
	}
	return
}
