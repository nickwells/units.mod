package units

// Numeric represents a dimensionless value
var Numeric = Family{
	BaseUnitName: "",
	Description:  "dimensionless value",
}

// DimensionlessUnit is a suitable default value for a Unit of Numeric
var DimensionlessUnit = Unit{1, Numeric, "", Numeric.BaseUnitName}

var DimensionlessNames = map[string]Unit{
	"y":       {y, Numeric, "y", "yocto"},
	"z":       {z, Numeric, "z", "zepto"},
	"a":       {a, Numeric, "a", "atto"},
	"f":       {f, Numeric, "f", "femto"},
	"p":       {p, Numeric, "p", "pico"},
	"n":       {n, Numeric, "n", "nano"},
	"u":       {u, Numeric, "u", "micro"},
	"m":       {m, Numeric, "m", "milli"},
	"c":       {c, Numeric, "c", "centi"},
	"d":       {d, Numeric, "d", "deci"},
	"1":       DimensionlessUnit,
	"":        DimensionlessUnit,
	"da":      {da, Numeric, "da", "deca"},
	"pony":    {25, Numeric, "pony", "pony"}, // UK slang: £25
	"h":       {h, Numeric, "h", "hecto"},
	"monkey":  {500, Numeric, "monkey", "monkey"}, // UK slang: £500
	"k":       {k, Numeric, "k", "kilo"},
	"grand":   {k, Numeric, "grand", "grand"}, // UK slang: £1000
	"lakh":    {1e5, Numeric, "lakh", "lakh"}, // Indian: 1,00,000
	"M":       {_M, Numeric, "M", "mega"},
	"million": {_M, Numeric, "M", "million"},
	"crore":   {1e7, Numeric, "crore", "crore"}, // Indian: 1,00,00,000
	"G":       {_G, Numeric, "G", "giga"},
	"yard":    {_G, Numeric, "yard", "yard"}, // Financial: 10^9 (from milliard)
	"T":       {_T, Numeric, "T", "tera"},
	"P":       {_P, Numeric, "P", "peta"},
	"E":       {_E, Numeric, "E", "exa"},
	"Z":       {_Z, Numeric, "Z", "zetta"},
	"Y":       {_Y, Numeric, "Y", "yotta"},
}
