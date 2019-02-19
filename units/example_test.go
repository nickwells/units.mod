package units_test

import (
	"fmt"

	"github.com/nickwells/units.mod/units"
)

// ExampleGetUnit demonstrates the usage of the Unit type
func ExampleGetUnit() {
	var metres = 1

	name := "inch"
	multToInch, err := units.GetUnit(units.Length, name)
	if err != nil {
		fmt.Printf("Couldn't find a mult called '%s': %s\n", name, err)
	}
	fmt.Printf("%d metres == %0.5f %s\n",
		metres, float64(metres)/multToInch.Factor, multToInch.Name)

	// Output: 1 metres == 39.37008 inch
}
