package app

import (
	"github.com/ceblay/payments/pkg/app/command"
	"github.com/ceblay/payments/pkg/app/query"
)

type Queries struct {
	AllGateways query.AllGatewayProvidersHandler
}

type Commands struct {
	AddGateway   command.AddGatewayHandler
	IssuePayment command.IssuePaymentHandler
}

type Application struct {
	Queries  Queries
	Commands Commands
}
