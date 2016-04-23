package decimal4

import "fmt"

func Example1() {
	fmt.Println("....... Example1 ..........")
	inputs := []float64{500.0025, 200.0005, 299.997, 1.5}
	data := make([]Decimal4, len(inputs))
	for i, v := range inputs {
		data[i] = New(v)
	}
	minValue := New(3)  // if data value below minValue, ignore it
	var total Decimal4
	var count int
	for _, v := range data {
		if v < minValue {
			continue
		}
		total += v
		count++
	}
	average := total.DivideInt(count)
	fmt.Printf("total: %s  average: %s \n", total.Fmt(.4), average)

	// Output:
	// ....... Example1 ..........
	// total: 1,000.0000  average: 333.3333
}
