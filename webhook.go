package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WebhookPayload struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

type Event struct {
	Type            string          `json:"type"`
	Message         *Message        `json:"message,omitempty"`
	Timestamp       int64           `json:"timestamp"`
	Source          Source          `json:"source"`
	ReplyToken      string          `json:"replyToken"`
	Mode            string          `json:"mode"`
	WebhookEventId  string          `json:"webhookEventId"`
	DeliveryContext DeliveryContext `json:"deliveryContext"`
}

type Message struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

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

	if err := h.validateSignature(r.Header, body); err != nil {
		Error(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		Error(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// user id 確認
	if len(payload.Events) > 0 {
		Debug(ctx, payload.Events[0].Source.UserID)
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
