package bookings

import (
	"fmt"

	"github.com/SahilMahale/OishiDes/oishi-server/constants"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/db"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/helper"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/tables"
	"github.com/SahilMahale/OishiDes/oishi-server/server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type BookingsController struct {
	DbInterface db.DbConnection
}

type BookingsOps interface {
	CreateBooking(username string, tables []uint) (*models.BookingsResponse, helper.MyHTTPErrors)
	DeleteBooking(bookingID string) helper.MyHTTPErrors
	GetBookings() ([]db.Bookings, helper.MyHTTPErrors)
	GetBookingsForUser(username string) ([]db.Bookings, helper.MyHTTPErrors)
}

func NewBookingController(db db.DbConnection) BookingsController {
	return BookingsController{
		DbInterface: db,
	}
}

func (b BookingsController) CreateBooking(username string, tablesToBook []uint) (*models.BookingsResponse, helper.MyHTTPErrors) {
	tbInt := tables.NewTablesController(b.DbInterface)
	bookid := uuid.NewString()
	booking := db.Bookings{
		BookingID:     bookid,
		UsernameRefer: username,
		Status:        constants.InProgress,
	}
	query := b.DbInterface.Db.Create(&booking)
	if query.Error != nil {
		myerr := helper.ErrorMatch(query.Error)
		return nil, myerr
	}
	bookingresp := models.BookingsResponse{
		BookingID: bookid,
		Status:    booking.Status,
		Username:  booking.UsernameRefer,
		CreatedAt: booking.CreatedAt,
		UpdatedAt: booking.UpdatedAt,
	}
	tblen := len(tablesToBook)
	tbarr := make([]*db.Table, tblen)
	errarr := make([]*helper.MyHTTPErrors, tblen)
	var bookFails uint
	for i, tabId := range tablesToBook {
		/* fmt.Println("-------------------------------------------------")
		fmt.Printf("Booking table: %d with bookid: %s\n", tabId, bookid)
		fmt.Println("-------------------------------------------------") */
		tbarr[i], errarr[i] = tbInt.BookTable(tabId, bookid)
		/* 		fmt.Printf("func resp: %v err: %v\n", tbarr[i], errarr[i]) */
		if errarr[i] != nil {
			bookFails++
			log.Errorf("Booking of Table: %s for Booking:%s with the Error:%s", tbarr[i].TableID, tbarr[i].BookingRef,
				errarr[i].Err.Error())
		}
	}
	if bookFails > 1 && tblen != 1 {
		return &bookingresp, helper.MyHTTPErrors{
			Err:      fmt.Errorf("Some tables failed to book"),
			HttpCode: fiber.StatusMultipleChoices,
		}
	}
	if bookFails == uint(tblen) {
		booking.Status = "failed"
		err := b.DbInterface.Db.Save(&booking)
		if err.Error != nil {
			return nil, helper.ErrorMatch(err.Error)
		}
		return nil, helper.MyHTTPErrors{
			Err:      fmt.Errorf("Booking Tables failed"),
			HttpCode: fiber.ErrConflict.Code,
		}
	}
	booking.Status = constants.Booked
	err := b.DbInterface.Db.Save(&booking)
	if err.Error != nil {
		return nil, helper.ErrorMatch(err.Error)
	}
	bookingresp.Status = booking.Status
	return &bookingresp, helper.MyHTTPErrors{
		Err: nil,
	}
}

func (b BookingsController) DeleteBooking(bookingID string) helper.MyHTTPErrors {
	booking := db.Bookings{BookingID: bookingID}
	tabCont := tables.NewTablesController(b.DbInterface)
	// check if entry is present
	err := b.DbInterface.Db.First(&booking)

	if err.Error != nil {
		myerr := helper.ErrorMatch(err.Error)
		return myerr
	}
	err = b.DbInterface.Db.Delete(&booking)

	if err.Error != nil {
		myerr := helper.ErrorMatch(err.Error)
		return myerr
	}
	cerr := tabCont.ClearBookingPossesion(bookingID)
	if cerr != nil {
		return *cerr
	}
	return helper.MyHTTPErrors{
		Err: nil,
	}
}

func (b BookingsController) GetBookings() ([]db.Bookings, helper.MyHTTPErrors) {
	var book []db.Bookings

	res := b.DbInterface.Db.Find(&book)
	if res.Error != nil {
		myerr := helper.ErrorMatch(res.Error)
		return []db.Bookings{}, myerr
	}

	return book, helper.MyHTTPErrors{
		Err: nil,
	}
}

func (b BookingsController) GetBookingsForUser(username string) ([]db.Bookings, helper.MyHTTPErrors) {
	var book []db.Bookings

	res := b.DbInterface.Db.Where(&db.Bookings{UsernameRefer: username}).Find(&book)
	if res.Error != nil {
		myerr := helper.ErrorMatch(res.Error)
		return []db.Bookings{}, myerr
	}

	return book, helper.MyHTTPErrors{
		Err: nil,
	}
}
