package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result []struct {
	ArtistName string `json:"artistName"`
	ArtistURL  string `json:"artistUrl"`
	ArtworkURL string `json:"artworkUrl100"`
	Copyright  string `json:"copyright"`
	Name       string `json:"name"`
	URL        string `json:"url"`
}

type Lanking struct {
	Outline struct {
		Updated string `json:"updated"`
		Results Result `json:"results"`
	} `json:"feed"`
}

func main() {
	res, err := http.Get("https://rss.itunes.apple.com/api/v1/jp/apple-music/top-albums/all/10/explicit.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	execute(res)
}

func execute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var lanking Lanking
	if err := json.Unmarshal(body, &lanking); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(lanking.Outline.Updated)
	for _, r := range lanking.Outline.Results {
		fmt.Printf("%s : %s : %s : %s : %s : %s\n",
			r.ArtistName, r.ArtistURL, r.Name, r.URL, r.ArtworkURL, r.Copyright)
	}
}
