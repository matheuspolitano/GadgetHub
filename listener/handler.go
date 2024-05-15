package listener

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

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
		log.Info().Time("receive", time.Now()).Msg("Webhook verified successfully!")
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "unexpected body", http.StatusBadRequest)
		return
	}
	log.Info().Time("receive", time.Now()).Fields(requestBody).Msg("Incoming webhook message:")

	entry, ok := requestBody["entry"].([]interface{})
	if !ok || len(entry) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	changes, ok := entry[0].(map[string]interface{})["changes"].([]interface{})
	if !ok || len(changes) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	value, ok := changes[0].(map[string]interface{})["value"].(map[string]interface{})
	if !ok {
		w.WriteHeader(http.StatusOK)
		return
	}

	messages, ok := value["messages"].([]interface{})
	if !ok || len(messages) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	message, ok := messages[0].(map[string]interface{})
	if !ok || message["type"] != "text" {
		w.WriteHeader(http.StatusOK)
		return
	}

	metadata, ok := value["metadata"].(map[string]interface{})
	if !ok {
		w.WriteHeader(http.StatusOK)
		return
	}

	businessPhoneNumberID, ok := metadata["phone_number_id"].(string)
	if !ok {
		w.WriteHeader(http.StatusOK)
		return
	}

	messageID, ok := message["id"].(string)
	if !ok {
		w.WriteHeader(http.StatusOK)
		return
	}

	from, ok := message["from"].(string)
	if !ok {
		w.WriteHeader(http.StatusOK)
		return
	}

	h.sendReplyMessage(businessPhoneNumberID, from, messageID)
	h.markMessageAsRead(businessPhoneNumberID, messageID)

	w.WriteHeader(http.StatusOK)
}

func (h *WebhookHandler) sendReplyMessage(phoneNumberID, to, messageID string) {
	url := "https://graph.facebook.com/v18.0/" + phoneNumberID + "/messages"
	data := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                to,
		"text":              map[string]string{"body": "Como vc esta?"},
		"context":           map[string]string{"message_id": messageID},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Err(err).Time("time", time.Now())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Err(err).Time("time", time.Now()).Msg("err in response")
		return
	}

	req.Header.Set("Authorization", "Bearer "+h.config.MetaApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Time("time", time.Now()).Msg("Error sending HTTP request:")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Err(err).Time("time", time.Now()).Msg("Error response from API:")
	}
}

func (h *WebhookHandler) markMessageAsRead(phoneNumberID, messageID string) {
	url := "https://graph.facebook.com/v18.0/" + phoneNumberID + "/messages"
	data := map[string]interface{}{
		"messaging_product": "whatsapp",
		"status":            "read",
		"message_id":        messageID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Err(err).Time("time", time.Now()).Msg("Error marshaling JSON:")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Err(err).Time("time", time.Now()).Msg("Error creating HTTP request:")
		return
	}

	req.Header.Set("Authorization", "Bearer "+h.config.MetaApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Time("time", time.Now()).Msg("Error sending HTTP request:")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Err(err).Time("time", time.Now()).Msg("Error response from API:")
	}
}
