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
		"The base unit of distance in the SI (metric) system." +
			" It was originally defined in 1791 by the" +
			" French National Assembly to be one ten millionth" +
			" of the distance from the equator to the" +
			" North Pole along a great circle passing through Paris." +
			"\n\n" +
			"Since 2019 it has been" +
			" defined in terms of the speed of light meaning" +
			" that the metre is measured in terms of the" +
			" speed of light not the other way round.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"m":      "",
			"metres": "plural",
			"meter":  "US spelling",         //nolint:misspell
			"meters": "US spelling, plural", //nolint:misspell
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
		map[string]string{
			"nanometre":  "",
			"nanometres": "plural",
			"nanometer":  "US spelling",         //nolint:misspell
			"nanometers": "US spelling, plural", //nolint:misspell
		},
		"", "",
	},
	"um": {
		0, 0, u,
		distanceFamily,
		"µm", "micrometre", "micrometres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"micrometre":  "",
			"micrometres": "plural",
			"micrometer":  "US spelling",         //nolint:misspell
			"micrometers": "US spelling, plural", //nolint:misspell
		},
		"", "",
	},
	"mm": {
		0, 0, m,
		distanceFamily,
		"mm", "millimetre", "millimetres", "a metric measure of distance.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"millimetre":  "",
			"millimetres": "plural",
			"millimeter":  "US spelling",         //nolint:misspell
			"millimeters": "US spelling, plural", //nolint:misspell
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
			"centimetres": "plural",
			"centimeter":  "US spelling",         //nolint:misspell
			"centimeters": "US spelling, plural", //nolint:misspell
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
			"kilometer":  "US spelling",         //nolint:misspell
			"kilometers": "US spelling, plural", //nolint:misspell
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
		map[string]string{
			"myriametre":  "",
			"myriametres": "plural",
		},
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
		"barleycorn", "barleycorn", "barleycorns",
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
			" there being 100 links in the chain.",
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
		map[string]string{
			"North German foot": "expanded",
			"North German feet": "expanded, plural",
		},
		"", "",
	},
	"foot (Roman)": {
		0, 0, 0.296,
		distanceFamily,
		"ft (Roman)", "foot (Roman)", "feet (Roman)",
		"a Roman measure of length. There are several alternative lengths" +
			" but the one used here is a popular value.",
		[]Tag{TagHist},
		map[string]string{
			"Roman foot": "expanded",
			"Roman feet": "expanded, plural",
		},
		"", "",
	},
	"foot (Parisian)": {
		0, 0, 0.325,
		distanceFamily,
		"ft (Parisian)", "foot (Parisian)", "feet (Parisian)",
		"a French measure of length.",
		[]Tag{TagHist},
		map[string]string{
			"Parisian foot": "expanded",
			"Parisian feet": "expanded, plural",
		},
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
		map[string]string{
			"Amsterdam foot": "expanded",
			"Amsterdam feet": "expanded, plural",
		},
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
		map[string]string{
			"Rijnland foot": "expanded",
			"Rijnland feet": "expanded, plural",
		},
		"", "",
	},
	"foot (metric)": {
		0, 0, 0.3,
		distanceFamily,
		"ft (metric)", "foot (metric)", "feet (metric)",
		"a measure of length - ISO 2848.",
		[]Tag{TagColloquial, TagMetric},
		map[string]string{
			"metric foot": "expanded",
			"metric feet": "expanded, plural",
			"metric-foot": "expanded, hyphenated",
			"metric-feet": "expanded, hyphenated, plural",
		},
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
			" late 13th century. It was standardised as a survey measure" +
			" by Edmund Gunter in 1607 as a quarter of a chain." +
			"\n\n" +
			"It is also known as a perch or pole and in some English" +
			" dialects as a lug." +
			"\n\n" +
			"Note that a rod can also be a measure of area" +
			" under certain circumstances.",
		[]Tag{TagImperial},
		map[string]string{
			"perch": "sometimes also a unit of area or volume",
			"pole":  "",
			"poles": "plural",
			"rods":  "plural",
		},
		"", "",
	},
	"chain": {
		0, 0, rodToMetre * 4,
		distanceFamily,
		"ch", "chain", "chains",
		"an imperial measure of length - four rods." +
			"\n\n" +
			"Gunter's chain was" +
			" a metal chain 66 feet long with 100 links." +
			" Edmund Gunter (1581-1626) was an English clergyman," +
			" mathematician, geometer and astronomer.",
		[]Tag{TagImperial},
		map[string]string{
			"Gunters chain":  "alternative",
			"Gunters chains": "alternative, plural",
			"gunters chain":  "alternative, lowercase",
			"gunters chains": "alternative, lowercase, plural",
			"Gunters-chain":  "alternative, hyphenated",
			"Gunters-chains": "alternative, hyphenated, plural",
			"gunters-chain":  "alternative, hyphenated, lowercase",
			"gunters-chains": "alternative, hyphenated, lowercase, plural",
		},
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

	// nautical
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
		map[string]string{
			"fathoms": "plural",
		},
		"", "",
	},
	"cable": {
		0, 0, nauticalMileToMetre / 10,
		distanceFamily,
		"cable", "cable", "cables",
		"an imperial measure of distance equal to" +
			" one tenth of a nautical mile." +
			" Historically it derives from the length of" +
			" a ship's anchor cable.",
		[]Tag{TagImperial, TagNautical},
		map[string]string{
			"cables": "plural",
		},
		"", "",
	},
	"nautical-mile": {
		0, 0, nauticalMileToMetre,
		distanceFamily,
		"M", "nautical mile", "nautical miles",
		"formerly one minute (1/60th of a degree) of latitude" +
			" around any line of latitude. The metre was originally" +
			" defined (in 1791) to be one ten-millionth of" +
			" a quarter meridian." +
			" This gave a value for a nautical mile of" +
			" 10,000,000/(90*60) or 1851.85 metres." +
			" In 1929 it was formally defined at the" +
			" First International Extraordinary Hydrographic Conference" +
			" to have the value given here.",
		[]Tag{TagImperial, TagNautical},
		map[string]string{
			"nautical mile":  "",
			"nautical miles": "",
		},
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
		0, 0, astronomicalUnitToMetre,
		distanceFamily,
		"au", "Astronomical Unit", "Astronomical Units",
		"the mean distance from the centre of the Earth" +
			" to the centre of the sun. In 2012 it was" +
			" defined to have the value given here." +
			" The abbreviation used here is as recommended" +
			" by the International Astronomical Union at" +
			" their XXVII General Assembly",
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
		0, 0, parsecToMetre,
		distanceFamily,
		"pc", "parsec", "parsecs",
		"the distance at which one astronomical unit" +
			" subtends an angle of one second of arc." +
			"\n\n" +
			"The term was coined in 1913 by the British astronomer" +
			" Herbert Hall Turner as a contraction of" +
			` "parallax of one arc second".` +
			"\n\n" +
			" In 2015, the International Astronomical Union" +
			" passed resolution B2 which explicitly defined" +
			" the parsec in terms of the astronomical unit." +
			" The definition gave it as being precisely" +
			"\n" +
			"648000/π au",
		[]Tag{TagAstro},
		map[string]string{
			"parsecs": "plural",
			"pc":      "abbreviation",
		},
		"", "",
	},
	"kiloparsec": {
		0, 0, parsecToMetre * k,
		distanceFamily,
		"kpc", "kiloparsec", "kiloparsecs",
		"a thousand parsecs",
		[]Tag{TagAstro},
		map[string]string{
			"kiloparsecs": "plural",
			"kpc":         "abbreviation",
		},
		"", "",
	},
	"megaparsec": {
		0, 0, parsecToMetre * _M,
		distanceFamily,
		"Mpc", "megaparsec", "megaparsecs",
		"a million (10^6) parsecs",
		[]Tag{TagAstro},
		map[string]string{
			"megaparsecs": "plural",
			"Mpc":         "abbreviation",
		},
		"", "",
	},
	"gigaparsec": {
		0, 0, parsecToMetre * _G,
		distanceFamily,
		"Gpc", "gigaparsec", "gigaparsecs",
		"a billion (10^9) parsecs",
		[]Tag{TagAstro},
		map[string]string{
			"gigaparsecs": "plural",
			"Gpc":         "abbreviation",
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
			" in one nanosecond. Also known as Mermin's foot." +
			" It differs from an imperial foot by about 5mm." +
			" It was proposed in 2005 by" +
			" professor Nathaniel David Mermin (born 30 March 1935)," +
			" a solid state physicist",
		[]Tag{TagColloquial},
		map[string]string{
			"pheet":            "plural",
			"Mermins foot":     "alternative",
			"Mermins feet":     "alternative, plural",
			"Mermin's foot":    "alternative",
			"Mermin's feet":    "alternative, plural",
			"mermins foot":     "alternative, lowercase",
			"mermins feet":     "alternative, lowercase, plural",
			"mermins-foot":     "alternative, hyphenated, lowercase",
			"mermins-feet":     "alternative, hyphenated, lowercase, plural",
			"light nanosecond": "alternative",
			"light foot":       "alternative",
		},
		"", "",
	},

	// printing
	"point": {
		0, 0, pointToMetre,
		distanceFamily,
		"pt", "point", "points",
		"a printing term.",
		[]Tag{TagPrint},
		map[string]string{
			"points": "plural",
		},
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
		0, 0, teuLength,
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
		map[string]string{
			"TEU": "20ft equivalent unit",
			"teu": "lowercase",
		},
		"", "",
	},
	"bus": {
		0, 0, 8.38,
		distanceFamily,
		"bus",
		"length of a bus", "lengths of a bus",
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
		"Eiffel Tower", "Eiffel Towers",
		"a colloquial term referencing the height of the Eiffel Tower." +
			" The value given is to the tip of the tower.",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},
	"smoot": {
		0, 0, 1.7018,
		distanceFamily,
		"smoot",
		"smoot", "smoots",
		"a colloquial term referencing the height of Oliver Smoot." +
			" Smoot was a freshman at MIT (a university in Boston, USA)" +
			" and in October 1958, as an initiation challenge for the" +
			" Lambda Chi Alpha fraternity he was tasked with measuring" +
			" the Harvard Bridge using his own height. A plaque on the" +
			" bridge's north side commemorates the event and the Smoot" +
			" markings are periodically repainted." +
			"\n\n" +
			"After leaving university he went on" +
			" to work in standards within the technology sector. He" +
			" served as chairman of the American National Standards" +
			" Institute (ANSI) from 2001 to 2002 and as president of the" +
			" International Organisation for Standards (OSI) from 2003 to" +
			" 2005.",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},
}
