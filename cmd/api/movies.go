package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.ilx.net/internal/data"
	"greenlight.ilx.net/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	v := validator.New()

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	if !data.ValidateMovie(v, movie) {
		app.failedValidationResponse(w, r, v.Errors)
		fmt.Fprintln(w, input.Title)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(w, r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	data := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casbalanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	app.writeJSON(w, r, http.StatusOK, envelope{"movie": data}, nil)

}
