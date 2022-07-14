package main

import (
	"fmt"
)

type dog struct {
	name   string
	weight float64
}
type cat struct {
	name   string
	weight float64
}
type cow struct {
	name   string
	weight float64
}

type Animal interface {
	fmt.Stringer
	getAnimalName() string
	getAnimalWeight() float64
	getAnimalFeedByMonth() float64
}

func (d dog) String() string {
	return "dog"
}
func (c cat) String() string {
	return "cat"
}
func (c cow) String() string {
	return "cow"
}

func (d dog) getAnimalName() string {
	return d.name
}
func (c cat) getAnimalName() string {
	return c.name
}
func (c cow) getAnimalName() string {
	return c.name
}

func (d dog) getAnimalWeight() float64 {
	return d.weight
}
func (c cat) getAnimalWeight() float64 {
	return c.weight
}
func (c cow) getAnimalWeight() float64 {
	return c.weight
}

func (d dog) getAnimalFeedByMonth() float64 {
	return d.weight / 5 * 10
}
func (c cat) getAnimalFeedByMonth() float64 {
	return c.weight * 7
}
func (c cow) getAnimalFeedByMonth() float64 {
	return c.weight * 25
}

func farmFeed(animals []Animal) float64 {
	s := 0.0
	for _, a := range animals {
		fmt.Printf("Food consumption by %v %q weighing %v kg in total is %v kg \n", a, a.getAnimalName(), a.getAnimalWeight(), a.getAnimalFeedByMonth())
		s += a.getAnimalFeedByMonth()
	}
	return s
}

func main() {

	animals := []Animal{
		dog{name: "Rex", weight: 5},
		dog{name: "Been", weight: 7},
		cat{name: "Iris", weight: 4},
		cow{name: "Mira", weight: 80},
		cow{name: "Sira", weight: 120},
		cow{name: "Bila", weight: 100},
	}

	sum := farmFeed(animals)
	fmt.Printf("All animals consume together %v kg of feed\n", sum)

}
