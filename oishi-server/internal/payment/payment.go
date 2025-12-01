package payment

import (
	"github.com/SahilMahale/OishiDes/oishi-server/constants"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/db"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/helper"
	"github.com/SahilMahale/OishiDes/oishi-server/server/models"
)

type PaymentController struct {
	DbInterface db.DbConnection
}

type PaymentOPs interface {
	CreatePayment(payreq models.TableRequest) (*models.TableResponse, *helper.MyHTTPErrors)
	DeletePayment(tableID uint, bookID string) *helper.MyHTTPErrors
	GetPayments() (*[]models.TableResponse, *helper.MyHTTPErrors)
	GetPaymentInfo(tableID uint) (*models.TableResponse, *helper.MyHTTPErrors)
	LinkPayment(tableID uint) (db.Table, *helper.MyHTTPErrors)
	ClearBookingPossesion(bookingID string) *helper.MyHTTPErrors
}

func NewTablesController(db db.DbConnection) PaymentController {
	return PaymentController{
		DbInterface: db,
	}
}

func (p PaymentController) CreatePayment(payreq models.PaymentRequest) (*models.PaymentResponse, *helper.MyHTTPErrors) {
	payment := db.Payment{
		PaymentID:  payreq.PaymentID,
		BookingRef: payreq.BookingRef,
		Method:     payreq.Method,
		AmtPaid:    payreq.AmtPaid,
		Status:     constants.INITIATED,
	}
	qr := p.DbInterface.Db.Create(payment)
	if qr.Error != nil {
		err := helper.ErrorMatch(qr.Error)
		return nil, &err
	}
	payResp := models.PaymentResponse{
		PaymentID:  payment.PaymentID,
		BookingRef: payment.BookingRef,
		Method:     payment.Method,
		AmtPaid:    payment.AmtPaid,
		Status:     payment.Status,
	}
	return &payResp, nil
}
