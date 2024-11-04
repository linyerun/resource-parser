package downloader

import "io"

type IDownloader interface {
	Download(url string, writer io.Writer) (*ResourceInfo, error)
}
