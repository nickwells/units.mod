package units

const (
	bottleToCubicMetre      = 0.75e-3
	gallonToCubicMetre      = 4.54609e-3
	fluidOzToCubicMetre     = gallonToCubicMetre / 160
	usFluidOzToCubicMetre   = 29.5735295625e-6
	usDryGallonToCubicMetre = 4.40488377086e-3
)

// UnitOfVolume represents the base unit of volume
var UnitOfVolume = Unit{
	BaseUnitName: "cubic metre",
	Description:  "unit of volume",
}

// CubicMetreMult is a suitable default value for a Mult of UnitOfVolume
var CubicMetreMult = Mult{1, UnitOfVolume, "m\u00B3", UnitOfVolume.BaseUnitName}

// LitreMult is an alternative suitable default value for a Mult of UnitOfVolume
var LitreMult = Mult{1e-3, UnitOfVolume, "l", "litre"}

var VolumeNames = map[string]Mult{
	"cc":               {m * 1e-3, UnitOfVolume, "cc", "cubic centimetre"},
	"cubic-centimetre": {m * 1e-3, UnitOfVolume, "cc", "cubic centimetre"},
	"cubic centimetre": {m * 1e-3, UnitOfVolume, "cc", "cubic centimetre"},
	"cubic-metre":      CubicMetreMult,
	"cubic metre":      CubicMetreMult,
	"cubic-kilometre":  {_G, UnitOfVolume, "km\u00B3", "cubic kilometre"},
	"cubic kilometre":  {_G, UnitOfVolume, "km\u00B3", "cubic kilometre"},

	"yl":    {y * 1e-3, UnitOfVolume, "yl", "yoctolitre"},
	"zl":    {z * 1e-3, UnitOfVolume, "zl", "zeptolitre"},
	"al":    {a * 1e-3, UnitOfVolume, "al", "attolitre"},
	"fl":    {f * 1e-3, UnitOfVolume, "fl", "femtolitre"},
	"pl":    {p * 1e-3, UnitOfVolume, "pl", "picolitre"},
	"nl":    {n * 1e-3, UnitOfVolume, "nl", "nanolitre"},
	"ul":    {u * 1e-3, UnitOfVolume, "ul", "microlitre"},
	"ml":    {m * 1e-3, UnitOfVolume, "ml", "millilitre"},
	"cl":    {c * 1e-3, UnitOfVolume, "cl", "centilitre"},
	"dl":    {d * 1e-3, UnitOfVolume, "dl", "decilitre"},
	"litre": LitreMult,
	"liter": LitreMult, // for American users
	"dal":   {da * 1e-3, UnitOfVolume, "dal", "decalitre"},
	"hl":    {h * 1e-3, UnitOfVolume, "hl", "hectolitre"},
	"kl":    {k * 1e-3, UnitOfVolume, "kl", "kilolitre"},
	"Ml":    {_M * 1e-3, UnitOfVolume, "Ml", "megalitre"},
	"Gl":    {_G * 1e-3, UnitOfVolume, "Gl", "gigalitre"},
	"Tl":    {_T * 1e-3, UnitOfVolume, "Tl", "teralitre"},
	"Pl":    {_P * 1e-3, UnitOfVolume, "Pl", "petalitre"},
	"El":    {_E * 1e-3, UnitOfVolume, "El", "exalitre"},
	"Zl":    {_Z * 1e-3, UnitOfVolume, "Zl", "zettalitre"},
	"Yl":    {_Y * 1e-3, UnitOfVolume, "Yl", "yottalitre"},

	"bottle-wine": {bottleToCubicMetre, UnitOfVolume, "bottle", "bottle (wine)"},
	"magnum": {bottleToCubicMetre * 2, UnitOfVolume,
		"Magnum", "Magnum (wine)"},
	"marie-jeanne": {bottleToCubicMetre * 3, UnitOfVolume,
		"Marie Jeanne", "Marie Jeanne (wine)"},
	"tappit-hen": {bottleToCubicMetre * 3, UnitOfVolume,
		"Tappit Hen", "Tappit Hen (port)"},
	"tregnum": {bottleToCubicMetre * 3, UnitOfVolume,
		"Tregnum", "Tregnum (port)"},
	"jeroboam": {bottleToCubicMetre * 4, UnitOfVolume,
		"Jeroboam", "Jeroboam (wine)"},
	"rehoboam": {bottleToCubicMetre * 6, UnitOfVolume,
		"Rehoboam", "Rehoboam (wine)"},
	"methuselah": {bottleToCubicMetre * 8, UnitOfVolume,
		"Methuselah", "Methuselah (wine)"},
	"salmanazar": {bottleToCubicMetre * 12, UnitOfVolume,
		"Salmanazar", "Salmanazar (wine)"},
	"balthazar": {bottleToCubicMetre * 16, UnitOfVolume,
		"Balthazar", "Balthazar (wine)"},
	"nebuchadnezzar": {bottleToCubicMetre * 20, UnitOfVolume,
		"Nebuchadnezzar", "Nebuchadnezzar (wine)"},

	"minim": {fluidOzToCubicMetre / 480, UnitOfVolume,
		"min", "minim"},
	"fluid-scruple": {fluidOzToCubicMetre / 24, UnitOfVolume,
		"fl s", "fluid scruple"},
	"fluid-drachm": {fluidOzToCubicMetre / 8, UnitOfVolume,
		"fl dr", "fluid drachm"},
	"fluid-ounce": {fluidOzToCubicMetre, UnitOfVolume, "fl oz", "fluid ounce"},
	"gill":        {fluidOzToCubicMetre * 5, UnitOfVolume, "gi", "gill"},
	"pint":        {fluidOzToCubicMetre * 20, UnitOfVolume, "pt", "pint"},
	"quart":       {fluidOzToCubicMetre * 40, UnitOfVolume, "qt", "quart"},
	"gallon":      {gallonToCubicMetre, UnitOfVolume, "gal", "gallon"},
	"pin":         {gallonToCubicMetre * 4.5, UnitOfVolume, "pin", "pin"},
	"firkin":      {gallonToCubicMetre * 9, UnitOfVolume, "firkin", "firkin"},
	"kilderkin":   {gallonToCubicMetre * 18, UnitOfVolume, "kilderkin", "kilderkin"},
	"barrel":      {gallonToCubicMetre * 36, UnitOfVolume, "barrel", "barrel"},
	"hogshead":    {gallonToCubicMetre * 54, UnitOfVolume, "hogshead", "hogshead"},

	"peck":   {gallonToCubicMetre * 2, UnitOfVolume, "peck", "peck"},
	"bushel": {gallonToCubicMetre * 8, UnitOfVolume, "bushel", "bushel"},

	"teaspoon": {usFluidOzToCubicMetre * 0.5 / 3.0, UnitOfVolume,
		"tsp", "teaspoon"},
	"tablespoon": {usFluidOzToCubicMetre * 0.5, UnitOfVolume,
		"Tbsp", "tablespoon"},
	"us-fluid-ounce": {usFluidOzToCubicMetre, UnitOfVolume,
		"US fl oz", "US fluid ounce"},
	"us-shot":  {usFluidOzToCubicMetre * 1.5, UnitOfVolume, "shot", "US shot"},
	"us-gill":  {usFluidOzToCubicMetre * 4, UnitOfVolume, "US gi", "US gill"},
	"us-cup":   {usFluidOzToCubicMetre * 8, UnitOfVolume, "US cp", "US cup"},
	"us-pint":  {usFluidOzToCubicMetre * 16, UnitOfVolume, "US pt", "US pint"},
	"us-quart": {usFluidOzToCubicMetre * 32, UnitOfVolume, "US qt", "US quart"},
	"us-wet-gallon": {usFluidOzToCubicMetre * 128, UnitOfVolume,
		"US gal", "US wet gallon"},
	"barrel-oil": {usFluidOzToCubicMetre * 128 * 42, UnitOfVolume,
		"bbl(oil)", "oil barrel"},
	"us-dry-gallon": {usDryGallonToCubicMetre, UnitOfVolume,
		"US dry gal", "US dry gallon"},
	"us-bushel": {usDryGallonToCubicMetre * 8, UnitOfVolume,
		"US bushel", "US bushel"},
	"us-dry-pint": {usDryGallonToCubicMetre * 0.125, UnitOfVolume,
		"US dry pint", "US dry pint"},
}
