package service

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ceblay/payments/pkg/adapters"
	"github.com/ceblay/payments/pkg/app"
	"github.com/ceblay/payments/pkg/app/command"
	"github.com/ceblay/payments/pkg/app/query"
)

func NewApplication() app.Application {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	repository := adapters.NewSqliteRepository(db)
	paymentGatewayService := adapters.NewPayDunyaPaymentService()

	return app.Application{
		Queries: app.Queries{
			AllGateways: query.NewAllGatewayProvidersHandler(repository),
		},
		Commands: app.Commands{
			AddGateway:   command.NewAddGatewayHandler(repository),
			IssuePayment: command.NewIssuePaymentHandler(paymentGatewayService),
		},
	}
}
