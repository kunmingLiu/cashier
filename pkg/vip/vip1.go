package vip

type VIP1 struct {
	*VIP
}

func NewVIP1(opts ...Option) *VIP1 {
	v := &VIP{
		discount: 0.95,
	}
	for _, opt := range opts {
		opt(v)
	}
	return &VIP1{
		VIP: v,
	}
}
