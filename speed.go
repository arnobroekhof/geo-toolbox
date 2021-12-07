package geo_toolbox

import (
	"time"
)

type Kph float64

type TimebasedPoint struct {
	Time  time.Time
	Point Point
}

func NewTimebasedPoint(time time.Time, point Point) TimebasedPoint {
	return TimebasedPoint{
		time,
		point,
	}
}

//Speed takes TimebasePoint and converts to meters per second and kilometers per hour
func Speed(p1, p2 TimebasedPoint) (mps, kph float64, err error) {
	_, k, err := Distance(p1.Point, p2.Point)
	if err != nil {
		return -1, -1, err
	}
	// convert km to meters
	meters := k * 1000.0
	t := p2.Time.Sub(p1.Time)
	mps = meters / t.Seconds()
	kph = (mps * 3600.0) / 1000.0
	return
}

func SpeedMultiple(points ...TimebasedPoint) (kph []float64) {
	l := len(points)

	for key, point := range points {
		switch {
		case key > 0:
			_, k, _ := Speed(points[(key-1)], point)
			kph = append(kph, k)
		case (key + 1) == l:
			_, k, _ := Speed(points[(key-1)], point)
			kph = append(kph, k)
		}
	}
	return
}
