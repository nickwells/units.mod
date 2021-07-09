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
