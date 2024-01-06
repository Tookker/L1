package menu

import (
	"errors"
	"fmt"
	"os"

	"L1/task-22/bigarifm"
	"L1/task-22/bigarifm/bigarifmint"
	"L1/task-22/bigval"
	"L1/task-22/bigval/bigvalint"
)

// Структура меню
type Menu struct{}

const (
	inputMsg           = "Enter %d value"
	choseTypeMsg       = "Enter 1 to work with BigInt"
	choseTypeArifmOper = "Enter 1 for +\nEnter 2 for -\nEnter 3 for *\nEnter 4 for /"
)

const (
	errInputMsg = "Wrong input"
)

type typeVal int

const (
	errInput typeVal = iota
	bigInt   typeVal = iota
)

// Основной цикл в котором осузествляет запрос из std::in
func (m *Menu) StartMenu() {
	for {
		typeV, err := m.choseType()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf(inputMsg+"\n", 1)
		val1, err := m.showEnterVals(typeV)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf(inputMsg+"\n", 2)
		val2, err := m.showEnterVals(typeV)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		res, err := m.chooseArifmOper(val1, val2, typeV)
		if err != nil {
			continue
		}

		fmt.Println("Result =", m.bigVaLToString(res, typeV))
	}
}

// Определение типа больших чисел (реализовано только Инт, позже можно добавить работу с флоатом)
func (m *Menu) choseType() (typeVal, error) {
	fmt.Println(choseTypeMsg)

	var str string
	fmt.Fscan(os.Stdin, &str)
	switch str {
	case "1":
		return bigInt, nil
	default:
		return errInput, errors.New(errInputMsg)
	}
}

// Ввод чисел
func (m *Menu) showEnterVals(typeV typeVal) (bigval.BigVal, error) {
	var str string
	fmt.Fscan(os.Stdin, &str)

	switch typeV {
	case bigInt:
		res, err := bigvalint.SetValFromString(str)
		if err != nil {
			fmt.Println(errInputMsg)
			return nil, err
		}

		return res, nil
	}

	return nil, nil
}

// Выбор арифмитической операции
func (m *Menu) chooseArifmOper(val1 bigval.BigVal, val2 bigval.BigVal, typeV typeVal) (bigval.BigVal, error) {
	fmt.Println(choseTypeArifmOper)
	var str string
	fmt.Fscan(os.Stdin, &str)

	var arifm bigarifm.BigArifm
	if typeV == bigInt {
		arifm = bigarifmint.NewArifm()
	}

	switch str {
	case "1":
		res, err := arifm.Sum(val1, val2)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return res, nil
	case "2":
		res, err := arifm.Sub(val1, val2)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return res, nil
	case "3":
		res, err := arifm.Multiplication(val1, val2)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return res, nil
	case "4":
		res, err := arifm.Division(val1, val2)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return res, nil
	}

	fmt.Println(errInputMsg)
	return nil, errors.New(errInputMsg)
}

// Преобразование результата в строку
func (m *Menu) bigVaLToString(val bigval.BigVal, typeV typeVal) string {
	var res string

	switch typeV {
	case bigInt:
		res = bigvalint.ToString(val)
	}

	return res
}
