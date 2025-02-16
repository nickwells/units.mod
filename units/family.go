package units

import (
	"fmt"
	"strings"
)

// The Family type records details of a family of units such as units of area
// or units of mass. It records the name of the base unit and a
// description. It also records all the available units in the family. The
// base unit is the unit in terms of which any conversion factors are
// defined.
//
// To get a Family value you should use the GetFamily func (or
// GetFamilyOrPanic) passing a Family name chosen ideally from the constant
// values provided.
//
// For testing purposes the SampleFamily and BadStandardFamily values are
// provided.
type Family struct {
	baseUnitName  string
	description   string
	name          string
	altUnits      map[string]Unit
	unitAliases   map[string]string
	familyAliases []string
}

// BaseUnitName returns the name of the base unit for this family.
func (f *Family) BaseUnitName() string {
	return f.baseUnitName
}

// Description returns the family description.
func (f *Family) Description() string {
	return f.description
}

// Name returns the family name (the value to pass to GetFamily to retrieve
// this Family).
func (f *Family) Name() string {
	return f.name
}

// String returns a string representation of the Family object.
func (f *Family) String() string {
	rval := fmt.Sprintf("%s: a %s. With %d units. Base unit: %s.",
		f.name, f.description, len(f.altUnits), f.baseUnitName)
	if len(f.familyAliases) > 0 {
		rval += " Aliases: " + strings.Join(f.familyAliases, ", ")
	}

	return rval
}

// These Family name constants should be used when retrieving unit families
// (using the GetFamily or GetFamilyOrPanic funcs).
const (
	Dimensionless = "dimensionless"
	Time          = "time"
	Data          = "data"
	Distance      = "distance"
	Area          = "area"
	Volume        = "volume"
	Velocity      = "velocity"
	Mass          = "mass"
	Temperature   = "temperature"
	Angle         = "angle"
	Energy        = "energy"
)

var unitFamilies = map[string]*Family{
	numericFamily.name:     numericFamily,
	timeFamily.name:        timeFamily,
	dataFamily.name:        dataFamily,
	distanceFamily.name:    distanceFamily,
	areaFamily.name:        areaFamily,
	volumeFamily.name:      volumeFamily,
	velocityFamily.name:    velocityFamily,
	massFamily.name:        massFamily,
	temperatureFamily.name: temperatureFamily,
	angleFamily.name:       angleFamily,
	energyFamily.name:      energyFamily,
}

// familyAlias is populated from the familyAliases values in the entries in
// unitFamilies.
var familyAlias = map[string]string{}

// populateFamilyAliases sets up the family aliases entries for this Family.
func (f *Family) populateFamilyAliases() {
	for _, a := range f.familyAliases {
		if a == f.name {
			panic(
				fmt.Errorf(
					"The alias %q on Family %q is redundant",
					a, f.name))
		}

		if _, ok := unitFamilies[a]; ok {
			panic(
				fmt.Errorf(
					"The alias %q on Family %q"+
						" is the name of an existing Family",
					a, f.name))
		}

		if fa, ok := familyAlias[a]; ok && fa != f.name {
			panic(
				fmt.Errorf(
					"There is a duplicate alias:"+
						" Family %q has an alias %q and so does Family %q",
					f.name, a, fa))
		}

		familyAlias[a] = f.name
	}
}

// populateUnitAliases sets up the aliases table for the Family from the aliases
// on each Unit.
func (f *Family) populateUnitAliases() {
	if f.unitAliases == nil {
		f.unitAliases = make(map[string]string)
	}

	for uName, u := range f.altUnits {
		for a := range u.aliases {
			if aliasVal, ok := f.unitAliases[a]; ok && aliasVal != uName {
				panic(
					fmt.Errorf(
						"There is a duplicate alias:"+
							" Unit %q has an alias %q and so does Unit %q",
						uName, a, aliasVal))
			}

			f.unitAliases[a] = uName
		}
	}
}

func init() {
	numericFamily.altUnits = dimensionlessNames
	timeFamily.altUnits = timeNames
	dataFamily.altUnits = dataNames
	distanceFamily.altUnits = distanceNames
	areaFamily.altUnits = areaNames
	volumeFamily.altUnits = volumeNames
	velocityFamily.altUnits = velocityNames
	massFamily.altUnits = massNames
	temperatureFamily.altUnits = temperatureNames
	angleFamily.altUnits = angleNames
	energyFamily.altUnits = energyNames

	for _, f := range unitFamilies {
		f.populateUnitAliases()
		f.populateFamilyAliases()
	}
}

// GetFamily returns the named Family. It returns a non-nil error if the
// Family name is not found.
//
// You are advised to use the constant Family names provided.
func GetFamily(name string) (*Family, error) {
	f, ok := unitFamilies[name]
	if !ok {
		alias, ok := familyAlias[name]
		if !ok {
			return nil, fmt.Errorf("There is no unit family called %q", name)
		}

		f, ok = unitFamilies[alias]
		if !ok {
			return nil,
				fmt.Errorf("There is no unit family called %q (alias: %q)",
					name, alias)
		}
	}

	return f, nil
}

// GetFamilyOrPanic returns the named Family. It panics if the Family name is
// not found.
func GetFamilyOrPanic(name string) *Family {
	f, err := GetFamily(name)
	if err != nil {
		panic(err)
	}

	return f
}

// GetFamilyAliases returns a slice holding the names of aliases to unit
// types. Note that the slice is not sorted and the order of elements may
// vary.
func GetFamilyAliases() []string {
	names := make([]string, 0, len(familyAlias))

	for name := range familyAlias {
		names = append(names, name)
	}

	return names
}

// GetFamilyNames returns a slice holding the canonical names (ie not
// aliases) of unit types. Note that the slice is not sorted and the order of
// elements may vary.
func GetFamilyNames() []string {
	names := make([]string, 0, len(unitFamilies))

	for name := range unitFamilies {
		names = append(names, name)
	}

	return names
}

// GetUnit gets the named unit from the UnitDetails. A non-nil error is
// returned if the name is not found.
func (f *Family) GetUnit(name string) (Unit, error) {
	u, ok := f.altUnits[name]
	if ok {
		u.id = name
		return u, nil
	}

	alias := name

	name, ok = f.unitAliases[alias]
	if !ok {
		return u, fmt.Errorf("there is no %s called %q",
			f.description, alias)
	}

	u, ok = f.altUnits[name]
	if !ok {
		return u, fmt.Errorf("there is no %s called %q (%q)",
			f.description, alias, name)
	}

	u.alias = alias
	u.id = name

	return u, nil
}

// GetUnitOrPanic will call GetUnit and check the error returned. If the
// error is non-nil it will panic, otherwise it will return the Unit.
func (f *Family) GetUnitOrPanic(name string) Unit {
	u, err := f.GetUnit(name)
	if err != nil {
		panic(err)
	}

	return u
}

// GetUnitNames returns a slice containing all the valid unit names (but not any
// aliases) in this family. Note that it is not sorted and so will have a
// random order which may vary between calls to this function.
func (f *Family) GetUnitNames() []string {
	names := make([]string, 0, len(f.altUnits))

	for n := range f.altUnits {
		names = append(names, n)
	}

	return names
}

// GetUnitAliases returns a slice containing all the valid unit aliases in
// this family. Note that it is not sorted and so will have a random order
// which may vary between calls to this function.
func (f *Family) GetUnitAliases() []string {
	names := make([]string, 0, len(f.unitAliases))

	for n := range f.unitAliases {
		names = append(names, n)
	}

	return names
}

// FamilyAliases returns all the aliases that can be used to retrieve this
// Family.
func (f *Family) FamilyAliases() []string {
	return f.familyAliases
}
