package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID        int64     `db:"id"`
	VIN       string    `vin:"vin"`
	Merk      string    `merk:"merk"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type HistoryDetail struct {
	ID         uuid.UUID `db:"id"`
	CustomerID uuid.UUID `db:"customer_id"`
	VehicleID  uuid.UUID `db:"history_id"`
	PIC        string    `db:"string"`
	VIN        string    `db:"string"`
	Notes      string    `db:"string"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type HistoryRepository interface {
	FindById(ctx context.Context, id int64) (History, error)
	FindByVIN(ctx context.Context, vin string) (History, error)
	FindDetailByHistory(ctx context.Context, id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *HistoryDetail) error
	InsertDetail(ctx context.Context, detail *HistoryDetail) error
}

type HistoryService interface {
}

type HistoryData struct {
	VehicleID   int64  `db:"vehicle_id"`
	CustomerID  int64  `db:"customer_id"`
	PIC         string `db:"pic"`
	PlateNumber string `db:"plate_number"`
	Notes       string `db:"notes"`
}
