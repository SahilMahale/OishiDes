package constants

type UserRole string

const (
	Admin UserRole = "admin"
	User  UserRole = "user"
)

const (
	Booked     string = "booked"
	Cancelled  string = "cancelled"
	Paid       string = "paid"
	Served     string = "served"
	InProgress string = "inProgress"
	COMPLETED  string = "completed"
	INITIATED  string = "initiated"
	FAILED     string = "failed"
)
