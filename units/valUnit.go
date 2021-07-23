package units

import (
	"fmt"

	"github.com/nickwells/mathutil.mod/mathutil"
)

// ValUnit associates a value with a unit, it is expected that the value
// is a measure expressed in the given units
type ValUnit struct {
	V float64
	U Unit
}

// String returns a string form of the ValUnit
func (v ValUnit) String() string {
	singularName := v.U.name
	pluralName := v.U.namePlural
	if v.U.alias != "" {
		singularName += " (" + v.U.alias + ")"
		pluralName += " (" + v.U.alias + ")"
	}

	epsilon := 0.0000001

	if mathutil.AlmostEqual(v.V, 1.0, epsilon) {
		return "1 " + singularName
	}
	if mathutil.AlmostEqual(v.V, 0.0, epsilon) {
		return "0 " + pluralName
	}
	return fmt.Sprintf("%.5g %s", v.V, pluralName)
}

// convertFromBaseUnits converts a value expressed in the base Units into
// 'to' Units. It will return a non-nil error if the Units are invalid (have
// a zero conversion factor). See also the ValUnit.Convert method.
func convertFromBaseUnits(v float64, to Unit) (float64, error) {
	if to.convFactor == 0 {
		return v,
			fmt.Errorf("Bad units - a zero conversion factor")
	}
	return ((v + to.convPreAdd) / to.convFactor) + to.convPostAdd,
		nil
}

// convertToBaseUnits converts a value expressed in the 'from' Units into
// base Units. It will return a non-nil error if the Units are invalid (have
// a zero conversion factor). See also the ValUnit.Convert method.
func convertToBaseUnits(v float64, from Unit) (float64, error) {
	if from.convFactor == 0 {
		return v,
			fmt.Errorf("Bad units - a zero conversion factor")
	}
	return ((v - from.convPostAdd) * from.convFactor) - from.convPreAdd,
		nil
}

// Convert will convert the value from its current units to the new units. If
// the new units are not in the same unit family as the existing units or are
// otherwise invalid (have a zero ConvFactor) the value is unchanged and a
// non-nil error is returned.
func (v ValUnit) Convert(u Unit) (ValUnit, error) {
	rval := ValUnit{U: u}
	if v.U.f != u.f {
		return rval,
			fmt.Errorf(
				"Mismatched unit families. Cannot convert units from %s to %s",
				v.U.f.name, u.f.name)
	}

	baseVal, err := convertToBaseUnits(v.V, v.U)
	if err != nil {
		return rval, err
	}
	rval.V, err = convertFromBaseUnits(baseVal, u)

	return rval, err
}

// ConvertOrPanic will call Convert and if the error returned is not nil it
// will panic, otherwise it will return the ValUnit value
func (v ValUnit) ConvertOrPanic(u Unit) ValUnit {
	convertedVal, err := v.Convert(u)
	if err != nil {
		panic(err)
	}
	return convertedVal
}
