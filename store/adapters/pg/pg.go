package pg

import (
	"TestTask/store/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

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

// Функция для создания новой записи в базе данных
func (adapter *PostgresAdapter) CreateExample(name string, email string) error {
	// Создаем объект модели ExampleModel
	example := model.ExampleModel{
		Name:  name,
		Email: email,
	}

	// Добавляем запись в базу данных через адаптер (DB)
	if err := adapter.DB.Create(&example).Error; err != nil {
		return err
	}

	log.Println("Запись успешно создана:", example)
	return nil
}
