package units

const (
	inchToMetre = 0.0254
	yardToMetre = inchToMetre * 12 * 3
)

// UnitOfDistance represents the base unit of distance
var UnitOfDistance = Family{
	BaseUnitName: "metre",
	Description:  "unit of distance",
}

// MetreUnit is a suitable default value for a UnitOfDistance
var MetreUnit = Unit{1, UnitOfDistance, "m", UnitOfDistance.BaseUnitName}

var DistanceNames = map[string]Unit{
	"ym":    {y, UnitOfDistance, "ym", "yoctometre"},
	"zm":    {z, UnitOfDistance, "zm", "zeptometre"},
	"am":    {a, UnitOfDistance, "am", "attometre"},
	"fm":    {f, UnitOfDistance, "fm", "femtometre"},
	"pm":    {p, UnitOfDistance, "pm", "picometre"},
	"nm":    {n, UnitOfDistance, "nm", "nanometre"},
	"um":    {u, UnitOfDistance, "um", "micrometre"},
	"mm":    {m, UnitOfDistance, "mm", "millimetre"},
	"cm":    {c, UnitOfDistance, "cm", "centimetre"},
	"dm":    {d, UnitOfDistance, "dm", "decimetre"},
	"metre": MetreUnit,
	"dam":   {da, UnitOfDistance, "dam", "decametre"},
	"hm":    {h, UnitOfDistance, "hm", "hectometre"},
	"km":    {k, UnitOfDistance, "km", "kilometre"},
	"Mm":    {_M, UnitOfDistance, "Mm", "megametre"},
	"Gm":    {_G, UnitOfDistance, "Gm", "gigametre"},
	"Tm":    {_T, UnitOfDistance, "Tm", "terametre"},
	"Pm":    {_P, UnitOfDistance, "Pm", "petametre"},
	"Em":    {_E, UnitOfDistance, "Em", "exametre"},
	"Zm":    {_Z, UnitOfDistance, "Zm", "zettametre"},
	"Ym":    {_Y, UnitOfDistance, "Ym", "yottametre"},

	"inch":          {inchToMetre, UnitOfDistance, "in", "inch"},
	"foot":          {inchToMetre * 12, UnitOfDistance, "ft", "foot"},
	"hand":          {inchToMetre * 4, UnitOfDistance, "hand", "hand"},
	"yard":          {yardToMetre, UnitOfDistance, "yd", "yard"},
	"rod":           {yardToMetre * 5.5, UnitOfDistance, "rod", "rod"},
	"chain":         {yardToMetre * 22, UnitOfDistance, "ch", "chain"},
	"furlong":       {yardToMetre * 220, UnitOfDistance, "fur", "furlong"},
	"mile":          {yardToMetre * 1760, UnitOfDistance, "mi", "mile"},
	"swimming-mile": {yardToMetre * 1650, UnitOfDistance, "swim-mi", "swimming mile"},
	"league":        {yardToMetre * 1760 * 3, UnitOfDistance, "lea", "league"},

	"admiralty-fathom": {yardToMetre * 2, UnitOfDistance, "adm ftm", "Admiralty fathom"},
	"fathom":           {1.852, UnitOfDistance, "ftm", "fathom"},
	"cable":            {185.2, UnitOfDistance, "cable", "cable"},
	"nautical-mile":    {1852, UnitOfDistance, "nautical mile", "nautical mile"},

	"astro-unit":   {1.49597871e11, UnitOfDistance, "AU", "Astronomical Unit"},
	"parsec":       {3.08567758767931e16, UnitOfDistance, "pc", "parsec"},
	"light-year":   {9.460528405106e15, UnitOfDistance, "ly", "light year"},
	"light-second": {2.99792458e8, UnitOfDistance, "ls", "light second"},

	"point": {0.0003528, UnitOfDistance, "pt", "point"},
	"pica":  {0.0042336, UnitOfDistance, "pc", "pica"},
}
