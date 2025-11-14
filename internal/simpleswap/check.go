package simpleswap

import (
	"fmt"
	"net/http"
	"encoding/json"
	"coin/internal/utils"
)

func CheckExchanges(pubId string) (status *CheckStatus) {
	apiUrl := fmt.Sprintf("https://simpleswap.io/api/v4/exchanges-new/%s", pubId)
	resp, err := http.Get(apiUrl)
	if utils.CheckError(err) {
		return nil
	}
	defer resp.Body.Close()

	var rp CheckResponse
	err = json.NewDecoder(resp.Body).Decode(&rp)
	if utils.CheckError(err) {
		return nil
	}

	info := rp.Result["extendedInfo"].(map[string]interface{})

	txid := ""
	if t, ok := info["txFrom"].(string); ok { txid = t }
	return &CheckStatus{
		Status: rp.Result["status"].(string),
		Address: info["addressFrom"].(string),
		Txid: txid,
	}
}