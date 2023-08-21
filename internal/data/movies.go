package data

import (
	"time"
)

type Movie struct {
	ID 			int64 		`json:"id"`
	CreatedAt 	time.Time 	`json:"-"`
	Title 		string 		`json:"title"`
	Year 		int32 		`json:"year,omitempty"`
	Runtime 	int32 		`json:"runtime,ompitempty"`// Runtime is measured in minutes
	Genres 		[]string 	`json:"genres,omitempty"`
	Version 	int32 		`json:"version"`
}