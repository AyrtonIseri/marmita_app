package clientModel

import (
	"errors"
	"marmita/types"
)

type model struct {
}

func NewModel() *model {
	return &model{}
}

func (m *model) CreateClient(Whatsapp string, Address string, Name string) error {
	return errors.New("not implemented error")
}

func (m *model) GetClientById(UserID uint) (types.Client, error) {
	return types.Client{}, errors.New("not implemented error")
}

func (m *model) GetClientByWhatsapp(Whatsapp string) (types.Client, error) {
	return types.Client{}, errors.New("not implemented error")
}

func (m *model) GetClients() ([]*types.Client, error) {
	var EmptyClients []*types.Client
	return EmptyClients, errors.New("not implemented error")
}
