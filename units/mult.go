package units

import (
	"errors"
	"fmt"
	"math"
)

// Mult represents a multiple that can be applied to a unit of size
// such as K (for 1000) or M (for 1000000)
//
// The Val gives the number of the BaseUnit that there are in the alternative
// unit. Given a measure in the base unit you can express this in the
// alternative unit by dividing by the Val
//
// for multiples of bytes the names follow the metric and IEC naming
// conventions so MB refers to a megabyte or 1000^2 bytes and MiB refers
// to a mebibyte or 1024^2 bytes
type Mult struct {
	Val      float64
	BaseUnit Unit
	Abbrev   string
	AltUnit  string
}

var (
	k  = 1000.0
	_M = math.Pow(1000, 2)
	_G = math.Pow(1000, 3)
	_T = math.Pow(1000, 4)
	_P = math.Pow(1000, 5)
	_E = math.Pow(1000, 6)
	_Z = math.Pow(1000, 7)
	_Y = math.Pow(1000, 8)

	h  = 100.0
	da = 10.0

	d = 0.1
	c = 0.01

	m = 1.0 / math.Pow(1000, 1)
	u = 1.0 / math.Pow(1000, 2)
	n = 1.0 / math.Pow(1000, 3)
	p = 1.0 / math.Pow(1000, 4)
	f = 1.0 / math.Pow(1000, 5)
	a = 1.0 / math.Pow(1000, 6)
	z = 1.0 / math.Pow(1000, 7)
	y = 1.0 / math.Pow(1000, 8)
)

// GetMult returns a Mult of the given type and unit name, a non-nil error is
// returned if the typeName or unitName is not found. It is recommended that
// the constant values for typeName should be used to reduce the liklihood of
// runtime errors.
func GetMult(typeName, unitName string) (Mult, error) {
	var m Mult

	ud, ok := validUnits[typeName]
	if !ok {
		return m, errors.New("no such unit type '" + typeName + "'")
	}
	m, ok = ud.AltU[unitName]
	if !ok {
		return m,
			fmt.Errorf("there is no %s with a name of '%s'",
				ud.U.Description, unitName)
	}
	return m, nil
}
