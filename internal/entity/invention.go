package entity

import (
	_ "github.com/xduck7/mini-go-server/docs"
	"time"
)

type Person struct {
	FirstName string `json:"firstname" binding:"required" example:"John"`
	LastName  string `json:"lastname" binding:"required" example:"Doe"`
	Age       int    `json:"age" binding:"gte=18,lte=100" example:"30"`
	Email     string `json:"email" validate:"required,email" example:"john.doe@example.com"`
}

type Invention struct {
	ID          string    `json:"id" example:"123"`
	Title       string    `json:"title" binding:"min=3" validate:"is_ok" example:"Amazing Invention"`
	Date        time.Time `json:"date" example:"2024-04-30T00:00:00Z"`
	Description string    `json:"description" binding:"min=3,max=300" example:"This is an amazing invention that will revolutionize the world."`
	Inventor    Person    `json:"inventor" binding:"required"`
}
