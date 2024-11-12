// made by recanman
package main

import (
	"carsearch/pkg/geocoder"
	"carsearch/pkg/models"
	"carsearch/pkg/scraper"
	"carsearch/pkg/store"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func addSearchTest() {
	yearMin := 2010
	yearMax := 2020
	mileageMin := 0
	mileageMax := 100000
	priceMin := 0
	priceMax := 10000
	radius := 50

	s := models.Search{
		Platform:   "facebook",
		Location:   "Banff, AB, Canada",
		CarMake:    "Toyota",
		YearMin:    &yearMin,
		YearMax:    &yearMax,
		MileageMin: &mileageMin,
		MileageMax: &mileageMax,
		PriceMin:   &priceMin,
		PriceMax:   &priceMax,
		Radius:     &radius,
	}

	_, err := store.CreateSearch(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.ID)
}

func main() {
	listen := flag.String("listen", ":8000", "listen address")
	flag.Parse()

	if err := store.Initialize(); err != nil {
		panic(err)
	}
	if err := models.Initialize(); err != nil {
		panic(err)
	}

	geocoder.Initialize()

	r := SetupRoutes()
	go scraper.Start()

	addSearchTest()

	fmt.Printf("Listening on %s\n", *listen)
	log.Fatal(http.ListenAndServe(*listen, r))
}
