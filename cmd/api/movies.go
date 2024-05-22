package main

import (
	"fmt"
	"github.com/markian/greenlight/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movieData envelope
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "The Godfather",
		Year:      1975,
		Runtime:   189,
		Genres:    []string{"crime", "drama"},
		Version:   1,
	}

	movieData = envelope{"movie": movie}
	err = app.writeJSON(w, http.StatusOK, movieData, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server could not process the request", http.StatusInternalServerError)
	}
}
