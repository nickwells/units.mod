package units

import (
	"errors"
	"fmt"
	"math"
)

// Unit represents a unit of measure
//
// There are conversion details that allow you to convert to and from the
// base units for the unit family. There are descriptive strings for display
// giving abbreviated and full names of the unit.
type Unit struct {
	ConvPreAdd  float64
	ConvPostAdd float64
	ConvFactor  float64

	Fam        Family
	Abbrev     string
	Name       string
	NamePlural string
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

// GetUnit returns a Unit of the given type and unit name, a non-nil error is
// returned if the familyName or unitName is not found. It is recommended
// that the constant values for familyName should be used to reduce the
// liklihood of runtime errors.
func GetUnit(familyName, unitName string) (Unit, error) {
	var u Unit

	ud, ok := validUnits[familyName]
	if !ok {
		return u, errors.New("no such unit type '" + familyName + "'")
	}
	u, ok = ud.AltU[unitName]
	if !ok {
		return u,
			fmt.Errorf("there is no %s with a name of '%s'",
				ud.Fam.Description, unitName)
	}
	return u, nil
}
