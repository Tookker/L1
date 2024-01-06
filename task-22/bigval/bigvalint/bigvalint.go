package bigvalint

import (
	"L1/task-22/bigval"
	"errors"
	"math/big"
)

var (
	errCast = "Cast error!"
)

// Реализации интерфейса БигВал
type BigValInt struct {
	Val *big.Int
}

// Преобразование строки в интерфейс БигВал
func SetValFromString(str string) (bigval.BigVal, error) {
	temp := new(big.Int)
	temp, ok := temp.SetString(str, 10)
	if !ok {
		return nil, errors.New(errCast)
	}

	return &BigValInt{
		Val: temp,
	}, nil
}

// Установить значение БигИнт в интерфейс БигВал
func SetVal(val *big.Int) bigval.BigVal {
	return &BigValInt{
		Val: val,
	}
}

// Преобразование БигВал в строку
func ToString(val bigval.BigVal) string {
	cast, _ := val.(*BigValInt)
	return cast.Val.String()
}
