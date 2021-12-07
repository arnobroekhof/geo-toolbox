package geo_toolbox

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
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
