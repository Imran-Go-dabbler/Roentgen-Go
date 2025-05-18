package main

import (
	"fmt"
	"math"
)

func main() {
	var invest = 1000
	var interest = 5.6
	var time = 10

	var final = float64(invest) * math.Pow(1+interest/100, float64(time))

	fmt.Println("Investment amount: ", final)

}
