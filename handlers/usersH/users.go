package usersH

import (
	"TestTask/store/model"
	"TestTask/store/usersS"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

// Handler — структура для работы с пользователями
type Handler struct {
	store  *usersS.Storage
	userID string
}

// NewHandler — создание нового хендлера с передачей store и userID
func NewHandler(storage *usersS.Storage, userID string) *Handler {
	return &Handler{
		store:  storage,
		userID: userID,
	}
}

// GetUserId — возвращает текущий userID
func (h *Handler) GetUserId() string {
	return h.userID
}

// СОЗДАНИЕ
func (h *Handler) CreateUserHandler(email string) (payload any, err error) {
	if email = strings.TrimSpace(email); email == "" {
		return nil, errors.New("email is required1")
	}

	var newUser *model.ExampleModel
	if newUser, err = h.store.CreateUser(email); err != nil {
		return nil, err
	}

	payload = map[string]interface{}{
		"user": newUser,
	}
	return
}

// ПОЛУЧЕНИЕ ПОЛЬЗОВАТЕЛЯ ПО ID (ЧТЕНИЕ)
func (h *Handler) GetUserByIdHandler(userID string) (payload any, err error) {
	// Получаем пользователя из хранилища по userID
	var user *model.ExampleModel
	if user, err = h.store.GetUserByID(userID); err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Формируем ответ
	payload = map[string]interface{}{
		"user": user,
	}
	return payload, nil
}

// ОБНОВЛЕНИЕ
func (h *Handler) UpdateUserHandler(userID string, newEmail string) (payload any, err error) {
	if userID == "0" {
		return nil, errors.New("invalid user ID")
	}
	log.Info().Msg("Прошел проверку на 0")
	// УБИРАЕМ ПРОБЕЛЫ и проверяем новый email
	if newEmail = strings.TrimSpace(newEmail); newEmail == "" {
		return nil, errors.New("new email is required3")
	}

	log.Info().Msg("После удаления пробелов в NewEmail идем дальше")
	// вызов метода
	var updatedUser *model.ExampleModel
	if updatedUser, err = h.store.UpdateUser(userID, newEmail); err != nil {
		return nil, err
	}

	// наш ответ
	payload = map[string]interface{}{
		"user": updatedUser,
	}
	return payload, nil
}

// удаление пользователя
func (h *Handler) DeleteUserHandler(userID string) (payload any, err error) {
	// Получаем пользователя по ID
	var user *model.ExampleModel
	if user, err = h.store.GetUserByID(userID); err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Удаляем пользователя
	if err = h.store.DeleteUser(userID); err != nil {
		return nil, err
	}

	// Возвращаем ответ, включая информацию о удалённом пользователе
	payload = map[string]interface{}{
		"message": "user deleted successfully",
		"user":    user, // Информация о пользователе
	}
	return payload, nil
}
