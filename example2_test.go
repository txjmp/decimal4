package decimal4

import "fmt"

// Determine a charge amt based on value.
// Higher values have different rates than lower values.
func Example2() {
	fmt.Println("....... Example2 ..........")
	type limitRate struct {
		limit Decimal4
		rate  Decimal6
	}
	// rates, as would typically be stored in database
	// **WARNING** DON'T USE LEADING ZEROS IN LITERALS
	rates := []limitRate{
		{1000000000, 31250}, // 100,000.0000 - 3 1/8%
		{2500000000, 40000}, // 250,000.0000 - 4 %
		{5000000000, 43750}, // 500,000.0000 - 4 3/8%
	}
	topRate := NewDecimal6(.05)
	for _, v := range rates {
		fmt.Println(v.limit.Fmt(.0), v.rate)
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
		fmt.Printf("value: %s,  rate: %s,  charge: %s\n", value.Fmt(12.4), valueRate, charge.Fmt(10.2, Dollar))
	}
	fmt.Println("total charges", totCharges)

	// Output:
	// ....... Example2 ..........
	// 100,000 0.031250
	// 250,000 0.040000
	// 500,000 0.043750
	// value: 100,000.0000,  rate: 0.031250,  charge:  $3,125.00
	// value: 250,000.0001,  rate: 0.043750,  charge: $10,937.50
	// value: 599,999.9999,  rate: 0.050000,  charge: $30,000.00
	// total charges 44062.5000
}
