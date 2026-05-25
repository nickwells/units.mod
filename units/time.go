package units

// bunTime is the base unit name for time
const bunTime = "second"

// timeFamily represents the collection of units of time
var timeFamily = &Family{
	baseUnitName: bunTime,
	description:  "unit of time",
	name:         Time,
}

// TimeNames maps names to units of time
var timeNames = map[string]Unit{
	bunTime: {
		0, 0, 1,
		timeFamily,
		"sec", bunTime, "seconds",
		"",
		[]Tag{TagSI},
		map[string]string{
			"sec":     "abbreviated",
			"seconds": "plural",
			"secs":    "abbreviated plural",
		},
		"", "",
	},

	"ysec": {
		0, 0, y, timeFamily, "ysec", "yoctosecond", "yoctoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"yoctosecond":  "unabbreviated",
			"yoctoseconds": "unabbreviated plural",
			"yoctosec":     "abbreviated",
			"yoctosecs":    "abbreviated plural",
		},
		"", "",
	},
	"zsec": {
		0, 0, z, timeFamily, "zsec", "zeptosecond", "zeptoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"zeptosecond":  "unabbreviated",
			"zeptoseconds": "unabbreviated plural",
			"zeptosec":     "abbreviated",
			"zeptosecs":    "abbreviated plural",
		},
		"", "",
	},
	"asec": {
		0, 0, a, timeFamily, "asec", "attosecond", "attoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"attosecond":  "unabbreviated",
			"attoseconds": "unabbreviated plural",
			"attosec":     "abbreviated",
			"attosecs":    "abbreviated plural",
		},
		"", "",
	},
	"fsec": {
		0, 0, f, timeFamily, "fsec", "femtosecond", "femtoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"femtosecond":  "unabbreviated",
			"femtoseconds": "unabbreviated plural",
			"femtosec":     "abbreviated",
			"femtosecs":    "abbreviated plural",
		},
		"", "",
	},
	"psec": {
		0, 0, p, timeFamily, "psec", "picosecond", "picoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"picosecond":  "unabbreviated",
			"picoseconds": "unabbreviated plural",
			"picosec":     "abbreviated",
			"picosecs":    "abbreviated plural",
		},
		"", "",
	},
	"nsec": {
		0, 0, n, timeFamily, "nsec", "nanosecond", "nanoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"nanosecond":  "unabbreviated",
			"nanoseconds": "unabbreviated plural",
			"nanosec":     "abbreviated",
			"nanosecs":    "abbreviated plural",
		},
		"", "",
	},
	"usec": {
		0, 0, u, timeFamily, "μsec", "microsecond", "microseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"microsecond":  "unabbreviated",
			"microseconds": "unabbreviated plural",
			"microsec":     "abbreviated",
			"microsecs":    "abbreviated plural",
		},
		"", "",
	},
	"msec": {
		0, 0, m, timeFamily, "msec", "millisecond", "milliseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"millisecond":  "unabbreviated",
			"milliseconds": "unabbreviated plural",
			"millisec":     "abbreviated",
			"millisecs":    "abbreviated plural",
		},
		"", "",
	},
	"csec": {
		0, 0, c, timeFamily, "csec", "centisecond", "centiseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"centisecond":  "unabbreviated",
			"centiseconds": "unabbreviated plural",
			"centisec":     "abbreviated",
			"centisecs":    "abbreviated plural",
		},
		"", "",
	},
	"dsec": {
		0, 0, d, timeFamily, "dsec", "decisecond", "deciseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"decisecond":  "unabbreviated",
			"deciseconds": "unabbreviated plural",
			"decisec":     "abbreviated",
			"decisecs":    "abbreviated plural",
		},
		"", "",
	},
	"minute": {
		0, 0, minToSec, timeFamily, "min", "minute", "minutes", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"min":     "abbreviated",
			"minutes": "plural",
			"mins":    "abbreviated plural",
		},
		"", "",
	},
	"hour": {
		0, 0, hourToSec, timeFamily, "hr", "hour", "hours", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"hr":    "abbreviated",
			"hours": "plural",
			"hrs":   "abbreviated plural",
		},
		"", "",
	},
	"day": {
		0, 0, dayToSec,
		timeFamily,
		"day", "day", "days",
		"An Ephemeris day." +
			" A solar day varies in length and tends to lengthen" +
			" due to tidal effects.",
		[]Tag{TagColloquial},
		map[string]string{
			"days": "plural",
		},
		"", "",
	},
	"week": {
		0, 0, dayToSec * 7,
		timeFamily,
		"week", "week", "weeks", "",
		[]Tag{TagColloquial},
		map[string]string{
			"weeks": "plural",
		},
		"", "",
	},
	"fortnight": {
		0, 0, dayToSec * 14,
		timeFamily,
		"fortnight", "fortnight", "fortnights", "two weeks",
		[]Tag{TagColloquial},
		map[string]string{
			"fortnights": "plural",
		},
		"", "",
	},
	"lunar month": {
		0, 0, dayToSec * 28,
		timeFamily,
		"lunar month", "lunar month", "lunar months",
		"Four weeks." +
			" In English common law a lunar month was taken to be" +
			" precisely 28 days." +
			"\n\n" +
			"For the average time between syzygies of" +
			" the Earth, Moon and Sun, see 'lunation'.",
		[]Tag{TagColloquial},
		map[string]string{
			"lunar months": "plural",
			"lunar-month":  "hyphenated",
			"lunar-months": "hyphenated, plural",
			"Lunar month":  "capitalised",
			"Lunar months": "capitalised, plural",
		},
		"", "",
	},
	"lunation": {
		0, 0, dayToSec * 29.530_588_861,
		timeFamily,
		"lunation", "lunation", "lunations",
		"The average time between successive syzygies (alignments) of the" +
			" Earth, Moon and Sun." +
			" This is the time between phases of the Moon," +
			" for example full moon to full moon." +
			" Because the appearance of the Moon depends on its position" +
			" with respect to the Sun, a lunation (or synodic month)" +
			" is longer by just over two days than the sidereal month" +
			"\n\n" +
			"Note that this is the average value;" +
			" there can be up to seven hours variation around the mean" +
			" in any given year." +
			" Note also that this value will drift" +
			" over time, lengthening by about a 50th of a second per century.",
		[]Tag{TagColloquial},
		map[string]string{
			"lunations":      "plural",
			"synodic month":  "alternative name",
			"synodic months": "alternative name, plural",
			"synodic-month":  "alternative name, hyphenated",
			"synodic-months": "alternative name, hyphenated, plural",
			"Lunation":       "capitalised",
			"Lunations":      "capitalised, plural",
			"Synodic month":  "capitalised, alternative name",
			"Synodic months": "capitalised, alternative name, plural",
		},
		"", "",
	},
	"sidereal month": {
		0, 0, dayToSec * 27.321_661_554,
		timeFamily,
		"sidereal month", "sidereal month", "sidereal months",
		"The period of the Moon's orbit as defined with respect to the" +
			" apparently fixed stars. It is the time it takes for the Moon" +
			" to return to a similar position among the stars." +
			"\n\n" +
			"Note that this value will drift over time," +
			" lengthening by about a 50th of a second per century.",
		[]Tag{TagColloquial},
		map[string]string{
			"sidereal months": "plural",
			"sidereal-month":  "hyphenated",
			"sidereal-months": "hyphenated, plural",
			"Sidereal month":  "capitalised",
			"Sidereal months": "capitalised plural",
		},
		"", "",
	},
	"draconic month": {
		0, 0, dayToSec * 27.212_220_815,
		timeFamily,
		"draconic month", "draconic month", "draconic months",
		"The period of the Moon's orbit as defined with respect to the" +
			" transitions of the ecliptic plane." +
			"\n\n" +
			"The orbit of the Moon lies in a plane inclined by about 5.14°" +
			" to the ecliptic plane." +
			" This orbit crosses the ecliptic plane at two points" +
			" called lunar nodes (the ascending and descending nodes)" +
			"\n\n" +
			"A draconic month is the average of the time between" +
			" two successive transits of the ecliptic plane" +
			" through the same node." +
			"\n\n" +
			"The name 'draconic' refers to a mythical dragon," +
			" said to live at the lunar nodes" +
			" and eat the Sun or Moon during an eclipse." +
			" Note that an eclipse is only possible" +
			" when the Moon is at one of these nodes as" +
			" only then can the Earth, Sun and Moon be in direct line" +
			" and able to eclipse one another." +
			"\n\n" +
			"Note that this value will drift over time," +
			" lengthening by about a third of a second per century.",
		[]Tag{TagColloquial},
		map[string]string{
			"draconic months":   "plural",
			"draconic-month":    "hyphenated",
			"draconic-months":   "hyphenated, plural",
			"Draconic month":    "capitalised",
			"Draconic months":   "capitalised plural",
			"draconitic month":  "alternative",
			"draconitic months": "alternative,plural",
			"draconitic-month":  "alternative,hyphenated",
			"draconitic-months": "alternative,hyphenated, plural",
			"Draconitic month":  "alternative,capitalised",
			"Draconitic months": "alternative,capitalised plural",
			"nodal month":       "alternative",
			"nodal months":      "alternative,plural",
			"nodal-month":       "alternative,hyphenated",
			"nodal-months":      "alternative,hyphenated,plural",
			"nodical month":     "alternative",
			"nodical months":    "alternative,plural",
			"nodical-month":     "alternative,hyphenated",
			"nodical-months":    "alternative,hyphenated,plural",
		},
		"", "",
	},
	"Julian year": {
		0, 0, 365.25 * dayToSec,
		timeFamily,
		"Julian Year", "Julian Year", "Julian Years",
		"365.25 days. A non-SI unit for use in astronomy",
		[]Tag{TagAstro},
		map[string]string{
			"Julian years": "plural",
			"Julian Year":  "capital year",
			"Julian-year":  "hyphenated",
			"Julian-years": "hyphenated, plural",
			"julian year":  "lowercase",
			"julian years": "lowercase, plural",
			"julian-year":  "lowercase, hyphenated",
			"julian-years": "lowercase, hyphenated, plural",
		},
		"", "",
	},
	"Gregorian year": {
		0, 0, daysPerGregorianYear * dayToSec,
		timeFamily,
		"year", "Gregorian Year", "Gregorian Years",
		"Introduced by Pope Gregory XIII in October 1582." +
			" Commonly used as the basis of the Gregorian calendar.",
		[]Tag{TagColloquial},
		map[string]string{
			"year":           "",
			"years":          "plural",
			"calendar year":  "alternative",
			"calendar years": "alternative, plural",
			"calendar-year":  "alternative, hyphenated",
			"calendar-years": "alternative, hyphenated, plural",
			"Gregorian Year": "capital year",
		},
		"", "",
	},
	"Sidereal year": {
		0, 0, 365.256_363_004 * dayToSec,
		timeFamily,
		"Sidereal Year", "Sidereal Year", "Sidereal Years",
		"The time taken for the Earth to orbit the Sun" +
			" with respect to the fixed stars. Note that this is" +
			" different from the time taken for the Earth to" +
			" complete a full cycle of seasons (the Tropical year)",
		[]Tag{TagAstro},
		map[string]string{
			"Sidereal years": "plural",
			"Sidereal Year":  "capital year",
			"sidereal year":  "lowercase",
			"sidereal years": "lowercase, plural",
			"sidereal-year":  "lowercase, hyphenated",
			"sidereal-years": "lowercase, hyphenated, plural",
		},
		"", "",
	},
	"Tropical year": {
		0, 0, 365.242_19 * dayToSec,
		timeFamily,
		"Tropical Year", "Tropical Year", "Tropical Years",
		"The time taken for the Earth to complete a full cycle of seasons." +
			" This is the target that the Julian and Gregorian years are" +
			" trying to hit with increasing degrees of accuracy",
		[]Tag{TagAstro},
		map[string]string{
			"Tropical years": "plural",
			"Tropical Year":  "capital year",
			"tropical year":  "lowercase",
			"tropical years": "lowercase, plural",
			"tropical-year":  "lowercase, hyphenated",
			"tropical-years": "lowercase, hyphenated, plural",
		},
		"", "",
	},
	"century": {
		0, 0, 100 * daysPerGregorianYear * dayToSec,
		timeFamily,
		"century", "century", "centuries",
		"one hundred Gregorian years." +
			" The name derives from the Latin word for a hundred",
		[]Tag{TagColloquial},
		map[string]string{
			"centuries": "plural",
		},
		"", "",
	},
	"millennium": {
		0, 0, 1000 * daysPerGregorianYear * dayToSec,
		timeFamily,
		"millennium", "millennium", "millennia",
		"one thousand Gregorian years." +
			" The name derives from the Latin words" +
			` "mille" (a thousand) and "annum" (a year)`,
		[]Tag{TagColloquial},
		map[string]string{
			"millennia": "plural",
		},
		"", "",
	},
	"aeon": {
		0, 0, 1e9 * daysPerGregorianYear * dayToSec,
		timeFamily,
		"aeon", "aeon", "aeons",
		"one billion Gregorian years." +
			" The name comes from the Greek where it means" +
			" one hundred years but in astronomy (or geology or cosmology)" +
			" it has the meaning used here." +
			" Colloquially it just means a very long time",
		[]Tag{TagAstro},
		map[string]string{
			"aeons": "plural",
			"eon":   "US or Australian",
			"eons":  "US or Australian plural",
		},
		"", "",
	},
	"dasec": {
		0, 0, da, timeFamily, "dasec", "decasecond", "decaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"decasecond":  "unabbreviated",
			"decaseconds": "unabbreviated plural",
			"decasec":     "abbreviated",
			"decasecs":    "abbreviated plural",
		},
		"", "",
	},
	"hsec": {
		0, 0, h, timeFamily, "hsec", "hectosecond", "hectoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"hectosecond":  "unabbreviated",
			"hectoseconds": "unabbreviated plural",
			"hectosec":     "abbreviated",
			"hectosecs":    "abbreviated plural",
		},
		"", "",
	},
	"ksec": {
		0, 0, k, timeFamily, "ksec", "kilosecond", "kiloseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"kilosecond":  "unabbreviated",
			"kiloseconds": "unabbreviated plural",
			"kilosec":     "abbreviated",
			"kilosecs":    "abbreviated plural",
		},
		"", "",
	},
	"Msec": {
		0, 0, _M, timeFamily, "Msec", "megasecond", "megaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"megasecond":  "unabbreviated",
			"megaseconds": "unabbreviated plural",
			"megasec":     "abbreviated",
			"megasecs":    "abbreviated plural",
		},
		"", "",
	},
	"Gsec": {
		0, 0, _G, timeFamily, "Gsec", "gigasecond", "gigaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"gigasecond":  "unabbreviated",
			"gigaseconds": "unabbreviated plural",
			"gigasec":     "abbreviated",
			"gigasecs":    "abbreviated plural",
		},
		"", "",
	},
	"Tsec": {
		0, 0, _T, timeFamily, "Tsec", "terasecond", "teraseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"terasecond":  "unabbreviated",
			"teraseconds": "unabbreviated plural",
			"terasec":     "abbreviated",
			"terasecs":    "abbreviated plural",
		},
		"", "",
	},
	"Psec": {
		0, 0, _P, timeFamily, "Psec", "petasecond", "petaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"petasecond":  "unabbreviated",
			"petaseconds": "unabbreviated plural",
			"petasec":     "abbreviated",
			"petasecs":    "abbreviated plural",
		},
		"", "",
	},
	"Esec": {
		0, 0, _E, timeFamily, "Esec", "exasecond", "exaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"exasecond":  "unabbreviated",
			"exaseconds": "unabbreviated plural",
			"exasec":     "abbreviated",
			"exasecs":    "abbreviated plural",
		},
		"", "",
	},
	"Zsec": {
		0, 0, _Z, timeFamily, "Zsec", "zettasecond", "zettaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"zettasecond":  "unabbreviated",
			"zettaseconds": "unabbreviated plural",
			"zettasec":     "abbreviated",
			"zettasecs":    "abbreviated plural",
		},
		"", "",
	},
	"Ysec": {
		0, 0, _Y, timeFamily, "Ysec", "yottasecond", "yottaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"yottasecond":  "unabbreviated",
			"yottaseconds": "unabbreviated plural",
			"yottasec":     "abbreviated",
			"yottasecs":    "abbreviated plural",
		},
		"", "",
	},
}
