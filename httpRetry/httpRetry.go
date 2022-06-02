package httpRetry

import (
	"errors"
	"log"
	"net/http"
	"time"
)

const (
	maxAttempts = 5
	attemptIntervalMilliseconds = 200
)

/// Tries multiple times an http.Get request.
/// Returns error after all posible atempts failed
func Get(url string) (*http.Response, error){
	attempts := 0
	for {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		attempts++
		log.Printf("%s [%d/%d attempts] on %s", resp.Status, attempts, maxAttempts, url)
		if attempts == maxAttempts {
			return nil, errors.New("maximum amount of attempts reached")
		}
		time.Sleep(time.Duration(attemptIntervalMilliseconds) * time.Millisecond)
	}
}