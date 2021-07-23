package units

import (
	"sort"
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestValidUnits(t *testing.T) {
	for fName, f := range unitFamilies {
		if f.altUnits == nil {
			t.Logf("Bad family: %q", fName)
			t.Error("\t: The altUnits map is not initialised\n")
		}

		if f.unitAliases == nil {
			t.Logf("Bad family: %q", fName)
			t.Error("\t: The unitAliases map is not initialised\n")
		}

		if u, ok := f.altUnits[f.baseUnitName]; !ok {
			t.Logf("Bad family: %q", fName)
			t.Logf("\t: Base unit name: %q", f.baseUnitName)
			t.Error("\t: The altUnits map does not contain the base unit\n")
		} else if u.convFactor != 1.0 {
			t.Logf("Bad family: %q", fName)
			t.Logf("\t:    Base unit name: %q", f.baseUnitName)
			t.Logf("\t: Conversion Factor: %g", u.convFactor)
			t.Error("\t: The base unit conversion factor must be: 1.0\n")
		} else if u.convPostAdd != 0.0 {
			t.Logf("Bad family: %q", fName)
			t.Logf("\t:      Base unit name: %q", f.baseUnitName)
			t.Logf("\t: Conversion Post Add: %g", u.convPostAdd)
			t.Error("\t: The base unit post-add value must be: 0.0\n")
		} else if u.convPreAdd != 0.0 {
			t.Logf("Bad family: %q", fName)
			t.Logf("\t:     Base unit name: %q", f.baseUnitName)
			t.Logf("\t: Conversion Pre Add: %g", u.convPreAdd)
			t.Error("\t: The base unit pre-add value must be: 0.0\n")
		}

		for uName, u := range f.altUnits {
			if u.f != f {
				t.Logf("Bad family: %q", fName)
				t.Log("\t: Bad entry in the altUnits map")
				t.Logf("\t: %q has the wrong family\n", uName)
				t.Logf("\t: Expected family: %q\n", f.description)
				t.Logf("\t:   Actual family: %q\n", u.f.description)
				t.Errorf("\t: wrong family")
			}

			if u.convFactor == 0 {
				t.Log("Bad entry in the validUnits map")
				t.Errorf("\t: %q: %q has a zero conversion factor",
					fName, uName)
			}
		}
	}
}

func TestGetFamily(t *testing.T) {
	const badFamilyName = "no-such-family"
	badFamilyNameErr := `There is no unit family called "` + badFamilyName + `"`
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
		_, err := GetFamily(tc.name)
		testhelper.CheckExpErr(t, err, tc)

		panicked, panicVal := testhelper.PanicSafe(func() {
			GetFamilyOrPanic(tc.name)
		})
		testhelper.CheckExpPanicError(t, panicked, panicVal, tc)
	}
}

func TestGetFamilyNames(t *testing.T) {
	fn := GetFamilyNames()
	if len(fn) != len(unitFamilies) {
		t.Logf("Expected name count: %d\n", len(unitFamilies))
		t.Logf("  Actual name count: %d\n", len(fn))
		t.Errorf("GetFamilyNames() failed\n")
	}
}

func TestFamily(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		f                *Family
		expNames         []string
		expFamilyAliases []string
		expBaseUnitName  string
		expDescription   string
		expString        string
	}{
		{
			ID: testhelper.MkID("SampleFamily"),
			f:  SampleFamily,
			expNames: []string{
				SampleUnitBase,
				SampleUnitA,
				SampleUnit2T,
				SampleUnit001,
				SampleUnit123,
				SampleUnitNeg123,
				SampleUnitBad,

				SampleUnitBaseAlias,
				SampleUnitAAlias,
			},
			expFamilyAliases: []string{SampleFamilyAlias},
			expBaseUnitName:  SampleUnitBase,
			expDescription:   SampleFamilyDescription,
			expString: "sample-family: a unit of test." +
				" With 7 units. Base unit: sample." +
				" Aliases: sample-family-alias",
		},
		{
			ID:               testhelper.MkID("BadSampleFamily"),
			f:                BadSampleFamily,
			expFamilyAliases: []string{},
			expBaseUnitName:  SampleUnitBase,
			expDescription:   SampleFamilyDescription,
			expString: "bad-sample-family: a unit of test." +
				" With 0 units. Base unit: sample.",
		},
	}

	for _, tc := range testCases {
		names := tc.f.GetUnitNames()
		names = append(names, tc.f.GetUnitAliases()...)
		sort.Strings(names)
		sort.Strings(tc.expNames)
		testhelper.DiffStringSlice(t, tc.IDStr(), "unit names",
			names, tc.expNames)

		names = tc.f.FamilyAliases()
		sort.Strings(names)
		sort.Strings(tc.expFamilyAliases)
		testhelper.DiffStringSlice(t, tc.IDStr(), "FamilyAliases()",
			names, tc.expFamilyAliases)

		testhelper.DiffString(t, tc.IDStr(), "BaseUnitName()",
			tc.f.BaseUnitName(), tc.expBaseUnitName)

		testhelper.DiffString(t, tc.IDStr(), "Description()",
			tc.f.Description(), tc.expDescription)

		testhelper.DiffString(t, tc.IDStr(), "String()",
			tc.f.String(), tc.expString)
	}
}
