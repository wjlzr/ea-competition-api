package xstring

import (
	"strconv"
	"strings"
)

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

func Matching(s string) string {
	if ok := strings.Contains(s, "_"); ok {
		r := strings.Split(s, "_")
		s = r[0]
	}
	return s
}
