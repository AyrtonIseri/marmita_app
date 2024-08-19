package whatsapp

import (
	"fmt"
	"marmita/controller/client"
	"marmita/env"
	"marmita/types"
	"net/http"

	"github.com/twilio/twilio-go/twiml"
)

var USERSTATE map[string]types.FlowState

type Chatbot struct {
	ClientController client.Controller
}

func NewChatbot(cc client.Controller) *Chatbot {
	return &Chatbot{ClientController: cc}
}

func (cb Chatbot) WriteToUser(UserWhatsapp string, Body string, w http.ResponseWriter) {
	message := &twiml.MessagingMessage{
		Body: Body,
		To:   UserWhatsapp,
		From: env.ENV.TwilioNumber,
	}

	twimlResult, err := twiml.Messages([]twiml.Element{message})
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, twimlResult)
}
