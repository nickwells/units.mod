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
	"The base unit of distance in the metric system.",
}

// DistanceNames maps names to units of distance
var DistanceNames = map[string]Unit{
	// metric
	"ym": {
		0, 0, y,
		UnitOfDistance,
		"ym", "yoctometre", "yoctometres", "a metric measure of distance.",
	},
	"zm": {
		0, 0, z,
		UnitOfDistance,
		"zm", "zeptometre", "zeptometres", "a metric measure of distance.",
	},
	"am": {
		0, 0, a,
		UnitOfDistance,
		"am", "attometre", "attometres", "a metric measure of distance.",
	},
	"fm": {
		0, 0, f,
		UnitOfDistance,
		"fm", "femtometre", "femtometres", "a metric measure of distance.",
	},
	"pm": {
		0, 0, p,
		UnitOfDistance,
		"pm", "picometre", "picometres", "a metric measure of distance.",
	},
	"nm": {
		0, 0, n,
		UnitOfDistance,
		"nm", "nanometre", "nanometres", "a metric measure of distance.",
	},
	"um": {
		0, 0, u,
		UnitOfDistance,
		"um", "micrometre", "micrometres", "a metric measure of distance.",
	},
	"mm": {
		0, 0, m,
		UnitOfDistance,
		"mm", "millimetre", "millimetres", "a metric measure of distance.",
	},
	"cm": {
		0, 0, c,
		UnitOfDistance,
		"cm", "centimetre", "centimetres", "a metric measure of distance.",
	},
	"dm": {
		0, 0, d,
		UnitOfDistance,
		"dm", "decimetre", "decimetres", "a metric measure of distance.",
	},
	"metre": MetreUnit,
	"dam": {
		0, 0, da,
		UnitOfDistance,
		"dam", "decametre", "decametres", "a metric measure of distance.",
	},
	"hm": {
		0, 0, h,
		UnitOfDistance,
		"hm", "hectometre", "hectometres", "a metric measure of distance.",
	},
	"km": {
		0, 0, k,
		UnitOfDistance,
		"km", "kilometre", "kilometres", "a metric measure of distance.",
	},
	"mym": {
		0, 0, 10000,
		UnitOfDistance,
		"mym", "myriametre", "myriametres",
		"an obsolete metric measure of distance.",
	},
	"Mm": {
		0, 0, _M,
		UnitOfDistance,
		"Mm", "megametre", "megametres", "a metric measure of distance.",
	},
	"Gm": {
		0, 0, _G,
		UnitOfDistance,
		"Gm", "gigametre", "gigametres", "a metric measure of distance.",
	},
	"Tm": {
		0, 0, _T,
		UnitOfDistance,
		"Tm", "terametre", "terametres", "a metric measure of distance.",
	},
	"Pm": {
		0, 0, _P,
		UnitOfDistance,
		"Pm", "petametre", "petametres", "a metric measure of distance.",
	},
	"Em": {
		0, 0, _E,
		UnitOfDistance,
		"Em", "exametre", "exametres", "a metric measure of distance.",
	},
	"Zm": {
		0, 0, _Z,
		UnitOfDistance,
		"Zm", "zettametre", "zettametres", "a metric measure of distance.",
	},
	"Ym": {
		0, 0, _Y,
		UnitOfDistance,
		"Ym", "yottametre", "yottametres", "a metric measure of distance.",
	},

	// Imperial / US
	"barleycorn": {
		0, 0, inchToMetre / 3.0,
		UnitOfDistance,
		"in", "inch", "inches",
		"a non-metric UK measure of distance. This was defined in" +
			" the 'Composition of Yards and Perches' an English" +
			" statute from between 1266-1303.",
	},
	"inch": {
		0, 0, inchToMetre,
		UnitOfDistance,
		"in", "inch", "inches",
		"imperial/US measure of distance.",
	},
	"hand": {
		0, 0, inchToMetre * 4,
		UnitOfDistance,
		"hand", "hand", "hands",
		"an imperial measure of length, typically used to measure" +
			" the height of horses.",
	},
	"link": {
		0, 0, footToMetre * 66 / 100,
		UnitOfDistance,
		"lnk", "link", "links",
		"an imperial measure of length, derived from Gunter's chain" +
			" - a metal chain 66 feet long with 100 links." +
			" Edmund Gunter (1581-1626) was an English clergyman," +
			" mathematician, geometer and astronomer.",
	},
	"foot": {
		0, 0, footToMetre,
		UnitOfDistance,
		"ft", "foot", "feet",
		"an imperial measure of length, 12 inches.",
	},
	"foot (North German)": {
		0, 0, footToMetre * 11.0 / 10.0,
		UnitOfDistance,
		"ft (North German)", "foot (North German)", "feet (North German)",
		"a medieval measure of length slightly longer than a standard foot.",
	},
	"foot (Roman)": {
		0, 0, 0.296,
		UnitOfDistance,
		"ft (Roman)", "foot (Roman)", "feet (Roman)",
		"a Roman measure of length. There are several alternative lengths" +
			" but the one used here is a popular value.",
	},
	"foot (Parisian)": {
		0, 0, 0.325,
		UnitOfDistance,
		"ft (Parisian)", "foot (Parisian)", "feet (Parisian)",
		"a French measure of length.",
	},
	"foot (metric)": {
		0, 0, 0.3,
		UnitOfDistance,
		"ft (metric)", "foot (metric)", "feet (metric)",
		"a measure of length - ISO 2848.",
	},
	"US survey foot": {
		0, 0, 1200.0 / 3937.0,
		UnitOfDistance,
		"ft", "foot", "feet",
		"a US Customary unit of length, longer than a standard foot" +
			" by approximately 2 parts in a million.",
	},
	"Indian survey foot": {
		0, 0, 0.3047996,
		UnitOfDistance,
		"ft", "foot", "feet",
		"an Indian unit of length.",
	},
	"yard": {
		0, 0, yardToMetre,
		UnitOfDistance,
		"yd", "yard", "yards",
		"an imperial measure of length, 3 feet.",
	},
	"rod": {
		0, 0, rodToMetre,
		UnitOfDistance,
		"rod", "rod", "rods",
		"an imperial measure of length, originally equal to 5 yards" +
			" it was redefined as 5.5 yards when England moved from" +
			" the North German to the shorter standard foot in the" +
			" late 13th century" +
			" Note that a rod can also be a measure of area" +
			" under certain circumstances.",
	},
	"chain": {
		0, 0, rodToMetre * 4,
		UnitOfDistance,
		"ch", "chain", "chains",
		"an imperial measure of length - four rods.",
	},
	"furlong": {
		0, 0, rodToMetre * 40,
		UnitOfDistance,
		"fur", "furlong", "furlongs",
		"an imperial measure of length - forty rods. Originally a" +
			" furrow's length - the distance a team of oxen" +
			" could plough without resting, later standardised as" +
			" 40 rods.",
	},
	"mile": {
		0, 0, yardToMetre * 1760,
		UnitOfDistance,
		"mi", "mile", "miles",
		"an imperial measure of length.",
	},
	"swimming-mile": {
		0, 0, yardToMetre * 1650,
		UnitOfDistance,
		"swim-mi", "swimming mile", "swimming miles",
		"a coloquial and ill-defined distance used by swimmers," +
			" sometimes taken to be 1500 metres.",
	},
	"league": {
		0, 0, yardToMetre * 1760 * 3,
		UnitOfDistance,
		"lea", "league", "leagues",
		"3 miles (perhaps the distance a person could walk in an hour).",
	},
	"nautical league": {
		0, 0, 5556,
		UnitOfDistance,
		"lea", "nautical league", "nautical leagues",
		"3 nautical miles.",
	},

	"fathom": {
		0, 0, yardToMetre * 2,
		UnitOfDistance,
		"ftm", "fathom", "fathoms",
		"an imperial measure of distance equal to two yards.",
	},
	"nautical-mile": {
		0, 0, 1852,
		UnitOfDistance,
		"M", "nautical mile", "nautical miles",
		"formerly one minute (1/60th of a degree) of latitude" +
			" around any line of latitude, now formally defined" +
			" to have the value given here.",
	},
	"nautical-mile (US)": {
		0, 0, 6080.2 * footToMetre,
		UnitOfDistance,
		"M", "nautical mile (US)", "nautical miles (US)",
		"one minute (1/60th of a degree) of latitude around" +
			" the Clarke 1866 ellipsoid (6080.210 feet) to 5" +
			" significant figures.",
	},
	"nautical-mile (Admiralty/UK)": {
		0, 0, 6080 * footToMetre,
		UnitOfDistance,
		"M", "Admiralty nautical mile", "Admiralty nautical miles",
		"an imperial measure of distance equal to" +
			" one minute (1/60th of a degree) of latitude around" +
			" the Clarke 1866 ellipsoid (6080.210 feet) to 4" +
			" significant figures.",
	},
	"admiralty-fathom": {
		0, 0, 6080 * footToMetre / 1000,
		UnitOfDistance,
		"adm ftm", "Admiralty fathom", "Admiralty fathoms",
		"an imperial measure of distance equal to 1/1000th of" +
			" an Admiralty nautical mile (a little over 2 yards).",
	},

	// regional
	"scots mile": {
		0, 0, yardToMetre * 1976,
		UnitOfDistance,
		"Scots mile", "Scots mile", "Scots miles",
		"a measure of distance used in Scotland (slightly" +
			" longer than the English mile), it has not had" +
			" any formal use for over a century.",
	},
	"irish mile": {
		0, 0, yardToMetre * 2240,
		UnitOfDistance,
		"Irish mile", "Irish mile", "Irish miles",
		"a measure of distance used in Ireland (slightly" +
			" longer than the English mile), no longer" +
			" commonly used.",
	},

	// astronomical
	"astro-unit": {
		0, 0, 1.49597871e11,
		UnitOfDistance,
		"AU", "Astronomical Unit", "Astronomical Units",
		"the mean distance from the centre of the Earth" +
			" to the centre of the sun.",
	},
	"parsec": {
		0, 0, 3.08567758767931e16,
		UnitOfDistance,
		"pc", "parsec", "parsecs",
		"the distance at which one astronomical unit" +
			" subtends an angle of one second of arc.",
	},
	"light-year": {
		0, 0, lightSecond * 86400 * 365.25,
		UnitOfDistance,
		"ly", "light year", "light years",
		"the distance travelled by light in a vacuum" +
			" in one year (365.25 days).",
	},
	"light-second": {
		0, 0, lightSecond,
		UnitOfDistance,
		"ls", "light second", "light seconds",
		"the distance travelled by light in a vacuum" +
			" in one second.",
	},
	"phoot": {
		0, 0, lightSecond * 1e-9,
		UnitOfDistance,
		"phoot", "phoot", "pheet",
		"the distance travelled by light in a vacuum" +
			" in one nanosecond. Also known as Mermin's foot.",
	},

	// printing
	"point": {
		0, 0, pointToMetre,
		UnitOfDistance,
		"pt", "point", "points",
		"a printing term.",
	},
	"pica": {
		0, 0, 12 * pointToMetre,
		UnitOfDistance,
		"pc", "pica", "pica",
		"a printing term (12 points).",
	},

	// Colloquial
	"20ft shipping container": {
		0, 0, (19 * footToMetre) + (10.5 * inchToMetre),
		UnitOfDistance,
		"TEU",
		"length of a 20ft shipping container",
		"lengths of a 20ft shipping container",
		"a colloquial term referencing the length of a standard 20 foot" +
			" shipping container" +
			" presumably to give an impression of scale." +
			" The TEU abbreviation stands for Twenty foot Equivalent Unit." +
			" Note that a 20 foot shipping container is not 20 feet long" +
			" (but a 40 foot container is 40 feet long).",
	},
	"bus": {
		0, 0, 8.38,
		UnitOfDistance,
		"bus",
		"length of a bus",
		"lengths of a bus",
		"a colloquial term referencing the length of a bus." +
			" The bus used for this value is" +
			" an AEC Routemaster standard, double-decker bus" +
			" as used by London Transport between 1956 and 2005.",
	},
	"Eiffel Tower": {
		0, 0, 324,
		UnitOfDistance,
		"Eiffel",
		"Eiffel Tower",
		"Eiffel Towers",
		"a colloquial term referencing the height of the Eiffel Tower." +
			" The value given is to the tip of the tower.",
	},
}

var distanceAliases = map[string]Alias{
	"perch": {"rod", "sometimes also a unit of area or volume"},
	"pole":  {"rod", ""},

	"AU":          {"astro-unit", ""},
	"parsecs":     {"parsec", "plural"},
	"light-years": {"light-year", "plural"},
	"light year":  {"light-year", "no hyphen"},
	"light years": {"light-year", "no hyphen, plural"},

	"m":      {"metre", ""},
	"metres": {"metre", "plural"},
	"meter":  {"metre", "US spelling"},
	"meters": {"metre", "US spelling, plural"},

	"kilometre":  {"km", ""},
	"kilometres": {"km", "plural"},
	"kilometer":  {"km", "US spelling"},
	"kilometers": {"km", "US spelling, plural"},
	"klick":      {"km", "US Army term for kilometre"},
	"klicks":     {"km", "US Army term for kilometre, plural"},

	"inches":   {"inch", "plural"},
	"feet":     {"foot", "plural"},
	"yards":    {"yard", "plural"},
	"miles":    {"mile", "plural"},
	"furlongs": {"furlong", "plural"},
	"leagues":  {"league", "plural"},
}
