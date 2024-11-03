package xhs

import (
	"testing"

	"go.uber.org/zap"

	"github.com/linyerun/resource-parser/common"
)

func TestXhs(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	parser := NewVideoParser(logger)
	parserProxy := common.NewParserProxy(logger, parser)
	urlDataList := []string{
		"79 【拍完美女很开心，碧蓝航线 cos - HzFlim📷 | 小红书 - 你的生活指南】 😆 eVctm6k2dAedyU9 😆 https://www.xiaohongshu.com/discovery/item/670bc39d000000001b02e0e9?source=webshare&xhsshare=pc_web&xsec_token=AB0KBLl1DGR4YiOwzKQSml-BW28EOdWJ3z1uJfT5DdCuc=&xsec_source=pc_share",
		"https://www.xiaohongshu.com/explore/670bc39d000000001b02e0e9?xsec_token=AB0KBLl1DGR4YiOwzKQSml-BW28EOdWJ3z1uJfT5DdCuc=&xsec_source=pc_feed",
		"53 【字节11月开始进入盘点hc，求爆料 - gogo学姐的职场侦察日记 | 小红书 - 你的生活指南】 😆 aOjAypx9M4Y5fTr 😆 https://www.xiaohongshu.com/discovery/item/67239eb4000000003c01af0c?source=webshare&xhsshare=pc_web&xsec_token=ABp6bd7QbZD9ws-Im8HXciRJUIiVADFQPRjjxayV8hvk0=&xsec_source=pc_share",
		"https://www.xiaohongshu.com/explore/6721b90a000000001b02c2b9?xsec_token=ABmcT80Kz2czMD_QioH9E7lnr5MX7Kwt6_ioKaqWWwNag=&xsec_source=pc_feed",
	}

	for _, urlData := range urlDataList {
		resp, err := parserProxy.Parse(urlData)
		if err != nil {
			logger.Error("parserProxy.Parse err", zap.String("urlData", urlData), zap.Error(err))
			continue
		}
		logger.Info("resp data", zap.Any("resp", resp))
	}
}
