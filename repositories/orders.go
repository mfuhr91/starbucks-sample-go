package repositories

import (
	"github.com/google/uuid"
	"log"
	"starbucks-app/database"
	"starbucks-app/models"
)

type orderRepository struct {
}

type OrderRepository interface {
	GetAll() ([]models.Order, error)
	GetById(id string) (models.Order, error)
	Save(models.Order) error
	Update(models.Order) (models.Order, error)
	Delete(id string) error
}

var productRepo ProductRepository

func NewOrderRepository(repository ProductRepository) OrderRepository {
	productRepo = repository
	return &orderRepository{}
}

func (orderRepo *orderRepository) GetAll() ([]models.Order, error) {
	
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
		
		err = rows.Scan(&order.ID, &order.Number, &order.Customer.ID, &order.Time, &order.FinalPrice)
		if err != nil {
			log.Printf("error when scanning the orders row from db, %s", err)
			return []models.Order{}, err
		}
		
		var items []models.Item
		items, err = getItemsByOrderID(order.ID)
		if err != nil {
			return []models.Order{}, err
		}
		
		for i, item := range items {
			items[i].Product, _ = productRepo.GetById(item.ProductID)
		}
		
		order.Items = items
		
		orders = append(orders, order)
	}
	
	return orders, nil
}

func (orderRepo *orderRepository) Save(order models.Order) error {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return err
	}
	
	defer db.Close()
	
	order.ID = uuid.New().String()
	
	_, err = db.Query("INSERT INTO orders (id, time, customer_id, final_price) VALUES ($1, $2, $3, $4)",
		order.ID, order.Time, order.Customer.ID, order.FinalPrice)
	if err != nil {
		log.Printf("error when saving the order into db, %s", err.Error())
		return err
	}
	
	err = saveItems(order)
	if err != nil {
		return err
	}
	
	return nil
}

func (orderRepo *orderRepository) GetById(id string) (models.Order, error) {
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
	
	err = selectStat.QueryRow(id).Scan(&order.ID, &order.Number, &order.Customer.ID, &order.Time, &order.FinalPrice)
	if err != nil {
		log.Printf("error when saving the order into db, %s", err.Error())
		return models.Order{}, err
	}
	
	var items []models.Item
	items, err = getItemsByOrderID(order.ID)
	if err != nil {
		return order, err
	}
	
	for i, item := range items {
		items[i].Product, _ = productRepo.GetById(item.ProductID)
	}
	
	order.Items = items
	
	return order, nil
}

func (orderRepo *orderRepository) Update(order models.Order) (models.Order, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return models.Order{}, err
	}
	
	defer db.Close()
	_, err = db.Query("UPDATE orders SET final_price=$1, time=$2, customer_id=$3 WHERE id=$4",
		order.FinalPrice, order.Time, order.Customer.ID, order.ID)
	if err != nil {
		log.Printf("error when updating the order from db, %s", err.Error())
		return models.Order{}, err
	}
	
	err = deleteItemsByOrderId(order.ID)
	if err != nil {
		return order, err
	}
	
	err = saveItems(order)
	if err != nil {
		return order, err
	}
	
	return order, nil
}

func (orderRepo *orderRepository) Delete(id string) error {
	db, err := database.DatabaseConnect()
	if err != nil {
		return err
	}
	
	defer db.Close()
	
	err = deleteItemsByOrderId(id)
	if err != nil {
		return err
	}
	
	_, err = db.Query("DELETE FROM orders WHERE id=$1", id)
	if err != nil {
		log.Printf("error when deleting the order from db, %s", err.Error())
		return err
	}
	
	return nil
}

func saveItems(order models.Order) error {
	db, err := database.DatabaseConnect()
	if err != nil {
		return err
	}
	
	defer db.Close()
	
	for _, item := range order.Items {
		item.ID = uuid.New().String()
		item.OrderID = order.ID
		_, err = db.Query("INSERT INTO items (id, order_id, product_id, quantity, price)"+
			" VALUES ($1, $2, $3, $4, $5)",
			item.ID, item.OrderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			log.Printf("error when saving the items into db, %s", err.Error())
			return err
		}
		
		err = updateProduct(item, "subs")
		if err != nil {
			return err
		}
	}
	return nil
}

func updateProduct(item models.Item, operation string) error {
	product, err := productRepo.GetById(item.ProductID)
	if err != nil {
		return err
	}
	
	if operation == "add" {
		product.Quantity = product.Quantity + item.Quantity
	} else {
		product.Quantity = product.Quantity - item.Quantity
	}
	
	product, err = productRepo.Update(product)
	if err != nil {
		return err
	}
	
	return nil
}
func getItemsByOrderID(orderID string) ([]models.Item, error) {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return []models.Item{}, err
	}
	
	defer db.Close()
	rows, err := db.Query("SELECT * FROM items WHERE order_id=$1", orderID)
	if err != nil {
		log.Printf("error when getting the orders from db")
		return []models.Item{}, err
	}
	
	var items []models.Item
	
	for rows.Next() {
		var item models.Item
		
		err = rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			log.Printf("error when scanning the orders row from db, %s", err)
			return []models.Item{}, err
		}
		
		items = append(items, item)
	}
	
	return items, nil
}

func deleteItemsByOrderId(orderID string) error {
	
	db, err := database.DatabaseConnect()
	if err != nil {
		return err
	}
	
	defer db.Close()
	
	var items []models.Item
	items, err = getItemsByOrderID(orderID)
	if err != nil {
		return err
	}
	for _, item := range items {
		err = updateProduct(item, "add")
		if err != nil {
			return err
		}
		
		_, err = db.Query("DELETE FROM items WHERE id=$1", item.ID)
		if err != nil {
			log.Printf("error when deleting the items from db, %s", err)
			return err
		}
		
	}
	
	return nil
}
