package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/bryutus/caspian-serverside/app/conf"
	"github.com/bryutus/caspian-serverside/app/db"
	"github.com/bryutus/caspian-serverside/app/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const datetimeFormat = "2006-01-02 15:04:05"

type Rsource []struct {
	ArtistName string `json:"artistName"`    // artist name
	ArtistURL  string `json:"artistUrl"`     // artist page URL
	ArtworkURL string `json:"artworkUrl100"` // jacket picture URL
	Copyright  string `json:"copyright"`     // copyright
	Name       string `json:"name"`          // album/song name
	URL        string `json:"url"`           // album/song URL
}

type Feed struct {
	Outline struct {
		Updated  string  `json:"updated"`
		APIURL   string  `json:"id"`
		Rsources Rsource `json:"results"`
	} `json:"feed"`
}

type FeedMap map[string]Feed
type HistoryMap map[string]models.History

func main() {
	// ロギングの設定
	logfile, err := os.OpenFile(conf.GetLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	defer logfile.Close()

	log.SetOutput(io.Writer(logfile))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	feeds := make(FeedMap)

	var waitGroup sync.WaitGroup

	types := conf.GetAppleApis()

	for k, v := range types {
		waitGroup.Add(1)

		go func(resource, apiUrl string) {
			defer waitGroup.Done()

			res, err := http.Get(apiUrl)

			if err != nil {
				fmt.Println(err)
				return
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			var feed Feed
			if err := json.Unmarshal(body, &feed); err != nil {
				fmt.Println(err)
				return
			}
			feeds[resource] = feed
		}(k, v)
	}

	waitGroup.Wait()

	db, err := db.Connect()
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		os.Exit(0)
	}
	defer db.Close()

	histories := make(HistoryMap)

	for k := range types {
		h := models.History{}
		db.Where("resource_type = ?", k).Last(&h)
		histories[k] = h
	}

	for resource := range types {
		f := feeds[resource]
		h := histories[resource]

		apiUpdated, _ := parseDatetime(f.Outline.Updated)
		historyUpdated, _ := parseDatetime(h.ApiUpdatedAt)

		if apiUpdated == historyUpdated {
			continue
		}

		history := models.History{
			ApiUpdatedAt: apiUpdated,
			ResourceType: resource,
			ApiUrl:       f.Outline.APIURL,
		}
		db.Create(&history)

		for _, r := range f.Outline.Rsources {
			db.Create(&models.Resource{
				HistoryId:  history.Model.ID,
				Name:       r.Name,
				Url:        r.URL,
				ArtworkUrl: r.ArtworkURL,
				ArtistName: r.ArtistName,
				ArtistUrl:  r.ArtistURL,
				Copyright:  r.Copyright,
			})
		}
	}
}

func parseDatetime(datetime string) (string, error) {

	if datetime == "" {
		return datetimeFormat, nil
	}

	timestamp, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return datetime, err
	}

	return timestamp.Format(datetimeFormat), nil
}
