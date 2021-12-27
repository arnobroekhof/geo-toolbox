package geo_toolbox

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance_RotterdamAndAmsterdam(t *testing.T) {
	rotterdam := Point{51.92250, 4.47917}
	amsterdam := Point{52.37403, 4.88969}

	miles, km, err := Distance(rotterdam, amsterdam)

	// the distance is actual 57.566046470020034 kilometers, but we are rounding it to 58
	assert.Equal(t, 58, int(math.Round(km)))
	// the distance is actual 35.769871861122816 miles, but we are rounding to 36
	assert.Equal(t, 36, int(math.Round(miles)))
	assert.NoError(t, err)
}

func TestSmallDistance_Rotterdam(t *testing.T) {
	// 2 points near Van Brienenoordbrug
	east := Point{51.90455, 4.55427}
	west := Point{51.90453, 4.55390}

	miles, km, err := Distance(east, west)
	expectedKm := 0.02552 // the distance according to Google Maps is 0.02552 km. Using Haversine the distance is 0.02546 km.
	expectedMiles := 0.01585739283
	allowedDeviation := 0.005 // 0.5 meters or appr. 16.4 foot
	println(km)
	println(miles)

	assert.NoError(t, err)
	assert.True(t,
		km >= expectedKm-allowedDeviation &&
			km <= expectedKm+allowedDeviation)
	assert.True(t,
		miles >= expectedMiles-(allowedDeviation/1.609) &&
			miles <= expectedMiles+(allowedDeviation/1.609))
}

func TestDistance_Rotterdam(t *testing.T) {
	// From Northsea port entry to van Brienenoordbrug
	northsea := Point{51.99097, 4.04375}
	bridge := Point{51.90455, 4.55427}

	miles, km, err := Distance(northsea, bridge)
	expectedKm := 36.29 // the distance according to Google Maps is 36.29 km. Using Haversine the distance is 36.25105 km.
	expectedMiles := 22.549561
	allowedDeviation := 0.11 // 110 meters or appr. 360 foot
	println(km)

	assert.NoError(t, err)
	assert.True(t,
		km >= expectedKm-allowedDeviation &&
			km <= expectedKm+allowedDeviation)
	assert.True(t,
		miles >= expectedMiles-(allowedDeviation/1.609) &&
			miles <= expectedMiles+(allowedDeviation/1.609))
}
