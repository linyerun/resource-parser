package dy

import (
	"errors"
	"net/url"

	"go.uber.org/zap"

	"github.com/linyerun/resource-parser/common"
)

type videoParser struct {
	logger *zap.Logger
}

func NewVideoParser(logger *zap.Logger) common.IVideoParser {
	return &videoParser{
		logger: logger,
	}
}

func (p *videoParser) Parse(videoUrl *url.URL) (resp *common.VideoInfo, err error) {
	switch videoUrl.Host {
	case "www.iesdouyin.com", "www.douyin.com": // 可以解析出videoId
		return p.getVideoInfoByPageUrl(videoUrl)
	case "v.douyin.com":
		return p.getVideoInfoByShareUrl(videoUrl)
	}

	p.logger.Error("can not parse video url", zap.String("videoUrl", videoUrl.String()))

	return nil, errors.New("can not parse video url")
}
