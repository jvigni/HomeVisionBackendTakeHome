package main

import (
	"fmt"
	"home_vision/api"
	"home_vision/models"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

const (
	amountOfPagesToProcess = 10
	housesImageRepositoryPath = "houses_images/"
	imagesExtension = "jpg"
)

var downloadedImagesCounter int64
var successfullyProcessedHousesCounter int64

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
	log.Printf("All available houses processed. Images downloaded: %d", downloadedImagesCounter)
}

func processHousesByPage(page int, wg *sync.WaitGroup) {
	defer wg.Done()
	houses, err := api.FetchHouses(page)
	if err != nil {
		log.Printf("Unable to fetch page %d... %v", page, err)
		return
	} else {
		var processHousesWG sync.WaitGroup
		for _, house := range houses {
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
		log.Printf("Failed to process house %d... %v", house.Id, err)
	} else {
		atomic.AddInt64(&successfullyProcessedHousesCounter, 1)
	}
}

func downloadHouseImage(house models.House) error {
	fileName := fmt.Sprintf("%d-%s.%s", house.Id, house.Address, imagesExtension)
	respBytes, err := api.FetchHouseImage(house);
	if err != nil {
		return fmt.Errorf("can't fetch image... %w", err)
	}
	if err := createNewFile(respBytes, fileName, housesImageRepositoryPath); err != nil {
		return fmt.Errorf("can't create image file... %w", err)
	}
	atomic.AddInt64(&downloadedImagesCounter, 1)
	return nil
}

func createNewFile(data []byte, fileName string, filePath string) error {
	err := os.WriteFile(filePath + fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}