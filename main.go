package main

import (
	"fmt"
	"gotest/utils"

	"github.com/stretchr/testify/mock"
)

func main() {

	c := CustomerRepositoryMock{}
	//"Bond", 18, nil => index of Bond = 0 , index of 18 = 1 , index of nil = 2
	c.On("GetCustomer", 1).Return("Bond", 18, nil)
	c.On("GetCustomer", 2).Return("", 0, utils.ErrNotFound)

	//case GetCustomer = 1
	//run
	//go run .
	// name, age, err := c.GetCustomer(1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(name, age)

	//case GetCustomer = 2
	//run
	//go run .
	name, age, err := c.GetCustomer(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)

}

type CustomerRepository interface {
	GetCustomer(id int) (name string, age int, err error)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}
