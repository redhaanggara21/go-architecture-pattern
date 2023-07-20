package customer

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"red21.id/learn/bengkel/domain"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.CustomerRepository.
func (r repository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Order(goqu.I("name").Asc())
	if err := dataset.ScanStructsContext(ctx, &customers); err != nil {
		return nil, err
	}
	return
}

// FindById implements domain.CustomerRepository.
func (r repository) FindById(ctx context.Context, id int64) (customers domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &customers)
	return
}

func (r repository) FindByIds(ctx context.Context, ids []int64) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": ids,
	})
	_, err = dataset.ScanStructContext(ctx, &customers)
	return
}

// FindByPhone implements domain.CustomerRepository.
func (r repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"phone": phone,
	})
	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, nil
	}
	return domain.Customer{}, nil
}

// Insert implements domain.CustomerRepository.
func (r repository) Insert(ctx context.Context, customer *domain.Customer) error {
	excutor := r.db.Insert("customers").Rows(*customer).Returning("id").Executor()
	var customerDb domain.Customer
	_, err := excutor.ScanStructContext(ctx, &customerDb)
	if err != nil {
		return err
	}
	customer.ID = customerDb.ID
	return err
}
