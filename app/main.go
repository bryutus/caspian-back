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
	ArtistURL  string `json:"artistUrl"`     // artist page URL
	ArtworkURL string `json:"artworkUrl100"` // jacket picture URL
	Copyright  string `json:"copyright"`     // copyright
	Name       string `json:"name"`          // album/song name
	URL        string `json:"url"`           // album/song URL
}

// Lanking RSS Feedのアウトライン
type Lanking struct {
	Outline struct {
		Updated string `json:"updated"`
		ApiUrl  string `json:"id"`
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

	apiUpdated := parseDatetime(lanking.Outline.Updated)

	db := gormConnect()
	defer db.Close()

	history := History{}
	db.Where("resource_type = ?", "album").Last(&history)

	updated := parseDatetime(history.ApiUpdatedAt)

	if apiUpdated == updated {
		return
	}

	history = History{}
	history.ApiUpdatedAt = apiUpdated
	history.ResourceType = "album"
	history.ApiUrl = lanking.Outline.ApiUrl

	db.Create(&history)

	for _, r := range lanking.Outline.Results {
		db.Create(&Resource{
			HistoryId:  history.Model.ID,
			Name:       r.ArtistName,
			Url:        r.URL,
			ArtworkUrl: r.ArtworkURL,
			ArtistName: r.ArtistName,
			ArtistUrl:  r.ArtistURL,
			Copyright:  r.Copyright,
		})
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
