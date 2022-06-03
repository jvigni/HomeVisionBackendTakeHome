package house

import (
	"fmt"
	"home_vision/httpClients"
	"home_vision/utils"
	"io"
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

func (s *HouseService) FetchHousesByPage(page int) ([]House, error) {
	fullUrl := s.Domain + housesEndpoint + "?page=" + strconv.Itoa(page)
	httpResponse, err := s.HttpClient.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	housesResponse, err := utils.ParseJson[HousesResponse](httpResponse)
	if err != nil {
		return nil, err
	}
	return housesResponse.Houses, nil
}

func (s *HouseService) FetchHouseImage(house House)  ([]byte, error) {
	resp, err := s.HttpClient.Get(house.PhotoURL)	
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