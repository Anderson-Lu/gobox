package string_helper

import "regexp"

var urlRegexp *regexp.Regexp

func init() {
	urlRegexp = regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
}

/*
* Find urls from specific string
 */
func FindUrl(raw string) []string {
	return urlRegexp.FindAllString(raw, -1)
}
