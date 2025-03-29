package units

import "math"

// common to all units

// Powers of 1000
const (
	k  = 1000.0                             // kilo
	_M = 1000_000.0                         // mega
	_G = 1000_000_000.0                     // giga
	_T = 1000_000_000_000.0                 // tera
	_P = 1000_000_000_000_000.0             // peta
	_E = 1000_000_000_000_000_000.0         // exa
	_Z = 1000_000_000_000_000_000_000.0     // zetta
	_Y = 1000_000_000_000_000_000_000_000.0 // yotta

	h  = 100.0 // hecto
	da = 10.0  // deca

	d = 0.1  // deci
	c = 0.01 // centi

	m = 1.0 / k  // milli
	u = 1.0 / _M // micro
	n = 1.0 / _G // nano
	p = 1.0 / _T // pico
	f = 1.0 / _P // femto
	a = 1.0 / _E // atto
	z = 1.0 / _Z // zepto
	y = 1.0 / _Y // yocto
)

// angles

const (
	degreePerRadian  = math.Pi / 180
	degreePerGradian = math.Pi / 200
)

// area

const (
	yard2ToMetre2 = yardToMetre * yardToMetre
	acreToMetre2  = yard2ToMetre2 * 4840
)

// data

// powers of 1024, the names are from the standard ISO/IEC 80000
//
//nolint:mnd
const (
	_  = 1 << (10 * iota)
	ki // kibi
	mi // mebi
	gi // gibi
	ti // tebi
	pi // pebi
	ei // exbi
	zi // zebi
	yi // yobi
)

// distance

const (
	inchToMetre = 0.0254
	footToMetre = inchToMetre * 12
	yardToMetre = footToMetre * 3
	mileToMetre = yardToMetre * 1760
	rodToMetre  = yardToMetre * 5.5

	nauticalMileToMetre = 1852

	// note that this is a precise conversion value as the metre is defined
	// as the distance light travels in a vacuum in this fraction of a second
	lightSecond  = 2.99792458e8 // speedOfLight * 1 second
	speedOfLight = lightSecond  // metres / second

	pointToMetre = 0.0003528

	// Note that the length of a TEU (a Twenty-foot Equivalent Unit or
	// Shipping Container) is not twenty feet. See the unit declaration for
	// an explanation.
	teuLength = (19 * footToMetre) + (10.5 * inchToMetre) //nolint:mnd
)

// energy

const (
	calorieToJoule     = 4.184
	footPoundToJoule   = 1.3558179483314004
	footPoundalToJoule = 0.0421401100938048

	electronVoltToJoule = 1.602176634e-19

	// Note that there are many different variants of the British thermal
	// unit with conversion ratios to Joules varying between 1054.3503 (for
	// the Thermochemical variant) and 1059.67 (using the calorie value of
	// water at its maximum density, when it is at 3.9 °C)
	btuToJoule = 1055.06 // International Standard ISO 31-4
)

// mass

const (
	grainToGram = 0.06479891
	ounceToGram = grainToGram * 437.5
	poundToGram = ounceToGram * 16
)

// temperature

const absZero = 273.15

// time

const (
	minToSec  = 60
	hourToSec = minToSec * 60
	dayToSec  = hourToSec * 24
)

// volume

const (
	litreToCubicMetre       = 1e-3
	cubicInchToCubicMetre   = (inchToMetre * inchToMetre * inchToMetre)
	cubicFootToCubicMetre   = (footToMetre * footToMetre * footToMetre)
	cubicYardToCubicMetre   = (yardToMetre * yardToMetre * yardToMetre)
	bottleToCubicMetre      = 0.75e-3
	gallonToCubicMetre      = 4.54609e-3
	fluidOzToCubicMetre     = gallonToCubicMetre / 160
	usFluidOzToCubicMetre   = 29.5735295625e-6
	usDryGallonToCubicMetre = 4.40488377086e-3
)
