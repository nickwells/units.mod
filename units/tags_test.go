package units_test

import (
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
