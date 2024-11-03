package dy

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"go.uber.org/zap"

	"github.com/linyerun/resource-parser/common"
)

// https://www.douyin.com/discover?modal_id=7421873649080093963
// https://www.douyin.com/video/7418951980405722419?modeFrom=

type MyLogger struct {
}

func (m MyLogger) Debug(msg string, args ...any) {
	log.Println("[debug]"+msg, args)
}

func (m MyLogger) Info(msg string, args ...any) {
	log.Println("[info]"+msg, args)
}

func (m MyLogger) Warn(msg string, args ...any) {
	log.Println("[warn]"+msg, args)
}

func (m MyLogger) Error(msg string, args ...any) {
	log.Println("[error]"+msg, args)
}

func TestNewVideoParser(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	parser := NewVideoParser(logger)
	parserProxy := common.NewParserProxy(parser, logger)
	resp, err := parserProxy.Parse("https://www.douyin.com/video/7418951980405722419?modeFrom=")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("resp01 data", zap.Any("resp01", resp))

	arg := `4.33 o@q.Rk Btr:/ 04/10 勇闯非洲最大垃圾城，遭遇恶霸拦路，难怪当地人都不敢来# 非洲小钟 # vlog十亿流量扶持计划 # 埃及  https://v.douyin.com/iSwqYEsa/ 复制此链接，打开Dou音搜索，直接观看视频！`
	resp, err = parserProxy.Parse(arg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("resp02 data", zap.Any("resp01", resp))
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
