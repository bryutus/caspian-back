package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/bryutus/caspian-serverside/app/conf"
	"github.com/bryutus/caspian-serverside/app/db"
	"github.com/bryutus/caspian-serverside/app/models"
	"github.com/labstack/echo"
)

type (
	Resource struct {
		Collection Collection `json:"collection"`
	}
	Collection struct {
		Title   string `json:"title"`
		Updated string `json:"updated"`
		Items   []Item `json:"items"`
	}
	Item struct {
		Name       string `json:"name"`
		Url        string `json:"url"`
		ArtworkUrl string `json:"artworkUrl"`
		ArtistName string `json:"artistName"`
		ArtistUrl  string `json:"artistUrl"`
		Copyright  string `json:"copyright"`
	}

	Err struct {
		Info ErrInfo `json:"error"`
	}
	ErrInfo struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func GetResources(resource string) echo.HandlerFunc {
	return func(c echo.Context) error {

		logfile, err := os.OpenFile(conf.GetLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return c.JSONPretty(http.StatusOK, Err{ErrInfo{Code: 80, Message: "Failed to file operation"}}, "  ")
		}
		defer logfile.Close()

		log.SetOutput(io.Writer(logfile))
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

		limit := c.QueryParam("limit")
		if limit != "" {
			if err := isNumeric(limit); err != nil {
				log.Printf("[INFO] %s", err.Error())
				limit = ""
			}
		}

		db, err := db.Connect()
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return c.JSONPretty(http.StatusOK, Err{ErrInfo{Code: 90, Message: "Failed to database operation"}}, "  ")
		}
		defer db.Close()

		h := models.History{}
		if err := db.Where("resource_type = ?", resource).Last(&h).Error; err != nil {
			if err.Error() != "record not found" {
				log.Printf("[ERROR] Database: `%s`", err.Error())
				return c.JSONPretty(http.StatusOK, Err{ErrInfo{Code: 91, Message: "Failed to database operation"}}, "  ")
			}
		}

		r := []models.Resource{}
		if err := db.Model(&h).Order("id").Limit(limit).Related(&r).Error; err != nil {
			if err.Error() != "record not found" {
				log.Printf("[ERROR] Database: %s", err.Error())
				return c.JSONPretty(http.StatusOK, Err{ErrInfo{Code: 92, Message: "Failed to database operation"}}, "  ")
			}
		}

		data := createResponseBody(resource, &h, &r)

		return c.JSONPretty(http.StatusOK, data, "  ")
	}
}

func isNumeric(str string) (err error) {
	if matched, _ := regexp.MatchString(`[0-9]`, str); !matched {
		return fmt.Errorf("limit: Invalid value `%s` is specified", str)
	}

	return nil
}

func createResponseBody(title string, h *models.History, r *[]models.Resource) (data *Resource) {
	var items []Item

	for _, v := range *r {
		t := Item{}
		t.Name = v.Name
		t.Url = v.Url
		t.ArtworkUrl = v.ArtworkUrl
		t.ArtistName = v.ArtistName
		t.ArtistUrl = v.ArtistUrl
		t.Copyright = v.Copyright
		items = append(items, t)
	}

	return &Resource{
		Collection{
			Title:   title,
			Updated: h.ApiUpdatedAt,
			Items:   items,
		},
	}
}
