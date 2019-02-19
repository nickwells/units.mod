package units

const yard2ToMetre2 = yardToMetre * yardToMetre

// UnitOfArea represents the base unit of area
var UnitOfArea = Family{
	BaseUnitName: "square metre",
	Description:  "unit of area",
}

// SquareMetreUnit is a suitable default value for a UnitOfArea
var SquareMetreUnit = Unit{1, UnitOfArea, "m\u00B2", UnitOfArea.BaseUnitName}

var AreaNames = map[string]Unit{
	"square-metre":     SquareMetreUnit,
	"square metre":     SquareMetreUnit,
	"hectare":          {1e4, UnitOfArea, "ha", "hectare"},
	"square-kilometre": {1e6, UnitOfArea, "km\u00B2", "square kilometre"},
	"square kilometre": {1e6, UnitOfArea, "km\u00B2", "square kilometre"},
	"square-yard":      {yard2ToMetre2, UnitOfArea, "yd\u00B2", "square yard"},
	"square yard":      {yard2ToMetre2, UnitOfArea, "yd\u00B2", "square yard"},
	"rood":             {yard2ToMetre2 * 1210, UnitOfArea, "rood", "rood"},
	"acre":             {yard2ToMetre2 * 4840, UnitOfArea, "acre", "acre"},
	"oxgang":           {yard2ToMetre2 * 4840 * 15, UnitOfArea, "oxgang", "oxgang"},
	"Wales":            {20779 * 1e6, UnitOfArea, "Wales", "Area the size of Wales"},
}
