package command

import (
	"errors"
	"log"
)

type IssuePayment struct {
	Amount float32
}

type IssuePaymentHandler interface {
	Handle(cmd IssuePayment) error
}

type issuePaymentHandler struct {
	paymentService PaymentService
}

func NewIssuePaymentHandler(ps PaymentService) IssuePaymentHandler {
	if ps == nil {
		panic("no payment service set")
	}
	return issuePaymentHandler{
		paymentService: ps,
	}
}

func (h issuePaymentHandler) Handle(cmd IssuePayment) error {
	log.Println("COMMAND: ", cmd)

	response, err := h.paymentService.InitiatePayment(cmd.Amount)
	log.Println("PAYMENT PROCESSED: ", response)

	if err != nil {
		log.Println("ACTUAL ERROR", err)
		return errors.New("Could not process payment successfully")
	}

	return nil
}
