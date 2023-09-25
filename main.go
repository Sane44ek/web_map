package main

import (
	"fmt"
	"mymap/mp"
)

func main() {
	m := mp.NewMap()
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("tree", 3)
	m.Add("four", 4)
	m.Add("aaa", 2)
	m.Add("bbb", 3)
	m.Add("ccc", 4)
	m.Add("ddd", 2)
	m.Add("eee", 3)
	m.Add("fff", 4)
	m.Print()
	length := m.Len() // Получаем длину карты
	fmt.Println("Length:", length)
	m.RemoveByIndex("two")
	m.Print()
	m.RemoveByValue(2)
	m.Print()
	m.RemoveAllByValue(3)
	m.Print()
	fmt.Println(m.GetByIndex("ccc"))
	fmt.Println(m.GetByValue(4))
	fmt.Println(m.GetAllByValue(4))
	fmt.Println(m.GetAll())
	m.Clear()
	m.Print()
}