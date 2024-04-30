package entity

import (
	_ "github.com/xduck7/mini-go-server/docs"
	"time"
)

type Inventor struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"type:varchar(32)" json:"firstname" binding:"required" example:"John"`
	LastName  string `gorm:"type:varchar(32)" json:"lastname" binding:"required" example:"Doe"`
	Age       int    `json:"age" binding:"gte=18,lte=100" example:"30"`
	Email     string `gorm:"type:varchar(256)" json:"email" validate:"required,email" example:"john.doe@example.com"`
}

type Invention struct {
	ID          string    `gorm:"primary_key;auto_increment" json:"id" example:"123"`
	Title       string    `gorm:"type:varchar(255)" json:"title" binding:"min=3" validate:"is_ok" example:"Amazing Invention"`
	Date        time.Time `gorm:"type:date" json:"date" example:"2024-04-30T00:00:00Z"`
	Description string    `gorm:"type:varchar(300)" json:"description" binding:"min=3,max=300" example:"This is an amazing invention that will revolutionize the world."`
	Inventor    Inventor  `gorm:"foreignkey:InventorID" json:"inventor" binding:"required"`
	InventorID  uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
