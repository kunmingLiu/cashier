package cashier

import (
	"fmt"

	"github.com/kunmingliu/cashier/pkg/paymentmethod"
)

type Cashier struct {
	paymenter   paymentmethod.Paymenter
	clientPoint uint
}

type Option func(*Cashier)

// WithPaymentMethod assign what payment method should be applied when transacting
func WithPaymentMethod(term paymentmethod.Paymenter) Option {
	return func(s *Cashier) {
		s.paymenter = term
	}
}

// WithClientPoint assign how many points a user owns
func WithClientPoint(point uint) Option {
	return func(s *Cashier) {
		s.clientPoint = point
	}
}

// NewCashier create a new cashier with specific options.
// Its method will be NormalPayment if the cashier doesn't assign any method.
func NewCashier(opts ...Option) *Cashier {
	c := &Cashier{
		paymenter: paymentmethod.NewNormalPayment(),
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Transact attempt to transact with the payment amount that the user pays, the points that the user owns and the method that the cashier chooses, and it will return change, remaining point and error.
// The error is nil when the transaction is success.
func (c *Cashier) Transact(price, payment float64) (change float64, remainPoint uint, err error) {
	if price <= 0 {
		err = fmt.Errorf("invalid Price %f", price)
		return
	}
	if payment < 0 {
		err = fmt.Errorf("invalid Payment %f", price)
		return
	}
	input := paymentmethod.PaymentMethodInput{
		Payment:     payment,
		Price:       price,
		ClientPoint: c.clientPoint,
	}
	output, err := c.paymenter.Transact(input)
	change = output.Change
	remainPoint = output.RemainingPoint
	if err != nil {
		return
	}
	return
}
