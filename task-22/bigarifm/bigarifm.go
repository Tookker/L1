package bigarifm

import "L1/task-22/bigval"

// Интерфейс арифмитических операций
type BigArifm interface {
	Sum(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error)
	Sub(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error)
	Multiplication(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error)
	Division(val1 bigval.BigVal, val2 bigval.BigVal) (bigval.BigVal, error)
}
