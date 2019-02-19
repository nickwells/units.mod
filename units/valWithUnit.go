package units

import "fmt"

// ValWithUnit associates a value with a unit, it is expected that the value
// is a measure expressed in the given units
type ValWithUnit struct {
	Val float64
	U   Unit
}

// String returns a string form of the ValWithUnit
func (v ValWithUnit) String() string {
	return fmt.Sprintf("%.5f %s", v.Val, v.U.Name)
}

// ConvertFromBaseUnits converts a value expressed in the base Units into
// 'to' Units. It will return a non-nil error if the Units are invalid (have
// a zero ConvFactor)
func ConvertFromBaseUnits(v float64, to Unit) (float64, error) {
	if to.ConvFactor == 0 {
		return v,
			fmt.Errorf("Bad units - a zero conversion factor")
	}
	return ((v - to.ConvPostAdd) / to.ConvFactor) + to.ConvPreAdd,
		nil
}

// ConvertToBaseUnits converts a value expressed in the 'from' Units into
// base Units. It will return a non-nil error if the Units are invalid (have
// a zero ConvFactor)
func ConvertToBaseUnits(v float64, from Unit) (float64, error) {
	if from.ConvFactor == 0 {
		return v,
			fmt.Errorf("Bad units - a zero conversion factor")
	}
	return ((v + from.ConvPreAdd) * from.ConvFactor) + from.ConvPostAdd,
		nil
}

// Convert will convert the value from its current units to the new units. If
// the new units are not in the same unit family as the existing units or are
// otherwise invalid (have a zero ConvFactor) the value is unchanged and a
// non-nil error is returned.
func (v ValWithUnit) Convert(u Unit) (ValWithUnit, error) {
	if v.U.Fam != u.Fam {
		return v,
			fmt.Errorf(
				"Mismatched unit families. Cannot convert units from %s to %s",
				v.U.Fam, u.Fam)
	}

	newVal, err := ConvertToBaseUnits(v.Val, v.U)
	if err != nil {
		return v, err
	}
	newVal, err = ConvertFromBaseUnits(newVal, u)
	if err != nil {
		return v, err
	}
	return ValWithUnit{
		Val: newVal,
		U:   u,
	}, nil
}
