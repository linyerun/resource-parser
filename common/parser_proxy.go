package common

import (
	"fmt"
	"net/url"
	"time"

	"go.uber.org/zap"

	"github.com/linyerun/resource-parser/util"
)

type parserProxy struct {
	logger *zap.Logger
	parser IVideoParser
}

func NewParserProxy(logger *zap.Logger, parser IVideoParser) IVideoParserProxy {
	return &parserProxy{
		logger: logger,
		parser: parser,
	}
}

func (p *parserProxy) Parse(arg string) (resourceInfo *VideoInfo, err error) {
	videoUrl, err := util.RegexpMatchUrlFromString(arg)
	if err != nil {
		return nil, err
	}

	urlObject, err := url.Parse(videoUrl)
	if err != nil {
		return nil, fmt.Errorf("parse video url fail, err: %v", err)
	}

	// 计算执行时间
	startTime := time.Now()
	defer func() {
		stopTime := time.Now()
		p.logger.Info("get video info data cost", zap.String("url", videoUrl), zap.Float64("cost-second", stopTime.Sub(startTime).Seconds()))
	}()

	return p.parser.Parse(urlObject)
}
