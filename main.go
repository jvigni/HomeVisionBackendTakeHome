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
	pagesAmountToProcess = 10
	housePhotoRepositoryPath = "houses_photos/"
	photosExtension = "jpg"
)

func main() {
	processAllHouses(pagesAmountToProcess)
}

func processAllHouses(amountOfPages int) {
	var pagesWG sync.WaitGroup
	for i := 1; i <= amountOfPages; i++ {
		pagesWG.Add(1)
		go processHousesByPage(i, &pagesWG)
	}
	pagesWG.Wait()
	log.Printf("All available houses processed")
}

func processHousesByPage(page int, housesWG *sync.WaitGroup) {
	log.Printf("Processing houses on page %d..", page)
	houses, err := api.FetchHouses(page)
	if err != nil {
		log.Printf("Unable to load page %d", page) //TODO CHEKEAR
		return
	} else {
		log.Printf("Houses from page %d fetched successfully", page)
		var housesWG sync.WaitGroup
		for _, house := range houses { //ENCAPSULAR EN -CONCURRENTPROCESSHOUSES()
			housesWG.Add(1)
			go processHouse(house, &housesWG)
		}
		log.Printf("Houses from page %d processed successfully", page)
		housesWG.Wait()
	}
}

func processHouse(house models.House, housesWG *sync.WaitGroup) {
	defer housesWG.Done()
	downloadHousePhoto(house)
}

func downloadHousePhoto(house models.House) {
	fileName := fmt.Sprintf("%d-%s.%s", house.Id, house.Address, photosExtension)
	respBytes, err := api.FetchHouseImage(house)
	if err != nil {
		log.Printf("Cant fetch photo on houseId: %d [%s]", house.Id, err)
	}
	if err := createNewFile(respBytes, fileName, housePhotoRepositoryPath); err != nil {
		log.Printf("Cant create photo file for houseId: %d [%s]", house.Id, err)	
	}
}

func createNewFile(data []byte, fileName string, filePath string) error {
	err := os.WriteFile(filePath + fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}