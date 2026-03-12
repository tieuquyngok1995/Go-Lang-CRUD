package product

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, p *Product) error
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(ctx context.Context) ([]Product, error) {
	products := make([]Product, 0)
	query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products ORDER BY id`
	err := r.db.SelectContext(ctx, &products, query)
	return products, err
}

func (r *repository) FindByID(ctx context.Context, id int64) (*Product, error) {
	var p Product
	query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = @p1`
	if err := r.db.GetContext(ctx, &p, query, id); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) Create(ctx context.Context, p *Product) error {
	query := `
		INSERT INTO products (name, description, price, stock)
		OUTPUT INSERTED.id, INSERTED.created_at, INSERTED.updated_at
		VALUES (@p1, @p2, @p3, @p4)`
	return r.db.QueryRowContext(ctx, query, p.Name, p.Description, p.Price, p.Stock).
		Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *repository) Update(ctx context.Context, p *Product) error {
	query := `
		UPDATE products
		SET name = @p1, description = @p2, price = @p3, stock = @p4, updated_at = GETDATE()
		WHERE id = @p5`
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.Price, p.Stock, p.ID)
	return err
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM products WHERE id = @p1`, id)
	return err
}
