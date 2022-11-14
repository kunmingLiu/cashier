package vip

type VIP3 struct {
	*VIP
}

func NewVIP3(opts ...Option) *VIP2 {
	v := &VIP{
		discount: 0.85,
	}
	for _, opt := range opts {
		opt(v)
	}
	return &VIP2{
		VIP: v,
	}
}
