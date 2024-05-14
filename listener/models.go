package listener

type WhatsAppMessage struct {
	Entry []struct {
		Changes []struct {
			Value struct {
				Messages []struct {
					Type string `json:"type"`
					From string `json:"from"`
					Text struct {
						Body string `json:"body"`
					} `json:"text"`
					ID string `json:"id"`
				} `json:"messages"`
				Metadata struct {
					PhoneNumberID string `json:"phone_number_id"`
				} `json:"metadata"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}

// TextContent represents the text body of a WhatsApp message.
type TextContent struct {
	Body string `json:"body"`
}

// MessageContext represents the context in which a message is sent, particularly for replies.
type MessageContext struct {
	MessageID string `json:"message_id"`
}

// MessageRequest represents the structure for sending a message via WhatsApp API.
type MessageRequest struct {
	MessagingProduct string         `json:"messaging_product"`
	To               string         `json:"to"`
	Text             TextContent    `json:"text"`
	Context          MessageContext `json:"context"`
}

// ReadStatusRequest represents the structure for updating a message's read status.
type ReadStatusRequest struct {
	MessagingProduct string `json:"messaging_product"`
	Status           string `json:"status"`
	MessageID        string `json:"message_id"`
}
