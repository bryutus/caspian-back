package handler

import (
	"fmt"
	"net/http"
	"regexp"

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
)

func GetResources(resource string) echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		if limit != "" {
			if err := isNumeric(limit); err != nil {
				return c.JSON(http.StatusOK, "Invalid value")
			}
		}

		db := db.Connect()
		defer db.Close()

		h := models.History{}
		if db.Where("resource_type = ?", resource).Last(&h).RecordNotFound() {
			return c.JSON(http.StatusOK, "not found")
		}

		r := []models.Resource{}
		if db.Model(&h).Order("id").Limit(limit).Related(&r).RecordNotFound() {
			return c.JSON(http.StatusOK, "not found")
		}

		data := createResponseBody(&h, &r)

		return c.JSONPretty(http.StatusOK, data, "  ")
	}
}

func isNumeric(str string) (err error) {
	if matched, _ := regexp.MatchString(`[0-9]`, str); !matched {
		return fmt.Errorf("limit: Invalid value `%s` is specified", str)
	}

	return nil
}

func createResponseBody(h *models.History, r *[]models.Resource) (data *Resource) {
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
			Title:   h.ResourceType,
			Updated: h.ApiUpdatedAt,
			Items:   items,
		},
	}
}
