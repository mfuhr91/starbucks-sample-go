package services

import (
	"log"
	"starbucks-app/models"
	"starbucks-app/repositories"
)

type productService struct{}

type ProductService interface {
	GetAll(bool) ([]models.Product, error)
	GetById(id string) (models.Product, error)
	Save(product models.Product) (models.Product, error)
	Delete(id string) error
}

var productRepository repositories.ProductRepository

func NewProductService(repository repositories.ProductRepository) ProductService {
	productRepository = repository
	return &productService{}
}

func (c productService) GetAll(allowEmptyProducts bool) ([]models.Product, error) {
	products, err := productRepository.GetAll(allowEmptyProducts)
	
	if err != nil {
		log.Printf("error when getting the products from service")
		return products, err
	}
	
	return products, nil
}

func (c productService) Save(product models.Product) (models.Product, error) {
	var err error
	if product.ID == "" {
		product, err = productRepository.Save(product)
	} else {
		product, err = productRepository.Update(product)
	}
	
	if err != nil {
		return models.Product{}, err
	}
	
	return product, nil
}

func (c *productService) GetById(id string) (models.Product, error) {
	product, err := productRepository.GetById(id)
	
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (c productService) Delete(id string) error {
	err := productRepository.Delete(id)
	if err != nil {
		return err
	}
	
	return nil
}
