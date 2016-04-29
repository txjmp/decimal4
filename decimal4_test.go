package decimal4

import "testing"
import "math"

type data struct {
	a float64 // input
	b float64 // input
	c float64 // result
}

func TestNew(t *testing.T) {
	minOk := New(200000000000) // 200,000,000,000
	var a, i int64
	var b float64
	var c Decimal4
	for i = 1; i < math.MaxInt64; i++ {
		a += i
		b = float64(a) / 10000
		c = New(b)
		if a != int64(c) {
			break
		}
	}
	if c < minOk {
		t.Fatal("did not reach minimum, largest value confirmed:", c)
	}
	t.Log(i, ", values checked, largest value confirmed: ", c)

	a = 10000000000000 // 1,000,000,000.0000
	for i = 1; i < 100000000; i++ {
		a++
		b = float64(a) / 10000
		c = New(b)
		if a != int64(c) {
			t.Fatal("New() not exact: should be:", a, "was:", c)
		}
	}
	t.Log("all possible values confirmed between 1,000,000,000 and ", c)

	var max float64 = float64(math.MaxInt64 / 1000000) // > 9.2 trillion
	max = max - 1 + .99
	x := New(max)
	t.Log("New max input:", max, " output:", x)
}

func TestNew6(t *testing.T) {
	minOk := Decimal6(20000000000000) // 2,000,000,000
	var a, i int64
	var b float64
	var c Decimal6
	for i = 1; i < math.MaxInt64; i++ {
		a += i
		b = float64(a) / 1000000
		c = NewDecimal6(b)
		if a != int64(c) {
			break
		}
	}
	if c < minOk {
		t.Fatal("did not reach minimum, largest value confirmed:", c)
	}
	t.Log(i, ", values checked, largest value confirmed: ", c)

	a = 10000000000000 // 1,000,000,000.0000
	for i = 1; i < 100000000; i++ {
		a++
		b = float64(a) / 1000000
		c = NewDecimal6(b)
		if a != int64(c) {
			t.Fatal("New() not exact: should be:", a, "was:", c)
		}
	}
	t.Log("all possible values confirmed between 1,000,000,000 and ", c)

	var max float64 = float64(math.MaxInt64 / 10000000) // > 922 billion
	max = max - 1 + .99
	x := NewDecimal6(max)
	t.Log("New6 max input:", max, " output:", x)
}

func TestMultiply(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 100000000) // > 92 billion
	var min float64 = float64(math.MinInt64 / 100000000) // < -92 billion
	t.Log("multiply max:", max, "min: ", min)
	data := []data{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{100000, .1, 10000},
		{100000, .0001, 10},
		{99999999.99, 1, 99999999.99},
		{555.5555, 333.3333, 185185.1481},
		{.5555, .5555, .3086},
		{.1111, .1111, .0123},
		{9999.9999, 9999.9999, 99999998.0000},
		{4500000, 20000, 90000000000}, // 90,000,000,000
		{-7321907.6324, -32.3976, 237212234.7114},
		{6111111.0039, 9388.0177, 57371218271.2780},
		{123.4567, 7654321.0025, 944977211.7093},
		{max, 1, max},
		{min, 1, min},
		//{max + 1, 1, max+1}, // should cause overflow panic
		//{min - 1, 1, min-1}, // should cause overflow panic (verified apr182016)
	}
	var a, b, c Decimal4
	for i, v := range data {
		a = New(v.a)
		b = New(v.b)
		c = a.Multiply(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestMultRound2(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 100000000) // > 92 billion
	var min float64 = float64(math.MinInt64 / 100000000) // < -92 billion
	t.Log("multiply max:", max, "min: ", min)
	data := []data{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{100000, .1, 10000},
		{100000, .0001, 10},
		{99999999.99, 1, 99999999.99},
		{555.5555, 333.3333, 185185.15},
		{.5555, .5555, .31},
		{.1111, .1111, .01},
		{9999.9999, 9999.9999, 99999998},
		{4500000, 20000, 90000000000}, // 90,000,000,000
		{-7321907.6324, -32.3976, 237212234.71},
		{8888.8888, -1, -8888.89},
		{456209.4178, 278.6814, 127137079.25},
		{100.75, 341, 34355.75},
		{max, 1, max},
		{min, 1, min},
		//{max + 1, 1, max+1}, // should cause overflow panic
		//{min - 1, 1, min-1}, // should cause overflow panic
	}
	var a, b, c Decimal4
	for i, v := range data {
		a = New(v.a)
		b = New(v.b)
		c = a.MultRound2(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestM(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 100000000) // > 92 billion
	var min float64 = float64(math.MinInt64 / 100000000) // < -92 billion
	t.Log("multiply max:", max, "min: ", min)
	data := []data{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{100000, .1, 10000},
		{100000, .0001, 10},
		{99999999.99, 1, 99999999.99},
		{555.5555, 333.3333, 185185.1481},
		{.5555, .5555, .3085},
		{.1111, .1111, .0123},
		{9999.9999, 9999.9999, 99999998.0000},
		{4500000, 20000, 90000000000}, // 90,000,000,000
		{-7321907.6324, -32.3976, 237212234.7114},
		{8888.8888, -1, -8888.8888},
		{123.4567, 7654321.0025, 944977211.7093},
		{max, 1, max},
		{min, 1, min},
		//{max + 1, 1, max+1}, // should cause overflow panic
		//{min - 1, 1, min-1}, // should cause overflow panic (verified apr182016)
	}
	var a, b, c Decimal4
	for i, v := range data {
		a = New(v.a)
		b = New(v.b)
		c = a.M(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestMultiplyBig(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 1000000) // > 9.2 trillion
	t.Log("multiplyBig max:", max)
	type input struct {
		a float64
		b float64
		c Decimal4
	}
	data := []input{
		{0, 0, 00000},
		{1, 1, 10000},
		{-1, 1, -10000},
		{100000000, .1, 100000000000},
		{999999999999.99, 1, 9999999999999900},
		{1000000000000.01, 1, 10000000000000100}, // 1 trillion
		{12345.67, 12345.67, 1524155677489},
		{-5000.0099, 2, -100000000}, // largest abs val should have 2 places truncated
		{83337222.0001, .0001, 83337222},
		{123999.4567, 7654321.0025, 9491316454074007},
		{803445821.1, 372.4701, 2992595453296991},
		{max, 1, 92233720368540000},
		// {max + 1, 1, 0}, // should cause overflow panic - verified 4/27/2016
	}
	var a, b, c Decimal4
	for i, v := range data {
		a = New(v.a)
		b = New(v.b)
		c = a.MultiplyBig(b)
		if c != v.c {
			t.Errorf("data[%d]: c should be %s, but is %s", i, v.c, c)
		}
	}
}

func TestMultiply6(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 10000000000) // > 900 million
	t.Log("multiply6 max:", max)
	data := []data{
		{0, 0, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{1, .123456, .1235},
		{555.5555, .000025, .0139},
		{1000000, .999999, 999999},
		{1000000, 10.000375, 10000375},
		{987654.4321, .987654, 975460.8505},
		{-33.3333, 73684.009999, -2456131.2105},
		//	{123.4567, 654321.002548, 949131645717.3993},		=====
		{max, 1, max},
		//{max + 1, 1, max + 1}, // should cause overflow panic, verified 4/18/2016
	}
	var a, c Decimal4
	var b Decimal6
	for i, v := range data {
		a = New(v.a)
		b = NewDecimal6(v.b)
		c = a.Multiply6(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestMultiplyBig6(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 100000000) // > 92 billion
	t.Log("multiplyBig6 max:", max)
	data := []data{
		{0, 0, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{1, .123456, .1235},
		{3849.27, .05125, 197.2751},
		{555.5555, .000025, .0139},
		{1000000, .999999, 999999},
		{1000000, 10.000375, 10000375},
		{987654.43, .987654, 975460.8484},
		{9000000000, -1.000001, -9000009000},
		{803445821.12, 72.470195, 58225875328.5015},
		{max, 1, max},
		//{max + 1, 1, 0}, // should cause overflow panic (verified 4/15/2016)
	}
	var a, c Decimal4
	var b Decimal6
	for i, v := range data {
		a = New(v.a)
		b = NewDecimal6(v.b)
		c = a.MultiplyBig6(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestMultiplyInt(t *testing.T) {
	type input struct {
		a float64
		b int
		c float64
	}
	data := []input{
		{0, 0, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{1.1111, 123456, 137171.9616},
		{9876543.1234, 599, 5916049330.9166},
		//{max + 1, 1, 0}, // should cause overflow panic (verified 4/15/2016)
	}
	var a, b Decimal4
	for i, v := range data {
		a = New(v.a)
		b = a.MultiplyInt(v.b)
		if b != New(v.c) {
			t.Errorf("data[%d]: b should be %f, but is %s", i, v.c, b)
		}
	}
}

func TestDivide(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 1000000000) // > 9 billion
	t.Log("divide max:", max)
	data := []data{
		{0, 1, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{100, .5, 200},
		{.05, 100, .0005},
		{9999.9999, 9, 1111.1111},
		{2222222.2222, 2222222.2222, 1},
		{1234567.0001, .123, 10037130.0821},
		{999, .4567, 2187.4316},
		{944977211.7093, 123.4567, 7654321.0025},
		{max, 1, max},
	}
	var a, b, c Decimal4
	for i, v := range data {
		a = New(v.a)
		b = New(v.b)
		c = a.Divide(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestDivideBig(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 10000000) // > 900 billion
	t.Log("divide max:", max)
	data := []data{
		{0, 1, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{9999.99, 9, 1111.11},
		{2222222.22, 2222222.22, 1},
		{1234567.01, .123, 10037130.162},
		{999, .4567, 2187.431},
		{2555444333, -9.125, -280048694.027},
		{max, 1, max},
		// {max + 1, 1, 0}, // should cause panic overflow, verified apr152016
	}
	var a, b, c Decimal4
	for i, v := range data {
		a = New(v.a)
		b = New(v.b)
		c = a.DivideBig(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestDivideInt(t *testing.T) {
	type input struct {
		a float64
		b int
		c float64
	}
	data := []input{
		{0, 1, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{100, 5, 20},
		{9999.9999, 9, 1111.1111},
		{1234567.0001, 123, 10037.1301},
		{999, -4567, -.2187},
	}
	var a, c Decimal4
	for i, v := range data {
		a = New(v.a)
		c = a.DivideInt(v.b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
		}
	}
}

func TestFmt(t *testing.T) {
	type input struct {
		val            Decimal4
		widthPrecision float64
		currency       string
		output         string
	}
	data := []input{
		{0, 1, "", "0"},
		{10000, 1, "", "1"},
		{11111, 5.2, "", " 1.11"},
		{-233987654, 13.4, Dollar, "$-23,398.7654"},
		{-7654, 13.4, Dollar, "     $-0.7654"},
		{1234567891239, 17.2, Dollar, "  $123,456,789.12"},
		{1234567891239, 16.3, Dollar, "$123,456,789.124"},
		{91234567891239, 18.4, "", "9,123,456,789.1239"},
		{9123456789551292, 20.3, "", " 912,345,678,955.129"},
		{12345600, 10.2, Dollar, " $1,234.56"},
		{12345600, .3, "", "1,234.560"},
	}
	for _, v := range data {
		result := v.val.Fmt(v.widthPrecision, v.currency)
		if result != v.output {
			t.Errorf("expected:%s   got:%s", v.output, result)
		}
	}
}

func TestRound(t *testing.T) {
	type input struct {
		input  float64
		places int
		output string
	}
	data := []input{
		{0, 0, "0.0000"},
		{1, 0, "1.0000"},
		{1.0001, 3, "1.0000"},
		{1.0009, 3, "1.0010"},
		{1.5, 0, "2.0000"},
		{99.9, 3, "99.9000"},
		{999.9999, 1, "1000.0000"},
		{999.9099, 1, "999.9000"},
		{999.9099, 2, "999.9100"},
		{999.9099, 3, "999.9100"},
		{10000000000.0006, 3, "10000000000.0010"},
	}
	var val, result Decimal4
	for _, v := range data {
		val = New(v.input)
		switch v.places {
		case 0:
			result = val.Round0()
		case 1:
			result = val.Round1()
		case 2:
			result = val.Round2()
		case 3:
			result = val.Round3()
		}
		if result.String() != v.output {
			t.Errorf("expected:%s   got:%s", v.output, result)
		}
	}
}

func TestTruncate(t *testing.T) {
	type input struct {
		input  float64
		places int
		output string
	}
	data := []input{
		{0, 0, "0.0000"},
		{1, 0, "1.0000"},
		{1.0001, 3, "1.0000"},
		{1.0009, 3, "1.0000"},
		{1.5, 0, "1.0000"},
		{99.9, 3, "99.9000"},
		{999.9999, 1, "999.9000"},
		{999.9099, 1, "999.9000"},
		{999.9099, 2, "999.9000"},
		{999.9099, 3, "999.9090"},
		{10000000001.1999, 3, "10000000001.1990"},
	}
	var val, result Decimal4
	for _, v := range data {
		val = New(v.input)
		switch v.places {
		case 0:
			result = val.Truncate0()
		case 1:
			result = val.Truncate1()
		case 2:
			result = val.Truncate2()
		case 3:
			result = val.Truncate3()
		}
		if result.String() != v.output {
			t.Errorf("expected:%s   got:%s", v.output, result)
		}
	}
}

func Benchmark_Multiply(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := New(123.4567)
		b := New(123.4567)
		for i := 0; i < 1000000; i++ { // 1 million iterations
			a.Multiply(b)
		}
	}
}

func Benchmark_M(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := New(123.4567)
		b := New(123.4567)
		for i := 0; i < 1000000; i++ { // 1 million iterations
			a.M(b)
		}
	}
}

func Benchmark_Divide(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := New(123.4567)
		b := New(123.4567)
		for i := 0; i < 1000000; i++ { // 1 million iterations
			a.Divide(b)
		}
	}
}
