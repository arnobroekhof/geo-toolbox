package geo_toolbox

import (
	"math"
)

type TotalDistance float64
type AverageSpeed float64

// KM returns the total distance rounded on 1 point after the decimal
func (t TotalDistance) KM() float64 {
	return math.Round((float64(t) * 10) / 10)
}

// KM returns the total distance rounded on 2 points after the decimal
func (t AverageSpeed) KM() float64 {
	return math.Round((float64(t) * 100) / 100)
}

type Journey struct {
	points []TimebasedPoint
}

func NewJourney(points ...TimebasedPoint) Journey {
	return Journey{points}
}

// AverageSpeed return the average speed in Kilometers per hour
func (j Journey) AverageSpeed() (averageSpeed AverageSpeed) {
	speeds := SpeedMultiple(j.points...)

	sum := float64(0)
	n := len(speeds)

	for i := 0; i < n; i++ {
		sum += speeds[i]
	}

	return AverageSpeed(sum / float64(n))
}

// TotalDistance return the total distance in kilometers
func (j Journey) TotalDistance() (totalDistance TotalDistance) {
	l := len(j.points)
	for key, point := range j.points {
		switch {
		case key > 0:
			_, km, _ := Distance(j.points[(key-1)].Point, point.Point)
			totalDistance += TotalDistance(km)
		case (key + 1) == l:
			_, km, _ := Distance(j.points[(key-1)].Point, point.Point)
			totalDistance += TotalDistance(km)
		}
	}
	return
}

// TotalTime calculates to total time off the journey
func (j Journey) TotalTime() (seconds float64) {
	l := len(j.points)
	for key, point := range j.points {
		switch {
		case key > 0:
			seconds += point.Time.Sub(j.points[key-1].Time).Seconds()
		case (key + 1) == l:
			seconds += point.Time.Sub(j.points[key-1].Time).Seconds()
		}
	}
	return
}
