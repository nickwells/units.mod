package units

const yard2ToMetre2 = yardToMetre * yardToMetre

// UnitOfArea represents the base unit of area
var UnitOfArea = Family{
	BaseUnitName: "square metre",
	Description:  "unit of area",
}

// SquareMetreUnit is a suitable default value for a UnitOfArea
var SquareMetreUnit = Unit{0, 0, 1, UnitOfArea, "m\u00B2", UnitOfArea.BaseUnitName, "square metres"}

// AreaNames maps names to units of area
var AreaNames = map[string]Unit{
	"square-metre":     SquareMetreUnit,
	"square metre":     SquareMetreUnit,
	"hectare":          {0, 0, 1e4, UnitOfArea, "ha", "hectare", "hectares"},
	"square-kilometre": {0, 0, 1e6, UnitOfArea, "km\u00B2", "square kilometre", "square kilometres"},
	"square kilometre": {0, 0, 1e6, UnitOfArea, "km\u00B2", "square kilometre", "square kilometres"},
	"square-yard":      {0, 0, yard2ToMetre2, UnitOfArea, "yd\u00B2", "square yard", "square yards"},
	"square yard":      {0, 0, yard2ToMetre2, UnitOfArea, "yd\u00B2", "square yard", "square yards"},
	"square-mile":      {0, 0, yard2ToMetre2 * 1760 * 1760, UnitOfArea, "mi\u00B2", "square mile", "square miles"},
	"square mile":      {0, 0, yard2ToMetre2 * 1760 * 1760, UnitOfArea, "mi\u00B2", "square mile", "square miles"},
	"rood":             {0, 0, yard2ToMetre2 * 1210, UnitOfArea, "rood", "rood", "roods"},
	"acre":             {0, 0, yard2ToMetre2 * 4840, UnitOfArea, "acre", "acre", "acres"},
	"oxgang":           {0, 0, yard2ToMetre2 * 4840 * 15, UnitOfArea, "oxgang", "oxgang", "oxgangs"},
	"Wales":            {0, 0, 20779 * 1e6, UnitOfArea, "Wales", "Area the size of Wales", "Areas the size of Wales"},
}
