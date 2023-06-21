package factory

// ToyotaFactory is the concrete factory for creating Toyota cars
type ToyotaFactory struct{}

// BuildSedan creates a Toyota sedan car
func (f ToyotaFactory) BuildSedan() Sedan {
	return &ToyotaSedan{}
}

// BuildSUV creates a Toyota SUV car
func (f ToyotaFactory) BuildSUV() SUV {
	return &ToyotaSUV{}
}
