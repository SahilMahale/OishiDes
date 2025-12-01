package server

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/SahilMahale/OishiDes/oishi-server/internal/bookings"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/db"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/helper"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/user"
	"github.com/SahilMahale/OishiDes/oishi-server/server/models"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	jwtware "github.com/gofiber/contrib/jwt"
)

// @title OishiDes Docs
// @version 1.0
// @description Swagger docs for OishiDes server
// @contact.name API Support
// @license.name Apache 2.0
// @host localhost:8001
// @BasePath /
var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type MyCustomClaims struct {
	Name string `json:"name"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

type bookingService struct {
	app          *fiber.App
	DbInterface  db.DbConnection
	ip           string
	totalTickets uint
}
type BookingServicer interface {
	GetBookings(c *fiber.Ctx) error
	GetTickets(c *fiber.Ctx) error
	DeleteBooking(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	BookTables(c *fiber.Ctx) error
	GetB(c *fiber.Ctx) error
	PayForBooking(c *fiber.Ctx) error
	CancelBooking(c *fiber.Ctx) error
	GetUserInfo(c *fiber.Ctx) error
	GetAllTables(c *fiber.Ctx) error
	CreatePayment(c *fiber.Ctx) error
	CreateCancellation(c *fiber.Ctx) error
	UpdatePayment(c *fiber.Ctx) error
	DeletePayment(c *fiber.Ctx) error
	UpdateCancellation(c *fiber.Ctx) error
	DeleteCancellation(c *fiber.Ctx) error
	CreateTables(c *fiber.Ctx) error
	UpdateTables(c *fiber.Ctx) error
	DeleteTables(c *fiber.Ctx) error
	StartBookingService()
}

func NewBookingService(appname, ip string, totalTickets uint, db db.DbConnection) bookingService {
	return bookingService{
		app: fiber.New(fiber.Config{
			AppName:       appname,
			StrictRouting: true,
			ServerHeader:  "Bookings",
		}),
		ip:           ip,
		totalTickets: totalTickets,
		DbInterface:  db,
	}
}

func (B *bookingService) initMiddleware() {
	// Adding logger to the app
	B.app.Use(requestid.New())
	B.app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	B.app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	B.app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,http://localhost:4000,http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
}

func (B *bookingService) initAuth() {
	secretsFolderPath := os.Getenv("APP_AUTH")
	if secretsFolderPath == "no-auth" || secretsFolderPath == "" {
		// run app without jwt auth
		return
	}
	privateKeyPath := fmt.Sprintf("%s/private_key.pem", secretsFolderPath)
	publicKeyPath := fmt.Sprintf("%s/private_key.pem.pub", secretsFolderPath)
	err := readPrivateKeyFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	err = readPublicKeyFile(publicKeyPath)
	if err != nil {
		panic(err)
	}
	B.app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    publicKey,
		},
		ContextKey: "acces-key-token",
	}))
}

func (B *bookingService) GetBookings(c *fiber.Ctx) error {
	authToken := getAuthToken(c)
	var bookarr []db.Bookings
	var err helper.MyHTTPErrors
	bookCtrl := bookings.NewBookingController(B.DbInterface)
	userNameSpecified := c.Query("user")
	claims, errp := getClaimsForThisCall(authToken)
	if errp != nil {
		panic(errp)
	}
	isAdmin := checkIfAdmin(claims.Type)

	if userNameSpecified == "" {
		if isAdmin {
			bookarr, err = bookCtrl.GetBookings()
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Not an admin")
		}
	} else {
		if userNameSpecified == claims.Name || isAdmin {
			bookarr, err = bookCtrl.GetBookingsForUser(userNameSpecified)
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Not authorized to make this request")
		}
	}

	if err.Err != nil {
		if err.HttpCode == fiber.StatusMultipleChoices {
			return c.Status(err.HttpCode).JSON(bookarr)
		}
		return c.Status(err.HttpCode).SendString(err.Err.Error())
	}
	bookRespArr := []models.BookingsResponse{}
	for _, book := range bookarr {
		bookEntry := models.BookingsResponse{
			BookingID: book.BookingID,
			Username:  book.UsernameRefer,
			Status:    book.Status,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		}
		bookRespArr = append(bookRespArr, bookEntry)
	}
	return c.JSON(bookRespArr)
}

func (B *bookingService) CreateUser(c *fiber.Ctx) error {
	var userCtrl user.UserOps
	u := new(models.UserSignup)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	userCtrl = user.NewUserController(B.DbInterface)

	errP := userCtrl.CreateUser(u.Username, u.Email, u.Password, u.IsAdmin)
	if errP.Err != nil {
		return c.Status(errP.HttpCode).SendString(errP.Err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (B *bookingService) LoginUser(c *fiber.Ctx) error {
	var userCtrl user.UserOps
	u := new(models.UserSignin)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	userCtrl = user.NewUserController(B.DbInterface)

	role, err := userCtrl.LoginUser(u.Username, u.Password)
	if err.Err != nil {
		return c.Status(err.HttpCode).SendString(err.Err.Error())
	}

	// Create a token based on user
	atoken, errp := makeTokenWithClaims(checkIfAdmin(role), u.Username)

	if errp != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(errp.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"auth_token": atoken})
}

func (B *bookingService) GetAllUsers(c *fiber.Ctx) error {
	authToken := getAuthToken(c)

	claims, errp := getClaimsForThisCall(authToken)
	if errp != nil {
		panic(errp)
	}
	isAdmin := checkIfAdmin(claims.Type)
	if !isAdmin {
		return c.Status(fiber.StatusUnauthorized).SendString("Operation only reserved for Admins")
	}
	userCtrl := user.NewUserController(B.DbInterface)
	usersList, err := userCtrl.GetAllUsers()
	if err.Err != nil {
		return c.Status(err.HttpCode).SendString(err.Err.Error())
	}
	return c.Status(fiber.StatusAccepted).JSON(usersList)
}

func (B *bookingService) BookTables(c *fiber.Ctx) error {
	authToken := getAuthToken(c)

	claims, errp := getClaimsForThisCall(authToken)
	if errp != nil {
		panic(errp)
	}
	isAdmin := checkIfAdmin(claims.Type)

	var bookCtrl bookings.BookingsController
	book := new(models.BookingsRequest)

	if err := c.BodyParser(book); err != nil {
		return err
	}
	if (book.Username != claims.Name) && !isAdmin {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized to book ticket for this user")
	}
	if len(book.Tables) < 1 {
		return c.Status(fiber.StatusBadRequest).SendString("Please Select Tables to book, Array is Empty")
	}
	if len(book.Tables) > 2 && !isAdmin {
		return c.Status(fiber.StatusBadRequest).SendString("A user can only book 2 Tables at once")
	}
	if len(book.Tables) > 5 {
		return c.Status(fiber.StatusBadRequest).SendString("Only 4 Tables can be booked at once")
	}
	bookCtrl = bookings.NewBookingController(B.DbInterface)
	bookings, err := bookCtrl.CreateBooking(book.Username, book.Tables)
	if err.Err != nil {
		return c.Status(err.HttpCode).SendString(err.Err.Error())
	}

	bookResp := models.BookingsResponse{
		BookingID: bookings.BookingID,
		Username:  bookings.Username,
		Status:    bookings.Status,
		CreatedAt: bookings.CreatedAt,
		UpdatedAt: bookings.UpdatedAt,
	}
	return c.Status(fiber.StatusAccepted).JSON(bookResp)
}

func (B *bookingService) DeleteBooking(c *fiber.Ctx) error {
	authToken := getAuthToken(c)

	bookID := ""
	if bookID = c.Params("bookid"); bookID == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Need to specify 'bookingId' param")
	}
	user := ""
	if user = c.Query("userName"); user == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Need to specify 'userName' param")
	}

	claims, errp := getClaimsForThisCall(authToken)
	if errp != nil {
		panic(errp)
	}
	isAdmin := checkIfAdmin(claims.Type)

	if user != claims.Name && !isAdmin {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized to delete booking for this user")
	}
	bookCtrl := bookings.NewBookingController(B.DbInterface)
	err := bookCtrl.DeleteBooking(bookID)

	if err.Err != nil {
		return c.Status(err.HttpCode).SendString(err.Err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("booking deleted")
}

func (B *bookingService) StartBookingService() {
	B.initMiddleware()
	// Unauthenticated routes
	userGroup := B.app.Group("/user")
	userGroup.Post("/signup", B.CreateUser)
	userGroup.Post("/signin", B.LoginUser)

	B.app.Get("/swagger/*", swagger.HandlerDefault)
	adminGroup := B.app.Group("/admin")
	adminGroup.Post("/signup", B.CreateUser)
	adminGroup.Post("/signin", B.LoginUser)

	B.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Booking APP Service is Running!")
	})

	B.initAuth()
	// authenticated routes
	userGroup.Get("/info", B.GetAllUsers)
	bookingGroup := B.app.Group("/bookings")
	// legacy concept testing routes need to remove
	bookingGroup.Get("", B.GetBookings)
	bookingGroup.Post("", B.BookTables)
	//--
	bookingGroup.Post("/tables", B.BookTables)
	bookingGroup.Delete("/:bookID", B.DeleteBooking)

	// Swagger Docs
	err := B.app.Listen(B.ip)
	if err != nil {
		log.Error(err)
		return
	}
}
