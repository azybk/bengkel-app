package domain

import (
	"context"
	"time"
)

type HistoryDetail struct {
	ID          int64     `db:"id"`
	VehicleID   int64     `db:"vehicle_id"`
	PIC         string    `db:"pic"`
	CreatedAt   time.Time `db:"created_at"`
	Notes       string    `db:"notes"`
	CustomerID  int64     `db:"customer_id"`
	PlateNumber string    `db:"plate_number"`
}

type HistoryRepository interface {
	FindByVehicle(ctx context.Context, id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, detail *HistoryDetail) error
}

type HistoryService interface {
}

type HistoryData struct {
	VehicleID   int64  `json:"vehicle_id"`
	PIC         string `json:"pic"`
	Notes       string `json:"notes"`
	CustomerID  int64  `json:"cutomer_id"`
	PlateNumber string `json:"plate_number"`
	ComeAt      string `json:"come_at"`
}
