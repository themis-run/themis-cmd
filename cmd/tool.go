package cmd

import "time"

var timeLayout = "2006-01-02 15:04:05"

func transforTimestamp(timestamp int64) string {
	t := time.UnixMilli(timestamp)
	return t.Format(timeLayout)
}
