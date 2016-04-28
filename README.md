#Go Decimal4 Package

This package provides decimal math using integers. 
  
To Install: go get github.com/txjmp/decimal4  
  
###Status

Passing all [tests](https://github.com/txjmp/decimal4/blob/master/decimal4_test.go). Need feedback on design, usability, features, and bugs.  

###Goals
1. Decimal accuracy
2. Ease of use
3. Precision & Max Values acceptable for a wide range of apps

###Key Points

* Values stored as int64: type Decimal4 int64.
* Values have 4 implied decimal places, for example: 1 stored as 10000.
* Computations can use one value with up to 6 decimal places, but result will only have 4.
* Values can easily be rounded or truncated to fewer decimal places.
* Values can be directly compared, added, and subtracted.
* Multiplication and division calculations should use provided methods.
* Value limits (inputs & results) imposed by Decimal4 methods:
    * .Multiply - 92,233,720,368 (~ -92 to +92 billion)
    * .MultiplyBig - 9,223,372,036,854 (~ -9 to +9 trillion)
    * .Multiply6 - 922,337,036 (~ -900 to +900 million)
    * .MutiplyBig6 - 92,233,720,368 (~ -92 to +92 billion)
    * .Divide - 9,223,372,036 (~ -9 to +9 billion)
    * .DivideBig - 922,337,203,685 (~ -922 to +922 billion)
* Multiply and Divide methods will panic on overflow.
* Values can be formatted with commas and currency sign.

###RECOMMENDATION - always use variables, not literals or constants

    y := New(1.1)  // stores as 11000
    x := y.Multiply(5)  // compiler converts 5 to Decimal4 type, but value is treated as .0005
    x += 5  // adds equivalent of .0005 to x
    five := 5
    x := y.Multiply(five) // will not compile
    z += five   // will not compile
    exception is zero, x = 0 is ok

###Comments on New() function:
  
Requires 1 parameter, a float64, and returns a Decimal4 (int64) value. Based on extensive testing, it returns an accurate value (rounded to 4 decimals) with the exception of very large values (testing indicates over 100 billion, but I have not tested every value). When rounded to 2 decimal places, even very large values will probably be correct (don't know limits/exceptions). If you need to create a new very large Decimal4 value and don't trust New() to be accurate to the required decimal places, type convert an int64 or literal value. Remember, the last 4 digits are for decimal places. Example:  

    billion := Decimal4(10000000000000)     // 13 zeros  
 
##[Link To API](https://github.com/txjmp/decimal4/blob/master/API.md)    

##Example  
  
    package main

    import "fmt"
    import d4 "github.com/txjmp/decimal4"

    func main() {
        inputs := []float64{500.0025, 200.0005, 299.997}
  
        data := make([]d4.Decimal4, len(inputs))
        for i, v := range inputs {
            data[i] = d4.New(v)
        }
        var total d4.Decimal4
        for _, v := range data {
            total += v
        }
        average := total.DivideInt(len(data))

        fmt.Printf("total: %s  average: %s", total.Fmt(.4), average)

        // Output:
        // total: 1,000.0000  average: 333.3333
    }
  
##Roll Your Own Decimal Module  

Methods in decimal4.go have few parameters. This design reduces method code size, complexity, and need for returning errors. It also limits the functionality of each method. You may need a multiply method that truncates to 3 decimal places. The desired result can be accomplished by chaining two methods together, but performance won't be so great. Creating your own decimal module might be the answer. Use decimal4.go for examples to speed up the process.