package customer

import (
	"bengkel_app/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{db: goqu.New("default", con)}
}

func (re repository) All(ctx context.Context) (customers []domain.Customer, err error) {
	dataset := re.db.From("customers").Order(goqu.I("name").Asc())

	if err := dataset.ScanStructsContext(ctx, &customers); err != nil {
		return nil, err
	}

	return
}

func (re repository) FindById(ctx context.Context, id int64) (customer domain.Customer, err error) {
	dataset := re.db.From("customers").Where(goqu.Ex{
		"id": id,
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}

	return
}

func (re repository) FindByIds(ctx context.Context, ids []int64) (customers []domain.Customer, err error) {
	dataset := re.db.From("customers").Where(goqu.Ex{
		"id": ids,
	})

	if err := dataset.ScanStructsContext(ctx, &customers); err != nil {
		return nil, err
	}

	return
}

func (re repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer, err error) {
	dataset := re.db.From("customers").Where(goqu.Ex{
		"phone": phone,
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}

	return
}

func (re repository) Insert(ctx context.Context, customer *domain.Customer) error {
	executor := re.db.Insert("customers").Rows(*customer).Returning("id").Executor()

	var customerDB domain.Customer

	_, err := executor.ScanStructContext(ctx, &customerDB)
	if err != nil {
		return err
	}

	customer.ID = customerDB.ID
	return err

}
