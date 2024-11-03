package dy

import (
	"fmt"
	"net/url"
	"testing"

	"go.uber.org/zap"

	"github.com/linyerun/resource-parser/common"
)

// https://www.douyin.com/discover?modal_id=7421873649080093963
// https://www.douyin.com/video/7418951980405722419?modeFrom=
// https://www.douyin.com/search/%E5%9B%BE%E6%96%87?aid=8cebfbab-a444-4163-846d-b186a34bd1df&modal_id=7306487778492058889&type=general

func TestNewVideoParser(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	parser := NewVideoParser(logger)
	parserProxy := common.NewParserProxy(logger, parser)

	resp, err := parserProxy.Parse("https://www.douyin.com/search/%E5%9B%BE%E6%96%87?aid=8cebfbab-a444-4163-846d-b186a34bd1df&modal_id=7306487778492058889&type=general")
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("resp", zap.Any("resp", resp))

	resp, err = parserProxy.Parse("https://www.douyin.com/video/7418951980405722419?modeFrom=")
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("resp data", zap.Any("resp01", resp))

	arg := `4.33 o@q.Rk Btr:/ 04/10 勇闯非洲最大垃圾城，遭遇恶霸拦路，难怪当地人都不敢来# 非洲小钟 # vlog十亿流量扶持计划 # 埃及  https://v.douyin.com/iSwqYEsa/ 复制此链接，打开Dou音搜索，直接观看视频！`
	resp, err = parserProxy.Parse(arg)
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("resp data", zap.Any("resp02", resp))
}

func TestGetVideoInfoByPageUrl(t *testing.T) {
	//pageUrl := "https://www.douyin.com/discover?modal_id=7421873649080093963"
	pageUrl := "https://www.douyin.com/video/7418951980405722419?modeFrom="
	logger, _ := zap.NewDevelopment()
	parser := &videoParser{logger: logger}
	pageUrlInfo, _ := url.Parse(pageUrl)
	uri := pageUrlInfo.String()
	t.Log(uri)
	videoInfo, err := parser.getVideoInfoByPageUrl(pageUrlInfo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(videoInfo)
}

func TestGetVideoInfoByShareUrl(t *testing.T) {
	fmt.Printf("TestGetVideoInfoByShareUrl")
	shareUrl, _ := url.Parse("https://v.douyin.com/iSw4H2Pa/")
	logger, _ := zap.NewDevelopment()
	parser := &videoParser{logger: logger}
	videoInfo, err := parser.getVideoInfoByShareUrl(shareUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(videoInfo)
}
