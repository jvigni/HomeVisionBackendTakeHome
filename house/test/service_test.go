package testHouse

import (
	"bytes"
	"home_vision/house"
	"home_vision/httpClients"
	"home_vision/utils"
	"testing"
)

var httpMockClient = httpClients.NewHttpMockClient()
var houseService = house.HouseService{ 
	HttpClient: httpMockClient,
	Domain: "testDomain",
}

func TestFetchHousesByPage(t *testing.T) {
	// Given
	responseBody, err := utils.LoadJson("getHousesResponseTest.json")
	if err != nil {
		t.Error(err)
	}
	url := "testDomain/api_project/houses?page=1"
	httpMockClient.Simulate(url, 200, responseBody)

	// When
	houses, err := houseService.FetchHousesByPage(1)
	
	// Then
	if err != nil {
		t.Error(err)
	}
	if len(houses) != 2 {
		t.Errorf("response len must be 2, got %d", len(houses))
	}
}

func TestFetchHouseImage(t *testing.T) {
	// Given
	imgUrl := "https://test/test.jpg"
	imgBytes := []byte{1,2,3}
	house := house.House{ PhotoURL: imgUrl }
	httpMockClient.Simulate(imgUrl, 200, string(imgBytes))

	// When
	resp, err := houseService.FetchHouseImage(house)
	
	// Then
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(resp, imgBytes) {
		t.Errorf("expected %v got %v", imgBytes, resp)
	}
}