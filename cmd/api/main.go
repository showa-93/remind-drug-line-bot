package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	bot "github.com/showa-93/remind-drug-line-bot"
)

func main() {
	config := bot.NewConfig()
	bot.InitLogger(config)

	wh := bot.NewWebhookHandler(config)

	r := chi.NewRouter()
	r.Post("/webhook", wh.Post)

	ctx := context.Background()
	bot.Info(ctx, "Starting Server... port="+config.Port)
	if err := http.ListenAndServe(":"+config.Port, r); err != nil {
		bot.Fatal(ctx, err.Error())
		os.Exit(1)
	}
}
