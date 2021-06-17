package units

// UnitOfTime represents the base unit of time
var UnitOfTime = Family{
	BaseUnitName: "second",
	Description:  "unit of time",
}

// SecondUnit is a suitable default value for a UnitOfTime
var SecondUnit = Unit{
	0, 0, 1,
	UnitOfTime,
	"sec", UnitOfTime.BaseUnitName, "seconds", "", "",
}

// TimeNames maps names to units of time
var TimeNames = map[string]Unit{
	"ysec":   {0, 0, y, UnitOfTime, "ysec", "yoctosecond", "yoctoseconds", "", ""},
	"zsec":   {0, 0, z, UnitOfTime, "zsec", "zeptosecond", "zeptoseconds", "", ""},
	"asec":   {0, 0, a, UnitOfTime, "asec", "attosecond", "attoseconds", "", ""},
	"fsec":   {0, 0, f, UnitOfTime, "fsec", "femtosecond", "femtoseconds", "", ""},
	"psec":   {0, 0, p, UnitOfTime, "psec", "picosecond", "picoseconds", "", ""},
	"nsec":   {0, 0, n, UnitOfTime, "nsec", "nanosecond", "nanoseconds", "", ""},
	"usec":   {0, 0, u, UnitOfTime, "μsec", "microsecond", "microseconds", "", ""},
	"msec":   {0, 0, m, UnitOfTime, "msec", "millisecond", "milliseconds", "", ""},
	"csec":   {0, 0, c, UnitOfTime, "csec", "centisecond", "centiseconds", "", ""},
	"dsec":   {0, 0, d, UnitOfTime, "dsec", "decisecond", "deciseconds", "", ""},
	"second": SecondUnit,
	"minute": {0, 0, 60, UnitOfTime, "min", "minute", "minutes", "", ""},
	"hour":   {0, 0, 60 * 60, UnitOfTime, "hr", "hour", "hours", "", ""},
	"day": {
		0, 0, 60 * 60 * 24,
		UnitOfTime,
		"day", "day", "days",
		"An Ephemeris day." +
			" A solar day varies in length and tends to lengthen" +
			" due to tidal effects.",
		"",
	},
	"week": {0, 0, 60 * 60 * 24 * 7, UnitOfTime, "week", "week", "weeks", "", ""},
	"fortnight": {
		0, 0, 60 * 60 * 24 * 14,
		UnitOfTime,
		"fortnight", "fortnight", "fortnights", "",
		"",
	},
	"Julian year": {
		0, 0, 365.25 * 86400,
		UnitOfTime,
		"Julian Year", "Julian Year", "Julian Years",
		"365.25 days. A non-SI unit for use in astronomy",
		"",
	},
	"Gregorian year": {
		0, 0, 365.2425 * 86400,
		UnitOfTime,
		"Gregorian Year", "Gregorian Year", "Gregorian Years",
		"Introduced by Pope Gregory XIII in October 1582." +
			" Commonly used as the basis of the Gregorian calendar.",
		"",
	},
	"Siderial year": {
		0, 0, 365.256363004 * 86400,
		UnitOfTime,
		"Siderial Year", "Siderial Year", "Siderial Years",
		"The time taken for the Earth to orbit the Sun" +
			" with respect to the fixed stars",
		"",
	},
	"dasec": {0, 0, da, UnitOfTime, "dasec", "decasecond", "decaseconds", "", ""},
	"hsec":  {0, 0, h, UnitOfTime, "hsec", "hectosecond", "hectoseconds", "", ""},
	"ksec":  {0, 0, k, UnitOfTime, "ksec", "kilosecond", "kiloseconds", "", ""},
	"Msec":  {0, 0, _M, UnitOfTime, "Msec", "megasecond", "megaseconds", "", ""},
	"Gsec":  {0, 0, _G, UnitOfTime, "Gsec", "gigasecond", "gigaseconds", "", ""},
	"Tsec":  {0, 0, _T, UnitOfTime, "Tsec", "terasecond", "teraseconds", "", ""},
	"Psec":  {0, 0, _P, UnitOfTime, "Psec", "petasecond", "petaseconds", "", ""},
	"Esec":  {0, 0, _E, UnitOfTime, "Esec", "exasecond", "exaseconds", "", ""},
	"Zsec":  {0, 0, _Z, UnitOfTime, "Zsec", "zettasecond", "zettaseconds", "", ""},
	"Ysec":  {0, 0, _Y, UnitOfTime, "Ysec", "yottasecond", "yottaseconds", "", ""},
}

var timeAliases = map[string]Alias{
	"year": {"Gregorian year", ""},
	"sec":  {"second", ""},
}
