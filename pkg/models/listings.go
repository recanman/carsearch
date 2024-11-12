// made by recanman
package models

type RawListing struct {
	Node struct {
		StoryType string `json:"story_type"`
		StoryKey  string `json:"story_key"`
		Tracking  string `json:"tracking"`
		Listing   struct {
			ID                  string `json:"id"`
			PrimaryListingPhoto struct {
				Image struct {
					URI string `json:"uri"`
				} `json:"image"`
				ID string `json:"id"`
			} `json:"primary_listing_photo"`
			IsMarketplaceListingRenderable string `json:"__isMarketplaceListingRenderable"`
			ListingPrice                   struct {
				FormattedAmount            string `json:"formatted_amount"`
				AmountWithOffsetInCurrency string `json:"amount_with_offset_in_currency"`
				Amount                     string `json:"amount"`
			} `json:"listing_price"`
			StrikethroughPrice                      string `json:"strikethrough_price"`
			IsMarketplaceListingWithComparablePrice string `json:"__isMarketplaceListingWithComparablePrice"`
			ComparablePrice                         string `json:"comparable_price"`
			ComparablePriceType                     string `json:"comparable_price_type"`
			Location                                struct {
				ReverseGeocode struct {
					City     string `json:"city"`
					State    string `json:"state"`
					CityPage struct {
						DisplayName string `json:"display_name"`
						ID          string `json:"id"`
					} `json:"city_page"`
				} `json:"reverse_geocode"`
			} `json:"location"`
			IsHidden                          bool   `json:"is_hidden"`
			IsLive                            bool   `json:"is_live"`
			IsPending                         bool   `json:"is_pending"`
			IsSold                            bool   `json:"is_sold"`
			IsViewerSeller                    bool   `json:"is_viewer_seller"`
			MinListingPrice                   string `json:"min_listing_price"`
			MaxListingPrice                   string `json:"max_listing_price"`
			MarketplaceListingCategoryID      string `json:"marketplace_listing_category_id"`
			MarketplaceListingTitle           string `json:"marketplace_listing_title"`
			CustomTitle                       string `json:"custom_title"`
			CustomSubTitlesWithRenderingFlags []struct {
				Subtitle string `json:"subtitle"`
			} `json:"custom_sub_titles_with_rendering_flags"`
			OriginGroup string `json:"origin_group"`
			//ListingVideo                            string  `json:"listing_video"`
			IsMarketplaceListingWithChildListings   string   `json:"__isMarketplaceListingWithChildListings"`
			ParentListing                           string   `json:"parent_listing"`
			MarketplaceListingSeller                string   `json:"marketplace_listing_seller"`
			IsMarketplaceListingWithDeliveryOptions string   `json:"__isMarketplaceListingWithDeliveryOptions"`
			DeliveryTypes                           []string `json:"delivery_types"`
		} `json:"listing"`
		ID string `json:"id"`
	} `json:"node"`
	Cursor string `json:"cursor"`
}

type Listing struct {
	ID                  string
	PrimaryListingPhoto string
	Amount              string
	Location            string
	Title               string
	IsSold              bool
}
