package query

type AllGatewayProvidersReadModel interface {
	FindAllGateways() ([]Gateway, error)
}

type AllGatewayProvidersHandler struct {
	readModel AllGatewayProvidersReadModel
}

func NewAllGatewayProvidersHandler(rm AllGatewayProvidersReadModel) AllGatewayProvidersHandler {
	return AllGatewayProvidersHandler{
		readModel: rm,
	}
}

func (h *AllGatewayProvidersHandler) Handle() ([]Gateway, error) {
	return h.readModel.FindAllGateways()
}
