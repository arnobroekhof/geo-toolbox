package geo_toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewJourney(t *testing.T) {
	rotterdam := Point{51.92250, 4.47917}
	amsterdam := Point{52.37403, 4.88969}
	utrecht := Point{52.0908, 5.1222}
	tn := time.Now()
	t1 := tn.Add(-3 * time.Hour)
	t2 := tn.Add(-2 * time.Hour)
	t3 := tn.Add(-1 * time.Hour)

	tpRotterdam := NewTimebasedPoint(t1, rotterdam)
	tpAmsterdam := NewTimebasedPoint(t2, amsterdam)
	tpUtrecht := NewTimebasedPoint(t3, utrecht)

	journey := NewJourney(tpRotterdam, tpAmsterdam, tpUtrecht)
	assert.Equal(t, float64(93), journey.TotalDistance().KM())
	assert.Equal(t, float64(46), journey.AverageSpeed().KM())
	assert.Equal(t, float64(7200), journey.TotalTime())

}
