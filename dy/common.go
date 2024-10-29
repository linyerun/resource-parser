package dy

const (
	videoBaseReqURL = "https://www.iesdouyin.com/share/video"
)

func generateVideoReqURLById(videoId string) string {
	return videoBaseReqURL + "/" + videoId
}
