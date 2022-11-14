package vip

//go:generate mockgen -destination=../../internal/mocks/vip/mock.go -package=redemption github.com/kunmingliu/cashier/pkg/vip VIPer
type VIPer interface {
	GetPrice(price float64) float64
	GetDiscount() float64
}

type VIP struct {
	discount float64
}

type Option func(*VIP)

// WithDiscount will replace the default discount of a VIP.
func WithDiscount(discount float64) Option {
	return func(v *VIP) {
		v.discount = discount
	}
}

func (v VIP) GetDiscount() float64 {
	return v.discount
}

func (v VIP) GetPrice(price float64) float64 {
	return price * v.discount
}
