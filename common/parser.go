package common

type VideoParser interface {
	// Parse arg: 可以是视频链接分享内容，也可以是视频页面url
	Parse(arg string) (resp *VideoInfo, err error)
}
