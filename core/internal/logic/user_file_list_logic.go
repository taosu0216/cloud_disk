package logic

import (
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/core/pkg"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UflReq, userIdentity string) (resp *types.UflResp, err error) {
	// todo: add your logic here and delete this line
	uf := make([]*types.UserFiles, 0)
	resp = &types.UflResp{}
	//分页参数
	size := req.Size
	if size == 0 {
		size = pkg.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	/* UserFiles
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	*/
	//查询用户文件列表
	err = models.XormEngine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id,user_repository.name,user_repository.ext,repository_pool.path,repository_pool.size,user_repository.identity,user_repository.repository_identity").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Limit(size, offset).
		Find(&uf)
	if err != nil {
		return nil, err
	}
	//查询用户文件总数
	cnt, err := models.XormEngine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	resp.List = uf
	resp.Count = cnt
	return
}
