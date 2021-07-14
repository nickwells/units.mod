package units

// UnitOfMass represents the base unit of mass
var UnitOfMass = Family{
	BaseUnitName: "gram",
	Description:  "unit of mass",
}

// GramUnit is a suitable default value for a UnitOfMass
var GramUnit = Unit{
	0, 0, 1,
	UnitOfMass,
	"g", UnitOfMass.BaseUnitName, "grams",
	"a metric measure of mass. Note that the SI unit is the kilogram.",
	"",
}

const (
	grainToGram = 0.06479891
	ounceToGram = grainToGram * 437.5
	poundToGram = ounceToGram * 16
)

// MassNames maps names to units of mass
//
// Note that the SI unit of weight is the kilogram not the gram. I have used
// the gram so as to avoid having to divide all the multiples by 1000 or else
// having all the multiples offset by one
var MassNames = map[string]Unit{
	// SI
	"yg": {
		0, 0, y,
		UnitOfMass,
		"yg", "yoctogram", "yoctograms", "a metric measure of mass.",
		"",
	},
	"zg": {
		0, 0, z,
		UnitOfMass,
		"zg", "zeptogram", "zeptograms", "a metric measure of mass.",
		"",
	},
	"ag": {
		0, 0, a,
		UnitOfMass,
		"ag", "attogram", "attograms", "a metric measure of mass.",
		"",
	},
	"fg": {
		0, 0, f,
		UnitOfMass,
		"fg", "femtogram", "femtograms", "a metric measure of mass.",
		"",
	},
	"pg": {
		0, 0, p,
		UnitOfMass,
		"pg", "picogram", "picograms", "a metric measure of mass.",
		"",
	},
	"ng": {
		0, 0, n,
		UnitOfMass,
		"ng", "nanogram", "nanograms", "a metric measure of mass.",
		"",
	},
	"ug": {
		0, 0, u,
		UnitOfMass,
		"µg", "microgram", "micrograms", "a metric measure of mass.",
		"",
	},
	"mg": {
		0, 0, m,
		UnitOfMass,
		"mg", "milligram", "milligrams", "a metric measure of mass.",
		"",
	},
	"cg": {
		0, 0, c,
		UnitOfMass,
		"cg", "centigram", "centigrams", "a metric measure of mass.",
		"",
	},
	"dg": {
		0, 0, d,
		UnitOfMass,
		"dg", "decigram", "decigrams", "a metric measure of mass.",
		"",
	},
	"gram": GramUnit,
	"dag": {
		0, 0, da,
		UnitOfMass,
		"dag", "decagram", "decagrams", "a metric measure of mass.",
		"",
	},
	"hg": {
		0, 0, h,
		UnitOfMass,
		"hg", "hectogram", "hectograms", "a metric measure of mass.",
		"",
	},
	"kg": {
		0, 0, k,
		UnitOfMass,
		"kg", "kilogram", "kilograms", "a metric measure of mass.",
		"",
	},
	"myg": {
		0, 0, 10000,
		UnitOfMass,
		"myg", "myriagram", "myriagrams",
		"an obsolete metric measure of mass.",
		"",
	},
	"tonne": {
		0, 0, _M,
		UnitOfMass,
		"T", "tonne", "tonnes",
		"a metric measure of mass." +
			" Also known (in the US as the 'metric ton').",
		"",
	},
	"kilotonne": {
		0, 0, _G,
		UnitOfMass,
		"KT", "kilotonne", "kilotonnes", "a metric measure of mass.",
		"",
	},
	"Tg": {
		0, 0, _T,
		UnitOfMass,
		"Tg", "teragram", "teragrams", "a metric measure of mass.",
		"",
	},
	"megatonne": {
		0, 0, _T,
		UnitOfMass,
		"MT", "megatonne", "megatonnes", "a metric measure of mass.",
		"",
	},
	"Pg": {
		0, 0, _P,
		UnitOfMass,
		"Pg", "petagram", "petagrams", "a metric measure of mass.",
		"",
	},
	"Eg": {
		0, 0, _E,
		UnitOfMass,
		"Eg", "exagram", "exagrams", "a metric measure of mass.",
		"",
	},
	"Zg": {
		0, 0, _Z,
		UnitOfMass,
		"Zg", "zettagram", "zettagrams", "a metric measure of mass.",
		"",
	},
	"Yg": {
		0, 0, _Y,
		UnitOfMass,
		"Yg", "yottagram", "yottagrams", "a metric measure of mass.",
		"",
	},

	// apothecaries weights
	"grain": {
		0, 0, grainToGram,
		UnitOfMass,
		"gr", "grain", "grains",
		"an imperial measure of mass (an apothecaries' measure).",
		"",
	},
	"troy-ounce": {
		0, 0, grainToGram * 480,
		UnitOfMass,
		"t oz", "troy ounce", "troy ounces",
		"an imperial measure of mass (an apothecaries' measure).",
		"",
	},
	"scruple": {
		0, 0, grainToGram * 20,
		UnitOfMass,
		"scruple", "scruple", "scruples",
		"an imperial measure of mass (an apothecaries' measure).",
		"",
	},
	"dram": {
		0, 0, grainToGram * 60,
		UnitOfMass,
		"dr", "dram", "drams",
		"an imperial measure of mass (an apothecaries' measure).",
		"",
	},
	"drachm": {
		0, 0, grainToGram * 60,
		UnitOfMass,
		"dr", "drachm", "drachms",
		"an imperial measure of mass (an apothecaries' measure).",
		"",
	},

	// Physics
	"electronvolt": {
		0, 0, electronVoltToJoule / (speedOfLight * speedOfLight),
		UnitOfMass,
		"eV", "electron volt", "electron volts",
		"derived from the unit of energy by" +
			" dividing by the speed of light squared (using e = mc^2).",
		"",
	},
	"gigaelectronvolt": {
		0, 0, _G * electronVoltToJoule / (speedOfLight * speedOfLight),
		UnitOfMass,
		"GeV", "giga electron volt", "giga electron volts",
		"see the description of the electron volt as a unit of mass." +
			" Note that, in general, the masses of all hadrons are" +
			" of the order of one giga electron volt which makes this" +
			" a convenient unit of mass for particle physics.",
		"",
	},

	// Imperial / US
	"ounce": {
		0, 0, ounceToGram,
		UnitOfMass,
		"oz", "ounce", "ounces",
		"an imperial measure of mass.",
		"",
	},
	"pound": {
		0, 0, poundToGram,
		UnitOfMass,
		"lb", "pound", "pounds",
		"an imperial measure of mass.",
		"",
	},
	"stone": {
		0, 0, poundToGram * 14,
		UnitOfMass,
		"st", "stone", "stones",
		"an imperial measure of mass.",
		"",
	},
	"hundredweight": {
		0, 0, poundToGram * 112,
		UnitOfMass,
		"cwt", "hundredweight", "hundredweight",
		"an imperial measure of mass.",
		"",
	},
	"short-hundredweight": {
		0, 0, poundToGram * 100, UnitOfMass,
		"cwt (short)", "short hundredweight", "short hundredweight",
		"a US customary measure of mass.",
		"",
	},
	"imperial-ton": {
		0, 0, poundToGram * 112 * 20,
		UnitOfMass,
		"t", "imperial ton", "imperial tons",
		"an imperial measure of mass." +
			" Also known as a 'long ton' or simply a 'ton'.",
		"",
	},
	"short-ton": {
		0, 0, poundToGram * 100 * 20,
		UnitOfMass,
		"short ton", "short ton", "short tons",
		"an imperial measure of mass.",
		"",
	},

	// astronomical
	"earth-mass": {
		0, 0, 5.9722e27,
		UnitOfMass,
		"M⊕", "Earth Mass", "Earth Masses",
		"an astronomical measure of mass, used to give a sense of scale.",
		"",
	},
	"solar-mass": {
		0, 0, 1.98847e33,
		UnitOfMass,
		"M⊙", "Solar Mass", "Solar Masses",
		"an astronomical measure of mass, used to give a sense of scale.",
		"",
	},
	"lunar-mass": {
		0, 0, 7.342e25,
		UnitOfMass,
		"ML", "Lunar Mass", "Lunar Masses",
		"an astronomical measure of mass, used to give a sense of scale.",
		"",
	},
}

var massAliases = map[string]Alias{
	"microgram":  {"ug", ""},
	"micrograms": {"ug", "plural"},

	"milligram":  {"mg", ""},
	"milligrams": {"mg", "plural"},

	"g":     {"gram", ""},
	"grams": {"gram", "plural"},

	"kilogram":  {"kg", ""},
	"kilograms": {"kg", "plural"},

	"Mg":         {"tonne", "megagram"},
	"tonnes":     {"tonne", "plural"},
	"metric ton": {"tonne", "US name"},

	"Gg":         {"kilotonne", "gigagram"},
	"kilotonnes": {"kilotonne", "plural"},

	"oz":     {"ounce", ""},
	"ounces": {"ounce", "plural"},

	"lb":     {"pound", ""},
	"pounds": {"pound", "plural"},

	"cwt":                    {"hundredweight", "abbreviation"},
	"imperial hundredweight": {"hundredweight", "US name"},
	"imperial-hundredweight": {"hundredweight", "US name, hyphenated"},
	"long hundredweight":     {"hundredweight", "US name"},
	"long-hundredweight":     {"hundredweight", "US name, hyphenated"},

	"cental": {"short-hundredweight", ""},

	"long ton": {"imperial-ton", "US name"},
	"long-ton": {"imperial-ton", "US name, hyphenated"},

	"eV":             {"electronvolt", "abbreviation"},
	"electron-volt":  {"electronvolt", ""},
	"electron-volts": {"electronvolt", "plural"},
	"electron volt":  {"electronvolt", ""},
	"electron volts": {"electronvolt", "plural"},

	"GeV":                 {"gigaelectronvolt", "abbreviation"},
	"giga-electron-volt":  {"gigaelectronvolt", ""},
	"giga-electron-volts": {"gigaelectronvolt", "plural"},
	"giga electron volt":  {"gigaelectronvolt", ""},
	"giga electron volts": {"gigaelectronvolt", "plural"},
}
