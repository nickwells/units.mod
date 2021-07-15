package units

import (
	"math"
)

// Unit represents a unit of measure
//
// There are conversion details that allow you to convert to and from the
// base units for the unit family. There are descriptive strings for display
// giving abbreviated and full names of the unit.
//
// The aliasName will only be set when the unit has been found through an
// alias rather than the canonical name.
type Unit struct {
	ConvPreAdd  float64
	ConvPostAdd float64
	ConvFactor  float64

	Fam        Family
	Abbrev     string
	Name       string
	NamePlural string
	Notes      string
	aliasName  string
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

// GetUnit returns a Unit of the given family and unit name, a non-nil error
// is returned if the familyName or unitName is not found. It is recommended
// that the constant values for familyName should be used to reduce the
// liklihood of runtime errors.
func GetUnit(familyName, unitName string) (Unit, error) {
	ud, err := GetUnitDetails(familyName)
	if err != nil {
		return Unit{}, err
	}

	return ud.GetUnit(unitName)
}

// GetUnitOrPanic will call GetUnit and check the error returned. If the
// error is non-nil it will panic, otherwise it will return the Unit
func GetUnitOrPanic(familyName, unitName string) Unit {
	u, err := GetUnit(familyName, unitName)
	if err != nil {
		panic(err)
	}
	return u
}
