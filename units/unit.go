package units

// Unit is the type of a unit. Only certain units are valid and you can check
// this using the IsValid func. The constants define the only valid units and
// should be used wherever possible
type Unit struct {
	Name string
	Type string
}

var (
	// NoUnit represents a typeless unit
	NoUnit = Unit{
		Name: "",
		Type: "dimensionless value",
	}
	// Second represents the base unit of time
	Second = Unit{
		Name: "second",
		Type: "unit of time",
	}
	// Byte represents the base unit of computer memory
	Byte = Unit{
		Name: "byte",
		Type: "unit of data",
	}
	// Metre represents the base unit of distance
	Metre = Unit{
		Name: "metre",
		Type: "unit of distance",
	}
	// SquareMetre represents the base unit of area
	SquareMetre = Unit{
		Name: "square metre",
		Type: "unit of area",
	}
	// CubicMetre represents the base unit of volume
	CubicMetre = Unit{
		Name: "cubic metre",
		Type: "unit of volume",
	}
	// Gram represents the base unit of mass
	Gram = Unit{
		Name: "gram",
		Type: "unit of mass",
	}
)

var validUnits = map[Unit]bool{
	NoUnit:      true,
	Second:      true,
	Byte:        true,
	Metre:       true,
	SquareMetre: true,
	CubicMetre:  true,
	Gram:        true,
}

// IsValid will test the validity of the supplied unit
func IsValid(u Unit) bool {
	return validUnits[u]
}
