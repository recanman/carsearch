// made by recanman
// Technically not thread-safe, but it doesn't matter
package store

import (
	"carsearch/pkg/models"
	"encoding/json"
	"os"
)

var searches = make(map[string]models.Search)
var listings = make([]models.Listing, 0)

const searchesFileName = "searches.json"
const listingsFileName = "listings.json"

func loadFromFile(fileName string, data interface{}, initialValue []byte) error {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(fileName)
			if err == nil {
				file.Write(initialValue)
			}
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}
	return nil
}

func saveToFile(fileName string, data interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func Initialize() error {
	if err := loadFromFile(searchesFileName, &searches, []byte("{}")); err != nil {
		return err
	}
	if err := loadFromFile(listingsFileName, &listings, []byte("[]")); err != nil {
		return err
	}

	return nil
}
