package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

const datetime_format = "2006-01-02 15:04:05"

var types = map[string]string{
	"album": "https://rss.itunes.apple.com/api/v1/jp/apple-music/top-albums/all/10/explicit.json",
	"song":  "https://rss.itunes.apple.com/api/v1/jp/apple-music/top-songs/all/10/explicit.json",
}

// historiesテーブル定義
type History struct {
	gorm.Model
	Id           int64 `gorm:"primary_key"`
	ApiUpdatedAt string
	ResourceType string
	ApiUrl       string
	Resources    []Resource
}

// resourcesテーブル定義
type Resource struct {
	gorm.Model
	Id         int64 `gorm:"primary_key"`
	HistoryId  uint
	Name       string
	Url        string
	ArtworkUrl string
	ArtistName string
	ArtistUrl  string
	Copyright  string
}

// Result アルバム/ソングの情報
type Result []struct {
	ArtistName string `json:"artistName"`    // artist name
	ArtistUrl  string `json:"artistUrl"`     // artist page URL
	ArtworkUrl string `json:"artworkUrl100"` // jacket picture URL
	Copyright  string `json:"copyright"`     // copyright
	Name       string `json:"name"`          // album/song name
	Url        string `json:"url"`           // album/song URL
}

// Lanking RSS Feedのアウトライン
type Lanking struct {
	Outline struct {
		Updated string `json:"updated"`
		ApiUrl  string `json:"id"`
		Results Result `json:"results"`
	} `json:"feed"`
}

type Lankings map[string]Lanking
type Histories map[string]History

func main() {
	lankings := make(Lankings)

	for k, v := range types {
		res, err := http.Get(v)

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

		var lanking Lanking
		if err := json.Unmarshal(body, &lanking); err != nil {
			fmt.Println(err)
			return
		}
		lankings[k] = lanking
	}

	db := gormConnect()
	defer db.Close()

	histories := make(Histories)

	for k, _ := range types {
		h := History{}
		db.Where("resource_type = ?", k).Last(&h)
		histories[k] = h
	}

	for resourceType, _ := range types {
		l := lankings[resourceType]
		h := histories[resourceType]

		apiUpdated := parseDatetime(l.Outline.Updated)
		updated := parseDatetime(h.ApiUpdatedAt)

		if apiUpdated == updated {
			continue
		}

		history := History{
			ApiUpdatedAt: apiUpdated,
			ResourceType: resourceType,
			ApiUrl:       l.Outline.ApiUrl,
		}
		db.Create(&history)

		for _, r := range l.Outline.Results {
			db.Create(&Resource{
				HistoryId:  history.Model.ID,
				Name:       r.Name,
				Url:        r.Url,
				ArtworkUrl: r.ArtworkUrl,
				ArtistName: r.ArtistName,
				ArtistUrl:  r.ArtistUrl,
				Copyright:  r.Copyright,
			})
		}
	}
}

func parseDatetime(datetime string) string {
	timestamp, err := time.Parse(time.RFC3339, datetime)

	if err != nil {
		fmt.Println(err)
		return "err"
	}

	return timestamp.Format(datetime_format)
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "app"
	OPTION := "charset=utf8&parseTime=True"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	return db
}
