package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	res := <-timer.C
	fmt.Println(res)
}
