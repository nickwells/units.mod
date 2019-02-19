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
//
// For units of data the names follow the metric and IEC naming
// conventions so MB refers to a megabyte or 1000^2 bytes and MiB refers
// to a mebibyte or 1024^2 bytes
var UnitOfData = Family{
	BaseUnitName: "byte",
	Description:  "unit of data",
}

// ByteUnit is a suitable default value for a UnitOfData
var ByteUnit = Unit{1, UnitOfData, "B", UnitOfData.BaseUnitName}

var DataNames = map[string]Unit{
	"bit":    {0.125, UnitOfData, "bit", "bit"},
	"nibble": {0.5, UnitOfData, "nibble", "nibble"},
	"byte":   ByteUnit,
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
