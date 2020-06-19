package units

const (
	inchToMetre = 0.0254
	footToMetre = inchToMetre * 12
	yardToMetre = footToMetre * 3
	rodToMetre  = yardToMetre * 5.5
	lightSecond = 2.99792458e8
)

// UnitOfDistance represents the base unit of distance
var UnitOfDistance = Family{
	BaseUnitName: "metre",
	Description:  "unit of distance",
}

// MetreUnit is a suitable default value for a UnitOfDistance
var MetreUnit = Unit{0, 0, 1, UnitOfDistance, "m", UnitOfDistance.BaseUnitName, "metres"}

// DistanceNames maps names to units of distance
var DistanceNames = map[string]Unit{
	// metric
	"ym":    {0, 0, y, UnitOfDistance, "ym", "yoctometre", "yoctometres"},
	"zm":    {0, 0, z, UnitOfDistance, "zm", "zeptometre", "zeptometres"},
	"am":    {0, 0, a, UnitOfDistance, "am", "attometre", "attometres"},
	"fm":    {0, 0, f, UnitOfDistance, "fm", "femtometre", "femtometres"},
	"pm":    {0, 0, p, UnitOfDistance, "pm", "picometre", "picometres"},
	"nm":    {0, 0, n, UnitOfDistance, "nm", "nanometre", "nanometres"},
	"um":    {0, 0, u, UnitOfDistance, "um", "micrometre", "micrometres"},
	"mm":    {0, 0, m, UnitOfDistance, "mm", "millimetre", "millimetres"},
	"cm":    {0, 0, c, UnitOfDistance, "cm", "centimetre", "centimetres"},
	"dm":    {0, 0, d, UnitOfDistance, "dm", "decimetre", "decimetres"},
	"metre": MetreUnit,
	"dam":   {0, 0, da, UnitOfDistance, "dam", "decametre", "decametres"},
	"hm":    {0, 0, h, UnitOfDistance, "hm", "hectometre", "hectometres"},
	"km":    {0, 0, k, UnitOfDistance, "km", "kilometre", "kilometres"},
	"klick": {0, 0, k, UnitOfDistance, "klick", "klick", "klicks"}, // US Army term
	"Mm":    {0, 0, _M, UnitOfDistance, "Mm", "megametre", "megametres"},
	"Gm":    {0, 0, _G, UnitOfDistance, "Gm", "gigametre", "gigametres"},
	"Tm":    {0, 0, _T, UnitOfDistance, "Tm", "terametre", "terametres"},
	"Pm":    {0, 0, _P, UnitOfDistance, "Pm", "petametre", "petametres"},
	"Em":    {0, 0, _E, UnitOfDistance, "Em", "exametre", "exametres"},
	"Zm":    {0, 0, _Z, UnitOfDistance, "Zm", "zettametre", "zettametres"},
	"Ym":    {0, 0, _Y, UnitOfDistance, "Ym", "yottametre", "yottametres"},

	// Imperial / US
	"inch":          {0, 0, inchToMetre, UnitOfDistance, "in", "inch", "inches"},
	"hand":          {0, 0, inchToMetre * 4, UnitOfDistance, "hand", "hand", "hands"},
	"foot":          {0, 0, footToMetre, UnitOfDistance, "ft", "foot", "feet"},
	"yard":          {0, 0, yardToMetre, UnitOfDistance, "yd", "yard", "yards"},
	"rod":           {0, 0, rodToMetre, UnitOfDistance, "rod", "rod", "rods"},
	"pole":          {0, 0, rodToMetre, UnitOfDistance, "pole", "pole", "poles"},
	"perch":         {0, 0, rodToMetre, UnitOfDistance, "perch", "perch", "perch"},
	"chain":         {0, 0, rodToMetre * 4, UnitOfDistance, "ch", "chain", "chains"},
	"furlong":       {0, 0, rodToMetre * 40, UnitOfDistance, "fur", "furlong", "furlongs"},
	"mile":          {0, 0, yardToMetre * 1760, UnitOfDistance, "mi", "mile", "miles"},
	"swimming-mile": {0, 0, yardToMetre * 1650, UnitOfDistance, "swim-mi", "swimming mile", "swimming miles"},
	"league":        {0, 0, yardToMetre * 1760 * 3, UnitOfDistance, "lea", "league", "leagues"},

	"fathom":           {0, 0, yardToMetre * 2, UnitOfDistance, "ftm", "fathom", "fathoms"},
	"nautical-mile":    {0, 0, 1852, UnitOfDistance, "M", "nautical mile", "nautical miles"},
	"admiralty-fathom": {0, 0, 6080 * yardToMetre / 3000, UnitOfDistance, "adm ftm", "Admiralty fathom", "Admiralty fathoms"},

	// regional
	"scots mile": {0, 0, yardToMetre * 1976, UnitOfDistance, "Scots mile", "Scots mile", "Scots miles"},
	"irish mile": {0, 0, yardToMetre * 2240, UnitOfDistance, "Irish mile", "Irish mile", "Irish miles"},

	// astronomical
	"astro-unit":   {0, 0, 1.49597871e11, UnitOfDistance, "AU", "Astronomical Unit", "Astronomical Units"},
	"parsec":       {0, 0, 3.08567758767931e16, UnitOfDistance, "pc", "parsec", "parsecs"},
	"light-year":   {0, 0, lightSecond * 86400 * 365.25, UnitOfDistance, "ly", "light year", "light years"},
	"light-second": {0, 0, lightSecond, UnitOfDistance, "ls", "light second", "light seconds"},

	// printing
	"point": {0, 0, 0.0003528, UnitOfDistance, "pt", "point", "points"},
	"pica":  {0, 0, 0.0042336, UnitOfDistance, "pc", "pica", "pica"},
}
