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
	expectedKm := 0.025562 // source: https://www.movable-type.co.uk/scripts/latlong-vincenty.html
	// Google Maps distance is 0.02552 km. Haversine distance is 0.02546 km.
	expectedMiles := 0.015883490416
	allowedDeviation := 0.001 // 10cm to allow for rounding differences

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
	expectedKm := 36.395404 // source: https://www.movable-type.co.uk/scripts/latlong-vincenty.html
	// Google Maps distance is 36.29 km. Haversine distance is 36.25105 km.
	expectedMiles := 22.6150555754
	allowedDeviation := 0.001 // 10cm to allow for rounding differences
	println(km)

	assert.NoError(t, err)
	assert.True(t,
		km >= expectedKm-allowedDeviation &&
			km <= expectedKm+allowedDeviation)
	assert.True(t,
		miles >= expectedMiles-(allowedDeviation/1.609) &&
			miles <= expectedMiles+(allowedDeviation/1.609))
}
