package main

import (
	"fmt"

	temp "sneath.org/gopl/ch2/ex2.1"
)

func main() {
	fmt.Println(temp.Celsius(100))
	fmt.Printf("%v\n", temp.Fahrenheit(32))
	fmt.Println(temp.Fahrenheit(temp.CtoF(temp.BoilingWaterC)))
}
