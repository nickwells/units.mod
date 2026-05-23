package units

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nickwells/mathutil.mod/v2/mathutil"
)

// ValUnit associates a value with a unit, it is expected that the value
// is a measure expressed in the given units
type ValUnit struct {
	V float64
	U Unit
}

// unitNames returns the pair of names (singular and plural) for the unit.
func (v ValUnit) unitNames() (string, string) {
	singularName := v.U.name
	pluralName := v.U.namePlural
	alias := v.U.alias
	singularNameHyphenated := strings.ReplaceAll(singularName, " ", "-")
	pluralNameHyphenated := strings.ReplaceAll(pluralName, " ", "-")

	if alias != "" &&
		alias != singularName && alias != pluralName &&
		alias != singularNameHyphenated && alias != pluralNameHyphenated {
		singularName += " (" + alias + ")"
		pluralName += " (" + alias + ")"
	}

	return singularName, pluralName
}

// effectiveVal returns the value or, if the value is within (10^-prec)/2 of
// zero or one then those values are returned.
func (v ValUnit) effectiveVal(prec int) float64 {
	epsilon := math.Pow10(-prec) / 2 //nolint:mnd

	if mathutil.AlmostEqual(v.V, 1.0, epsilon) {
		return 1.0
	}

	if mathutil.AlmostEqual(v.V, 0.0, epsilon) {
		return 0.0
	}

	return v.V
}

// String returns a string form of the ValUnit
func (v ValUnit) String() string {
	singularName, name := v.unitNames()

	const prec = 5

	eVal := v.effectiveVal(prec)
	if eVal == 1.0 {
		name = singularName
	}

	return fmt.Sprintf("%.5g %s", eVal, name)
}

// makeFmtStart translates the fmt.State into the start of a format
// string. This can then have the verb character added at the end to generate
// the appropriate formatting string. Two such strings are returned, one for
// strings and one for non-strings. This is to avoid formats of strings with
// zero maximum length being generated when the intention was only to limit
// or suppress decimal precision.
func makeFmtStart(f fmt.State) (string, string) {
	fmtStart := "%"

	if f.Flag('+') {
		fmtStart += "+"
	}

	if f.Flag(' ') {
		fmtStart += " "
	}

	if f.Flag('0') {
		fmtStart += "0"
	}

	if f.Flag('#') {
		fmtStart += "#"
	}

	if wid, ok := f.Width(); ok {
		fmtStart += strconv.Itoa(wid)
	}

	fmtStrStart := fmtStart
	if prec, ok := f.Precision(); ok {
		fmtStart += "." + strconv.Itoa(prec)
	}

	return fmtStrStart, fmtStart
}

// uFormat performs the formatting for the 'u' verb
func (v ValUnit) uFormat(f fmt.State) {
	_, fmtNonStrStart := makeFmtStart(f)
	numFmt := fmtNonStrStart + "f"
	singularName, name := v.unitNames()
	nameWidth := max(len(singularName), len(name))
	nameFmt := "%" +
		strconv.Itoa(nameWidth) +
		"." + strconv.Itoa(nameWidth) +
		"s"

	prec, _ := f.Precision()

	eVal := v.effectiveVal(prec)
	if eVal == 1.0 {
		name = singularName
	}

	fmt.Fprintf(f, numFmt+" "+nameFmt, eVal, name)
}

// vFormat performs the formatting for the 'v' verb
func (v ValUnit) vFormat(f fmt.State) {
	_, fmtNonStrStart := makeFmtStart(f)
	numFmt := fmtNonStrStart + "f"
	valFmt := fmtNonStrStart + "v"

	if f.Flag('#') {
		fmt.Fprint(f, "units.ValUnit")
	}

	fmt.Fprint(f, "{")

	if f.Flag('#') || f.Flag('+') {
		fmt.Fprint(f, "V:")
	}

	fmt.Fprintf(f, numFmt, v.V)

	if f.Flag('#') {
		fmt.Fprint(f, ", ")
	} else {
		fmt.Fprint(f, " ")
	}

	if f.Flag('#') || f.Flag('+') {
		fmt.Fprint(f, "U:")
	}

	fmt.Fprintf(f, valFmt, v.U)

	fmt.Fprint(f, "}")
}

// Format provides a custom formatter for a ValUnit
func (v ValUnit) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmtStrStart, _ := makeFmtStart(f)
		fmt.Fprintf(f, fmtStrStart+"s", v.String())
	case 'q':
		fmtStrStart, _ := makeFmtStart(f)
		fmt.Fprintf(f, fmtStrStart+"q", v.String())
	case 'u':
		v.uFormat(f)
	case 'v':
		v.vFormat(f)
	default:
		_, _ = f.Write([]byte("%!" +
			string(verb) +
			"(units.ValUnit=" + v.String() + ")"))
	}
}

// convertFromBaseUnits converts a value expressed in the base Units into
// 'to' Units. It will return a non-nil error if the Units are invalid (have
// a zero conversion factor). See also the ValUnit.Convert method.
func convertFromBaseUnits(v float64, to Unit) (float64, error) {
	if to.convFactor == 0 {
		return v, fmt.Errorf("bad units - a zero conversion factor")
	}

	return ((v + to.convPreAdd) / to.convFactor) + to.convPostAdd, nil
}

// convertToBaseUnits converts a value expressed in the 'from' Units into
// base Units. It will return a non-nil error if the Units are invalid (have
// a zero conversion factor). See also the ValUnit.Convert method.
func convertToBaseUnits(v float64, from Unit) (float64, error) {
	if from.convFactor == 0 {
		return v, fmt.Errorf("bad units - a zero conversion factor")
	}

	return ((v - from.convPostAdd) * from.convFactor) - from.convPreAdd, nil
}

// Convert will convert the value from its current units to the new units. If
// the new units are not in the same unit family as the existing units or are
// otherwise invalid (have a zero ConvFactor) a non-nil error is returned.
func (v ValUnit) Convert(u Unit) (ValUnit, error) {
	rval := ValUnit{U: u}
	if v.U.f != u.f {
		return rval,
			fmt.Errorf(
				"mismatched unit families. Cannot convert units from %s to %s",
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
