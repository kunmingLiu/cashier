package vip

import (
	"testing"
)

func TestVIP3(t *testing.T) {
	type args struct {
		opts  []Option
		price float64
	}

	type wants struct {
		price    float64
		discount float64
	}

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "default discount",
			args: args{
				price: 100,
			},
			wants: wants{
				discount: 0.85,
				price:    100 * 0.85,
			},
		},
		{
			name: "assign new discount",
			args: args{
				opts: []Option{
					WithDiscount(0.7),
				},
				price: 100,
			},
			wants: wants{
				discount: 0.7,
				price:    100 * 0.7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVIP3(tt.args.opts...)
			if got.discount != tt.wants.discount {
				t.Errorf("name = %s, discount = %v, want %v", tt.name, got.discount, tt.wants.discount)
			}
			if got.GetPrice(tt.args.price) != tt.wants.price {
				t.Errorf("name = %s, price = %v, want %v", tt.name, got.GetPrice(tt.args.price), tt.wants.price)
			}
		})
	}
}
