package repositories

import (
	"fmt"
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
	Delete(id string) error
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
		log.Printf("error when getting the customers from db, %s", err)
		return []models.Customer{}, err
	}
	
	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		
		err = rows.Scan(&customer.ID, &customer.Name, &customer.LastName,
			&customer.Phone, &customer.Address, &customer.Fav)
		if err != nil {
			log.Printf("error when scanning the customers row from db")
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
	insertStat, err := db.Prepare("INSERT INTO customers (id, name, lastname, phone, address, fav) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Customer{}, err
	}
	
	customer.ID = uuid.New().String()
	
	_, err = insertStat.Exec(customer.ID, customer.Name, customer.LastName, customer.Phone, customer.Address, customer.Fav)
	if err != nil {
		log.Printf("error when saving the customer to db, %s", err.Error())
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
	selectStat, err := db.Prepare("SELECT * FROM customers WHERE id=$1")
	if err != nil {
		log.Printf("error when getting the customers from db")
		return models.Customer{}, err
	}
	
	var customer models.Customer
	
	err = selectStat.QueryRow(id).Scan(&customer.ID, &customer.Name, &customer.LastName,
		&customer.Phone, &customer.Address, &customer.Fav)
	if err != nil {
		log.Printf("error when saving the customer to db, %s", err.Error())
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
	updateStat, err := db.Prepare("UPDATE customers SET name=$1, lastname=$2, phone=$3, address=$4, fav=$5 WHERE id=$6")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Customer{}, err
	}
	
	_, err = updateStat.Exec(customer.Name, customer.LastName, customer.Phone, customer.Address, customer.Fav, customer.ID)
	if err != nil {
		log.Printf("error when updating the customer from db, %s", err.Error())
		return models.Customer{}, err
	}
	
	return customer, nil
}

func (c *customerRepository) Delete(id string) error {
	db, err := database.DatabaseConnect()
	if err != nil {
		return err
	}
	
	defer db.Close()
	deleteStat, err := db.Prepare("DELETE FROM customers WHERE id=$1")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return err
	}
	
	_, err = deleteStat.Exec(id)
	if err != nil {
		log.Printf("error when deleting the customer from db, %s", err.Error())
		return err
	}
	
	return nil
}
