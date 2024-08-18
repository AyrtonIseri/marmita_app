package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/twilio/twilio-go/twiml"
)

func handler(w http.ResponseWriter, r *http.Request) {

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New request received. Information: \n %s \n", string(dump))

	message := &twiml.MessagingMessage{
		Body: "message received.",
		To:   "mock_whatsapp",
		From: "twilio_whatsapp",
	}
	twimlResult, err := twiml.Messages([]twiml.Element{message})
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, twimlResult)
}

func main() {
	http.HandleFunc("/whatsapp", handler)
	http.ListenAndServe(":80", nil)
}
