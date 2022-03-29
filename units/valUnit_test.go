package units

import (
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestValUnit_String(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vwu       ValUnit
		expString string
	}{
		{
			ID: testhelper.MkID("zero val, no alias"),
			vwu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expString: "0 samples",
		},
		{
			ID: testhelper.MkID("almost zero val, no alias"),
			vwu: ValUnit{
				V: 0.00000001,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expString: "0 samples",
		},
		{
			ID: testhelper.MkID("unit val, no alias"),
			vwu: ValUnit{
				V: 1,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expString: "1 sample",
		},
		{
			ID: testhelper.MkID("almost unit val (too big), no alias"),
			vwu: ValUnit{
				V: 1.00000001,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expString: "1 sample",
		},
		{
			ID: testhelper.MkID("almost unit val (too small), no alias"),
			vwu: ValUnit{
				V: 0.9999999,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expString: "1 sample",
		},
		{
			ID: testhelper.MkID("multi val, no alias"),
			vwu: ValUnit{
				V: 42,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expString: "42 samples",
		},
		{
			ID: testhelper.MkID("multi val, with alias"),
			vwu: ValUnit{
				V: 42,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBaseAlias),
			},
			expString: "42 samples (smpl)",
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
			u:  SampleFamily.GetUnitOrPanic(SampleUnit001),
		},
		{
			ID: testhelper.MkID("full conversion"),
			v:  41,
			u:  SampleFamily.GetUnitOrPanic(SampleUnit123),
		},
		{
			ID:     testhelper.MkID("bad unit"),
			v:      41,
			u:      SampleFamily.GetUnitOrPanic(SampleUnitBad),
			ExpErr: testhelper.MkExpErr("Bad units - a zero conversion factor"),
		},
	}

	for _, tc := range testCases {
		v, err := convertToBaseUnits(tc.v, tc.u)
		testhelper.CheckExpErr(t, err, tc)

		v, err = convertFromBaseUnits(v, tc.u)
		testhelper.CheckExpErr(t, err, tc)

		testhelper.DiffFloat(t, tc.IDStr(), "", v, tc.v, 0.0000001)
	}
}

func TestValUnit_Convert(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vwu    ValUnit
		toUnit Unit
		expVal float64
		testhelper.ExpErr
	}{
		{
			ID: testhelper.MkID("convert-001-123"),
			vwu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnit001),
			},
			toUnit: SampleFamily.GetUnitOrPanic(SampleUnit123),
			expVal: 2.0 + 1.0/3.0,
		},
		{
			ID: testhelper.MkID("convert-bad-from-unit"),
			vwu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBad),
			},
			toUnit: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			expVal: 0,
			ExpErr: testhelper.MkExpErr("Bad units - a zero conversion factor"),
		},
		{
			ID: testhelper.MkID("convert-bad-to-unit"),
			vwu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			toUnit: SampleFamily.GetUnitOrPanic(SampleUnitBad),
			expVal: 0,
			ExpErr: testhelper.MkExpErr("Bad units - a zero conversion factor"),
		},
	}

	for _, tc := range testCases {
		v, err := tc.vwu.Convert(tc.toUnit)
		testhelper.CheckExpErr(t, err, tc)
		testhelper.DiffFloat(t, tc.IDStr(), "", v.V, tc.expVal, 0.0000001)
	}
}
