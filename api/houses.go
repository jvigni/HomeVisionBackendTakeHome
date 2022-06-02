package api

import (
	"encoding/json"
	"errors"
	"home_vision/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	homevisionDomain = "http://app-homevision-staging.herokuapp.com"
	housesEndpoint = "/api_project/houses"
	housesPerPage = 10
	requestMaxAttempts = 5
	attemptIntervalMilliseconds = 200
)

func FetchHouses(page int) ([]models.House, error) {
	fullUrl := homevisionDomain + housesEndpoint + "?page=" + strconv.Itoa(page) + "&per_page=" + strconv.Itoa(housesPerPage)
	resp, err := tryGet(fullUrl, requestMaxAttempts, attemptIntervalMilliseconds)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var housesResponse models.HousesResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	// Parse json response
	if err := json.Unmarshal(body, &housesResponse); err != nil {
		return nil, err
	}

	return housesResponse.Houses, nil
}

func FetchHouseImage(house models.House)  ([]byte, error) {
	resp, err := http.Get(house.PhotoURL)	
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

/// Tries multiple times an http.Get request.
/// Returns error after all posible atempts failed
func tryGet(url string, maxAttempts int, intervalMilliseconds int) (resp *http.Response, err error) {
	attempts := 0
	for {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		attempts++
		log.Printf("%s (%d/%d attempts) on %s", resp.Status, attempts, maxAttempts, url)
		if attempts == maxAttempts {
			return nil, errors.New("maximum amount of attempts reached")
		}
		time.Sleep(time.Duration(intervalMilliseconds) * time.Millisecond)
	}
}
