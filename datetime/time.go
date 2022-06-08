package datetime

import "time"

func NowSec() int64 {
	return time.Now().Unix()
}

func NowNanoSec() int64 {
	return time.Now().UnixNano()
}

func NowMillSec() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func MorningDateTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func MiddayDateTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
}

func NightDateTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}
