package listener

import (
	"log"
	"net/http"

	"github.com/matheuspolitano/GadgetHub/utils"
)

type WebhookHandler struct {
	config utils.Config
}

func NewWebhookHandler(config utils.Config) *WebhookHandler {
	return &WebhookHandler{
		config: config,
	}
}

func (h *WebhookHandler) VerifyWebhook(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	if mode == "subscribe" && token == h.config.WebhookVerifyToken {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
		log.Println("Webhook verified successfully!")
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}
