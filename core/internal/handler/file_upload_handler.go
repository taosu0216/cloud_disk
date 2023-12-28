package handler

import (
	"cloud_disk/core/models"
	"cloud_disk/core/utils"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud_disk/core/internal/logic"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		//以下为自定义内容,非框架生成
		//读取文件信息
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		buff := make([]byte, fileHeader.Size)
		//获取hash值
		hash := fmt.Sprintf("%x", md5.Sum(buff))
		//判断文件是否存在
		rp := new(models.RepositoryPool)
		isExist, err := models.XormEngine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if isExist {
			httpx.OkJsonCtx(r.Context(), w, &types.FileUploadResp{
				Identity: rp.Identity,
			})
			return

		}
		//向腾讯COS上传文件
		cosPath, err := utils.COSUploadFile(r)
		if err != nil {
			return
		}
		//向login中传递req正确的数据
		req.Name = fileHeader.Filename
		req.Hash = hash
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Path = cosPath

		//自动生成
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
