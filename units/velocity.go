package units

// bunVelocity is the base unit name for velocity
const bunVelocity = "metre/second"

// velocityFamily represents the base unit of Velocity
var velocityFamily = &Family{
	baseUnitName:  bunVelocity,
	description:   "unit of velocity",
	name:          Velocity,
	familyAliases: []string{"speed"},
}

// VelocityNames maps names to units of velocity
var velocityNames = map[string]Unit{
	// metric
	bunVelocity: {
		0, 0, 1,
		velocityFamily,
		"m/s", bunVelocity, "metres/second",
		"The base unit of velocity in the metric system.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"metres/second": "plural",
			"meters/second": "US spelling, plural",
			"meter/second":  "US spelling",
		},
		"", "",
	},

	"kilometre/hour": {
		0, 0, 1000.0 / 3600.0,
		velocityFamily,
		"km/h", "kilometre/hour", "kilometres/hour",
		"a metric measure of velocity.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"kmph":            "abbreviation",
			"km/h":            "abbreviation",
			"kilometres/hour": "plural",
			"kilometers/hour": "US spelling, plural",
			"kilometer/hour":  "US spelling",
		},
		"", "",
	},
	"foot/second": {
		0, 0, footToMetre,
		velocityFamily,
		"ft/s", "foot/sec", "feet/sec",
		"an imperial measure of velocity.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"feet/second": "plural",
			"foot/sec":    "abbreviation",
			"feet/sec":    "abbreviation, plural",
		},
		"", "",
	},
	"mile/hour": {
		0, 0, mileToMetre / 3600.0,
		velocityFamily,
		"mph", "mile/hour", "miles/hour",
		"an imperial measure of velocity.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"mph":        "abbreviation",
			"m/h":        "abbreviation",
			"miles/hour": "plural",
		},
		"", "",
	},
	"percentOfSpeedOfLight": {
		0, 0, lightSecond / 100,
		velocityFamily,
		"%c", "%c", "%c",
		"as a percentage of the speed of light.",
		[]Tag{TagPhysics},
		map[string]string{},
		"", "",
	},
	"knot": {
		0, 0, nauticalMileToMetre / 3600.0,
		velocityFamily,
		"kn", "knot", "knots",
		"an imperial measure of velocity." +
			" It is equal to one nautical mile per hour." +
			"\n\n" +
			"It was historically calculated by dropping overboard" +
			" a 'chip log' attached to a cord and measuring how many" +
			" knots tied in the cord passed through a sailor's fingers" +
			" Another sailor timed the operation using a sand glass." +
			" The knots were spaced every 47 feet 3 inches and" +
			" the sand timer measured 28 seconds." +
			" This method differs from the modern definition by" +
			" less than 0.02% (presumably much less than any errors in" +
			" the measurements).",
		[]Tag{TagImperial, TagUScustomary, TagNautical},
		map[string]string{},
		"", "",
	},
}
