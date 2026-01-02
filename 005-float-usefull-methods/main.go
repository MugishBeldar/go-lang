package main

import (
	"fmt"
	"math"
	"strconv"
)

const epsilon = 1e-9

func main() {
	// 1. float types in go
	fmt.Println("================================ FLOAT TYPE IN GO ====================================")
	// float32: 32-bit floating point (~6-7 decimal digit of precision)
	// float64: 64-bit floating point (~15-16 decimal digit of precision)
	// NOTE:
	// 	always use float64 unless you have a spicif reason for float32 (like a memory constraint)

	// 2. the math package
	fmt.Println("=============================== THE MATH PACKAGE ==============================")
	x := 3.7
	y := -1.2
	fmt.Println("x:", x, "y:", y)
	math.Abs(y)         // 1.2 absolute value
	math.Ceil(x)        // 4.0  round up
	math.Floor(x)       // 3.0 round down
	math.Round(x)       // 4.0 round to nearest
	math.Trunc(x)       // 3.0 truncate decimal part
	math.RoundToEven(x) // 4.0 banker's rounding
	math.Pow(2, 10)     // 1024.0  - 2¹⁰
	math.Pow10(3)       // 1000.0  - 10³
	math.Sqrt(16)       // 4.0     - Square root
	math.Cbrt(27)       // 3.0     - Cube root
	math.Exp(1)         // 2.718.. - e¹
	math.Exp2(3)        // 8.0     - 2³
	math.Hypot(3, 4)    // 5.0     - √(3² + 4²)
	math.Log(2.718)     // ~1.0    - Natural log (ln)
	math.Log10(100)     // 2.0     - Log base 10
	math.Log2(8)        // 3.0     - Log base 2
	math.Log1p(x)       // Log(1+x) - More accurate for small x
	math.Max(5.5, 10.2) // 10.2
	math.Min(5.5, 10.2) // 5.5
	math.IsNaN(x)       // Check if Not-a-Number
	math.IsInf(x, 1)    // Check if +Infinity
	math.IsInf(x, -1)   // Check if -Infinity
	math.IsInf(x, 0)    // Check if any Infinity

	// 3. the strconv package
	fmt.Println("=============================== THE STRCONV PACKAGE ==============================")
	z := 3.1415926
	// float to string
	str1 := strconv.FormatFloat(z, 'f', 2, 64) // 3.14
	fmt.Println(str1)
	str2 := strconv.FormatFloat(z, 'f', 4, 64) // 3.1416
	fmt.Println(str2)
	str3 := strconv.FormatFloat(z, 'f', 6, 64) // 3.141593
	fmt.Println(str3)

	// format specifiers:
	// 'f' - decimal point, no exponent (-ddd.ddd)
	// 'e' - scientific notation (d.ddde±dd)
	// 'E' - scientific notation (d.dddE±dd)
	// 'g' - 'e' for large exponents, 'f' otherwise
	// 'G' - 'E' for large exponents, 'f' otherwise

	// string to float
	str4 := "3.1415926"
	a, err := strconv.ParseFloat(str4, 64)
	if err != nil {
		fmt.Println("ParseFloat failed:", err)
	} else {
		fmt.Println(a)
	}

	// 4. float comparison
	fmt.Println("=============================== FLOAT COMPARISON ==============================")
	// bad - don't do this
	// a:=0.1 + 0.2
	// b:=0.3
	// fmt.Println(a == b) // false (due to precision issues)
	b := 3.1415926
	c := 3.1415926
	fmt.Println("direct float comparison false due to precision issues:", b == c) // false (due to precision issues)

	// good - use epsilon comparison
	fmt.Println("epsilon float comparison:", floatEquals(b, c)) // true

	// 5. nan behavior (weird but important)
	fmt.Println("=============================== NAN BEHAVIOR ==============================")
	nan := math.NaN()
	fmt.Println(nan == nan)      // false (NaN is not equal to anything, even itself)
	fmt.Println(nan > 0)         // false
	fmt.Println(nan < 0)         // false
	fmt.Println(nan == 0)        // false
	fmt.Println(nan != 0)        // true
	fmt.Println(math.IsNaN(nan)) // true

	// 6. infinity behavior (also weird but important)
	fmt.Println("=============================== INFINITY BEHAVIOR ==============================")
	inf := math.Inf(1)
	fmt.Println(inf == inf)          // true
	fmt.Println(inf > 0)             // true
	fmt.Println(inf < 0)             // false
	fmt.Println(inf == 0)            // false
	fmt.Println(inf != 0)            // true
	fmt.Println(math.IsInf(inf, 1))  // true
	fmt.Println(math.IsInf(inf, -1)) // false
	fmt.Println(math.IsInf(inf, 0))  // true
	fmt.Println(math.IsInf(inf, 2))  // false
	fmt.Println(math.IsInf(inf, -2)) // false
	fmt.Println(inf + 100)           // +Inf
	fmt.Println(inf - inf)           // NaN
	fmt.Println(inf * 0)             // NaN

	// below operation is not allowed thows an error division by 0 is not allowed
	// fmt.Println(1.0 / 0.0)           
	// fmt.Println(-1.0 / 0.0)          
	// fmt.Println(0.0 / 0.0)           

}

func floatEquals(a float64, b float64) bool {
	return math.Abs(a-b) < epsilon
}
