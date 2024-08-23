package main

import (
	"marmita/controller/client"
	clientModel "marmita/model/client"
	whatsapp "marmita/view/chatbot"
	"marmita/view/flows"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ClientModel := clientModel.NewModel()
	ClientController := client.NewController(ClientModel)
	RegisterFlow := flows.NewRegisterFlow()
	WhatsappView := whatsapp.NewChatbot(ClientController, RegisterFlow)

	WhatsappView.MessageHandler(w, r)
}

func main() {
	http.HandleFunc("/whatsapp", handler)
	http.ListenAndServe(":80", nil)
}
