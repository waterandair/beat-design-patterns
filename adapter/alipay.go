package adapter

// AliTransaction 支付宝支付接口
type AliPayer interface {
	AliPay()
}

type AliPay struct {

}

func(a *AliPay) AliPay() {

}
