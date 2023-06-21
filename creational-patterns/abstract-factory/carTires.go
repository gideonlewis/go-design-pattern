package main

import "fmt"

type ITires interface {
	getBrand() string
	getSize() int64
	getWeight() int64
}

type Tires struct {
	Brand  string
	Size   int64
	Weight int64
}

// Concrete product
type MercedesTires struct {
	Tires
}

func (mt *MercedesTires) getBrand() string {
	return mt.Brand
}

func (mt *MercedesTires) getSize() int64 {
	return mt.Size
}

func (mt *MercedesTires) getWeight() int64 {
	return mt.Weight
}

type ToyotaTires struct {
	Tires
}

func (mt *ToyotaTires) getBrand() string {
	return mt.Brand
}

func (mt *ToyotaTires) getSize() int64 {
	return mt.Size
}

func (mt *ToyotaTires) getWeight() int64 {
	return mt.Weight
}

// Abstract Factory interface
type ITiresFactory interface {
	makeTires(size, weight int64) ITires
}

// Concrete Factory
type Mercedes struct{}

func (m *Mercedes) makeTires(size, weight int64) ITires {
	return &MercedesTires{
		Tires{
			Brand:  "Mercedes",
			Size:   size,
			Weight: weight,
		},
	}
}

type Toyota struct{}

func (m *Toyota) makeTires(size, weight int64) ITires {
	return &MercedesTires{
		Tires{
			Brand:  "Toyota",
			Size:   size,
			Weight: weight,
		},
	}
}

func main() {
	mercedesFactory := Mercedes{}
	mercedesTires := mercedesFactory.makeTires(60, 10)
	showDetails(mercedesTires)

	toyotaFactory := Toyota{}
	toyotaTires := toyotaFactory.makeTires(75, 15)
	showDetails(toyotaTires)
}

func showDetails(t ITires) {
	fmt.Println("Brand: ", t.getBrand())
	fmt.Println("Size: ", t.getSize())
	fmt.Println("Weight: ", t.getWeight())

}
