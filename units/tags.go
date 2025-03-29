package units

// Tag is the type of a descriptive tag that can be attached to a unit
type Tag string

// These Tag name constants should be used when setting tags on a new unit or
// when checking unit tags.
const (
	TagSI            = Tag("SI")
	TagMetric        = Tag("metric")
	TagHist          = Tag("historic")
	TagApothecary    = Tag("apothecary")
	TagImperial      = Tag("imperial")
	TagUScustomary   = Tag("US customary")
	TagColloquial    = Tag("colloquial")
	TagAstro         = Tag("astronomical")
	TagComputing     = Tag("computing")
	TagPrint         = Tag("printing")
	TagPhysics       = Tag("physics")
	TagTrig          = Tag("trigonometric")
	TagNautical      = Tag("nautical")
	TagDrinks        = Tag("drink")
	TagDimensionless = Tag("dimensionless")
)

var tags = map[Tag]string{
	TagSI: "an SI (Système Internationale) unit." +
		" Note that this includes SI derived units",
	TagMetric: "based on the metre." +
		" It is like the SI and is a forerunner of it" +
		" but includes some non-standard units.",
	TagHist: "a historical unit. No longer in common usage.",
	TagApothecary: "an apothecary's unit." +
		" This system of units is based on" +
		" the grain as the unit of mass and" +
		" the fluid ounce as the unit of volume." +
		" Apothecary is a historic term for" +
		" someone who would prepare and sell medicine," +
		" the precursor of today's pharmacist." +
		" It was the standard for the practice of pharmacy" +
		" in the United States until 1971" +
		" when it was replaced by the metric system.",
	TagImperial: "an imperial unit." +
		" This system of units was" +
		" originally in use in the United Kingdom" +
		" and other Commonwealth countries." +
		" It is now mostly replaced by the metric system." +
		" It is similar to the system used in the United States" +
		" though it differs in some values." +
		" It was first defined in the British Weights and Measures Act 1824.",
	TagUScustomary: "a United States customary unit." +
		" The units are derived from the English units" +
		" used in the British Empire before American independence" +
		" but there are significant differences from the imperial units" +
		" particularly in units of volume and fluid measures." +
		" It is now (2024) only used in" +
		" the United States, Myanmar and Liberia.",
	TagColloquial: "a colloquial unit." +
		" Not formally defined but sometimes used in conversation.",
	TagAstro:     "a unit sometimes used in astronomical calculations.",
	TagComputing: "a unit sometimes used in comnputing.",
	TagPrint:     "a unit used in printing",
	TagPhysics: "a unit used in physics." +
		" Of course most units are used in physics" +
		" but these are not used generally outside physics.",
	TagTrig:          "a unit used in trigonometric calculations.",
	TagNautical:      "a unit used in nautical discussions",
	TagDrinks:        "a unit used in the drinks industry (wine/brewing)",
	TagDimensionless: "a dimensionless unit.",
}

// Notes returns the description of the given tag or a blank string if no
// notes are available for the tag.
func (t Tag) Notes() string {
	return tags[t]
}

// IsValid returns the true if the tag is found in the set of valid tags,
// false otherwise.
func (t Tag) IsValid() bool {
	_, ok := tags[t]
	return ok
}

// GetTagNames returns a list of the tag names (as strings). Note that the
// list is unsorted and its order may vary between calls to this function.
func GetTagNames() []string {
	rval := []string{}

	for t := range tags {
		rval = append(rval, string(t))
	}

	return rval
}
