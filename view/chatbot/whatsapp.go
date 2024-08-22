package whatsapp

import (
	"fmt"
	"marmita/env"
	"marmita/internals/cookies"
	"marmita/types"
	"net/http"
	"strconv"

	"github.com/twilio/twilio-go/twiml"
)

var USERSTATE map[string]types.FlowState

type Chatbot struct {
	ClientController types.ClientController
}

func NewChatbot(cc types.ClientController) *Chatbot {
	return &Chatbot{ClientController: cc}
}

func (cb Chatbot) MessageHandler(w http.ResponseWriter, r *http.Request) error {
	Request, err := ReadIncomingMessage(w, r)
	if err != nil {
		return err
	}

	cookieVal, err := cookies.ReadCookies(r, "stage")
	if err != nil {
		cookieVal = "1"
	}

	stage, err := strconv.Atoi(cookieVal)
	if err != nil {
		fmt.Println("Cookie sent is not a valid integer:", cookieVal)
		stage = 1
	}

	FinalBody := "Message received. Submitting a copy: " + Request.Body + "\nYour current stage is: " + strconv.Itoa(stage)

	stage += 1
	if stage > 3 {
		stage = 1
	}

	ResponseCookie := http.Cookie{
		Name:   "stage",
		Value:  strconv.Itoa(stage),
		MaxAge: 3600,
	}

	cookies.WriteCookies(ResponseCookie, w)
	cb.WriteToUser(Request.WppUser, FinalBody, w)

	return nil
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
