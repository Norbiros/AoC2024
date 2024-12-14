package utils

import (
	"strconv"
)

type Pair struct {
	First, Second int
}

func ToInt(arg string) int {
	var val, err = strconv.Atoi(arg)
	if err != nil {
		panic("error converting string to int " + err.Error())
	}
	return val
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
