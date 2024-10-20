package handlers

import (
	"TestTask/store/model"
	"github.com/D0K-ich/types/iface"
	"github.com/D0K-ich/types/message"
	"github.com/valyala/fasthttp"
)

func RequestHandler(ctx *fasthttp.RequestCtx) {
	var incoming *message.Message
	var err error
	if incoming, err = message.FromJson(ctx.PostBody(), string(ctx.Path())); err != nil {
		return
	}
	var exmo *model.ExampleModel
	exmo = iface.ReMarshalMust[*model.ExampleModel](incoming.Get("data"))

}
