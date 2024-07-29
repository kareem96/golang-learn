package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		Prefork: true, //prefork
	})

	app.Use("/api", func (ctx *fiber.Ctx) error {
		fmt.Println("im middleware before processing request")
		err := ctx.Next()
		fmt.Println("im middleware after processing request")
		return err
	})

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	if fiber.IsChild(){
		fmt.Println("im child process")
	}else{
		fmt.Println("im parent process")
	}

	err := app.Listen("localhost:3000")
	if err != nil{
		panic(err)	
	}
}


