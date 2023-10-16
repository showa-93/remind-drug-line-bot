package bot

import (
	"io"
	"net/http"
)

type WebhookHandler struct {
}

func (h *WebhookHandler) Post(w http.ResponseWriter, r *http.Request) {
	Info(r.Context(), "handle webhook")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Warn(r.Context(), string(b))
	w.WriteHeader(http.StatusOK)
}
