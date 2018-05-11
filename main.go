package main

import (
	"fmt"

	"github.com/aaronarduino/adder/api"
)

func main() {
	//a := api.NewBin4(1, 0, 1, 0)
	b := api.NewBin4(0, 1, 1, 1)
	//c := a.Add(b)
	fmt.Println(b)
}