package units_test

import (
	"fmt"

	"github.com/nickwells/units.mod/v2/units"
)

// ExampleValUnit_Convert demonstrates the use of the ValUnit.Convert
// function. Note that we use the ...OrPanic variants to streamline the
// code. In this example we know that the various units exist and are
// well-formed - it is a programming error if any of the calls fail (such as
// a typo in the unit name).
func ExampleValUnit_Convert() {
	f := units.GetFamilyOrPanic(units.Distance)
	inch := f.GetUnitOrPanic("inch")
	foot := f.GetUnitOrPanic("foot")

	vInInches := units.ValUnit{12.0, inch}

	vInFeet := vInInches.ConvertOrPanic(foot)

	fmt.Printf("%s = %s\n", vInInches, vInFeet)
	// Output: 12 inches = 1 foot
}
