#Decimal4 Package

This package provides decimal math using  integers. 

###Goals
1. Decimal accuracy
2. Ease of use
3. Precision & Max Values acceptable for a wide range of apps

###Key Points

* Values stored as int64: type Decimal4 int64.
* Values have 4 implied decimal places, for example: 1 stored as 10000.
* Computations can use one value with up to 6 decimal places, but result will only have 4.
* Values can easily be rounded to fewer decimal places.
* Values can be directly compared, added, and subtracted.
* Multiplication and division should use provided methods.
* Value limits (inputs & results) imposed by Decimal4 methods:
    * .Multiply - 92,233,720,368 (~92 billion)
    * .MultiplyBig - 9,223,372,036,854 (~9 trillion)
    * .Multiply6 - 922,337,036 (~900 million)
    * .Mutiply6Big - 92,233,720,368 (~92 billion)
    * .Divide - 9,223,372,036 (~9 billion)

Multiply and Divide methods will panic on overflow.

###RECOMMENDATION - always use variables, not literals or constants  

    y := New(1.1)  // stores as 11000
    x := y.Multiply(5)  // compiler converts 5 to Decimal4 type, but value is treated as .0005
    x += 5  // adds equivalent of .0005 to x
    five := 5
    x := y.Multiply(five) // will not compile
    z += five   // will not compile
    exception is zero, x = 0 is ok

###Comments on New() function:
  
Requires 1 parameter, a float64, and returns a Decimal4 (int64) value. Based on extensive testing, it returns an accurate value (rounded to 4 decimals) with the exception of very large values (testing indicates over 100 billion, but I have not tested every value). When rounded to 2 decimal places, even very large values will probably be correct (don't know the limit). If you need to create a new very large Decimal4 value and don't trust New() to be accurate to the required decimal places, type convert an int64 or literal value. Remember, the last 4 digits are for decimal places. For this example, use 13 zeros: billion := Decimal4(10000000000000).