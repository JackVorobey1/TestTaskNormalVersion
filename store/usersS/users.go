package usersS

import "gorm.io/gorm"

type Storage struct {
	Db *gorm.DB
}

// создания нового стора с базой данных
func NewStore(db *gorm.DB) *Storage {
	return &Storage{
		Db: db,
	}
}
