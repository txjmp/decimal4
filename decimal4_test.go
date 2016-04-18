package decimal4

import "testing"
import "math"

type data struct {
	a float64 // input
	b float64 // input
	c float64 // result
}

func TestNew(t *testing.T) {
	minOk := New(200000000000)  // 200,000,000,000
	var a, i int64
	var b float64
	var c Decimal4
	for i = 1; i<math.MaxInt64 ;i++ {
		a += i
		b = float64(a) / 10000
		c = New(b)
		if a != int64(c) {
			break
		}
	}
	if c < minOk {
		t.Fatal("did not reach minimum, largest value confirmed:", c.Format(4))
	}
	t.Log(i, ", values checked, largest value confirmed: ", c.Format(4))

	a = 10000000000000  // 1,000,000,000.0000
	for i = 1; i<100000000; i++ {
		a++
		b = float64(a) / 10000
		c = New(b)
		if a != int64(c) {
			t.Fatal("New() not exact: should be:", a, "was:", c)
		}
	}
	t.Log("all possible values confirmed between 1,000,000,000 and ", c.Format(4))

	var max float64 = float64(math.MaxInt64 / 1000000) // > 9.2 trillion
	max = max - 1 + .99
	x := New(max)
	t.Log("New max input:", max, " output:", x.Format(4))
}

func TestMultiply(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 100000000) // > 92 billion
	t.Log("multiply max:", max)
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
		{9999.9999, 9999.9999, 99999998.0000 },
		{4500000, 20000, 90000000000},  // 90,000,000,000
		{max, 1, max},
		//{max + 1, 1, 0}, // should cause overflow panic
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

func TestMultiplyBig(t *testing.T) {
	var max float64 = float64(math.MaxInt64 / 1000000) // > 9.2 trillion
	t.Log("multiplyBig max:", max)
	data := []data{
		{0, 0, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{100000000, .1, 10000000},
		{999999999999.99, 1, 999999999999.99},
		{1000000000000.01, 1, 1000000000000.01}, // 1 trillion
		{12345.67, 12345.67, 152415567.7489},
		{83337222.0001, .0001, 8333.7222},
		{max, 1, max},
		//{max + 1, 1, 0}, // should cause overflow panic
	}
	var a, b, c Decimal4
	for i, v := range data {
		a = New(v.a)
		b = New(v.b)
		c = a.MultiplyBig(b)
		if c != New(v.c) {
			t.Errorf("data[%d]: c should be %f, but is %s", i, v.c, c)
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
		{max, 1, max},
		//{max + 1, 1, 0}, // should cause overflow panic
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
		{3849.27, .05125, 197.2751 },
		{555.5555, .000025, .0139},
		{1000000, .999999, 999999},
		{1000000, 10.000375, 10000375},
		{987654.43, .987654, 975460.8484},
		{9000000000, 1.000001, 9000009000},
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
		{2555444333, 9.125, 280048694.027},
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

func Benchmark_Multiply(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var c Decimal4
		a := New(123.4567)
		b := New(123.4567)
		for i := 0; i < 1000000; i++ { // 1 million iterations
			c = a.Multiply(b)
		}
		if c == 0 { // need ref to c for compile
			c = 0
		}
	}
}

func Benchmark_Divide(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var c Decimal4
		a := New(123.4567)
		b := New(123.4567)
		for i := 0; i < 1000000; i++ { // 1 million iterations
			c = a.Divide(b)
		}
		if c == 0 { // need ref to c for compile
			c = 0
		}
	}
}

func Benchmark_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var c Decimal4
		a := New(123.4567)
		b := New(123.4567)
		for i := 0; i < 1000000; i++ { // 1 million iterations
			c = a + b
		}
		if c == 0 { // need ref to c for compile
			c = 0
		}
	}
}
