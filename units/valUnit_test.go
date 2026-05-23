package units

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestValUnit_String(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vu     ValUnit
		expStr string
	}{
		{
			ID: testhelper.MkID("zero val, no alias"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expStr: "0 samples",
		},
		{
			ID: testhelper.MkID("almost zero val, no alias"),
			vu: ValUnit{
				V: 0.00000001,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expStr: "0 samples",
		},
		{
			ID: testhelper.MkID("unit val, no alias"),
			vu: ValUnit{
				V: 1,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expStr: "1 sample",
		},
		{
			ID: testhelper.MkID("almost unit val (too big), no alias"),
			vu: ValUnit{
				V: 1.00000001,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expStr: "1 sample",
		},
		{
			ID: testhelper.MkID("almost unit val (too small), no alias"),
			vu: ValUnit{
				V: 0.9999999,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expStr: "1 sample",
		},
		{
			ID: testhelper.MkID("multi val, no alias"),
			vu: ValUnit{
				V: 42,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			expStr: "42 samples",
		},
		{
			ID: testhelper.MkID("multi val, with alias"),
			vu: ValUnit{
				V: 42,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBaseAlias),
			},
			expStr: "42 samples (smpl)",
		},
	}

	for _, tc := range testCases {
		s := tc.vu.String()
		testhelper.DiffString(t, tc.IDStr(), "", s, tc.expStr)
	}
}

// hexEdit replaces hex values in the string with a constant value. The
// intention is to allow pointer values represented in a string to be
// compared when we don't care about the actual value.
func hexEdit(s string) string {
	re := regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)

	return re.ReplaceAllString(s, "<HEX>")
}

func TestValUnit_Format(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vu     ValUnit
		format string
		expStr string
	}{
		{
			ID: testhelper.MkID("zero-val-s"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%s",
			expStr: "0 samples",
		},
		{
			ID: testhelper.MkID("zero-val-q"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%q",
			expStr: `"0 samples"`,
		},
		{
			ID: testhelper.MkID("zero-val-T"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%T",
			expStr: `units.ValUnit`,
		},
		{
			ID: testhelper.MkID("zero-val-u"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%u",
			expStr: `0.000000 samples`,
		},
		{
			ID: testhelper.MkID("one-val-u"),
			vu: ValUnit{
				V: 1,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%u",
			expStr: `1.000000  sample`,
		},
		{
			ID: testhelper.MkID("zero-val-u.0"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%.0u",
			expStr: `0 samples`,
		},
		{
			ID: testhelper.MkID("zero-val-u.3"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%.3u",
			expStr: `0.000 samples`,
		},
		{
			ID: testhelper.MkID("zero-val-v"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%v",
			expStr: `{0.000000 {0 0 1 <HEX> s sample samples some brief notes about the base unit [SampleTag] map[smpl:Alias]  sample}}`,
		},
		{
			ID: testhelper.MkID("zero-val-+v"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%+v",
			expStr: `{V:+0.000000 U:{convPreAdd:0 convPostAdd:0 convFactor:1 f:<HEX> abbrev:s name:sample namePlural:samples notes:some brief notes about the base unit tags:[SampleTag] aliases:map[smpl:Alias] alias: id:sample}}`,
		},
		{
			ID: testhelper.MkID("zero-val-#v"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%#v",
			expStr: `units.ValUnit{V:0.000000, U:units.Unit{convPreAdd:0, convPostAdd:0, convFactor:1, f:(*units.Family)(<HEX>), abbrev:"s", name:"sample", namePlural:"samples", notes:"some brief notes about the base unit", tags:[]units.Tag{"SampleTag"}, aliases:map[string]string{"smpl":"Alias"}, alias:"", id:"sample"}}`,
		},
		{
			ID: testhelper.MkID("zero-val-d"),
			vu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			format: "%d",
			expStr: `%!d(ValUnit=0 samples)`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			s := hexEdit(fmt.Sprintf(tc.format, tc.vu))

			testhelper.DiffString(t, tc.IDStr(), tc.format, s, tc.expStr)
		})
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
			ExpErr: testhelper.MkExpErr("bad units - a zero conversion factor"),
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
			ExpErr: testhelper.MkExpErr("bad units - a zero conversion factor"),
		},
		{
			ID: testhelper.MkID("convert-bad-to-unit"),
			vwu: ValUnit{
				V: 0,
				U: SampleFamily.GetUnitOrPanic(SampleUnitBase),
			},
			toUnit: SampleFamily.GetUnitOrPanic(SampleUnitBad),
			expVal: 0,
			ExpErr: testhelper.MkExpErr("bad units - a zero conversion factor"),
		},
	}

	for _, tc := range testCases {
		v, err := tc.vwu.Convert(tc.toUnit)
		testhelper.CheckExpErr(t, err, tc)
		testhelper.DiffFloat(t, tc.IDStr(), "", v.V, tc.expVal, 0.0000001)
	}
}
