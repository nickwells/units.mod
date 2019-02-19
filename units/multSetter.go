package units

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/nickwells/param.mod/v2/param"
	"github.com/nickwells/strdist.mod/strdist"
)

// MultCheckFunc is the type of the check function for this setter. It takes
// a Mult parameter and returns an error (or nil)
type MultCheckFunc func(Mult) error

// MultSetter allows you to specify a parameter that can be used to set a
// Mult value. You can also supply a check function that will validate the
// Value.
type MultSetter struct {
	Value  *Mult
	UD     UnitDetails
	Checks []MultCheckFunc
}

// ValueReq returns Mandatory indicating that some value must follow
// the parameter
func (s MultSetter) ValueReq() param.ValueReq {
	return param.Mandatory
}

// Set (called when there is no following value) returns an error
func (s MultSetter) Set(_ string) error {
	return errors.New("no value given (it should be followed by a unit name)")
}

// suggestAltVal will suggest a possible alternative value for the parameter
// value. It will find those strings in the set of possible values that are
// closest to the given value
func (s MultSetter) suggestAltVal(val string) string {
	suggestedNames := ""
	matches :=
		strdist.CaseBlindCosineFinder.FindNStrLike(3, val, s.validNames()...)

	if len(matches) > 0 {
		suggestedNames = " Did you mean: " +
			strings.Join(matches, " or ") +
			"?"
	}
	return suggestedNames
}

// SetWithVal (called when a value follows the parameter) checks that the
// value can be found in the map of Mults, if it cannot it returns an
// error. If there are checks and any check is violated it returns an
// error. Only if the value is parsed successfully and no checks are violated
// is the Value set.
func (s MultSetter) SetWithVal(_ string, paramVal string) error {
	v, ok := s.UD.AltU[paramVal]
	if !ok {

		return fmt.Errorf("'%s' is not a recognised %s.%s",
			paramVal, s.UD.U.Description, s.suggestAltVal(paramVal))
	}

	if len(s.Checks) != 0 {
		for _, check := range s.Checks {
			if check == nil {
				continue
			}

			err := check(v)
			if err != nil {
				return err
			}
		}
	}

	*s.Value = v
	return nil
}

// allowedValNames returns slice containing the names of allowed values
func (s MultSetter) validNames() []string {
	var names []string

	for k := range s.UD.AltU {
		names = append(names, k)
	}

	return names
}

// AllowedValues returns a string describing the allowed values
func (s MultSetter) AllowedValues() string {
	var names = s.validNames()
	if len(names) == 0 {
		return "there are no valid conversions for this unit type: " +
			s.UD.U.BaseUnitName
	}

	sort.Strings(names)
	rval := ""
	sep := ""
	for _, name := range names {
		rval += sep + name
		sep = ", "
	}

	if len(s.Checks) != 0 {
		rval += " subject to checks"
	}
	return rval
}

// CurrentValue returns the current setting of the parameter value
func (s MultSetter) CurrentValue() string {
	return s.Value.AltUnit
}

// CheckSetter panics if the setter has not been properly created - if the
// Value is nil, if the base unit is invalid or if one of the check functions
// is nil.
func (s MultSetter) CheckSetter(name string) {
	if s.Value == nil {
		panic(name +
			": MultSetter Check failed: the Value to be set is nil")
	}
	if s.UD.AltU == nil {
		panic(name +
			": MultSetter Check failed: there are no valid alternative units")
	}
	for _, check := range s.Checks {
		if check == nil {
			panic(name +
				": MultSetter Check failed: one of the check functions is nil")
		}
	}
}
