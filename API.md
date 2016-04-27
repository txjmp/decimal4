#Decimal4 API

### Global Public Variables & Constants

var Decimal4StringPlaces string = "4" // precision used by String method  
  
const Dollar = "\u0024"  
const Euro = "\u20AC"  
const Yen = "\u00A5"  
const Yuan = "\u00A5"  
const Rupee = "\u20B9"  
const Ruble = "\u20BD"  
const Pound = "\u00A3"  
  
### Types  

type Decimal4 int64  
type Decimal6 int64  

### Functions  

func New(x float64) Decimal4  
* returns x * 10000 rounded and converted to Decimal4

func NewDecimal6(x float64) Decimal6  
* returns x * 1000000 rounded and converted to Decimal6

###Decimal4 Computation Methods 

* all return a single Decimal4 value
* all panic on overflow (except M)
* Decimal4 values have implied 4 decimal places
* Decimal6 values have implied 6 decimal places
* receiver used by all: (this Decimal4)

Multiply(x Decimal4)    
* returns *this* * x, rounded to 4 places  

MultRound2(x Decimal4)    
* returns *this* * x, rounded to 2 places  

M(x Decimal4)    
* a fast version of Multiply
* no rounding (result truncated after 4th decimal), no check for overflow  

MultiplyBig(x Decimal4)  
* returns *this* * x, rounded to 4 places
* the larger absolute value between *this* and x has last 2 decimal places truncated
* allows larger values than Multiply without causing overflow

Multiply6(x Decimal6)  
* returns *this* * x, rounded to 4 places
* x is type Decimal6, providing up to 6 decimal places precision
* allows smaller values than Multiply without causing overflow

MultiplyBig6(x Decimal6)  
* returns *this* * x, rounded to 4 places
* x is type Decimal6, providing up to 6 decimal places precision
* last 2 decimal places of *this* value are truncated
* allows larger values than Multiply6 without causing overflow

MultiplyInt(x int)  
* returns *this* * x

Divide(x Decimal4)  
* returns *this* / x, rounded to 4 places

DivideBig(x Decimal4)  
* returns *this* / x, 3 decimal places precision, no rounding
* allows larger values than Divide without causing overflow

DivideInt(x int)  
* returns *this* / x, rounded to 4 places

---

####Decimal4 Output Methods 

Round0(), Round1(), Round2(), Round3()
* no parameters
* each rounds value of *this* to specified decimal places and returns Decimal4

Truncate0(), Truncate1(), Truncate2(), Truncate3()
* no parameters
* each truncates value of *this* to specified decimal places and returns Decimal4

Fmt(widthPrecision float64, currency ...string) string
* returns *this* formatted with width.precision and comma thousands separators
* if optional currency is specified, output will have symbol prefixed to value
* see included currency examples (Dollar, Euro, Yen, ...) at top
  
Fmt Examples:  
  
    d4Val := Decimal4(12345600)  
    d4Val.Fmt(10.2, Dollar) -> " $1,234.56"   
    d4Val.Fmt(.3) -> "1,234.560"   
  
String() string
* converts *this* to float64
* returns output of fmt.Sprintf using Decimal4StringPlaces to set decimal places shown

