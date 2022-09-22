package main

import (
	"errors"
	"fmt"
)

type Subject interface {
	Subscribe(o Observer) (bool, error)
	Unsubscribe(o Observer) (bool, error)
	Notify() (bool, error)
}

// Observer Interface
type Observer interface {
	Update(string)
}

type UserObserver struct {
	name string
}

func (s *UserObserver) Update(t string) {
	println(s.name, "has been updated", t)
}

type Product struct {
	nameOfProduct []string
	observers     []Observer
}

func (s *Product) AddProduct(item string) []string {
	s.nameOfProduct = append(s.nameOfProduct, item)
	return s.nameOfProduct
}

func (s *Product) Subscribe(o Observer) (bool, error) {
	for _, observer := range s.observers {
		if observer == o {
			return false, errors.New("Observer already exists")
		}
	}
	s.observers = append(s.observers, o)
	return true, nil
}

func (s *Product) Unsubscribe(o Observer) (bool, error) {

	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("not found")
}

func (s *Product) Notify() {
	for _, observer := range s.observers {
		fmt.Println(observer)
		for _, ticker := range s.nameOfProduct {
			fmt.Println(ticker)
		}
	}
}

func main() {

	Product1 := &Product{
		nameOfProduct: nil,
	}

	Product1.AddProduct("Apple")
	Product1.AddProduct("Banana")
	Product1.AddProduct("Orange")

	observerA := &UserObserver{
		name: "Dinara",
	}
	observerB := &UserObserver{
		name: "Gaziz",
	}

	Product1.Subscribe(observerA)
	Product1.Subscribe(observerB)

	Product1.Notify()
	Product1.Unsubscribe(observerA)
	Product1.Notify()

}
