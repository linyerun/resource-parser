package xigua

import (
	"bytes"
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"

	"github.com/linyerun/resource-parser/common"
)

func (p *videoParser) getVideoInfoByPageUrl(pageUrl *url.URL) (*common.VideoInfo, error) {
	// 获取videoId
	videoId := ""
	if pathParams := strings.Split(strings.Trim(pageUrl.Path, "/"), "/"); len(pathParams) != 0 && len(pathParams[len(pathParams)-1]) != 0 {
		videoId = pathParams[len(pathParams)-1]
		p.logger.Debug("parser video id from pageUrl success", zap.String("videoId", videoId))
	}

	if len(videoId) == 0 {
		p.logger.Error("can not get video id", zap.String("pageUrl", pageUrl.String()))
		return nil, errors.New("get video id fail")
	}

	// 校验videoId是否合法
	for _, c := range videoId {
		if c >= '0' && c <= '9' {
			continue
		}
		p.logger.Error("video id is invalid", zap.String("videoId", videoId))
		return nil, errors.New("video id is invalid")
	}

	return p._getVideoInfoById(videoId)
}

func (p *videoParser) _getVideoInfoById(videoId string) (*common.VideoInfo, error) {
	reqUrl := "https://m.ixigua.com/douyin/share/video/" + videoId + "?aweme_type=107&schema_type=1&utm_source=copy&utm_campaign=client_share&utm_medium=android&app=aweme"
	headers := map[string]string{
		common.HttpHeaderUserAgent: common.XiguaUserAgent,
	}

	client := resty.New()
	res, err := client.R().SetHeaders(headers).Get(reqUrl)
	if err != nil {
		p.logger.Error("get video info error", zap.String("reqUrl", reqUrl), zap.Error(err))
		return nil, err
	}

	re := regexp.MustCompile(`window._ROUTER_DATA\s*=\s*(.*?)</script>`)
	findRes := re.FindSubmatch(res.Body())
	if len(findRes) < 2 {
		p.logger.Error("get video info error", zap.String("reqUrl", reqUrl), zap.Error(err))
		return nil, errors.New("parse video json info from html fail")
	}

	jsonBytes := bytes.TrimSpace(findRes[1])
	videoData := gjson.GetBytes(jsonBytes, "loaderData.video_(id)/page.videoInfoRes.item_list.0")

	userId := videoData.Get("author.user_id").String()
	userName := videoData.Get("author.nickname").String()
	userAvatar := videoData.Get("author.avatar_thumb.url_list.0").String()
	videoDesc := videoData.Get("desc").String()
	videoAddr := videoData.Get("video.play_addr.url_list.0").String()
	coverUrl := videoData.Get("video.cover.url_list.0").String()

	parseRes := &common.VideoInfo{
		Desc:     videoDesc,
		VideoUrl: videoAddr,
		CoverUrl: coverUrl,
		Author:   new(common.AuthorInfo),
	}
	parseRes.Author.Uid = userId
	parseRes.Author.Name = userName
	parseRes.Author.Avatar = userAvatar

	return parseRes, nil
}
