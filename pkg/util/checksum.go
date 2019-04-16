package util

import (
	"regexp"
)

func GetCheckSum(message string) string {
	r := regexp.MustCompile(`original:\w+,(\w+),\w+`)
	res := r.FindStringSubmatch(message)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}
