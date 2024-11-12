// made by recanman
package main

import (
	"carsearch/pkg/models"
	"carsearch/pkg/store"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

const CONTENT_TYPE = "Content-Type"
const CONTENT_TYPE_JSON = "application/json"

const SEARCH_NOT_FOUND = "Search not found"

var validate = validator.New()

func CreateSearch(w http.ResponseWriter, r *http.Request) {
	var search models.Search
	if err := json.NewDecoder(r.Body).Decode(&search); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(search); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if search.Platform != "facebook" {
		http.Error(w, "Invalid platform, only 'facebook' is allowed", http.StatusBadRequest)
		return
	}

	if search.CarMake != "Any" {
		if _, exists := models.Makes[search.CarMake]; !exists {
			http.Error(w, "Invalid car make", http.StatusBadRequest)
			return
		}
	}

	if search.YearMin != nil && search.YearMax != nil && *search.YearMin > *search.YearMax {
		http.Error(w, "Year min must be less than year max", http.StatusBadRequest)
		return
	}
	if search.MileageMin != nil && search.MileageMax != nil && *search.MileageMin > *search.MileageMax {
		http.Error(w, "Mileage min must be less than mileage max", http.StatusBadRequest)
		return
	}
	if search.PriceMin != nil && search.PriceMax != nil && *search.PriceMin > *search.PriceMax {
		http.Error(w, "Price min must be less than price max", http.StatusBadRequest)
		return
	}

	id, err := store.CreateSearch(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_JSON)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	search, exists := store.GetSearch(id)
	if !exists {
		http.Error(w, SEARCH_NOT_FOUND, http.StatusNotFound)
		return
	}

	w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_JSON)
	json.NewEncoder(w).Encode(search)
}

func UpdateSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedSearch models.Search
	if err := json.NewDecoder(r.Body).Decode(&updatedSearch); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedSearch.Platform != "facebook" {
		http.Error(w, "Invalid platform, only 'facebook' is allowed", http.StatusBadRequest)
		return
	}

	if err := store.UpdateSearch(id, updatedSearch); err != nil {
		http.Error(w, SEARCH_NOT_FOUND, http.StatusNotFound)
		return
	}

	w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_JSON)
	json.NewEncoder(w).Encode(updatedSearch)
}

func DeleteSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := store.DeleteSearch(id); err != nil {
		http.Error(w, SEARCH_NOT_FOUND, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllSearches(w http.ResponseWriter, r *http.Request) {
	allSearches := store.GetAllSearches()
	w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_JSON)
	json.NewEncoder(w).Encode(allSearches)
}

func GetAllMakes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_JSON)
	json.NewEncoder(w).Encode(models.Makes)
}
