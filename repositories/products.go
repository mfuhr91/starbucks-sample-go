package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"starbucks-app/database"
	"starbucks-app/models"
)

type productRepository struct{}

type ProductRepository interface {
	GetAll(bool) ([]models.Product, error)
	GetById(id string) (models.Product, error)
	Save(models.Product) (models.Product, error)
	Update(models.Product) (models.Product, error)
	Delete(id string) error
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (c *productRepository) GetAll(allowEmptyProducts bool) ([]models.Product, error) {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return []models.Product{}, err
	}
	defer db.Close()
	
	query := fmt.Sprintf("SELECT * FROM products WHERE disabled = false")
	
	if !allowEmptyProducts {
		query += fmt.Sprintf(" AND quantity != 0")
	}
	
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("error when getting the products from db, %s", err)
		return []models.Product{}, err
	}
	
	var products []models.Product
	for rows.Next() {
		var product models.Product
		
		err = rows.Scan(&product.ID, &product.Quantity, &product.Name,
			&product.Price, &product.Disabled)
		if err != nil {
			log.Printf("error when scanning the products row from db, %s", err)
			return []models.Product{}, err
		}
		
		fmt.Println(product)
		products = append(products, product)
	}
	
	return products, nil
}

func (c *productRepository) Save(product models.Product) (models.Product, error) {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Product{}, err
	}
	
	defer db.Close()
	insertStat, err := db.Prepare("INSERT INTO products (id, quantity, name, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Product{}, err
	}
	
	product.ID = uuid.New().String()
	
	_, err = insertStat.Exec(product.ID, product.Quantity, product.Name, product.Price)
	if err != nil {
		log.Printf("error when saving the product into db, %s", err.Error())
		return models.Product{}, err
	}
	
	return product, nil
}

func (c *productRepository) GetById(id string) (models.Product, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Product{}, err
	}
	
	defer db.Close()
	selectStat, err := db.Prepare("SELECT * FROM products WHERE id=$1")
	if err != nil {
		log.Printf("error when getting the products from db")
		return models.Product{}, err
	}
	
	var product models.Product
	
	err = selectStat.QueryRow(id).Scan(&product.ID, &product.Quantity, &product.Name,
		&product.Price, &product.Disabled)
	if err != nil {
		log.Printf("error when saving the product into db, %s", err.Error())
		return models.Product{}, err
	}
	
	return product, nil
}

func (c *productRepository) Update(product models.Product) (models.Product, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Product{}, err
	}
	
	defer db.Close()
	updateStat, err := db.Prepare("UPDATE products SET quantity=$1, name=$2, price=$3 WHERE id=$4")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Product{}, err
	}
	
	_, err = updateStat.Exec(product.Quantity, product.Name, product.Price, product.ID)
	if err != nil {
		log.Printf("error when updating the product from db, %s", err.Error())
		return models.Product{}, err
	}
	
	return product, nil
}

func (c *productRepository) Delete(id string) error {
	db, err := database.DatabaseConnect()
	if err != nil {
		return err
	}
	
	defer db.Close()
	_, err = db.Query("UPDATE products SET disabled=true WHERE id=$1", id)
	if err != nil {
		log.Printf("error  when deleting the product from db, %s", err.Error())
		return err
	}
	
	return nil
}
