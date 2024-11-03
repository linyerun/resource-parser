package xhs

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
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
	client := resty.New()
	videoRes, err := client.R().
		SetHeader(common.HttpHeaderUserAgent, common.XhsUserAgent).
		Get(url.String())
	if err != nil {
		p.logger.Error("get xhs video info error", zap.Error(err))
		return nil, err
	}

	re := regexp.MustCompile(`window.__INITIAL_STATE__\s*=\s*(.*?)</script>`)
	findRes := re.FindSubmatch(videoRes.Body())
	if len(findRes) < 2 {
		p.logger.Error("get xhs video info from resp content error", zap.Error(err))
		return nil, errors.New("parse video json info from html fail")
	}

	jsonBytes := bytes.TrimSpace(findRes[1])

	nodeId := gjson.GetBytes(jsonBytes, "note.currentNoteId").String()
	data := gjson.GetBytes(jsonBytes, fmt.Sprintf("note.noteDetailMap.%s.note", nodeId))

	videoUrl := data.Get("video.media.stream.h264.0.masterUrl").String()

	// 获取图集图片地址
	imagesObjArr := data.Get("imageList").Array()
	images := make([]string, 0, len(imagesObjArr))
	if len(videoUrl) <= 0 {
		for _, imageItem := range imagesObjArr {
			imageUrl := imageItem.Get("urlDefault").String()
			if len(imageUrl) > 0 {
				images = append(images, imageUrl)
			}
		}
	}

	parseInfo := &common.VideoInfo{
		Desc:      data.Get("title").String(),
		VideoUrl:  data.Get("video.media.stream.h264.0.masterUrl").String(),
		CoverUrl:  data.Get("imageList.0.urlDefault").String(),
		ImageURLs: images,
		Author:    new(common.AuthorInfo),
	}
	parseInfo.Author.Uid = data.Get("user.userId").String()
	parseInfo.Author.Name = data.Get("user.nickname").String()
	parseInfo.Author.Avatar = data.Get("user.avatar").String()

	return parseInfo, nil
}
