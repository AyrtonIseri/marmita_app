package whatsapp

import (
	"errors"
	"fmt"
	"marmita/env"
	"marmita/internals/cookies"
	"marmita/types"
	"marmita/utils"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/twilio/twilio-go/twiml"
)

type Chatbot struct {
	ClientController types.ClientController
	RegisterFlow     types.Flow
	// QueryFlow        types.Flow
	// UpdateFlow       types.Flow
}

// func NewChatbot(cc types.ClientController, rf types.Flow, qf types.Flow, uf types.Flow) *Chatbot {
// 	return &Chatbot{ClientController: cc, RegisterFlow: rf, QueryFlow: qf, UpdateFlow: uf}
// }

func NewChatbot(cc types.ClientController, rf types.Flow) *Chatbot {
	return &Chatbot{ClientController: cc, RegisterFlow: rf}
}

func (cb Chatbot) MessageHandler(w http.ResponseWriter, r *http.Request) error {
	Request, err := ReadIncomingMessage(w, r)
	if err != nil {
		return err
	}

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	var Flow, Step int

	CookieFlow, err := cookies.ReadCookies(r, "flow")
	if err != nil {
		fmt.Println("Couldn't read the Flow Cookie ")
		fmt.Println(err)
		Flow = -1
	} else {
		Flow, err = strconv.Atoi(CookieFlow)
		if err != nil {
			fmt.Println("Couldn't decode the flow back to int ", CookieFlow)
			fmt.Println(err)
			Flow = -1
		}
	} // if there is anything not set, the Flow should be at the start

	CookieStep, err := cookies.ReadCookies(r, "step")
	if err != nil {
		Step = 0
	} else {
		Step, err = strconv.Atoi(CookieStep)
		if err != nil {
			Step = 0
		}
	} // if no step can be found, it should be at the start

	var Message string = "Por favor escolha sua opção:\n1. Fazer um Cadastro\n2.Fazer uma busca\n3.Corrigir dados"
	defer func() {
		fmt.Println("Writing back to user: ", Message)
		cb.WriteToUser(Request.WppUser, Message, w)
	}()
	defer func() {
		fmt.Println("Updating current Step: ", Step)
		cb.UpdateState("step", strconv.Itoa(Step), w)
	}()
	defer func() {
		fmt.Println("Updating current Flow: ", Flow)
		cb.UpdateState("flow", strconv.Itoa(Flow), w)
	}()

	fmt.Printf("Current Flow: %d, current Step: %d\n", Flow, Step)

	// Makes sure to write the welcome message at least once!
	if Flow == -1 {
		Flow = utils.NO_FLOW
		return nil
	}

	if Flow == utils.NO_FLOW {
		FlowHandler, err := cb.HandleHomeSection(&Flow, Request.Body)
		if err != nil {
			cb.LogError(Request.WppUser, err, w)
			http.Error(w, "Client error response", http.StatusBadRequest)
			return err
		}
		cb.NextStep(&Step)
		FlowHandler.WelcomeMessage(&Message)
		return nil
	}

	// get a flow handler (choose from everyone)

	return nil
}

func (cb Chatbot) HandleHomeSection(Flow *int, Response string) (types.Flow, error) {
	ValidResponses := map[string]bool{"1": true, "2": true, "3": true}

	if ok := ValidResponses[Response]; !ok {
		return nil, errors.New("A resposta nao eh valida. Tente novamente")
	}

	NewFlow, err := strconv.Atoi(Response)
	if err != nil {
		panic(err)
	}

	*Flow = NewFlow

	// if *Flow == utils.FLOW_REGISTER {
	// 	return cb.RegisterFlow, nil
	// }
	// if *Flow == utils.FLOW_QUERY {
	// 	return cb.QueryFlow, nil
	// }

	// return cb.UpdateFlow, nil
	return cb.RegisterFlow, nil
}

func (cb Chatbot) NextStep(Step *int) {
	*Step += 1
}

func (cb Chatbot) UpdateState(State string, Value string, w http.ResponseWriter) {
	UserCookie := http.Cookie{
		Name:   State,
		Value:  Value,
		MaxAge: 3600,
	}

	http.SetCookie(w, &UserCookie)
}

func (cb Chatbot) ValidateResponse(Flow int, Step int, Response string) error {

	switch Flow {
	case utils.FLOW_REGISTER:
		return cb.RegisterFlow.ValidateResponse(Response, Step)
		// case utils.FLOW_QUERY:
		// 	return cb.QueryFlow.ValidateResponse(Response, Step)
		// case utils.FLOW_UPDATE:
		// 	return cb.UpdateFlow.ValidateResponse(Response, Step)
	}

	return nil
}

func (cb Chatbot) LogError(UserWhatsapp string, err error, w http.ResponseWriter) {
	Message := err.Error()
	cb.WriteToUser(UserWhatsapp, Message, w)
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
