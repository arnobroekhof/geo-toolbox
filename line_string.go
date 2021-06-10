package geo_toolbox

type LineString struct {
	Points []Point
}

func splitMultipleCoordinates(l LineString) (lines []LineString) {
	count := len(l.Points)
	if count > 2 {
		for k := range l.Points {
			if k == count-1 {
				break
			}
			line := LineString{
				[]Point{
					l.Points[k],
					l.Points[k+1],
				}}
			lines = append(lines, line)
		}
	}
	return lines
}

func (l LineString) Intersect(line LineString) (crossed bool, point Point) {
	if lp := len(line.Points); lp > 2 {
		for _, v := range splitMultipleCoordinates(line) {
			if crossed, point = intersect(l, v); crossed {
				return crossed, point
			}
		}
	}
	return intersect(l, line)
}
func intersect(line1 LineString, line2 LineString) (bool, Point) {

	switch lp := len(line1.Points); {
	case lp > 2:
		for _, v := range splitMultipleCoordinates(line1) {
			if c := crossed(v, line2); c {
				if p := extractIntersectionPoint(v, line2); p != nil {
					return true, *p
				}
			}
		}
	case lp == 2:
		if c := crossed(line1, line2); c {
			if p := extractIntersectionPoint(line1, line2); p != nil {
				return true, *p
			}
		}

	default:
		return false, Point{}
	}

	return false, Point{}
}

func extractIntersectionPoint(line1, line2 LineString) *Point {
	x1 := line1.Points[0].X()
	y1 := line1.Points[0].Y()
	x2 := line1.Points[1].X()
	y2 := line1.Points[1].Y()

	x3 := line2.Points[0].X()
	y3 := line2.Points[0].Y()
	x4 := line2.Points[1].X()
	y4 := line2.Points[1].Y()

	return intersectionPoint(x1, y1, x2, y2, x3, y3, x4, y4)

}

func intersectionPoint(x1, y1, x2, y2, x3, y3, x4, y4 float64) *Point {
	px := ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	py := ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))

	if px > 0 && py > 0 {
		return &Point{
			Longitude: px,
			Latitude:  py,
		}
	}

	return nil
}

// crossed check if intersect
// got the formula from http://www.bryceboe.com/2006/10/23/line-segment-intersection-algorithm/
func crossed(line1, line2 LineString) bool {
	a := line1.Points[0]
	b := line1.Points[1]
	c := line2.Points[0]
	d := line2.Points[1]
	return ccw(a, c, d) != ccw(b, c, d) && ccw(a, b, c) != ccw(a, b, d)
}

func ccw(a, b, c Point) bool {
	return (c.Y()-a.Y())*(b.X()-a.X()) > (b.Y()-a.Y())*(c.X()-a.X())
}
