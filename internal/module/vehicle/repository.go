package vehicle

import (
	"bengkel_app/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.VehicleRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

func (r repository) FindByID(ctx context.Context, id string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

func(r *repository) FindByVIN(ctx context.Context, vin string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{
		"vin": vin,
	})

	if _, err = dataset.ScanStructContext(ctx, &vehicle); err != nil {
		return domain.Vehicle{}, err
	}
	return
}

func (r *repository) Insert(ctx context.Context, vehicle *domain.Vehicle) error {
	executor := r.db.Insert("vehicles").Rows(goqu.Record{
		"vin": vehicle.VIN,
		"brand": vehicle.Brand,
		"updated_at": vehicle.UpdatedAt,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, vehicle)
	return err
}