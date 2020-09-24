package units

const (
	cubicInchToCubicMetre   = (inchToMetre * inchToMetre * inchToMetre)
	cubicFootToCubicMetre   = (footToMetre * footToMetre * footToMetre)
	cubicYardToCubicMetre   = (yardToMetre * yardToMetre * yardToMetre)
	bottleToCubicMetre      = 0.75e-3
	gallonToCubicMetre      = 4.54609e-3
	fluidOzToCubicMetre     = gallonToCubicMetre / 160
	usFluidOzToCubicMetre   = 29.5735295625e-6
	usDryGallonToCubicMetre = 4.40488377086e-3
)

// UnitOfVolume represents the base unit of volume
var UnitOfVolume = Family{
	BaseUnitName: "cubic metre",
	Description:  "unit of volume",
}

// CubicMetreUnit is a suitable default value for a UnitOfVolume
var CubicMetreUnit = Unit{0, 0, 1,
	UnitOfVolume,
	"m\u00B3", UnitOfVolume.BaseUnitName, "cubic metres",
	"a metric measure of volume."}

// LitreUnit is an alternative suitable default value for a UnitOfVolume
var LitreUnit = Unit{0, 0, 1e-3,
	UnitOfVolume,
	"l", "litre", "litres",
	"a metric measure of volume."}

// VolumeNames maps names to units of volume
var VolumeNames = map[string]Unit{
	// metric
	"cc": {0, 0, m * 1e-3,
		UnitOfVolume,
		"cc", "cubic centimetre", "cubic centimetres",
		"a metric measure of volume."},
	"cubic-centimetre": {0, 0, m * 1e-3,
		UnitOfVolume,
		"cc", "cubic centimetre", "cubic centimetres",
		"a metric measure of volume."},
	"cubic centimetre": {0, 0, m * 1e-3,
		UnitOfVolume,
		"cc", "cubic centimetre", "cubic centimetres",
		"a metric measure of volume."},
	"cubic-metre": CubicMetreUnit,
	"cubic metre": CubicMetreUnit,
	"cubic-kilometre": {0, 0, _G,
		UnitOfVolume,
		"km\u00B3", "cubic kilometre", "cubic kilometres",
		"a metric measure of volume."},
	"cubic kilometre": {0, 0, _G,
		UnitOfVolume,
		"km\u00B3", "cubic kilometre", "cubic kilometres",
		"a metric measure of volume."},

	"yl": {0, 0, y * 1e-3,
		UnitOfVolume,
		"yl", "yoctolitre", "yoctolitres", "a metric measure of volume."},
	"zl": {0, 0, z * 1e-3,
		UnitOfVolume,
		"zl", "zeptolitre", "zeptolitres", "a metric measure of volume."},
	"al": {0, 0, a * 1e-3,
		UnitOfVolume,
		"al", "attolitre", "attolitres", "a metric measure of volume."},
	"fl": {0, 0, f * 1e-3,
		UnitOfVolume,
		"fl", "femtolitre", "femtolitres", "a metric measure of volume."},
	"pl": {0, 0, p * 1e-3,
		UnitOfVolume,
		"pl", "picolitre", "picolitres", "a metric measure of volume."},
	"nl": {0, 0, n * 1e-3,
		UnitOfVolume,
		"nl", "nanolitre", "nanolitres", "a metric measure of volume."},
	"ul": {0, 0, u * 1e-3,
		UnitOfVolume,
		"ul", "microlitre", "microlitres", "a metric measure of volume."},
	"ml": {0, 0, m * 1e-3,
		UnitOfVolume,
		"ml", "millilitre", "millilitres", "a metric measure of volume."},
	"cl": {0, 0, c * 1e-3,
		UnitOfVolume,
		"cl", "centilitre", "centilitres", "a metric measure of volume."},
	"dl": {0, 0, d * 1e-3,
		UnitOfVolume,
		"dl", "decilitre", "decilitres", "a metric measure of volume."},
	"litre": LitreUnit,
	// for American users
	"liter": LitreUnit, // nolint:misspell
	"dal": {0, 0, da * 1e-3,
		UnitOfVolume,
		"dal", "decalitre", "decalitres", "a metric measure of volume."},
	"hl": {0, 0, h * 1e-3,
		UnitOfVolume,
		"hl", "hectolitre", "hectolitres", "a metric measure of volume."},
	"kl": {0, 0, k * 1e-3,
		UnitOfVolume,
		"kl", "kilolitre", "kilolitres", "a metric measure of volume."},
	"Ml": {0, 0, _M * 1e-3,
		UnitOfVolume,
		"Ml", "megalitre", "megalitres", "a metric measure of volume."},
	"Gl": {0, 0, _G * 1e-3,
		UnitOfVolume,
		"Gl", "gigalitre", "gigalitres", "a metric measure of volume."},
	"Tl": {0, 0, _T * 1e-3,
		UnitOfVolume,
		"Tl", "teralitre", "teralitres", "a metric measure of volume."},
	"Pl": {0, 0, _P * 1e-3,
		UnitOfVolume,
		"Pl", "petalitre", "petalitres", "a metric measure of volume."},
	"El": {0, 0, _E * 1e-3,
		UnitOfVolume,
		"El", "exalitre", "exalitres", "a metric measure of volume."},
	"Zl": {0, 0, _Z * 1e-3,
		UnitOfVolume,
		"Zl", "zettalitre", "zettalitres", "a metric measure of volume."},
	"Yl": {0, 0, _Y * 1e-3,
		UnitOfVolume,
		"Yl", "yottalitre", "yottalitres", "a metric measure of volume."},

	// bottle
	"bottle-wine": {0, 0, bottleToCubicMetre,
		UnitOfVolume,
		"bottle", "bottle (wine)", "bottles (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade. Note that there is a now obsolete bottle size of" +
			" 70cl with corresponding multiples."},
	"magnum": {0, 0, bottleToCubicMetre * 2, UnitOfVolume,
		"Magnum", "Magnum (wine)", "Magnums (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},
	"marie-jeanne": {0, 0, bottleToCubicMetre * 3, UnitOfVolume,
		"Marie Jeanne", "Marie Jeanne (wine)", "Marie Jeannes (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},
	"tappit-hen": {0, 0, bottleToCubicMetre * 3, UnitOfVolume,
		"Tappit Hen", "Tappit Hen (port)", "Tappit Hens (port)",
		"a measure of volume typically used in the wine and spirits" +
			" trade. This is specifically for port. The name" +
			" derives from the shape of the bottle which is said to" +
			" resemble a hen."},
	"tregnum": {0, 0, bottleToCubicMetre * 3, UnitOfVolume,
		"Tregnum", "Tregnum (port)", "Tregnums (port)",
		"a measure of volume typically used in the wine and spirits" +
			" trade, specifically for port."},
	"jeroboam": {0, 0, bottleToCubicMetre * 4, UnitOfVolume,
		"Jeroboam", "Jeroboam (wine)", "Jeroboams (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},
	"rehoboam": {0, 0, bottleToCubicMetre * 6, UnitOfVolume,
		"Rehoboam", "Rehoboam (wine)", "Rehoboams (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},
	"methuselah": {0, 0, bottleToCubicMetre * 8, UnitOfVolume,
		"Methuselah", "Methuselah (wine)", "Methuselahs (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},
	"salmanazar": {0, 0, bottleToCubicMetre * 12, UnitOfVolume,
		"Salmanazar", "Salmanazar (wine)", "Salmanazars (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},
	"balthazar": {0, 0, bottleToCubicMetre * 16, UnitOfVolume,
		"Balthazar", "Balthazar (wine)", "Balthazars (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},
	"nebuchadnezzar": {0, 0, bottleToCubicMetre * 20, UnitOfVolume,
		"Nebuchadnezzar", "Nebuchadnezzar (wine)", "Nebuchadnezzars (wine)",
		"a measure of volume typically used in the wine and spirits" +
			" trade."},

	// Imperial / US
	"cubic-inch": {0, 0, cubicInchToCubicMetre,
		UnitOfVolume,
		"in\u00B3", "cubic inch", "cubic inches",
		"an imperial measure of volume."},
	"cubic inch": {0, 0, cubicInchToCubicMetre,
		UnitOfVolume,
		"in\u00B3", "cubic inch", "cubic inches",
		"an imperial measure of volume."},
	"cubic-foot": {0, 0, cubicFootToCubicMetre,
		UnitOfVolume,
		"ft\u00B3", "cubic foot", "cubic feet",
		"an imperial measure of volume."},
	"cubic foot": {0, 0, cubicFootToCubicMetre,
		UnitOfVolume,
		"ft\u00B3", "cubic foot", "cubic feet",
		"an imperial measure of volume."},
	"cubic-yard": {0, 0, cubicYardToCubicMetre,
		UnitOfVolume,
		"yd\u00B3", "cubic yard", "cubic yards",
		"an imperial measure of volume."},
	"cubic yard": {0, 0, cubicYardToCubicMetre,
		UnitOfVolume,
		"yd\u00B3", "cubic yard", "cubic yards",
		"an imperial measure of volume."},

	"perch (masonry)": {0, 0, 16.5 * 1.5 * 1 * cubicFootToCubicMetre,
		UnitOfVolume,
		"masonry perch", "masonry perch", "masonry perches",
		"a traditional measure of volume for stone and other" +
			" masonry. Note that this refers to the measure" +
			" for dressed stone, a masonry perch of brickwork" +
			" or rubble is a third smaller being one perch" +
			" long by one foot high by one foot wide." +
			" Note that a perch can also be a" +
			" measure of length or area."},

	"minim": {0, 0, fluidOzToCubicMetre / 480, UnitOfVolume,
		"min", "minim", "minims",
		"an imperial measure of volume (an apothecaries' measure)."},
	"fluid-scruple": {0, 0, fluidOzToCubicMetre / 24, UnitOfVolume,
		"fl s", "fluid scruple", "fluid scruples",
		"an imperial measure of volume (an apothecaries' measure)."},
	"fluid-drachm": {0, 0, fluidOzToCubicMetre / 8, UnitOfVolume,
		"fl dr", "fluid drachm", "fluid drachms",
		"an imperial measure of volume (an apothecaries' measure)."},

	"fluid-ounce": {0, 0, fluidOzToCubicMetre,
		UnitOfVolume,
		"fl oz", "fluid ounce", "fluid ounces",
		"an imperial measure of volume."},
	"gill": {0, 0, fluidOzToCubicMetre * 5,
		UnitOfVolume,
		"gi", "gill", "gills",
		"an imperial measure of volume."},
	"pint": {0, 0, fluidOzToCubicMetre * 20,
		UnitOfVolume,
		"pt", "pint", "pints",
		"an imperial measure of volume. Note that this refers to" +
			" the British pint of 20 fluid ounces." +
			" See also the US pint."},
	"quart": {0, 0, fluidOzToCubicMetre * 40,
		UnitOfVolume,
		"qt", "quart", "quarts",
		"an imperial measure of volume."},
	"gallon": {0, 0, gallonToCubicMetre,
		UnitOfVolume,
		"gal", "gallon", "gallons",
		"an imperial measure of volume."},

	"pin": {0, 0, gallonToCubicMetre * 4.5,
		UnitOfVolume,
		"pin", "pin", "pins",
		"an imperial measure of volume, used in the brewing industry" +
			" for large quantities of beer."},
	"firkin": {0, 0, gallonToCubicMetre * 9,
		UnitOfVolume,
		"firkin", "firkin", "firkins",
		"an imperial measure of volume, used in the brewing industry" +
			" for large quantities of beer."},
	"kilderkin": {0, 0, gallonToCubicMetre * 18,
		UnitOfVolume,
		"kilderkin", "kilderkin", "kilderkins",
		"an imperial measure of volume, used in the brewing industry" +
			" for large quantities of beer."},
	"barrel": {0, 0, gallonToCubicMetre * 36,
		UnitOfVolume,
		"barrel", "barrel", "barrels",
		"an imperial measure of volume, used in the brewing industry" +
			" for large quantities of beer."},
	"hogshead": {0, 0, gallonToCubicMetre * 54,
		UnitOfVolume,
		"hogshead", "hogshead", "hogsheads",
		"an imperial measure of volume, used in the brewing industry" +
			" for large quantities of beer."},

	"peck": {0, 0, gallonToCubicMetre * 2,
		UnitOfVolume,
		"peck", "peck", "pecks",
		"an imperial measure of volume, typically for dry goods."},
	"bushel": {0, 0, gallonToCubicMetre * 8,
		UnitOfVolume,
		"bushel", "bushel", "bushels",
		"an imperial measure of volume, typically for dry goods."},

	"teaspoon": {0, 0, usFluidOzToCubicMetre * 0.5 / 3.0, UnitOfVolume,
		"tsp", "teaspoon", "teaspoons",
		"a US Customary measure of volume."},
	"tablespoon": {0, 0, usFluidOzToCubicMetre * 0.5, UnitOfVolume,
		"Tbsp", "tablespoon", "tablespoons",
		"a US Customary measure of volume."},
	"US-fluid-ounce": {0, 0, usFluidOzToCubicMetre, UnitOfVolume,
		"US fl oz", "US fluid ounce", "US fluid ounces",
		"a US Customary measure of volume."},
	"US-shot": {0, 0, usFluidOzToCubicMetre * 1.5,
		UnitOfVolume,
		"shot", "US shot", "US shots",
		"a US Customary measure of volume."},
	"US-gill": {0, 0, usFluidOzToCubicMetre * 4,
		UnitOfVolume,
		"US gi", "US gill", "US gills",
		"a US Customary measure of volume."},
	"US-cup": {0, 0, usFluidOzToCubicMetre * 8,
		UnitOfVolume,
		"US cp", "US cup", "US cups",
		"a US Customary measure of volume."},
	"US-pint": {0, 0, usFluidOzToCubicMetre * 16,
		UnitOfVolume,
		"US pt", "US pint", "US pints",
		"a US Customary measure of volume."},
	"US-quart": {0, 0, usFluidOzToCubicMetre * 32,
		UnitOfVolume,
		"US qt", "US quart", "US quarts",
		"a US Customary measure of volume."},
	"US-wet-gallon": {0, 0, usFluidOzToCubicMetre * 128,
		UnitOfVolume,
		"US gal", "US wet gallon", "US wet gallons",
		"a US Customary measure of volume."},
	"barrel-oil": {0, 0, usFluidOzToCubicMetre * 128 * 42,
		UnitOfVolume,
		"bbl(oil)", "oil barrel", "oil barrels",
		"a US Customary measure of volume."},
	"US-dry-gallon": {0, 0, usDryGallonToCubicMetre,
		UnitOfVolume,
		"US dry gal", "US dry gallon", "US dry gallons",
		"a US Customary measure of volume, typically for dry goods."},
	"US-bushel": {0, 0, usDryGallonToCubicMetre * 8,
		UnitOfVolume,
		"US bushel", "US bushel", "US bushels",
		"a US Customary measure of volume, typically for dry goods."},
	"US-dry-pint": {0, 0, usDryGallonToCubicMetre * 0.125,
		UnitOfVolume,
		"US dry pint", "US dry pint", "US dry pints",
		"a US Customary measure of volume, typically for dry goods."},

	// Colloquial
	"20ft shipping container": {0, 0, 33.14,
		UnitOfVolume,
		"vol. of a TEU",
		"volume of a 20ft shipping container",
		"volumes of a 20ft shipping container",
		"a colloquial term referencing the volume of a standard 20 foot" +
			" shipping container" +
			" presumably to give an impression of scale." +
			" The TEU abbreviation stands for Twenty foot Equivalent Unit."},
}
