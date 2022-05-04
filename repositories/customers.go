package repositories

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"starbucks-app/database"
	"starbucks-app/models"
)

type customerRepository struct{}

type CustomerRepository interface {
	GetAll() ([]models.Customer, error)
	GetById(id string) (models.Customer, error)
	Save(models.Customer) (models.Customer, error)
	Update(models.Customer) (models.Customer, error)
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}

func (c *customerRepository) GetAll() ([]models.Customer, error) {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return []models.Customer{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		log.Printf("error when getting the customers from starbucks_db")
		return []models.Customer{}, err
	}
	
	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		
		err = rows.Scan(&customer.ID, &customer.Name, &customer.LastName,
			&customer.Phone, &customer.Address, &customer.Fav)
		if err != nil {
			log.Printf("error when scanning the customers row from starbucks_db")
			return []models.Customer{}, err
		}
		
		fmt.Println(customer)
		customers = append(customers, customer)
	}
	
	return customers, nil
}

func (c *customerRepository) Save(customer models.Customer) (models.Customer, error) {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Customer{}, err
	}
	
	defer db.Close()
	insertStat, err := db.Prepare("INSERT INTO customers (id, name, lastname, phone, address, fav) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Customer{}, err
	}
	
	customer.ID = uuid.New().String()
	
	_, err = insertStat.Exec(customer.ID, customer.Name, customer.LastName, customer.Phone, customer.Address, customer.Fav)
	if err != nil {
		log.Printf("error when saving the customer to starbucks_db, %s", err.Error())
		return models.Customer{}, err
	}
	
	return customer, nil
}

func (c *customerRepository) GetById(id string) (models.Customer, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Customer{}, err
	}
	
	defer db.Close()
	selectStat, err := db.Prepare("SELECT * FROM customers WHERE id=?")
	if err != nil {
		log.Printf("error when getting the customers from starbucks_db")
		return models.Customer{}, err
	}
	
	var customer models.Customer
	
	err = selectStat.QueryRow(id).Scan(&customer.ID, &customer.Name, &customer.LastName,
		&customer.Phone, &customer.Address, &customer.Fav)
	if err != nil {
		log.Printf("error when saving the customer to starbucks_db, %s", err.Error())
		return models.Customer{}, err
	}
	
	return customer, nil
}

func (c *customerRepository) Update(customer models.Customer) (models.Customer, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Customer{}, err
	}
	
	defer db.Close()
	updateStat, err := db.Prepare("UPDATE customers SET name=?, lastname=?, phone=?, address=?, fav=? WHERE id=?")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Customer{}, err
	}
	
	_, err = updateStat.Exec(customer.Name, customer.LastName, customer.Phone, customer.Address, customer.Fav, customer.ID)
	if err != nil {
		log.Printf("error when updating the customer from starbucks_db, %s", err.Error())
		return models.Customer{}, err
	}
	
	return customer, nil
}
