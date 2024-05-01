package repository

import (
	"fmt"
	"github.com/xduck7/mini-go-server/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// "postgres://username:password@addr:port/dbname
const (
	username = "postgres"
	password = "postgres"
	addr     = "localhost"
	port     = "5432"
	dbname   = "postgres"
)

type InventionRepository interface {
	Save(invention entity.Invention)
	Update(invention entity.Invention)
	Delete(invention entity.Invention)
	FindAll() []entity.Invention
	FindById(idx int) entity.Invention
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewInventionRepository() InventionRepository {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, addr, port, dbname)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	err = db.AutoMigrate(&entity.Invention{}, &entity.Inventor{})
	if err != nil {
		return nil
	}
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		panic("Failed to close to database")
	}
	err = sqlDB.Close()
	if err != nil {
		return
	}
}

// TODO: fix repo methods
func (db *database) Save(invention entity.Invention) {
	db.connection.Create(&invention)
}

func (db *database) Update(invention entity.Invention) {
	db.connection.Save(&invention)
}
func (db *database) Delete(invention entity.Invention) {
	db.connection.Delete(&invention)
}
func (db *database) FindAll() []entity.Invention {
	var allInventions []entity.Invention
	db.connection.Set("gorm:auto_preload", true).Find(&allInventions)
	return allInventions
}

func (db *database) FindById(idx int) entity.Invention {
	var invention entity.Invention
	result := db.connection.Set("gorm:auto_preload", true).First(&invention, idx)
	if result.Error != nil {
		// Обработка ошибки
		return entity.Invention{}
	}
	return invention
}
