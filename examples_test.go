package decimal4 

import "testing"
import "fmt"

func TestFloatCompare(t *testing.T) {
	var f, fMax, fAdder float64 = 0, 100000, .0125
	for f < fMax {
		f += fAdder
	}
	fmt.Println(f)

	var d4, d4Max, d4Adder Decimal4 = 0, New(100000), New(.0125)
	for d4 < d4Max {
		d4 += d4Adder
	}
	fmt.Println(d4)
}

func Test2(t *testing.T) {
	fmt.Println("....... Test2 ..........")
	inputs := []float64{4000.1, 6000.01, 10000.001, 20000.0001, .002}
	d4Data := make([]Decimal4, len(inputs))
	for i, v := range inputs {
		d4Data[i] = New(v)
	}
	fmt.Println(d4Data)
	var tot, cnt Decimal4
	one := New(1)
	min := New(3)  // if value < min, skip it	
	for _, v := range d4Data {
		if v < min {
			continue
		}
		tot += v
		cnt += one
	}
	avg := tot.Divide(cnt)
	fmt.Printf("tot:%s  cnt:%s  avg:%s \n", tot.Format(4), cnt.Format(0), avg)
	if avg != New(10000.0278) {
		t.Fail()
	}
}	
func Test3(t *testing.T) {
	fmt.Println("....... Test3 ..........")
	type limitRate struct {
		limit  Decimal4
		rate Decimal6
	}    	
	// rates, as would typically be stored in database
	// **WARNING** DON'T USE LEADING ZEROS IN LITERALS
	rates := []limitRate{
		{1000000000, 31250},   // 100,000.0000 - 3 1/8%
		{2500000000, 40000},   // 250,000.0000 - 4 %
		{5000000000, 43750},   // 500,000.0000 - 4 3/8%
	}
	topRate := NewDecimal6(.05)
	for _, v := range rates{
		fmt.Println(v.limit.Format(0), v.rate)
	}
	// values, as would typically be stored in database
	values := []Decimal4{1000000000, 2500000001, 5999999999} //100,000, 250,000.0001, 599,999.9999 
	
	// compute charge = value * rate (rate determined by value)
	var charge, totCharges Decimal4
	var valueRate Decimal6
	for _, value := range values {
		valueRate = topRate
		for _, rate := range rates {
			if value <= rate.limit {
				valueRate = rate.rate
				break
			}
		}
		charge = value.Multiply6(valueRate).Round2()
		totCharges += charge
		fmt.Printf("value: %s,  rate: %s,  charge: %s \n", value, valueRate, charge)
	}
	fmt.Println("total charges", totCharges)
	//100,000 0.031250
	//250,000 0.040000
	//500,000 0.043750
	//value: 100000.0000,  rate: 0.031250,  charge: 3125.0000
	//value: 250000.0001,  rate: 0.043750,  charge: 10937.5000
	//value: 599999.9999,  rate: 0.050000,  charge: 30000.0000
	//total charges 44062.5000
}
