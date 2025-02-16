//nolint:mnd
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
		[]Tag{TagTrig, TagSI},
		map[string]string{
			"radians": "plural",
			"rad":     "abbreviation",
			"rads":    "abbreviation, plural",
		},
		"", "",
	},

	"milliradian": {
		0, 0, m,
		angleFamily,
		"mrad", "milliradian", "milliradians",
		"a unit in which angles are measured.",
		[]Tag{TagTrig, TagSI},
		map[string]string{},
		"", "",
	},

	// gradians
	"gradian": {
		0, 0, degreePerGradian,
		angleFamily,
		"gon", "gradian", "gradians",
		"The gradian originated during the French Revolution" +
			" along with the metric system." +
			" It has 100 degrees to the right angle rather than 90." +
			" It is a legal unit of measurement in the European Union" +
			" and in Switzerland. principally used in surveying.",
		[]Tag{TagTrig, TagMetric},
		map[string]string{
			"gradians": "plural",
			"gon":      "abbreviation",
			"gons":     "abbreviation, plural",
			"grad":     "abbreviation",
			"grads":    "abbreviation, plural",
		},
		"", "",
	},

	// degrees
	"degree": {
		0, 0, degreePerRadian,
		angleFamily,
		"°", "degree", "degrees",
		"a unit in which angles are measured.",
		[]Tag{TagTrig},
		map[string]string{},
		"", "",
	},
	"minute": {
		0, 0, degreePerRadian / 60,
		angleFamily,
		"′", "arc minute", "arc minutes",
		"a unit in which angles are measured.",
		[]Tag{TagTrig},
		map[string]string{},
		"", "",
	},
	"second": {
		0, 0, degreePerRadian / 3600,
		angleFamily,
		"″", "arc second", "arc seconds",
		"a unit in which angles are measured.",
		[]Tag{TagTrig},
		map[string]string{},
		"", "",
	},
}
