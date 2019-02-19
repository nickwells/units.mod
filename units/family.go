package units

import "errors"

// Family records details of a family of units such as units of area or units
// of mass, it records the name of the base unit and a description.  The base
// unit is the unit in terms of which any conversion factors are defined.
type Family struct {
	BaseUnitName string
	Description  string
}

// String returns a string representation of the Family object
func (f Family) String() string {
	return f.BaseUnitName + " (a " + f.Description + ")"
}

// UnitDetails bundles together the Family and the associated collection of
// named alternative units
type UnitDetails struct {
	Fam  Family
	AltU map[string]Unit
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
	Dimensionless: {Numeric, DimensionlessNames},
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
	ud, ok := validUnits[name]
	if !ok {
		return UnitDetails{}, errors.New("no such unit type '" + name + "'")
	}
	return ud, nil
}
