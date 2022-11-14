package redemption

import (
	"testing"
)

func TestRedeem(t *testing.T) {
	type args struct {
		opts        []Option
		price       float64
		clientPoint uint
	}
	type wants struct {
		price       float64
		clientPoint uint
		rate        float64
	}

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "default rate",
			args: args{
				price:       100,
				clientPoint: 10,
			},
			wants: wants{
				price:       90,
				clientPoint: 0,
				rate:        1,
			},
		},
		{
			name: "assign new rate and remainingPrice is not zero",
			args: args{
				price:       100,
				clientPoint: 10,
				opts: []Option{
					WithRate(2),
				},
			},
			wants: wants{
				price:       80,
				clientPoint: 0,
				rate:        2,
			},
		},
		{
			name: "assign new rate and remainingPrice is zero",
			args: args{
				price:       100,
				clientPoint: 10,
				opts: []Option{
					WithRate(15),
				},
			},
			wants: wants{
				price:       0,
				clientPoint: 3,
				rate:        15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRedeem(tt.args.opts...)
			if d := got.GetRate(); d != tt.wants.rate {
				t.Errorf("name = %s, rate = %v, want %v", tt.name, d, tt.wants.rate)
			}
			if price, clientPoint := got.GetPrice(tt.args.price, tt.args.clientPoint); price != tt.wants.price || clientPoint != tt.wants.clientPoint {
				if price != tt.wants.price {
					t.Errorf("name = %s, price = %v, want %v", tt.name, price, tt.wants.price)
				}
				if clientPoint != tt.wants.clientPoint {
					t.Errorf("name = %s, clientPoint = %v, want %v", tt.name, clientPoint, tt.wants.clientPoint)
				}
			}
		})
	}
}
