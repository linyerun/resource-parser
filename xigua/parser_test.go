package xigua

import (
	"testing"

	"go.uber.org/zap"

	"github.com/linyerun/resource-parser/common"
)

func TestXigua(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	parser := NewVideoParser(logger)
	parserProxy := common.NewParserProxy(logger, parser)

	urlList := []string{
		"https://www.ixigua.com/7420283740938568207?logTag=2a4d1922395e5d6c1092",
		"https://v.ixigua.com/iSokg5YE/点击链接直接打开",
		"https://www.ixigua.com/7417027106590982667",
	}

	for _, url := range urlList {
		videoInfo, err := parserProxy.Parse(url)
		if err != nil {
			logger.Error("get video info err", zap.Error(err))
			continue
		}
		logger.Info("resp data", zap.Any("videoInfo", videoInfo))
	}
}
