package main

import (
	"backend/cmd/api/internals/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World! %s", app.Domain)

	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:version`
	}{
		Status:  "active",
		Message: "go to movies",
		Version: "1,0,0",
	}
	out, err := json.Marshal(payload)
	//mashal is used in encoding and decoding
	//of json, hence weve used it to decode the json above in the struct
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "applcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rd, _ := time.Parse("2006-07-02", "1986-07-07")
	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		MPPAARAting: "R",
		Runtime:     116,
		Description: "A very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, highlander)
	rd, _ = time.Parse("2006-07-02", "1986-07-07")

	rotla := models.Movie{
		ID:          2,
		Title:       "Raiders of the lost ark",
		ReleaseDate: rd,
		MPPAARAting: "pg-13",
		Runtime:     115,
		Description: "A very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, rotla)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
