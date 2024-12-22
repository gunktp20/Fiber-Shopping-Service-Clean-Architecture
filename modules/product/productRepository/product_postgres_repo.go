package repository

import (
	"fmt"
	productDto "shopping-service-be/modules/product/productDto"
	productEntity "shopping-service-be/modules/product/productEntity"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type productPostgresRepository struct {
	db *sqlx.DB
}

func NewProductPostgresRepository(db *sqlx.DB) ProductRepositoryService {
	return &productPostgresRepository{db}
}

func (r *productPostgresRepository) CreateOneProduct(createProductReq *productDto.CreateProductReq) (*productEntity.Product, error) {
	// ? sqlx query row x is not support named parameters (:name , :description , :price ) that why we use $ instead
	var query string
	var args []interface{}

	if createProductReq.ProductImage != "" {
		query = `
			INSERT INTO products (name, description, price, image_url) 
			VALUES ($1, $2, $3, $4) 
			RETURNING id, name, description, price, image_url, created_at, updated_at
		`
		args = []interface{}{createProductReq.Name, createProductReq.Description, createProductReq.Price, createProductReq.ProductImage}
	} else {
		query = `
			INSERT INTO products (name, description, price) 
			VALUES ($1, $2, $3) 
			RETURNING id, name, description, price, created_at, updated_at
		`
		args = []interface{}{createProductReq.Name, createProductReq.Description, createProductReq.Price}
	}

	product := new(productEntity.Product)

	err := r.db.QueryRowx(query, args...).StructScan(product)
	if err != nil {
		return nil, err
	}

	return product, nil

}

func (r *productPostgresRepository) GetAllProducts(keyword string) ([]*productEntity.Product, error) {
	// Base query
	query := `SELECT id, name, description, price, image_url, created_at, updated_at FROM products`
	args := map[string]interface{}{}

	if keyword != "" {
		query += ` WHERE name ILIKE :keyword OR description ILIKE :keyword`
		args["keyword"] = "%" + keyword + "%"
	}

	var products []*productEntity.Product

	nstmt, err := r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}
	defer nstmt.Close()

	for nstmt.Next() {
		product := new(productEntity.Product)
		if err := nstmt.StructScan(product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productPostgresRepository) DeleteProductById(id string) error {
	query := `DELETE FROM products WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return fmt.Errorf("no product found with id %s", id)
	}

	return nil
}

func (r *productPostgresRepository) UpdateProductById(selectedId string, updateProductReq *productDto.UpdateProductReq) (*productEntity.Product, error) {

	query := `UPDATE Products SET name = $2, description = $3, price = $4 WHERE id = $1;`
	args := []interface{}{selectedId, updateProductReq.Name, updateProductReq.Description, updateProductReq.Price}

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return &productEntity.Product{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &productEntity.Product{}, err
	}

	if rowsAffected == 0 {
		return &productEntity.Product{}, fmt.Errorf("no rows were updated, product with id %s not found", selectedId)
	}

	id, err := strconv.ParseInt(selectedId, 10, 64)
	if err != nil {
		return &productEntity.Product{}, fmt.Errorf("invalid id format: %v", err)
	}

	product := new(productEntity.Product)
	product.ID = id
	product.Name = updateProductReq.Name
	product.Description = updateProductReq.Description
	product.Price = updateProductReq.Price

	return product, nil
}
