package repository

import (
	"database/sql"

	"github.com/VadimBoganov/fulgur/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRespository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query("select * from products")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err = rows.Scan(&product.Id, &product.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) GetAllTypes() ([]domain.ProductType, error) {
	return nil, nil
}

func (r *ProductRepository) GetAllSubTypes() ([]domain.ProductSubType, error) {
	return nil, nil
}

func (r *ProductRepository) Insert(products []domain.Product) error {
	query := "INSERT INTO products (id, name) VALUES "
	var vals []interface{}

	id := 1
	for _, product := range products {
		query += "(?, ?),"
		vals = append(vals, id, product.Name)
		id++
	}

	query = query[0 : len(query)-1]

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}

	return err
}

func (r *ProductRepository) RemoveById(id int) error {
	_, err := r.db.Exec("delete from products where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) RemoveAll() error {
	_, err := r.db.Exec("delete from products")
	if err != nil {
		return err
	}

	return nil
}
