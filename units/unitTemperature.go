package units

// UnitOfTemperature represents the base unit of temperature
var UnitOfTemperature = Family{
	BaseUnitName: "C",
	Description:  "unit of temperature",
}

const absZero = 273.15

// DegCUnit is a suitable default value for a UnitOfTemperature
var DegCUnit = Unit{
	0, 0, 1,
	UnitOfTemperature,
	"℃", UnitOfTemperature.BaseUnitName, "degrees Celsius",
	"a measure of temperature." +
		" It is named to honour the Swedish astronomer Anders Celsius." +
		" It was formerly known as the Centigrade scale" +
		" with units of 'centigrade'.",
}

var DegKUnit = Unit{
	0, absZero, 1,
	UnitOfTemperature,
	"K", "kelvin", "kelvin",
	"a measure of temperature based on the Celsius scale but" +
		" having zero at absolute zero (-273.15 on the Celsius" +
		" scale)." +
		" It is named to honour" +
		" the Glasgow university engineer and physicist" +
		" William Thomson, 1st Baron Kelvin.",
}

var DegFUnit = Unit{
	0, 32, 5.0 / 9.0,
	UnitOfTemperature,
	"℉", "degree Fahrenheit", "degrees Fahrenheit",
	"a measure of temperature. It is named after the physicist" +
		" Daniel Gabriel Fahrenheit." +
		" It has the freezing point of water at 32 degrees" +
		" and the boiling point of water at 212 degrees." +
		" Both using pure water at sea level.",
}

// The following are antique measures of historical interest only (like
// Fahrenheit?)

// DegRaUnit (Rankine) is the Fahrenheit equivalent of Kelvin
var DegRaUnit = Unit{
	absZero, 0, 5.0 / 9.0,
	UnitOfTemperature,
	"°R", "degree Rankine", "degrees Rankine",
	"a measure of temperature using degrees Fahrenheit but" +
		" having zero at absolute zero." +
		" It is named to honour" +
		" the Glasgow university engineer and physicist" +
		" William John Macquorn Rankine.",
}

var DegRoUnit = Unit{
	0, 7.5, 40.0 / 21.0,
	UnitOfTemperature,
	"°Rø", "degree Rømer", "degrees Rømer",
	"a measure of temperature. It is named after the Danish astronomer" +
		"  Ole Christensen Rømer." +
		"\n\n" +
		"No longer used." +
		"\n\n" +
		"It has the freezing point of water at 7.5 degrees" +
		" and the boiling point of water at 60 degrees." +
		" Both using pure water at sea level.",
}

var DegReUnit = Unit{
	0, 0, 4.0 / 5.0,
	UnitOfTemperature,
	"°Ré", "degree Réaumur", "degrees Réaumur",
	"a measure of temperature It is named after" +
		" René Antoine Ferchault de Réaumur." +
		"\n\n" +
		"No longer used." +
		"\n\n" +
		"It has the freezing point of water at 0 degrees" +
		" and the boiling point of water at 80 degrees.",
}

var DegNUnit = Unit{
	0, 0, 100.0 / 33.0,
	UnitOfTemperature,
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
}

// DegDUnit grows in the opposite direction to all the other units so that
// the larger the value, the colder it is
var DegDUnit = Unit{
	-100, 0, -2.0 / 3.0,
	UnitOfTemperature,
	"°D", "degree Delisle", "degrees Delisle",
	"a measure of temperature invented by Joseph-Nicolas Delisle." +
		"\n\n" +
		"No longer used." +
		"\n\n" +
		"It has the unusual property of hotter temperatures" +
		" having a lower value than colder ones. It runs from 0" +
		" at the boiling point of water to 150 at the freezing" +
		" point.",
}

// TemperatureNames maps names to units of temperature
var TemperatureNames = map[string]Unit{
	"C":  DegCUnit,
	"F":  DegFUnit,
	"K":  DegKUnit,
	"Ra": DegRaUnit,
	"Ro": DegRoUnit,
	"Re": DegReUnit,
	"N":  DegNUnit,
	"D":  DegDUnit,
}

var temperatureAliases = map[string]Alias{
	"Celsius":         {"C", ""},
	"degree Celsius":  {"C", ""},
	"degrees Celsius": {"C", "plural"},
	"Centigrade":      {"C", ""},

	"Fahrenheit":         {"F", ""},
	"degree Fahrenheit":  {"F", ""},
	"degrees Fahrenheit": {"F", "plural"},

	"Kelvin":         {"K", ""},
	"degree Kelvin":  {"K", ""},
	"degrees Kelvin": {"K", "plural"},

	"Rankine":         {"Ra", ""},
	"degree Rankine":  {"Ra", ""},
	"degrees Rankine": {"Ra", "plural"},

	"Romer":         {"Ro", ""},
	"degree Romer":  {"Ro", ""},
	"degrees Romer": {"Ro", "plural"},
	"Rømer":         {"Ro", ""},
	"Roemer":        {"Ro", ""},

	"Réaumur":         {"Re", ""},
	"Reaumur":         {"Re", ""},
	"degree Reaumur":  {"Re", ""},
	"degrees Reaumur": {"Re", "plural"},

	"Newton":         {"N", ""},
	"degree Newton":  {"N", ""},
	"degrees Newton": {"N", "plural"},

	"Delisle":         {"D", ""},
	"degree Delisle":  {"D", ""},
	"degrees Delisle": {"D", "plural"},
}
