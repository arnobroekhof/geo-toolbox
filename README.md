# Geo Toolbox

toolbox for performing intersect operations.

currently only LineStrings are supported

install:

```
go get github.com/arnobroekhof/geo-toolbox
```

for more examples look inside the tests

simple usage:

```
l1p1 := geo_toolbox.Point{
		Latitude:  51.93052146629074,
		Longitude: 4.458684325218201,
	}
	l1p2 := geo_toolbox.Point{
		Latitude:  51.930346153123985,
		Longitude: 4.458823800086975,
	}
	l1 := geo_toolbox.LineString{Points: []geo_toolbox.Point{
		l1p1, l1p2,
	}}

	l2p1 := geo_toolbox.Point{
		Latitude:  51.83040734742036,
		Longitude: 4.458622634410858,
	}

	l2p2 := geo_toolbox.Point{
		Latitude:  51.83043380979297,
		Longitude: 4.459003508090973,
	}

	l2 := geo_toolbox.LineString{Points: []geo_toolbox.Point{
		l2p1, l2p2,
	}}

	crossed, intersect := l1.Intersect(l2)

```

![Example Map](docs/img1.png?raw=true "Example")
