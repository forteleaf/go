package SplitLineRe

import "regexp"

func SplitLineRe(raw string) []string {
	return regexp.MustCompile("((\r)\n)").Split(raw, -1)
}
