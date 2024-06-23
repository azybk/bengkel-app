package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID        int64     `db:"id"`
	Name      string    `db:"Name"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

type CustomerRepository interface {
	FindById(ctx context.Context, id int64) (Customer, error)
	FindByIds(ctx context.Context, ids []int64) ([]Customer, error)
	FindByPhone(ctx context.Context, phone string) (Customer, error)
	Insert(ctx context.Context, customer *Customer) error
}

type CustomerService interface {
}
