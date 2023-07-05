package product

import (
	"errors"
	"strconv"
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
