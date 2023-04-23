package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK");
}
func findVideo(c *fiber.Ctx) error {
	fmt.Println("Hit video route")
	video, err := os.ReadFile("../videos/rain.mp4");
    if err != nil {
        log.Fatal(err)
    }
	return c.Status(fiber.StatusOK).SendStream(bytes.NewReader([]byte(video)))

}

func main(){
	app := fiber.New();

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	app.Get("/healthcheck", healthcheck)
	app.Get("/video", findVideo)
	fmt.Println("Hello from main");
	log.Fatal(app.Listen(":3000"));
}