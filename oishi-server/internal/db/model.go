package db

import (
	"time"

	"github.com/SahilMahale/OishiDes/oishi-server/constants"
	"gorm.io/gorm"
)

/* Tables with DeletedAt field have soft delete automatically configured with db.Delete*/
type User struct {
	Username  string `gorm:"primaryKey;type:varchar(36);"`
	Email     string
	Pass      string
	Role      constants.UserRole `gorm:"type:enum('admin','user')"`
	CreatedAt time.Time          // Automatically managed by GORM for creation time
	UpdatedAt time.Time          // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt     // field addded for soft Delete
}

type Bookings struct {
	BookingID     string `gorm:"primaryKey;type:varchar(36);"`
	Status        string
	User          User `gorm:"foreignKey:UsernameRefer"`
	UsernameRefer string
	CreatedAt     time.Time      // Automatically managed by GORM for creation time
	UpdatedAt     time.Time      // Automatically managed by GORM for update time
	DeletedAt     gorm.DeletedAt // field addded for soft Delete
}

type Table struct {
	Bookings   Bookings `gorm:"foreignKey:BookingRef"`
	BookingRef string
	NonVeg     bool
	AirCon     bool
	TableID    uint `gorm:"primaryKey"`
	Seats      uint
}

type Cancellation struct {
	CancellationID string    `gorm:"primaryKey"`
	CreatedAt      time.Time // Automatically managed by GORM for creation time
	UpdatedAt      time.Time // Automatically managed by GORM for update time
	Bookings       Bookings  `gorm:"foreignKey:BookingRef"`
	BookingRef     string
	DeletedAt      gorm.DeletedAt // field addded for soft Delete
	RefundAmt      uint
}

type Payment struct {
	PaymentID  string `gorm:"primaryKey"`
	BookingRef string
	Bookings   Bookings       `gorm:"foreignKey:BookingRef"`
	Method     string         `gorm:"type:enum('card','cash','UPI')"`
	Status     string         `gorm:"type:enum('completed','pending','initiated','failed')"`
	CreatedAt  time.Time      // Automatically managed by GORM for creation time
	UpdatedAt  time.Time      // Automatically managed by GORM for update time
	DeletedAt  gorm.DeletedAt // field addded for soft Delete
	AmtPaid    uint
}

// type Admins struct {
// 	Name string `gorm:"primaryKey;type:varchar(36);"`
// 	Pass string
// }
