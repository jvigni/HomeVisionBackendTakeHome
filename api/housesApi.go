package api

import (
	"encoding/json"
	"fmt"
	"home_vision/httpRetry"
	"home_vision/models"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	homevisionDomain = "http://app-homevision-staging.herokuapp.com"
	housesEndpoint = "/api_project/houses"
)

func FetchHouses(page int) ([]models.House, error) {
	fullUrl := homevisionDomain + housesEndpoint + "?page=" + strconv.Itoa(page)
	resp, err := httpRetry.Get(fullUrl)
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
	resp, err := httpRetry.Get(house.PhotoURL)	
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error code %v fetching houseID %v", resp.Status, house.Id)
	}
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}