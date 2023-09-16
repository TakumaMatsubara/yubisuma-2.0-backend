package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Animal response for single animal
type animalResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	NumFinger     int    `json:"numFinger"`
	Skill         string `json:"skill"`
	HandUrl       string `json:"handUrl"`
	UpFingerUrl   string `json:"upFingerUrl"`
	DownFingerUrl string `json:"downFingerUrl"`
}

// Animals response for list of animals
type animalsResponse struct {
	Animals []animalResponse `json:"animals"`
}

// HandleAnimalsGet handles the GET request for all animals
func HandleAnimalsGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		animals, err := GetAnimals()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch animals",
			})
		}

		if len(animals) == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No animals found",
			})
		}

		// Convert database animals to response format
		resp := animalsResponse{}
		for _, animal := range animals {
			resp.Animals = append(resp.Animals, animalResponse{
				ID:            animal.ID,
				Name:          animal.Name,
				NumFinger:     animal.NumFinger,
				Skill:         animal.Skill,
				HandUrl:       animal.HandUrl,
				UpFingerUrl:   animal.UpFingerUrl,
				DownFingerUrl: animal.DownFingerUrl,
			})
		}

		return c.JSON(http.StatusOK, resp)
	}
}
