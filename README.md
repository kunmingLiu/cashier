# cashier

`cashier` provides different payment methods and it's easy to customize.

## Install

```
go get github.com/kunmingliu/cashier
```

## Quick Start

### Normal payment method

```go
p := paymentmethod.NewNormalPayment()
c := cashier.NewCashier(cashier.WithPaymentMethod(p))

change, point, err = c.Transact(100, 50)
```

### VIP payment method

```go
v := vip.NewVIP2()
p := paymentmethod.NewVIPPayment(v)
c := cashier.NewCashier(cashier.WithPaymentMethod(p))
```

You can use `WithDiscount` to override the default discount when initialing.

### Point redemption method

```go
r := redemption.NewRedeem()
p := paymentmethod.NewRedemptionPayment(r)
c := cashier.NewCashier(cashier.WithPaymentMethod(p), cashier.WithClientPoint(100))
```

You can use `WithRate` to override the default rate when initialing.

Please remember to set how many points user owns when using the method.

### Point and VIP redemption method (new)

```go
r := redemption.NewVIPRedeem(r)
p := paymentmethod.NewRedemptionPayment(p)
c := cashier.NewCashier(cashier.WithPaymentMethod(vipRedeemPayment), cashier.WithClientPoint(100))
```

You can use `WithPointFloor` and `WithDiscount` to override the default point floor and discount when initialing.

Please remember to set how many points user owns when using the method.

The method is extended by `Point redemption method`.
