package geo_toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineString_NoIntersectDoublePoint(t *testing.T) {
	l1p1 := Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1 := LineString{[]Point{
		l1p1, l1p2,
	}}

	l2p1 := Point{
		Latitude:  51.83040734742036,
		Longitude: 4.458622634410858,
	}

	l2p2 := Point{
		Latitude:  51.83043380979297,
		Longitude: 4.459003508090973,
	}

	l2 := LineString{[]Point{
		l2p1, l2p2,
	}}

	crossed, _ := l1.Intersect(l2)
	assert.False(t, crossed)

}

func TestLineString_IntersectDoublePoint(t *testing.T) {
	l1p1 := Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1 := LineString{[]Point{
		l1p1, l1p2,
	}}

	l2p1 := Point{
		Latitude:  51.93040734742036,
		Longitude: 4.458622634410858,
	}

	l2p2 := Point{
		Latitude:  51.93043380979297,
		Longitude: 4.459003508090973,
	}

	l2 := LineString{[]Point{
		l2p1, l2p2,
	}}

	crossed, point := l1.Intersect(l2)
	assert.True(t, crossed)
	assert.Equal(t, Point{51.93041738656155, 4.458767128471792}, point)

	t.Logf("crossed: %v and point is [ %v , %v ] \n", crossed, point.Longitude, point.Latitude)

}

func TestLineString_IntersectTriplePoint(t *testing.T) {
	l1p1 := Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1p3 := Point{
		51.930346153123985,
		4.45824900196975,
	}

	l1 := LineString{[]Point{
		l1p1, l1p2, l1p3,
	}}

	l2p1 := Point{
		Latitude:  51.93040734742036,
		Longitude: 4.458622634410858,
	}

	l2p2 := Point{
		Latitude:  51.93043380979297,
		Longitude: 4.459003508090973,
	}

	l2 := LineString{[]Point{
		l2p1, l2p2,
	}}

	crossed, point := l1.Intersect(l2)
	assert.True(t, crossed)
	assert.Equal(t, Point{51.93041738656155, 4.458767128471792}, point)

	t.Logf("crossed: %v and point is [ %v , %v ] \n", crossed, point.Longitude, point.Latitude)

}

func TestLineString_IntersectTriplePointSwapped(t *testing.T) {
	l1p1 := Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1p3 := Point{
		51.930346153123985,
		4.45824900196975,
	}

	l1 := LineString{[]Point{
		l1p1, l1p2, l1p3,
	}}

	l2p1 := Point{
		Latitude:  51.93040734742036,
		Longitude: 4.458622634410858,
	}

	l2p2 := Point{
		Latitude:  51.93043380979297,
		Longitude: 4.459003508090973,
	}

	l2 := LineString{[]Point{
		l2p1, l2p2,
	}}

	crossed, point := l2.Intersect(l1)
	assert.True(t, crossed)
	assert.Equal(t, Point{51.93041738656155, 4.458767128471792}, point)

	t.Logf("crossed: %v and point is [ %v , %v ] \n", crossed, point.Longitude, point.Latitude)

}

func TestLineString_IntersectTriplePointReverse(t *testing.T) {
	l1p3 := Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1p1 := Point{
		51.930346153123985,
		4.45824900196975,
	}

	l1 := LineString{[]Point{
		l1p1, l1p2, l1p3,
	}}

	l2p1 := Point{
		Latitude:  51.93040734742036,
		Longitude: 4.458622634410858,
	}

	l2p2 := Point{
		Latitude:  51.93043380979297,
		Longitude: 4.459003508090973,
	}

	l2 := LineString{[]Point{
		l2p1, l2p2,
	}}

	crossed, point := l1.Intersect(l2)
	assert.True(t, crossed)
	assert.Equal(t, Point{51.93041738656155, 4.458767128471792}, point)

	t.Logf("crossed: %v and point is [ %v , %v ] \n", crossed, point.Longitude, point.Latitude)

}

func TestLineString_IntersectTriplePointReverseSwapped(t *testing.T) {
	l1p3 := Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1p1 := Point{
		51.930346153123985,
		4.45824900196975,
	}

	l1 := LineString{[]Point{
		l1p1, l1p2, l1p3,
	}}

	l2p1 := Point{
		Latitude:  51.93040734742036,
		Longitude: 4.458622634410858,
	}

	l2p2 := Point{
		Latitude:  51.93043380979297,
		Longitude: 4.459003508090973,
	}

	l2 := LineString{[]Point{
		l2p1, l2p2,
	}}

	crossed, point := l2.Intersect(l1)
	assert.True(t, crossed)
	assert.Equal(t, Point{51.93041738656155, 4.458767128471792}, point)

	t.Logf("crossed: %v and point is [ %v , %v ] \n", crossed, point.Longitude, point.Latitude)

}

func TestLineString_multiLineStringExtract(t *testing.T) {
	l1p1 := Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1p3 := Point{
		51.930346153123985,
		4.45824900196975,
	}
	l1 := LineString{[]Point{
		l1p1, l1p2, l1p3,
	}}

	lines := splitMultipleCoordinates(l1)
	assert.Equal(t, 2, len(lines))
	assert.Equal(t, l1p1, lines[0].Points[0])
	assert.Equal(t, l1p2, lines[0].Points[1])
	assert.Equal(t, l1p2, lines[1].Points[0])
	assert.Equal(t, l1p3, lines[1].Points[1])
}
