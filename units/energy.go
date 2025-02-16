//nolint:mnd
package units

const bunEnergy = "joule"

// energyFamily represents the base unit of energy
var energyFamily = &Family{
	baseUnitName: bunEnergy,
	description:  "unit of energy",
	name:         Energy,
}

// EnergyNames maps names to units of energy
var energyNames = map[string]Unit{
	bunEnergy: {
		0, 0, 1,
		energyFamily,
		"J", bunEnergy, "joules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"J":      "abbreviation",
			"joules": "plural",
			"Joule":  "with initial capital",
			"Joules": "with initial capital, plural",
		},
		"", "",
	},

	// SI
	"yJ": {
		0, 0, y,
		energyFamily,
		"yJ", "yoctojoule", "yoctojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"zJ": {
		0, 0, z,
		energyFamily,
		"zJ", "zeptojoule", "zeptojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"aJ": {
		0, 0, a,
		energyFamily,
		"aJ", "attojoule", "attojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"fJ": {
		0, 0, f,
		energyFamily,
		"fJ", "femtojoule", "femtojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"pJ": {
		0, 0, p,
		energyFamily,
		"pJ", "picojoule", "picojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"nJ": {
		0, 0, n,
		energyFamily,
		"nJ", "nanojoule", "nanojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"uJ": {
		0, 0, u,
		energyFamily,
		"uJ", "microjoule", "microjoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"mJ": {
		0, 0, m,
		energyFamily,
		"mJ", "millijoule", "millijoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"cJ": {
		0, 0, c,
		energyFamily,
		"cJ", "centijoule", "centijoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"dJ": {
		0, 0, d,
		energyFamily,
		"dJ", "decijoule", "decijoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"daJ": {
		0, 0, da,
		energyFamily,
		"daJ", "decajoule", "decajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"hJ": {
		0, 0, h,
		energyFamily,
		"hJ", "hectojoule", "hectojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"kJ": {
		0, 0, k,
		energyFamily,
		"kJ", "kilojoule", "kilojoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"MJ": {
		0, 0, _M,
		energyFamily,
		"MJ", "megajoule", "megajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"GJ": {
		0, 0, _G,
		energyFamily,
		"GJ", "gigajoule", "gigajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"TJ": {
		0, 0, _T,
		energyFamily,
		"TJ", "terajoule", "terajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"PJ": {
		0, 0, _P,
		energyFamily,
		"PJ", "petajoule", "petajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"EJ": {
		0, 0, _E,
		energyFamily,
		"EJ", "exajoule", "exajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"ZJ": {
		0, 0, _Z,
		energyFamily,
		"ZJ", "zettajoule", "zettajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"YJ": {
		0, 0, _Y,
		energyFamily,
		"YJ", "yottajoule", "yottajoules",
		"a metric measure of energy.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},

	"kWh": {
		0, 0, 3.6e6,
		energyFamily,
		"kW h", "kilowatt-hour", "killowatt-hours",
		"a measure of energy commonly used as a billing unit" +
			" for energy delivered to consumers by electric utilities.",
		[]Tag{TagMetric},
		map[string]string{},
		"", "",
	},

	"erg": {
		0, 0, 1e-7,
		energyFamily,
		"erg", "erg", "ergs",
		"a measure of energy proposed by" +
			" the German physicist and mathematician Rudolf Clausius.",
		[]Tag{TagMetric},
		map[string]string{},
		"", "",
	},
	"foe": {
		0, 0, 1e44,
		energyFamily,
		"foe", "foe", "foes",
		"a measure of energy equivalent to 10^51 ergs," +
			" the approximate energy released by a supernova." +
			" It was introduced by Gerald E. Brown.",
		[]Tag{TagColloquial},
		map[string]string{
			"bethe": "An alternative proposed by Steven Weinberg," +
				" after Hans Bethe",
		},
		"", "",
	},

	"cal": {
		0, 0, calorieToJoule,
		energyFamily,
		"cal", "calorie", "calories",
		"a measure of energy (deprecated in 1948 - use kcal)." +
			" It was introduced by Pierre Antoine Favre" +
			" and Johann T. Silbermann in 1852",
		[]Tag{TagMetric, TagHist},
		map[string]string{
			"calorie":       "",
			"calories":      "plural",
			"small calorie": "",
		},
		"", "",
	},
	"kcal": {
		0, 0, calorieToJoule * 1000,
		energyFamily,
		"kcal", "kilocalorie", "kilocalories",
		"a measure of energy introduced by Nicolas Clément." +
			" It is widely used in nutrition." +
			" It is the amount of heat energy required to" +
			" raise the temperature of a kilogram of water" +
			" by one degree Celsius." +
			" This is dependent on atmopspheric pressure and the" +
			" starting temperature, The value used is that for" +
			" the Thermochemical calorie," +
			"\n\n" +
			" It is sometimes known as the large calorie." +
			" It was introduced by Nicolas Clément",
		[]Tag{TagMetric},
		map[string]string{
			"kilocalorie":   "full name",
			"kilocalories":  "full name, plural",
			"kilocal":       "",
			"large calorie": "",
		},
		"", "",
	},

	// Physics
	"electronvolt": {
		0, 0, electronVoltToJoule,
		energyFamily,
		"eV", "electron volt", "electron volts",
		"The kinetic energy gained by a single electron" +
			" accelerating from rest through" +
			" an electric potential difference of one volt.",
		[]Tag{TagPhysics},
		map[string]string{
			"eV":             "abbreviation",
			"electron-volt":  "",
			"electron-volts": "plural",
			"electron volt":  "",
			"electron volts": "plural",
		},
		"", "",
	},

	// Imperial / US
	"foot-pound": {
		0, 0, footPoundToJoule,
		energyFamily,
		"ft lb", "foot-pound", "foot-pounds",
		"an imperial measure of energy.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{},
		"", "",
	},
	"foot-poundal": {
		0, 0, footPoundalToJoule,
		energyFamily,
		"ft pdl", "foot-poundal", "foot-poundals",
		"an imperial measure of energy.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{},
		"", "",
	},
	"BTU": {
		0, 0, btuToJoule,
		energyFamily,
		"Btu", "British thermal unit", "British thermal units",
		"an imperial measure of energy." +
			" This uses the ISO 31-4 conversion to Joules",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"Btu": "",
		},
		"", "",
	},
	"therm": {
		0, 0, btuToJoule * 1e5,
		energyFamily,
		"therm", "therm", "therms",
		"an imperial measure of energy.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{},
		"", "",
	},

	// Coloquial
	"tonOfTNT": {
		0, 0, calorieToJoule * 1e9,
		energyFamily,
		"ton of TNT", "ton of TNT", "tons of TNT",
		"a coloquial measure of energy typically used to describe" +
			" the yield of atomic weapons",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},
}
