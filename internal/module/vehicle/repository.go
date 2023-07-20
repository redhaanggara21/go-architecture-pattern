package vehicle

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"red21.id/learn/bengkel/domain"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.VehicleRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

func (r repository) FindById(ctx context.Context, id int64) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{"id": id})
	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

func (r repository) FindByVin(ctx context.Context, vin string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{"vin": vin})
	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

func (r repository) Insert(ctx context.Context, vehicle *domain.Vehicle) error {
	excutor := r.db.Insert("vehicles").
		Rows(goqu.Record{
			"vin":       vehicle.VIN,
			"brand":     vehicle.Brand,
			"CreatedAt": vehicle.CreatedAt,
			"UpdatedAt": vehicle.UpdatedAt,
		}).Returning("id").Executor()

	_, err := excutor.ScanStructContext(ctx, &vehicle)
	return err
}
