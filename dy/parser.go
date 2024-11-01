package dy

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"

	"github.com/linyerun/resource-parser/common"
)

type videoParser struct {
}

func NewVideoParser() common.VideoParser {
	return &videoParser{}
}

func (p *videoParser) Parse(arg string) (resourceInfo *common.VideoInfo, err error) {
	return nil, nil
}

func (p *videoParser) parseVideoByPageUrl(pageUrl string) (resourceInfo *common.VideoInfo, err error) {
	if len(pageUrl) <= 0 {
		return nil, errors.New("page url is empty")
	}

	pagePath := strings.TrimPrefix(pageUrl, "https://")
	pagePath = strings.TrimPrefix(pageUrl, "http://")
	pagePath = strings.Trim(pagePath, "/") // 去除末尾的/

	urlPathParts := strings.Split(pageUrl, "/")

	if len(urlPathParts) == 0 {
		return nil, errors.New("get video id from page url fail")
	}

	videoId := urlPathParts[len(urlPathParts)-1]
	_ = videoId

	return nil, nil
}

// parseVideoById 通过视频ID获取视频信息
func (p *videoParser) parseVideoById(videoId string) (videoInfo *common.VideoInfo, err error) {
	// 发送请求获取视频信息的响应
	client := resty.New()
	reqURL := generateVideoReqURLById(videoId)
	res, err := client.R().SetHeader(common.HttpHeaderUserAgent, common.DouyinUserAgent).Get(reqURL)
	if err != nil {
		return
	}

	// (.*?)匹配到的是一个json字符串，通过findRes[1]获取到的是一个json字符串的字节数组
	re := regexp.MustCompile(`window._ROUTER_DATA\s*=\s*(.*?)</script>`)
	findRes := re.FindSubmatch(res.Body())
	if len(findRes) < 2 {
		err = errors.New("parse video json info from html fail")
		return
	}

	// 去除换行、空值、制表符等
	jsonBytes := bytes.TrimSpace(findRes[1])

	// 属性loaderData.video_(id)/page的属性videoInfoRes的属性item_list的第0个元素
	data := gjson.GetBytes(jsonBytes, "loaderData.video_(id)/page.videoInfoRes.item_list.0")
	if !data.Exists() {
		filterObj := gjson.GetBytes(
			jsonBytes,
			fmt.Sprintf(`loaderData.video_(id)/page.videoInfoRes.filter_list.#(aweme_id=="%s")`, videoId),
		)

		return nil, fmt.Errorf(
			"get video info fail: %s - %s",
			filterObj.Get("filter_reason"),
			filterObj.Get("detail_msg"),
		)
	}

	// 获取图集图片地址（视频没有图集）
	imagesObjArr := data.Get("images").Array()
	images := make([]string, 0, len(imagesObjArr))
	for _, imageItem := range imagesObjArr {
		imageUrl := imageItem.Get("url_list.0").String()
		if len(imageUrl) > 0 {
			images = append(images, imageUrl)
		}
	}

	// 获取视频播放地址
	videoUrl := data.Get("video.play_addr.url_list.0").String()
	videoUrl = strings.ReplaceAll(videoUrl, "playwm", "play")

	// 如果图集地址不为空时，因为没有视频，上面抖音返回的视频地址无法访问，置空处理
	if len(images) > 0 {
		videoUrl = ""
	}

	videoInfo = &common.VideoInfo{
		Desc:      data.Get("desc").String(),
		VideoUrl:  videoUrl,
		MusicUrl:  "",
		CoverUrl:  data.Get("video.cover.url_list.0").String(),
		ImageURLs: images,
		Author:    new(common.AuthorInfo),
	}
	videoInfo.Author.Uid = data.Get("author.sec_uid").String()
	videoInfo.Author.Name = data.Get("author.nickname").String()
	videoInfo.Author.Avatar = data.Get("author.avatar_thumb.url_list.0").String()

	// 视频地址非空时，获取302重定向之后的视频地址；图集时，视频地址为空，不处理。
	if len(videoInfo.VideoUrl) > 0 {
		p.getRedirectUrl(videoInfo)
	}

	return videoInfo, nil
}

// parseVideoByShareUrl 通过分享链接获取视频信息
func (p *videoParser) parseVideoByShareUrl(shareUrl string) (videoInfo *common.VideoInfo, err error) {
	return nil, nil
}

// getRedirectUrl 获取视频302重定向之后的地址
func (p *videoParser) getRedirectUrl(videoInfo *common.VideoInfo) {
	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	res, _ := client.R().
		SetHeader(common.HttpHeaderUserAgent, common.DefaultUserAgent).
		Get(videoInfo.VideoUrl)
	locationRes, _ := res.RawResponse.Location()
	if locationRes != nil {
		(*videoInfo).VideoUrl = locationRes.String()
	}
}
