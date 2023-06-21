package factory

import "fmt"

// MercedesSedan represents a Mercedes sedan car
type MercedesSedan struct{}

// Drive drives the Mercedes sedan
func (s MercedesSedan) Drive() {
	fmt.Println("Driving Mercedes sedan")
}

// MercedesSUV represents a Mercedes SUV car
type MercedesSUV struct{}

// Drive drives the Mercedes SUV
func (s MercedesSUV) Drive() {
	fmt.Println("Driving Mercedes SUV")
}
