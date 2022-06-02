package main

import (
	"fmt"
	"home_vision/api"
	"home_vision/models"
	"log"
	"os"
	"sync"
)

const (
	amountOfPagesToProcess = 10
	housesImageRepositoryPath = "houses_images/"
	imagesExtension = "jpg"
)

var downloadedImagesCount int
var successfullyProcessedHousesCount int


func main() {
	processPages(amountOfPagesToProcess)
}

func processPages(amountOfPages int) {
	log.Printf("Processing houses from pages 1 to %d", amountOfPages)
	var processPagesWG sync.WaitGroup
	for i := 1; i <= amountOfPages; i++ {
		processPagesWG.Add(1)
		go processHousesByPage(i, &processPagesWG)
	}
	processPagesWG.Wait()
	log.Printf("All available houses processed. Images downloaded: %d", downloadedImagesCount)
}

func processHousesByPage(page int, wg *sync.WaitGroup) {
	defer wg.Done()
	//log.Printf("Processing houses on page %d..", page)
	houses, err := api.FetchHouses(page)
	if err != nil {
		log.Printf("Unable to load page %d... %v", page, err)
		return
	} else {
		//log.Printf("Houses from page %d fetched successfully", page)
		var processHousesWG sync.WaitGroup
		for _, house := range houses { //ENCAPSULAR EN -CONCURRENTPROCESSHOUSES()
			processHousesWG.Add(1)
			go processHouse(house, &processHousesWG)
		}
		log.Printf("Page %d Done", page)
		processHousesWG.Wait()
	}
}

func processHouse(house models.House, wg *sync.WaitGroup) {
	defer wg.Done()
	err := downloadHouseImage(house)
	if err != nil {
		log.Printf("failed to process house %d... %v", house.Id, err)
	}
	successfullyProcessedHousesCount++
}

func downloadHouseImage(house models.House) error {
	fileName := fmt.Sprintf("%d-%s.%s", house.Id, house.Address, imagesExtension)
	respBytes, err := api.FetchHouseImage(house);
	if err != nil {
		return fmt.Errorf("can't fetch image on houseID: %d... %w", house.Id, err)
	}
	if err := createNewFile(respBytes, fileName, housesImageRepositoryPath); err != nil {
		return fmt.Errorf("can't create image file on houseID: %d... %w", house.Id, err)
	}

	downloadedImagesCount++
	return nil
}

func createNewFile(data []byte, fileName string, filePath string) error {
	err := os.WriteFile(filePath + fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}