package logic

import (
	"context"
	"io/ioutil"
	"logic-server/common"
	"logic-server/service/download/internal/svc"
	"logic-server/service/download/internal/types"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadPublicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadPublicLogic(ctx context.Context, svcCtx *svc.ServiceContext) DownloadPublicLogic {
	return DownloadPublicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadPublicLogic) DownloadPublic(urlPath string) (resp *types.ResponseByte, err error) {

	filePath := filepath.Join(l.svcCtx.Config.DownloadPath, urlPath)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, common.NewCodeError(common.FileErrorCode, err.Error(), nil)
	}

	return &types.ResponseByte{
		Data: data,
	}, nil
}
