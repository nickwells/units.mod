package units

// These unit name constants should only be used when retrieving units from
// the SampleFamily for testing. Similarly for the aliases and Tags.
const (
	SampleFamilyDescription = "unit of test"

	SampleFamilyName    = "sample-family"
	SampleFamilyAlias   = "sample-family-alias"
	SampleFamilyBadName = "bad-sample-family"

	SampleUnitBase   = "sample"
	SampleUnitA      = "sa"
	SampleUnit2T     = "s2t"
	SampleUnit001    = "s001"
	SampleUnit123    = "s123"
	SampleUnitNeg123 = "s-1-2+3"
	SampleUnitBad    = "bad"

	SampleUnitBaseAlias = "smpl"
	SampleUnitAAlias    = "sample A"

	SampleTagName      = Tag("SampleTag")
	SampleTagNameOther = Tag("SampleTagOther")
)

// SampleFamily is a Family that can be used to test the behaviour of other
// packages that want to use the units package.
//
// Note that it is not a true unit family and cannot be retrieved using the
// GetFamily func; it should be used directly and only for testing.
var SampleFamily = &Family{
	baseUnitName:  SampleUnitBase,
	description:   SampleFamilyDescription,
	name:          SampleFamilyName,
	familyAliases: []string{SampleFamilyAlias},
}

var sampleNames = map[string]Unit{
	SampleUnitBase: {
		0, 0, 1,
		SampleFamily,
		"s", SampleUnitBase, "samples",
		"some brief notes about the base unit",
		[]Tag{SampleTagName},
		map[string]string{
			SampleUnitBaseAlias: "Alias",
		},
		"", "",
	},

	SampleUnitA: {
		0, 0, 2,
		SampleFamily,
		"sa", "sample A", "samples A",
		"notes about unit sa, which has an alias",
		[]Tag{SampleTagName},
		map[string]string{
			SampleUnitAAlias: "full name",
		},
		"", "",
	},

	SampleUnit2T: {
		0, 0, 3,
		SampleFamily,
		"s2t", "sample 2-tags", "samples 2-tags",
		"notes about unit s2t which has 2 tags",
		[]Tag{SampleTagName, SampleTagNameOther},
		map[string]string{},
		"", "",
	},

	SampleUnit001: {
		0, 0, 1,
		SampleFamily,
		"s001", "sample 0, 0, 1", "samples 0, 0, 1",
		"notes about unit s001",
		[]Tag{SampleTagName},
		map[string]string{},
		"", "",
	},

	SampleUnit123: {
		1, 2, 3,
		SampleFamily,
		"s123", "sample 1, 2, 3", "samples 1, 2, 3",
		"notes about unit s123",
		[]Tag{SampleTagName},
		map[string]string{},
		"", "",
	},

	SampleUnitNeg123: {
		-1, -2, 3,
		SampleFamily,
		"s-1-2+3", "sample -1, -2, 3", "samples -1, -2, 3",
		"notes about unit s-1-2+3",
		[]Tag{SampleTagName},
		map[string]string{},
		"", "",
	},

	SampleUnitBad: {
		0, 0, 0, // has a zero conversion factor
		SampleFamily,
		"bad", "bad", "bad",
		"bad unit, which has a zero conversion factor",
		[]Tag{SampleTagName},
		map[string]string{},
		"", "",
	},
}

func init() {
	SampleFamily.altUnits = sampleNames
	SampleFamily.populateUnitAliases()
}

// BadSampleFamily is a Family that can be used to test the behaviour of
// other packages that want to use the units package. It is 'bad' in the
// sense that it has no alternative Units.
//
// Note that it is not a true unit family and cannot be retrieved using the
// GetFamily func; it should be used directly and only for testing.
var BadSampleFamily = &Family{
	baseUnitName: SampleUnitBase,
	description:  SampleFamilyDescription,
	name:         SampleFamilyBadName,
	altUnits:     map[string]Unit{},
}
