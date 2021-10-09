package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 12
	age := "22"
	ageInt, _ := strconv.Atoi(age) //El _ es para desechar variables.
	ageInt = ageInt + 1
	fmt.Println("Niao es " + strconv.Itoa(number) + " Y el sapoperro tiene " + strconv.Itoa(ageInt))
}
