package units

const (
	yard2ToMetre2 = yardToMetre * yardToMetre
	acreToMetre2  = yard2ToMetre2 * 4840
)

// UnitOfArea represents the base unit of area
var UnitOfArea = Family{
	BaseUnitName: "square metre",
	Description:  "unit of area",
}

// SquareMetreUnit is a suitable default value for a UnitOfArea
var SquareMetreUnit = Unit{
	0, 0, 1,
	UnitOfArea,
	"m\u00B2", UnitOfArea.BaseUnitName, "square metres",
	"a metric measure of area.",
	"",
}

// AreaNames maps names to units of area
var AreaNames = map[string]Unit{
	// metric
	"square metre": SquareMetreUnit,
	"are": {
		0, 0, 1e2,
		UnitOfArea,
		"are", "are", "ares",
		"a non-SI metric measure of area.",
		"",
	},
	"decare": {
		0, 0, 1e3,
		UnitOfArea,
		"decare", "decare", "decares",
		"a non-SI metric measure of area.",
		"",
	},
	"hectare": {
		0, 0, 1e4,
		UnitOfArea,
		"ha", "hectare", "hectares",
		"a non-SI metric measure of area.",
		"",
	},
	"square kilometre": {
		0, 0, 1e6,
		UnitOfArea,
		"km\u00B2", "square kilometre", "square kilometres",
		"a metric measure of area.",
		"",
	},

	// Imperial / US
	"square yard": {
		0, 0, yard2ToMetre2,
		UnitOfArea,
		"yd\u00B2", "square yard", "square yards",
		"an imperial measure of area.",
		"",
	},
	"square mile": {
		0, 0, yard2ToMetre2 * 1760 * 1760,
		UnitOfArea,
		"mi\u00B2", "square mile", "square miles",
		"an imperial measure of area.",
		"",
	},
	"perch": {
		0, 0, yard2ToMetre2 * 30.25,
		UnitOfArea,
		"square perch", "square perch", "square perches",
		"an imperial measure of area. It is equal to one square rod." +
			" Note that a perch can also be a" +
			" measure of length or volume.",
		"",
	},
	"rood": {
		0, 0, acreToMetre2 * 0.25,
		UnitOfArea,
		"ro", "rood", "roods",
		"an imperial measure of area. It may also be called a rod" +
			" though a rod is typically a measure of length. Note" +
			" however that a rood is not a square rod but is 40" +
			" square rods (one furlong by one rod).",
		"",
	},
	"acre": {
		0, 0, acreToMetre2,
		UnitOfArea,
		"acre", "acre", "acres",
		"an imperial measure of area.",
		"",
	},

	"oxgang": {
		0, 0, acreToMetre2 * 15,
		UnitOfArea,
		"oxgang", "oxgang", "oxgangs",
		"an obsolete imperial measure of area. Derived from the" +
			" area of land that could be ploughed by a single ox" +
			" in one season." +
			" Sometimes known as 'bovate' or, in Scotland, 'damh-imir'." +
			" Sometimes given as 20 acres." +
			" See also 'virgate' and 'carucate'.",
		"",
	},
	"virgate": {
		0, 0, acreToMetre2 * 30,
		UnitOfArea,
		"virgate", "virgate", "virgates",
		"an obsolete imperial measure of area. Derived from the" +
			" area of land that could be ploughed by a team of" +
			" 2 oxen in one season." +
			" See also 'oxgang' and 'carucate'.",
		"",
	},
	"carucate": {
		0, 0, acreToMetre2 * 120,
		UnitOfArea,
		"carucate", "carucate", "carucates",
		"an obsolete imperial measure of area. Derived from the" +
			" area of land that could be ploughed by a team of" +
			" 8 oxen in one season." +
			" See also 'oxgang' and 'virgate'.",
		"",
	},

	// Colloquial
	"Wales": {
		0, 0, 20779 * 1e6,
		UnitOfArea,
		"times the size of Wales",
		"times the size of Wales", "times the size of Wales",
		"a colloquial measure of area, used primarily to" +
			" give an impression of size.",
		"",
	},
}

var areaAliases = map[string]Alias{
	"square metres": {"square metre", "plural"},
	"square-metre":  {"square metre", "hyphenated"},
	"square-metres": {"square metre", "hyphenated, plural"},
	"square-meter":  {"square metre", "US spelling, hyphenated"},
	"square meter":  {"square metre", "US spelling"},
	"square-meters": {"square metre", "US spelling, hyphenated, plural"},
	"square meters": {"square metre", "US spelling, plural"},

	"square kilometres": {"square kilometre", "plural"},
	"square-kilometre":  {"square kilometre", "hyphenated"},
	"square-kilometer":  {"square kilometre", "US spelling, hyphenated"},
	"square kilometer":  {"square kilometre", "US spelling"},
	"square-kilometres": {"square kilometre", "hyphenated, plural"},
	"square-kilometers": {"square kilometre", "US spelling, hyphenated, plural"},
	"square kilometers": {"square kilometre", "US spelling, plural"},

	"square miles": {"square mile", "plural"},
	"square-mile":  {"square mile", "hyphenated"},
	"square-miles": {"square mile", "hyphenated, plural"},

	"square yards": {"square yard", "plural"},
	"square-yard":  {"square yard", "hyphenated"},
	"square-yards": {"square yard", "hyphenated, plural"},

	"acres":    {"acre", "plural"},
	"hectares": {"hectare", "plural"},
}
