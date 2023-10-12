package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	h := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(h)
	slog.SetDefault(logger)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info(string(b))
		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("%w", err)
		os.Exit(1)
	}
}
