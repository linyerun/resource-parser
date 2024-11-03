package dy

import (
	"errors"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/linyerun/resource-parser/common"
)

// parseVideoByShareUrl 通过分享链接获取视频信息
func (p *videoParser) getVideoInfoByShareUrl(shareUrl *url.URL) (videoInfo *common.VideoInfo, err error) {
	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	res, err := client.R().
		SetHeader(common.HttpHeaderUserAgent, common.DefaultUserAgent).
		Get(shareUrl.String())

	// 非 resty.ErrAutoRedirectDisabled 错误时，返回错误
	if !errors.Is(err, resty.ErrAutoRedirectDisabled) {
		return nil, err
	}

	pageUrl, err := res.RawResponse.Location()
	if err != nil {
		return nil, err
	}

	// 跳转到西瓜视频了, 使用西瓜视频解析器来处理
	if strings.Contains(pageUrl.Host, "ixigua.com") {
		return p.xiguaVideoParser.Parse(pageUrl)
	}

	return p.getVideoInfoByPageUrl(pageUrl)
}
