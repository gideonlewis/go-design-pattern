package main

import "github.com/quaan2hand/go-design-pattern/creational-patterns/builder/factory"

func main() {
	// Create a Toyota factory
	toyotaFactory := factory.ToyotaFactory{}

	// Build and drive a Toyota sedan
	toyotaSedan := toyotaFactory.BuildSedan()
	toyotaSedan.Drive()

	// Build and drive a Toyota SUV
	toyotaSUV := toyotaFactory.BuildSUV()
	toyotaSUV.Drive()

	// Create a Mercedes factory
	mercedesFactory := factory.MercedesFactory{}

	// Build and drive a Mercedes sedan
	mercedesSedan := mercedesFactory.BuildSedan()
	mercedesSedan.Drive()

	// Build and drive a Mercedes SUV
	mercedesSUV := mercedesFactory.BuildSUV()
	mercedesSUV.Drive()
}
