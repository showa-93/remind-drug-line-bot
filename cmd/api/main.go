package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	bot "github.com/showa-93/remind-drug-line-bot"
)

func main() {
	ctx := context.Background()

	wh := bot.WebhookHandler{}

	r := chi.NewRouter()
	r.Post("/webhook", wh.Post)

	bot.Info(ctx, "Starting Server...")
	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		bot.Fatal(ctx, err.Error())
		os.Exit(1)
	}
}
