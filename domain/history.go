package domain

import (
	"context"
	"time"
)

type History struct {
	ID        int64     `db:"id"`
	NoRangka  string    `db:"no_rangka"`
	Merek     string    `db:merek`
	UpdatedAt time.Time `db:updated_at`
}

type HistoryDetail struct {
	ID         int64     `db:"id"`
	HistoryID  int64     `db:"history_id"`
	PIC        string    `db:"pic"`
	CreatedAt  time.Time `db:created_at`
	Notes      string    `db:"notes"`
	CustomerID int64     `db:"cutomer_id"`
	PlatNomor  string    `db:"plat_nomor"`
}


type HistoryRepository interface {
	FindById(ctx context.Context, id int64) (History, error)
	FindByNoRangka(ctx context.Context, no string) (History, error)
	FindDetailByHistory(ctx context.Context, id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *History) error
	InsertDetail(ctx context.Context, detail *HistoryDetail) error
}

type HistoryService interface {
}
