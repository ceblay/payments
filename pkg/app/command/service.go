package command

type PaymentService interface {
	InitiatePayment(float32) (string, error)
}
