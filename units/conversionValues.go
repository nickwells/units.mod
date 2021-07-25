package units

import "math"

// common to all units

// Powers of 1000
var (
	k  = 1000.0
	_M = math.Pow(1000, 2)
	_G = math.Pow(1000, 3)
	_T = math.Pow(1000, 4)
	_P = math.Pow(1000, 5)
	_E = math.Pow(1000, 6)
	_Z = math.Pow(1000, 7)
	_Y = math.Pow(1000, 8)

	h  = 100.0
	da = 10.0

	d = 0.1
	c = 0.01

	m = 1.0 / math.Pow(1000, 1)
	u = 1.0 / math.Pow(1000, 2)
	n = 1.0 / math.Pow(1000, 3)
	p = 1.0 / math.Pow(1000, 4)
	f = 1.0 / math.Pow(1000, 5)
	a = 1.0 / math.Pow(1000, 6)
	z = 1.0 / math.Pow(1000, 7)
	y = 1.0 / math.Pow(1000, 8)
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
	cubicInchToCubicMetre   = (inchToMetre * inchToMetre * inchToMetre)
	cubicFootToCubicMetre   = (footToMetre * footToMetre * footToMetre)
	cubicYardToCubicMetre   = (yardToMetre * yardToMetre * yardToMetre)
	bottleToCubicMetre      = 0.75e-3
	gallonToCubicMetre      = 4.54609e-3
	fluidOzToCubicMetre     = gallonToCubicMetre / 160
	usFluidOzToCubicMetre   = 29.5735295625e-6
	usDryGallonToCubicMetre = 4.40488377086e-3
)
