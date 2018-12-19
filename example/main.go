package main

import (
	"fmt"

	lastfm "github.com/irlndts/go-lastfm"
)

func main() {
	client, err := lastfm.New("")
	if err != nil {
		fmt.Printf("failed to init client: %s", err)
		return
	}

	top, err := client.User.TopArtists("kleto4kin", "7day", 1, 0)
	if err != nil {
		fmt.Printf("failed to get top: %s", err)
		return
	}

	if top.Total != 0 {
		for _, artist := range top.Artists {
			fmt.Println(artist)
		}
	}
}
