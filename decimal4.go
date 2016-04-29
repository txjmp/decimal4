package decimal4

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
)

const Dollar = "\u0024"
const Euro = "\u20AC"
const Yen = "\u00A5"
const Yuan = "\u00A5"
const Rupee = "\u20B9"
const Ruble = "\u20BD"
const Pound = "\u00A3"

var Decimal4StringPlaces string = "4" // precision used by String method

type Decimal4 int64
type Decimal6 int64

// Multiply returns product of this * x, rounded to 4 decimal places.
func (this Decimal4) Multiply(x Decimal4) Decimal4 {
	a := this * x
	if a == 0 {
		return 0
	}
	if a/x != this {
		log.Panic("Decimal4 Multiply Overflow, this=", this, " x=", x)
	}
	if a > 0 {
		a += 5000
	} else {
		a -= 5000
	}
	return a / 10000
}

// MultRound2 returns product of this * x, rounded to 2 decimal places.
func (this Decimal4) MultRound2(x Decimal4) Decimal4 {
	a := this * x
	if a == 0 {
		return 0
	}
	if a/x != this {
		log.Panic("Decimal4 MultRound2 Overflow, this=", this, " x=", x)
	}
	if a > 0 {
		a += 500000
	} else {
		a -= 500000
	}
	return (a / 1000000) * 100
}

// M is a fast version of Multiply, no rounding, no check for overflow.
func (this Decimal4) M(x Decimal4) Decimal4 {
	return (this * x) / 10000
}

// Multiply6 returns product of this * x rounded to 4 decimal places.
// Parameter x is type Decimal6, providing up to 6 places precision.
func (this Decimal4) Multiply6(x Decimal6) Decimal4 {
	a := Decimal6(this) * x
	if a == 0 {
		return 0
	}
	if a/x != Decimal6(this) {
		log.Panic("Decimal4 Multiply6 Overflow, this=", this, " x=", x)
	}
	if a > 0 {
		a += 500000
	} else {
		a -= 500000
	}
	return Decimal4(a / 1000000)
}

// MultiplyBig allows for a larger maximum value than Multiply (before exceeding int64 max).
// Last 2 decimal places are truncated on largest input value.
func (this Decimal4) MultiplyBig(x Decimal4) Decimal4 {
	var a, b, c Decimal4
	if Abs(this) > Abs(x) {
		a = this
		b = x
	} else {
		a = x
		b = this
	}
	a = a / 100 // knock off last 2 decimal places of largest value
	c = a * b
	if c == 0 {
		return 0
	}
	if c/a != b {
		log.Panic("Decimal4 MultiplyBig Overflow, this=", this, " x=", x)
	}
	if c > 0 {
		c += 50
	} else {
		c -= 50
	}
	return c / 100
}

// MultiplyBig6 allows for a larger maximum value than Multiply6 (before exceeding int64 max).
// Last 2 decimal places are truncated on this value.
func (this Decimal4) MultiplyBig6(x Decimal6) Decimal4 {
	a := Decimal6(this / 100) // knock off last 2 decimal places
	b := a * x
	if b == 0 {
		return 0
	}
	if b/x != a {
		log.Panic("Decimal4 MultiplyBig6 Overflow, this=", this, " x=", x)
	}
	if b > 0 {
		b += 5000
	} else {
		b -= 5000
	}
	return Decimal4(b / 10000)
}

// MultiplyInt returns product of this * x.
// Parameter x is type int.
func (this Decimal4) MultiplyInt(x int) Decimal4 {
	a := int64(this) * int64(x)
	if a == 0 {
		return 0
	}
	if a/int64(x) != int64(this) {
		log.Panic("Decimal4 MultiplyInt Overflow, this=", this, " x=", x)
	}
	return Decimal4(a)
}

// Divide returns quotient of this / x rounded to 4 decimal places.
func (this Decimal4) Divide(x Decimal4) Decimal4 {
	if this == 0 {
		return 0
	}
	a := this * 100000 // shift over 5 places rather than 4, so result can be rounded
	if a/100000 != this {
		log.Panic("Decimal4 Divide Overflow, this=", this, " x=", x, " a=", a)
	}
	b := a / x
	if b > 0 {
		b += 5
	} else {
		b -= 5
	}
	return b / 10
}

// DivideBig returns quotient of this / x, 3 decimal places precision, no rounding.
func (this Decimal4) DivideBig(x Decimal4) Decimal4 {
	if this == 0 {
		return 0
	}
	a := this * 1000
	if a/1000 != this {
		log.Panic("Decimal4 DivideBig Overflow, this=", this, " x=", x, " a=", a)
	}
	b := (a / x) * 10
	if b/10 != (a / x) {
		log.Panic("Decimal4 DivideBig Overflow, this=", this, " x=", x, " a=", a)
	}
	return b
}

// DivideInt returns quotient of this / x rounded to 4 decimal places.
// Parameter x is type int.
func (this Decimal4) DivideInt(x int) Decimal4 {
	if this == 0 {
		return 0
	}
	a := int64(this) * 10 // shift over 1 position so result can be rounded
	if a/10 != int64(this) {
		log.Panic("Decimal4 DivideInt Overflow, this=", this, " x=", x, " a=", a)
	}
	b := a / int64(x)
	if b > 0 {
		b += 5
	} else {
		b -= 5
	}
	return Decimal4(b / 10)
}

// return true if difference in values is < .1
func (this Decimal4) CloseTo(x Decimal4) bool {
	diff := this - x
	if diff == 0 {
		return true
	}
	if diff > 0 && diff < 1000 {
		return true
	} else if diff < 0 && diff > -1000 {
		return true
	}
	return false
}

// Round Methods
// Result rounded to specified number of decimal places
// Result still has 4 implied decimal places
func (this Decimal4) Round0() Decimal4 { // 1235555 -> 1240000
	if this == 0 {
		return 0
	}
	if this > 0 {
		return ((this + 5000) / 10000) * 10000
	}
	return ((this - 5000) / 10000) * 10000
}
func (this Decimal4) Round1() Decimal4 { // 1235555 -> 1236000
	if this == 0 {
		return 0
	}
	if this > 0 {
		return ((this + 500) / 1000) * 1000
	}
	return ((this - 500) / 1000) * 1000
}
func (this Decimal4) Round2() Decimal4 { // 1235555 -> 1235600
	if this == 0 {
		return 0
	}
	if this > 0 {
		return ((this + 50) / 100) * 100
	}
	return ((this - 50) / 100) * 100
}
func (this Decimal4) Round3() Decimal4 { // 1235555 -> 1235560
	if this == 0 {
		return 0
	}
	if this > 0 {
		return ((this + 5) / 10) * 10
	}
	return ((this - 5) / 10) * 10
}

// Truncate Methods
// Result truncated to specified number of decimal places
// Result still has 4 implied decimal places, but truncated places are zero
func (this Decimal4) Truncate0() Decimal4 {
	return (this / 10000) * 10000 // 1235555 -> 1230000
}
func (this Decimal4) Truncate1() Decimal4 {
	return (this / 1000) * 1000 // 1235555 -> 1235000
}
func (this Decimal4) Truncate2() Decimal4 {
	return (this / 100) * 100 // 1235555 -> 1235500
}
func (this Decimal4) Truncate3() Decimal4 {
	return (this / 10) * 10 // 1235555 -> 1235550
}

func (this Decimal4) String() string {
	format := "%." + Decimal4StringPlaces + "f"
	return fmt.Sprintf(format, float64(this)/10000)
}

func (this Decimal6) String() string {
	format := "%.6f"
	return fmt.Sprintf(format, float64(this)/1000000)
}

func (this Decimal4) Fmt(widthPrecision float64, currency ...string) string {
	format := "%" + strconv.FormatFloat(widthPrecision, 'f', 1, 64) + "f"
	fmtNum := fmt.Sprintf(format, float64(this)/10000)
	if len(currency) == 0 && Abs(this) < 10000000 { // < 1 thousand
		return fmtNum
	}
	if len(currency) == 0 {
		return addCommas(fmtNum, "")
	}
	return addCommas(fmtNum, currency[0])
}

var space byte = 32
var comma byte = 44
var minusSign byte = 45

func addCommas(in, currency string) string {
	inBytes := []byte(in)
	spaceCount := bytes.Count(inBytes, []byte(" ")) // number of leading spaces
	dotNdx := bytes.Index(inBytes, []byte("."))     // location of decimal point

	if dotNdx == -1 { // if no decimal point, assumed after last digit
		dotNdx = len(inBytes)
	}
	// commaLocations are indexes where a comma should be inserted
	commaLocations := []int{dotNdx - 4, dotNdx - 7, dotNdx - 10, dotNdx - 13}
	commaNdx := 0 // used in loop below, indicates which commaLocation to use in comparison

	outBytes := make([]byte, 50)

	// load outBytes from inBytes, beginning with last byte, adding commas at commaLocations
	outNdx := len(outBytes)
	for i := len(inBytes) - 1; i > -1; i-- {
		if inBytes[i] == space { // done when 1st leading space hit
			break
		}
		outNdx--
		if i == commaLocations[commaNdx] && inBytes[i] != minusSign {
			outBytes[outNdx] = comma
			outNdx--
			commaNdx++
			if spaceCount > 0 {
				spaceCount-- // for each comma added, reduce number of leading spaces
			}
		}
		outBytes[outNdx] = inBytes[i]
	}
	result := make([]byte, 0, 50)
	if currency == "" {
		if spaceCount > 0 {
			result = append(result, inBytes[0:spaceCount]...) // add leading spaces
		}
	} else {
		if spaceCount > 1 {
			result = append(result, inBytes[0:spaceCount-1]...) // add leading spaces
		}
		symbol := []byte(currency)
		result = append(result, symbol...) // add currency symbol
	}
	result = append(result, outBytes[outNdx:]...) // add outBytes to result, removing leading unloaded bytes
	return string(result)
}

func New(x float64) Decimal4 {
	//fixer := math.Copysign(.00003, x)  // if x < 0, fixer = -.00003 ... much slower, don't use
	var fixer float64 = .00003
	if x > 0 {
		return Decimal4((x + fixer) * 10000)
	} else if x < 0 {
		return Decimal4((x - fixer) * 10000)
	}
	return 0
}

func NewDecimal6(x float64) Decimal6 {
	var fixer float64 = .0000003
	if x > 0 {
		return Decimal6((x + fixer) * 1000000)
	} else if x < 0 {
		return Decimal6((x - fixer) * 1000000)
	}
	return 0
}

func Abs(x Decimal4) Decimal4 {
	if x < 0 {
		return -x
	}
	return x
}

// RoundFloat2 - round float64 to 2 decimal places
func RoundFloat2(x float64) float64 {
	var multiplier float64 = 100
	var rounder float64 = .005
	var intX int64 = int64((x + rounder) * multiplier)
	return float64(intX) / multiplier
}

// RoundFloat - round float64 to precision decimal places
func RoundFloat(x float64, precision int) float64 {
	var multiplier float64 = math.Pow(10, float64(precision))         // 1, 10, 100, etc.
	var rounder float64 = math.Pow(10, float64((precision+1)*-1)) * 5 // .5, .05, .005, etc.
	var intX int64 = int64((x + rounder) * multiplier)
	return float64(intX) / multiplier
}
