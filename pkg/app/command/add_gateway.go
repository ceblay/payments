package command

import (
	"errors"
	"log"

	pg "github.com/ceblay/payments/pkg/domain/paymentgateway"
)

type AddGateway struct {
	Name string
}

type AddGatewayWriteModel interface {
	AddNewGateway(p pg.Provider) (*pg.Provider, error)
}

type AddGatewayHandler interface {
	Handle(cmd AddGateway) error
}

type addGatewayHandler struct {
	writeModel AddGatewayWriteModel
}

func NewAddGatewayHandler(wm AddGatewayWriteModel) AddGatewayHandler {
	if wm == nil {
		panic("nil write model")
	}
	return addGatewayHandler{
		writeModel: wm,
	}
}

func (h addGatewayHandler) Handle(cmd AddGateway) error {
	log.Println("COMMAND: ", cmd)

	provider, err := pg.NewProvider(cmd.Name)
	if err != nil {
		return errors.New("Failed domain layer invariants")
	}
	gateway, err := h.writeModel.AddNewGateway(*provider)
	log.Println("GATEWAY: ", gateway)

	if err != nil {
		return errors.New("Could not be created at DB layer")
	}

	return nil
}
