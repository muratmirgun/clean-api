package main

import "github.com/labstack/echo/v4"

func main() {
	instance := echo.New()
	instance.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	err := instance.Start(":8080")
	if err != nil {
		return
	}
}

