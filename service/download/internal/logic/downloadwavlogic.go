package logic

import (
	"context"
	"io/ioutil"
	"logic-server/common"
	"logic-server/service/download/internal/svc"
	"logic-server/service/download/internal/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadWavLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadWavLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadWavLogic {
	return &DownloadWavLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadWavLogic) DownloadWav(urlPath string) (resp *types.ResponseByte, err error) {

	filePath := filepath.Join(l.svcCtx.Config.DownloadPath, urlPath)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, common.NewCodeError(common.WavErrorCode, err.Error(), nil)
	}

	FileSuffixs := strings.Split(urlPath, ".")
	FileSuffix := FileSuffixs[len(FileSuffixs)-1]

	switch FileSuffix {
	case "wav":
	default:
		return nil, common.NewCodeError(common.WavTypeErrorCode, "不支持音频类型", nil)
	}

	if l.svcCtx.Config.WavIsDelete {
		go func() {
			os.Remove(filePath)
		}()
	}

	return &types.ResponseByte{
		Data: data,
	}, nil
}
