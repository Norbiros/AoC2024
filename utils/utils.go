package utils

import (
	"strconv"
)

func ToInt(arg string) int {
	var val, err = strconv.Atoi(arg)
	if err != nil {
		panic("error converting string to int " + err.Error())
	}
	return val
}
