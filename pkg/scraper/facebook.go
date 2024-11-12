// made by recanman
package scraper

import (
	"carsearch/pkg/geocoder"
	"carsearch/pkg/models"
	"carsearch/pkg/store"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func constructFacebookURL(search models.Search) (string, error) {
	base := "https://www.facebook.com/marketplace/category/vehicles?"

	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	if search.YearMin != nil {
		params.Add("minYear", fmt.Sprintf("%d", *search.YearMin))
	}
	if search.YearMax != nil {
		params.Add("maxYear", fmt.Sprintf("%d", *search.YearMax))
	}

	if search.MileageMin != nil {
		params.Add("minMileage", fmt.Sprintf("%d", *search.MileageMin))
	}
	if search.MileageMax != nil {
		params.Add("maxMileage", fmt.Sprintf("%d", *search.MileageMax))
	}

	if search.PriceMin != nil {
		params.Add("minPrice", fmt.Sprintf("%d", *search.PriceMin))
	}
	if search.PriceMax != nil {
		params.Add("maxPrice", fmt.Sprintf("%d", *search.PriceMax))
	}

	if search.Radius != nil {
		params.Add("radius", fmt.Sprintf("%d", *search.Radius))
	}
	if search.Location != "" {
		loc, err := geocoder.Geocode(search.Location)
		if err != nil {
			return "", err
		}

		params.Add("latitude", fmt.Sprintf("%.2f", loc.Lat))
		params.Add("longitude", fmt.Sprintf("%.2f", loc.Lng))
	}

	if search.CarMake != "Any" {
		params.Add("make", models.Makes[search.CarMake])
	}

	params.Add("sortBy", "creation_time_descend")
	u.RawQuery = params.Encode()
	return u.String(), nil
}

var LISTINGS_START = "\"edges\":[{"
var LISTINGS_END = ",\"page_info\":"

func addListing(listing models.Listing) {
	new, err := store.CreateListing(listing)
	if err != nil {
		fmt.Println("Error creating listing:", err)
	}

	if new {
		fmt.Println("New listing: https://www.facebook.com/marketplace/item/" + listing.ID)
	}

}

func FacebookMarketplace(search models.Search) {
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true

	c.OnHTML("body", func(e *colly.HTMLElement) {
		html := e.Text

		start := strings.Index(html, LISTINGS_START)
		end := strings.Index(html, LISTINGS_END)

		if start == -1 || end == -1 {
			fmt.Println("Could not find listings")
			return
		}
		if end < start {
			fmt.Println("End is before start")
			return
		}

		edgesData := "[{" + html[start+len(LISTINGS_START):end]

		var rawListings []models.RawListing
		if err := json.Unmarshal([]byte(edgesData), &rawListings); err != nil {
			fmt.Println("Error unmarshalling listings:", err)
			return
		}

		fmt.Println("Found", len(rawListings), "listings")
		for _, rawListing := range rawListings {
			listing := models.Listing{
				ID:                  rawListing.Node.Listing.ID,
				PrimaryListingPhoto: rawListing.Node.Listing.PrimaryListingPhoto.Image.URI,
				Amount:              rawListing.Node.Listing.ListingPrice.FormattedAmount,
				Location:            rawListing.Node.Listing.Location.ReverseGeocode.CityPage.DisplayName,
				Title:               rawListing.Node.Listing.MarketplaceListingTitle,
				IsSold:              rawListing.Node.Listing.IsSold,
			}

			addListing(listing)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("TE", "Trailers")
	})
	c.OnResponse(func(r *colly.Response) {
		addScrape(search.ID, r.Body)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)
	})

	url, err := constructFacebookURL(search)
	if err != nil {
		fmt.Println("Error constructing URL:", err)
	}

	if err = c.Visit(url); err != nil {
		fmt.Println("Error visiting Facebook Marketplace:", err)
	}
}

func Start() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Scraping...")
		allSearches := store.GetAllSearches()

		for _, search := range allSearches {
			FacebookMarketplace(search)
		}
	}
}
