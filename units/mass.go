package units

// bunMass is the base unit name for mass
const bunMass = "gram"

const daltonToGram = 1.660_539_068_92e-24

// massFamily represents the collection of units of mass
var massFamily = &Family{
	baseUnitName: bunMass,
	description:  "unit of mass",
	name:         Mass,
}

// MassNames maps names to units of mass
//
// Note that the SI unit of weight is the kilogram not the gram. I have used
// the gram so as to avoid having to divide all the multiples by 1000 or else
// having all the multiples offset by one
var massNames = map[string]Unit{
	bunMass: {
		0, 0, 1,
		massFamily,
		"g", bunMass, "grams",
		"a metric measure of mass. Note that the SI unit is the kilogram.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"g":     "",
			"grams": "plural",
		},
		"", "",
	},
	// SI
	"yg": {
		0, 0, y,
		massFamily,
		"yg", "yoctogram", "yoctograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"zg": {
		0, 0, z,
		massFamily,
		"zg", "zeptogram", "zeptograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"ag": {
		0, 0, a,
		massFamily,
		"ag", "attogram", "attograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"fg": {
		0, 0, f,
		massFamily,
		"fg", "femtogram", "femtograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"pg": {
		0, 0, p,
		massFamily,
		"pg", "picogram", "picograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"ng": {
		0, 0, n,
		massFamily,
		"ng", "nanogram", "nanograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"ug": {
		0, 0, u,
		massFamily,
		"µg", "microgram", "micrograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"microgram":  "",
			"micrograms": "plural",
		},
		"", "",
	},
	"mg": {
		0, 0, m,
		massFamily,
		"mg", "milligram", "milligrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"milligram":  "",
			"milligrams": "plural",
		},
		"", "",
	},
	"cg": {
		0, 0, c,
		massFamily,
		"cg", "centigram", "centigrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"dg": {
		0, 0, d,
		massFamily,
		"dg", "decigram", "decigrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"dag": {
		0, 0, da,
		massFamily,
		"dag", "decagram", "decagrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"hg": {
		0, 0, h,
		massFamily,
		"hg", "hectogram", "hectograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"kg": {
		0, 0, k,
		massFamily,
		"kg", "kilogram", "kilograms", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"kilogram":  "",
			"kilograms": "plural",
		},
		"", "",
	},
	"myg": {
		0, 0, 10000,
		massFamily,
		"myg", "myriagram", "myriagrams",
		"an obsolete metric measure of mass.",
		[]Tag{TagHist, TagMetric},
		map[string]string{},
		"", "",
	},
	"tonne": {
		0, 0, _M,
		massFamily,
		"T", "tonne", "tonnes",
		"a metric measure of mass." +
			" Also known (in the US as the 'metric ton').",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"Mg":         "megagram",
			"tonnes":     "plural",
			"metric ton": "US name",
		},
		"", "",
	},
	"kilotonne": {
		0, 0, _G,
		massFamily,
		"KT", "kilotonne", "kilotonnes", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"Gg":         "gigagram",
			"kilotonnes": "plural",
		},
		"", "",
	},
	"Tg": {
		0, 0, _T,
		massFamily,
		"Tg", "teragram", "teragrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"megatonne": {
		0, 0, _T,
		massFamily,
		"MT", "megatonne", "megatonnes", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Pg": {
		0, 0, _P,
		massFamily,
		"Pg", "petagram", "petagrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Eg": {
		0, 0, _E,
		massFamily,
		"Eg", "exagram", "exagrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Zg": {
		0, 0, _Z,
		massFamily,
		"Zg", "zettagram", "zettagrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Yg": {
		0, 0, _Y,
		massFamily,
		"Yg", "yottagram", "yottagrams", "a metric measure of mass.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},

	// apothecaries weights
	"grain": {
		0, 0, grainToGram,
		massFamily,
		"gr", "grain", "grains",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},
	"troy-ounce": {
		0, 0, grainToGram * 480,
		massFamily,
		"t oz", "troy ounce", "troy ounces",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},
	"scruple": {
		0, 0, grainToGram * 20,
		massFamily,
		"scruple", "scruple", "scruples",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},
	"dram": {
		0, 0, grainToGram * 60,
		massFamily,
		"dr", "dram", "drams",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},
	"drachm": {
		0, 0, grainToGram * 60,
		massFamily,
		"dr", "drachm", "drachms",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},

	// Physics
	"electronvolt": {
		0, 0, electronVoltToJoule / (speedOfLight * speedOfLight),
		massFamily,
		"eV", "electron volt", "electron volts",
		"derived from the unit of energy by" +
			" dividing by the speed of light squared (using e = mc^2).",
		[]Tag{TagMetric, TagPhysics},
		map[string]string{
			"eV":             "abbreviation",
			"electron-volt":  "hyphenated",
			"electron-volts": "hyphenated, plural",
			"electron volt":  "",
			"electron volts": "plural",
		},
		"", "",
	},
	"gigaelectronvolt": {
		0, 0, _G * electronVoltToJoule / (speedOfLight * speedOfLight),
		massFamily,
		"GeV", "giga electron volt", "giga electron volts",
		"see the description of the electron volt as a unit of mass." +
			" Note that, in general, the masses of all hadrons are" +
			" of the order of one giga electron volt which makes this" +
			" a convenient unit of mass for particle physics.",
		[]Tag{TagMetric, TagPhysics},
		map[string]string{
			"GeV":                 "abbreviation",
			"giga-electron-volt":  "hyphenated",
			"giga-electron-volts": "hyphenated, plural",
			"giga electron volt":  "",
			"giga electron volts": "plural",
		},
		"", "",
	},
	"dalton": {
		0, 0, daltonToGram,
		massFamily,
		"Da", "dalton", "daltons",
		"One twelfth of the mass of an unbound neutral atom of" +
			" carbon-12 in its nuclear and electronic ground" +
			" state and at rest." +
			" It is commonly used in physics and chemistry to" +
			" express the mass of atomic scale objects such as" +
			" atoms, molecules and elementary particles." +
			" The mass in daltons of an atom is numerically close" +
			" but not exactly equal to the number of nucleons in its nucleus." +
			" It is named for John Dalton (1766-1844), an English chemist," +
			" physicist and meteorologist." +
			" It is a non-SI unit accepted for use with SI.",
		[]Tag{TagMetric},
		map[string]string{
			"Da":                       "abbreviation",
			"unified atomic mass unit": "",
			"daltons":                  "plural",
		},
		"", "",
	},
	"kilodalton": {
		0, 0, k * daltonToGram,
		massFamily,
		"kDa", "kilodalton", "kilodaltons",
		"see the description of the dalton as a unit of mass.",
		[]Tag{TagMetric},
		map[string]string{
			"kDa":         "abbreviation",
			"kilodaltons": "plural",
		},
		"", "",
	},
	"megadalton": {
		0, 0, _M * daltonToGram,
		massFamily,
		"MDa", "megadalton", "megadaltons",
		"see the description of the dalton as a unit of mass.",
		[]Tag{TagMetric},
		map[string]string{
			"MDa":         "abbreviation",
			"megadaltons": "plural",
		},
		"", "",
	},

	// Imperial / US
	"ounce": {
		0, 0, ounceToGram,
		massFamily,
		"oz", "ounce", "ounces",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"oz":     "",
			"ounces": "plural",
		},
		"", "",
	},
	"pound": {
		0, 0, poundToGram,
		massFamily,
		"lb", "pound", "pounds",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"lb":     "",
			"pounds": "plural",
		},
		"", "",
	},
	"stone": {
		0, 0, poundToGram * 14,
		massFamily,
		"st", "stone", "stones",
		"an imperial measure of mass.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{},
		"", "",
	},
	"hundredweight": {
		0, 0, poundToGram * 112,
		massFamily,
		"cwt", "hundredweight", "hundredweight",
		"an imperial measure of mass.",
		[]Tag{TagImperial},
		map[string]string{
			"cwt":                    "abbreviation",
			"imperial hundredweight": "US name",
			"imperial-hundredweight": "US name, hyphenated",
			"long hundredweight":     "US name",
			"long-hundredweight":     "US name, hyphenated",
		},
		"", "",
	},
	"short-hundredweight": {
		0, 0, poundToGram * 100, massFamily,
		"cwt (short)", "short hundredweight", "short hundredweight",
		"a US customary measure of mass.",
		[]Tag{TagUScustomary},
		map[string]string{
			"cental": "",
		},
		"", "",
	},
	"imperial-ton": {
		0, 0, poundToGram * 112 * 20,
		massFamily,
		"t", "imperial ton", "imperial tons",
		"an imperial measure of mass." +
			" Also known as a 'long ton' or simply a 'ton'.",
		[]Tag{TagImperial},
		map[string]string{
			"long ton": "US name",
			"long-ton": "US name, hyphenated",
		},
		"", "",
	},
	"short-ton": {
		0, 0, poundToGram * 100 * 20,
		massFamily,
		"short ton", "short ton", "short tons",
		"an imperial measure of mass.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},

	// astronomical
	"earth-mass": {
		0, 0, 5.9722e27,
		massFamily,
		"M⊕", "Earth Mass", "Earth Masses",
		"an astronomical measure of mass, used to give a sense of scale.",
		[]Tag{TagAstro},
		map[string]string{},
		"", "",
	},
	"solar-mass": {
		0, 0, 1.98847e33,
		massFamily,
		"M⊙", "Solar Mass", "Solar Masses",
		"an astronomical measure of mass, used to give a sense of scale.",
		[]Tag{TagAstro},
		map[string]string{},
		"", "",
	},
	"lunar-mass": {
		0, 0, 7.342e25,
		massFamily,
		"ML", "Lunar Mass", "Lunar Masses",
		"an astronomical measure of mass, used to give a sense of scale.",
		[]Tag{TagAstro},
		map[string]string{},
		"", "",
	},
}
