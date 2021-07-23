package units

const bunAngle = "radian"

// angleFamily represents the collection of units of angular measure
var angleFamily = &Family{
	baseUnitName: bunAngle,
	description:  "unit of angular measure",
	name:         Angle,
}

// AngleNames maps names to units of anglular measure
var angleNames = map[string]Unit{
	// SI
	bunAngle: {
		0, 0, 1,
		angleFamily,
		"rad", bunAngle, "radians",
		"a unit in which angles are measured.",
		nil,
		map[string]string{},
		"", "",
	},

	"milliradian": {
		0, 0, m,
		angleFamily,
		"mrad", "milliradian", "milliradians",
		"a unit in which angles are measured.",
		nil,
		map[string]string{},
		"", "",
	},

	// degrees
	"degree": {
		0, 0, degreePerRadian,
		angleFamily,
		"°", "degree", "degrees",
		"a unit in which angles are measured.",
		nil,
		map[string]string{},
		"", "",
	},
	"minute": {
		0, 0, degreePerRadian / 60,
		angleFamily,
		"′", "arc minute", "arc minutes",
		"a unit in which angles are measured.",
		nil,
		map[string]string{},
		"", "",
	},
	"second": {
		0, 0, degreePerRadian / 3600,
		angleFamily,
		"″", "arc second", "arc seconds",
		"a unit in which angles are measured.",
		nil,
		map[string]string{},
		"", "",
	},
}
