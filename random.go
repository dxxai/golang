package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MyRand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min

}

func main() {
	ranRange := MyRand(1, 3)
	fmt.Println(ranRange)

}
