package downloader

type ResourceInfo struct {
	ContentType   string `json:"contentType"`
	ContentLength int64  `json:"contentLength"`
}
