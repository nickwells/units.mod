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
		"", nil,
		map[string]string{},
		"", "",
	},
	"y": {
		0, 0, y,
		numericFamily,
		"y", "yocto", "yocto", "", nil,
		map[string]string{},
		"", "",
	},
	"z": {
		0, 0, z,
		numericFamily,
		"z", "zepto", "zepto", "", nil,
		map[string]string{},
		"", "",
	},
	"a": {
		0, 0, a,
		numericFamily,
		"a", "atto", "atto", "", nil,
		map[string]string{},
		"", "",
	},
	"f": {
		0, 0, f,
		numericFamily,
		"f", "femto", "femto", "", nil,
		map[string]string{},
		"", "",
	},
	"p": {
		0, 0, p,
		numericFamily,
		"p", "pico", "pico", "", nil,
		map[string]string{},
		"", "",
	},
	"n": {
		0, 0, n,
		numericFamily,
		"n", "nano", "nano", "", nil,
		map[string]string{},
		"", "",
	},
	"u": {
		0, 0, u,
		numericFamily,
		"u", "micro", "micro", "", nil,
		map[string]string{},
		"", "",
	},
	"m": {
		0, 0, m,
		numericFamily,
		"m", "milli", "milli", "", nil,
		map[string]string{},
		"", "",
	},
	"c": {
		0, 0, c,
		numericFamily,
		"c", "centi", "centi", "", nil,
		map[string]string{},
		"", "",
	},
	"d": {
		0, 0, d,
		numericFamily,
		"d", "deci", "deci", "", nil,
		map[string]string{},
		"", "",
	},
	"da": {
		0, 0, da,
		numericFamily,
		"da", "deca", "deca", "", nil,
		map[string]string{},
		"", "",
	},
	"dozen": {
		0, 0, 12,
		numericFamily,
		"doz", "dozen", "dozen", "", nil,
		map[string]string{},
		"", "",
	},
	"h": {
		0, 0, h,
		numericFamily,
		"h", "hecto", "hecto", "", nil,
		map[string]string{},
		"", "",
	},
	"k": {
		0, 0, k,
		numericFamily,
		"k", "kilo", "kilo", "", nil,
		map[string]string{},
		"", "",
	},
	"M": {
		0, 0, _M,
		numericFamily,
		"M", "mega", "mega", "", nil,
		map[string]string{},
		"", "",
	},
	"G": {
		0, 0, _G,
		numericFamily,
		"G", "giga", "giga", "", nil,
		map[string]string{},
		"", "",
	},
	"T": {
		0, 0, _T,
		numericFamily,
		"T", "tera", "tera", "", nil,
		map[string]string{},
		"", "",
	},
	"P": {
		0, 0, _P,
		numericFamily,
		"P", "peta", "peta", "", nil,
		map[string]string{},
		"", "",
	},
	"E": {
		0, 0, _E,
		numericFamily,
		"E", "exa", "exa", "", nil,
		map[string]string{},
		"", "",
	},
	"Z": {
		0, 0, _Z,
		numericFamily,
		"Z", "zetta", "zetta", "", nil,
		map[string]string{},
		"", "",
	},
	"Y": {
		0, 0, _Y,
		numericFamily,
		"Y", "yotta", "yotta", "", nil,
		map[string]string{},
		"", "",
	},

	"myriad": {
		0, 0, 10000,
		numericFamily,
		"myriad", "myriad", "myriads",
		"historically, ten thousand but latterly meaning" +
			" a countless number of things",
		nil,
		map[string]string{},
		"", "",
	},

	"million": {
		0, 0, _M,
		numericFamily,
		"M", "million", "million", "",
		nil,
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
		nil,
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
		nil,
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
		nil,
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
		nil,
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
		nil,
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
		nil,
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
		nil,
		map[string]string{},
		"", "",
	},
	"trilliard": {
		0, 0, 1e18 * 1000,
		numericFamily,
		"trilliard", "trilliard", "trilliard",
		"an English (UK) term for a thousand (UK) trillion," +
			" now obsolete.",
		nil,
		map[string]string{},
		"", "",
	},

	"lakh": {
		0, 0, 1e5,
		numericFamily,
		"lakh", "lakh", "lakh",
		"Indian: 1,00,000.",
		nil,
		map[string]string{},
		"", "",
	},
	"crore": {
		0, 0, 1e7,
		numericFamily,
		"crore", "crore", "crore",
		"Indian: 1,00,00,000.",
		nil,
		map[string]string{},
		"", "",
	},

	"pony": {
		0, 0, 25,
		numericFamily,
		"pony", "pony", "ponies",
		"UK slang: £25.",
		nil,
		map[string]string{},
		"", "",
	},
	"monkey": {
		0, 0, 500,
		numericFamily,
		"monkey", "monkey", "monkeys",
		"UK slang: £500.",
		nil,
		map[string]string{},
		"", "",
	},
	"grand": {
		0, 0, k,
		numericFamily,
		"grand", "grand", "grand",
		"UK slang: £1000.",
		nil,
		map[string]string{},
		"", "",
	},
}
