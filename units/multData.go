package units

const (
	_ = 1 << (10 * iota)
	ki
	mi
	gi
	ti
	pi
	ei
	zi
	yi
)

// UnitOfData represents the base unit of data
var UnitOfData = Unit{
	BaseUnitName: "byte",
	Description:  "unit of data",
}

// ByteMult is a suitable default value for a Mult of UnitOfData
var ByteMult = Mult{1, UnitOfData, "B", UnitOfData.BaseUnitName}

var DataNames = map[string]Mult{
	"bit":    {0.125, UnitOfData, "bit", "bit"},
	"nibble": {0.5, UnitOfData, "nibble", "nibble"},
	"byte":   ByteMult,
	"KB":     {k, UnitOfData, "KB", "kilobyte"},
	"MB":     {_M, UnitOfData, "MB", "megabyte"},
	"GB":     {_G, UnitOfData, "GB", "gigabyte"},
	"TB":     {_T, UnitOfData, "TB", "terabyte"},
	"PB":     {_P, UnitOfData, "PB", "petabyte"},
	"EB":     {_E, UnitOfData, "EB", "exabyte"},
	"ZB":     {_Z, UnitOfData, "ZB", "zettabyte"},
	"YB":     {_Y, UnitOfData, "YB", "yottabyte"},

	"KiB": {ki, UnitOfData, "KiB", "kibibyte"},
	"MiB": {mi, UnitOfData, "MiB", "mebibyte"},
	"GiB": {gi, UnitOfData, "GiB", "gibibyte"},
	"TiB": {ti, UnitOfData, "TiB", "tebibyte"},
	"PiB": {pi, UnitOfData, "PiB", "pebibyte"},
	"EiB": {ei, UnitOfData, "EiB", "exbibyte"},
	"ZiB": {zi, UnitOfData, "ZiB", "zebibyte"},
	"YiB": {yi, UnitOfData, "YiB", "yobibyte"},
}
