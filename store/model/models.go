package model

import "gorm.io/gorm"

type ExampleModel struct {
	gorm.Model

	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}
