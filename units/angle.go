package units

import "math"

const degreePerRadian = math.Pi / 180

// UnitOfAngle represents the base unit of angular measure
var UnitOfAngle = Family{
	BaseUnitName: "radian",
	Description:  "unit of angular measure",
	Name:         Angle,
}

// AngularUnitRadian is a suitable default value for a UnitOfAngle
var AngularUnitRadian = Unit{
	0, 0, 1,
	UnitOfAngle,
	"rad", UnitOfAngle.BaseUnitName, "radians",
	"a unit in which angles are measured.",
	"",
}

// AngleNames maps names to units of anglular measure
var AngleNames = map[string]Unit{
	// SI
	"radian": AngularUnitRadian,
	"milliradian": {
		0, 0, m,
		UnitOfAngle,
		"mrad", "milliradian", "milliradians",
		"a unit in which angles are measured.",
		"",
	},

	// degrees
	"degree": {
		0, 0, degreePerRadian,
		UnitOfAngle,
		"°", "degree", "degrees",
		"a unit in which angles are measured.",
		"",
	},
	"minute": {
		0, 0, degreePerRadian / 60,
		UnitOfAngle,
		"′", "arc minute", "arc minutes",
		"a unit in which angles are measured.",
		"",
	},
	"second": {
		0, 0, degreePerRadian / 3600,
		UnitOfAngle,
		"″", "arc second", "arc seconds",
		"a unit in which angles are measured.",
		"",
	},
}

var angleAliases = map[string]Alias{}
