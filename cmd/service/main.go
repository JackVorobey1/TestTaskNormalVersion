package main

import (
	"TestTask"
	"TestTask/cmd/handlers"
	"TestTask/store"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"syscall"
)

var globalConfig *TestTask.Config

// Загрузка конфигурации, результат сохранятется в globalConfig
func init() {
	var err error
	globalConfig, err = TestTask.LoadConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func main() {

	var err error
	store.Default, err = store.InitStore(globalConfig.Store)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = store.Default.Adapter.CreateExample("John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	log.Logger = log.With().Caller().Logger().Level(zerolog.DebugLevel).Output(zerolog.ConsoleWriter{Out: os.Stderr})

	/*router, err := NewRouter(globalConfig.Network)
	if err != nil {
		log.Fatal(err)
	}*/
	log.Print("Сервер запущен на http://localhost:8080")

	// Запускаем HTTP-сервер и указываем функцию обработчика из пакета handler
	if err := fasthttp.ListenAndServe(":8080", handlers.RequestHandler); err != nil {
		fmt.Printf("Ошибка прослушивания сервера")
	}
	// Основная программа не завершится, пока не получит сигнал завершения
	sig := make(chan os.Signal, 1)
	log.Info().Msg("Ready")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	log.Info().Msg("Программа завершена")

}
