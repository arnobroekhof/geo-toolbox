# Geo Toolbox

toolbox for performing intersect operation.

currently only LineStrings are supported

install:

```
go get github.com/arnobroekhof/geo-toolbox
```

for more examples look inside the tests

simple usage:

```
l1p1 := Point{
    latitude:  51.93052146629074,
    longitude: 4.458684325218201,
}
l1p2 := Point{
    latitude:  51.930346153123985,
    longitude: 4.458823800086975,
}
l1 := LineString{[]Point{
    l1p1, l1p2,
}}

l2p1 := Point{
    latitude:  51.93040734742036,
    longitude: 4.458622634410858,
}

l2p2 := Point{
    latitude:  51.93043380979297,
    longitude: 4.459003508090973,
}

l2 := LineString{[]Point{
    l2p1, l2p2,
}}

// crossed: true and point is [ 4.458767128471792 , 51.93041738656155 ]
crossed, point := l1.Intersect(l2) 

```

![Example Map](docs/img1.png?raw=true "Example")
