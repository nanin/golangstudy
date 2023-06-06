package service

import (
	"test/crm/model"
)

type CustomerService struct {
	customers     []model.Customer
	custoemrIndex int
}

func CreateCustomerService() *CustomerService {
	return &CustomerService{
		customers:     make([]model.Customer, 0),
		custoemrIndex: 0,
	}
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) Add(custoemr model.Customer) bool {
	cs.custoemrIndex++
	custoemr.SetId(cs.custoemrIndex)
	cs.customers = append(cs.customers, custoemr)
	return true
}

func (cs *CustomerService) GetIndex(id int) (int, model.Customer) {
	result := -1
	cus := model.Customer{}
	for index, va := range cs.customers {
		if va.GetId() == id {
			result = index
			cus = va
			break
		}
	}
	return result, cus
}

func (cs *CustomerService) Delete(index int) bool {
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

func (cs *CustomerService) UpdateName(index int, customer model.Customer) bool {
	cs.customers[index] = customer

	return true
}
