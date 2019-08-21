package units

import "math"

const degreePerRadian = math.Pi / 180

// UnitOfAngle represents the base unit of angular measure
var UnitOfAngle = Family{
	BaseUnitName: "radian",
	Description:  "unit of angular measure",
}

// DegCUnit is a suitable default value for a UnitOfAngle
var AngularUnitRadian = Unit{0, 0, 1, UnitOfAngle, "rad", UnitOfAngle.BaseUnitName, "radians"}

// AngleNames maps names to units of temperature
var AngleNames = map[string]Unit{
	// SI
	"radian":      AngularUnitRadian,
	"milliradian": {0, 0, m, UnitOfAngle, "mrad", "milliradian", "milliradians"},

	// degrees
	"degree": {0, 0, degreePerRadian, UnitOfAngle, "°", "degree", "degrees"},
	"minute": {0, 0, degreePerRadian / 60, UnitOfAngle, "′", "arc minute", "arc minutes"},
	"second": {0, 0, degreePerRadian / 3600, UnitOfAngle, "″", "arc second", "arc seconds"},
}
