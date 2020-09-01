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
var ByteUnit = Unit{0, 0, 1, UnitOfData, "B", UnitOfData.BaseUnitName, "bytes", ""}

// DataNames maps names to units of data
var DataNames = map[string]Unit{
	"bit":    {0, 0, 0.125, UnitOfData, "bit", "bit", "bits", ""},
	"nibble": {0, 0, 0.5, UnitOfData, "nibble", "nibble", "nibbles", "a nibble is half a byte"},
	"byte":   ByteUnit,
	// powers of 1000
	"KB": {0, 0, k, UnitOfData, "KB", "kilobyte", "kilobytes", "metric value (in powers of 10^3)"},
	"MB": {0, 0, _M, UnitOfData, "MB", "megabyte", "megabytes", "metric value (in powers of 10^3)"},
	"GB": {0, 0, _G, UnitOfData, "GB", "gigabyte", "gigabytes", "metric value (in powers of 10^3)"},
	"TB": {0, 0, _T, UnitOfData, "TB", "terabyte", "terabytes", "metric value (in powers of 10^3)"},
	"PB": {0, 0, _P, UnitOfData, "PB", "petabyte", "petabytes", "metric value (in powers of 10^3)"},
	"EB": {0, 0, _E, UnitOfData, "EB", "exabyte", "exabytes", "metric value (in powers of 10^3)"},
	"ZB": {0, 0, _Z, UnitOfData, "ZB", "zettabyte", "zettabytes", "metric value (in powers of 10^3)"},
	"YB": {0, 0, _Y, UnitOfData, "YB", "yottabyte", "yottabytes", "metric value (in powers of 10^3)"},

	// powers of 2 (1024 = 2^10)
	"KiB": {0, 0, ki, UnitOfData, "KiB", "kibibyte", "kibibytes", "traditional computing measure (in powers of 2^10)"},
	"MiB": {0, 0, mi, UnitOfData, "MiB", "mebibyte", "mebibytes", "traditional computing measure (in powers of 2^10)"},
	"GiB": {0, 0, gi, UnitOfData, "GiB", "gibibyte", "gibibytes", "traditional computing measure (in powers of 2^10)"},
	"TiB": {0, 0, ti, UnitOfData, "TiB", "tebibyte", "tebibytes", "traditional computing measure (in powers of 2^10)"},
	"PiB": {0, 0, pi, UnitOfData, "PiB", "pebibyte", "pebibytes", "traditional computing measure (in powers of 2^10)"},
	"EiB": {0, 0, ei, UnitOfData, "EiB", "exbibyte", "exbibytes", "traditional computing measure (in powers of 2^10)"},
	"ZiB": {0, 0, zi, UnitOfData, "ZiB", "zebibyte", "zebibytes", "traditional computing measure (in powers of 2^10)"},
	"YiB": {0, 0, yi, UnitOfData, "YiB", "yobibyte", "yobibytes", "traditional computing measure (in powers of 2^10)"},
}
