package factory

import "fmt"

// Sedan represents a sedan car
type Sedan interface {
	Drive()
}

// ToyotaSedan represents a Toyota sedan car
type ToyotaSedan struct{}

// Drive drives the Toyota sedan
func (s *ToyotaSedan) Drive() {
	fmt.Println("Driving Toyota sedan")
}

// SUV represents an SUV car
type SUV interface {
	Drive()
}

// ToyotaSUV represents a Toyota SUV car
type ToyotaSUV struct{}

// Drive drives the Toyota SUV
func (s *ToyotaSUV) Drive() {
	fmt.Println("Driving Toyota SUV")
}
