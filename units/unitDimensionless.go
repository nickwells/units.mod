package units

// Numeric represents a dimensionless value
var Numeric = Family{
	BaseUnitName: "",
	Description:  "dimensionless value",
}

// DimensionlessUnit is a suitable default value for a Unit of Numeric
var DimensionlessUnit = Unit{
	0, 0, 1,
	Numeric, "", Numeric.BaseUnitName, Numeric.BaseUnitName, "", "",
}

// DimensionlessNames maps names to numeric (dimensionless) units
var DimensionlessNames = map[string]Unit{
	"y":     {0, 0, y, Numeric, "y", "yocto", "yocto", "", ""},
	"z":     {0, 0, z, Numeric, "z", "zepto", "zepto", "", ""},
	"a":     {0, 0, a, Numeric, "a", "atto", "atto", "", ""},
	"f":     {0, 0, f, Numeric, "f", "femto", "femto", "", ""},
	"p":     {0, 0, p, Numeric, "p", "pico", "pico", "", ""},
	"n":     {0, 0, n, Numeric, "n", "nano", "nano", "", ""},
	"u":     {0, 0, u, Numeric, "u", "micro", "micro", "", ""},
	"m":     {0, 0, m, Numeric, "m", "milli", "milli", "", ""},
	"c":     {0, 0, c, Numeric, "c", "centi", "centi", "", ""},
	"d":     {0, 0, d, Numeric, "d", "deci", "deci", "", ""},
	"1":     DimensionlessUnit,
	"":      DimensionlessUnit,
	"da":    {0, 0, da, Numeric, "da", "deca", "deca", "", ""},
	"dozen": {0, 0, 12, Numeric, "doz", "dozen", "dozen", "", ""},
	"h":     {0, 0, h, Numeric, "h", "hecto", "hecto", "", ""},
	"k":     {0, 0, k, Numeric, "k", "kilo", "kilo", "", ""},
	"M":     {0, 0, _M, Numeric, "M", "mega", "mega", "", ""},
	"G":     {0, 0, _G, Numeric, "G", "giga", "giga", "", ""},
	"T":     {0, 0, _T, Numeric, "T", "tera", "tera", "", ""},
	"P":     {0, 0, _P, Numeric, "P", "peta", "peta", "", ""},
	"E":     {0, 0, _E, Numeric, "E", "exa", "exa", "", ""},
	"Z":     {0, 0, _Z, Numeric, "Z", "zetta", "zetta", "", ""},
	"Y":     {0, 0, _Y, Numeric, "Y", "yotta", "yotta", "", ""},

	"myriad": {
		0, 0, 10000,
		Numeric,
		"myriad", "myriad", "myriads",
		"historically, ten thousand but latterly meaning" +
			" a countless number of things",
		"",
	},

	"million": {
		0, 0, _M,
		Numeric,
		"M", "million", "million", "",
		"",
	},
	"billion": {
		0, 0, _G,
		Numeric,
		"B", "billion", "billion",
		"Note that this reflects the, now universal, meaning of" +
			" 1000 million." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million squared which was the" +
			" original usage in England. If the text also uses" +
			" milliard then this is most likely to mean a million" +
			" million (what is now generally referred to as a trillion).",
		"",
	},
	"trillion": {
		0, 0, _T,
		Numeric,
		"Tr", "trillion", "trillion",
		"Note that this reflects the, now universal, meaning of" +
			" 1000 billion." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million cubed which was the" +
			" original usage in England.",
		"",
	},
	"quadrillion": {
		0, 0, _P,
		Numeric,
		"Qu", "quadrillion", "quadrillion",
		"Note that this reflects the, now universal, meaning of" +
			" 1000 trillion." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million to the power of four" +
			" which was the original usage in England.",
		"",
	},
	"milliard": {
		0, 0, 1e6 * 1000,
		Numeric,
		"milliard", "milliard", "milliard",
		"an English (UK) term for a thousand million," +
			" now obsolete, it has been superseded by" +
			" the US meaning of billion.",
		"",
	},
	"billion (UK)": {
		0, 0, 1e12,
		Numeric,
		"B (UK)", "billion (UK)", "billion (UK)",
		"Note that this reflects the, now obsolete, meaning of" +
			" a million squared." +
			" Be aware that some older texts, especially in the UK" +
			" might use this to mean a million squared which was the" +
			" original usage in England. You will need to judge by" +
			" context whether the original or inflated meaning is" +
			" intended.",
		"",
	},
	"billiard": {
		0, 0, 1e12 * 1000,
		Numeric,
		"billiard", "billiard", "billiard",
		"an English (UK) term for a thousand (UK) billion," +
			" now obsolete, it has been superseded by" +
			" the US meaning of quadrillion.",
		"",
	},
	"trillion (UK)": {
		0, 0, 1e18,
		Numeric,
		"Tr (UK)", "trillion (UK)", "trillion (UK)",
		"Note that this reflects the, now obsolete, meaning of" +
			" a million cubed." +
			" Be aware that some older texts, especially in the UK" +
			" might use trillion to mean a million cubed which was the" +
			" original usage in the UK. You will need to judge by" +
			" context whether the original or inflated meaning is intended.",
		"",
	},
	"trilliard": {
		0, 0, 1e18 * 1000,
		Numeric,
		"trilliard", "trilliard", "trilliard",
		"an English (UK) term for a thousand (UK) trillion," +
			" now obsolete.",
		"",
	},

	"lakh": {
		0, 0, 1e5,
		Numeric,
		"lakh", "lakh", "lakh",
		"Indian: 1,00,000.",
		"",
	},
	"crore": {
		0, 0, 1e7,
		Numeric,
		"crore", "crore", "crore",
		"Indian: 1,00,00,000.",
		"",
	},

	"pony": {
		0, 0, 25,
		Numeric,
		"pony", "pony", "ponies",
		"UK slang: £25.",
		"",
	},
	"monkey": {
		0, 0, 500,
		Numeric,
		"monkey", "monkey", "monkeys",
		"UK slang: £500.",
		"",
	},
	"grand": {
		0, 0, k,
		Numeric,
		"grand", "grand", "grand",
		"UK slang: £1000.",
		"",
	},
}

var dimensionlessAliases = map[string]Alias{
	"yard": {
		"milliard",
		"Financial slang meaning a billion pounds/dollars/etc." +
			" It is a contraction of milliard",
	},
}
