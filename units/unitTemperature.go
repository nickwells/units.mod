package units

// UnitOfTemperature represents the base unit of temperature
var UnitOfTemperature = Family{
	BaseUnitName: "degree Celsius",
	Description:  "unit of temperature",
}

const absZero = 273.15

// DegCUnit is a suitable default value for a UnitOfTemperature
var DegCUnit = Unit{0, 0, 1, UnitOfTemperature,
	"℃", UnitOfTemperature.BaseUnitName, "degrees Celsius"}
var DegKUnit = Unit{0, absZero, 1, UnitOfTemperature,
	"K", "Kelvin", "Kelvin"}

var DegFUnit = Unit{0, 32, 5.0 / 9.0, UnitOfTemperature,
	"℉", "degree Fahrenheit", "degrees Fahrenheit"}

// The following are antique measures of historical interest only (like
// Fahrenheit?)

// DegRaUnit (Rankine) is the Fahrenheit equivalent of Kelvin
var DegRaUnit = Unit{absZero, 0, 5.0 / 9.0, UnitOfTemperature,
	"°R", "degree Rankine", "degrees Rankine"}

var DegRoUnit = Unit{0, 7.5, 40.0 / 21.0, UnitOfTemperature,
	"°Rø", "degree Rømer", "degrees Rømer"}
var DegReUnit = Unit{0, 0, 4.0 / 5.0, UnitOfTemperature,
	"°Ré", "degree Réaumur", "degrees Réaumur"}
var DegNUnit = Unit{0, 0, 100.0 / 33.0, UnitOfTemperature,
	"°N", "degree Newton", "degrees Newton"}

// DegDUnit grows in the opposite direction to all the other units so that
// the larger the value, the colder it is
var DegDUnit = Unit{-100, 0, -2.0 / 3.0, UnitOfTemperature,
	"°D", "degree Delisle", "degrees Delisle"}

// TemperatureNames maps names to units of temperature
var TemperatureNames = map[string]Unit{
	"C":          DegCUnit,
	"Celsius":    DegCUnit,
	"Centigrade": DegCUnit,

	"F":          DegFUnit,
	"Fahrenheit": DegFUnit,

	"K":      DegKUnit,
	"Kelvin": DegKUnit,

	"Ra":      DegRaUnit,
	"Rankine": DegRaUnit,

	"Ro":     DegRoUnit,
	"Romer":  DegRoUnit,
	"Rømer":  DegRoUnit,
	"Roemer": DegRoUnit,

	"Re":      DegReUnit,
	"Réaumur": DegReUnit,
	"Reaumur": DegReUnit,

	"N":      DegNUnit,
	"Newton": DegNUnit,

	"D":       DegDUnit,
	"Delisle": DegDUnit,
}
