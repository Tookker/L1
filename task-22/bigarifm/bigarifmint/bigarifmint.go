package bigarifmint

import (
	"L1/task-22/bigarifm"
	"L1/task-22/bigval"
	"L1/task-22/bigval/bigvalint"
	"errors"
	"fmt"
)

var (
	errCast = "Error cast"
)

// Реализация арифмитических операций под БигИнт
type BigArifmInt struct{}

func NewArifm() bigarifm.BigArifm {
	return &BigArifmInt{}
}

// Сумма
func (b *BigArifmInt) Sum(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error) {
	castVal1, ok := checkCastVal(val1)
	if !ok {
		return nil, errors.New(errCast)
	}

	castVal2, ok := checkCastVal(val2)
	if !ok {
		return nil, errors.New(errCast)
	}

	return bigvalint.SetVal(castVal1.Val.Add(castVal1.Val, castVal2.Val)), nil
}

// Вычитание
func (b *BigArifmInt) Sub(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error) {
	castVal1, ok := checkCastVal(val1)
	if !ok {
		return nil, errors.New(errCast)
	}

	castVal2, ok := checkCastVal(val2)
	if !ok {
		return nil, errors.New(errCast)
	}

	return bigvalint.SetVal(castVal1.Val.Sub(castVal1.Val, castVal2.Val)), nil
}

// Умножение
func (b *BigArifmInt) Multiplication(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error) {
	castVal1, ok := checkCastVal(val1)
	if !ok {
		return nil, errors.New(errCast)
	}

	castVal2, ok := checkCastVal(val2)
	if !ok {
		return nil, errors.New(errCast)
	}

	return bigvalint.SetVal(castVal1.Val.Mul(castVal1.Val, castVal2.Val)), nil
}

// Деление
func (b *BigArifmInt) Division(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error) {
	castVal1, ok := checkCastVal(val1)
	if !ok {
		return nil, errors.New(errCast)
	}

	castVal2, ok := checkCastVal(val2)
	if !ok {
		return nil, errors.New(errCast)
	}

	return bigvalint.SetVal(castVal1.Val.Div(castVal1.Val, castVal2.Val)), nil
}

// Проверка что в интерфейс передан БигИнт
func checkCastVal(val bigval.BigVal) (bigvalint.BigValInt, bool) {
	res, ok := val.(*bigvalint.BigValInt)
	if !ok {
		fmt.Println(errCast)
		return bigvalint.BigValInt{}, false
	}

	return *res, true
}
