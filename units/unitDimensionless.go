package units

// Numeric represents a dimensionless value
var Numeric = Family{
	BaseUnitName: "",
	Description:  "dimensionless value",
}

// DimensionlessUnit is a suitable default value for a Unit of Numeric
var DimensionlessUnit = Unit{
	0, 0, 1,
	Numeric, "", Numeric.BaseUnitName, Numeric.BaseUnitName, "",
}

// DimensionlessNames maps names to numeric (dimensionless) units
var DimensionlessNames = map[string]Unit{
	"y":           {0, 0, y, Numeric, "y", "yocto", "yocto", ""},
	"z":           {0, 0, z, Numeric, "z", "zepto", "zepto", ""},
	"a":           {0, 0, a, Numeric, "a", "atto", "atto", ""},
	"f":           {0, 0, f, Numeric, "f", "femto", "femto", ""},
	"p":           {0, 0, p, Numeric, "p", "pico", "pico", ""},
	"n":           {0, 0, n, Numeric, "n", "nano", "nano", ""},
	"u":           {0, 0, u, Numeric, "u", "micro", "micro", ""},
	"m":           {0, 0, m, Numeric, "m", "milli", "milli", ""},
	"c":           {0, 0, c, Numeric, "c", "centi", "centi", ""},
	"d":           {0, 0, d, Numeric, "d", "deci", "deci", ""},
	"1":           DimensionlessUnit,
	"":            DimensionlessUnit,
	"da":          {0, 0, da, Numeric, "da", "deca", "deca", ""},
	"dozen":       {0, 0, 12, Numeric, "doz", "dozen", "dozen", ""},
	"pony":        {0, 0, 25, Numeric, "pony", "pony", "ponies", ""}, // UK slang: £25
	"h":           {0, 0, h, Numeric, "h", "hecto", "hecto", ""},
	"monkey":      {0, 0, 500, Numeric, "monkey", "monkey", "monkeys", ""}, // UK slang: £500
	"k":           {0, 0, k, Numeric, "k", "kilo", "kilo", ""},
	"grand":       {0, 0, k, Numeric, "grand", "grand", "grand", ""}, // UK slang: £1000
	"lakh":        {0, 0, 1e5, Numeric, "lakh", "lakh", "lakh", ""},  // Indian: 1,00,000
	"M":           {0, 0, _M, Numeric, "M", "mega", "mega", ""},
	"million":     {0, 0, _M, Numeric, "M", "million", "million", ""},
	"billion":     {0, 0, _G, Numeric, "B", "billion", "billion", ""},
	"trillion":    {0, 0, _T, Numeric, "Tr", "trillion", "trillion", ""},
	"quadrillion": {0, 0, _P, Numeric, "Qu", "quadrillion", "quadrillion", ""},
	"milliard":    {0, 0, 1e6 * 1000, Numeric, "milliard", "milliard", "milliard", ""},
	"yard":        {0, 0, 1e6 * 1000, Numeric, "yard", "yard", "yards", ""}, // Financial: 10^9 (from milliard)
	"billiard":    {0, 0, 1e12 * 1000, Numeric, "billiard", "billiard", "billiard", ""},
	"trilliard":   {0, 0, 1e18 * 1000, Numeric, "trilliard", "trilliard", "trilliard", ""},
	"crore":       {0, 0, 1e7, Numeric, "crore", "crore", "crore", ""}, // Indian: 1,00,00,000
	"G":           {0, 0, _G, Numeric, "G", "giga", "giga", ""},
	"T":           {0, 0, _T, Numeric, "T", "tera", "tera", ""},
	"P":           {0, 0, _P, Numeric, "P", "peta", "peta", ""},
	"E":           {0, 0, _E, Numeric, "E", "exa", "exa", ""},
	"Z":           {0, 0, _Z, Numeric, "Z", "zetta", "zetta", ""},
	"Y":           {0, 0, _Y, Numeric, "Y", "yotta", "yotta", ""},
}
