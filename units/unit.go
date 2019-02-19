package units

import "errors"

// Unit is the type of a unit, it records the base unit and a description.
// The base unit is the unit in terms of which any multiples are defined.
type Unit struct {
	BaseUnitName string
	Description  string
}

// UnitDetails bundles together the Unit and the associated collection of
// named alternative units
type UnitDetails struct {
	U    Unit
	AltU map[string]Mult
}

// These constants should be used when retrieving unit details
const (
	Dimensionless = "dimensionless"
	Time          = "time"
	Data          = "data"
	Distance      = "distance"
	Length        = "distance"
	Area          = "area"
	Volume        = "volume"
	Mass          = "mass"
)

var validUnits = map[string]UnitDetails{
	Dimensionless: {NoUnit, NoUnitNames},
	Time:          {UnitOfTime, TimeNames},
	Data:          {UnitOfData, DataNames},
	Distance:      {UnitOfDistance, DistanceNames},
	Area:          {UnitOfArea, AreaNames},
	Volume:        {UnitOfVolume, VolumeNames},
	Mass:          {UnitOfMass, MassNames},
}

// GetUnitDetails retrieves the unit details. The error value will be non-nil
// if the name is not recognised.
func GetUnitDetails(name string) (UnitDetails, error) {
	u, ok := validUnits[name]
	if !ok {
		return UnitDetails{}, errors.New("no such unit type '" + name + "'")
	}
	return u, nil
}
