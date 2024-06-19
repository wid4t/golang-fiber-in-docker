package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	pb "golang-fiber-in-docker/proto"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"

	"go-micro.dev/v4"
)

var (
	serviceName = "golang-micro-api"
	version     = "latest"
)

func main() {

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
	)

	service.Init()

	client := pb.NewGolangMicroProductService("golang-micro-product", service.Client())

	app := fiber.New()

	api := app.Group("/module/partner", logger.New())

	api.Get("/check", func(c *fiber.Ctx) error {

		name := c.Query("name", "good people")

		req := &pb.CallRequest{
			Name: name,
		}

		resp, err := client.Call(context.Background(), req)

		if err != nil {
			log.Fatal(err)
		}

		return c.SendString("Hello, i'm from golang-fiber-in-docker, data from golang-micro-product : " + resp.Msg)
	})

	log.Fatal(app.Listen(":3000"))
}
