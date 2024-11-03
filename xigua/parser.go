package xigua

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

func (p *videoParser) Parse(url *url.URL) (resp *common.VideoInfo, err error) {
	switch url.Host {
	case "v.ixigua.com":
		return p.getVideoInfoByShareUrl(url)
	case "www.ixigua.com":
		return p.getVideoInfoByPageUrl(url)
	}

	p.logger.Error("unknown host", zap.String("host", url.Host))

	return nil, errors.New("unknown host")
}
