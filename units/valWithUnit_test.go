package units

import (
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

var (
	dfltFamily = Family{
		"base unit name",
		"unit family description",
		"family", "",
	}

	unit0_0_1 = Unit{
		0, 0, 1,
		dfltFamily,
		"u", "unit", "units", "unit notes",
		"",
	}
	aliasUnit = Unit{
		0, 0, 1,
		dfltFamily,
		"u", "unit", "units", "unit notes",
		"alias",
	}
	unit1_2_3 = Unit{
		1, 2, 3,
		dfltFamily,
		"u", "unit", "units", "unit notes",
		"",
	}
	badUnit = Unit{
		0, 0, 0,
		dfltFamily,
		"u", "unit", "units", "unit notes",
		"",
	}
)

func TestValWithUnit_String(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vwu       ValWithUnit
		expString string
	}{
		{
			ID: testhelper.MkID("zero val, no alias"),
			vwu: ValWithUnit{
				Val: 0,
				U:   unit0_0_1,
			},
			expString: "0 units",
		},
		{
			ID: testhelper.MkID("almost zero val, no alias"),
			vwu: ValWithUnit{
				Val: 0.00000001,
				U:   unit0_0_1,
			},
			expString: "0 units",
		},
		{
			ID: testhelper.MkID("unit val, no alias"),
			vwu: ValWithUnit{
				Val: 1,
				U:   unit0_0_1,
			},
			expString: "1 unit",
		},
		{
			ID: testhelper.MkID("almost unit val (too big), no alias"),
			vwu: ValWithUnit{
				Val: 1.00000001,
				U:   unit0_0_1,
			},
			expString: "1 unit",
		},
		{
			ID: testhelper.MkID("almost unit val (too small), no alias"),
			vwu: ValWithUnit{
				Val: 0.9999999,
				U:   unit0_0_1,
			},
			expString: "1 unit",
		},
		{
			ID: testhelper.MkID("multi val, no alias"),
			vwu: ValWithUnit{
				Val: 42,
				U:   unit0_0_1,
			},
			expString: "42 units",
		},
		{
			ID: testhelper.MkID("multi val, with alias"),
			vwu: ValWithUnit{
				Val: 42,
				U:   aliasUnit,
			},
			expString: "42 units (alias)",
		},
	}

	for _, tc := range testCases {
		s := tc.vwu.String()
		testhelper.DiffString(t, tc.IDStr(), "", s, tc.expString)
	}
}

func TestConvertToAndFromBaseUnits(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		v float64
		u Unit
		testhelper.ExpErr
	}{
		{
			ID: testhelper.MkID("no change conversion"),
			v:  42,
			u:  unit0_0_1,
		},
		{
			ID: testhelper.MkID("full conversion"),
			v:  41,
			u:  unit1_2_3,
		},
		{
			ID:     testhelper.MkID("bad unit"),
			v:      41,
			u:      badUnit,
			ExpErr: testhelper.MkExpErr("Bad units - a zero conversion factor"),
		},
	}

	for _, tc := range testCases {
		v, err := ConvertToBaseUnits(tc.v, tc.u)
		testhelper.CheckExpErr(t, err, tc)

		v, err = ConvertFromBaseUnits(v, tc.u)
		testhelper.CheckExpErr(t, err, tc)

		testhelper.DiffFloat64(t, tc.IDStr(), "", v, tc.v, 0.0000001)
	}
}

func TestValWithUnit_Convert(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vwu    ValWithUnit
		toUnit Unit
		expVal float64
		testhelper.ExpErr
	}{
		{
			ID: testhelper.MkID("..."),
			vwu: ValWithUnit{
				Val: 0,
				U:   unit0_0_1,
			},
			toUnit: unit1_2_3,
			expVal: 2.0 + 1.0/3.0,
		},
		{
			ID: testhelper.MkID("..."),
			vwu: ValWithUnit{
				Val: 0,
				U:   badUnit,
			},
			toUnit: unit1_2_3,
			expVal: 0,
			ExpErr: testhelper.MkExpErr("Bad units - a zero conversion factor"),
		},
		{
			ID: testhelper.MkID("..."),
			vwu: ValWithUnit{
				Val: 0,
				U:   unit0_0_1,
			},
			toUnit: badUnit,
			expVal: 0,
			ExpErr: testhelper.MkExpErr("Bad units - a zero conversion factor"),
		},
	}

	for _, tc := range testCases {
		v, err := tc.vwu.Convert(tc.toUnit)
		testhelper.CheckExpErr(t, err, tc)
		testhelper.DiffFloat64(t, tc.IDStr(), "", v.Val, tc.expVal, 0.0000001)
	}
}
