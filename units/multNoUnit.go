package units

// NoUnit represents a dimensionless value
var NoUnit = Unit{
	BaseUnitName: "",
	Description:  "dimensionless value",
}

// NoUnitMult is a suitable default value for a Mult of NoUnit
var NoUnitMult = Mult{1, NoUnit, "", NoUnit.BaseUnitName}

var NoUnitNames = map[string]Mult{
	"y":       {y, NoUnit, "y", "yocto"},
	"z":       {z, NoUnit, "z", "zepto"},
	"a":       {a, NoUnit, "a", "atto"},
	"f":       {f, NoUnit, "f", "femto"},
	"p":       {p, NoUnit, "p", "pico"},
	"n":       {n, NoUnit, "n", "nano"},
	"u":       {u, NoUnit, "u", "micro"},
	"m":       {m, NoUnit, "m", "milli"},
	"c":       {c, NoUnit, "c", "centi"},
	"d":       {d, NoUnit, "d", "deci"},
	"1":       NoUnitMult,
	"":        NoUnitMult,
	"da":      {da, NoUnit, "da", "deca"},
	"pony":    {25, NoUnit, "pony", "pony"}, // UK slang: £25
	"h":       {h, NoUnit, "h", "hecto"},
	"monkey":  {500, NoUnit, "monkey", "monkey"}, // UK slang: £500
	"k":       {k, NoUnit, "k", "kilo"},
	"grand":   {k, NoUnit, "grand", "grand"}, // UK slang: £1000
	"lakh":    {1e5, NoUnit, "lakh", "lakh"}, // Indian: 1,00,000
	"M":       {_M, NoUnit, "M", "mega"},
	"million": {_M, NoUnit, "M", "million"},
	"crore":   {1e7, NoUnit, "crore", "crore"}, // Indian: 1,00,00,000
	"G":       {_G, NoUnit, "G", "giga"},
	"yard":    {_G, NoUnit, "yard", "yard"}, // Financial: 10^9 (from milliard)
	"T":       {_T, NoUnit, "T", "tera"},
	"P":       {_P, NoUnit, "P", "peta"},
	"E":       {_E, NoUnit, "E", "exa"},
	"Z":       {_Z, NoUnit, "Z", "zetta"},
	"Y":       {_Y, NoUnit, "Y", "yotta"},
}
