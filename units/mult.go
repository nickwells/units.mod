package units

import (
	"errors"
	"math"
)

// Mult represents a multiple that can be applied to a unit of size
// such as K (for 1000) or M (for 1000000)
//
// The Val gives the number of the BaseUnit that there are in the alternative
// unit. Given a measure in the base unit you can express this in the
// alternative unit by dividing by the Val
//
// for multiples of bytes the names follow the metric and IEC naming
// conventions so MB refers to a megabyte or 1000^2 bytes and MiB refers
// to a mebibyte or 1024^2 bytes
type Mult struct {
	Val      float64
	BaseUnit Unit
	Abbrev   string
	AltUnit  string
}

var (
	k  = 1000.0
	_M = math.Pow(1000, 2)
	_G = math.Pow(1000, 3)
	_T = math.Pow(1000, 4)
	_P = math.Pow(1000, 5)
	_E = math.Pow(1000, 6)
	_Z = math.Pow(1000, 7)
	_Y = math.Pow(1000, 8)

	h  = 100.0
	da = 10.0

	d = 0.1
	c = 0.01

	m = 1.0 / math.Pow(1000, 1)
	u = 1.0 / math.Pow(1000, 2)
	n = 1.0 / math.Pow(1000, 3)
	p = 1.0 / math.Pow(1000, 4)
	f = 1.0 / math.Pow(1000, 5)
	a = 1.0 / math.Pow(1000, 6)
	z = 1.0 / math.Pow(1000, 7)
	y = 1.0 / math.Pow(1000, 8)
)

const (
	_ = 1 << (10 * iota)
	ki
	mi
	gi
	ti
	pi
	ei
	zi
	yi
)

// NoUnitMult is a suitable default value for a Mult of NoUnit
var NoUnitMult = Mult{1, NoUnit, "", NoUnit.Name}

// SecondMult is a suitable default value for a Mult of Seconds
var SecondMult = Mult{1, Second, "sec", Second.Name}

// ByteMult is a suitable default value for a Mult of Bytes
var ByteMult = Mult{1, Byte, "byte", Byte.Name}

// MetreMult is a suitable default value for a Mult of Metres
var MetreMult = Mult{1, Metre, "m", Metre.Name}

// SquareMetreMult is a suitable default value for a Mult of SquareMetres
var SquareMetreMult = Mult{1, SquareMetre, "m^2", SquareMetre.Name}

// CubicMetreMult is a suitable default value for a Mult of CubicMetres
var CubicMetreMult = Mult{1, CubicMetre, "m^3", CubicMetre.Name}

// LitreMult is an alternative suitable default value for a Mult of CubicMetres
var LitreMult = Mult{1e-3, CubicMetre, "l", "litre"}

// GramMult is a suitable default value for a Mult of Grams
var GramMult = Mult{1, Gram, "g", Gram.Name}

const (
	grainToGram = 0.06479891
	ounceToGram = grainToGram * 437.5
	poundToGram = ounceToGram * 16

	inchToMetre = 0.0254
	yardToMetre = inchToMetre * 12 * 3

	yard2ToMetre2 = yardToMetre * yardToMetre

	bottleToCubicMetre    = 0.75e-3
	gallonToCubicMetre    = 4.54609e-3
	fluidOzToCubicMetre   = gallonToCubicMetre / 160
	usFluidOzToCubicMetre = 29.5735295625e-6
)

var umNames = map[string]Mult{
	"y":              {y, NoUnit, "y", "yocto"},
	"z":              {z, NoUnit, "z", "zepto"},
	"a":              {a, NoUnit, "a", "atto"},
	"f":              {f, NoUnit, "f", "femto"},
	"p":              {p, NoUnit, "p", "pico"},
	"n":              {n, NoUnit, "n", "nano"},
	"u":              {u, NoUnit, "u", "micro"},
	"m":              {m, NoUnit, "m", "milli"},
	"c":              {c, NoUnit, "c", "centi"},
	"d":              {d, NoUnit, "d", "deci"},
	"1":              NoUnitMult,
	"da":             {da, NoUnit, "da", "deca"},
	"pony":           {25, NoUnit, "pony", "pony"}, // London slang: £25
	"h":              {h, NoUnit, "h", "hecto"},
	"monkey":         {500, NoUnit, "monkey", "monkey"}, // London slang: £500
	"k":              {k, NoUnit, "k", "kilo"},
	"grand":          {k, NoUnit, "grand", "grand"}, // UK slang: £1000
	"lakh":           {1e5, NoUnit, "lakh", "lakh"}, // Indian: 1,00,000
	"M":              {_M, NoUnit, "M", "mega"},
	"million":        {_M, NoUnit, "M", "million"},
	"crore":          {1e7, NoUnit, "crore", "crore"}, // Indian: 1,00,00,000
	"G":              {_G, NoUnit, "G", "giga"},
	"yard-financial": {_G, NoUnit, "yard", "yard"}, // from milliard (thousand million)
	"T":              {_T, NoUnit, "T", "tera"},
	"P":              {_P, NoUnit, "P", "peta"},
	"E":              {_E, NoUnit, "E", "exa"},
	"Z":              {_Z, NoUnit, "Z", "zetta"},
	"Y":              {_Y, NoUnit, "Y", "yotta"},

	"ysec":   {y, Second, "ysec", "yoctosecond"},
	"zsec":   {z, Second, "zsec", "zeptosecond"},
	"asec":   {a, Second, "asec", "attosecond"},
	"fsec":   {f, Second, "fsec", "femtosecond"},
	"psec":   {p, Second, "psec", "picosecond"},
	"nsec":   {n, Second, "nsec", "nanosecond"},
	"usec":   {u, Second, "μsec", "microsecond"},
	"msec":   {m, Second, "msec", "millisecond"},
	"csec":   {c, Second, "csec", "centisecond"},
	"dsec":   {d, Second, "dsec", "decisecond"},
	"sec":    SecondMult,
	"second": SecondMult,
	"minute": {60, Second, "min", "minute"},
	"hour":   {60 * 60, Second, "hr", "hour"},
	"day":    {60 * 60 * 24, Second, "day", "day"},
	"week":   {60 * 60 * 24 * 7, Second, "week", "week"},
	"dasec":  {da, Second, "dasec", "decasecond"},
	"hsec":   {h, Second, "hsec", "hectosecond"},
	"ksec":   {k, Second, "ksec", "kilosecond"},
	"Msec":   {_M, Second, "Msec", "megasecond"},
	"Gsec":   {_G, Second, "Gsec", "gigasecond"},
	"Tsec":   {_T, Second, "Tsec", "terasecond"},
	"Psec":   {_P, Second, "Psec", "petasecond"},
	"Esec":   {_E, Second, "Esec", "exasecond"},
	"Zsec":   {_Z, Second, "Zsec", "zettasecond"},
	"Ysec":   {_Y, Second, "Ysec", "yottasecond"},

	"bit":    {0.125, Byte, "bit", "bit"},
	"nibble": {0.5, Byte, "nibble", "nibble"},
	"byte":   ByteMult,
	"KB":     {k, Byte, "KB", "kilobyte"},
	"MB":     {_M, Byte, "MB", "megabyte"},
	"GB":     {_G, Byte, "GB", "gigabyte"},
	"TB":     {_T, Byte, "TB", "terabyte"},
	"PB":     {_P, Byte, "PB", "petabyte"},
	"EB":     {_E, Byte, "EB", "exabyte"},
	"ZB":     {_Z, Byte, "ZB", "zettabyte"},
	"YB":     {_Y, Byte, "YB", "yottabyte"},

	"KiB": {ki, Byte, "KiB", "kibibyte"},
	"MiB": {mi, Byte, "MiB", "mebibyte"},
	"GiB": {gi, Byte, "GiB", "gibibyte"},
	"TiB": {ti, Byte, "TiB", "tebibyte"},
	"PiB": {pi, Byte, "PiB", "pebibyte"},
	"EiB": {ei, Byte, "EiB", "exbibyte"},
	"ZiB": {zi, Byte, "ZiB", "zebibyte"},
	"YiB": {yi, Byte, "YiB", "yobibyte"},

	"ym":    {y, Metre, "ym", "yoctometre"},
	"zm":    {z, Metre, "zm", "zeptometre"},
	"am":    {a, Metre, "am", "attometre"},
	"fm":    {f, Metre, "fm", "femtometre"},
	"pm":    {p, Metre, "pm", "picometre"},
	"nm":    {n, Metre, "nm", "nanometre"},
	"um":    {u, Metre, "um", "micrometre"},
	"mm":    {m, Metre, "mm", "millimetre"},
	"cm":    {c, Metre, "cm", "centimetre"},
	"dm":    {d, Metre, "dm", "decimetre"},
	"metre": MetreMult,
	"dam":   {da, Metre, "dam", "decametre"},
	"hm":    {h, Metre, "hm", "hectometre"},
	"km":    {k, Metre, "km", "kilometre"},
	"Mm":    {_M, Metre, "Mm", "megametre"},
	"Gm":    {_G, Metre, "Gm", "gigametre"},
	"Tm":    {_T, Metre, "Tm", "terametre"},
	"Pm":    {_P, Metre, "Pm", "petametre"},
	"Em":    {_E, Metre, "Em", "exametre"},
	"Zm":    {_Z, Metre, "Zm", "zettametre"},
	"Ym":    {_Y, Metre, "Ym", "yottametre"},

	"inch":          {inchToMetre, Metre, "in", "inch"},
	"foot":          {inchToMetre * 12, Metre, "ft", "foot"},
	"hand":          {inchToMetre * 4, Metre, "hand", "hand"},
	"yard":          {yardToMetre, Metre, "yd", "yard"},
	"rod":           {yardToMetre * 5.5, Metre, "rod", "rod"},
	"chain":         {yardToMetre * 22, Metre, "ch", "chain"},
	"furlong":       {yardToMetre * 220, Metre, "fur", "furlong"},
	"mile":          {yardToMetre * 1760, Metre, "mi", "mile"},
	"swimming-mile": {yardToMetre * 1650, Metre, "swim-mi", "swimming mile"},
	"league":        {yardToMetre * 1760 * 3, Metre, "lea", "league"},

	"admiralty-fathom": {yardToMetre * 2, Metre, "adm ftm", "Admiralty fathom"},
	"fathom":           {1.852, Metre, "ftm", "fathom"},
	"cable":            {185.2, Metre, "cable", "cable"},
	"nautical-mile":    {1852, Metre, "nautical mile", "nautical mile"},

	"astro-unit":   {1.49597871e11, Metre, "AU", "Astronomical Unit"},
	"parsec":       {3.08567758767931e16, Metre, "pc", "parsec"},
	"light-year":   {9.460528405106e15, Metre, "ly", "light year"},
	"light-second": {2.99792458e8, Metre, "ls", "light second"},

	"point": {0.0003528, Metre, "pt", "point"},
	"pica":  {0.0042336, Metre, "pc", "pica"},

	"square-metre":     SquareMetreMult,
	"hectare":          {1e4, SquareMetre, "ha", "hectare"},
	"square-kilometre": {1e6, SquareMetre, "sq km", "square kilometre"},
	"square-yard":      {yard2ToMetre2, SquareMetre, "sq yd", "square yard"},
	"rood":             {yard2ToMetre2 * 1210, SquareMetre, "rood", "rood"},
	"acre":             {yard2ToMetre2 * 4840, SquareMetre, "acre", "acre"},
	"oxgang": {yard2ToMetre2 * 4840 * 15, SquareMetre,
		"oxgang", "oxgang"},
	"Wales": {20779 * 1e6, SquareMetre,
		"Wales", "Area the size of Wales"},

	"cubic-metre": CubicMetreMult,

	"yl":    {y * 1e-3, CubicMetre, "yl", "yoctolitre"},
	"zl":    {z * 1e-3, CubicMetre, "zl", "zeptolitre"},
	"al":    {a * 1e-3, CubicMetre, "al", "attolitre"},
	"fl":    {f * 1e-3, CubicMetre, "fl", "femtolitre"},
	"pl":    {p * 1e-3, CubicMetre, "pl", "picolitre"},
	"nl":    {n * 1e-3, CubicMetre, "nl", "nanolitre"},
	"ul":    {u * 1e-3, CubicMetre, "ul", "microlitre"},
	"ml":    {m * 1e-3, CubicMetre, "ml", "millilitre"},
	"cc":    {m * 1e-3, CubicMetre, "cc", "cubic centimetre"},
	"cl":    {c * 1e-3, CubicMetre, "cl", "centilitre"},
	"dl":    {d * 1e-3, CubicMetre, "dl", "decilitre"},
	"litre": LitreMult,
	"liter": LitreMult, // for American users
	"dal":   {da * 1e-3, CubicMetre, "dal", "decalitre"},
	"hl":    {h * 1e-3, CubicMetre, "hl", "hectolitre"},
	"kl":    {k * 1e-3, CubicMetre, "kl", "kilolitre"},
	"Ml":    {_M * 1e-3, CubicMetre, "Ml", "megalitre"},
	"Gl":    {_G * 1e-3, CubicMetre, "Gl", "gigalitre"},
	"Tl":    {_T * 1e-3, CubicMetre, "Tl", "teralitre"},
	"cubic-kilometre": {_T * 1e-3, CubicMetre,
		"cubic kilometre", "cubic kilometre"},
	"Pl": {_P * 1e-3, CubicMetre, "Pl", "petalitre"},
	"El": {_E * 1e-3, CubicMetre, "El", "exalitre"},
	"Zl": {_Z * 1e-3, CubicMetre, "Zl", "zettalitre"},
	"Yl": {_Y * 1e-3, CubicMetre, "Yl", "yottalitre"},

	"bottle-wine": {bottleToCubicMetre, CubicMetre, "bottle", "bottle (wine)"},
	"magnum": {bottleToCubicMetre * 2, CubicMetre,
		"Magnum", "Magnum (wine)"},
	"marie-jeanne": {bottleToCubicMetre * 3, CubicMetre,
		"Marie Jeanne", "Marie Jeanne (wine)"},
	"tappit-hen": {bottleToCubicMetre * 3, CubicMetre,
		"Tappit Hen", "Tappit Hen (port)"},
	"tregnum": {bottleToCubicMetre * 3, CubicMetre,
		"Tregnum", "Tregnum (port)"},
	"jeroboam": {bottleToCubicMetre * 4, CubicMetre,
		"Jeroboam", "Jeroboam (wine)"},
	"rehoboam": {bottleToCubicMetre * 6, CubicMetre,
		"Rehoboam", "Rehoboam (wine)"},
	"methuselah": {bottleToCubicMetre * 8, CubicMetre,
		"Methuselah", "Methuselah (wine)"},
	"salmanazar": {bottleToCubicMetre * 12, CubicMetre,
		"Salmanazar", "Salmanazar (wine)"},
	"balthazar": {bottleToCubicMetre * 16, CubicMetre,
		"Balthazar", "Balthazar (wine)"},
	"nebuchadnezzar": {bottleToCubicMetre * 20, CubicMetre,
		"Nebuchadnezzar", "Nebuchadnezzar (wine)"},

	"minim": {fluidOzToCubicMetre / 480, CubicMetre,
		"min", "minim"},
	"fluid-scruple": {fluidOzToCubicMetre / 24, CubicMetre,
		"fl s", "fluid scruple"},
	"fluid-drachm": {fluidOzToCubicMetre / 8, CubicMetre,
		"fl dr", "fluid drachm"},
	"fluid-ounce": {fluidOzToCubicMetre, CubicMetre, "fl oz", "fluid ounce"},
	"gill":        {fluidOzToCubicMetre * 5, CubicMetre, "gi", "gill"},
	"pint":        {fluidOzToCubicMetre * 20, CubicMetre, "pt", "pint"},
	"quart":       {fluidOzToCubicMetre * 40, CubicMetre, "qt", "quart"},
	"gallon":      {gallonToCubicMetre, CubicMetre, "gal", "gallon"},
	"pin":         {gallonToCubicMetre * 4.5, CubicMetre, "pin", "pin"},
	"firkin":      {gallonToCubicMetre * 9, CubicMetre, "firkin", "firkin"},
	"kilderkin":   {gallonToCubicMetre * 18, CubicMetre, "kilderkin", "kilderkin"},
	"barrel":      {gallonToCubicMetre * 36, CubicMetre, "barrel", "barrel"},
	"hogshead":    {gallonToCubicMetre * 54, CubicMetre, "hogshead", "hogshead"},

	"teaspoon": {usFluidOzToCubicMetre * 0.5 / 3.0, CubicMetre,
		"tsp", "teaspoon"},
	"tablespoon": {usFluidOzToCubicMetre * 0.5, CubicMetre,
		"Tbsp", "tablespoon"},
	"us-fluid-ounce": {usFluidOzToCubicMetre, CubicMetre,
		"US fl oz", "US fluid ounce"},
	"us-shot":  {usFluidOzToCubicMetre * 1.5, CubicMetre, "US gi", "US gill"},
	"us-gill":  {usFluidOzToCubicMetre * 4, CubicMetre, "US gi", "US gill"},
	"us-cup":   {usFluidOzToCubicMetre * 8, CubicMetre, "US cp", "US cup"},
	"us-pint":  {usFluidOzToCubicMetre * 16, CubicMetre, "US pt", "US pint"},
	"us-quart": {usFluidOzToCubicMetre * 32, CubicMetre, "US qt", "US quart"},
	"us-gallon": {usFluidOzToCubicMetre * 128, CubicMetre,
		"US gal", "US gallon"},
	"barrel-oil": {usFluidOzToCubicMetre * 128 * 42, CubicMetre,
		"bbl(oil)", "oil barrel"},

	"yg":        {y, Gram, "yg", "yoctogram"},
	"zg":        {z, Gram, "zg", "zeptogram"},
	"ag":        {a, Gram, "ag", "attogram"},
	"fg":        {f, Gram, "fg", "femtogram"},
	"pg":        {p, Gram, "pg", "picogram"},
	"ng":        {n, Gram, "ng", "nanogram"},
	"ug":        {u, Gram, "ug", "microgram"},
	"mg":        {m, Gram, "mg", "milligram"},
	"cg":        {c, Gram, "cg", "centigram"},
	"dg":        {d, Gram, "dg", "decigram"},
	"gram":      GramMult,
	"dag":       {da, Gram, "dag", "decagram"},
	"hg":        {h, Gram, "hg", "hectogram"},
	"kg":        {k, Gram, "kg", "kilogram"},
	"Mg":        {_M, Gram, "Mg", "megagram"},
	"tonne":     {_M, Gram, "T", "tonne"},
	"Gg":        {_G, Gram, "Gg", "gigagram"},
	"kilotonne": {_G, Gram, "KT", "kilotonne"},
	"Tg":        {_T, Gram, "Tg", "teragram"},
	"megatonne": {_T, Gram, "MT", "megatonne"},
	"Pg":        {_P, Gram, "Pg", "petagram"},
	"Eg":        {_E, Gram, "Eg", "exagram"},
	"Zg":        {_Z, Gram, "Zg", "zettagram"},
	"Yg":        {_Y, Gram, "Yg", "yottagram"},

	"grain":      {grainToGram, Gram, "gr", "grain"},
	"troy-ounce": {grainToGram * 480, Gram, "t oz", "troy ounce"},
	"scruple":    {grainToGram * 20, Gram, "scruple", "scruple"},
	"dram":       {grainToGram * 60, Gram, "dr", "dram"},
	"drachm":     {grainToGram * 60, Gram, "dr", "drachm"},

	"ounce":         {ounceToGram, Gram, "oz", "ounce"},
	"pound":         {poundToGram, Gram, "lb", "pound"},
	"stone":         {poundToGram * 14, Gram, "st", "stone"},
	"hundredweight": {poundToGram * 112, Gram, "cwt", "hundredweight"},
	"short-hundredweight": {poundToGram * 100, Gram,
		"cental", "short hundredweight"},
	"imperial-ton": {poundToGram * 112 * 20, Gram, "t", "imperial ton"},
	"short-ton":    {poundToGram * 2000, Gram, "short ton", "short ton"},

	"earth-mass": {5.9722e27, Gram, "M\u2295", "Earth Mass"},
	"solar-mass": {1.98847e33, Gram, "M\u2299", "Solar Mass"},
}

// GetMult returns a Mult having the passed name, a non-nil error is returned
// if the name is not found
func GetMult(name string) (Mult, error) {
	var m Mult
	m, ok := umNames[name]
	if !ok {
		return m, errors.New("no mult found for '" + name + "'")
	}
	return m, nil
}
