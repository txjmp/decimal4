#Decimal4 API

var Decimal4StringPlaces string = "4" // precision used by String method

type Decimal4 int64
type Decimal6 int64

func New(x float64) Decimal4  
* returns x * 10000 rounded as int64

func NewDecimal6(x float64) Decimal6  
* returns x * 1000000 rounded as int64

##Decimal4 Methods

####Computation Methods 

* all return a single Decimal4 value
* all panic on overflow 
* Decimal4 values have implied 4 decimal places
* Decimal6 values have implied 6 decimal places

Multiply(x Decimal4)    
* returns this * x, rounded to 4 places  

MultiplyBig(x Decimal4)  
* returns this * x, rounded to 4 places
* the larger absolute value of this and x has last 2 decimal places truncated
* allows larger values than Multiply without causing overflow

Multiply6(x Decimal6)  
* returns this * x, rounded to 4 places
* x is type Decimal6, providing up to 6 decimal places precision
* allows smaller values than Multiply without causing overflow

Multiply6Big(x Decimal6)  
* returns this * x, rounded to 4 places
* x is type Decimal6, providing up to 6 decimal places precision
* last 2 decimal places of this value are truncated
* allows larger values than Multiply6 without causing overflow