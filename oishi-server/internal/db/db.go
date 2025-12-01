package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConnection struct {
	Db *gorm.DB
}

// type DbInterface interface {
// }

func NewDBConnection() (DbConnection, error) {
	uName := os.Getenv("MYSQL_USERNAME")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbIP := os.Getenv("MYSQL_IP")
	if uName == "" || pass == "" || dbIP == "" {
		panic(fmt.Errorf("one or more ENV variables required for DB connection are not set"))
	}
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/OishiDatabase?charset=utf8mb4&parseTime=True&loc=Local", uName, pass, dbIP)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return DbConnection{}, err
	}
	// Automigrate and create the tables
	db.AutoMigrate(User{})
	db.AutoMigrate(Bookings{})
	db.AutoMigrate(Table{})
	db.AutoMigrate(Cancellation{})
	db.AutoMigrate(Payment{})
	// db.AutoMigrate(Admins{})
	return DbConnection{
		Db: db,
	}, nil
}
