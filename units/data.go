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
		map[string]string{
			"bytes": "plural",
		},
		"", "",
	},

	"bit": {
		0, 0, 0.125,
		dataFamily,
		"bit", "bit", "bits",
		"a bit is a single binary digit taking a value of either 0 or 1." +
			" It is one eighth of a byte and is indivisible.",
		[]Tag{TagComputing},
		map[string]string{
			"bits": "plural",
		},
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
		map[string]string{
			"kB":        "lowercase",
			"kilobyte":  "expanded",
			"kilobytes": "expanded, plural",
		},
		"", "",
	},
	"MB": {
		0, 0, _M,
		dataFamily,
		"MB", "megabyte", "megabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{
			"megabyte":  "expanded",
			"megabytes": "expanded, plural",
		},
		"", "",
	},
	"GB": {
		0, 0, _G,
		dataFamily,
		"GB", "gigabyte", "gigabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{
			"gigabyte":  "expanded",
			"gigabytes": "expanded, plural",
		},
		"", "",
	},
	"TB": {
		0, 0, _T,
		dataFamily,
		"TB", "terabyte", "terabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{
			"terabyte":  "expanded",
			"terabytes": "expanded, plural",
		},
		"", "",
	},
	"PB": {
		0, 0, _P,
		dataFamily,
		"PB", "petabyte", "petabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{
			"petabyte":  "expanded",
			"petabytes": "expanded, plural",
		},
		"", "",
	},
	"EB": {
		0, 0, _E,
		dataFamily,
		"EB", "exabyte", "exabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{
			"exabyte":  "expanded",
			"exabytes": "expanded, plural",
		},
		"", "",
	},
	"ZB": {
		0, 0, _Z,
		dataFamily,
		"ZB", "zettabyte", "zettabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{
			"zettabyte":  "expanded",
			"zettabytes": "expanded, plural",
		},
		"", "",
	},
	"YB": {
		0, 0, _Y,
		dataFamily,
		"YB", "yottabyte", "yottabytes",
		"a metric value (in powers of 10^3).",
		[]Tag{TagComputing},
		map[string]string{
			"yottabyte":  "expanded",
			"yottabytes": "expanded, plural",
		},
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
		map[string]string{
			"kibibyte":  "expanded",
			"kibibytes": "expanded, plural",
		},
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
		map[string]string{
			"mebibyte":  "expanded",
			"mebibytes": "expanded, plural",
		},
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
		map[string]string{
			"gibibyte":  "expanded",
			"gibibytes": "expanded, plural",
		},
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
		map[string]string{
			"tebibyte":  "expanded",
			"tebibytes": "expanded, plural",
		},
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
		map[string]string{
			"pebibyte":  "expanded",
			"pebibytes": "expanded, plural",
		},
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
		map[string]string{
			"exbibyte":  "expanded",
			"exbibytes": "expanded, plural",
		},
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
		map[string]string{
			"zebibyte":  "expanded",
			"zebibytes": "expanded, plural",
		},
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
		map[string]string{
			"yobibyte":  "expanded",
			"yobibytes": "expanded, plural",
		},
		"", "",
	},
}
