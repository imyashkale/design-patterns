package observer

import "fmt"

type item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) UpdateAvailability(){
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) Deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer)[]Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
            observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
		   return observerList[:observerListLength-1]
		}
	}
	return observerList
}