package millis

import "time"

func NowInMillis() int64 {
	now := time.Now().Truncate(time.Millisecond)
	return now.UnixNano() / int64(time.Millisecond)
}

func NowInMillisAddHours(hourOffset int64) int64 {
	now := NowInMillis()
	delta := int64(time.Hour) * hourOffset / int64(time.Millisecond)
	return now + delta
}

func InMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func IntMillisInSeconds(t int64) int64 {
	return t / int64(time.Second/time.Millisecond)
}
