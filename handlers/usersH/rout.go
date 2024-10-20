package usersH

import (
	"TestTask/store/model"
	"fmt"
	"github.com/D0K-ich/types/iface"
	"github.com/D0K-ich/types/message"
	"github.com/kr/pretty"
)

// путь начинается с ОДНОГО user
func (h *Handler) Rout(incoming *message.Message) (payload any, err error) {
	switch incoming.SubjectAction() {
	case "info/get":
		//ReMarshalMust для структуры
		pretty.Print(iface.ReMarshalMust[*model.ExampleModel](incoming.Get("user")))
		payload = map[string]interface{}{
			"Its okay": true,
		}

	case "user/create":
		// Создание
		payload, err = h.CreateUserHandler(incoming.String("email"))

	case "user/update":
		// Обновление
		userID := incoming.String("userID")

		// Обновление пользователя
		payload, err = h.UpdateUserHandler(userID, incoming.String("newEmail"))

	case "user/delete":
		// Удаление
		userID := incoming.String("userID")

		payload, err = h.DeleteUserHandler(userID)

	default:
		err = fmt.Errorf("unknown action: %s", incoming.SubjectAction())
	}
	return
}
