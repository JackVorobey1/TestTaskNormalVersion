package server

import (
	"TestTask/handlers/usersH"
	"TestTask/store/usersS"
	"encoding/json"
	"github.com/D0K-ich/types/message"
	"github.com/fasthttp/router"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"runtime/debug"
)

//ДЕФЕР

// создает маршрутизатор, который принимает POST-запросы на пути /user/{sub}/{act},
// разбирает тело запроса как JSON, передает данные в обработчик, который их обрабатывает, и возвращает JSON-ответ клиенту.
func NewRouter(storage *usersS.Storage) (mainrouter *router.Router) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("(network) >> Panic recovered in write %s %v", "err", r)
			debug.PrintStack()
			return
		}
	}()
	mainrouter = router.New()
	//ошибки при панике маршрутизатора
	//mainrouter.PanicHandler = func(ctx *fasthttp.RequestCtx, i interface{}) {
	//	ctx.Error("Err", fasthttp.StatusInternalServerError)
	//}

	mainrouter.POST("/user/{sub}/{act}", func(ctx *fasthttp.RequestCtx) {
		var err error
		//входящий
		var incoming *message.Message
		if incoming, err = message.FromJson(ctx.PostBody(), string(ctx.Path())); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
		//создает новый экземпляр обработчика Handler, вызывая конструктор NewHandler из пакета users
		//и передавая строку "aaa" в качестве аргумента.
		var handler = usersH.NewHandler(storage, "31")
		//ответ
		var payload any
		if payload, err = handler.Rout(incoming); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}
		var data []byte
		if data, err = json.Marshal(payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
		ctx.SetBody(data)
		ctx.SetStatusCode(fasthttp.StatusOK)
	})

	return
}
