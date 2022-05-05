package services

import (
	"log"
	"starbucks-app/models"
	"starbucks-app/repositories"
)

type customerService struct{}

type CustomerService interface {
	GetAll() ([]models.Customer, error)
	GetById(id string) (models.Customer, error)
	Save(customer models.Customer) (models.Customer, error)
	Delete(id string) error
}

var customerRepository repositories.CustomerRepository

func NewCustomerService(repository repositories.CustomerRepository) CustomerService {
	customerRepository = repository
	return &customerService{}
}

func (c customerService) GetAll() ([]models.Customer, error) {
	customers, err := customerRepository.GetAll()
	if err != nil {
		log.Printf("error when getting the customers from service")
		return customers, err
	}
	
	return customers, nil
}

func (c customerService) Save(customer models.Customer) (models.Customer, error) {
	var err error
	if customer.ID == "" {
		customer, err = customerRepository.Save(customer)
	} else {
		customer, err = customerRepository.Update(customer)
	}
	if err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

func (c *customerService) GetById(id string) (models.Customer, error) {
	customer, err := customerRepository.GetById(id)
	if err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

func (c customerService) Delete(id string) error {
	err := customerRepository.Delete(id)
	if err != nil {
		return err
	}
	
	return nil
}
