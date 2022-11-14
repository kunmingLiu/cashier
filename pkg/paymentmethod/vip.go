package paymentmethod

import (
	"fmt"

	"github.com/kunmingliu/cashier/pkg/vip"
)

type VIPPayment struct {
	vip.VIPer
}

func NewVIPPayment(vip vip.VIPer) *VIPPayment {
	return &VIPPayment{
		VIPer: vip,
	}
}

func (v VIPPayment) Transact(input PaymentMethodInput) (output PaymentMethodOutput, err error) {
	change := input.Payment - v.VIPer.GetPrice(input.Price)
	output.Change = change
	output.RemainingPoint = input.ClientPoint
	if change < 0 {
		err = fmt.Errorf("payment is not enough, you are short by %f", (change * -1))
		return
	}
	return
}
