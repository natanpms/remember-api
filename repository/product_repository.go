package repository

import (
	"database/sql"
	"fmt"

	"github.com/natanpms/remember-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price, created_at FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObject model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price,
			&productObject.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObject)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	const query = `
		INSERT INTO product (product_name, price)
		VALUES ($1, $2)
		RETURNING id, product_name, price, created_at`

	var created model.Product

	err := pr.connection.QueryRow(query, product.Name, product.Price).Scan(
		&created.ID,
		&created.Name,
		&created.Price,
		&created.CreatedAt,
	)

	if err != nil {
		return model.Product{}, err
	}

	return created, nil
}

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
	const query = `
		SELECT id, product_name, price, created_at
		FROM product
		WHERE id = $1`

	var product model.Product

	err := pr.connection.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.CreatedAt,
	)
	fmt.Printf("Produto encontrado 3: %+v\n", product)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, fmt.Errorf("Produto não encontrado")
		}

		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	const query = `
		DELETE FROM product
		WHERE id = $1`

	result, err := pr.connection.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("Erro ao realizar consulta.")
	}

	fmt.Printf("Linhas afetadas: %d\n", rows)

	if rows == 0 {
		return fmt.Errorf("Produto não encontrado")
	}

	return nil
}
