package units

import (
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestGetUnit(t *testing.T) {
	const (
		goodFamilyName    = Distance
		goodUnitName      = "metre"
		goodUnitNameAlias = "meter"
		badFamilyName     = "no such family name"
		badUnitName       = "no such unit name"
	)

	ud, err := GetUnitDetails(goodFamilyName)
	if err != nil {
		t.Fatalf("Cannot get the unit details for %q: %v", goodFamilyName, err)
	}

	badFamilyNameErr := `no such unit type "` + badFamilyName + `"`
	badUnitNameErr := `there is no ` + ud.Fam.Description + ` with a name of "` + badUnitName + `"`
	testCases := []struct {
		testhelper.ID
		fName string
		uName string
		testhelper.ExpErr
		testhelper.ExpPanic
		expAlias string
	}{
		{
			ID:       testhelper.MkID("bad-family"),
			fName:    badFamilyName,
			ExpErr:   testhelper.MkExpErr(badFamilyNameErr),
			ExpPanic: testhelper.MkExpPanic(badFamilyNameErr),
		},
		{
			ID:       testhelper.MkID("good-family-bad-unit"),
			fName:    goodFamilyName,
			uName:    badUnitName,
			ExpErr:   testhelper.MkExpErr(badUnitNameErr),
			ExpPanic: testhelper.MkExpPanic(badUnitNameErr),
		},
		{
			ID:    testhelper.MkID("good-family-good-unit"),
			fName: goodFamilyName,
			uName: goodUnitName,
		},
		{
			ID:       testhelper.MkID("good-family-good-unit-alias"),
			fName:    goodFamilyName,
			uName:    goodUnitNameAlias,
			expAlias: goodUnitNameAlias,
		},
	}

	for _, tc := range testCases {
		u, err := GetUnit(tc.fName, tc.uName)
		testhelper.CheckExpErr(t, err, tc)
		testhelper.DiffString(t, tc.IDStr(), "", u.aliasName, tc.expAlias)

		panicked, panicVal := testhelper.PanicSafe(func() {
			GetUnitOrPanic(tc.fName, tc.uName)
		})
		testhelper.CheckExpPanicError(t, panicked, panicVal, tc)
	}
}
