package main

import (
	"fmt"
)

type animal struct {
	name     string
	species  string
	weight   int
	food     int
	onweight int
}

type farm struct {
	animals []animal
}

type Animal interface {
	getAnimalName() string
	getAnimalSpecies() string
	getAnimalWeight() int
	getAnimalEatsByKg() float64
}

type Farm interface {
	getAnimals() []animal
	getAnimalsCount() int
}

func (f farm) getAnimals() []animal {
	return f.animals
}
func (f farm) getAnimalsCount() int {
	var countAnimals int
	for i, _ := range f.animals {
		countAnimals += i
	}
	return countAnimals
}

func (a animal) getAnimalName() string {
	return a.name
}
func (a animal) getAnimalSpecies() string {
	return a.species
}
func (a animal) getAnimalWeight() int {
	return a.weight
}
func (a animal) getAnimalEatsByKg() float64 {
	if aw := a.onweight; aw <= 0 {
		a.onweight = 1
	}
	return float64(a.food / a.onweight)
}

func farmFeed(f Farm) float64 {
	var i Animal
	var s float64
	animals := f.getAnimals()

	for _, a := range animals {
		i = &a
		feedByAnimal := i.getAnimalEatsByKg() * float64(i.getAnimalWeight())
		fmt.Printf("Food consumption by %v %q weighing %v kg in total is %v kg \n", i.getAnimalSpecies(), i.getAnimalName(), i.getAnimalWeight(), feedByAnimal)
		s += feedByAnimal
	}
	return s
}

func main() {
	var f Farm

	animals := []animal{
		{name: "Rex", species: "dog", weight: 5, food: 10, onweight: 5},
		{name: "Been", species: "dog", weight: 7, food: 10, onweight: 5},
		{name: "Isis", species: "cat", weight: 4, food: 7, onweight: 1},
		{name: "Mira", species: "cow", weight: 80, food: 25, onweight: 1},
		{name: "Sira", species: "cow", weight: 120, food: 25, onweight: 1},
		{name: "Bila", species: "cow", weight: 100, food: 25, onweight: 1},
	}

	myfarm := farm{animals: animals}
	f = &myfarm

	sum := farmFeed(f)
	fmt.Printf("All animals (%v) consume together %v kg of feed\n", myfarm.getAnimalsCount(), sum)

}
