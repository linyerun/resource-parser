package common

import "net/url"

type IVideoParser interface {
	// Parse arg: 可以是视频链接分享内容，也可以是视频页面url
	Parse(url *url.URL) (resp *VideoInfo, err error)
}

type IVideoParserProxy interface {
	Parse(arg string) (resp *VideoInfo, err error)
}
