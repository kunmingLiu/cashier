package redemption

type VIPRedeemer interface {
	GetDiscount() float64
	GetPointFloor() uint
}

type VIPRedeem struct {
	clientPointFloor uint
	discount         float64
	redeemer         Redeemer
}

type VIPOption func(*VIPRedeem)

// WithPointFloor will set how many points user redeems and it can get an extra discount.
func WithPointFloor(floor uint) VIPOption {
	return func(p *VIPRedeem) {
		p.clientPointFloor = floor
	}
}

// WithDiscount will set how many discounts user gets when he has enough points.
func WithDiscount(discount float64) VIPOption {
	return func(p *VIPRedeem) {
		p.discount = discount
	}
}

// NewVIPRedeem will return a new VIPRedeem with specific options.
// The discount will be 0.9 if the cashier doesn't set it.
// The clientPointFloor will be 100 if the cashier doesn't set it.
func NewVIPRedeem(redeemer Redeemer, opts ...VIPOption) *VIPRedeem {
	v := &VIPRedeem{
		clientPointFloor: 100,
		discount:         0.9,
		redeemer:         redeemer,
	}

	for _, opt := range opts {
		opt(v)
	}
	return v
}

func (v VIPRedeem) GetPrice(price float64, clientPoint uint) (remainingPrice float64, remainingPoint uint) {
	remainingPrice, remainingPoint = v.redeemer.GetPrice(price, clientPoint)
	if remainingPrice == 0.0 || clientPoint < v.clientPointFloor {
		return
	}
	remainingPrice *= v.discount
	return
}

func (v VIPRedeem) GetDiscount() float64 {
	return v.discount
}

func (v VIPRedeem) GetPointFloor() uint {
	return v.clientPointFloor
}

func (v VIPRedeem) GetRate() float64 {
	return v.redeemer.GetRate()
}
