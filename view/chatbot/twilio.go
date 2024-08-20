package whatsapp

import (
	"fmt"
	"marmita/types"
	"net/http"

	"github.com/gorilla/schema"
)

// aqui vamos formatar as mensagens para retornar para o twilio

func ReadIncomingMessage(w http.ResponseWriter, r *http.Request) (types.TwilioRequest, error) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Couldn't parse url properly. Abording request")
		panic(err)
	}

	var Request types.TwilioRequest

	if MessageType := r.FormValue("MessageType"); MessageType != "text" {
		err = fmt.Errorf("formato invalido de mensagem. reiniciando chat")
		return Request, err
	}

	Decoder := schema.NewDecoder()
	Decoder.IgnoreUnknownKeys(true)
	err = Decoder.Decode(&Request, r.Form)
	if err != nil {
		fmt.Println("Couldn't decode the original string into the Request type")
		panic(err)
	}

	fmt.Println("\nUser Request: ", Request)

	return Request, nil
}
