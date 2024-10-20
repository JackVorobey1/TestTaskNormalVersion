package pg

import (
	"TestTask/store/model"

	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//ТУТ МЕТОДЫ СТОРА

type PostgresAdapter struct {
	DB *gorm.DB
}

func NewAdapter(cfg *Config) (*PostgresAdapter, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ", cfg.Host, cfg.Username, cfg.Password, cfg.DBname, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return nil, err
	}

	log.Println("Успешное подключение к базе данных.")

	if err := db.AutoMigrate(&model.ExampleModel{}); err != nil {
		log.Fatalf("Ошибка при миграции базы данных: %v", err)
		return nil, err
	}

	// Возвращаем адаптер с подключением к базе данных
	return &PostgresAdapter{
		DB: db,
	}, nil
}
