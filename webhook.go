package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WebhookHandler struct {
	c Config
}

func NewWebhookHandler(c Config) *WebhookHandler {
	return &WebhookHandler{
		c: c,
	}
}

func (h *WebhookHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Warn(ctx, string(body))

	if err := h.validateSignature(r.Header, body); err != nil {
		Error(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *WebhookHandler) validateSignature(header http.Header, body []byte) error {
	originSig := header.Get("x-line-signature")
	if originSig == "" {
		return errors.New("署名が取得できない")
	}

	mac := hmac.New(sha256.New, []byte(h.c.LineSecret))
	if _, err := mac.Write(body); err != nil {
		return fmt.Errorf("ダイジェスト値の生成に失敗 %w", err)
	}

	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	if sig != originSig {
		return fmt.Errorf("署名が一致しない 生成した署名=%s, ヘッダーの署名=%s", string(sig), originSig)
	}

	return nil
}
