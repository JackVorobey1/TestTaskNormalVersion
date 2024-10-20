package model

import "gorm.io/gorm"

type ExampleModel struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `json:"email"`
}

// проверка на пустоту
func (u *ExampleModel) IsEmpty() bool {
	return u.ID == 0
}
