package usersS

import (
	"TestTask/store/model"
	"errors"
	"gorm.io/gorm"
)

//достают инфу с БД

// ищем и получаем пользователя по ID
func (s *Storage) GetUserById(userID int) (user *model.ExampleModel, err error) {
	if err = s.Db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		return
	}
	return
}
