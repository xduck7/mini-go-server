package entity

import "time"

type Invention struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Inventor    string    `json:"inventor"`
}
