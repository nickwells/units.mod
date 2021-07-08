package units

import (
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestVelocity(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		from    string
		to      string
		fromVal float64
		expVal  float64
		epsilon float64
	}{
		{
			ID:      testhelper.MkID("metresPerSec-FeetPerSec"),
			from:    "metresPerSecond",
			to:      "feetPerSecond",
			fromVal: 1,
			expVal:  3.28,
			epsilon: 0.1,
		},
		{
			ID:      testhelper.MkID("metresPerSec-kilometresPerHour"),
			from:    "metresPerSecond",
			to:      "kilometresPerHour",
			fromVal: 1,
			expVal:  3.6,
			epsilon: 0.00001,
		},
		{
			ID:      testhelper.MkID("metresPerSec-percentageOfSpeedOfLight"),
			from:    "metresPerSecond",
			to:      "percentOfSpeedOfLight",
			fromVal: 2.997e8,
			expVal:  99.9,
			epsilon: 0.1,
		},
		{
			ID:      testhelper.MkID("milesPerH-kmPerH"),
			from:    "milesPerHour",
			to:      "kilometresPerHour",
			fromVal: 100,
			expVal:  161,
			epsilon: 0.1,
		},
	}

	for _, tc := range testCases {
		fromUnit := GetUnitOrPanic(Velocity, tc.from)
		toUnit := GetUnitOrPanic(Velocity, tc.to)
		fv := ValWithUnit{tc.fromVal, fromUnit}
		tv := fv.ConvertOrPanic(toUnit)

		testhelper.DiffFloat64(t, tc.IDStr(), "velocity",
			tv.Val, tc.expVal, tc.epsilon)
	}
}
