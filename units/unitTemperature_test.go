package units_test

import (
	"testing"

	"github.com/nickwells/mathutil.mod/mathutil"
	"github.com/nickwells/testhelper.mod/testhelper"
	"github.com/nickwells/units.mod/units"
)

func TestTemperature(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		degC float64
		degF float64
		K    float64
	}{
		{
			ID:   testhelper.MkID("0℃"),
			degC: 0,
			degF: 32,
			K:    237.15,
		},
		{
			ID:   testhelper.MkID("100℃"),
			degC: 100,
			degF: 212,
			K:    337.15,
		},
	}

	for _, tc := range testCases {
		vc := units.ValWithUnit{Val: tc.degC, U: units.DegCUnit}

		vf, err := vc.Convert(units.DegFUnit)
		testhelper.CheckExpErr(t, err, tc)
		if err != nil {
			continue
		}
		if !mathutil.AlmostEqual(vf.Val, tc.degF, 0.000001) {
			t.Log(tc.IDStr())
			t.Logf("\t:    Given %6.2f℃\n", tc.degC)
			t.Logf("\t: Expected %6.2f℉\n", tc.degF)
			t.Logf("\t:      Got %6.2f℉\n", vf.Val)
			t.Errorf("\t: bad conversion from C to F\n")
		}

		vk, err := vc.Convert(units.DegKUnit)
		testhelper.CheckExpErr(t, err, tc)
		if err != nil {
			continue
		}
		if !mathutil.AlmostEqual(vk.Val, tc.K, 0.000001) {
			t.Log(tc.IDStr())
			t.Logf("\t:    Given %6.2f℃\n", tc.degC)
			t.Logf("\t: Expected %6.2fK\n", tc.K)
			t.Logf("\t:      Got %6.2fK\n", vk.Val)
			t.Errorf("\t: bad conversion from C to K\n")
		}

		vc, err = vf.Convert(units.DegCUnit)
		testhelper.CheckExpErr(t, err, tc)
		if err != nil {
			continue
		}
		if !mathutil.AlmostEqual(vc.Val, tc.degC, 0.000001) {
			t.Log(tc.IDStr())
			t.Logf("\t:    Given %6.2f℉\n", vf.Val)
			t.Logf("\t: Expected %6.2f℃\n", tc.degC)
			t.Logf("\t:      Got %6.2f℃\n", vc.Val)
			t.Errorf("\t: bad conversion from F back to C\n")
		}

		vc, err = vk.Convert(units.DegCUnit)
		testhelper.CheckExpErr(t, err, tc)
		if err != nil {
			continue
		}
		if !mathutil.AlmostEqual(vc.Val, tc.degC, 0.000001) {
			t.Log(tc.IDStr())
			t.Logf("\t:    Given %6.2fK\n", vk.Val)
			t.Logf("\t: Expected %6.2f℃\n", tc.degC)
			t.Logf("\t:      Got %6.2f℃\n", vc.Val)
			t.Errorf("\t: bad conversion from K back to C\n")
		}
	}
}
