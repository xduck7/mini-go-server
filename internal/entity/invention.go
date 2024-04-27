package entity

import "time"

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"gte=18,lte=100"`
	Email     string `json:"email" validate:"required,email"`
}

type Invention struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"min=3" validate:"is_ok"`
	Date        time.Time `json:"date"`
	Description string    `json:"description" binding:"min=3,max=300"`
	Inventor    Person    `json:"inventor" binding:"required"`
}
