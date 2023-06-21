package factory

// MercedesFactory is the concrete factory for creating Mercedes cars
type MercedesFactory struct{}

// BuildSedan creates a Mercedes sedan car
func (f MercedesFactory) BuildSedan() Sedan {
	return &MercedesSedan{}
}

// BuildSUV creates a Mercedes SUV car
func (f MercedesFactory) BuildSUV() SUV {
	return &MercedesSUV{}
}
