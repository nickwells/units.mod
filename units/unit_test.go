package units

import (
	"sort"
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestGetUnit(t *testing.T) {
	const (
		goodFamilyName    = Distance
		goodUnitName      = "metre"
		goodUnitNameAlias = "meter"
		badFamilyName     = "no-such-family"
		badUnitName       = "no-such-unit"
	)

	f, err := GetFamily(goodFamilyName)
	if err != nil {
		t.Fatalf("Cannot get the unit details for %q: %v", goodFamilyName, err)
	}

	badFamilyNameErr := `There is no unit family called "` + badFamilyName + `"`
	badUnitNameErr := `there is no ` + f.description +
		` called "` + badUnitName + `"`
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
		var err error
		var f *Family
		var u Unit
		f, err = GetFamily(tc.fName)
		if err == nil {
			u, err = f.GetUnit(tc.uName)
		}
		testhelper.CheckExpErr(t, err, tc)
		testhelper.DiffString(t, tc.IDStr(), "", u.alias, tc.expAlias)

		panicked, panicVal := testhelper.PanicSafe(func() {
			f = GetFamilyOrPanic(tc.fName)
			u = f.GetUnitOrPanic(tc.uName)
		})
		testhelper.CheckExpPanicError(t, panicked, panicVal, tc)
	}
}

func TestUnitInternal(t *testing.T) {
	baseUnit := SampleFamily.GetUnitOrPanic(SampleFamily.baseUnitName)
	const epsilon = 0.000001

	testCases := []struct {
		testhelper.ID
		uName         string
		expPreAdd     float64
		expPostAdd    float64
		expFactor     float64
		expFormula    string
		testVal       float64
		expConvVal    float64
		expFamily     *Family
		expAbbrev     string
		expName       string
		expNamePlural string
		expNotes      string
		expAliasName  string
		expAliases    []string
		expID         string
		expTags       []string
		testTag       Tag
		expHasTag     bool
	}{
		{
			ID:            testhelper.MkID("sample unit: s123"),
			uName:         SampleUnit123,
			expPreAdd:     1,
			expPostAdd:    2,
			expFactor:     3,
			expFormula:    "add 1 divide by 3 add 2",
			testVal:       2.0,
			expConvVal:    3.0,
			expFamily:     SampleFamily,
			expAbbrev:     "s123",
			expName:       "sample 1, 2, 3",
			expNamePlural: "samples 1, 2, 3",
			expNotes:      "notes about unit s123",
			expAliasName:  "",
			expAliases:    []string{},
			expID:         SampleUnit123,
			expTags:       []string{string(SampleTagName)},
			testTag:       SampleTagName,
			expHasTag:     true,
		},
		{
			ID:            testhelper.MkID("sample unit: s-1-2+3"),
			uName:         SampleUnitNeg123,
			expPreAdd:     -1,
			expPostAdd:    -2,
			expFactor:     3,
			expFormula:    "subtract 1 divide by 3 subtract 2",
			testVal:       4.0,
			expConvVal:    -1.0,
			expFamily:     SampleFamily,
			expAbbrev:     "s-1-2+3",
			expName:       "sample -1, -2, 3",
			expNamePlural: "samples -1, -2, 3",
			expNotes:      "notes about unit s-1-2+3",
			expAliasName:  "",
			expAliases:    []string{},
			expID:         SampleUnitNeg123,
			expTags:       []string{string(SampleTagName)},
			testTag:       SampleTagName,
			expHasTag:     true,
		},
		{
			ID:            testhelper.MkID("sample unit: sa (with ID)"),
			uName:         SampleUnitA,
			expPreAdd:     0,
			expPostAdd:    0,
			expFactor:     2,
			expFormula:    "divide by 2",
			testVal:       2.0,
			expConvVal:    1.0,
			expFamily:     SampleFamily,
			expAbbrev:     "sa",
			expName:       "sample A",
			expNamePlural: "samples A",
			expNotes:      "notes about unit sa, which has an alias",
			expAliasName:  "",
			expAliases:    []string{SampleUnitAAlias},
			expID:         SampleUnitA,
			expTags:       []string{string(SampleTagName)},
			testTag:       SampleTagName,
			expHasTag:     true,
		},
		{
			ID:            testhelper.MkID("sample unit: sa (with alias)"),
			uName:         SampleUnitAAlias,
			expPreAdd:     0,
			expPostAdd:    0,
			expFactor:     2,
			expFormula:    "divide by 2",
			testVal:       2.0,
			expConvVal:    1.0,
			expFamily:     SampleFamily,
			expAbbrev:     "sa",
			expName:       "sample A",
			expNamePlural: "samples A",
			expNotes:      "notes about unit sa, which has an alias",
			expAliasName:  SampleUnitAAlias,
			expAliases:    []string{SampleUnitAAlias},
			expID:         SampleUnitA,
			expTags:       []string{string(SampleTagName)},
			testTag:       SampleTagNameOther,
			expHasTag:     false,
		},
		{
			ID:            testhelper.MkID("sample unit: baseUnit"),
			uName:         SampleUnitBase,
			expPreAdd:     0,
			expPostAdd:    0,
			expFactor:     1,
			expFormula:    "no conversion needed (already in the base units)",
			testVal:       2.0,
			expConvVal:    2.0,
			expFamily:     SampleFamily,
			expAbbrev:     "s",
			expName:       "sample",
			expNamePlural: "samples",
			expNotes:      "some brief notes about the base unit",
			expAliasName:  "",
			expAliases:    []string{SampleUnitBaseAlias},
			expID:         SampleUnitBase,
			expTags:       []string{string(SampleTagName)},
			testTag:       SampleTagNameOther,
			expHasTag:     false,
		},
	}

	for _, tc := range testCases {
		u, err := SampleFamily.GetUnit(tc.uName)
		if err != nil {
			t.Log(tc.IDStr())
			t.Errorf("\t: Couldn't get the unit: %v\n", err)
			continue
		}

		testhelper.DiffFloat64(t, tc.IDStr(), "ConvPreAdd",
			u.ConvPreAdd(), tc.expPreAdd, epsilon)
		testhelper.DiffFloat64(t, tc.IDStr(), "ConvPostAdd",
			u.ConvPostAdd(), tc.expPostAdd, epsilon)
		testhelper.DiffFloat64(t, tc.IDStr(), "ConvFactor",
			u.ConvFactor(), tc.expFactor, epsilon)
		testhelper.DiffString(t, tc.IDStr(), "ConversionFormula",
			u.ConversionFormula(), tc.expFormula)
		convResult := ValUnit{V: tc.testVal, U: baseUnit}.ConvertOrPanic(u)
		testhelper.DiffFloat64(t, tc.IDStr(), "Conversion",
			convResult.V, tc.expConvVal, epsilon)
		testhelper.DiffString(t, tc.IDStr(), "Family",
			u.Family().Name(), tc.expFamily.name)
		testhelper.DiffString(t, tc.IDStr(), "Abbrev",
			u.Abbrev(), tc.expAbbrev)
		testhelper.DiffString(t, tc.IDStr(), "Name",
			u.Name(), tc.expName)
		testhelper.DiffString(t, tc.IDStr(), "NamePlural",
			u.NamePlural(), tc.expNamePlural)
		testhelper.DiffString(t, tc.IDStr(), "Notes",
			u.Notes(), tc.expNotes)
		testhelper.DiffString(t, tc.IDStr(), "ID",
			u.ID(), tc.expID)
		aliasMap := u.Aliases()
		aliases := []string{}
		for a := range aliasMap {
			aliases = append(aliases, a)
		}
		sort.Strings(aliases)
		testhelper.DiffStringSlice(t, tc.IDStr(), "Aliases",
			aliases, tc.expAliases)
		testhelper.DiffString(t, tc.IDStr(), "AliasName",
			u.AliasName(), tc.expAliasName)
		tags := u.Tags()
		strTags := []string{}
		for _, t := range tags {
			strTags = append(strTags, string(t))
		}
		testhelper.DiffStringSlice(t, tc.IDStr(), "Tags", strTags, tc.expTags)
		testhelper.DiffBool(t, tc.IDStr(), "HasTag",
			u.HasTag(tc.testTag), tc.expHasTag)
	}
}
