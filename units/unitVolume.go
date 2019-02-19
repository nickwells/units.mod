package units

const (
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
var CubicMetreUnit = Unit{0, 0, 1, UnitOfVolume, "m\u00B3", UnitOfVolume.BaseUnitName, "cubic metres"}

// LitreUnit is an alternative suitable default value for a UnitOfVolume
var LitreUnit = Unit{0, 0, 1e-3, UnitOfVolume, "l", "litre", "litres"}

// VolumeNames maps names to units of volume
var VolumeNames = map[string]Unit{
	"cc":               {0, 0, m * 1e-3, UnitOfVolume, "cc", "cubic centimetre", "cubic centimetres"},
	"cubic-centimetre": {0, 0, m * 1e-3, UnitOfVolume, "cc", "cubic centimetre", "cubic centimetres"},
	"cubic centimetre": {0, 0, m * 1e-3, UnitOfVolume, "cc", "cubic centimetre", "cubic centimetres"},
	"cubic-metre":      CubicMetreUnit,
	"cubic metre":      CubicMetreUnit,
	"cubic-kilometre":  {0, 0, _G, UnitOfVolume, "km\u00B3", "cubic kilometre", "cubic kilometres"},
	"cubic kilometre":  {0, 0, _G, UnitOfVolume, "km\u00B3", "cubic kilometre", "cubic kilometres"},

	"yl":    {0, 0, y * 1e-3, UnitOfVolume, "yl", "yoctolitre", "yoctolitres"},
	"zl":    {0, 0, z * 1e-3, UnitOfVolume, "zl", "zeptolitre", "zeptolitres"},
	"al":    {0, 0, a * 1e-3, UnitOfVolume, "al", "attolitre", "attolitres"},
	"fl":    {0, 0, f * 1e-3, UnitOfVolume, "fl", "femtolitre", "femtolitres"},
	"pl":    {0, 0, p * 1e-3, UnitOfVolume, "pl", "picolitre", "picolitres"},
	"nl":    {0, 0, n * 1e-3, UnitOfVolume, "nl", "nanolitre", "nanolitres"},
	"ul":    {0, 0, u * 1e-3, UnitOfVolume, "ul", "microlitre", "microlitres"},
	"ml":    {0, 0, m * 1e-3, UnitOfVolume, "ml", "millilitre", "millilitres"},
	"cl":    {0, 0, c * 1e-3, UnitOfVolume, "cl", "centilitre", "centilitres"},
	"dl":    {0, 0, d * 1e-3, UnitOfVolume, "dl", "decilitre", "decilitres"},
	"litre": LitreUnit,
	"liter": LitreUnit, // for American users
	"dal":   {0, 0, da * 1e-3, UnitOfVolume, "dal", "decalitre", "decalitres"},
	"hl":    {0, 0, h * 1e-3, UnitOfVolume, "hl", "hectolitre", "hectolitres"},
	"kl":    {0, 0, k * 1e-3, UnitOfVolume, "kl", "kilolitre", "kilolitres"},
	"Ml":    {0, 0, _M * 1e-3, UnitOfVolume, "Ml", "megalitre", "megalitres"},
	"Gl":    {0, 0, _G * 1e-3, UnitOfVolume, "Gl", "gigalitre", "gigalitres"},
	"Tl":    {0, 0, _T * 1e-3, UnitOfVolume, "Tl", "teralitre", "teralitres"},
	"Pl":    {0, 0, _P * 1e-3, UnitOfVolume, "Pl", "petalitre", "petalitres"},
	"El":    {0, 0, _E * 1e-3, UnitOfVolume, "El", "exalitre", "exalitres"},
	"Zl":    {0, 0, _Z * 1e-3, UnitOfVolume, "Zl", "zettalitre", "zettalitres"},
	"Yl":    {0, 0, _Y * 1e-3, UnitOfVolume, "Yl", "yottalitre", "yottalitres"},

	"bottle-wine": {0, 0, bottleToCubicMetre, UnitOfVolume, "bottle", "bottle (wine)", "bottles (wine)"},
	"magnum": {0, 0, bottleToCubicMetre * 2, UnitOfVolume,
		"Magnum", "Magnum (wine)", "Magnums (wine)"},
	"marie-jeanne": {0, 0, bottleToCubicMetre * 3, UnitOfVolume,
		"Marie Jeanne", "Marie Jeanne (wine)", "Marie Jeannes (wine)"},
	"tappit-hen": {0, 0, bottleToCubicMetre * 3, UnitOfVolume,
		"Tappit Hen", "Tappit Hen (port)", "Tappit Hens (port)"},
	"tregnum": {0, 0, bottleToCubicMetre * 3, UnitOfVolume,
		"Tregnum", "Tregnum (port)", "Tregnums (port)"},
	"jeroboam": {0, 0, bottleToCubicMetre * 4, UnitOfVolume,
		"Jeroboam", "Jeroboam (wine)", "Jeroboams (wine)"},
	"rehoboam": {0, 0, bottleToCubicMetre * 6, UnitOfVolume,
		"Rehoboam", "Rehoboam (wine)", "Rehoboams (wine)"},
	"methuselah": {0, 0, bottleToCubicMetre * 8, UnitOfVolume,
		"Methuselah", "Methuselah (wine)", "Methuselahs (wine)"},
	"salmanazar": {0, 0, bottleToCubicMetre * 12, UnitOfVolume,
		"Salmanazar", "Salmanazar (wine)", "Salmanazars (wine)"},
	"balthazar": {0, 0, bottleToCubicMetre * 16, UnitOfVolume,
		"Balthazar", "Balthazar (wine)", "Balthazars (wine)"},
	"nebuchadnezzar": {0, 0, bottleToCubicMetre * 20, UnitOfVolume,
		"Nebuchadnezzar", "Nebuchadnezzar (wine)", "Nebuchadnezzars (wine)"},

	"minim": {0, 0, fluidOzToCubicMetre / 480, UnitOfVolume,
		"min", "minim", "minims"},
	"fluid-scruple": {0, 0, fluidOzToCubicMetre / 24, UnitOfVolume,
		"fl s", "fluid scruple", "fluid scruples"},
	"fluid-drachm": {0, 0, fluidOzToCubicMetre / 8, UnitOfVolume,
		"fl dr", "fluid drachm", "fluid drachms"},
	"fluid-ounce": {0, 0, fluidOzToCubicMetre, UnitOfVolume, "fl oz", "fluid ounce", "fluid ounces"},
	"gill":        {0, 0, fluidOzToCubicMetre * 5, UnitOfVolume, "gi", "gill", "gills"},
	"pint":        {0, 0, fluidOzToCubicMetre * 20, UnitOfVolume, "pt", "pint", "pints"},
	"quart":       {0, 0, fluidOzToCubicMetre * 40, UnitOfVolume, "qt", "quart", "quarts"},
	"gallon":      {0, 0, gallonToCubicMetre, UnitOfVolume, "gal", "gallon", "gallons"},
	"pin":         {0, 0, gallonToCubicMetre * 4.5, UnitOfVolume, "pin", "pin", "pins"},
	"firkin":      {0, 0, gallonToCubicMetre * 9, UnitOfVolume, "firkin", "firkin", "firkins"},
	"kilderkin":   {0, 0, gallonToCubicMetre * 18, UnitOfVolume, "kilderkin", "kilderkin", "kilderkins"},
	"barrel":      {0, 0, gallonToCubicMetre * 36, UnitOfVolume, "barrel", "barrel", "barrels"},
	"hogshead":    {0, 0, gallonToCubicMetre * 54, UnitOfVolume, "hogshead", "hogshead", "hogsheads"},

	"peck":   {0, 0, gallonToCubicMetre * 2, UnitOfVolume, "peck", "peck", "pecks"},
	"bushel": {0, 0, gallonToCubicMetre * 8, UnitOfVolume, "bushel", "bushel", "bushels"},

	"teaspoon": {0, 0, usFluidOzToCubicMetre * 0.5 / 3.0, UnitOfVolume,
		"tsp", "teaspoon", "teaspoons"},
	"tablespoon": {0, 0, usFluidOzToCubicMetre * 0.5, UnitOfVolume,
		"Tbsp", "tablespoon", "tablespoons"},
	"us-fluid-ounce": {0, 0, usFluidOzToCubicMetre, UnitOfVolume,
		"US fl oz", "US fluid ounce", "US fluid ounces"},
	"us-shot":  {0, 0, usFluidOzToCubicMetre * 1.5, UnitOfVolume, "shot", "US shot", "US shots"},
	"us-gill":  {0, 0, usFluidOzToCubicMetre * 4, UnitOfVolume, "US gi", "US gill", "US gills"},
	"us-cup":   {0, 0, usFluidOzToCubicMetre * 8, UnitOfVolume, "US cp", "US cup", "US cups"},
	"us-pint":  {0, 0, usFluidOzToCubicMetre * 16, UnitOfVolume, "US pt", "US pint", "US pints"},
	"us-quart": {0, 0, usFluidOzToCubicMetre * 32, UnitOfVolume, "US qt", "US quart", "US quarts"},
	"us-wet-gallon": {0, 0, usFluidOzToCubicMetre * 128, UnitOfVolume,
		"US gal", "US wet gallon", "US wet gallons"},
	"barrel-oil": {0, 0, usFluidOzToCubicMetre * 128 * 42, UnitOfVolume,
		"bbl(oil)", "oil barrel", "oil barrels"},
	"us-dry-gallon": {0, 0, usDryGallonToCubicMetre, UnitOfVolume,
		"US dry gal", "US dry gallon", "US dry gallons"},
	"us-bushel": {0, 0, usDryGallonToCubicMetre * 8, UnitOfVolume,
		"US bushel", "US bushel", "US bushels"},
	"us-dry-pint": {0, 0, usDryGallonToCubicMetre * 0.125, UnitOfVolume,
		"US dry pint", "US dry pint", "US dry pints"},
}
