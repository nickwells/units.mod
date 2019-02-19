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
var ByteUnit = Unit{0, 0, 1, UnitOfData, "B", UnitOfData.BaseUnitName, "bytes"}

// DataNames maps names to units of data
var DataNames = map[string]Unit{
	"bit":    {0, 0, 0.125, UnitOfData, "bit", "bit", "bits"},
	"nibble": {0, 0, 0.5, UnitOfData, "nibble", "nibble", "nibbles"},
	"byte":   ByteUnit,
	"KB":     {0, 0, k, UnitOfData, "KB", "kilobyte", "kilobytes"},
	"MB":     {0, 0, _M, UnitOfData, "MB", "megabyte", "megabytes"},
	"GB":     {0, 0, _G, UnitOfData, "GB", "gigabyte", "gigabytes"},
	"TB":     {0, 0, _T, UnitOfData, "TB", "terabyte", "terabytes"},
	"PB":     {0, 0, _P, UnitOfData, "PB", "petabyte", "petabytes"},
	"EB":     {0, 0, _E, UnitOfData, "EB", "exabyte", "exabytes"},
	"ZB":     {0, 0, _Z, UnitOfData, "ZB", "zettabyte", "zettabytes"},
	"YB":     {0, 0, _Y, UnitOfData, "YB", "yottabyte", "yottabytes"},

	"KiB": {0, 0, ki, UnitOfData, "KiB", "kibibyte", "kibibytes"},
	"MiB": {0, 0, mi, UnitOfData, "MiB", "mebibyte", "mebibytes"},
	"GiB": {0, 0, gi, UnitOfData, "GiB", "gibibyte", "gibibytes"},
	"TiB": {0, 0, ti, UnitOfData, "TiB", "tebibyte", "tebibytes"},
	"PiB": {0, 0, pi, UnitOfData, "PiB", "pebibyte", "pebibytes"},
	"EiB": {0, 0, ei, UnitOfData, "EiB", "exbibyte", "exbibytes"},
	"ZiB": {0, 0, zi, UnitOfData, "ZiB", "zebibyte", "zebibytes"},
	"YiB": {0, 0, yi, UnitOfData, "YiB", "yobibyte", "yobibytes"},
}
