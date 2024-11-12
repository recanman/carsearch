// made by recanman
package store

import (
	"carsearch/pkg/models"
)

func saveListingsToFile() error {
	return saveToFile(listingsFileName, listings)
}

func CreateListing(listing models.Listing) (bool, error) {
	for _, l := range listings {
		if l.ID == listing.ID {
			// Skip if listing already exists
			return false, nil
		}
	}

	listings = append(listings, listing)
	err := saveListingsToFile()

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetListings() []models.Listing {
	return listings
}
