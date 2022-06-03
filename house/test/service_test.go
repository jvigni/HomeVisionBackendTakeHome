package testHouse

import (
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

func TestFetchTwoHousesAndReturnSizeTwo(t *testing.T) {
	responseBody, err := utils.LoadJson("getHousesResponseTest.json")
	if err != nil {
		t.Error(err)
	}

	url := "testDomain/api_project/houses?page=1"
	httpMockClient.Simulate(url, 200, responseBody)

	houses, err := houseService.FetchHousesByPage(1)
	if err != nil {
		t.Error(err)
	}
	if len(houses) != 2 {
		t.Errorf("response len must be 2, got %d", len(houses))
	}
}
