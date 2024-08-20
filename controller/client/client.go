package client

import (
	"fmt"
	"marmita/types"
)

type Controller struct {
	store types.ClientModel
}

func NewController(s types.ClientModel) *Controller {
	return &Controller{store: s}
}

func (c *Controller) RegisterClient(Whatsapp string, Address string, Name string) (err error) {
	Client, err := c.store.GetClientByWhatsapp(Whatsapp)

	if err == nil {
		err = fmt.Errorf("este whatsapp pertence a um cliente ja registrado: %s", Client.Name)
		return err
	}

	err = c.store.CreateClient(Whatsapp, Address, Name)
	if err != nil {
		return err
	}

	return
}

func (c *Controller) GetClients() ([]*types.Client, error) {
	ClientSlice, err := c.store.GetClients()

	if err != nil {
		return nil, err
	}

	if len(ClientSlice) == 0 {
		err = fmt.Errorf("nao ha nenhum cliente registrado")
		return nil, err
	}

	return ClientSlice, nil
}
