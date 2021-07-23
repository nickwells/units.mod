package units

// bunData is the base unit name for data
const bunData = "byte"

// dataFamily represents the base unit of data
//
// For units of data the names follow the metric and IEC naming
// conventions so MB refers to a megabyte or 1000^2 bytes and MiB refers
// to a mebibyte or 1024^2 bytes
var dataFamily = &Family{
	baseUnitName: bunData,
	description:  "unit of data",
	name:         Data,
}

// DataNames maps names to units of data
var dataNames = map[string]Unit{
	bunData: {
		0, 0, 1,
		dataFamily,
		"B", bunData, "bytes",
		"eight bits",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},

	"bit": {
		0, 0, 0.125,
		dataFamily,
		"bit", "bit", "bits",
		"a bit is a single binary digit taking a value of either 0 or 1." +
			" It is one eighth of a byte and is indivisible.",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"nibble": {
		0, 0, 0.5,
		dataFamily,
		"nibble", "nibble", "nibbles",
		"a nibble is half a byte (four bits).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	// powers of 1000
	"KB": {
		0, 0, k,
		dataFamily,
		"KB", "kilobyte", "kilobytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"MB": {
		0, 0, _M,
		dataFamily,
		"MB", "megabyte", "megabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"GB": {
		0, 0, _G,
		dataFamily,
		"GB", "gigabyte", "gigabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"TB": {
		0, 0, _T,
		dataFamily,
		"TB", "terabyte", "terabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"PB": {
		0, 0, _P,
		dataFamily,
		"PB", "petabyte", "petabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"EB": {
		0, 0, _E,
		dataFamily,
		"EB", "exabyte", "exabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"ZB": {
		0, 0, _Z,
		dataFamily,
		"ZB", "zettabyte", "zettabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"YB": {
		0, 0, _Y,
		dataFamily,
		"YB", "yottabyte", "yottabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},

	// powers of 2 (1024 = 2^10)
	"KiB": {
		0, 0, ki,
		dataFamily,
		"KiB", "kibibyte", "kibibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"MiB": {
		0, 0, mi,
		dataFamily,
		"MiB", "mebibyte", "mebibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"GiB": {
		0, 0, gi,
		dataFamily,
		"GiB", "gibibyte", "gibibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"TiB": {
		0, 0, ti,
		dataFamily,
		"TiB", "tebibyte", "tebibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"PiB": {
		0, 0, pi,
		dataFamily,
		"PiB", "pebibyte", "pebibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"EiB": {
		0, 0, ei,
		dataFamily,
		"EiB", "exbibyte", "exbibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"ZiB": {
		0, 0, zi,
		dataFamily,
		"ZiB", "zebibyte", "zebibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
	"YiB": {
		0, 0, yi,
		dataFamily,
		"YiB", "yobibyte", "yobibytes",
		"a traditional computing measure (in powers of 2^10)." +
			" The name is as given by ISO/IEC 80000 and" +
			" supersedes the traditional name which is" +
			" now reserved for the power of 1000",
		[]Tag{TagComputing},
		map[string]string{},
		"", "",
	},
}
