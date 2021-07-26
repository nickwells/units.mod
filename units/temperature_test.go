package units

import (
	"testing"

	"github.com/nickwells/mathutil.mod/mathutil"
	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestTemperature(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		degC float64
		vals []ValUnit
	}{
		{
			ID:   testhelper.MkID("0℃"),
			degC: 0,
			vals: []ValUnit{
				{U: degFUnit, V: 32},
				{U: degReUnit, V: 0},
				{U: degRaUnit, V: 491.67},
				{U: degRoUnit, V: 7.5},
				{U: degKUnit, V: 273.15},
				{U: degNUnit, V: 0},
				{U: degDUnit, V: 150},
			},
		},
		{
			ID:   testhelper.MkID("100℃"),
			degC: 100,
			vals: []ValUnit{
				{U: degFUnit, V: 212},
				{U: degRaUnit, V: 671.67},
				{U: degReUnit, V: 80},
				{U: degRoUnit, V: 60},
				{U: degKUnit, V: 373.15},
				{U: degNUnit, V: 33.0},
				{U: degDUnit, V: 0},
			},
		},
		{
			ID:   testhelper.MkID("absolute zero"),
			degC: -273.15,
			vals: []ValUnit{
				{U: degFUnit, V: -459.67},
				{U: degRaUnit, V: 0},
				{U: degReUnit, V: -218.52},
				{U: degRoUnit, V: -135.90375},
				{U: degKUnit, V: 0},
				{U: degNUnit, V: -90.1395},
				{U: degDUnit, V: 559.725},
			},
		},
		{
			ID: testhelper.MkID("bad unit family"),
			ExpErr: testhelper.MkExpErr(
				"Mismatched unit families." +
					" Cannot convert units from" +
					" temperature to sample-family"),
			degC: 99,
			vals: []ValUnit{
				{U: SampleFamily.GetUnitOrPanic(SampleUnitBase), V: 0},
			},
		},
	}

	for _, tc := range testCases {
		vc := ValUnit{V: tc.degC, U: degCUnit}
		for _, v := range tc.vals {
			unitVal, err := vc.Convert(v.U)
			testhelper.CheckExpErr(t, err, tc)
			if err != nil {
				continue
			}
			if !mathutil.AlmostEqual(unitVal.V, v.V, 0.000001) {
				t.Log(tc.IDStr())
				t.Logf("\t:    Given %6.2f℃\n", tc.degC)
				t.Logf("\t: Expected %6.2f%s\n", v.V, v.U.abbrev)
				t.Logf("\t:      Got %6.2f%s\n", unitVal.V, v.U.abbrev)
				t.Errorf("\t: bad conversion from degrees Celsius to %s\n",
					v.U.namePlural)
			}

			vCelsius, err := v.Convert(degCUnit)
			testhelper.CheckExpErr(t, err, tc)
			if err != nil {
				continue
			}
			if !mathutil.AlmostEqual(vCelsius.V, tc.degC, 0.000001) {
				t.Log(tc.IDStr())
				t.Logf("\t:    Given %6.2f%s\n", v.V, v.U.abbrev)
				t.Logf("\t: Expected %6.2f℃\n", tc.degC)
				t.Logf("\t:      Got %6.2f℃\n", vCelsius.V)
				t.Errorf("\t: bad conversion from %s back to degrees Celsius\n",
					v.U.namePlural)
			}
		}
	}
}
