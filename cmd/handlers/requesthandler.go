package handlers

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func RequestHandler(ctx *fasthttp.RequestCtx) { // Изменено: передаем указатель
	switch string(ctx.Path()) {
	case "/":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("text/plain; charset=utf-8")
		ctx.SetBodyString("You connect to server")
		log.Info().Msg("User connect to server")
	case "/post":
		if ctx.IsPost() {
			body := ctx.PostBody()
			log.Info().Msg(" Post Request\n")
			fmt.Print(ctx, body)
		} else if ctx.IsGet() {

			log.Info().Msg(" Get Request\n")
			fmt.Print(ctx)
		} else {
			ctx.Error("\n ERROR\n", fasthttp.StatusMethodNotAllowed)
		}

	case "/main":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString("Вы на главной сттанице")
		log.Info().Msg(" User in main\n")
		fmt.Print(ctx)

	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString("Страница не найдена")
		log.Info().Msg(" User didn't find the page\n")

	}
}

/*func PostHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "post":
		if ctx.IsPost() {
			body := ctx.PostBody()
			fmt.Printf(ctx, "Post request", body)
		} else {
			ctx.Error("Post request ERROR", fasthttp.StatusMethodNotAllowed)
		}
	}
}
*/
