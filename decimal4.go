package decimal4

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"log"
)

var Decimal4StringPlaces string = "4" // precision used by String method

type Decimal4 int64
type Decimal6 int64

func (this Decimal4) Multiply(x Decimal4) Decimal4 {
	a := this * x
	if a == 0 {
		return 0
	}
	if a / x != this {
		log.Panic("Decimal4 Multiply Overflow, this=", this, " x=", x)
	}
	if a < 0 {
		a -= 5000
	} else {
		a += 5000
	}
	return a / 10000
}

func (this Decimal4) Multiply6(x Decimal6) Decimal4 {
	a := Decimal6(this) * x
	if a == 0 {
		return 0
	}
	if a / x != Decimal6(this) {
		log.Panic("Decimal4 Multiply6 Overflow, this=", this, " x=", x)
	}
	if a < 0 {
		a -= 500000
	} else {
		a += 500000
	}
	return Decimal4(a / 1000000)
}

var bigLimit Decimal4 = 100000000 // 10000.0000, if inputs less than, use normal method

// MultiplyBig - Allows for a larger maximum value (before exceeding int64 max).
// Last 2 decimal places are truncated on input value > bigLimit
// Intermediate product value will only contain 6 implied decimal places rather than 8.
// If neither value is > bigLimit, calls Multiply method.
func (this Decimal4) MultiplyBig(x Decimal4) Decimal4 {
	var a, b Decimal4
	if (this > 0 && this > x) || (this < 0 && this < x) {
		a = this
		b = x
	} else {
		a = x 
		b = this
	}
	if a < bigLimit {  // both values small enough that regular multiply can be used
		return this.Multiply(x)
	}
	a2 := a / 100  // knock off last 2 decimal places of largest value
	c := a2 * b
	if c == 0 {
		return 0
	}
	if c / a2 != b {
		log.Panic("Decimal4 MultiplyBig Overflow, this=", this, " x=", x)
	}
	if c < 0 {
		c -= 50
	} else {
		c += 50
	}
	return c / 100
}

func (this Decimal4) MultiplyBig6(x Decimal6) Decimal4 {
	this2 := Decimal6(this / 100)  // knock off last 2 decimal places
	a := this2 * x
	if a == 0 {
		return 0
	}
	if a / x != this2 {
		log.Panic("Decimal4 MultiplyBig6 Overflow, this=", this, " x=", x)
	}
	if a < 0 {
		a -= 5000
	} else {
		a += 5000
	}
	return Decimal4(a / 10000)
}
// 4 decimal place precision with rounding
func (this Decimal4) Divide(x Decimal4) Decimal4 {
	a := this * 100000 // shift over 5 places rather than 4, so result can be rounded
	if a / 100000 != this {
		log.Panic("Decimal4 Divide Overflow, this=", this, " x=", x, " a=", a)
	}
	b := a / x
	if b == 0 {
		return 0
	}
	if b < 0 {
		b -= 5
	} else {
		b += 5
	}
	return b / 10
}

// 3 decimal place precision, no rounding
func (this Decimal4) DivideBig(x Decimal4) Decimal4 {
	a := this * 1000
	if a / 1000 != this {
		log.Panic("Decimal4 DivideBig Overflow, this=", this, " x=", x, " a=", a)
	}
	b := (a / x) * 10
	if b / 10 != (a / x) {
		log.Panic("Decimal4 DivideBig Overflow, this=", this, " x=", x, " a=", a)
	}
	return b
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

func (this Decimal4) Round0() Decimal4 {
	if this < 0 {
		return ((this - 5000) / 10000) * 10000
	} else {
		return ((this + 5000) / 10000) * 10000
	}
}
func (this Decimal4) Round1() Decimal4 {
	if this < 0 {
		return ((this - 500) / 1000) * 1000
	} else {
		return ((this + 500) / 1000) * 1000
	}
}
func (this Decimal4) Round2() Decimal4 {
	if this < 0 {
		return ((this - 50) / 100) * 100
	} else {
		return ((this + 50) / 100) * 100
	}
}
func (this Decimal4) Round3() Decimal4 {
	if this < 0 {
		return ((this - 5) / 10) * 10
	} else {
		return ((this + 5) / 10) * 10
	}
}

func (this Decimal4) String() string {
	format := "%." + Decimal4StringPlaces + "f"
	return fmt.Sprintf(format, float64(this)/10000)
}

func (this Decimal6) String() string {
	format := "%.6f"
	return fmt.Sprintf(format, float64(this)/1000000)
}

func (this Decimal4) Format(places int) string {
	if places > 4 || places < 0 {
		places = 4
	}
	format := "%." + strconv.Itoa(places) + "f"
	fmtNum := fmt.Sprintf(format, float64(this)/10000)
	if this < 10000000 && this > -10000000 { // 1000.0000
		return fmtNum
	} else {
		return addCommas([]byte(fmtNum))
	}
}

func addCommas(in []byte) string {
	dotNdx := bytes.Index(in, []byte("."))
	if dotNdx == -1 { // if no dot, assumed after last digit
		dotNdx = len(in)
	}
	commaLocations := []int{dotNdx - 4, dotNdx - 7, dotNdx - 10, dotNdx - 13}
	commaNdx := 0

	out := make([]byte, len(in)+4)
	outNdx := len(out)

	var comma byte = 44 // ascii values
	var minusSign byte = 45

	for i := len(in) - 1; i > -1; i-- {
		outNdx--
		if i == commaLocations[commaNdx] && in[i] != minusSign {
			commaNdx++
			out[outNdx] = comma
			outNdx--
		}
		out[outNdx] = in[i]
	}
	return string(out[outNdx:]) // remove leading bytes not loaded
}

func New(x float64) Decimal4 {
	if x < 0 {
		return Decimal4((x - .00005) * 10000)
	} else {
		return Decimal4((x + .00005) * 10000)
	}
}

func NewDecimal6(x float64) Decimal6 {
	if x < 0 {
		return Decimal6((x - .0000005) * 1000000)
	} else {
		return Decimal6((x + .0000005) * 1000000)
	}
}

// --- funcs that round float numbers ------------------------
func Exact2(x float64) float64 {
	var multiplier float64 = 100
	var rounder float64 = .005
	var intX int64 = int64((x + rounder) * multiplier)
	return float64(intX) / multiplier
}
func Exact(x float64, dec int) float64 {
	var multiplier float64 = math.Pow(10, float64(dec))         // 1, 10, 100, etc.
	var rounder float64 = math.Pow(10, float64((dec+1)*-1)) * 5 // .5, .05, .005, etc.
	var intX int64 = int64((x + rounder) * multiplier)
	return float64(intX) / multiplier
}
