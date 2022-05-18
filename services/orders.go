package services

import (
	"log"
	"starbucks-app/models"
	"starbucks-app/repositories"
	"strings"
)

type orderService struct{}

type OrderService interface {
	GetAll() ([]models.Order, error)
	GetById(id string) (models.Order, error)
	Save(order models.Order) error
	Delete(id string) error
}

var orderRepository repositories.OrderRepository

func NewOrderService(repository repositories.OrderRepository) OrderService {
	orderRepository = repository
	return &orderService{}
}

func (c orderService) GetAll() ([]models.Order, error) {
	orders, err := orderRepository.GetAll()
	
	for i, order := range orders {
		customer, err := customerRepository.GetById(order.Customer.ID)
		if err != nil {
			log.Printf("error when getting the orders from service")
			return orders, err
		}
		date := orders[i].Time[0:16]
		date = strings.Replace(date, "T", " ", 1)
		
		orders[i].Time = date
		orders[i].Customer = customer
		
	}
	
	if err != nil {
		log.Printf("error when getting the orders from service")
		return orders, err
	}
	
	return orders, nil
}

func (c orderService) Save(order models.Order) error {
	var err error
	if order.ID == "" {
		err = orderRepository.Save(order)
	} else {
		order, err = orderRepository.Update(order)
	}
	
	if err != nil {
		return err
	}
	
	return nil
}

func (c *orderService) GetById(id string) (models.Order, error) {
	order, err := orderRepository.GetById(id)
	
	customer, err := customerRepository.GetById(order.Customer.ID)
	order.Customer = customer
	
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (c orderService) Delete(id string) error {
	err := orderRepository.Delete(id)
	if err != nil {
		return err
	}
	
	return nil
}
