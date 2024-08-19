package types

import (
	"net/http"
	"time"
)

type FlowState int
type RegisterState int

type TwilioResponse struct {
	WppUser   string `json:"WaId"`
	TwilioWpp string `json:"To"`
	Body      string `json:"Body"`
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
	MessageHandler(r *http.Request, w http.ResponseWriter) error
}
