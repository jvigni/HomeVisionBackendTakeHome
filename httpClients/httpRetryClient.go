package httpClients

import (
	"errors"
	"log"
	"net/http"
	"time"
)

const (
	baseRetryIntervalMilliseconds = 200
	backoffRetryIntervalMultiplier = 2
)

type HttpRetryClient struct {
	MaxAttempts int
}

/// Tries multiple times an http.Get request.
/// Returns error after all posible atempts failed.
/// Uses a simple exponential backoff.
func (c HttpRetryClient) Get(url string) (*http.Response, error){
	attempts := 0
	sleepTimeMilliseconds := baseRetryIntervalMilliseconds
	for {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		attempts++
		log.Printf("%s [%d/%d attempts] on %s", resp.Status, attempts, c.MaxAttempts, url)
		if attempts == c.MaxAttempts {
			return nil, errors.New("maximum amount of attempts reached")
		}
		time.Sleep(time.Duration(sleepTimeMilliseconds) * time.Millisecond)
		sleepTimeMilliseconds *= backoffRetryIntervalMultiplier
	}
}
