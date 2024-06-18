package domain

import "time"

type Customer struct {
	ID        int64     `db:"id"`
	Name      string    `db:"Name"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

type CustomerService interface {
}

type CustomerRepository interface {
}
