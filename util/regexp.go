package util

import (
	"fmt"
	"regexp"
)

func RegexpMatchUrlFromString(str string) (string, error) {
	//urlReg, err := regexp.Compile(`https?://[\w.-]+[\w/-]*[\w.-:]*\??[\w=&:\-+%.]*/*`)
	pattern := `https://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[^?\s]*)?(\?[^\s]*)?`
	urlReg, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("match url regexp compile error: %s", err.Error())
	}

	findStr := urlReg.FindString(str)
	if len(findStr) <= 0 {
		return "", fmt.Errorf("str not have url")
	}

	return findStr, nil
}
