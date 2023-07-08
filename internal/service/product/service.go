package product

import (
	"errors"
	"strconv"
	"strings"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if len(allProducts)-1 >= idx {
		return &allProducts[idx], nil
	} else {
		return nil, errors.New("no such index, list capacity: " + strconv.Itoa(len(allProducts)) + "start from 0")
	}
}

func (s *Service) New(name string) ([]Product, error) {
	if strings.TrimSpace(name) != "" {
		allProducts = append(allProducts, Product{name})
		return allProducts, nil
	}
	return nil, errors.New("value should not be an empty string or string, that conducts only spaces")
}
