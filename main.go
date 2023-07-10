package main

import (
	"fmt"

	"github.com/yuttasakcom/go-reflection-simple/report"
)

type Person struct {
	name string `report:"ชื่อ,uppercase"`
	age  int    `report:"อายุ"`
}

type Employee struct {
	name string `report:"ชื่อ"`
	age  int    `report:"อายุ"`
}

func main() {
	fmt.Println(report.Text(Person{"Yuttasak", 40}))
	fmt.Println(report.Text(Employee{"Yea", 35}))
	fmt.Println(report.Text(struct {
		name string `report:"ชื่อ"`
		age  int    `report:"อายุ"`
	}{"Sri", 36}))
}
