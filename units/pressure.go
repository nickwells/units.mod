package units

// bunPressure is the base unit name for pressure
const bunPressure = "pascal"

// pressureFamily represents the base unit of pressure
var pressureFamily = &Family{
	baseUnitName:  bunPressure,
	description:   "unit of pressure or stress",
	name:          Pressure,
	familyAliases: []string{"stress"},
}

// PressureNames maps names to units of pressure
var pressureNames = map[string]Unit{
	bunPressure: {
		0, 0, 1,
		pressureFamily,
		"Pa", bunPressure, "pascals",
		"The base unit of pressure in the SI (metric) system." +
			" It is defined as one Newton per square metre." +
			" It is named after Blaise Pascal (1623-1662) a" +
			" French mathematician and physicist (among other things)",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"Pa":      "abbreviation",
			"pascals": "plural",
		},
		"", "",
	},
	"mPa": {
		0, 0, m,
		pressureFamily,
		"mPa", "millipascal", "millipascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"millipascal":  "",
			"millipascals": "plural",
		},
		"", "",
	},
	"cPa": {
		0, 0, c,
		pressureFamily,
		"cPa", "centipascal", "centipascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"centipascal":  "",
			"centipascals": "plural",
		},
		"", "",
	},
	"dPa": {
		0, 0, d,
		pressureFamily,
		"dPa", "decipascal", "decipascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"decipascal":  "",
			"decipascals": "plural",
		},
		"", "",
	},
	"hPa": {
		0, 0, h,
		pressureFamily,
		"hPa", "hectopascal", "hectopascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"hectopascal":  "",
			"hectopascals": "plural",
		},
		"", "",
	},
	"kPa": {
		0, 0, k,
		pressureFamily,
		"kPa", "kilopascal", "kilopascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"kilopascal":  "",
			"kilopascals": "plural",
		},
		"", "",
	},
	"MPa": {
		0, 0, _M,
		pressureFamily,
		"MPa", "megapascal", "megapascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"megapascal":  "",
			"megapascals": "plural",
		},
		"", "",
	},
	"GPa": {
		0, 0, _G,
		pressureFamily,
		"GPa", "gigapascal", "gigapascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"gigapascal":  "",
			"gigapascals": "plural",
		},
		"", "",
	},
	"TPa": {
		0, 0, _T,
		pressureFamily,
		"TPa", "terapascal", "terapascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"terapascal":  "",
			"terapascals": "plural",
		},
		"", "",
	},
	"PPa": {
		0, 0, _P,
		pressureFamily,
		"PPa", "petapascal", "petapascals",
		"an SI unit of pressure",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"petapascal":  "",
			"petapascals": "plural",
		},
		"", "",
	},
	"standard atmosphere": {
		0, 0, atmosphereToPascal,
		pressureFamily,
		"atm", "atmosphere", "atmospheres",
		"A unit of pressure approximately equal to" +
			" Earth's average atmospheric pressure at sea level",
		[]Tag{TagMetric},
		map[string]string{
			"atm":                  "abbreviation",
			"standard atmospheres": "plural",
			"atmosphere":           "abbreviation",
			"atmospheres":          "plural",
		},
		"", "",
	},
	"bar": {
		0, 0, barToPascal,
		pressureFamily,
		"bar", "bar", "bars",
		"A unit of pressure introduced by the Norwegian meteorologist" +
			" Vilhelm Bjerkenes. Its use is now deprecated" +
			" it is roughly equal to the atmospheric pressure on Earth" +
			" at an altitude of 111 metres at 15 ℃.",
		[]Tag{TagMetric},
		map[string]string{
			"bars": "plural",
		},
		"", "",
	},
	"millibar": {
		0, 0, barToPascal * m,
		pressureFamily,
		"mbar", "millibar", "millibars",
		"A unit of pressure.",
		[]Tag{TagMetric},
		map[string]string{
			"mbar":      "abbreviation",
			"millibars": "plural",
		},
		"", "",
	},
	"centibar": {
		0, 0, barToPascal * c,
		pressureFamily,
		"cbar", "centibar", "centibars",
		"A unit of pressure.",
		[]Tag{TagMetric},
		map[string]string{
			"cbar":      "abbreviation",
			"centibars": "plural",
		},
		"", "",
	},
	"decibar": {
		0, 0, barToPascal * d,
		pressureFamily,
		"dbar", "decibar", "decibars",
		"A unit of pressure.",
		[]Tag{TagMetric},
		map[string]string{
			"dbar":     "abbreviation",
			"decibars": "plural",
		},
		"", "",
	},
	"kilobar": {
		0, 0, barToPascal * k,
		pressureFamily,
		"kbar", "kilobar", "kilobars",
		"A unit of pressure.",
		[]Tag{TagMetric},
		map[string]string{
			"kbar":     "abbreviation",
			"kilobars": "plural",
		},
		"", "",
	},
	"megabar": {
		0, 0, barToPascal * _M,
		pressureFamily,
		"Mbar", "megabar", "megabars",
		"A unit of pressure.",
		[]Tag{TagMetric},
		map[string]string{
			"Mbar":     "abbreviation",
			"megabars": "plural",
		},
		"", "",
	},
	"psi": {
		0, 0, psiToPascal,
		pressureFamily,
		"psi", "pound per square inch", "pounds per square inch",
		"A unit of pressure measured in pounds force per square inch.",
		[]Tag{TagUScustomary},
		map[string]string{
			"pound per square inch":  "",
			"pounds per square inch": "plural",
		},
		"", "",
	},
	"kpsi": {
		0, 0, psiToPascal * k,
		pressureFamily,
		"kpsi", "kilopound per square inch", "kilopounds per square inch",
		"A unit of pressure measured in pounds force per square inch.",
		[]Tag{TagUScustomary},
		map[string]string{
			"kilopound per square inch":  "",
			"kilopounds per square inch": "plural",
		},
		"", "",
	},
	"Mpsi": {
		0, 0, psiToPascal * _M,
		pressureFamily,
		"Mpsi", "megapound per square inch", "megapounds per square inch",
		"A unit of pressure measured in pounds force per square inch.",
		[]Tag{TagUScustomary},
		map[string]string{
			"megapound per square inch":  "",
			"megapounds per square inch": "plural",
		},
		"", "",
	},
	"mmHg": {
		0, 0, mmHgToPascal,
		pressureFamily,
		"mmHg", "millimetre of mercury", "millimetres of mercury",
		"A unit of pressure measured in millimetres of mercury with" +
			" a density of 13,595.1 kilograms per cubic metre at 0 ℃" +
			" at standard gravity (9.80665 metres per second squared)." +
			" It is almost equal to the Torr, differing by only one part" +
			" in seven million.",
		[]Tag{TagUScustomary},
		map[string]string{
			"millimetre of mercury":  "",
			"millimetres of mercury": "plural",
		},
		"", "",
	},
	"Torr": {
		0, 0, torrToPascal,
		pressureFamily,
		"Torr", "torr", "torrs",
		"A unit of pressure defined to be 1/760 of a standard atmosphere." +
			" It is named after Evangelista Torricelli an Italiam physicist" +
			" and mathematician." +
			"\n\n" +
			" It was originally intended to be equal to the mmHg," +
			" but subsequent redefinitions made it slightly lower by" +
			" approximately one part in seven million.",
		[]Tag{TagMetric},
		map[string]string{
			"torr": "capitalisation",
		},
		"", "",
	},
	"barye": {
		0, 0, d,
		pressureFamily,
		"Ba", "barye", "baryes",
		"A unit of pressure in the" +
			" centimetre-gram-second (CGS) system of units.",
		[]Tag{TagMetric},
		map[string]string{
			"Ba":     "abbreviation",
			"baryes": "plural",
			"barad":  "alternative name",
			"barrie": "alternative name",
			"bary":   "alternative name",
			"baryd":  "alternative name",
			"baried": "alternative name",
			"barie":  "alternative name",
		},
		"", "",
	},
	"millibarye": {
		0, 0, d * m,
		pressureFamily,
		"mBa", "millibarye", "millibaryes",
		"A unit of pressure in the" +
			" centimetre-gram-second (CGS) system of units.",
		[]Tag{TagMetric},
		map[string]string{
			"mBa":         "abbreviation",
			"millibaryes": "plural",
			"millibarad":  "alternative name",
			"millibarrie": "alternative name",
			"millibary":   "alternative name",
			"millibaryd":  "alternative name",
			"millibaried": "alternative name",
			"millibarie":  "alternative name",
		},
		"", "",
	},
	"kilobarye": {
		0, 0, d * k,
		pressureFamily,
		"kBa", "kilobarye", "kilobaryes",
		"A unit of pressure in the" +
			" centimetre-gram-second (CGS) system of units.",
		[]Tag{TagMetric},
		map[string]string{
			"kBa":        "abbreviation",
			"kilobaryes": "plural",
			"kilobarad":  "alternative name",
			"kilobarrie": "alternative name",
			"kilobary":   "alternative name",
			"kilobaryd":  "alternative name",
			"kilobaried": "alternative name",
			"kilobarie":  "alternative name",
		},
		"", "",
	},
}
