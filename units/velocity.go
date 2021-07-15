package units

// UnitOfVelocity represents the base unit of Velocity
var UnitOfVelocity = Family{
	BaseUnitName: "metrePerSecond",
	Description:  "unit of velocity",
	Name:         Velocity,
}

// MetrePerSecondUnit is a suitable default value for a UnitOfVelocity
var MetrePerSecondUnit = Unit{
	0, 0, 1,
	UnitOfVelocity, "m/s", UnitOfVelocity.BaseUnitName + "/sec", "metres/sec",
	"The base unit of velocity in the metric system.",
	"",
}

// VelocityNames maps names to units of velocity
var VelocityNames = map[string]Unit{
	// metric
	"metresPerSecond": MetrePerSecondUnit,
	"kilometresPerHour": {
		0, 0, 1000.0 / 3600.0,
		UnitOfVelocity,
		"km/h", "kilometre/hour", "kilometres/hour",
		"a metric measure of velocity.",
		"",
	},
	"feetPerSecond": {
		0, 0, footToMetre,
		UnitOfVelocity,
		"ft/s", "foot/sec", "feet/sec",
		"an imperial measure of velocity.",
		"",
	},
	"milesPerHour": {
		0, 0, (yardToMetre * 1760) / 3600,
		UnitOfVelocity,
		"mph", "mile/hour", "miles/hour",
		"an imperial measure of velocity.",
		"",
	},
	"percentOfSpeedOfLight": {
		0, 0, lightSecond / 100,
		UnitOfVelocity,
		"%c", "%c", "%c",
		"as a percentage of the speed of light.",
		"",
	},
}

var velocityAliases = map[string]Alias{
	"mph":        {"milesPerHour", "abbreviation"},
	"m/h":        {"milesPerHour", "abbreviation"},
	"miles/hour": {"milesPerHour", "abbreviation"},

	"metrePerSecond":  {"metresPerSecond", "singular"},
	"metersPerSecond": {"metresPerSecond", "US spelling"},
	"meterPerSecond":  {"metresPerSecond", "US spelling, singular"},

	"kmph":              {"kilometresPerHour", "abbreviation"},
	"km/h":              {"kilometresPerHour", "abbreviation"},
	"kilometrePerHour":  {"kilometresPerHour", "singular"},
	"kilometersPerHour": {"kilometresPerHour", "US spelling"},
	"kilometerPerHour":  {"kilometresPerHour", "US spelling, singular"},
}
