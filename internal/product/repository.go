package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/agustinrabini/Gocker/internal/database"
	"github.com/agustinrabini/Gocker/internal/domain"
)

//CRUDE
type Repository interface {
	GetAll(ctx context.Context) (domain.Products, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: database.DbConecction(),
	}
}

func (r *repository) GetAll(ctx context.Context) (domain.Products, error) {

	var products domain.Products
	var product domain.Product

	db := r.db

	result, err := r.db.Query("select * from product")
	if err != nil {
		fmt.Print("Error en la consulta", err.Error())
	}

	for result.Next() {

		err := result.Scan(&product.Id_product, &product.Name, &product.Brand, &product.Description, &product.Image, &product.Price, &product.Stock)
		if err != nil {

			fmt.Print("Error en la consulta", err.Error())
			return nil, err
		}

		products = append(products, product)
	}

	defer result.Close()
	defer db.Close()

	return products, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {

	var product domain.Product

	dbProducts := r.db

	result, err := dbProducts.Query("select * from product where id_product = ?", id)
	if err != nil {

		fmt.Print("Error en la consulta", err.Error())
		return domain.Product{}, err
	}

	for result.Next() {

		err := result.Scan(&product.Id_product, &product.Name, &product.Brand, &product.Description, &product.Image, &product.Price, &product.Stock)
		if err != nil {

			fmt.Print("Error en la consulta", err.Error())
			return domain.Product{}, err
		}
	}

	defer dbProducts.Close()
	defer result.Close()

	return product, nil
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {

	db := r.db

	stmt, err := db.Prepare("insert into product (id_product, name, brand, description, image, price, stock) values (?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(nil, p.Name, p.Brand, p.Description, p.Image, p.Price, p.Stock)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	defer db.Close()

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {

	db := r.db

	stmt, err := db.Prepare("update product set id_product=?, name=?, blade=?, brand=?, description=?, image=?, price=?, length=?, stock=? where id_product = ? ")
	if err != nil {

		return err
	}

	res, err := stmt.Exec(p.Id_product, p.Name, p.Brand, p.Description, p.Image, p.Price, p.Stock, p.Id_product)
	if err != nil {

		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return err
	}

	defer stmt.Close()
	defer db.Close()

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {

	db := r.db

	stmt, err := db.Prepare("delete from product where id_product = ? ")
	if err != nil {

		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {

		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return err
	}

	defer stmt.Close()
	defer db.Close()

	return nil
}
