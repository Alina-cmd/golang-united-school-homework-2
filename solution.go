package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

const Pi float64 = math.Pi

const (
	SidesTriangle int = 3
	SidesSquare   int = 4
	SidesCircle   int = 0
)

type MyType int

func CalcSquare(sideLen float64, sidesNum MyType) float64 {
	switch sidesNum {
	case 0:
		return Pi * math.Pow(sideLen, 2)
	case 3:
		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case 4:
		return sideLen * sideLen
	default:
		return 0
	}
}
