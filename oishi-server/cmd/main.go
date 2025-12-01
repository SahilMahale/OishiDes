package main

import (
	"fmt"
	"os"

	"github.com/SahilMahale/OishiDes/oishi-server/internal/db"
	"github.com/SahilMahale/OishiDes/oishi-server/server"
)

// @title OishiDes Docs
// @version 1.0
// @description Swagger docs for OishiDes server
// @contact.name API Support
// @license.name Apache 2.0
// @host localhost:8001
// @BasePath /
const totalTIckets = 50

func main() {
	db, err := db.NewDBConnection()
	if err != nil {
		panic(err)
	}
	ipAddrNPort := os.Getenv("SERVER_BIND_TO")
	if ipAddrNPort == "" {
		ipAddrNPort = "localhost:8001"
	}
	fmt.Println("Staring server....")
	bookingService := server.NewBookingService("OishiDes Server", ipAddrNPort, totalTIckets, db)
	bookingService.StartBookingService()
}
