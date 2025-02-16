//nolint:mnd
package units

const bunArea = "square metre"

// areaFamily represents the collection of units of area
var areaFamily = &Family{
	baseUnitName: bunArea,
	description:  "unit of area",
	name:         Area,
}

// AreaNames maps names to units of area
var areaNames = map[string]Unit{
	// metric
	bunArea: {
		0, 0, 1,
		areaFamily,
		"m\u00B2", bunArea, "square metres",
		"a metric measure of area.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"square metres": "plural",
			"square-metre":  "hyphenated",
			"square-metres": "hyphenated, plural",
			"square-meter":  "US spelling, hyphenated",
			"square meter":  "US spelling",
			"square-meters": "US spelling, hyphenated, plural",
			"square meters": "US spelling, plural",
		},
		"", "",
	},

	"are": {
		0, 0, 1e2,
		areaFamily,
		"are", "are", "ares",
		"a non-SI metric measure of area.",
		[]Tag{TagMetric},
		map[string]string{},
		"", "",
	},
	"decare": {
		0, 0, 1e3,
		areaFamily,
		"decare", "decare", "decares",
		"a non-SI metric measure of area.",
		[]Tag{TagMetric},
		map[string]string{
			"new dunam": `An Ottoman Turkish measure of area,` +
				` formerly "forty standard paces in length and breadth".`,
			"stremma": "A Greek unit of land area",
		},
		"", "",
	},
	"hectare": {
		0, 0, 1e4,
		areaFamily,
		"ha", "hectare", "hectares",
		"a non-SI metric measure of area." +
			" Primariy used in the measurement of land",
		[]Tag{TagMetric},
		map[string]string{
			"hectares": "plural",
		},
		"", "",
	},
	"square kilometre": {
		0, 0, 1e6,
		areaFamily,
		"km\u00B2", "square kilometre", "square kilometres",
		"a metric measure of area.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"square kilometres": "plural",
			"square-kilometre":  "hyphenated",
			"square-kilometer":  "US spelling, hyphenated",
			"square kilometer":  "US spelling",
			"square-kilometres": "hyphenated, plural",
			"square-kilometers": "US spelling, hyphenated, plural",
			"square kilometers": "US spelling, plural",
		},
		"", "",
	},

	// Imperial / US
	"square yard": {
		0, 0, yard2ToMetre2,
		areaFamily,
		"yd\u00B2", "square yard", "square yards",
		"an imperial measure of area.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"square yards": "plural",
			"square-yard":  "hyphenated",
			"square-yards": "hyphenated, plural",
		},
		"", "",
	},
	"square mile": {
		0, 0, yard2ToMetre2 * 1760 * 1760,
		areaFamily,
		"mi\u00B2", "square mile", "square miles",
		"an imperial measure of area.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"square miles": "plural",
			"square-mile":  "hyphenated",
			"square-miles": "hyphenated, plural",
		},
		"", "",
	},
	"perch": {
		0, 0, yard2ToMetre2 * 30.25,
		areaFamily,
		"square perch", "square perch", "square perches",
		"an imperial measure of area. It is equal to one square rod." +
			" Note that a perch can also be a" +
			" measure of length or volume.",
		[]Tag{TagImperial, TagHist},
		map[string]string{},
		"", "",
	},
	"rood": {
		0, 0, acreToMetre2 * 0.25,
		areaFamily,
		"ro", "rood", "roods",
		"an imperial measure of area. It may also be called a rod" +
			" though a rod is typically a measure of length. Note" +
			" however that a rood is not a square rod but is 40" +
			" square rods (one furlong by one rod).",
		[]Tag{TagImperial, TagHist},
		map[string]string{},
		"", "",
	},
	"acre": {
		0, 0, acreToMetre2,
		areaFamily,
		"acre", "acre", "acres",
		"an imperial measure of area.",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"acres": "plural",
		},
		"", "",
	},

	"oxgang": {
		0, 0, acreToMetre2 * 15,
		areaFamily,
		"oxgang", "oxgang", "oxgangs",
		"an obsolete imperial measure of area. Derived from the" +
			" area of land that could be ploughed by a single ox" +
			" in one season." +
			" Sometimes known as 'bovate' or, in Scotland, 'damh-imir'." +
			" Sometimes given as 20 acres." +
			" See also 'virgate' and 'carucate'.",
		[]Tag{TagImperial, TagHist},
		map[string]string{},
		"", "",
	},
	"virgate": {
		0, 0, acreToMetre2 * 30,
		areaFamily,
		"virgate", "virgate", "virgates",
		"an obsolete imperial measure of area. Derived from the" +
			" area of land that could be ploughed by a team of" +
			" 2 oxen in one season." +
			" See also 'oxgang' and 'carucate'.",
		[]Tag{TagImperial, TagHist},
		map[string]string{},
		"", "",
	},
	"carucate": {
		0, 0, acreToMetre2 * 120,
		areaFamily,
		"carucate", "carucate", "carucates",
		"an obsolete imperial measure of area. Derived from the" +
			" area of land that could be ploughed by a team of" +
			" 8 oxen in one season." +
			" See also 'oxgang' and 'virgate'.",
		[]Tag{TagImperial, TagHist},
		map[string]string{},
		"", "",
	},

	// Colloquial
	"Wales": {
		0, 0, 20779 * 1e6,
		areaFamily,
		"Wales",
		"times the size of Wales", "times the size of Wales",
		"Wales is a mountainous country that is part of the United Kingdom" +
			" lying west of the greater part of England." +
			" In the Welsh language it is called Cymru.",
		[]Tag{TagColloquial},
		map[string]string{},
		"", "",
	},

	"football pitch": {
		0, 0, 7140,
		areaFamily,
		"football pitch", "football pitch", "football pitches",
		"the size of a football pitch can vary considerably but" +
			" the preferred size for many professional teams' stadiums" +
			" is 105m x 68m giving the figure used here.",
		[]Tag{TagColloquial},
		map[string]string{
			"soccer pitch": "alternative",
		},
		"", "",
	},

	"American football pitch": {
		0, 0, 6400 * yard2ToMetre2,
		areaFamily,
		"football pitch", "football pitch", "football pitches",
		"120 x 53 1/3 yards.",
		[]Tag{TagColloquial},
		map[string]string{
			"US football pitch": "alternative",
		},
		"", "",
	},
}
