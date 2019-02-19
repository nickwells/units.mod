package units

// UnitOfMass represents the base unit of mass
var UnitOfMass = Unit{
	BaseUnitName: "gram",
	Description:  "unit of mass",
}

// GramMult is a suitable default value for a Mult of UnitOfMass
var GramMult = Mult{1, UnitOfMass, "g", UnitOfMass.BaseUnitName}

const (
	grainToGram = 0.06479891
	ounceToGram = grainToGram * 437.5
	poundToGram = ounceToGram * 16
)

var MassNames = map[string]Mult{
	"yg":        {y, UnitOfMass, "yg", "yoctogram"},
	"zg":        {z, UnitOfMass, "zg", "zeptogram"},
	"ag":        {a, UnitOfMass, "ag", "attogram"},
	"fg":        {f, UnitOfMass, "fg", "femtogram"},
	"pg":        {p, UnitOfMass, "pg", "picogram"},
	"ng":        {n, UnitOfMass, "ng", "nanogram"},
	"ug":        {u, UnitOfMass, "ug", "microgram"},
	"mg":        {m, UnitOfMass, "mg", "milligram"},
	"cg":        {c, UnitOfMass, "cg", "centigram"},
	"dg":        {d, UnitOfMass, "dg", "decigram"},
	"gram":      GramMult,
	"dag":       {da, UnitOfMass, "dag", "decagram"},
	"hg":        {h, UnitOfMass, "hg", "hectogram"},
	"kg":        {k, UnitOfMass, "kg", "kilogram"},
	"Mg":        {_M, UnitOfMass, "Mg", "megagram"},
	"tonne":     {_M, UnitOfMass, "T", "tonne"},
	"Gg":        {_G, UnitOfMass, "Gg", "gigagram"},
	"kilotonne": {_G, UnitOfMass, "KT", "kilotonne"},
	"Tg":        {_T, UnitOfMass, "Tg", "teragram"},
	"megatonne": {_T, UnitOfMass, "MT", "megatonne"},
	"Pg":        {_P, UnitOfMass, "Pg", "petagram"},
	"Eg":        {_E, UnitOfMass, "Eg", "exagram"},
	"Zg":        {_Z, UnitOfMass, "Zg", "zettagram"},
	"Yg":        {_Y, UnitOfMass, "Yg", "yottagram"},

	"grain":      {grainToGram, UnitOfMass, "gr", "grain"},
	"troy-ounce": {grainToGram * 480, UnitOfMass, "t oz", "troy ounce"},
	"scruple":    {grainToGram * 20, UnitOfMass, "scruple", "scruple"},
	"dram":       {grainToGram * 60, UnitOfMass, "dr", "dram"},
	"drachm":     {grainToGram * 60, UnitOfMass, "dr", "drachm"},

	"ounce":         {ounceToGram, UnitOfMass, "oz", "ounce"},
	"pound":         {poundToGram, UnitOfMass, "lb", "pound"},
	"stone":         {poundToGram * 14, UnitOfMass, "st", "stone"},
	"hundredweight": {poundToGram * 112, UnitOfMass, "cwt", "hundredweight"},
	"short-hundredweight": {poundToGram * 100, UnitOfMass,
		"cental", "short hundredweight"},
	"imperial-ton": {poundToGram * 112 * 20, UnitOfMass, "t", "imperial ton"},
	"short-ton":    {poundToGram * 2000, UnitOfMass, "short ton", "short ton"},

	"earth-mass": {5.9722e27, UnitOfMass, "M\u2295", "Earth Mass"},
	"solar-mass": {1.98847e33, UnitOfMass, "M\u2299", "Solar Mass"},
}
