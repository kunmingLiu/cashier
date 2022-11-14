package vip

type VIP2 struct {
	*VIP
}

func NewVIP2(opts ...Option) *VIP2 {
	v := &VIP{
		discount: 0.9,
	}
	for _, opt := range opts {
		opt(v)
	}
	return &VIP2{
		VIP: v,
	}
}
