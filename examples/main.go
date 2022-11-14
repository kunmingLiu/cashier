package main

import (
	"fmt"

	"github.com/kunmingliu/cashier/pkg/cashier"
	"github.com/kunmingliu/cashier/pkg/paymentmethod"
	"github.com/kunmingliu/cashier/pkg/redemption"
	"github.com/kunmingliu/cashier/pkg/vip"
)

func main() {
	var (
		change float64
		point  uint
		err    error
	)
	t := paymentmethod.NewNormalPayment()
	c := cashier.NewCashier(cashier.WithPaymentMethod(t))

	change, point, err = c.Transact(100, 50)
	fmt.Printf("change = %v, point = %v, err = %v\n", change, point, err)

	vip := vip.NewVIP2()
	vipPayment := paymentmethod.NewVIPPayment(vip)
	c1 := cashier.NewCashier(cashier.WithPaymentMethod(vipPayment))

	change, point, err = c1.Transact(100, 50)
	fmt.Printf("change = %v, point = %v, err = %v\n", change, point, err)

	redeem := redemption.NewRedeem()
	redeemPayment := paymentmethod.NewRedemptionPayment(redeem)
	c2 := cashier.NewCashier(cashier.WithPaymentMethod(redeemPayment), cashier.WithClientPoint(100))

	change, point, err = c2.Transact(100, 50)
	fmt.Printf("change = %v, point = %v, err = %v\n", change, point, err)

	vipRedeem := redemption.NewVIPRedeem(redeem)
	vipRedeemPayment := paymentmethod.NewRedemptionPayment(vipRedeem)
	c3 := cashier.NewCashier(cashier.WithPaymentMethod(vipRedeemPayment), cashier.WithClientPoint(100))

	change, point, err = c3.Transact(160, 50)
	fmt.Printf("change = %v, point = %v, err = %v\n", change, point, err)
}
