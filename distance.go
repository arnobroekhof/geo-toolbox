package geo_toolbox

import (
	"errors"
	"math"
)

// these constants are used for Distance() fomulae in Vincenty
// ellipsoid in WGS-84
// length of semi-major axis of the ellipsoid (radius at equator);
const distanceA = 6378137

// 	length of semi-minor axis of the ellipsoid (radius at the poles);
const distanceB = 6356752.314245

// 	flattening of the ellipsoid;
const f = 1 / 298.257223563

//Distance calculate the distance between 2 points based on the Vincenty Alghoritm
// and return the miles and kilometers between the 2 points
func Distance(p1, p2 Point) (miles float64, kilometers float64, err error) {
	// convert from degrees to radians
	piRad := math.Pi / 180
	p1.Latitude = p1.Latitude * piRad
	p1.Longitude = p1.Longitude * piRad
	p2.Latitude = p2.Latitude * piRad
	p2.Longitude = p2.Longitude * piRad

	l := p2.Longitude - p1.Longitude

	u1 := math.Atan((1 - f) * math.Tan(p1.Latitude))
	u2 := math.Atan((1 - f) * math.Tan(p2.Latitude))

	sinU1 := math.Sin(u1)
	cosU1 := math.Cos(u1)
	sinU2 := math.Sin(u2)
	cosU2 := math.Cos(u2)

	lambda := l
	lambdaP := 2 * math.Pi
	iterLimit := 20

	var sinLambda, cosLambda, sinSigma float64
	var cosSigma, sigma, sinAlpha, cosSqAlpha, cos2SigmaM, C float64

	for {
		if math.Abs(lambda-lambdaP) > 1e-12 && (iterLimit > 0) {
			iterLimit -= 1
		} else {
			break
		}
		sinLambda = math.Sin(lambda)
		cosLambda = math.Cos(lambda)

		sinSigma = math.Sqrt((cosU2*sinLambda)*(cosU2*sinLambda) + (cosU1*sinU2-sinU1*cosU2*cosLambda)*(cosU1*sinU2-sinU1*cosU2*cosLambda))
		if sinSigma == 0 {
			return 0, 0, nil // co-incident points
		}

		cosSigma = sinU1*sinU2 + cosU1*cosU2*cosLambda
		sigma = math.Atan2(sinSigma, cosSigma)
		sinAlpha = cosU1 * cosU2 * sinLambda / sinSigma
		cosSqAlpha = 1 - sinAlpha*sinAlpha
		cos2SigmaM = cosSigma - 2*sinU1*sinU2/cosSqAlpha
		if math.IsNaN(cos2SigmaM) {
			cos2SigmaM = 0 // equatorial line: cosSqAlpha=0
		}

		C = f / 16 * cosSqAlpha * (4 + f*(4-3*cosSqAlpha))
		lambdaP = lambda
		lambda = l + (1-C)*f*sinAlpha*(sigma+C*sinSigma*(cos2SigmaM+C*cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)))
	}
	if iterLimit == 0 {
		return -1, -1, errors.New("formula failed to converge")
	}

	uSq := cosSqAlpha * (distanceA*distanceA - distanceB*distanceB) / (distanceB * distanceB)
	A := 1 + uSq/16384*(4096+uSq*(-768+uSq*(320-175*uSq)))
	B := uSq / 1024 * (256 + uSq*(-128+uSq*(74-47*uSq)))
	deltaSigma := B * sinSigma * (cos2SigmaM + B/4*(cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)-B/6*cos2SigmaM*(-3+4*sinSigma*sinSigma)*(-3+4*cos2SigmaM*cos2SigmaM)))
	meters := distanceB * A * (sigma - deltaSigma)
	kilometers = meters / 1000
	miles = kilometers * 0.621371
	return miles, kilometers, nil

}
