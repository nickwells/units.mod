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
			"sec":     "",
			"seconds": "plural",
			"secs":    "plural",
		},
		"", "",
	},

	"ysec": {
		0, 0, y, timeFamily, "ysec", "yoctosecond", "yoctoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"zsec": {
		0, 0, z, timeFamily, "zsec", "zeptosecond", "zeptoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"asec": {
		0, 0, a, timeFamily, "asec", "attosecond", "attoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"fsec": {
		0, 0, f, timeFamily, "fsec", "femtosecond", "femtoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"psec": {
		0, 0, p, timeFamily, "psec", "picosecond", "picoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"nsec": {
		0, 0, n, timeFamily, "nsec", "nanosecond", "nanoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"usec": {
		0, 0, u, timeFamily, "μsec", "microsecond", "microseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"msec": {
		0, 0, m, timeFamily, "msec", "millisecond", "milliseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"csec": {
		0, 0, c, timeFamily, "csec", "centisecond", "centiseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"dsec": {
		0, 0, d, timeFamily, "dsec", "decisecond", "deciseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"minute": {
		0, 0, minToSec, timeFamily, "min", "minute", "minutes", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"min":     "",
			"minutes": "plural",
			"mins":    "plural",
		},
		"", "",
	},
	"hour": {
		0, 0, hourToSec, timeFamily, "hr", "hour", "hours", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"hr":    "",
			"hours": "plural",
			"hrs":   "plural",
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
	"Julian year": {
		0, 0, 365.25 * dayToSec,
		timeFamily,
		"Julian Year", "Julian Year", "Julian Years",
		"365.25 days. A non-SI unit for use in astronomy",
		[]Tag{TagAstro},
		map[string]string{},
		"", "",
	},
	"Gregorian year": {
		0, 0, 365.2425 * dayToSec,
		timeFamily,
		"year", "Gregorian Year", "Gregorian Years",
		"Introduced by Pope Gregory XIII in October 1582." +
			" Commonly used as the basis of the Gregorian calendar.",
		[]Tag{TagColloquial},
		map[string]string{
			"year":  "",
			"years": "plural",
		},
		"", "",
	},
	"Sidereal year": {
		0, 0, 365.256363004 * dayToSec,
		timeFamily,
		"Sidereal Year", "Sidereal Year", "Sidereal Years",
		"The time taken for the Earth to orbit the Sun" +
			" with respect to the fixed stars. Note that this is" +
			" different from the time taken for the Earth to" +
			" complete a full cycle of seasons (the Tropical year)",
		[]Tag{TagAstro},
		map[string]string{},
		"", "",
	},
	"Tropical year": {
		0, 0, 365.24219 * dayToSec,
		timeFamily,
		"Tropical Year", "Tropical Year", "Tropical Years",
		"The time taken for the Earth to complete a full cycle of seasons." +
			" This is the target that the Julian and Gregorian years are" +
			" trying to hit with increasing degrees of accuracy",
		[]Tag{TagAstro},
		map[string]string{},
		"", "",
	},
	"dasec": {
		0, 0, da, timeFamily, "dasec", "decasecond", "decaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"hsec": {
		0, 0, h, timeFamily, "hsec", "hectosecond", "hectoseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"ksec": {
		0, 0, k, timeFamily, "ksec", "kilosecond", "kiloseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Msec": {
		0, 0, _M, timeFamily, "Msec", "megasecond", "megaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Gsec": {
		0, 0, _G, timeFamily, "Gsec", "gigasecond", "gigaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Tsec": {
		0, 0, _T, timeFamily, "Tsec", "terasecond", "teraseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Psec": {
		0, 0, _P, timeFamily, "Psec", "petasecond", "petaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Esec": {
		0, 0, _E, timeFamily, "Esec", "exasecond", "exaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Zsec": {
		0, 0, _Z, timeFamily, "Zsec", "zettasecond", "zettaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Ysec": {
		0, 0, _Y, timeFamily, "Ysec", "yottasecond", "yottaseconds", "",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
}
