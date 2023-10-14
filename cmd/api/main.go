package main

import (
	"context"
	"io"
	"net/http"
	"os"

	bot "github.com/showa-93/remind-drug-line-bot"
)

func main() {
	ctx := context.Background()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bot.Info(r.Context(), "start handle")
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bot.Warn(r.Context(), string(b))
		w.WriteHeader(http.StatusOK)
	})

	bot.Info(ctx, "Starting Server...")
	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		bot.Fatal(ctx, err.Error())
		os.Exit(1)
	}
}
