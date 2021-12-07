# Geo Toolbox

toolbox for performing Geo related operations.

* instersects
* speed
* distance
* journey


install:

```
go get github.com/arnobroekhof/geo-toolbox
```


## Intersect
*NOTE* currently only LineStrings are supported


## Distance between 2 points
calculate the distance between 2 points based on the [Vincenty Alghoritm](https://en.wikipedia.org/wiki/Vincenty%27s_formulae)

## Examples
for more examples look inside the tests

### Intersect point:

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

### Distance
```
rotterdam := Point{51.92250, 4.47917}
amsterdam := Point{52.37403, 4.88969}

miles, km, err := Distance(rotterdam, amsterdam)

```

### Speed

```
rotterdam := Point{51.92250, 4.47917}
amsterdam := Point{52.37403, 4.88969}
utrecht := Point{52.0908, 5.1222}

tn := time.Now()
t1 := tn.Add(- 3 * time.Hour)
t2 := tn.Add(- 2 * time.Hour)
t3 := tn.Add(- 1 * time.Hour)

tpRotterdam := NewTimebasedPoint(t1, rotterdam)
tpAmsterdam := NewTimebasedPoint(t2, amsterdam)
tpUtrecht := NewTimebasedPoint(t3, utrecht)

kph  := SpeedMultiple(tpRotterdam, tpAmsterdam, tpUtrecht)
```

### Journey

```
rotterdam := Point{51.92250, 4.47917}
amsterdam := Point{52.37403, 4.88969}
utrecht := Point{52.0908, 5.1222}
tn := time.Now()
t1 := tn.Add(- 3 * time.Hour)
t2 := tn.Add(- 2 * time.Hour)
t3 := tn.Add(- 1 * time.Hour)

tpRotterdam := NewTimebasedPoint(t1, rotterdam)
tpAmsterdam := NewTimebasedPoint(t2, amsterdam)
tpUtrecht := NewTimebasedPoint(t3, utrecht)

journey := NewJourney(tpRotterdam, tpAmsterdam, tpUtrecht)
```
