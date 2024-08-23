package types

import (
	"net/http"
	"time"
)

type TwilioRequest struct {
	WppUser   string `schema:"From"`
	TwilioWpp string `schema:"To"`
	Body      string `schema:"Body"`
}

type Client struct {
	ID       uint
	Whatsapp string
	Address  string
	Name     string
}

type Delivery struct {
	ID     uint
	Date   time.Time
	Region string
}

type Order struct {
	ID         uint
	UserID     uint
	DeliveryId uint
}

type Marmita struct {
	ID       uint
	OrderID  uint
	Protein  string
	Carb     string
	Size     string
	Extras   string // this can be something like 'feijao' or 'creme de milho'
	Comments string
}

type ClientModel interface {
	CreateClient(Whatsapp string, Address string, Name string) error
	GetClientById(UserID uint) (Client, error)
	GetClientByWhatsapp(Whatsapp string) (Client, error)
	GetClients() ([]*Client, error)
}

type ClientController interface {
}

type View interface {
	MessageHandler(r *http.Request, w http.ResponseWriter) error // this should route the request to the correct flow! And also reply
	WriteToUser(message string) error                            // write a message to the user
	SaveInput(value string) error                                // save the input somewhere so that the user can input its information 1-by-1
	ForwardCommand() error                                       // calls controller to further commands
}

type Flow interface {
	ListEntities() error                              // list all entities available for registering/update/etc (should be granular at flow level)
	GetEntityFields(entity string) error              // returns all fields of an entity in a slice lets say
	ValidateResponse(response string, step int) error // validate whether the response is acceptable (multi choice for example)
	WelcomeMessage(*string)
}
