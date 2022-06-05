package util

import "time"

func NowTimestampMS() int64 {
	return time.Now().UnixNano() / 1e6
}

func NowTimestamp() int64 {
	return time.Now().Unix()
}

func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
