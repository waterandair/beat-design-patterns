package adapter

// PayAdapter 交易接口适配器
type PayAdapter interface {
	// 对外提供统一的 Pay() 支付接口, channel 表示要使用的支付渠道
	Pay(channel string)
}

var _ PayAdapter = (*UnityPay)(nil)

func NewUnipay() PayAdapter {
	return &UnityPay{
		&AliPay{},
		&WeChatPay{},
	}
}

type UnityPay struct {
	AliPayer
	WeChatPayer
}

func (t *UnityPay) Pay(channel string) {
	switch channel {
	case "wechatpay":
		t.WeChatPay()
	case "alipay":
		t.AliPay()
	}
}
