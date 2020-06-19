package units

// UnitOfEnergy represents the base unit of energy
var UnitOfEnergy = Family{
	BaseUnitName: "joule",
	Description:  "unit of energy",
}

// JouleUnit is a suitable default value for a UnitOfEnergy
var JouleUnit = Unit{0, 0, 1, UnitOfEnergy, "J", UnitOfEnergy.BaseUnitName, "joules"}

const (
	calorieToJoule     = 4.184
	footPoundToJoule   = 1.3558179483314004
	footPoundalToJoule = 0.0421401100938048

	// Note that there are many different variants of the British thermal
	// unit with conversion ratios to Joules varying between 1054.3503 (for
	// the Thermochemical variant) and 1059.67 (using the calorie value of
	// water at its maximum density, when it is at 3.9 ℃)
	BtuToJoule = 1055.06 // International Standard ISO 31-4
)

// EnergyNames maps names to units of energy
var EnergyNames = map[string]Unit{
	// SI
	"yJ":    {0, 0, y, UnitOfEnergy, "yJ", "yoctojoule", "yoctojoules"},
	"zJ":    {0, 0, z, UnitOfEnergy, "zJ", "zeptojoule", "zeptojoules"},
	"aJ":    {0, 0, a, UnitOfEnergy, "aJ", "attojoule", "attojoules"},
	"fJ":    {0, 0, f, UnitOfEnergy, "fJ", "femtojoule", "femtojoules"},
	"pJ":    {0, 0, p, UnitOfEnergy, "pJ", "picojoule", "picojoules"},
	"nJ":    {0, 0, n, UnitOfEnergy, "nJ", "nanojoule", "nanojoules"},
	"uJ":    {0, 0, u, UnitOfEnergy, "uJ", "microjoule", "microjoules"},
	"mJ":    {0, 0, m, UnitOfEnergy, "mJ", "millijoule", "millijoules"},
	"cJ":    {0, 0, c, UnitOfEnergy, "cJ", "centijoule", "centijoules"},
	"dJ":    {0, 0, d, UnitOfEnergy, "dJ", "decijoule", "decijoules"},
	"joule": JouleUnit,
	"daJ":   {0, 0, da, UnitOfEnergy, "daJ", "decajoule", "decajoules"},
	"hJ":    {0, 0, h, UnitOfEnergy, "hJ", "hectojoule", "hectojoules"},
	"kJ":    {0, 0, k, UnitOfEnergy, "kJ", "kilojoule", "kilojoules"},
	"MJ":    {0, 0, _M, UnitOfEnergy, "MJ", "megajoule", "megajoules"},
	"GJ":    {0, 0, _G, UnitOfEnergy, "GJ", "gigajoule", "gigajoules"},
	"TJ":    {0, 0, _T, UnitOfEnergy, "TJ", "terajoule", "terajoules"},
	"PJ":    {0, 0, _P, UnitOfEnergy, "PJ", "petajoule", "petajoules"},
	"EJ":    {0, 0, _E, UnitOfEnergy, "EJ", "exajoule", "exajoules"},
	"ZJ":    {0, 0, _Z, UnitOfEnergy, "ZJ", "zettajoule", "zettajoules"},
	"YJ":    {0, 0, _Y, UnitOfEnergy, "YJ", "yottajoule", "yottajoules"},

	"kWh": {0, 0, 3.6e6, UnitOfEnergy, "kW h", "kilowatt-hour", "killowatt-hours"},

	"erg": {0, 0, 1e7, UnitOfEnergy, "erg", "erg", "ergs"},
	// foe = ten to the fifty-one ergs - the energy released by a supernova
	"foe": {0, 0, 1e-44, UnitOfEnergy, "foe", "foe", "foes"},

	"cal":         {0, 0, calorieToJoule, UnitOfEnergy, "cal", "calorie", "calories"},
	"calorie":     {0, 0, calorieToJoule, UnitOfEnergy, "cal", "calorie", "calories"},
	"kcal":        {0, 0, calorieToJoule * 1000, UnitOfEnergy, "kcal", "kilocalorie", "kilocalories"},
	"kilocalorie": {0, 0, calorieToJoule * 1000, UnitOfEnergy, "kcal", "kilocalorie", "kilocalories"},

	"foot-pound":   {0, 0, footPoundToJoule, UnitOfEnergy, "ft lb", "foot-pound", "foot-pounds"},
	"foot-poundal": {0, 0, footPoundalToJoule, UnitOfEnergy, "ft pdl", "foot-poundal", "foot-poundals"},
	"BTU":          {0, 0, BtuToJoule, UnitOfEnergy, "Btu", "British thermal unit", "British thermal units"},
	"Btu":          {0, 0, BtuToJoule, UnitOfEnergy, "Btu", "British thermal unit", "British thermal units"},
	"therm":        {0, 0, BtuToJoule * 1e5, UnitOfEnergy, "therm", "therm", "therms"},

	"tonOfTNT": {0, 0, calorieToJoule * 1e9, UnitOfEnergy, "ton of TNT", "ton of TNT", "tons of TNT"},
}
