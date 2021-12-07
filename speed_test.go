package geo_toolbox

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestSpeed(t *testing.T) {
	rotterdam := Point{51.92250, 4.47917}
	tn := time.Now()
	t1 := tn.Add(-2 * time.Hour)
	t2 := tn.Add(-1 * time.Hour)
	amsterdam := Point{52.37403, 4.88969}

	tpRotterdam := NewTimebasedPoint(t1, rotterdam)
	tpAmsterdam := NewTimebasedPoint(t2, amsterdam)

	mps, kph, err := Speed(tpRotterdam, tpAmsterdam)
	// because the distance is 58 kilometers and the time is 1 hour we expect
	// that the speed is 58 kilometers per hour which is 16 miles per second
	assert.Equal(t, 58, int(math.Round(kph)))
	assert.Equal(t, 16, int(math.Round(mps)))
	assert.NoError(t, err)

}

func TestSpeedMultiple_2Records(t *testing.T) {
	rotterdam := Point{51.92250, 4.47917}
	tn := time.Now()
	t1 := tn.Add(-2 * time.Hour)
	t2 := tn.Add(-1 * time.Hour)
	amsterdam := Point{52.37403, 4.88969}

	tpRotterdam := NewTimebasedPoint(t1, rotterdam)
	tpAmsterdam := NewTimebasedPoint(t2, amsterdam)

	kph := SpeedMultiple(tpRotterdam, tpAmsterdam)
	assert.Equal(t, 1, len(kph))
	assert.Equal(t, 58, int(math.Round(kph[0])))

}

func TestSpeedMultiple_3Records(t *testing.T) {
	rotterdam := Point{51.92250, 4.47917}
	tn := time.Now()
	t1 := tn.Add(-3 * time.Hour)
	t2 := tn.Add(-2 * time.Hour)
	t3 := tn.Add(-1 * time.Hour)
	amsterdam := Point{52.37403, 4.88969}
	utrecht := Point{52.0908, 5.1222}

	tpRotterdam := NewTimebasedPoint(t1, rotterdam)
	tpAmsterdam := NewTimebasedPoint(t2, amsterdam)
	tpUtrecht := NewTimebasedPoint(t3, utrecht)

	kph := SpeedMultiple(tpRotterdam, tpAmsterdam, tpUtrecht)
	assert.Equal(t, 2, len(kph))
	assert.Equal(t, 58, int(math.Round(kph[0])))
	assert.Equal(t, 35, int(math.Round(kph[1])))
}
