package paymentmethod

type PaymentMethodInput struct {
	Payment     float64
	Price       float64
	ClientPoint uint
}

type PaymentMethodOutput struct {
	Change         float64
	RemainingPoint uint
}

type Paymenter interface {
	Transact(input PaymentMethodInput) (output PaymentMethodOutput, err error)
}
