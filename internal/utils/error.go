package utils

import (
	"log"
)

func CheckError(err error) (status bool) {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}