package units

// bunNumeric is the base unit name for dimensionless units
const bunNumeric = "1"

// numericFamily represents a dimensionless value
var numericFamily = &Family{
	baseUnitName: bunNumeric,
	description:  "dimensionless value",
	name:         Dimensionless,
}

// DimensionlessNames maps names to numeric (dimensionless) units
var dimensionlessNames = map[string]Unit{
	bunNumeric: {
		0, 0, 1,
		numericFamily,
		"", bunNumeric, bunNumeric,
		"", []Tag{TagDimensionless},
		map[string]string{
			"one":  "",
			"unit": "",
		},
		"", "",
	},
	"y": {
		0, 0, y,
		numericFamily,
		"y", "yocto", "yocto", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"z": {
		0, 0, z,
		numericFamily,
		"z", "zepto", "zepto", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"a": {
		0, 0, a,
		numericFamily,
		"a", "atto", "atto", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"f": {
		0, 0, f,
		numericFamily,
		"f", "femto", "femto", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"p": {
		0, 0, p,
		numericFamily,
		"p", "pico", "pico", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"n": {
		0, 0, n,
		numericFamily,
		"n", "nano", "nano", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"u": {
		0, 0, u,
		numericFamily,
		"u", "micro", "micro", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"m": {
		0, 0, m,
		numericFamily,
		"m", "milli", "milli", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"c": {
		0, 0, c,
		numericFamily,
		"c", "centi", "centi", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"d": {
		0, 0, d,
		numericFamily,
		"d", "deci", "deci", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"da": {
		0, 0, da,
		numericFamily,
		"da", "deca", "deca", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"dozen": {
		0, 0, 12,
		numericFamily,
		"doz", "dozen", "dozen", "",
		[]Tag{TagDimensionless},
		map[string]string{
			"dozens": "plural",
		},
		"", "",
	},
	"bakers dozen": {
		0, 0, 13,
		numericFamily,
		"doz", "dozen", "dozen",
		"The name may have originated from a practice of" +
			" bakers adding an extra item to an order of 12" +
			" to ensure that the items sold met a standard weight." +
			" This was to avoid a fine",
		[]Tag{TagDimensionless, TagColloquial},
		map[string]string{
			"devils dozen":  "",
			"baker's dozen": "with the correct apostrophe",
			"devil's dozen": "with the correct apostrophe",
			"long dozen":    "",
			"big dozen":     "",
		},
		"", "",
	},
	"h": {
		0, 0, h,
		numericFamily,
		"h", "hecto", "hecto", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"k": {
		0, 0, k,
		numericFamily,
		"k", "kilo", "kilo", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"M": {
		0, 0, _M,
		numericFamily,
		"M", "mega", "mega", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"G": {
		0, 0, _G,
		numericFamily,
		"G", "giga", "giga", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"T": {
		0, 0, _T,
		numericFamily,
		"T", "tera", "tera", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"P": {
		0, 0, _P,
		numericFamily,
		"P", "peta", "peta", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"E": {
		0, 0, _E,
		numericFamily,
		"E", "exa", "exa", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"Z": {
		0, 0, _Z,
		numericFamily,
		"Z", "zetta", "zetta", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"Y": {
		0, 0, _Y,
		numericFamily,
		"Y", "yotta", "yotta", "", []Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},

	"myriad": {
		0, 0, 10000,
		numericFamily,
		"myriad", "myriad", "myriads",
		"historically, ten thousand but latterly meaning" +
			" a countless number of things",
		[]Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},

	"million": {
		0, 0, _M,
		numericFamily,
		"M", "million", "million", "",
		[]Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"billion": {
		0, 0, _G,
		numericFamily,
		"B", "billion", "billion",
		"Note that this reflects the, now universal, meaning of" +
			" 1000 million." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million squared which was the" +
			" original usage in England. If the text also uses" +
			" milliard then this is most likely to mean a million" +
			" million (what is now generally referred to as a trillion).",
		[]Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"trillion": {
		0, 0, _T,
		numericFamily,
		"Tr", "trillion", "trillion",
		"Note that this reflects the, now universal, meaning of" +
			" 1000 billion." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million cubed which was the" +
			" original usage in England.",
		[]Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"quadrillion": {
		0, 0, _P,
		numericFamily,
		"Qu", "quadrillion", "quadrillion",
		"Note that this reflects the, now universal, meaning of" +
			" 1000 trillion." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million to the power of four" +
			" which was the original usage in England.",
		[]Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"milliard": {
		0, 0, 1e6 * 1000,
		numericFamily,
		"milliard", "milliard", "milliard",
		"an English (UK) term for a thousand million," +
			" now obsolete, it has been superseded by" +
			" the US meaning of billion.",
		[]Tag{TagDimensionless, TagHist},
		map[string]string{
			"yard": "milliard",
		},
		"", "",
	},
	"billion (UK)": {
		0, 0, 1e12,
		numericFamily,
		"B (UK)", "billion (UK)", "billion (UK)",
		"Note that this reflects the, now obsolete, meaning of" +
			" a million squared." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million squared which was the" +
			" original usage in England. You will need to judge by" +
			" context whether the original or inflated meaning is" +
			" intended.",
		[]Tag{TagDimensionless, TagHist},
		map[string]string{},
		"", "",
	},
	"billiard": {
		0, 0, 1e12 * 1000,
		numericFamily,
		"billiard", "billiard", "billiard",
		"an English (UK) term for a thousand (UK) billion," +
			" now obsolete, it has been superseded by" +
			" the US meaning of quadrillion.",
		[]Tag{TagDimensionless, TagHist},
		map[string]string{},
		"", "",
	},
	"trillion (UK)": {
		0, 0, 1e18,
		numericFamily,
		"Tr (UK)", "trillion (UK)", "trillion (UK)",
		"Note that this reflects the, now obsolete, meaning of" +
			" a million cubed." +
			" Be aware that some older texts, especially in the UK" +
			" might use trillion to mean a million cubed which was the" +
			" original usage in the UK. You will need to judge by" +
			" context whether the original or inflated meaning is intended.",
		[]Tag{TagDimensionless, TagHist},
		map[string]string{},
		"", "",
	},
	"trilliard": {
		0, 0, 1e18 * 1000,
		numericFamily,
		"trilliard", "trilliard", "trilliard",
		"an English (UK) term for a thousand (UK) trillion," +
			" now obsolete.",
		[]Tag{TagDimensionless, TagHist},
		map[string]string{},
		"", "",
	},

	"lakh": {
		0, 0, 1e5,
		numericFamily,
		"lakh", "lakh", "lakh",
		"Indian: 1,00,000.",
		[]Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},
	"crore": {
		0, 0, 1e7,
		numericFamily,
		"crore", "crore", "crore",
		"Indian: 1,00,00,000.",
		[]Tag{TagDimensionless},
		map[string]string{},
		"", "",
	},

	"pony": {
		0, 0, 25,
		numericFamily,
		"pony", "pony", "ponies",
		"UK slang: £25.",
		[]Tag{TagDimensionless, TagColloquial},
		map[string]string{},
		"", "",
	},
	"monkey": {
		0, 0, 500,
		numericFamily,
		"monkey", "monkey", "monkeys",
		"UK slang: £500.",
		[]Tag{TagDimensionless, TagColloquial},
		map[string]string{},
		"", "",
	},
	"grand": {
		0, 0, k,
		numericFamily,
		"grand", "grand", "grand",
		"UK slang: £1000.",
		[]Tag{TagDimensionless, TagColloquial},
		map[string]string{},
		"", "",
	},
}
