// made by recanman
package main

import (
	"carsearch/pkg/geocoder"
	"carsearch/pkg/models"
	"carsearch/pkg/notify"
	"carsearch/pkg/scraper"
	"carsearch/pkg/store"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func listingToEmbed(l models.Listing) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Title: l.Title,
		Description: fmt.Sprintf(
			"[%s](%s)",
			l.Title,
			"https://www.facebook.com/marketplace/item/"+l.ID,
		),
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Marketplace search bot",
			URL:  "https://github.com/recanman",
		},
		URL: "https://www.facebook.com/marketplace/item/" + l.ID,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Price",
				Value:  l.Amount,
				Inline: true,
			},
			{
				Name:   "Distance",
				Value:  l.Mileage,
				Inline: true,
			},
			{
				Name:   "Location",
				Value:  l.Location,
				Inline: true,
			},
			{
				Name:   "Sold?",
				Value:  fmt.Sprintf("%t", l.IsSold),
				Inline: true,
			},
		},
	}

	if l.PrimaryListingPhoto != "" {
		embed.Image = &discordgo.MessageEmbedImage{
			URL: l.PrimaryListingPhoto,
		}
	}

	return embed
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

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		panic("DISCORD_TOKEN environment variable not set")
	}
	channelID := os.Getenv("DISCORD_CHANNEL_ID")
	if channelID == "" {
		panic("DISCORD_CHANNEL_ID environment variable not set")
	}

	notifier, err := notify.NewDiscordNotifier(os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_CHANNEL_ID"))
	if err != nil {
		fmt.Println("Error creating Discord notifier")
	}

	geocoder.Initialize()

	r := SetupRoutes()
	go scraper.Start()

	go func() {
		for n := range scraper.Notifications {
			fmt.Println("New listing:", n.ID)
			embed := listingToEmbed(n)
			notifier.Notify(embed)
		}
	}()

	go func() {
		for f := range scraper.Failure {
			fmt.Println("Error:", f)
			notifier.NotifyError(f)

		}
	}()

	fmt.Printf("Listening on %s\n", *listen)
	log.Fatal(http.ListenAndServe(*listen, r))
}
