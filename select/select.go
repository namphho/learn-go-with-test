package _select

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(link string) time.Duration{
	start := time.Now()
	_, _ = http.Get(link)
	return time.Since(start)
}



func ping(link string) chan struct{}{
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(link)
		close(ch)
	}()
	return ch
}