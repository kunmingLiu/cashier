package redemption

//go:generate mockgen -destination=../../internal/mocks/redemption/mock.go -package=redemption github.com/kunmingliu/cashier/pkg/redemption Redeemer
type Redeemer interface {
	GetPrice(price float64, point uint) (remainingPrice float64, remainingPoint uint)
	GetRate() float64
}

type Option func(*Redeem)

type Redeem struct {
	rate float64 // how many amounts could be redeemed by 1 point
}

func WithRate(ratio float64) Option {
	return func(p *Redeem) {
		p.rate = ratio
	}
}

// NewRedeem will return a new Redeem with specific options.
// The rate will be 1 if the cashier doesn't set it.
func NewRedeem(opts ...Option) *Redeem {
	p := &Redeem{
		rate: 1,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p Redeem) GetPrice(price float64, point uint) (remainingPrice float64, remainingPoint uint) {
	redeemAmount := float64(point) * p.rate

	diff := redeemAmount - price
	if diff >= 0 {
		remainingPrice = 0
		remainingPoint = uint((diff / p.rate))
		return
	}
	remainingPrice = diff * -1
	return
}

func (p Redeem) GetRate() float64 {
	return p.rate
}
