package server

import (
	"fmt"

	"github.com/SahilMahale/Booking-App/booking-server/server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type bookingService struct {
	totalTickets uint
	app          *fiber.App
	ip           string
}
type BookingServicer interface {
	GetBookings(c *fiber.Ctx) error
	GetTickets(c *fiber.Ctx) error
	BookTickets(c *fiber.Ctx) error
	DeleteBooking(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	StartBookingService()
}

func NewBookingService(appname, ip string, totalTickets uint) bookingService {
	return bookingService{
		app: fiber.New(fiber.Config{
			AppName:       appname,
			StrictRouting: true,
			ServerHeader:  "Bookings",
		}),
		ip:           ip,
		totalTickets: totalTickets,
	}
}

func (B *bookingService) initLogger() {
	// Adding logger to the app
	B.app.Use(requestid.New())
	B.app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
}

func (B *bookingService) GetTickets(c *fiber.Ctx) error {
	if user := c.Query("user"); user != "" {
		userData := models.UserTicketsResponse{Username: user, TicketsBooked: 5}
		return c.JSON(userData)
	}
	return c.JSON(models.TicketsResponse{TicketsLeft: B.totalTickets})
}

func (B *bookingService) GetBookings(c *fiber.Ctx) error {
	if user := c.Query("user"); user != "" {
		return c.SendString(fmt.Sprintf("%s has %d bookings\n", user, 3))
	}
	return c.SendString("will return all the bookings")
}

func (B *bookingService) CreateUser(c *fiber.Ctx) error {
	u := new(models.UserSignup)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}

func (B *bookingService) LoginUser(c *fiber.Ctx) error {
	u := new(models.UserSignin)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(u)
}

func (B *bookingService) BookTickets(c *fiber.Ctx) error {
	book := new(models.BookingsRequest)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(book)
}

func (B *bookingService) DeleteBooking(c *fiber.Ctx) error {
	if bookID := c.Params("bookid"); bookID != "" {
		return c.SendString(fmt.Sprintf("%s booking is delete\n", bookID))
	}
	return c.Status(fiber.StatusBadRequest).SendString("Need the bookingID")
}

func (B *bookingService) StartBookingService() {

	B.initLogger()

	B.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Booking APP!")
	})

	userGroup := B.app.Group("/user")
	userGroup.Post("/signup", B.CreateUser)
	userGroup.Post("/signin", B.LoginUser)

	ticketGroup := B.app.Group("/tickets")
	ticketGroup.Get("", B.GetTickets)

	bookingGroup := B.app.Group("/bookings")
	bookingGroup.Get("", B.GetBookings)
	bookingGroup.Post("", B.BookTickets)
	bookingGroup.Delete("/:bookID", B.DeleteBooking)

	B.app.Listen(B.ip)
}