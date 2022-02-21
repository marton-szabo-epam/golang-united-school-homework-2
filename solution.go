package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type Square int64

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle Square = 3
	SidesSquare          = 4
	SidesCircle          = 0
)

func CalcSquare(sideLen float64, sidesNum Square) float64 {
	switch sidesNum {
	case SidesTriangle:
		return sideLen * sideLen / 2.0
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	}
	return 0
}
