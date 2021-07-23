package units

import (
	"fmt"
	"strings"
)

// Unit represents a unit of measure
//
// There are conversion details that allow you to convert to and from the
// base units for the unit family. There are descriptive strings for display
// giving abbreviated and full names of the unit.
//
// The tags provide extra detail about the unit. For instance a unit might be
// tagged as an SI unit or of historical use only.
//
// The aliasName will only be set when the unit has been found through an
// alias rather than the canonical name.
type Unit struct {
	convPreAdd  float64
	convPostAdd float64
	convFactor  float64

	f          *Family
	abbrev     string
	name       string
	namePlural string
	notes      string
	tags       []Tag
	aliases    map[string]string

	alias string
	id    string
}

// ConversionFormula returns a string describing the formula this Unit uses
// to convert a value expressed in base units into these units.
func (u Unit) ConversionFormula() string {
	formula := ""
	if u.convPreAdd != 0 {
		if u.convPreAdd > 0 {
			formula += fmt.Sprintf(" add %g", u.convPreAdd)
		} else {
			formula += fmt.Sprintf(" subtract %g", -u.convPreAdd)
		}
	}
	if u.convFactor != 1.0 {
		formula += fmt.Sprintf(" divide by %g", u.convFactor)
	}
	if u.convPostAdd != 0 {
		if u.convPostAdd > 0 {
			formula += fmt.Sprintf(" add %g", u.convPostAdd)
		} else {
			formula += fmt.Sprintf(" subtract %g", -u.convPostAdd)
		}
	}
	if formula == "" {
		formula = "no conversion needed"
		if u.id == u.f.baseUnitName {
			formula += " (already in the base units)"
		}
	}
	return strings.TrimSpace(formula)
}

// ConvPreAdd returns the conversion pre-add value for this Unit.
func (u Unit) ConvPreAdd() float64 {
	return u.convPreAdd
}

// ConvPostAdd returns the conversion post-add value for this Unit.
func (u Unit) ConvPostAdd() float64 {
	return u.convPostAdd
}

// ConvFactor returns the conversion factor for this Unit.
func (u Unit) ConvFactor() float64 {
	return u.convFactor
}

// Family returns a pointer to the Family that this Unit belongs to
func (u Unit) Family() *Family {
	return u.f
}

// Abbrev returns the abbreviated unit name
func (u Unit) Abbrev() string {
	return u.abbrev
}

// Name returns the name of the unit (in singular form)
func (u Unit) Name() string {
	return u.name
}

// NamePlural returns the name of the unit (in plural form)
func (u Unit) NamePlural() string {
	return u.namePlural
}

// Notes returns the notes about the unit
func (u Unit) Notes() string {
	return u.notes
}

// AliasName returns the name of the alias (if any) used to find this unit
func (u Unit) AliasName() string {
	return u.alias
}

// ID returns the name under which the unit is held in the Family, the
// canonical name.
func (u Unit) ID() string {
	return u.id
}

// HasTag returns true if the unit has the given tag, false otherwise
func (u Unit) HasTag(t Tag) bool {
	for _, ut := range u.tags {
		if ut == t {
			return true
		}
	}
	return false
}

// Tags returns a copy of the tags associated with this unit
func (u Unit) Tags() []Tag {
	rval := []Tag{}
	rval = append(rval, u.tags...)

	return rval
}

// Aliases returns the aliases for this unit. The returned map is keyed on
// the alias name and the associated value is a, possibly empty, description
// of the alias.
func (u Unit) Aliases() map[string]string {
	rval := map[string]string{}
	for k, v := range u.aliases {
		rval[k] = v
	}
	return rval
}
