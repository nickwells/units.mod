//nolint:mnd
package units

// bunDistance is the base unit name for distance
const bunDistance = "metre"

// distanceFamily represents the base unit of distance
var distanceFamily = &Family{
	baseUnitName:  bunDistance,
	description:   "unit of distance",
	name:          Distance,
	familyAliases: []string{"length", "len"},
}

// DistanceNames maps names to units of distance
var distanceNames = map[string]Unit{
	bunDistance: {
		0, 0, 1,
		distanceFamily,
		"m", bunDistance, "metres",
		"The base unit of distance in the SI (metric) system.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"m":      "",
			"metres": "plural",
			"meter":  "US spelling",
			"meters": "US spelling, plural",
		},
		"", "",
	},
	// metric
	"ym": {
		0, 0, y,
		distanceFamily,
		"ym", "yoctometre", "yoctometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"zm": {
		0, 0, z,
		distanceFamily,
		"zm", "zeptometre", "zeptometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"am": {
		0, 0, a,
		distanceFamily,
		"am", "attometre", "attometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"fm": {
		0, 0, f,
		distanceFamily,
		"fm", "femtometre", "femtometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"pm": {
		0, 0, p,
		distanceFamily,
		"pm", "picometre", "picometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"nm": {
		0, 0, n,
		distanceFamily,
		"nm", "nanometre", "nanometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"um": {
		0, 0, u,
		distanceFamily,
		"µm", "micrometre", "micrometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"mm": {
		0, 0, m,
		distanceFamily,
		"mm", "millimetre", "millimetres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"millimetre":  "",
			"millimeter":  "US spelling",
			"millimetres": "plural",
			"millimeters": "US spelling, plural",
		},
		"", "",
	},
	"cm": {
		0, 0, c,
		distanceFamily,
		"cm", "centimetre", "centimetres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"centimetre":  "",
			"centimeter":  "US spelling",
			"centimetres": "plural",
			"centimeters": "US spelling, plural",
		},
		"", "",
	},
	"dm": {
		0, 0, d,
		distanceFamily,
		"dm", "decimetre", "decimetres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"dam": {
		0, 0, da,
		distanceFamily,
		"dam", "decametre", "decametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"hm": {
		0, 0, h,
		distanceFamily,
		"hm", "hectometre", "hectometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"km": {
		0, 0, k,
		distanceFamily,
		"km", "kilometre", "kilometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"kilometre":  "",
			"kilometres": "plural",
			"kilometer":  "US spelling",
			"kilometers": "US spelling, plural",
			"klick":      "US Army term for kilometre",
			"klicks":     "US Army term for kilometre, plural",
		},
		"", "",
	},
	"mym": {
		0, 0, 10000,
		distanceFamily,
		"mym", "myriametre", "myriametres",
		"an obsolete metric measure of distance.",
		[]Tag{TagSI, TagMetric, TagHist},
		map[string]string{},
		"", "",
	},
	"Mm": {
		0, 0, _M,
		distanceFamily,
		"Mm", "megametre", "megametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Gm": {
		0, 0, _G,
		distanceFamily,
		"Gm", "gigametre", "gigametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Tm": {
		0, 0, _T,
		distanceFamily,
		"Tm", "terametre", "terametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Pm": {
		0, 0, _P,
		distanceFamily,
		"Pm", "petametre", "petametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Em": {
		0, 0, _E,
		distanceFamily,
		"Em", "exametre", "exametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Zm": {
		0, 0, _Z,
		distanceFamily,
		"Zm", "zettametre", "zettametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Ym": {
		0, 0, _Y,
		distanceFamily,
		"Ym", "yottametre", "yottametres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},

	// Imperial / US
	"barleycorn": {
		0, 0, inchToMetre / 3.0,
		distanceFamily,
		"in", "inch", "inches",
		"a non-metric UK measure of distance. This was defined in" +
			" the 'Composition of Yards and Perches' an English" +
			" statute from between 1266-1303.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"inch": {
		0, 0, inchToMetre,
		distanceFamily,
		"in", "inch", "inches",
		"imperial/US measure of distance.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"inches": "plural",
		},
		"", "",
	},
	"hand": {
		0, 0, inchToMetre * 4,
		distanceFamily,
		"hand", "hand", "hands",
		"an imperial measure of length, typically used to measure" +
			" the height of horses.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{},
		"", "",
	},
	"link": {
		0, 0, footToMetre * 66 / 100,
		distanceFamily,
		"lnk", "link", "links",
		"an imperial measure of length, derived from Gunter's chain" +
			" - a metal chain 66 feet long with 100 links." +
			" Edmund Gunter (1581-1626) was an English clergyman," +
			" mathematician, geometer and astronomer.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"foot": {
		0, 0, footToMetre,
		distanceFamily,
		"ft", "foot", "feet",
		"an imperial measure of length, 12 inches.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"feet": "plural",
		},
		"", "",
	},
	"foot (North German)": {
		0, 0, footToMetre * 11.0 / 10.0,
		distanceFamily,
		"ft (North German)", "foot (North German)", "feet (North German)",
		"a medieval measure of length slightly longer than a standard foot.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"foot (Roman)": {
		0, 0, 0.296,
		distanceFamily,
		"ft (Roman)", "foot (Roman)", "feet (Roman)",
		"a Roman measure of length. There are several alternative lengths" +
			" but the one used here is a popular value.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"foot (Parisian)": {
		0, 0, 0.325,
		distanceFamily,
		"ft (Parisian)", "foot (Parisian)", "feet (Parisian)",
		"a French measure of length.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"foot (Amsterdam)": {
		0, 0, 0.283133,
		distanceFamily,
		"ft (Amsterdam)",
		"Amsterdamse voet", "Amsterdamse voeten",
		"a measure of distance used in Amsterdam, no longer" +
			" commonly used. Note that the Amsterdam foot" +
			" had 11 inches (duimen) rather than the more usual 12",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"foot (Rijnland)": {
		0, 0, 0.314,
		distanceFamily,
		"ft (Rijnland)",
		"Rijnlandse voet", "Rijnlandse voeten",
		"a measure of distance used in the Netherlands, no longer" +
			" commonly used. Note that, unlike the Amsterdam foot," +
			" the Rijnland foot had 12 inches (duimen)",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"foot (metric)": {
		0, 0, 0.3,
		distanceFamily,
		"ft (metric)", "foot (metric)", "feet (metric)",
		"a measure of length - ISO 2848.",
		[]Tag{TagColloquial, TagMetric},
		map[string]string{},
		"", "",
	},
	"US survey foot": {
		0, 0, 1200.0 / 3937.0,
		distanceFamily,
		"ft", "foot", "feet",
		"a US Customary unit of length, longer than a standard foot" +
			" by approximately 2 parts in a million.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"Indian survey foot": {
		0, 0, 0.3047996,
		distanceFamily,
		"ft", "foot", "feet",
		"an Indian unit of length.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"yard": {
		0, 0, yardToMetre,
		distanceFamily,
		"yd", "yard", "yards",
		"an imperial measure of length, 3 feet.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"yards": "plural",
		},
		"", "",
	},
	"rod": {
		0, 0, rodToMetre,
		distanceFamily,
		"rod", "rod", "rods",
		"an imperial measure of length, originally equal to 5 yards" +
			" it was redefined as 5.5 yards when England moved from" +
			" the North German to the shorter standard foot in the" +
			" late 13th century" +
			" Note that a rod can also be a measure of area" +
			" under certain circumstances.",
		[]Tag{TagImperial},
		map[string]string{
			"perch": "sometimes also a unit of area or volume",
			"pole":  "",
		},
		"", "",
	},
	"chain": {
		0, 0, rodToMetre * 4,
		distanceFamily,
		"ch", "chain", "chains",
		"an imperial measure of length - four rods.",
		[]Tag{TagImperial},
		map[string]string{},
		"", "",
	},
	"furlong": {
		0, 0, rodToMetre * 40,
		distanceFamily,
		"fur", "furlong", "furlongs",
		"an imperial measure of length - forty rods. Originally a" +
			" furrow's length - the distance a team of oxen" +
			" could plough without resting, later standardised as" +
			" 40 rods.",
		[]Tag{TagImperial},
		map[string]string{
			"furlongs": "plural",
		},
		"", "",
	},
	"mile": {
		0, 0, mileToMetre,
		distanceFamily,
		"mi", "mile", "miles",
		"an imperial measure of length.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"miles": "plural",
		},
		"", "",
	},
	"swimming-mile": {
		0, 0, yardToMetre * 1650,
		distanceFamily,
		"swim-mi", "swimming mile", "swimming miles",
		"a coloquial and ill-defined distance used by swimmers," +
			" sometimes taken to be 1500 metres.",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},
	"league": {
		0, 0, mileToMetre * 3,
		distanceFamily,
		"lea", "league", "leagues",
		"3 miles (perhaps the distance a person could walk in an hour).",
		[]Tag{TagImperial},
		map[string]string{
			"leagues": "plural",
		},
		"", "",
	},
	"nautical league": {
		0, 0, nauticalMileToMetre * 3,
		distanceFamily,
		"lea", "nautical league", "nautical leagues",
		"3 nautical miles.",
		[]Tag{TagImperial, TagNautical},
		map[string]string{},
		"", "",
	},

	"fathom": {
		0, 0, yardToMetre * 2,
		distanceFamily,
		"ftm", "fathom", "fathoms",
		"an imperial measure of distance equal to two yards.",
		[]Tag{TagImperial, TagNautical},
		map[string]string{},
		"", "",
	},
	"nautical-mile": {
		0, 0, nauticalMileToMetre,
		distanceFamily,
		"M", "nautical mile", "nautical miles",
		"formerly one minute (1/60th of a degree) of latitude" +
			" around any line of latitude, now formally defined" +
			" to have the value given here.",
		[]Tag{TagImperial, TagNautical},
		map[string]string{},
		"", "",
	},
	"nautical-mile (US)": {
		0, 0, 6080.2 * footToMetre,
		distanceFamily,
		"M", "nautical mile (US)", "nautical miles (US)",
		"one minute (1/60th of a degree) of latitude around" +
			" the Clarke 1866 ellipsoid (6080.210 feet) to 5" +
			" significant figures.",
		[]Tag{TagUScustomary, TagNautical},
		map[string]string{},
		"", "",
	},
	"nautical-mile (Admiralty/UK)": {
		0, 0, 6080 * footToMetre,
		distanceFamily,
		"M", "Admiralty nautical mile", "Admiralty nautical miles",
		"an imperial measure of distance equal to" +
			" one minute (1/60th of a degree) of latitude around" +
			" the Clarke 1866 ellipsoid (6080.210 feet) to 4" +
			" significant figures.",
		[]Tag{TagImperial, TagNautical},
		map[string]string{},
		"", "",
	},
	"admiralty-fathom": {
		0, 0, 6080 * footToMetre / 1000,
		distanceFamily,
		"adm ftm", "Admiralty fathom", "Admiralty fathoms",
		"an imperial measure of distance equal to 1/1000th of" +
			" an Admiralty nautical mile (a little over 2 yards).",
		[]Tag{TagImperial, TagNautical},
		map[string]string{},
		"", "",
	},

	// regional
	"scots mile": {
		0, 0, yardToMetre * 1976,
		distanceFamily,
		"Scots mile", "Scots mile", "Scots miles",
		"a measure of distance used in Scotland (slightly" +
			" longer than the English mile), it has not had" +
			" any formal use for over a century.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},
	"irish mile": {
		0, 0, yardToMetre * 2240,
		distanceFamily,
		"Irish mile", "Irish mile", "Irish miles",
		"a measure of distance used in Ireland (slightly" +
			" longer than the English mile), no longer" +
			" commonly used.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},

	// astronomical
	"astro-unit": {
		0, 0, 1.495978707e11,
		distanceFamily,
		"AU", "Astronomical Unit", "Astronomical Units",
		"the mean distance from the centre of the Earth" +
			" to the centre of the sun. In 2012 it was" +
			" defined to have the value given here",
		[]Tag{TagAstro},
		map[string]string{
			"AU":                 "abbreviation",
			"au":                 "abbreviation",
			"astronomical-unit":  "full name",
			"astronomical-units": "full name, plural",
			"astronomical unit":  "full name, no hyphen",
			"astronomical units": "full name, no hyphen, plural",
		},
		"", "",
	},
	"parsec": {
		0, 0, 3.08567758767931e16,
		distanceFamily,
		"pc", "parsec", "parsecs",
		"the distance at which one astronomical unit" +
			" subtends an angle of one second of arc.",
		[]Tag{TagAstro},
		map[string]string{
			"parsecs": "plural",
		},
		"", "",
	},
	"light-year": {
		0, 0, lightSecond * dayToSec * 365.25,
		distanceFamily,
		"ly", "light year", "light years",
		"the distance travelled by light in a vacuum" +
			" in one year (365.25 days).",
		[]Tag{TagAstro},
		map[string]string{
			"light-years": "plural",
			"light year":  "no hyphen",
			"light years": "no hyphen, plural",
		},
		"", "",
	},
	"light-second": {
		0, 0, lightSecond,
		distanceFamily,
		"ls", "light second", "light seconds",
		"the distance travelled by light in a vacuum" +
			" in one second.",
		[]Tag{TagAstro},
		map[string]string{},
		"", "",
	},
	"phoot": {
		0, 0, lightSecond * 1e-9,
		distanceFamily,
		"phoot", "phoot", "pheet",
		"the distance travelled by light in a vacuum" +
			" in one nanosecond. Also known as Mermin's foot.",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},

	// printing
	"point": {
		0, 0, pointToMetre,
		distanceFamily,
		"pt", "point", "points",
		"a printing term.",
		[]Tag{TagPrint},
		map[string]string{},
		"", "",
	},
	"pica": {
		0, 0, 12 * pointToMetre,
		distanceFamily,
		"pc", "pica", "pica",
		"a printing term (12 points).",
		[]Tag{TagPrint},
		map[string]string{},
		"", "",
	},

	// Colloquial
	"20ft shipping container": {
		0, 0, (19 * footToMetre) + (10.5 * inchToMetre),
		distanceFamily,
		"TEU",
		"length of a 20ft shipping container",
		"lengths of a 20ft shipping container",
		"a colloquial term referencing the length of a standard 20 foot" +
			" shipping container" +
			" presumably to give an impression of scale." +
			" The TEU abbreviation stands for Twenty foot Equivalent Unit." +
			" Note that a 20 foot shipping container is not 20 feet long" +
			" (but a 40 foot container is 40 feet long).",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},
	"bus": {
		0, 0, 8.38,
		distanceFamily,
		"bus",
		"length of a bus",
		"lengths of a bus",
		"a colloquial term referencing the length of a bus." +
			" The bus used for this value is" +
			" an AEC Routemaster standard, double-decker bus" +
			" as used by London Transport between 1956 and 2005.",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},
	"Eiffel Tower": {
		0, 0, 324,
		distanceFamily,
		"Eiffel",
		"Eiffel Tower",
		"Eiffel Towers",
		"a colloquial term referencing the height of the Eiffel Tower." +
			" The value given is to the tip of the tower.",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},
}
