package redemption

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockRedemption "github.com/kunmingliu/cashier/internal/mocks/redemption"
)

func TestVIPRedeem(t *testing.T) {
	type args struct {
		opts        []VIPOption
		price       float64
		clientPoint uint
	}
	type wants struct {
		price            float64
		discount         float64
		clientPoint      uint
		clientPointFloor uint
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRedemption := mockRedemption.NewMockRedeemer(ctl)

	gomock.InOrder(
		mockRedemption.EXPECT().GetPrice(100.0, uint(10)).Return(0.0, uint(10)),
		mockRedemption.EXPECT().GetPrice(100.0, uint(10)).Return(20.0, uint(10)),
		mockRedemption.EXPECT().GetPrice(100.0, uint(10)).Return(20.0, uint(10)),
		mockRedemption.EXPECT().GetPrice(100.0, uint(21)).Return(20.0, uint(1)),
	)

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "default discount and clientPointFloor and remainingPrice is 0",
			args: args{
				price:       100,
				clientPoint: 10,
			},
			wants: wants{
				price:            0,
				discount:         0.9,
				clientPointFloor: 100,
				clientPoint:      10,
			},
		},
		{
			name: "default discount and clientPointFloor but not achieve clientPointFloor",
			args: args{
				price:       100,
				clientPoint: 10,
			},
			wants: wants{
				price:            20,
				discount:         0.9,
				clientPointFloor: 100,
				clientPoint:      10,
			},
		},
		{
			name: "assign new discount and new clientPointFloor but not achieve clientPointFloor",
			args: args{
				opts: []VIPOption{
					WithDiscount(0.8),
					WithPointFloor(20),
				},
				price:       100,
				clientPoint: 10,
			},
			wants: wants{
				price:            20,
				discount:         0.8,
				clientPointFloor: 20,
				clientPoint:      10,
			},
		},
		{
			name: "assign new discount and clientPointFloor and achieve clientPointFloor",
			args: args{
				opts: []VIPOption{
					WithDiscount(0.8),
					WithPointFloor(20),
				},
				price:       100,
				clientPoint: 21,
			},
			wants: wants{
				price:            16,
				discount:         0.8,
				clientPointFloor: 20,
				clientPoint:      1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVIPRedeem(mockRedemption, tt.args.opts...)
			if d := got.GetDiscount(); d != tt.wants.discount {
				t.Errorf("name = %s, discount = %v, want %v", tt.name, d, tt.wants.discount)
			}
			if f := got.GetPointFloor(); f != tt.wants.clientPointFloor {
				t.Errorf("name = %s, clientPointFloor = %v, want %v", tt.name, f, tt.wants.clientPointFloor)
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
