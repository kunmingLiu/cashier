package paymentmethod

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	mockRedemption "github.com/kunmingliu/cashier/internal/mocks/redemption"
	"github.com/kunmingliu/cashier/pkg/redemption"
)

func TestRedemptionPayment(t *testing.T) {
	type args struct {
		redeemer    redemption.Redeemer
		payment     float64
		price       float64
		clientPoint uint
	}

	type wants struct {
		change      float64
		clientPoint uint
		fail        bool
		err         string
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRedemption := mockRedemption.NewMockRedeemer(ctl)
	mockRedemption.EXPECT().GetPrice(100.0, uint(30)).Return(30.0, uint(20)).AnyTimes()

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "user has enough money",
			args: args{
				redeemer:    mockRedemption,
				payment:     40,
				price:       100,
				clientPoint: 30,
			},
			wants: wants{
				change:      10,
				clientPoint: 20,
				fail:        false,
			},
		},
		{
			name: "user has no enough money",
			args: args{
				redeemer:    mockRedemption,
				payment:     10,
				price:       100,
				clientPoint: 30,
			},
			wants: wants{
				change:      -20,
				clientPoint: 20,
				fail:        true,
				err:         fmt.Sprintf("payment is not enough, you are short by %f", 20.0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRedemptionPayment(mockRedemption)
			output, err := got.Transact(PaymentMethodInput{
				Payment:     tt.args.payment,
				Price:       tt.args.price,
				ClientPoint: tt.args.clientPoint,
			})

			if tt.wants.fail {
				if err.Error() != tt.wants.err {
					t.Errorf("name = %s, err = %v, want %v", tt.name, err, tt.wants.err)
				}
			} else {
				if err != nil {
					t.Errorf("name = %s, err = %v, want %v", tt.name, err, nil)
				}
			}

			if change := output.Change; change != tt.wants.change {
				t.Errorf("name = %s, change = %v, want %v", tt.name, change, tt.wants.change)
			}

			if clientPoint := output.RemainingPoint; clientPoint != tt.wants.clientPoint {
				t.Errorf("name = %s, clientPoint = %v, want %v", tt.name, clientPoint, tt.wants.clientPoint)
			}
		})
	}
}
