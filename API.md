#Decimal4 API

var Decimal4StringPlaces string = "4" // precision used by String method

type Decimal4 int64
type Decimal6 int64

##Decimal4 Methods

####Computation Methods 

* all return a single Decimal4 value
* all panic on overflow 

.Multiply(x Decimal4)  
* returns this * x, rounded to 4 places  

.MultiplyBig(x Decimal4) - 
