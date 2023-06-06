package model

import "fmt"

type Customer struct {
	id     int
	name   string
	tel    int
	addres string
}

func (cs *Customer) SetId(id int) {
	cs.id = id
}

func (cs *Customer) SetTel(tel int) {
	cs.tel = tel
}

func (cs *Customer) SetAddress(add string) {
	cs.addres = add
}

func (cs *Customer) GetId() int {
	return cs.id
}

func (cs *Customer) GetTel() int {
	return cs.tel
}

func (cs *Customer) GetAddress() string {
	return cs.addres
}

func (cs *Customer) Update(tel int, add string) {
	cs.tel = tel
	cs.addres = add
}

func (cs *Customer) DetilStringInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v", cs.id, cs.name, cs.tel, cs.addres)
}

func CreateCustomer(id int, name string, tel int, add string) *Customer {
	return &Customer{
		id, name, tel, add,
	}
}

func CreateCustomerWithoutId(name string, tel int, add string) *Customer {
	return &Customer{
		name:   name,
		tel:    tel,
		addres: add,
	}
}
