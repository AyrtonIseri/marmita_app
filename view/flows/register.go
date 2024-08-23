package flows

type RegisterFlow struct{}

func NewRegisterFlow() *RegisterFlow {
	return &RegisterFlow{}
}

func (r *RegisterFlow) ValidateResponse(response string, step int) error {
	return nil
}

func (r *RegisterFlow) ListEntities() error {
	return nil
}

func (r *RegisterFlow) GetEntityFields(entity string) error {
	return nil
}

func (r *RegisterFlow) WelcomeMessage(Message *string) {
	*Message = "Bem vindo a primeira etapa do registro!"
}
