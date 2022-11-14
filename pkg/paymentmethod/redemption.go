package paymentmethod

import (
	"fmt"

	"github.com/kunmingliu/cashier/pkg/redemption"
)

type RedemptionPayment struct {
	redemption.Redeemer
}

func NewRedemptionPayment(redeemer redemption.Redeemer) *RedemptionPayment {
	return &RedemptionPayment{
		Redeemer: redeemer,
	}
}

func (p RedemptionPayment) Transact(input PaymentMethodInput) (output PaymentMethodOutput, err error) {
	price, point := p.Redeemer.GetPrice(input.Price, input.ClientPoint)
	change := input.Payment - price
	output.Change = change
	output.RemainingPoint = point
	if change < 0 {
		err = fmt.Errorf("payment is not enough, you are short by %f", (change * -1))
		return
	}
	return
}
