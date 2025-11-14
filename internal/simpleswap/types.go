package simpleswap

type Payload struct {
	AddressTo string `json:"addressTo"`
	Amount string `json:"amount"`
	Campaign *string `json:"campaign"`
	ExtraIdTo string `json:"extraIdTo"`
	FirstPage string `json:"firstPage"`
	Fixed bool `json:"fixed"`
	IsExternalProvider bool `json:"isExternalProvider"`
	Medium string `json:"medium"`
	NetworkFrom string `json:"networkFrom"`
	NetworkTo string `json:"networkTo"`
	NotProvideMemo bool `json:"notProvideMemo"`
	NotProvideRefundMemo bool `json:"notProvideRefundMemo"`
	Referral *string `json:"referral"`
	Reverse bool `json:"reverse"`
	Source string `json:"source"`
	TickerFrom string `json:"tickerFrom"`
	TickerTo string `json:"tickerTo"`
	UnregEmail string `json:"unregEmail"`
	UserRefundAddress string `json:"userRefundAddress"`
	UserRefundExtraId string `json:"userRefundExtraId"`
}
type ExchangeResponse struct {
	Result map[string]interface{} `json:"result"`
}


type CheckStatus struct {
	Status string
	Address string
	Txid string
}

type CheckResponse struct {
	Result map[string]interface{} `json:"result"`
}