package units

// UnitOfMass represents the base unit of mass
var UnitOfMass = Family{
	BaseUnitName: "gram",
	Description:  "unit of mass",
}

// GramUnit is a suitable default value for a UnitOfMass
var GramUnit = Unit{0, 0, 1, UnitOfMass, "g", UnitOfMass.BaseUnitName, "grams"}

const (
	grainToGram = 0.06479891
	ounceToGram = grainToGram * 437.5
	poundToGram = ounceToGram * 16
)

// MassNames maps names to units of mass
var MassNames = map[string]Unit{
	"yg":        {0, 0, y, UnitOfMass, "yg", "yoctogram", "yoctograms"},
	"zg":        {0, 0, z, UnitOfMass, "zg", "zeptogram", "zeptograms"},
	"ag":        {0, 0, a, UnitOfMass, "ag", "attogram", "attograms"},
	"fg":        {0, 0, f, UnitOfMass, "fg", "femtogram", "femtograms"},
	"pg":        {0, 0, p, UnitOfMass, "pg", "picogram", "picograms"},
	"ng":        {0, 0, n, UnitOfMass, "ng", "nanogram", "nanograms"},
	"ug":        {0, 0, u, UnitOfMass, "ug", "microgram", "micrograms"},
	"mg":        {0, 0, m, UnitOfMass, "mg", "milligram", "milligrams"},
	"cg":        {0, 0, c, UnitOfMass, "cg", "centigram", "centigrams"},
	"dg":        {0, 0, d, UnitOfMass, "dg", "decigram", "decigrams"},
	"gram":      GramUnit,
	"dag":       {0, 0, da, UnitOfMass, "dag", "decagram", "decagrams"},
	"hg":        {0, 0, h, UnitOfMass, "hg", "hectogram", "hectograms"},
	"kg":        {0, 0, k, UnitOfMass, "kg", "kilogram", "kilograms"},
	"Mg":        {0, 0, _M, UnitOfMass, "Mg", "megagram", "megagrams"},
	"tonne":     {0, 0, _M, UnitOfMass, "T", "tonne", "tonnes"},
	"Gg":        {0, 0, _G, UnitOfMass, "Gg", "gigagram", "gigagrams"},
	"kilotonne": {0, 0, _G, UnitOfMass, "KT", "kilotonne", "kilotonnes"},
	"Tg":        {0, 0, _T, UnitOfMass, "Tg", "teragram", "teragrams"},
	"megatonne": {0, 0, _T, UnitOfMass, "MT", "megatonne", "megatonnes"},
	"Pg":        {0, 0, _P, UnitOfMass, "Pg", "petagram", "petagrams"},
	"Eg":        {0, 0, _E, UnitOfMass, "Eg", "exagram", "exagrams"},
	"Zg":        {0, 0, _Z, UnitOfMass, "Zg", "zettagram", "zettagrams"},
	"Yg":        {0, 0, _Y, UnitOfMass, "Yg", "yottagram", "yottagrams"},

	"grain":      {0, 0, grainToGram, UnitOfMass, "gr", "grain", "grains"},
	"troy-ounce": {0, 0, grainToGram * 480, UnitOfMass, "t oz", "troy ounce", "troy ounces"},
	"scruple":    {0, 0, grainToGram * 20, UnitOfMass, "scruple", "scruple", "scruples"},
	"dram":       {0, 0, grainToGram * 60, UnitOfMass, "dr", "dram", "drams"},
	"drachm":     {0, 0, grainToGram * 60, UnitOfMass, "dr", "drachm", "drachms"},

	"ounce":         {0, 0, ounceToGram, UnitOfMass, "oz", "ounce", "ounces"},
	"pound":         {0, 0, poundToGram, UnitOfMass, "lb", "pound", "pounds"},
	"stone":         {0, 0, poundToGram * 14, UnitOfMass, "st", "stone", "stones"},
	"hundredweight": {0, 0, poundToGram * 112, UnitOfMass, "cwt", "hundredweight", "hundredweights"},
	"short-hundredweight": {0, 0, poundToGram * 100, UnitOfMass,
		"cental", "short hundredweight", "short hundredweights"},
	"imperial-ton": {0, 0, poundToGram * 112 * 20, UnitOfMass, "t", "imperial ton", "imperial tons"},
	"short-ton":    {0, 0, poundToGram * 2000, UnitOfMass, "short ton", "short ton", "short tons"},

	"earth-mass": {0, 0, 5.9722e27, UnitOfMass, "M\u2295", "Earth Mass", "Earth Masses"},
	"solar-mass": {0, 0, 1.98847e33, UnitOfMass, "M\u2299", "Solar Mass", "Solar Masses"},
}
