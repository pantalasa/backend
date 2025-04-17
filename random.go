package main

import (
	"fmt"
	"math/rand"
)

func CheckRandom() {
	r := rand.Intn(100)
	if r < 50 {
		fmt.Println("r < 50")
	} else {
		fmt.Println("5 >= 50")
	}
}
