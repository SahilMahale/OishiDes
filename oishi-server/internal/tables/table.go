package tables

import (
	"fmt"

	"github.com/SahilMahale/OishiDes/oishi-server/internal/db"
	"github.com/SahilMahale/OishiDes/oishi-server/internal/helper"
	"github.com/SahilMahale/OishiDes/oishi-server/server/models"
	"github.com/gofiber/fiber/v2"
)

type TablesController struct {
	DbInterface db.DbConnection
}

type TablesOPs interface {
	CreateTables(tabreq models.TableRequest) (*models.TableResponse, *helper.MyHTTPErrors)
	DeleteTable(tableID uint, bookID string) *helper.MyHTTPErrors
	GetTables() (*[]models.TableResponse, *helper.MyHTTPErrors)
	GetTableInfo(tableID uint) (*models.TableResponse, *helper.MyHTTPErrors)
	BookTable(tableID uint, bookID string) (db.Table, *helper.MyHTTPErrors)
	ClearBookingPossesion(bookingID string) *helper.MyHTTPErrors
	UpdateTable(table models.TableRequest) (*models.TableResponse, *helper.MyHTTPErrors)
}

func NewTablesController(db db.DbConnection) TablesController {
	return TablesController{
		DbInterface: db,
	}
}
func (t TablesController) BookTable(tableID uint, bookID string) (*db.Table, *helper.MyHTTPErrors) {
	table := db.Table{TableID: tableID}
	qr := t.DbInterface.Db.First(&table)
	if qr.Error != nil {
		myerr := helper.ErrorMatch(qr.Error)
		return nil, &myerr
	}
	if table.BookingRef != "" {
		if table.BookingRef == bookID {
			return &table, nil
		}
		return nil, &helper.MyHTTPErrors{Err: fmt.Errorf("Table already booked")}
	}
	table.BookingRef = bookID
	qr = t.DbInterface.Db.Save(&table)
	if qr.Error != nil {
		myerr := helper.ErrorMatch(qr.Error)
		return nil, &myerr
	}
	return &table, nil
}
func (t TablesController) CreateTables(tabreq models.TableRequest) (*models.TableResponse, *helper.MyHTTPErrors) {
	table := db.Table{
		TableID: tabreq.TableID,
		NonVeg:  tabreq.NonVeg,
		AirCon:  tabreq.AirCon,
		Seats:   tabreq.Seats,
	}
	qr := t.DbInterface.Db.Create(&table)
	if qr.Error != nil {
		err := helper.ErrorMatch(qr.Error)
		return nil, &err
	}
	tbresp := models.TableResponse{
		TableID: table.TableID,
		NonVeg:  table.NonVeg,
		AirCon:  table.AirCon,
		Seats:   table.Seats,
	}
	return &tbresp, nil
}

func (t TablesController) ClearBookingPossesion(bookingID string) *helper.MyHTTPErrors {
	tables := make([]db.Table, 1)
	err := t.DbInterface.Db.Where(&db.Table{BookingRef: bookingID}).Find(&tables)

	if err.Error != nil {
		merr := helper.ErrorMatch(err.Error)
		return &merr
	}
	for _, table := range tables {
		table.BookingRef = ""
		err = t.DbInterface.Db.Save(&table)
		if err.Error != nil {
			merr := helper.ErrorMatch(err.Error)
			return &merr
		}
	}
	return &helper.MyHTTPErrors{
		Err:      nil,
		HttpCode: fiber.StatusOK,
	}
}
func (t TablesController) GetTableInfo(tableID uint) (*models.TableResponse, *helper.MyHTTPErrors) {
	var table = db.Table{TableID: tableID}
	err := t.DbInterface.Db.Find(&table)
	if err.Error != nil {
		merr := helper.ErrorMatch(err.Error)
		return nil, &merr
	}
	tablesResp := models.TableResponse{}
	tablesResp = models.TableResponse{
		TableID: table.TableID,
		Seats:   table.Seats,
		NonVeg:  table.NonVeg,
		AirCon:  table.AirCon,
	}
	return &tablesResp, nil
}
func (t TablesController) GetTables() (*[]models.TableResponse, *helper.MyHTTPErrors) {
	var tables = make([]db.Table, 1)
	err := t.DbInterface.Db.Find(&tables)
	if err.Error != nil {
		merr := helper.ErrorMatch(err.Error)
		return nil, &merr
	}
	tablesResp := make([]models.TableResponse, 1)
	for i, tval := range tables {
		tablesResp[i] = models.TableResponse{
			TableID: tval.TableID,
			Seats:   tval.Seats,
			NonVeg:  tval.NonVeg,
			AirCon:  tval.AirCon,
		}
	}
	return &tablesResp, nil
}

func (t TablesController) DeleteTable(tableID uint) *helper.MyHTTPErrors {
	table := db.Table{TableID: tableID}
	//check if entry is present
	err := t.DbInterface.Db.First(&table)

	if err.Error != nil {
		myerr := helper.ErrorMatch(err.Error)
		return &myerr
	}

	err = t.DbInterface.Db.Delete(&table)

	if err.Error != nil {
		myerr := helper.ErrorMatch(err.Error)
		return &myerr
	}
	return &helper.MyHTTPErrors{
		Err:      nil,
		HttpCode: fiber.StatusOK,
	}
}

func (t TablesController) UpdateTable(tabReq *models.TableRequest) (*models.TableResponse, *helper.MyHTTPErrors) {
	table := db.Table{TableID: tabReq.TableID}
	err := t.DbInterface.Db.Find(&table)
	if err.Error != nil {
		myerr := helper.ErrorMatch(err.Error)
		return nil, &myerr
	}
	table.Seats = tabReq.Seats
	table.AirCon = tabReq.AirCon
	table.NonVeg = tabReq.NonVeg
	err = t.DbInterface.Db.Save(&table)
	if err.Error != nil {
		myerr := helper.ErrorMatch(err.Error)
		return nil, &myerr
	}
	return &models.TableResponse{
		TableID: table.TableID,
		Seats:   table.Seats,
		NonVeg:  table.NonVeg,
		AirCon:  table.AirCon,
	}, nil
}
