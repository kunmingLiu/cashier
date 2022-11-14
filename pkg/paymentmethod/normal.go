package paymentmethod

import "fmt"

type NormalPayment struct {
}

func NewNormalPayment() *NormalPayment {
	return &NormalPayment{}
}

func (n NormalPayment) Transact(input PaymentMethodInput) (output PaymentMethodOutput, err error) {
	change := input.Payment - input.Price
	output.Change = change
	output.RemainingPoint = input.ClientPoint
	if change < 0 {
		err = fmt.Errorf("payment is not enough, you are short by %f", (change * -1))
		return
	}
	return
}
