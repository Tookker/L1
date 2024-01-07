package main

import (
	"errors"
	"fmt"
	"strconv"
)

//8.Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

type BitVal bool

const (
	ZERO BitVal = false
	ONE  BitVal = true
)

func bitInt64Swap(bit int64, val int64, bitVal BitVal) (int64, error) {
	var errInputBit = "Error bit input"

	if bit >= 512 {
		return 0, errors.New(errInputBit)
	}

	mask := int64(1 << bit)

	if bitVal {
		return val | mask, nil
	}

	return val &^ mask, nil
}

func main() {
	val := int64(10)

	fmt.Println(strconv.FormatInt(val, 2))
	res, err := bitInt64Swap(7, val, ONE)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strconv.FormatInt(res, 2))
}
