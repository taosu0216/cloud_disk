package logic

import (
	"cloud_disk/core/models"
	"cloud_disk/core/utils"
	"context"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq) (resp *types.FileUploadResp, err error) {
	rp := &models.RepositoryPool{
		Identity:  utils.GenerateUUID(),
		Hash:      req.Hash,
		Name:      req.Name,
		Ext:       req.Ext,
		Size:      req.Size,
		Path:      req.Path,
		CreatedAt: time.Now(),
	}
	_, err = models.XormEngine.Insert(rp)
	if err != nil {
		return nil, err
	}
	resp = &types.FileUploadResp{}
	resp.Identity = rp.Identity
	return
}
