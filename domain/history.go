package domain

import "time"

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

type HistoryService interface {
}

type HistoryRepository interface {
}
