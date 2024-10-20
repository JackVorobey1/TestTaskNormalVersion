package logs

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// Твоя функция настройки логирования
func SetupLog(logLevel string) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	switch logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
		fmt.Print("Panic status")
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
		fmt.Print("Fatal stastus")

	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func AAA() {
	SetupLog("info")
	log.Info().Msg("status info")
}
