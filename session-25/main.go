package main

import (
	"github.com/gofiber/fiber/v2"
	"session-25/internal/handlers"
)

func main() {

	//cmd := exec.Command("./mainprocess")
	//err := cmd.Start()
	//if err != nil {
	//	fmt.Println("Failed to start process:", err)
	//} else {
	//	fmt.Println("Process started with PID:", cmd.Process.Pid)
	//}

	app := fiber.New()
	handler := handlers.Handler{}

	app.Get("/", handler.AddUser)

	app.Listen(":3001")
}
