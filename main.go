package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"net/http"
)

// ua TelegramBot (like TwitterBot)

func main() {

	e := echo.New()
	// e.Debug = true
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, " ___       __      ___    ___ ________  ___  ___      ________  ________     \n|\\  \\     |\\  \\   |\\  \\  /  /|\\   ____\\|\\  \\|\\  \\    |\\   ____\\|\\   ____\\    \n\\ \\  \\    \\ \\  \\  \\ \\  \\/  / | \\  \\___|\\ \\  \\\\\\  \\   \\ \\  \\___|\\ \\  \\___|    \n \\ \\  \\  __\\ \\  \\  \\ \\    / / \\ \\_____  \\ \\  \\\\\\  \\   \\ \\_____  \\ \\  \\  ___  \n  \\ \\  \\|\\__\\_\\  \\  /     \\/ __\\|____|\\  \\ \\  \\\\\\  \\ __\\|____|\\  \\ \\  \\|\\  \\ \n   \\ \\____________\\/  /\\   \\|\\__\\____\\_\\  \\ \\_______\\\\__\\____\\_\\  \\ \\_______\\\n    \\|____________/__/ /\\ __\\|__|\\_________\\|_______\\|__|\\_________\\|_______|\n                  |__|/ \\|__|   \\|_________|            \\|_________|         \n                                                                             \n                                                                             \nOpen Graph Meta Proxy\n\nUsage: wx.su.sg/https://mp.weixin.qq.com/s/xxx\nor wx.su.sg/mp.weixin.qq.com/s/xxx \nor wx.su.sg/s/xxx\nor wx.su.sg/xxx\n\nExample:\nwx.su.sg/https://mp.weixin.qq.com/s/9COs4RUL7v8TTeqgrONsoA\nwx.su.sg/9COs4RUL7v8TTeqgrONsoA")
	})
	e.GET("/robots.txt", func(c echo.Context) error {
		return c.String(http.StatusOK, "User-agent: *\nDisallow: /")
	})
	e.GET("/*", Handler)

	e.Logger.Fatal(e.Start("127.0.0.1:7233"))

}
