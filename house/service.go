package house

import (
	"encoding/json"
	"fmt"
	"home_vision/httpClients"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	housesEndpoint = "/api_project/houses"
)

type HouseService struct {
	HttpClient httpClients.HttpClient
	Domain string
}

func (h *HouseService) FetchHousesByPage(page int) ([]House, error) {
	fullUrl := h.Domain + housesEndpoint + "?page=" + strconv.Itoa(page)
	resp, err := h.HttpClient.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	// Parse json response
	var housesResponse HousesResponse
	if err := json.Unmarshal(body, &housesResponse); err != nil {
		return nil, err
	}

	return housesResponse.Houses, nil
}

func (h *HouseService) FetchHouseImage(house House)  ([]byte, error) {
	resp, err := h.HttpClient.Get(house.PhotoURL)	
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