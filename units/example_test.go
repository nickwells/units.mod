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
// function. Note that we use the ...OrPanic variants to streamline the
// code. In this example we know that the various units exist and are
// well-formed - it is a programming error if any of the calls fail (such as
// a typo in the unit name).
func ExampleValWithUnit_Convert() {
	inch := units.GetUnitOrPanic(units.Length, "inch")
	foot := units.GetUnitOrPanic(units.Length, "foot")

	var vInInches = units.ValWithUnit{12.0, inch}

	vInFeet := vInInches.ConvertOrPanic(foot)

	fmt.Printf("%s = %s\n", vInInches, vInFeet)
	// Output: 12 inches = 1 foot
}
