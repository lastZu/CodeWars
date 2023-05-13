package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 4; i > 0; i-- {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(i) + 1
		fmt.Println(num)
	}

}
