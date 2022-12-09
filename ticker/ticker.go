package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	ticker := time.NewTicker(2 * time.Second)
	i := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				i++
				fmt.Println(<-ticker.C)
				if i == 4 {
					ticker.Stop()
				}
			}

		}
	}()
	e.Logger.Fatal(e.Start(":8086"))
}
