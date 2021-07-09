package units

import (
	"errors"
	"fmt"
	"sort"
)

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

// Alias records details about a particular Alias
type Alias struct {
	// UnitName is the name of the unit that this is aliasing
	UnitName string
	// Notes describes the alias itself
	Notes string
}

// UnitDetails bundles together the Family and the associated collection of
// named alternative units
type UnitDetails struct {
	Fam     Family
	AltU    map[string]Unit
	Aliases map[string]Alias
}

// These constants should be used when retrieving unit details
const (
	Dimensionless = "dimensionless"
	Time          = "time"
	Data          = "data"
	Distance      = "distance"
	Length        = "distance" // synonym
	Area          = "area"
	Volume        = "volume"
	Velocity      = "velocity"
	Speed         = "velocity" // synonym
	Mass          = "mass"
	Temperature   = "temperature"
	Angle         = "angle"
	Energy        = "energy"
)

var validUnits = map[string]UnitDetails{
	Dimensionless: {Numeric, DimensionlessNames, dimensionlessAliases},
	Time:          {UnitOfTime, TimeNames, timeAliases},
	Data:          {UnitOfData, DataNames, dataAliases},
	Distance:      {UnitOfDistance, DistanceNames, distanceAliases},
	Area:          {UnitOfArea, AreaNames, areaAliases},
	Volume:        {UnitOfVolume, VolumeNames, volumeAliases},
	Velocity:      {UnitOfVelocity, VelocityNames, velocityAliases},
	Mass:          {UnitOfMass, MassNames, massAliases},
	Temperature:   {UnitOfTemperature, TemperatureNames, temperatureAliases},
	Angle:         {UnitOfAngle, AngleNames, angleAliases},
	Energy:        {UnitOfEnergy, EnergyNames, energyAliases},
}

// reverseAlias holds a list of aliases for the family/unit entry
var reverseAlias = make(map[string]map[string][]string)

func init() {
	for f, ud := range validUnits {
		reverseAlias[f] = make(map[string][]string)
		for aliasName, alias := range ud.Aliases {
			if _, ok := ud.AltU[aliasName]; ok {
				panic(fmt.Errorf(
					"unit family %q: alias %q"+
						" already exists in the map of unit names",
					f, aliasName))
			}
			if _, ok := ud.AltU[alias.UnitName]; !ok {
				panic(fmt.Errorf(
					"unit family %q: alias %q references unit %q"+
						" which does not exist in the map of unit names",
					f, aliasName, alias))
			}

			reverseAlias[f][alias.UnitName] =
				append(reverseAlias[f][alias.UnitName], aliasName)
		}
	}
}

// GetAliases returns all the aliases for the given family/unit
func GetAliases(fName, uName string) []string {
	ua, ok := reverseAlias[fName]
	if !ok {
		return []string{}
	}
	return ua[uName]
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

// GetUnitDetailsOrPanic retrieves the unit details. It will panic if the
// name is not recognised.
func GetUnitDetailsOrPanic(name string) UnitDetails {
	ud, err := GetUnitDetails(name)
	if err != nil {
		panic(err)
	}
	return ud
}

// GetFamilyNames returns a sorted slice holding the allowed names of unit
// types
func GetFamilyNames() []string {
	names := make([]string, 0, len(validUnits))
	for name := range validUnits {
		names = append(names, name)
	}

	sort.Strings(names)
	return names
}
