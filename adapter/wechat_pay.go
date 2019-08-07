package adapter

// WeChatPay 微信支付接口
type WeChatPayer interface {
	WeChatPay()
}

type WeChatPay struct {
}

func (a *WeChatPay) WeChatPay() {

}
