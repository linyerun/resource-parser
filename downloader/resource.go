package downloader

import (
	"bufio"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type resourceDownloader struct {
	logger *zap.Logger
}

func NewResourceDownloader(logger *zap.Logger) IDownloader {
	return &resourceDownloader{
		logger: logger,
	}
}

func (d *resourceDownloader) Download(url string, writer io.Writer) (info *ResourceInfo, err error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			d.logger.Error("failed to close response body", zap.Error(err))
		}
	}()

	n, err := bufio.NewReader(resp.Body).WriteTo(writer)
	if err != nil {
		d.logger.Error("failed to write response body to writer", zap.Error(err))
		return
	}

	d.logger.Info("read resp body data to writer successfully", zap.Int64("bytes", n))

	info = &ResourceInfo{
		ContentType:   resp.Header.Get("Content-Type"),
		ContentLength: resp.ContentLength,
	}

	return
}
