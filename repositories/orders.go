package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"starbucks-app/database"
	"starbucks-app/models"
)

type orderRepository struct{}

type OrderRepository interface {
	GetAll() ([]models.Order, error)
	GetById(id string) (models.Order, error)
	Save(models.Order) (models.Order, error)
	Update(models.Order) (models.Order, error)
	Delete(id string) error
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (c *orderRepository) GetAll() ([]models.Order, error) {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return []models.Order{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Printf("error when getting the orders from db, %s", err)
		return []models.Order{}, err
	}
	
	var orders []models.Order
	for rows.Next() {
		var order models.Order
		
		err = rows.Scan(&order.ID, &order.Number, &order.Customer.ID,
			&order.Product, &order.Time)
		if err != nil {
			log.Printf("error when scanning the orders row from db, %s", err)
			return []models.Order{}, err
		}
		
		fmt.Println(order)
		orders = append(orders, order)
	}
	
	return orders, nil
}

func (c *orderRepository) Save(order models.Order) (models.Order, error) {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Order{}, err
	}
	
	defer db.Close()
	insertStat, err := db.Prepare("INSERT INTO orders (id, product, time, customer_id) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Order{}, err
	}
	
	order.ID = uuid.New().String()
	
	_, err = insertStat.Exec(order.ID, order.Product, order.Time, order.Customer.ID)
	if err != nil {
		log.Printf("error when saving the order to db, %s", err.Error())
		return models.Order{}, err
	}
	
	return order, nil
}

func (c *orderRepository) GetById(id string) (models.Order, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Order{}, err
	}
	
	defer db.Close()
	selectStat, err := db.Prepare("SELECT * FROM orders WHERE id=$1")
	if err != nil {
		log.Printf("error when getting the orders from db")
		return models.Order{}, err
	}
	
	var order models.Order
	
	err = selectStat.QueryRow(id).Scan(&order.ID, &order.Number, &order.Customer.ID,
		&order.Product, &order.Time)
	if err != nil {
		log.Printf("error when saving the order to db, %s", err.Error())
		return models.Order{}, err
	}
	
	return order, nil
}

func (c *orderRepository) Update(order models.Order) (models.Order, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Order{}, err
	}
	
	defer db.Close()
	updateStat, err := db.Prepare("UPDATE orders SET product=$1, time=$2, customer_id=$3 WHERE id=$4")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return models.Order{}, err
	}
	
	_, err = updateStat.Exec(order.Product, order.Time, order.Customer.ID, order.ID)
	if err != nil {
		log.Printf("error when updating the order from db, %s", err.Error())
		return models.Order{}, err
	}
	
	return order, nil
}

func (c *orderRepository) Delete(id string) error {
	db, err := database.DatabaseConnect()
	if err != nil {
		return err
	}
	
	defer db.Close()
	updateStat, err := db.Prepare("DELETE FROM orders WHERE id=$1")
	if err != nil {
		log.Printf("error when preparing the query, %s", err.Error())
		return err
	}
	
	_, err = updateStat.Exec(id)
	if err != nil {
		log.Printf("error when deleting the order from db, %s", err.Error())
		return err
	}
	
	return nil
}
