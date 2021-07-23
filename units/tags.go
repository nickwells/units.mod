package units

type Tag string

// These Tag name constants should be used when setting tags on a new unit or
// when checking unit tags.
const (
	TagSI          = Tag("SI")
	TagMetric      = Tag("metric")
	TagHist        = Tag("historic")
	TagImperial    = Tag("imperial")
	TagUScustomary = Tag("US customary")
	TagColloquial  = Tag("colloquial")
	TagAstro       = Tag("astronomical")
	TagComputing   = Tag("computing")
	TagPrint       = Tag("printing")
	TagPhysics     = Tag("physics")
	TagNautical    = Tag("nautical")
	TagDrinks      = Tag("drink")
)

var tags = map[Tag]string{
	TagSI: "an SI (Système Internationale) unit.",
	TagMetric: "based on the metre." +
		" It is like SI but includes some non-standard units.",
	TagHist: "a historical unit. No longer in common usage.",
	TagImperial: "an imperial unit." +
		" This system of units was" +
		" originally in use in the United Kingdom" +
		" and other Commonwealth countries." +
		" It is now mostly replaced by the metric system." +
		" It is similar to the system used in the United States" +
		" though it differs in some values." +
		" It was first defined in the British Weights and Measures Act 1824.",
	TagUScustomary: "a United States customary units." +
		" The units are derived from the English units" +
		" used in the British Empire before American independence" +
		" but there are significant differences from the imperial units.",
	TagColloquial: "a colloquial unit." +
		" Not formally defined but sometimes used in conversation.",
	TagAstro:     "a unit sometimes used in astronomical calculations.",
	TagComputing: "a unit sometimes used in comnputing.",
	TagPrint:     "a unit used in printing",
	TagPhysics: "a unit used in physics." +
		" Of course most units are used in physics" +
		" but these are not used generally outside physics.",
	TagNautical: "a unit used in nautical discussions",
	TagDrinks:   "a unit used in the drinks industry (wine/brewing)",
}

// Notes returns the description of the given tag or a blank string if no
// notes are available for the tag.
func (t Tag) Notes() string {
	return tags[t]
}
