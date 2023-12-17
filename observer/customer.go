package observer

import "fmt"

type customer struct {
	id string
}

func NewCustomer(id string) Observer {
	return &customer{}
}

func (c *customer) Update(itemName string) {
	fmt.Printf("Sending email to customere %s for item %s \n", c.id, itemName)
}

func (c *customer) GetID()string {
	return c.id
}