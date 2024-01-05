package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.

// Интерфейс записи в общую бд
type TelephoneDB interface {
	Write(key string, value uint)
}

// Реализация интерфейса через map
type TelephoneDBMap struct {
	mutex sync.Mutex
	mapDb map[string]uint
}

// Функция создания новой мапы
func NewDBMap() TelephoneDB {
	return &TelephoneDBMap{
		mutex: sync.Mutex{},
		mapDb: make(map[string]uint),
	}
}

// Запись данных в мапу используя  mutex, чтобы избежать гонок
func (t *TelephoneDBMap) Write(key string, value uint) {
	defer t.mutex.Unlock()
	t.mutex.Lock()
	t.mapDb[key] = value
}

func (t *TelephoneDBMap) GetInfo() map[string]uint {
	copyMap := make(map[string]uint)
	for key, value := range t.mapDb {
		copyMap[key] = value
	}

	return copyMap
}

// Запись данных в мапу
func writeToMap(db TelephoneDB, key string, value uint) {
	db.Write(key, value)
}

func main() {
	db := NewDBMap()

	go writeToMap(db, "000000", 000000)
	go writeToMap(db, "111111", 111111)
	go writeToMap(db, "222222", 222222)
	go writeToMap(db, "333333", 333333)
	go writeToMap(db, "444444", 444444)
	go writeToMap(db, "555555", 555555)
	go writeToMap(db, "666666", 666666)
	go writeToMap(db, "777777", 777777)
	go writeToMap(db, "888888", 888888)
	go writeToMap(db, "999999", 999999)

	time.Sleep(time.Second * 3)

	newDb := db.(*TelephoneDBMap).GetInfo()
	newDb["111111"] = 12356776

	fmt.Println(db.(*TelephoneDBMap).GetInfo())
	fmt.Println(newDb)
}
