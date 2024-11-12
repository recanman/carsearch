// made by recanman
package store

import (
	"carsearch/pkg/models"
)

const SAVE_ERROR = "Error saving to file:"

func saveSearchesToFile() error {
	return saveToFile(searchesFileName, searches)
}

func CreateSearch(search models.Search) (string, error) {
	id, err := models.NewID()
	if err != nil {
		return "", err
	}

	search.ID = id
	searches[search.ID] = search

	err = saveSearchesToFile()
	return search.ID, err
}

func GetSearch(id string) (models.Search, bool) {
	search, ok := searches[id]
	return search, ok
}

func UpdateSearch(id string, search models.Search) error {
	searches[id] = search
	return saveSearchesToFile()
}

func DeleteSearch(id string) error {
	delete(searches, id)
	return saveSearchesToFile()
}

func GetAllSearches() []models.Search {
	allSearches := make([]models.Search, 0, len(searches))
	for _, search := range searches {
		allSearches = append(allSearches, search)
	}
	return allSearches
}
