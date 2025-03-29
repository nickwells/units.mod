package units

// bunTemp is the base unit name for temperature
const bunTemp = "C"

// temperatureFamily represents the base unit of temperature
var temperatureFamily = &Family{
	baseUnitName:  bunTemp,
	description:   "unit of temperature",
	name:          Temperature,
	familyAliases: []string{"temp"},
}

// degCUnit is a suitable default value for a temperatureFamily
var degCUnit = Unit{
	0, 0, 1,
	temperatureFamily,
	"°C", temperatureFamily.baseUnitName, "degrees Celsius",
	"a measure of temperature." +
		" It is named to honour the Swedish astronomer Anders Celsius." +
		" It was formerly known as the Centigrade scale" +
		" with units of 'centigrade'." +
		"\n\n" +
		"When Celsius created the scale in 1742 he had" +
		" 0 degrees as the boiling point of water and" +
		" 100 degrees as the freezing point (like the Delisle scale)." +
		" It was inverted to the more familiar modern scale" +
		" by Jean-Pierre Christin in 1743.", //nolint:misspell
	[]Tag{TagMetric},
	map[string]string{
		"Celsius":         "full name",
		"degree Celsius":  "full name, with degree",
		"degrees Celsius": "full name, with degree, plural",
		"Centigrade":      "alternative name",
		"c":               "lower-case",
	},
	"", "",
}

var degKUnit = Unit{
	0, absZero, 1,
	temperatureFamily,
	"K", "kelvin", "kelvin",
	"a measure of temperature based on the Celsius scale but" +
		" having zero at absolute zero (-273.15 on the Celsius" +
		" scale)." +
		" It is named to honour" +
		" the Glasgow university engineer and physicist" +
		" William Thomson, 1st Baron Kelvin.",
	[]Tag{TagSI, TagMetric},
	map[string]string{
		"Kelvin":         "full name",
		"degree Kelvin":  "full name, with degree",
		"degrees Kelvin": "full name, with degree, plural",
	},
	"", "",
}

var degFUnit = Unit{
	0, 32, 5.0 / 9.0,
	temperatureFamily,
	"°F", "degree Fahrenheit", "degrees Fahrenheit",
	"a measure of temperature. It is named after the physicist" +
		" Daniel Gabriel Fahrenheit." +
		" It has the freezing point of water at 32 degrees" +
		" and the boiling point of water at 212 degrees." +
		" Both using pure water at sea level." +
		"\n\n" +
		"It is only still used in the United States and its" +
		" territories and a few small countries.",
	[]Tag{TagUScustomary},
	map[string]string{
		"Fahrenheit":         "full name",
		"degree Fahrenheit":  "full name, with degree",
		"degrees Fahrenheit": "full name, with degree, plural",
		"f":                  "lower-case",
	},
	"", "",
}

// The following are antique measures of historical interest only (like
// Fahrenheit?)

// degRaUnit (Rankine) is the Fahrenheit equivalent of Kelvin
var degRaUnit = Unit{
	absZero, 0, 5.0 / 9.0,
	temperatureFamily,
	"°R", "degree Rankine", "degrees Rankine",
	"a measure of temperature using degrees Fahrenheit but" +
		" having zero at absolute zero." +
		" It is named to honour" +
		" the Glasgow university engineer and physicist" +
		" William John Macquorn Rankine.",
	[]Tag{TagHist},
	map[string]string{
		"Rankine":         "full name",
		"degree Rankine":  "full name, with degree",
		"degrees Rankine": "full name, with degree, plural",
	},
	"", "",
}

var degRoUnit = Unit{
	0, 7.5, 40.0 / 21.0,
	temperatureFamily,
	"°Rø", "degree Rømer", "degrees Rømer",
	"a measure of temperature. It is named after the Danish astronomer" +
		"  Ole Christensen Rømer." +
		"\n\n" +
		"No longer used." +
		"\n\n" +
		"It has the freezing point of water at 7.5 degrees" +
		" and the boiling point of water at 60 degrees." +
		" Both using pure water at sea level.",
	[]Tag{TagHist},
	map[string]string{
		"Romer":         "full name",
		"degree Romer":  "full name, with degree",
		"degrees Romer": "full name, with degree, plural",
		"Rømer":         "full name, with correct spelling",
		"Roemer":        "full name, with oe-spelling",
	},
	"", "",
}

var degReUnit = Unit{
	0, 0, 5.0 / 4.0,
	temperatureFamily,
	"°Ré", "degree Réaumur", "degrees Réaumur",
	"a measure of temperature It is named after" +
		" René Antoine Ferchault de Réaumur." +
		" The Réaumur scale is also known as" +
		" the octogesimal division." +
		"\n\n" +
		"No longer used." +
		"\n\n" +
		"It has the freezing point of water at 0 degrees" +
		" and the boiling point of water at 80 degrees." +
		"\n\n" +
		"It was used by Charles Joseph Minard in his famous" +
		" infographic depicting Napoleon's disastrous 1812 Russian campaign" +
		" despite France having adopted the Celsius scale in the 1790s.",
	[]Tag{TagHist},
	map[string]string{
		"Réaumur":         "full name, with accent",
		"Reaumur":         "full name",
		"degree Reaumur":  "full name, with degree",
		"degrees Reaumur": "full name, with degree, plural",
	},
	"", "",
}

var degNUnit = Unit{
	0, 0, 100.0 / 33.0,
	temperatureFamily,
	"°N", "degree Newton", "degrees Newton",
	"a measure of temperature devised by Isaac Newton." +
		"\n\n" +
		"No longer used." +
		"\n\n" +
		"It has the freezing point of water at 0 degrees" +
		" and the boiling point of water at 33 degrees." +
		" It is poorly defined and no unambiguous" +
		" conversion to other scales is possible but the" +
		" figures used refer to points given by Newton.",
	[]Tag{TagHist},
	map[string]string{
		"Newton":         "full name",
		"degree Newton":  "full name, with degree",
		"degrees Newton": "full name, with degree, plural",
	},
	"", "",
}

// degDUnit grows in the opposite direction to all the other units so that
// the larger the value, the colder it is
var degDUnit = Unit{
	-100, 0, -2.0 / 3.0,
	temperatureFamily,
	"°D", "degree Delisle", "degrees Delisle",
	"a measure of temperature invented by Joseph-Nicolas Delisle." +
		"\n\n" +
		"No longer used." +
		"\n\n" +
		"It has the unusual property of hotter temperatures" +
		" having a lower value than colder ones. It runs from 0" +
		" at the boiling point of water to 150 at the freezing" +
		" point.",
	[]Tag{TagHist},
	map[string]string{
		"Delisle":         "full name",
		"degree Delisle":  "full name, with degree",
		"degrees Delisle": "full name, with degree, plural",
	},
	"", "",
}

// TemperatureNames maps names to units of temperature
var temperatureNames = map[string]Unit{
	bunTemp: degCUnit,
	"F":     degFUnit,
	"K":     degKUnit,
	"Ra":    degRaUnit,
	"Ro":    degRoUnit,
	"Re":    degReUnit,
	"N":     degNUnit,
	"D":     degDUnit,
}
