package units

import "testing"

func TestValidUnits(t *testing.T) {
	for famName, ud := range validUnits {
		for unitName, unitVal := range ud.AltU {
			if unitVal.Fam != ud.Fam {
				t.Log("Bad entry in the validUnits map")
				t.Logf("\t: %q: %q has the wrong family\n", famName, unitName)
				t.Logf("\t: Expected family: %q\n", ud.Fam.Description)
				t.Logf("\t:   Actual family: %q\n", unitVal.Fam.Description)
				t.Errorf("\t: wrong family")
			}
		}
	}
}

func TestAliasValues(t *testing.T) {
	for f, ud := range validUnits {
		reverseAlias[f] = make(map[string][]string)
		for aliasName, alias := range ud.Aliases {
			if _, ok := ud.AltU[aliasName]; ok {
				t.Errorf(
					"unit family %q: alias %q"+
						" already exists in the map of unit names",
					f, aliasName)
			}
			if _, ok := ud.AltU[alias.UnitName]; !ok {
				t.Errorf(
					"unit family %q: alias %q references unit %q"+
						" which does not exist in the map of unit names",
					f, aliasName, alias)
			}
		}
	}
}
