package usersS

import (
	"TestTask/store/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
)

//сеттит в базу!!!

// проверки формата email
func ValidateEmail(email string) bool {
	//проверка
	var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)

}

// СОЗДАНИЕ ЮЗЕРА
func (s *Storage) CreateUser(email string) (*model.ExampleModel, error) {
	if email == "" {
		return nil, errors.New("email is required2")
	}

	if !ValidateEmail(email) {
		return nil, errors.New("invalid email format")
	}

	//Проверка на существующего пользователя
	var existingUser model.ExampleModel
	if err := s.Db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user := &model.ExampleModel{
		Email: email,
	}
	//сощдаем пользователя по почте и сохраняем
	if err := s.Db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ЧТЕНИЕ ПОЛЬЗОВАТЕЛЯ ПО ID
func (s *Storage) GetUserByID(userID string) (*model.ExampleModel, error) {
	var user model.ExampleModel

	// Поиск пользователя по ID
	if err := s.Db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user withh ID %d not found", userID)
		}
		return nil, err
	}

	return &user, nil
}

// ОБНОВЛЕНИЕ ПОЛЬЗОВАТЕЛЯ
func (s *Storage) UpdateUser(userID string, newEmail string) (*model.ExampleModel, error) {
	// Находим пользователя по ID
	var user model.ExampleModel
	if err := s.Db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d not found", userID)
		}
		return nil, err
	}

	// Проверка корректности нового email
	if !ValidateEmail(newEmail) {
		return nil, errors.New("invalid email format")
	}

	// Обновляем email пользователя
	user.Email = newEmail
	if err := s.Db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// УДАЛЕНИЕ ПОЛЬЗОВАТЕЛЯ
func (s *Storage) DeleteUser(userID string) error {
	// Проверяем существование пользователя
	var user model.ExampleModel
	if err := s.Db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %d not found", userID)
		}
		return err
	}

	// Удаляем пользователя
	if err := s.Db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
