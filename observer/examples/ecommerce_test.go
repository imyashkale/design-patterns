package examples

import (
	"testing"

	"github.com/imyashkale/go-design-patterns/observer"
)

func TestExampleShirt(t *testing.T){
	shirtItem := observer.NewItem("Jacket")

	observerFirst := observer.NewCustomer("imyashkale@gmail.com")
	observerSecond := observer.NewCustomer("yashkale46@gmail.com")

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()

}