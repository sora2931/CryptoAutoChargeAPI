package handlers

import (
	"coin/internal/simpleswap"
	"github.com/labstack/echo/v4"
)

func CheckTransaction(c echo.Context) error {
	id := c.QueryParam("id")

	if id == "" {
		return c.JSON(400, map[string]string{
			"status": "error",
			"message": "id 값을 반드시 채워주세요.",
		})
	}

	status := simpleswap.CheckExchanges(id)
	if status == nil {
		return c.JSON(500, map[string]string{
			"status": "error",
			"message": "거래를 조회하던중 문제가 발생했습니다.",
		})
	}

	return c.JSON(200, map[string]interface{} {
		"status": "success",
		"data": map[string]string{
			"status": status.Status,
			"address": status.Address,
			"txid": status.Txid,
		},
	})
}