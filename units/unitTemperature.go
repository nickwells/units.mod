package units

// UnitOfTemperature represents the base unit of temperature
var UnitOfTemperature = Family{
	BaseUnitName: "degree Celsius",
	Description:  "unit of temperature",
}

// DegCUnit is a suitable default value for a UnitOfTemperature
var DegCUnit = Unit{0, 0, 1, UnitOfTemperature, "℃", UnitOfTemperature.BaseUnitName, "degrees Celsius"}
var DegFUnit = Unit{0, 32, 5.0 / 9.0, UnitOfTemperature, "℉", "degree Fahrenheit", "degrees Fahrenheit"}
var DegKUnit = Unit{0, 237.15, 1, UnitOfTemperature, "K", "Kelvin", "Kelvin"}

// TemperatureNames maps names to units of temperature
var TemperatureNames = map[string]Unit{
	"C": DegCUnit,
	"F": DegFUnit,
	"K": DegKUnit,
}
