package simpleswap

import (
	"bytes"
	"net/http"
	"encoding/json"
	"coin/internal/utils"
)

func buildPayload(cryptoType, amount string) (payload *Payload) {
	return Payload{
		AddressTo: utils.Cfg.Address,
		Amount: amount,
		Campaign: nil,
		ExtraIdTo: "",
		FirstPage: "https://simpleswap.io/",
		Fixed: false,
		IsExternalProvider: false,
		Medium: "organic",
		NetworkFrom: cryptoType,
		NetworkTo: utils.Cfg.Crypto,
		NotProvideMemo: false,
		NotProvideRefundMemo: false,
		Referral: nil,
		Reverse: false,
		Source: "firefox",
		TickerFrom: cryptoType,
		TickerTo: utils.Cfg.Crypto,
		UnregEmail: "",
		UserRefundAddress: "",
		UserRefundExtraId: "",
	}
}

func postJSON(apiUrl string, payload *Payload) (resp *http.Response) {
	data, err := json.Marshal(payload)
	if utils.CheckError(err) {
		return nil
	}

	resp, err = http.Post(apiUrl, "application/json", bytes.NewBuffer(data))
	if utils.CheckError(err) {
		return nil
	}
	return resp
}

func NewExchange(cryptoType string, amount string) (pubId string) {
	payload := buildPayload(cryptoType, amount)
	resp:= postJSON("https://simpleswap.io/api/v4/exchanges-new", payload)
	if resp == nil {
		return ""
	}
	defer resp.Body.Close()

	var rp ExchangeResponse
    err := json.NewDecoder(resp.Body).Decode(&rp)
	if utils.CheckError(err) {
		return ""
	}

	return rp.Result["publicId"].(string)
}
