package units_test

import (
	"fmt"

	"github.com/nickwells/units.mod/units"
)

// ExampleGetUnit demonstrates the usage of the GetUnit function
func ExampleGetUnit() {
	var metres = 1

	inch, err := units.GetUnit(units.Length, "inch")
	if err != nil {
		fmt.Printf("Couldn't find a length unit called 'inch': %s\n", err)
	}
	valInInches, err := units.ConvertFromBaseUnits(float64(metres), inch)
	if err != nil {
		fmt.Printf("Couldn't convert to 'inch': %s\n", err)
	}
	fmt.Printf("%d metres == %0.5f %s\n", metres, valInInches, inch.Name)

	// Output: 1 metres == 39.37008 inch
}

// ExampleValWithUnit_Convert demonstrates the use of the ValWithUnit.Convert
// function
func ExampleValWithUnit_Convert() {
	inch, err := units.GetUnit(units.Length, "inch")
	if err != nil {
		fmt.Printf("Couldn't find a length unit called 'inch': %s\n", err)
	}
	foot, err := units.GetUnit(units.Length, "foot")
	if err != nil {
		fmt.Printf("Couldn't find a length unit called 'foot': %s\n", err)
	}
	var vInInches = units.ValWithUnit{
		Val: 12.0,
		U:   inch,
	}
	vInFeet, err := vInInches.Convert(foot)
	if err != nil {
		fmt.Printf("Couldn't convert units to 'foot': %s\n", err)
	}
	fmt.Printf("%s = %s\n", vInInches, vInFeet)
	// Output: 12.00000 inch = 1.00000 foot
}
