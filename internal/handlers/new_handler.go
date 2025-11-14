package handlers

import (
	"coin/internal/simpleswap"
	"github.com/labstack/echo/v4"
)

func NewTransaction(c echo.Context) error {
	crypto := c.QueryParam("crypto")
	amount := c.QueryParam("amount")

	if crypto == "" || amount == "" {
		return c.JSON(400, map[string]string{
			"status": "error",
			"message": "crypto, amount 값을 반드시 채워주세요.",
		})
	}

	pubId := simpleswap.NewExchange(crypto, amount)
	if pubId == "" {
		return c.JSON(500, map[string]string{
			"status": "error",
			"message": "새로운 거래를 생성중 에러가 발생했습니다.",
		})
	}

	return c.JSON(200, map[string]string{
		"status": "success",
		"pubId": pubId,
	})
}