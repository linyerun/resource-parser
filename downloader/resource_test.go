package downloader

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestVideoDownloader(t *testing.T) {
	url := "https://sns-webpic-qc.xhscdn.com/202411042346/5576e2560d4b1b2a05b64df5215bf364/1040g2sg319ik2j22mu7g5pof029nc03bcd6u21g!nd_dft_wlteh_jpg_3"

	logger, _ := zap.NewDevelopment()
	buffer := bytes.NewBuffer(nil)
	resp, err := NewResourceDownloader(logger).Download(url, buffer)
	if err != nil {
		t.Fatal(err)
	}

	a := strings.Split(resp.ContentType, "/")
	if len(a) == 0 {
		t.Fatal("invalid content-type", resp)
	}

	filename := fmt.Sprintf("%d.%s", rand.Intn(100)+1, a[len(a)-1])
	file, err := os.OpenFile("H:\\code\\github\\resource-parser\\downloader\\"+filename, os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}

	_, err = buffer.WriteTo(file)
	if err != nil {
		t.Fatal(err)
	}

	logger.Info("download success", zap.Any("resp", resp))
}
