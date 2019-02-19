package units

// UnitOfTime represents the base unit of time
var UnitOfTime = Unit{
	BaseUnitName: "second",
	Description:  "unit of time",
}

// SecondMult is a suitable default value for a Mult of UnitOfTime
var SecondMult = Mult{1, UnitOfTime, "sec", UnitOfTime.BaseUnitName}

var TimeNames = map[string]Mult{
	"ysec":   {y, UnitOfTime, "ysec", "yoctosecond"},
	"zsec":   {z, UnitOfTime, "zsec", "zeptosecond"},
	"asec":   {a, UnitOfTime, "asec", "attosecond"},
	"fsec":   {f, UnitOfTime, "fsec", "femtosecond"},
	"psec":   {p, UnitOfTime, "psec", "picosecond"},
	"nsec":   {n, UnitOfTime, "nsec", "nanosecond"},
	"usec":   {u, UnitOfTime, "μsec", "microsecond"},
	"msec":   {m, UnitOfTime, "msec", "millisecond"},
	"csec":   {c, UnitOfTime, "csec", "centisecond"},
	"dsec":   {d, UnitOfTime, "dsec", "decisecond"},
	"sec":    SecondMult,
	"second": SecondMult,
	"minute": {60, UnitOfTime, "min", "minute"},
	"hour":   {60 * 60, UnitOfTime, "hr", "hour"},
	"day":    {60 * 60 * 24, UnitOfTime, "day", "day"},
	"week":   {60 * 60 * 24 * 7, UnitOfTime, "week", "week"},
	"dasec":  {da, UnitOfTime, "dasec", "decasecond"},
	"hsec":   {h, UnitOfTime, "hsec", "hectosecond"},
	"ksec":   {k, UnitOfTime, "ksec", "kilosecond"},
	"Msec":   {_M, UnitOfTime, "Msec", "megasecond"},
	"Gsec":   {_G, UnitOfTime, "Gsec", "gigasecond"},
	"Tsec":   {_T, UnitOfTime, "Tsec", "terasecond"},
	"Psec":   {_P, UnitOfTime, "Psec", "petasecond"},
	"Esec":   {_E, UnitOfTime, "Esec", "exasecond"},
	"Zsec":   {_Z, UnitOfTime, "Zsec", "zettasecond"},
	"Ysec":   {_Y, UnitOfTime, "Ysec", "yottasecond"},
}
