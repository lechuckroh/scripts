package timeutil

import "time"

// FromMillis 함수는 주어진 밀리초 값을 기준으로 time.Time 객체를 생성합니다.
// 매개변수:
//
//	millis - Unix 기준 밀리초(millisecond) 값.
//
// 반환값:
//
//	time.Time - millis를 기준으로 생성된 시간 객체.
func FromMillis(millis int64) time.Time {
	seconds := millis / 1000
	nanos := (millis % 1000) * int64(time.Millisecond)
	return time.Unix(seconds, nanos)
}
