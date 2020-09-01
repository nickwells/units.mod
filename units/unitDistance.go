package units

const (
	inchToMetre  = 0.0254
	footToMetre  = inchToMetre * 12
	yardToMetre  = footToMetre * 3
	rodToMetre   = yardToMetre * 5.5
	lightSecond  = 2.99792458e8
	pointToMetre = 0.0003528
)

// UnitOfDistance represents the base unit of distance
var UnitOfDistance = Family{
	BaseUnitName: "metre",
	Description:  "unit of distance",
}

// MetreUnit is a suitable default value for a UnitOfDistance
var MetreUnit = Unit{
	0, 0, 1,
	UnitOfDistance, "m", UnitOfDistance.BaseUnitName, "metres",
	"The base unit of distance in the metric system"}

// DistanceNames maps names to units of distance
var DistanceNames = map[string]Unit{
	// metric
	"ym": {0, 0, y,
		UnitOfDistance,
		"ym", "yoctometre", "yoctometres",
		"a metric measure of distance (fractional)"},
	"zm": {0, 0, z,
		UnitOfDistance,
		"zm", "zeptometre", "zeptometres",
		"a metric measure of distance (fractional)"},
	"am": {0, 0, a,
		UnitOfDistance,
		"am", "attometre", "attometres",
		"a metric measure of distance (fractional)"},
	"fm": {0, 0, f,
		UnitOfDistance,
		"fm", "femtometre", "femtometres",
		"a metric measure of distance (fractional)"},
	"pm": {0, 0, p,
		UnitOfDistance,
		"pm", "picometre", "picometres",
		"a metric measure of distance (fractional)"},
	"nm": {0, 0, n,
		UnitOfDistance,
		"nm", "nanometre", "nanometres",
		"a metric measure of distance (fractional)"},
	"um": {0, 0, u,
		UnitOfDistance,
		"um", "micrometre", "micrometres",
		"a metric measure of distance (fractional)"},
	"mm": {0, 0, m,
		UnitOfDistance,
		"mm", "millimetre", "millimetres",
		"a metric measure of distance (fractional)"},
	"cm": {0, 0, c,
		UnitOfDistance,
		"cm", "centimetre", "centimetres",
		"a metric measure of distance (fractional)"},
	"dm": {0, 0, d,
		UnitOfDistance,
		"dm", "decimetre", "decimetres",
		"a metric measure of distance (fractional)"},
	"m":     MetreUnit,
	"metre": MetreUnit,
	"dam": {0, 0, da,
		UnitOfDistance,
		"dam", "decametre", "decametres",
		"a metric measure of distance"},
	"hm": {0, 0, h,
		UnitOfDistance,
		"hm", "hectometre", "hectometres",
		"a metric measure of distance"},
	"km": {0, 0, k,
		UnitOfDistance,
		"km", "kilometre", "kilometres",
		"a metric measure of distance"},
	"klick": {0, 0, k,
		UnitOfDistance,
		"klick", "klick", "klicks",
		"a US Army name for a kilometre"},
	"Mm": {0, 0, _M,
		UnitOfDistance,
		"Mm", "megametre", "megametres",
		"a metric measure of distance"},
	"Gm": {0, 0, _G,
		UnitOfDistance,
		"Gm", "gigametre", "gigametres",
		"a metric measure of distance"},
	"Tm": {0, 0, _T,
		UnitOfDistance,
		"Tm", "terametre", "terametres",
		"a metric measure of distance"},
	"Pm": {0, 0, _P,
		UnitOfDistance,
		"Pm", "petametre", "petametres",
		"a metric measure of distance"},
	"Em": {0, 0, _E,
		UnitOfDistance,
		"Em", "exametre", "exametres",
		"a metric measure of distance"},
	"Zm": {0, 0, _Z,
		UnitOfDistance,
		"Zm", "zettametre", "zettametres",
		"a metric measure of distance"},
	"Ym": {0, 0, _Y,
		UnitOfDistance,
		"Ym", "yottametre", "yottametres",
		"a metric measure of distance"},

	// Imperial / US
	"inch": {0, 0, inchToMetre,
		UnitOfDistance,
		"in", "inch", "inches",
		"imperial/US measure of distance"},
	"hand": {0, 0, inchToMetre * 4,
		UnitOfDistance,
		"hand", "hand", "hands",
		"an imperial measure of length, typically used to measure" +
			" the height of horses"},
	"link": {0, 0, footToMetre * 66 / 100,
		UnitOfDistance,
		"lnk", "link", "links",
		"an imperial measure of length, derived from Gunter's chain" +
			" - a metal chain 66 feet long with 100 links"},
	"foot": {0, 0, footToMetre,
		UnitOfDistance,
		"ft", "foot", "feet",
		"an imperial measure of length, 12 inches"},
	"US survey foot": {0, 0, 1200.0 / 3937.0,
		UnitOfDistance,
		"ft", "foot", "feet",
		"a US Customary unit of length, longer than a standard foot" +
			" by approximately 2 parts in a million"},
	"yard": {0, 0, yardToMetre,
		UnitOfDistance,
		"yd", "yard", "yards",
		"an imperial measure of length, 3 feet"},
	"rod": {0, 0, rodToMetre,
		UnitOfDistance,
		"rod", "rod", "rods",
		"an imperial measure of length, originally equal to 5 yards" +
			" it was redefined as 5.5 yards when England moved from" +
			" the North German to the shorter standard foot in the" +
			" late 13th century"},
	"pole": {0, 0, rodToMetre,
		UnitOfDistance,
		"pole", "pole", "poles",
		"an imperial measure of length, an alternative name for a rod"},
	"perch": {0, 0, rodToMetre,
		UnitOfDistance,
		"perch", "perch", "perches",
		"an imperial measure of length, an alternative name for a rod." +
			" Note that a perch can also be a measure of length or" +
			" volume under certain circumstances"},
	"chain": {0, 0, rodToMetre * 4,
		UnitOfDistance,
		"ch", "chain", "chains",
		"an imperial measure of length - four rods."},
	"furlong": {0, 0, rodToMetre * 40,
		UnitOfDistance,
		"fur", "furlong", "furlongs",
		"an imperial measure of length - forty rods. Originally a" +
			" furrow's length - the distance a team of oxen" +
			" could plough without resting, later standardised as" +
			" 40 rods."},
	"mile": {0, 0, yardToMetre * 1760,
		UnitOfDistance,
		"mi", "mile", "miles",
		"an imperial measure of length"},
	"swimming-mile": {0, 0, yardToMetre * 1650,
		UnitOfDistance,
		"swim-mi", "swimming mile", "swimming miles",
		"a coloquial and ill-defined distance used by swimmers," +
			" sometimes taken to be 1500 metres"},
	"league": {0, 0, yardToMetre * 1760 * 3,
		UnitOfDistance,
		"lea", "league", "leagues",
		"3 miles (perhaps the distance a person could walk in an hour)"},
	"nautical league": {0, 0, 5556,
		UnitOfDistance,
		"lea", "nautical league", "nautical leagues",
		"3 nautical miles"},

	"fathom": {0, 0, yardToMetre * 2,
		UnitOfDistance,
		"ftm", "fathom", "fathoms",
		"an imperial measure of distance equal to two yards"},
	"nautical-mile": {0, 0, 1852,
		UnitOfDistance,
		"M", "nautical mile", "nautical miles",
		"formerly one minute (1/60th of a degree) of latitude" +
			" around any line of latitude, now formally defined" +
			" to have the value given here"},
	"nautical-mile (US)": {0, 0, 6080.2 * footToMetre,
		UnitOfDistance,
		"M", "nautical mile (US)", "nautical miles (US)",
		"one minute (1/60th of a degree) of latitude around" +
			" the Clarke 1866 ellipsoid (6080.210 feet) to 5" +
			" significant figures"},
	"nautical-mile (Admiralty/UK)": {0, 0, 6080 * footToMetre,
		UnitOfDistance,
		"M", "Admiralty nautical mile", "Admiralty nautical miles",
		"an imperial measure of distance equal to" +
			"one minute (1/60th of a degree) of latitude around" +
			" the Clarke 1866 ellipsoid (6080.210 feet) to 4" +
			" significant figures"},
	"admiralty-fathom": {0, 0, 6080 * footToMetre / 1000,
		UnitOfDistance,
		"adm ftm", "Admiralty fathom", "Admiralty fathoms",
		"an imperial measure of distance equal to 1/1000th of" +
			" an Admiralty nautical mile (a little over 2 yards)"},

	// regional
	"scots mile": {0, 0, yardToMetre * 1976,
		UnitOfDistance,
		"Scots mile", "Scots mile", "Scots miles",
		"a measure of distance used in Scotland (slightly" +
			" longer than the English mile), it has not had" +
			" any formal use for over a century"},
	"irish mile": {0, 0, yardToMetre * 2240,
		UnitOfDistance,
		"Irish mile", "Irish mile", "Irish miles",
		"a measure of distance used in Ireland (slightly" +
			" longer than the English mile), no longer" +
			" commonly used"},

	// astronomical
	"astro-unit": {0, 0, 1.49597871e11,
		UnitOfDistance,
		"AU", "Astronomical Unit", "Astronomical Units",
		"the mean distance from the centre of the Earth" +
			" to the centre of the sun"},
	"parsec": {0, 0, 3.08567758767931e16,
		UnitOfDistance,
		"pc", "parsec", "parsecs",
		"the distance at which one astronomical unit" +
			" subtends an angle of one second of arc"},
	"light-year": {0, 0, lightSecond * 86400 * 365.25,
		UnitOfDistance,
		"ly", "light year", "light years",
		"the distance travelled by light in a vacuum" +
			" in one year (365.25 days)"},
	"light-second": {0, 0, lightSecond,
		UnitOfDistance,
		"ls", "light second", "light seconds",
		"the distance travelled by light in a vacuum" +
			" in one second"},

	// printing
	"point": {0, 0, pointToMetre,
		UnitOfDistance,
		"pt", "point", "points",
		"a printing term"},
	"pica": {0, 0, 12 * pointToMetre,
		UnitOfDistance,
		"pc", "pica", "pica",
		"a printing term"},
}
