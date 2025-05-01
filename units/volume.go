package units

// bunVolume is the base unit name for volume
const bunVolume = "cubic metre"

// volumeFamily represents the collection of units of volume
var volumeFamily = &Family{
	baseUnitName: bunVolume,
	description:  "unit of volume",
	name:         Volume,
}

// VolumeNames maps names to units of volume
var volumeNames = map[string]Unit{
	// metric
	bunVolume: {
		0, 0, 1,
		volumeFamily,
		"m\u00B3", bunVolume, "cubic metres",
		"a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"cubic-metre":  "hyphenated",
			"cubic-metres": "hyphenated, plural",
			"cubic metres": "plural",
			"cubic-meter":  "US spelling, hyphenated",         //nolint:misspell
			"cubic-meters": "US spelling, hyphenated, plural", //nolint:misspell
			"cubic meter":  "US spelling",                     //nolint:misspell
			"cubic meters": "US spelling, plural",             //nolint:misspell
		},
		"", "",
	},

	"yl": {
		0, 0, y * litreToCubicMetre,
		volumeFamily,
		"yl", "yoctolitre", "yoctolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"zl": {
		0, 0, z * litreToCubicMetre,
		volumeFamily,
		"zl", "zeptolitre", "zeptolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"al": {
		0, 0, a * litreToCubicMetre,
		volumeFamily,
		"al", "attolitre", "attolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"fl": {
		0, 0, f * litreToCubicMetre,
		volumeFamily,
		"fl", "femtolitre", "femtolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"pl": {
		0, 0, p * litreToCubicMetre,
		volumeFamily,
		"pl", "picolitre", "picolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"nl": {
		0, 0, n * litreToCubicMetre,
		volumeFamily,
		"nl", "nanolitre", "nanolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"ul": {
		0, 0, u * litreToCubicMetre,
		volumeFamily,
		"ul", "microlitre", "microlitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"ml": {
		0, 0, m * litreToCubicMetre,
		volumeFamily,
		"ml", "millilitre", "millilitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"cc":                "abbreviation",
			"cubic-centimetre":  "hyphenated",
			"cubic-centimetres": "hyphenated, plural",
			"cubic centimetre":  "full name",
			"cubic centimetres": "full name, plural",
			"cubic-centimeter":  "full name, US spelling, hyphenated",         //nolint:misspell
			"cubic-centimeters": "full name, US spelling, hyphenated, plural", //nolint:misspell
			"cubic centimeter":  "full name, US spelling",                     //nolint:misspell
			"cubic centimeters": "full name, US spelling, plural",             //nolint:misspell

			"millilitre":  "full name",
			"millilitres": "full name, plural",
			"milliliter":  "full name, US spelling",         //nolint:misspell
			"milliliters": "full name, US spelling, plural", //nolint:misspell
		},
		"", "",
	},
	"cl": {
		0, 0, c * litreToCubicMetre,
		volumeFamily,
		"cl", "centilitre", "centilitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"dl": {
		0, 0, d * litreToCubicMetre,
		volumeFamily,
		"dl", "decilitre", "decilitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"litre": {
		0, 0, litreToCubicMetre,
		volumeFamily,
		"l", "litre", "litres",
		"a metric measure of volume." +
			" It is not an SI unit but may be used alongside them",
		[]Tag{TagMetric},
		map[string]string{
			"l":      "abbreviation",
			"litres": "plural",
			"liter":  "US spelling",         //nolint:misspell
			"liters": "US spelling, plural", //nolint:misspell
		},
		"", "",
	},
	"dal": {
		0, 0, da * litreToCubicMetre,
		volumeFamily,
		"dal", "decalitre", "decalitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"hl": {
		0, 0, h * litreToCubicMetre,
		volumeFamily,
		"hl", "hectolitre", "hectolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"kl": {
		0, 0, k * litreToCubicMetre,
		volumeFamily,
		"kl", "kilolitre", "kilolitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Ml": {
		0, 0, _M * litreToCubicMetre,
		volumeFamily,
		"Ml", "megalitre", "megalitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Gl": {
		0, 0, _G * litreToCubicMetre,
		volumeFamily,
		"Gl", "gigalitre", "gigalitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Tl": {
		0, 0, _T * litreToCubicMetre,
		volumeFamily,
		"Tl", "teralitre", "teralitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{
			"cubic kilometre":  "",
			"cubic kilometres": "plural",
			"cubic-kilometre":  "hyphenated",
			"cubic-kilometres": "hyphenated, plural",
			"cubic-kilometer":  "US spelling, hyphenated",         //nolint:misspell
			"cubic-kilometers": "US spelling, hyphenated, plural", //nolint:misspell
			"cubic kilometer":  "US spelling",                     //nolint:misspell
			"cubic kilometers": "US spelling, plural",             //nolint:misspell
		},
		"", "",
	},
	"Pl": {
		0, 0, _P * litreToCubicMetre,
		volumeFamily,
		"Pl", "petalitre", "petalitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"El": {
		0, 0, _E * litreToCubicMetre,
		volumeFamily,
		"El", "exalitre", "exalitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Zl": {
		0, 0, _Z * litreToCubicMetre,
		volumeFamily,
		"Zl", "zettalitre", "zettalitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},
	"Yl": {
		0, 0, _Y * litreToCubicMetre,
		volumeFamily,
		"Yl", "yottalitre", "yottalitres", "a metric measure of volume.",
		[]Tag{TagSI, TagMetric},
		map[string]string{},
		"", "",
	},

	// bottle
	"bottle-wine": {
		0, 0, bottleToCubicMetre,
		volumeFamily,
		"bottle", "bottle (wine)", "bottles (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade. Note that there is a now obsolete bottle size of" +
			" 70cl with corresponding multiples.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"magnum": {
		0, 0, bottleToCubicMetre * 2,
		volumeFamily,
		"Magnum", "Magnum (wine)", "Magnums (wine)",
		"two bottles.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"marie-jeanne": {
		0, 0, bottleToCubicMetre * 3,
		volumeFamily,
		"Marie Jeanne", "Marie Jeanne (wine)", "Marie Jeannes (wine)",
		"three bottles.",
		[]Tag{TagDrinks},
		map[string]string{
			"tappit-hen": "specifically for port, the name" +
				" derives from the shape of the bottle which is said to" +
				" resemble a hen.",
			"tregnum": "specifically for port.",
		},
		"", "",
	},
	"jeroboam": {
		0, 0, bottleToCubicMetre * 4,
		volumeFamily,
		"Jeroboam", "Jeroboam (wine)", "Jeroboams (wine)",
		"four bottles.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"rehoboam": {
		0, 0, bottleToCubicMetre * 6,
		volumeFamily,
		"Rehoboam", "Rehoboam (wine)", "Rehoboams (wine)",
		"6 bottles.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"methuselah": {
		0, 0, bottleToCubicMetre * 8,
		volumeFamily,
		"Methuselah", "Methuselah (wine)", "Methuselahs (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"salmanazar": {
		0, 0, bottleToCubicMetre * 12,
		volumeFamily,
		"Salmanazar", "Salmanazar (wine)", "Salmanazars (wine)",
		"12 bottles.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"balthazar": {
		0, 0, bottleToCubicMetre * 16,
		volumeFamily,
		"Balthazar", "Balthazar (wine)", "Balthazars (wine)",
		"16 bottles.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"nebuchadnezzar": {
		0, 0, bottleToCubicMetre * 20,
		volumeFamily,
		"Nebuchadnezzar", "Nebuchadnezzar (wine)", "Nebuchadnezzars (wine)",
		"20 bottles.",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},

	// Imperial / US
	"cubic inch": {
		0, 0, cubicInchToCubicMetre,
		volumeFamily,
		"in\u00B3", "cubic inch", "cubic inches",
		"",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"cubic-inch":   "hyphenated",
			"cubic-inches": "hyphenated, plural",
			"cubic inches": "plural",
		},
		"", "",
	},
	"cubic foot": {
		0, 0, cubicFootToCubicMetre,
		volumeFamily,
		"ft\u00B3", "cubic foot", "cubic feet",
		"",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"cubic-foot": "hyphenated",
			"cubic-feet": "hyphenated, plural",
			"cubic feet": "plural",
		},
		"", "",
	},
	"cubic yard": {
		0, 0, cubicYardToCubicMetre,
		volumeFamily,
		"yd\u00B3", "cubic yard", "cubic yards",
		"",
		[]Tag{TagImperial, TagUScustomary},
		map[string]string{
			"cubic-yard":  "hyphenated",
			"cubic-yards": "hyphenated, plural",
			"cubic yards": "plural",
		},
		"", "",
	},

	"perch (masonry)": {
		0, 0, 16.5 * 1.5 * 1 * cubicFootToCubicMetre,
		volumeFamily,
		"masonry perch", "masonry perch", "masonry perches",
		"a traditional measure of volume for stone and other" +
			" masonry. Note that this refers to the measure" +
			" for dressed stone, a masonry perch of brickwork" +
			" or rubble is a third smaller being one perch" +
			" long by one foot high by one foot wide." +
			" Note that a perch can also be a" +
			" measure of length or area.",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},

	"minim": {
		0, 0, fluidOzToCubicMetre / 480,
		volumeFamily,
		"min", "minim", "minims",
		"an apothecaries' measure.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},
	"fluid-scruple": {
		0, 0, fluidOzToCubicMetre / 24,
		volumeFamily,
		"fl s", "fluid scruple", "fluid scruples",
		"an apothecaries' measure.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},
	"fluid-drachm": {
		0, 0, fluidOzToCubicMetre / 8,
		volumeFamily,
		"fl dr", "fluid drachm", "fluid drachms",
		"an apothecaries' measure.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},

	"fluid-ounce": {
		0, 0, fluidOzToCubicMetre,
		volumeFamily,
		"fl oz", "fluid ounce", "fluid ounces",
		"this is an imperial fluid ounce." +
			" It is 1/160th of an imperial gallon" +
			" which was defined as the volume of" +
			" 10 pounds of water at 62°F" +
			" meaning that a fluid ounce weighed an ounce.",
		[]Tag{TagImperial, TagApothecary},
		map[string]string{},
		"", "",
	},
	"gill": {
		0, 0, fluidOzToCubicMetre * 5,
		volumeFamily,
		"gi", "gill", "gills",
		"5 imperial fluid ounces.",
		[]Tag{TagImperial},
		map[string]string{},
		"", "",
	},
	"pint": {
		0, 0, fluidOzToCubicMetre * 20,
		volumeFamily,
		"pt", "pint", "pints",
		"20 imperial fluid ounces (4 gills)." +
			" See also the US pint.",
		[]Tag{TagImperial},
		map[string]string{
			"pints": "plural",
		},
		"", "",
	},
	"quart": {
		0, 0, fluidOzToCubicMetre * 40,
		volumeFamily,
		"qt", "quart", "quarts",
		"40 imperial fluid ounces (2 pints).",
		[]Tag{TagImperial},
		map[string]string{},
		"", "",
	},
	"gallon": {
		0, 0, gallonToCubicMetre,
		volumeFamily,
		"gal", "gallon", "gallons",
		"160 imperial fluid ounces (4 quarts).",
		[]Tag{TagImperial},
		map[string]string{},
		"", "",
	},

	"pin": {
		0, 0, gallonToCubicMetre * 4.5,
		volumeFamily,
		"pin", "pin", "pins",
		"4.5 imperial gallons (36 imperial pints).",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"firkin": {
		0, 0, gallonToCubicMetre * 9,
		volumeFamily,
		"firkin", "firkin", "firkins",
		"9 imperial gallons (72 imperial pints).",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"kilderkin": {
		0, 0, gallonToCubicMetre * 18,
		volumeFamily,
		"kilderkin", "kilderkin", "kilderkins",
		"18 imperial gallons (144 imperial pints).",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"barrel": {
		0, 0, gallonToCubicMetre * 36,
		volumeFamily,
		"barrel", "barrel", "barrels",
		"36 imperial gallons (288 imperial pints).",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},
	"hogshead": {
		0, 0, gallonToCubicMetre * 54,
		volumeFamily,
		"hogshead", "hogshead", "hogsheads",
		"54 imperial gallons (432 imperial pints).",
		[]Tag{TagDrinks},
		map[string]string{},
		"", "",
	},

	"peck": {
		0, 0, gallonToCubicMetre * 2,
		volumeFamily,
		"peck", "peck", "pecks",
		"an imperial measure of volume, typically for dry goods.",
		[]Tag{TagImperial},
		map[string]string{},
		"", "",
	},
	"bushel": {
		0, 0, gallonToCubicMetre * 8,
		volumeFamily,
		"bushel", "bushel", "bushels",
		"an imperial measure of volume, typically for dry goods.",
		[]Tag{TagImperial},
		map[string]string{},
		"", "",
	},

	"teaspoon": {
		0, 0, 5 * m * litreToCubicMetre,
		volumeFamily,
		"tsp", "teaspoon", "teaspoons",
		"used for cooking, nutritional labelling or medical measures.",
		[]Tag{TagMetric},
		map[string]string{
			"tsp":              "abbreviation",
			"teaspoons":        "plural",
			"metric teaspoon":  "alternative",
			"metric teaspoons": "alternative, plural",
		},
		"", "",
	},

	"tablespoon": {
		0, 0, 15 * m * litreToCubicMetre,
		volumeFamily,
		"Tbsp", "tablespoon", "tablespoons",
		"used for cooking or nutritional labelling.",
		[]Tag{TagMetric},
		map[string]string{
			"Tbsp":               "abbreviation",
			"tbsp":               "abbreviation, all lower-case",
			"tablespoons":        "plural",
			"metric tablespoon":  "alternative",
			"metric tablespoons": "alternative, plural",
		},
		"", "",
	},

	"Australian tablespoon": {
		0, 0, 20 * m * litreToCubicMetre,
		volumeFamily,
		"Tbsp", "tablespoon", "tablespoons",
		"A slightly larger Australian version.",
		[]Tag{TagMetric, TagColloquial},
		map[string]string{},
		"", "",
	},

	"US teaspoon": {
		0, 0, usFluidOzToCubicMetre * 0.5 / 3.0,
		volumeFamily,
		"tsp", "teaspoon", "teaspoons",
		"a US measure of volume," +
			" not used for nutritional labelling" +
			" or for measuring medicines" +
			" for which the metric variant is used.",
		[]Tag{TagUScustomary},
		map[string]string{
			"US-teaspoon":  "hyphenated",
			"US-teaspoons": "hyphenated, plural",
			"US teaspoon":  "plural",
		},
		"", "",
	},

	"US tablespoon": {
		0, 0, usFluidOzToCubicMetre * 0.5,
		volumeFamily,
		"Tbsp", "tablespoon", "tablespoons",
		"a US measure of volume," +
			" not used for nutritional labelling" +
			" for which the metric variant is used.",
		[]Tag{TagUScustomary},
		map[string]string{
			"US-tablespoon":  "hyphenated",
			"US-tablespoons": "hyphenated, plural",
			"US tablespoon":  "plural",
		},
		"", "",
	},

	"US-fluid-ounce": {
		0, 0, usFluidOzToCubicMetre,
		volumeFamily,
		"US fl oz", "US fluid ounce", "US fluid ounces",
		"The US fluid ounce is defined as 1/128 of a US gallon" +
			" which is based on the wine gallon of 231 cubic inches" +
			" which was used in the UK prior to the adoption of" +
			" the imperial system after 1824.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-shot": {
		0, 0, usFluidOzToCubicMetre * 1.5,
		volumeFamily,
		"shot", "US shot", "US shots",
		"1.5 US fluid ounces.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-gill": {
		0, 0, usFluidOzToCubicMetre * 4,
		volumeFamily,
		"US gi", "US gill", "US gills",
		"4 US fluid ounces.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-cup": {
		0, 0, usFluidOzToCubicMetre * 8,
		volumeFamily,
		"US cp", "US cup", "US cups",
		"8 US fluid ounces.",
		[]Tag{TagUScustomary},
		map[string]string{
			"cup":  "",
			"cups": "plural",
		},
		"", "",
	},
	"US-pint": {
		0, 0, usFluidOzToCubicMetre * 16,
		volumeFamily,
		"US pt", "US pint", "US pints",
		"16 US fluid ounces.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-quart": {
		0, 0, usFluidOzToCubicMetre * 32,
		volumeFamily,
		"US qt", "US quart", "US quarts",
		"32 US fluid ounces.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-gallon": {
		0, 0, usFluidOzToCubicMetre * 128,
		volumeFamily,
		"US gal", "US gallon", "US gallons",
		"128 US fluid ounces.",
		[]Tag{TagUScustomary},
		map[string]string{
			"US-wet-gallon": "",
			"US gallon":     "no hyphen",
		},
		"", "",
	},
	"barrel-oil": {
		0, 0, usFluidOzToCubicMetre * 128 * 42,
		volumeFamily,
		"bbl(oil)", "oil barrel", "oil barrels",
		"42 US gallons.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-dry-gallon": {
		0, 0, usDryGallonToCubicMetre,
		volumeFamily,
		"US dry gal", "US dry gallon", "US dry gallons",
		"a US Customary measure of volume, typically for dry goods.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-bushel": {
		0, 0, usDryGallonToCubicMetre * 8,
		volumeFamily,
		"US bushel", "US bushel", "US bushels",
		"a US Customary measure of volume, typically for dry goods.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"US-dry-pint": {
		0, 0, usDryGallonToCubicMetre * 0.125,
		volumeFamily,
		"US dry pint", "US dry pint", "US dry pints",
		"a US Customary measure of volume, typically for dry goods.",
		[]Tag{TagUScustomary},
		map[string]string{},
		"", "",
	},
	"wine-gallon": {
		0, 0, cubicInchToCubicMetre * 231,
		volumeFamily,
		"wine gallon", "wine gallon", "wine gallons",
		"a measure of volume in the UK, customarily from the 16th century" +
			" and by law since 1707 during the reign of Queen Anne." +
			" It is of interest principally because" +
			" it is the basis of US fluid measures",
		[]Tag{TagHist},
		map[string]string{},
		"", "",
	},

	// Colloquial
	"20ft shipping container": {
		0, 0, 6.058 * 2.44 * 2.59,
		volumeFamily,
		"vol. of a TEU",
		"volume of a 20ft shipping container",
		"volumes of a 20ft shipping container",
		"a colloquial term referencing the volume of a standard 20 foot" +
			" shipping (or intermodal) container" +
			" presumably to give an impression of scale." +
			" The TEU abbreviation stands for Twenty foot Equivalent Unit." +
			"\n\n" +
			"The standard intermodal container is 19 feet 10.5 inches long" +
			" by 8 feet wide and, most commonly, 8 feet 6 inches high." +
			"\n\n" +
			"The volume used in this definition is based on this most common" +
			" height though the height ranges between 4 feet 3 inches and" +
			" 9 feet 6 inches." +
			" Note that the internal volume is smaller at" +
			" only 1172 cubic feet or 33.2 cubic metres." +
			"\n\n" +
			"The reason for the 1.5 inch difference in length is because" +
			" when they are joined the connecting twistlock is three inches" +
			" wide allowing a forty-foot unit to fit precisely on top.",
		[]Tag{TagColloquial},
		map[string]string{
			"TEU": "abbreviation",
		},
		"", "",
	},
}
