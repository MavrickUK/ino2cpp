package utils

import (
	"regexp"
)

const (
	regEx = `(?m)[\\\/:"*?<>|\x00-\x1F]+`
)

func RemoveInvalidFilenameChars(fn string) string {
	var re = regexp.MustCompile(regEx)
	var str = fn
	var substitution = "_"

	return re.ReplaceAllString(str, substitution)
}
