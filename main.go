package main

import (
	"fmt"
)

type Dog struct {
	name     string
	weight   int
	food     int
	onweight int
}
type Cat struct {
	name     string
	weight   int
	food     int
	onweight int
}
type Cow struct {
	name     string
	weight   int
	food     int
	onweight int
}
type Farm struct {
	Dog
	Cat
	Cow
	dogCount int
	catCount int
	cowCount int
}

type Animal interface {
	getDogName() string
	getCatName() string
	getCowName() string
	getDogWeight() int
	getCatWeight() int
	getCowWeight() int
	getDogEatsByKg() float64
	getCatEatsByKg() float64
	getCowEatsByKg() float64
}

type iFarm interface {
	Animal
	getSumFeedByMonth() float64
	getDogCount() int
	getCatCount() int
	getCowCount() int
}

func (f *Farm) getSumFeedByMonth() float64 {

	eatByOneDog := f.getDogEatsByKg() * float64(f.getDogWeight())
	eatByOneCat := f.getCatEatsByKg() * float64(f.getCatWeight())
	eatByOneCow := f.getCowEatsByKg() * float64(f.getCowWeight())
	fmt.Printf("Food consumption by %q weighing %v kg in total is %v kg \n", f.getDogName(), f.getDogWeight(), eatByOneDog)
	fmt.Printf("Food consumption by %q weighing %v kg in total is %v kg \n", f.getCatName(), f.getCatWeight(), eatByOneCat)
	fmt.Printf("Food consumption by %q weighing %v kg in total is %v kg \n", f.getCowName(), f.getCowWeight(), eatByOneCow)

	eatByAllDog := eatByOneDog * float64(f.getDogCount())
	eatByAllCat := eatByOneCat * float64(f.getCatCount())
	eatByAllCow := eatByOneCow * float64(f.getCowCount())

	return eatByAllDog + eatByAllCat + eatByAllCow
}

func (d Dog) getDogName() string {
	return d.name
}
func (c Cat) getCatName() string {
	return c.name
}
func (c Cow) getCowName() string {
	return c.name
}

func (d Dog) getDogWeight() int {
	return d.weight
}
func (c Cat) getCatWeight() int {
	return c.weight
}
func (c Cow) getCowWeight() int {
	return c.weight
}

func (d Dog) getDogEatsByKg() float64 {
	if dw := d.onweight; dw <= 0 {
		d.onweight = 1
	}
	return float64(d.food / d.onweight)
}
func (c Cat) getCatEatsByKg() float64 {
	if cw := c.onweight; cw <= 0 {
		c.onweight = 1
	}
	return float64(c.food / c.onweight)
}
func (c Cow) getCowEatsByKg() float64 {
	if cw := c.onweight; cw <= 0 {
		c.onweight = 1
	}
	return float64(c.food / c.onweight)
}

func (f Farm) getDogCount() int {
	return f.dogCount
}
func (f Farm) getCatCount() int {
	return f.catCount
}
func (f Farm) getCowCount() int {
	return f.cowCount
}

func main() {

	myDogs := Dog{
		name:     "dog",
		weight:   15,
		food:     10,
		onweight: 5,
	}

	myCats := Cat{
		name:     "cat",
		weight:   8,
		food:     7,
		onweight: 1,
	}

	myCows := Cow{
		name:     "cow",
		weight:   100,
		food:     25,
		onweight: 1,
	}

	var f iFarm

	f = &Farm{
		Dog:      myDogs,
		Cat:      myCats,
		Cow:      myCows,
		dogCount: 74,
		catCount: 52,
		cowCount: 6,
	}

	sum := f.getSumFeedByMonth()

	fmt.Printf("All animals (%v dogs, %v cats, %v cows) need a total of %v kg of feed per month\n", f.getDogCount(), f.getCatCount(), f.getCowCount(), sum)

}
