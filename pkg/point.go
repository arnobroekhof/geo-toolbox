package pkg

type Point struct {
	latitude  float64
	longitude float64
}

func (p Point) Lat() float64 {
	return p.latitude
}

func (p Point) Lon() float64 {
	return p.longitude
}

func (p Point) X() float64 {
	return p.longitude
}

func (p Point) Y() float64 {
	return p.latitude
}

// Coordinates return a float64 array of [ Lon, Lat ]
func (p Point) Coordinates() []float64 {
	return []float64{
		p.Lat(),
		p.Lon(),
	}
}
