package ports

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/ceblay/payments/pkg/app"
	"github.com/ceblay/payments/pkg/app/command"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

type NewGatewayRequest struct {
	Name string `json:"name"`
}

type IssuePaymentRequest struct {
	Amount float32 `json:"amount"`
}

func (h HttpServer) Run() {
	_app := fiber.New()

	_app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "HTTP Server is ready",
		})
	})

	_app.Get("/gateways", func(c *fiber.Ctx) error {
		result, err := h.app.Queries.AllGateways.Handle()
		if err != nil {
			return c.SendString("An error occurred")
		}
		return c.JSON(result)
	})

	_app.Post("/gateways", func(c *fiber.Ctx) error {
		dto := &NewGatewayRequest{}
		if err := c.BodyParser(dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Ensure the payload conforms to the expected",
			})
		}

		cmd := command.AddGateway{
			Name: dto.Name,
		}

		err := h.app.Commands.AddGateway.Handle(cmd)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Something unexpected happened",
			})
		}
		c.Status(fiber.StatusCreated)
		return nil
	})

	_app.Get("/ipns", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Webhook for payment processing called.",
		})
	})

	_app.Post("/payments", func(c *fiber.Ctx) error {
		dto := &IssuePaymentRequest{}
		if err := c.BodyParser(dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Ensure the payload conforms to the expected",
			})
		}

		cmd := command.IssuePayment{
			Amount: dto.Amount,
		}

		err := h.app.Commands.IssuePayment.Handle(cmd)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "An error occurred while processing payment",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Payment successfully made",
		})
	})

	log.Fatal(_app.Listen(":7000"))
}
