package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Macbook Pro M1",
			stock: 20,
		},
	}
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "iMac",
			stock: 10,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "laptop":
		return NewLaptop(), nil
	case "desktop":
		return NewDesktop(), nil
	default:
		return nil, fmt.Errorf("invalid computer type")
	}
}

func PrintNameAndStock(p IProduct) {
	fmt.Printf("Product Name: %s\nStock: %d\n\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
