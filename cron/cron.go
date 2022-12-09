package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/robfig/cron"
)

func main() {
	e := echo.New()

	var orderNo string
	var res string
	time.Sleep(5 * time.Second)

	fmt.Println("拿到了订单号")
	fmt.Println("开始执行定时任务")
	orderNo = "fasdfsda"
	if orderNo != "" {
		c := cron.New()
		c.AddFunc("*/2 * * * * ?", func() {
			fmt.Println("业务逻辑")
			res = "ok"
			if res == "ok" {
				c.Stop()
			}
		})
		c.Start()
	}

	e.Logger.Fatal(e.Start(":8085"))
}
