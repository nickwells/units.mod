package units_test

import (
	"sort"
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
	"github.com/nickwells/units.mod/v2/units"
)

func TestTagNote(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		tag     units.Tag
		expNote string
	}{
		{
			ID:      testhelper.MkID("..."),
			tag:     units.Tag("nonesuch"),
			expNote: "",
		},
	}

	for _, tc := range testCases {
		note := tc.tag.Notes()
		testhelper.DiffString(t, tc.IDStr(), "note", note, tc.expNote)
	}
}

func TestTagNames(t *testing.T) {
	expTagNames := []string{
		"SI",
		"metric",
		"historic",
		"apothecary",
		"imperial",
		"US customary",
		"colloquial",
		"astronomical",
		"computing",
		"printing",
		"physics",
		"trigonometric",
		"nautical",
		"drink",
		"dimensionless",
	}
	sort.Strings(expTagNames)

	actTagNames := units.GetTagNames()
	sort.Strings(actTagNames)

	testhelper.DiffStringSlice(t, "GetTagNames", "", actTagNames, expTagNames)
}
