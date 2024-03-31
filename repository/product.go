package repository

import (
    "database/sql"
    "example.com/m/model"
)

type ProductRepository interface {
    FindByID(id int) (*model.Product, error)
    Save(Product *model.Product) error
    Update(Product *model.Product) error
    Delete(id int) error
	GetAllProducts() ([]*model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}


func (r *productRepository) FindByID(id int) (*model.Product, error) {
    return nil, nil
}

func (r *productRepository) Save(Product *model.Product) error {
    return nil
}

func (r *productRepository) Update(Product *model.Product) error {
    return nil
}

func (r *productRepository) Delete(id int) error {
    return nil
}
func (r *productRepository) GetAllProducts() ([]*model.Product, error) {
	rows, err := r.db.Query("SELECT * FROM Products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var Products []*model.Product
    for rows.Next() {
        Product := &model.Product{}
        err := rows.Scan(&Product.ID, &Product.ProductName, &Product.ProductDescription, &Product.Price)
        if err != nil {
            return nil, err
        }
        Products = append(Products, Product)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return Products, nil
}