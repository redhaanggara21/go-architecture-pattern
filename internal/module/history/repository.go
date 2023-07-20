package history

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"red21.id/learn/bengkel/domain"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.HistoryRepository {
	return &repository{db: goqu.New("default", con)}
}

func (r repository) FindAll(ctx context.Context) (history []domain.History, err error) {
	dataset := r.db.From("history").Order(goqu.I("name").Asc())
	if _, err := dataset.ScanStructContext(ctx, &history); err != nil {
		return nil, err
	}
	return
}

func (r repository) FindByVehicle(ctx context.Context, name string) (result []domain.History, err error) {
	dataset := r.db.From("history_details").Where(goqu.Ex{
		"name": name,
	}).Order(goqu.I("id").Asc())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (r repository) Insert(ctx context.Context, detail *domain.HistoryDetail) error {
	detail.CreatedAt = time.Now()
	detail.UpdatedAt = time.Now()
	executor := r.db.Insert("history").Rows(goqu.Record{
		"vehicle_id":  detail.VehicleID,
		"customer_id": detail.CustomerID,
		"notes":       detail.Notes,
		"pic":         detail.PIC,
		"createdAt":   detail.CreatedAt,
		"updatedAt":   detail.UpdatedAt,
	}).Returning("id").Executor()
	_, err := executor.ScanStructContext(ctx, detail)
	return err
}
