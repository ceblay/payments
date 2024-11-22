package paymentgateway

type Repository interface {
	AddNewGateway(id, platformName, country string) (*Provider, error)
	GetGatewayByID(id string) (*Provider, error)
	FindGateways() []*Provider
}
