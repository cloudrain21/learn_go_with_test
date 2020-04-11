package concurrency

import (
	"errors"
	"net/http"
	"time"
)

func WebsiteRacer(url1, url2 string) (winner string) {
	startA := time.Now()
	http.Get(url1)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(url2)
	bDuration := time.Since(startB)

	if aDuration > bDuration {
		winner = url1
	} else {
		winner = url2
	}
	return
}

var tenSecondTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <- pingUrl(url1):
		return url1, nil
	case <- pingUrl(url2):
		return url2, nil
	case <- time.After(timeout):
		return "", errors.New("timeout")
	}
}

func measureDuration(url string) (duration time.Duration) {
	startA := time.Now()
	http.Get(url)
	duration = time.Since(startA)
	return
}

func pingUrl(url string) (chan struct{}){
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}