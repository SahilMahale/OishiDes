package models

import "time"

type UserTicketsResponse struct {
	Username      string
	TicketsBooked uint
}

type UserSignup struct {
	Username string `json:"user" xml:"user" form:"user"`
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"pass" xml:"pass" form:"pass"`
	IsAdmin  bool   `json:"isadmin" xml:"isadmin" form:"isadmin"`
}
type UserSignin struct {
	Username string `json:"user" xml:"user" form:"user"`
	Password string `json:"pass" xml:"pass" form:"pass"`
}

type AdminSignup struct {
	Username string `json:"user" xml:"user" form:"user"`
	Password string `json:"pass" xml:"pass" form:"pass"`
}
type AdminSignin struct {
	Username string `json:"user" xml:"user" form:"user"`
	Password string `json:"pass" xml:"pass" form:"pass"`
}

type TicketsResponse struct {
	TicketsLeft uint
}
type TableRequest struct {
	TableID uint `json:"tableID" xml:"tableID" form:"tableID"`
	Seats   uint `json:"seats" xml:"seats" form:"seats"`
	NonVeg  bool `json:"nonVeg" xml:"nonVeg" form:"nonVeg"`
	AirCon  bool `json:"ac" xml:"ac" form:"ac"`
}
type TableResponse struct {
	BookingRef string `json:"bookingID" xml:"bookingID" form:"bookingID"`
	NonVeg     bool   `json:"nonVeg" xml:"nonVeg" form:"nonVeg"`
	AirCon     bool   `json:"ac" xml:"ac" form:"ac"`
	TableID    uint   `json:"tableID" xml:"tableID" form:"tableID"`
	Seats      uint   `json:"seats" xml:"seats" form:"seats"`
}

type BookingsResponse struct {
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
	BookingID string    `json:"bookingID" xml:"bookingID" form:"bookingID"`
	Username  string    `json:"user" xml:"user" form:"user"`
	Status    string    `json:"status" xml:"status" form:"status"`
}

type BookingsRequest struct {
	Username string `json:"user" xml:"user" form:"user"`
	Tables   []uint `json:"tables" xml:"tables" form:"tables"`
}
type PaymentRequest struct {
	PaymentID  string `json:"paymentID" xml:"paymentID" form:"paymentID"`
	Method     string `json:"method" xml:"method" form:"method"`
	BookingRef string `json:"bookingRef" xml:"bookingRef" form:"bookingRef"`
	AmtPaid    uint   `json:"amtPaid" xml:"amtPaid" form:"amtPaid"`
}

type PaymentResponse struct {
	PaymentID  string `json:"paymentID" xml:"paymentID" form:"paymentID"`
	Method     string `json:"method" xml:"method" form:"method"`
	Status     string `json:"status" xml:"status" form:"status"`
	BookingRef string `json:"bookingRef" xml:"bookingRef" form:"bookingRef"`
	AmtPaid    uint   `json:"amtPaid" xml:"amtPaid" form:"amtPaid"`
}
