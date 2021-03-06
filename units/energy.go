package units

// UnitOfEnergy represents the base unit of energy
var UnitOfEnergy = Family{
	BaseUnitName: "joule",
	Description:  "unit of energy",
	Name:         Energy,
}

// JouleUnit is a suitable default value for a UnitOfEnergy
var JouleUnit = Unit{
	0, 0, 1,
	UnitOfEnergy,
	"J", UnitOfEnergy.BaseUnitName, "joules",
	"a metric measure of energy.",
	"",
}

const (
	calorieToJoule     = 4.184
	footPoundToJoule   = 1.3558179483314004
	footPoundalToJoule = 0.0421401100938048

	electronVoltToJoule = 1.602176634e-19

	// Note that there are many different variants of the British thermal
	// unit with conversion ratios to Joules varying between 1054.3503 (for
	// the Thermochemical variant) and 1059.67 (using the calorie value of
	// water at its maximum density, when it is at 3.9 ℃)
	BtuToJoule = 1055.06 // International Standard ISO 31-4
)

// EnergyNames maps names to units of energy
var EnergyNames = map[string]Unit{
	// SI
	"yJ": {
		0, 0, y,
		UnitOfEnergy,
		"yJ", "yoctojoule", "yoctojoules",
		"a metric measure of energy.",
		"",
	},
	"zJ": {
		0, 0, z,
		UnitOfEnergy,
		"zJ", "zeptojoule", "zeptojoules",
		"a metric measure of energy.",
		"",
	},
	"aJ": {
		0, 0, a,
		UnitOfEnergy,
		"aJ", "attojoule", "attojoules",
		"a metric measure of energy.",
		"",
	},
	"fJ": {
		0, 0, f,
		UnitOfEnergy,
		"fJ", "femtojoule", "femtojoules",
		"a metric measure of energy.",
		"",
	},
	"pJ": {
		0, 0, p,
		UnitOfEnergy,
		"pJ", "picojoule", "picojoules",
		"a metric measure of energy.",
		"",
	},
	"nJ": {
		0, 0, n,
		UnitOfEnergy,
		"nJ", "nanojoule", "nanojoules",
		"a metric measure of energy.",
		"",
	},
	"uJ": {
		0, 0, u,
		UnitOfEnergy,
		"uJ", "microjoule", "microjoules",
		"a metric measure of energy.",
		"",
	},
	"mJ": {
		0, 0, m,
		UnitOfEnergy,
		"mJ", "millijoule", "millijoules",
		"a metric measure of energy.",
		"",
	},
	"cJ": {
		0, 0, c,
		UnitOfEnergy,
		"cJ", "centijoule", "centijoules",
		"a metric measure of energy.",
		"",
	},
	"dJ": {
		0, 0, d,
		UnitOfEnergy,
		"dJ", "decijoule", "decijoules",
		"a metric measure of energy.",
		"",
	},
	"joule": JouleUnit,
	"daJ": {
		0, 0, da,
		UnitOfEnergy,
		"daJ", "decajoule", "decajoules",
		"a metric measure of energy.",
		"",
	},
	"hJ": {
		0, 0, h,
		UnitOfEnergy,
		"hJ", "hectojoule", "hectojoules",
		"a metric measure of energy.",
		"",
	},
	"kJ": {
		0, 0, k,
		UnitOfEnergy,
		"kJ", "kilojoule", "kilojoules",
		"a metric measure of energy.",
		"",
	},
	"MJ": {
		0, 0, _M,
		UnitOfEnergy,
		"MJ", "megajoule", "megajoules",
		"a metric measure of energy.",
		"",
	},
	"GJ": {
		0, 0, _G,
		UnitOfEnergy,
		"GJ", "gigajoule", "gigajoules",
		"a metric measure of energy.",
		"",
	},
	"TJ": {
		0, 0, _T,
		UnitOfEnergy,
		"TJ", "terajoule", "terajoules",
		"a metric measure of energy.",
		"",
	},
	"PJ": {
		0, 0, _P,
		UnitOfEnergy,
		"PJ", "petajoule", "petajoules",
		"a metric measure of energy.",
		"",
	},
	"EJ": {
		0, 0, _E,
		UnitOfEnergy,
		"EJ", "exajoule", "exajoules",
		"a metric measure of energy.",
		"",
	},
	"ZJ": {
		0, 0, _Z,
		UnitOfEnergy,
		"ZJ", "zettajoule", "zettajoules",
		"a metric measure of energy.",
		"",
	},
	"YJ": {
		0, 0, _Y,
		UnitOfEnergy,
		"YJ", "yottajoule", "yottajoules",
		"a metric measure of energy.",
		"",
	},

	"kWh": {
		0, 0, 3.6e6,
		UnitOfEnergy,
		"kW h", "kilowatt-hour", "killowatt-hours",
		"a measure of energy commonly used as a billing unit" +
			" for energy delivered to consumers by electric utilities.",
		"",
	},

	"erg": {
		0, 0, 1e-7,
		UnitOfEnergy,
		"erg", "erg", "ergs",
		"a measure of energy proposed by" +
			" the German physicist and mathematician Rudolf Clausius.",
		"",
	},
	"foe": {
		0, 0, 1e44,
		UnitOfEnergy,
		"foe", "foe", "foes",
		"a measure of energy equivalent to 10^51 ergs," +
			" the approximate energy released by a supernova." +
			" It was introduced by Gerald E. Brown.",
		"",
	},

	"cal": {
		0, 0, calorieToJoule,
		UnitOfEnergy,
		"cal", "calorie", "calories",
		"a measure of energy (deprecated - use kcal).",
		"",
	},
	"kcal": {
		0, 0, calorieToJoule * 1000,
		UnitOfEnergy,
		"kcal", "kilocalorie", "kilocalories",
		"a measure of energy introduced by Nicolas Clément." +
			" It is widely used in nutrition." +
			" It is the amount of heat energy required to" +
			" raise the temperature of a kilogram of water" +
			" by one degree Celsius." +
			" It is sometimes known as the large calorie.",
		"",
	},

	// Physics
	"electronvolt": {
		0, 0, electronVoltToJoule,
		UnitOfEnergy,
		"eV", "electron volt", "electron volts",
		"The kinetic energy gained by a single electron" +
			" accelerating from rest through" +
			" an electric potential difference of one volt.",
		"",
	},

	// Imperial / US
	"foot-pound": {
		0, 0, footPoundToJoule,
		UnitOfEnergy,
		"ft lb", "foot-pound", "foot-pounds",
		"an imperial measure of energy.",
		"",
	},
	"foot-poundal": {
		0, 0, footPoundalToJoule,
		UnitOfEnergy,
		"ft pdl", "foot-poundal", "foot-poundals",
		"an imperial measure of energy.",
		"",
	},
	"BTU": {
		0, 0, BtuToJoule,
		UnitOfEnergy,
		"Btu", "British thermal unit", "British thermal units",
		"an imperial measure of energy." +
			" This uses the ISO 31-4 conversion to Joules",
		"",
	},
	"therm": {
		0, 0, BtuToJoule * 1e5,
		UnitOfEnergy,
		"therm", "therm", "therms",
		"an imperial measure of energy.",
		"",
	},

	// Coloquial
	"tonOfTNT": {
		0, 0, calorieToJoule * 1e9,
		UnitOfEnergy,
		"ton of TNT", "ton of TNT", "tons of TNT",
		"a coloquial measure of energy typically used to describe" +
			" the yield of atomic weapons",
		"",
	},
}

var energyAliases = map[string]Alias{
	"calorie":       {"cal", ""},
	"calories":      {"cal", "plural"},
	"kilocalorie":   {"kcal", ""},
	"kilocalories":  {"kcal", "plural"},
	"kilocal":       {"kcal", ""},
	"large calorie": {"kcal", ""},

	"Btu": {"BTU", ""},

	"bethe": {
		"foe",
		"An alternative proposed by Steven Weinberg, after Hans Bethe",
	},

	"J":      {"joule", "abbreviation"},
	"joules": {"joule", "plural"},
	"Joule":  {"joule", "with initial capital"},
	"Joules": {"joule", "with initial capital, plural"},

	"eV":             {"electronvolt", "abbreviation"},
	"electron-volt":  {"electronvolt", ""},
	"electron-volts": {"electronvolt", "plural"},
	"electron volt":  {"electronvolt", ""},
	"electron volts": {"electronvolt", "plural"},
}
