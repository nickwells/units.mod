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
		vals map[units.Unit]float64
	}{
		{
			ID:   testhelper.MkID("0℃"),
			degC: 0,
			vals: map[units.Unit]float64{
				units.DegFUnit:  32,
				units.DegRaUnit: 491.67,
				units.DegRoUnit: 7.5,
				units.DegKUnit:  273.15,
				units.DegNUnit:  0,
				units.DegDUnit:  150,
			},
		},
		{
			ID:   testhelper.MkID("100℃"),
			degC: 100,
			vals: map[units.Unit]float64{
				units.DegFUnit:  212,
				units.DegRaUnit: 671.67,
				units.DegRoUnit: 60,
				units.DegKUnit:  373.15,
				units.DegNUnit:  33.0,
				units.DegDUnit:  0,
			},
		},
		{
			ID:   testhelper.MkID("absolute zero"),
			degC: -273.15,
			vals: map[units.Unit]float64{
				units.DegFUnit:  -459.67,
				units.DegRaUnit: 0,
				units.DegRoUnit: -135.90375,
				units.DegKUnit:  0,
				units.DegNUnit:  -90.1395,
				units.DegDUnit:  559.725,
			},
		},
		{
			ID: testhelper.MkID("bad unit family"),
			ExpErr: testhelper.MkExpErr(
				"Mismatched unit families." +
					" Cannot convert units from" +
					" C (a unit of temperature)" +
					" to gram (a unit of mass)"),
			degC: 99,
			vals: map[units.Unit]float64{
				units.GramUnit: 0,
			},
		},
		{
			ID:     testhelper.MkID("bad unit - zero conversion factor"),
			ExpErr: testhelper.MkExpErr("Bad units - a zero conversion factor"),
			degC:   99,
			vals: map[units.Unit]float64{
				{
					ConvPreAdd:  0,
					ConvPostAdd: 0,
					ConvFactor:  0,
					Fam:         units.UnitOfTemperature,
					Abbrev:      "Bad",
					Name:        "Bad singular",
					NamePlural:  "Bad plural",
					Notes:       "Description",
				}: 0,
			},
		},
	}

	for _, tc := range testCases {
		vc := units.ValWithUnit{Val: tc.degC, U: units.DegCUnit}
		for altUnit, v := range tc.vals {
			unitVal, err := vc.Convert(altUnit)
			testhelper.CheckExpErr(t, err, tc)
			if err != nil {
				continue
			}
			if !mathutil.AlmostEqual(unitVal.Val, v, 0.000001) {
				t.Log(tc.IDStr())
				t.Logf("\t:    Given %6.2f℃\n", tc.degC)
				t.Logf("\t: Expected %6.2f%s\n", v, altUnit.Abbrev)
				t.Logf("\t:      Got %6.2f%s\n", unitVal.Val, altUnit.Abbrev)
				t.Errorf("\t: bad conversion from degrees Celsius to %s\n",
					altUnit.NamePlural)
			}
			vAlt := units.ValWithUnit{Val: v, U: altUnit}

			vCelsius, err := vAlt.Convert(units.DegCUnit)
			testhelper.CheckExpErr(t, err, tc)
			if err != nil {
				continue
			}
			if !mathutil.AlmostEqual(vCelsius.Val, tc.degC, 0.000001) {
				t.Log(tc.IDStr())
				t.Logf("\t:    Given %6.2f%s\n", v, altUnit.Abbrev)
				t.Logf("\t: Expected %6.2f℃\n", tc.degC)
				t.Logf("\t:      Got %6.2f℃\n", vCelsius.Val)
				t.Errorf("\t: bad conversion from %s back to degrees Celsius\n",
					altUnit.NamePlural)
			}
		}
	}
}
