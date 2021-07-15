package units

import (
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

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

			if unitVal.ConvFactor == 0 {
				t.Log("Bad entry in the validUnits map")
				t.Errorf("\t: %q: %q has a zero conversion factor",
					famName, unitName)
			}
		}
	}
}

func TestGetAliases(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		fname         string
		uname         string
		expectAliases bool
	}{
		{
			ID:    testhelper.MkID("bad-family"),
			fname: "no such unit family",
		},
		{
			ID:    testhelper.MkID("bad-unitname"),
			fname: "distance",
			uname: "no such unit name",
		},
		{
			ID:            testhelper.MkID("all-good"),
			fname:         "distance",
			uname:         "metre",
			expectAliases: true,
		},
		{
			ID:            testhelper.MkID("all-good-with-family-alias"),
			fname:         "length",
			uname:         "metre",
			expectAliases: true,
		},
	}

	for _, tc := range testCases {
		aliases := GetAliases(tc.fname, tc.uname)
		if len(aliases) == 0 && tc.expectAliases {
			t.Log(tc.IDStr())
			t.Logf("\t: family name: %q\n", tc.fname)
			t.Logf("\t:   unit name: %q\n", tc.uname)
			t.Errorf("\t: Aliases were expected but none were found\n")
		} else if len(aliases) != 0 && !tc.expectAliases {
			t.Log(tc.IDStr())
			t.Logf("\t: family name: %q\n", tc.fname)
			t.Logf("\t:   unit name: %q\n", tc.uname)
			t.Logf("\t:   aliases: %q\n", aliases)
			t.Errorf("\t: Aliases were not expected but were found\n")
		}
	}
}

func TestGetUnitDetails(t *testing.T) {
	const badFamilyName = "no such family name"
	badFamilyNameErr := `no such unit type "` + badFamilyName + `"`
	testCases := []struct {
		testhelper.ID
		name string
		testhelper.ExpPanic
		testhelper.ExpErr
	}{
		{
			ID:       testhelper.MkID("bad-name"),
			ExpPanic: testhelper.MkExpPanic(badFamilyNameErr),
			ExpErr:   testhelper.MkExpErr(badFamilyNameErr),
			name:     badFamilyName,
		},
		{
			ID:   testhelper.MkID("good-name"),
			name: Distance,
		},
	}

	for _, tc := range testCases {
		_, err := GetUnitDetails(tc.name)
		testhelper.CheckExpErr(t, err, tc)

		panicked, panicVal := testhelper.PanicSafe(func() {
			GetUnitDetailsOrPanic(tc.name)
		})
		testhelper.CheckExpPanicError(t, panicked, panicVal, tc)
	}
}

func TestGetFamilyNames(t *testing.T) {
	fn := GetFamilyNames()
	if len(fn) != (len(validUnits) + len(familyAlias)) {
		t.Logf("Expected name count: %d\n", len(validUnits))
		t.Logf("  Actual name count: %d\n", len(fn))
		t.Errorf("GetFamilyNames() failed\n")
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
