package common

// AuthorInfo 作者信息
type AuthorInfo struct {
	Uid    string `json:"uid"`    // 作者id
	Name   string `json:"name"`   // 作者名称
	Avatar string `json:"avatar"` // 作者头像
}

// VideoInfo 视频信息
type VideoInfo struct {
	Author    *AuthorInfo `json:"author,omitempty"`    // 作者信息
	Desc      string      `json:"desc,omitempty"`      // 描述
	VideoUrl  string      `json:"video_url,omitempty"` // 视频播放地址
	MusicUrl  string      `json:"music_url,omitempty"` // 音乐播放地址
	CoverUrl  string      `json:"cover_url,omitempty"` // 视频封面地址
	ImageURLs []string    `json:"images,omitempty"`    // 图集图片地址列表
}
