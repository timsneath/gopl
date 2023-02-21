package main

import (
	"fmt"
	"math"
)

func main() {
	s := 0
	addr := &s

	flt := math.Pi
	pFlt := &flt

	fmt.Println(addr, pFlt)
	fmt.Println(*pFlt)
}
