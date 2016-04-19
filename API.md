#Decimal4 API

var Decimal4StringPlaces string = "4" // precision used by String method

type Decimal4 int64
type Decimal6 int64

##Decimal4 Methods

All computation methods return a Decimal4 value.
All computation methods will panic on overflow.

###Computation Methods  

.Multiply(x Decimal4) Decimal4  
    returns this * x, rounded to 4 places  
    
.MultiplyBig()
