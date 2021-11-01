package xtime

import "time"

func Second() int64 {
	return time.Now().Unix()
}
