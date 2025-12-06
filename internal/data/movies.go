package data

import (
	"greenlight.ilx.net/internal/validator"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, input *Movie) bool {
	v.Check(len(input.Title) > 3, "title", "tetle must be greater than 3 characters")
	v.Check(len(input.Title) < 18, "title", "title must be lower than 18 characters")
	v.Check(validator.Unique(input.Genres), "genres", "you should avoid using duplicate values for genres")
	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(input.Runtime != 0, "runtime", "must be provided")
	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

	return v.Valid()

}
