package geo_toolbox

type Point struct {
	Latitude  float64
	Longitude float64
}

func (p Point) Lat() float64 {
	return p.Latitude
}

func (p Point) Lon() float64 {
	return p.Longitude
}

func (p Point) X() float64 {
	return p.Longitude
}

func (p Point) Y() float64 {
	return p.Latitude
}

// Coordinates return a float64 array of [ Lon, Lat ]
func (p Point) Coordinates() []float64 {
	return []float64{
		p.Lat(),
		p.Lon(),
	}
}
