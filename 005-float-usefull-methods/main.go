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

// ============================================================================
// INTERVIEW QUESTIONS - FLOAT OPERATIONS IN GO
// ============================================================================

// ============================================================================
// SECTION 1: BASIC LEVEL QUESTIONS
// ============================================================================

// Q1. What is the difference between float32 and float64?
/*
Answer:
float32:
- 32-bit floating point
- ~6-7 decimal digits of precision
- Range: ±1.18e-38 to ±3.4e38
- Uses less memory (4 bytes)

float64:
- 64-bit floating point
- ~15-16 decimal digits of precision
- Range: ±2.23e-308 to ±1.80e308
- Uses more memory (8 bytes)
- Default for floating-point literals in Go

Best practice: Use float64 unless memory is constrained
*/

// Q2. What is the output?
/*
func main() {
    x := 3.7
    fmt.Println(math.Ceil(x))
    fmt.Println(math.Floor(x))
    fmt.Println(math.Round(x))
}

Answer:
4
3
4
Explanation:
- Ceil rounds UP to next integer: 3.7 → 4
- Floor rounds DOWN to previous integer: 3.7 → 3
- Round rounds to nearest integer: 3.7 → 4
*/

// Q3. What is the output?
/*
func main() {
    fmt.Println(math.Abs(-5.5))
}

Answer: 5.5
Explanation: Abs returns absolute value (removes negative sign)
*/

// Q4. What is the output?
/*
func main() {
    fmt.Println(math.Sqrt(16))
    fmt.Println(math.Cbrt(27))
}

Answer:
4
3
Explanation:
- Sqrt(16) = √16 = 4
- Cbrt(27) = ∛27 = 3
*/

// Q5. What is the output?
/*
func main() {
    fmt.Println(math.Pow(2, 10))
    fmt.Println(math.Pow10(3))
}

Answer:
1024
1000
Explanation:
- Pow(2, 10) = 2^10 = 1024
- Pow10(3) = 10^3 = 1000
*/

// Q6. What is the output?
/*
func main() {
    fmt.Println(0.1 + 0.2 == 0.3)
}

Answer: false
Explanation: Floating-point precision issue!
0.1 + 0.2 = 0.30000000000000004 (not exactly 0.3)
This is due to binary representation of decimals
*/

// Q7. How should you compare floating-point numbers?
/*
Answer: Use epsilon comparison, not direct equality

Bad:
if a == b { ... }

Good:
const epsilon = 1e-9
func floatEquals(a, b float64) bool {
    return math.Abs(a - b) < epsilon
}

if floatEquals(a, b) { ... }

Explanation: Floating-point arithmetic has precision errors.
Use a small threshold (epsilon) to check if values are "close enough"
*/

// Q8. What is the output?
/*
func main() {
    str := strconv.FormatFloat(3.1415926, 'f', 2, 64)
    fmt.Println(str)
}

Answer: 3.14
Explanation: FormatFloat with precision 2 rounds to 2 decimal places
*/

// ============================================================================
// SECTION 2: INTERMEDIATE LEVEL QUESTIONS
// ============================================================================

// Q9. What is the output?
/*
func main() {
    fmt.Println(math.Trunc(3.7))
    fmt.Println(math.Trunc(-3.7))
}

Answer:
3
-3
Explanation: Trunc removes decimal part (truncates toward zero)
- 3.7 → 3
- -3.7 → -3
*/

// Q10. What is the difference between math.Round and math.RoundToEven?
/*
Answer:
math.Round:
- Standard rounding (round half away from zero)
- 2.5 → 3, 3.5 → 4, -2.5 → -3

math.RoundToEven (Banker's rounding):
- Round half to nearest even number
- 2.5 → 2 (even), 3.5 → 4 (even), 4.5 → 4 (even)
- Reduces cumulative rounding errors in financial calculations

Example:
fmt.Println(math.Round(2.5))      // 3
fmt.Println(math.RoundToEven(2.5)) // 2
*/

// Q11. What is the output?
/*
func main() {
    fmt.Println(math.Log(math.E))
    fmt.Println(math.Log10(100))
    fmt.Println(math.Log2(8))
}

Answer:
1
2
3
Explanation:
- Log(e) = ln(e) = 1
- Log10(100) = log₁₀(100) = 2 (because 10² = 100)
- Log2(8) = log₂(8) = 3 (because 2³ = 8)
*/

// Q12. What is the output?
/*
func main() {
    fmt.Println(math.Hypot(3, 4))
}

Answer: 5
Explanation: Hypot calculates hypotenuse: √(3² + 4²) = √(9 + 16) = √25 = 5
Useful for Pythagorean theorem calculations
*/

// Q13. What is the output?
/*
func main() {
    val, err := strconv.ParseFloat("3.14", 64)
    fmt.Println(val, err)
}

Answer: 3.14 <nil>
Explanation: ParseFloat successfully converts "3.14" to float64
*/

// Q14. What is the output?
/*
func main() {
    val, err := strconv.ParseFloat("abc", 64)
    fmt.Println(val, err)
}

Answer: 0 strconv.ParseFloat: parsing "abc": invalid syntax
Explanation: ParseFloat fails, returns 0 and error
*/

// Q15. What are the different format specifiers for FormatFloat?
/*
Answer:
'f': Fixed decimal point, no exponent
     Example: 123.456 → "123.456"

'e': Scientific notation (lowercase e)
     Example: 123.456 → "1.23456e+02"

'E': Scientific notation (uppercase E)
     Example: 123.456 → "1.23456E+02"

'g': Use 'e' for large exponents, 'f' otherwise (compact)
     Example: 0.0001 → "1e-04", 123.456 → "123.456"

'G': Use 'E' for large exponents, 'f' otherwise
     Example: 0.0001 → "1E-04", 123.456 → "123.456"

'b': Binary exponent (rarely used)
'x': Hexadecimal (rarely used)

Most common: 'f' for fixed decimal, 'e'/'E' for scientific notation
*/

// Q16. What is the output?
/*
func main() {
    fmt.Println(strconv.FormatFloat(1234.5678, 'f', 2, 64))
    fmt.Println(strconv.FormatFloat(1234.5678, 'e', 2, 64))
    fmt.Println(strconv.FormatFloat(1234.5678, 'g', 4, 64))
}

Answer:
1234.57
1.23e+03
1235
Explanation:
- 'f' with precision 2: 1234.5678 → 1234.57 (2 decimal places)
- 'e' with precision 2: 1234.5678 → 1.23e+03 (scientific notation)
- 'g' with precision 4: 1234.5678 → 1235 (4 significant digits)
*/

// Q17. What is the output?
/*
func main() {
    fmt.Println(math.Max(5.5, 10.2))
    fmt.Println(math.Min(5.5, 10.2))
}

Answer:
10.2
5.5
Explanation: Max returns larger value, Min returns smaller value
*/

// Q18. What is the output?
/*
func main() {
    fmt.Println(math.Exp(1))
    fmt.Println(math.Exp2(3))
}

Answer:
2.718281828459045
8
Explanation:
- Exp(1) = e^1 = 2.718... (Euler's number)
- Exp2(3) = 2^3 = 8
*/

// ============================================================================
// SECTION 3: ADVANCED LEVEL - NaN BEHAVIOR
// ============================================================================

// Q19. What is the output?
/*
func main() {
    nan := math.NaN()
    fmt.Println(nan == nan)
}

Answer: false
Explanation: NaN is NOT equal to anything, including itself!
This is IEEE 754 standard behavior.
Use math.IsNaN() to check for NaN.
*/

// Q20. What is the output?
/*
func main() {
    nan := math.NaN()
    fmt.Println(nan > 0)
    fmt.Println(nan < 0)
    fmt.Println(nan == 0)
}

Answer:
false
false
false
Explanation: NaN is not comparable to any value (always false)
*/

// Q21. How do you properly check if a value is NaN?
/*
Answer:
Bad:
if x == math.NaN() { ... }  // Always false!

Good:
if math.IsNaN(x) { ... }

Example:
nan := math.NaN()
fmt.Println(nan == math.NaN())  // false
fmt.Println(math.IsNaN(nan))    // true
*/

// Q22. What operations produce NaN?
/*
Answer:
Operations that produce NaN:
1. 0.0 / 0.0                    // Indeterminate form
2. math.Inf(1) - math.Inf(1)    // Infinity minus infinity
3. math.Inf(1) * 0              // Infinity times zero
4. math.Sqrt(-1)                // Square root of negative
5. math.Log(-1)                 // Log of negative
6. math.Asin(2)                 // Arcsin of value > 1

Example:
fmt.Println(math.Sqrt(-1))      // NaN
fmt.Println(0.0 / 0.0)          // NaN
fmt.Println(math.Inf(1) - math.Inf(1))  // NaN
*/

// Q23. What is the output?
/*
func main() {
    nan := math.NaN()
    fmt.Println(nan + 5)
    fmt.Println(nan * 2)
    fmt.Println(nan / 10)
}

Answer:
NaN
NaN
NaN
Explanation: Any arithmetic operation with NaN produces NaN
NaN "propagates" through calculations
*/

// ============================================================================
// SECTION 4: ADVANCED LEVEL - INFINITY BEHAVIOR
// ============================================================================

// Q24. What is the output?
/*
func main() {
    inf := math.Inf(1)
    fmt.Println(inf == inf)
}

Answer: true
Explanation: Unlike NaN, infinity IS equal to itself
*/

// Q25. What is the output?
/*
func main() {
    fmt.Println(1.0 / 0.0)
    fmt.Println(-1.0 / 0.0)
}

Answer:
+Inf
-Inf
Explanation: Division by zero for floats produces infinity (not panic!)
Positive divided by zero = +Inf
Negative divided by zero = -Inf
*/

// Q26. What is the output?
/*
func main() {
    fmt.Println(0.0 / 0.0)
}

Answer: NaN
Explanation: 0/0 is indeterminate form, produces NaN
*/

// Q27. What is the output?
/*
func main() {
    inf := math.Inf(1)
    fmt.Println(inf + 100)
    fmt.Println(inf - 100)
    fmt.Println(inf * 2)
    fmt.Println(inf / 2)
}

Answer:
+Inf
+Inf
+Inf
+Inf
Explanation: Arithmetic with infinity (except special cases) produces infinity
*/

// Q28. What is the output?
/*
func main() {
    inf := math.Inf(1)
    fmt.Println(inf - inf)
    fmt.Println(inf / inf)
    fmt.Println(inf * 0)
}

Answer:
NaN
NaN
NaN
Explanation: Indeterminate forms produce NaN:
- Infinity minus infinity
- Infinity divided by infinity
- Infinity times zero
*/

// Q29. How do you check for infinity?
/*
Answer:
math.IsInf(x, sign)

Parameters:
- sign = 1: Check for +Infinity only
- sign = -1: Check for -Infinity only
- sign = 0: Check for any infinity (+Inf or -Inf)

Example:
posInf := math.Inf(1)
negInf := math.Inf(-1)

fmt.Println(math.IsInf(posInf, 1))   // true
fmt.Println(math.IsInf(posInf, -1))  // false
fmt.Println(math.IsInf(posInf, 0))   // true
fmt.Println(math.IsInf(negInf, -1))  // true
fmt.Println(math.IsInf(negInf, 0))   // true
*/

// Q30. What is the output?
/*
func main() {
    inf := math.Inf(1)
    fmt.Println(math.IsInf(inf, 2))
    fmt.Println(math.IsInf(inf, -2))
}

Answer:
false
false
Explanation: IsInf only accepts -1, 0, or 1 as sign parameter
Any other value (like 2 or -2) returns false
*/

// ============================================================================
// SECTION 5: TYPE CONVERSION AND PRECISION
// ============================================================================

// Q31. What is the output?
/*
func main() {
    var x float32 = 1.23456789
    fmt.Println(x)
}

Answer: 1.2345679
Explanation: float32 has ~6-7 digits precision
Extra digits are rounded/truncated
*/

// Q32. What is the output?
/*
func main() {
    var x float64 = 1.234567890123456789
    fmt.Println(x)
}

Answer: 1.2345678901234567
Explanation: float64 has ~15-16 digits precision
Extra digits beyond that are lost
*/

// Q33. Will this code compile?
/*
func main() {
    var x float32 = 3.14
    var y float64 = x
    fmt.Println(y)
}

Answer: NO - Compilation Error
Error: "cannot use x (type float32) as type float64 in assignment"
Explanation: Go does not allow implicit type conversion
Use: var y float64 = float64(x)
*/

// Q34. What is the output?
/*
func main() {
    var x float64 = 3.7
    var y int = int(x)
    fmt.Println(y)
}

Answer: 3
Explanation: Converting float to int TRUNCATES (doesn't round)
3.7 → 3 (decimal part removed)
*/

// Q35. What is the output?
/*
func main() {
    var x float64 = 1e308
    var y float64 = x * 10
    fmt.Println(y)
    fmt.Println(math.IsInf(y, 1))
}

Answer:
+Inf
true
Explanation: Overflow in float64 produces infinity
1e308 * 10 exceeds max float64 value
*/

// Q36. What is the output?
/*
func main() {
    var x float64 = 1e-308
    var y float64 = x / 10
    fmt.Println(y)
}

Answer: 1e-309 (or very small number close to 0)
Explanation: Underflow in float64 produces very small number (denormalized)
Eventually becomes 0 if too small
*/

// ============================================================================
// SECTION 6: SPECIAL MATH FUNCTIONS
// ============================================================================

// Q37. What is the output?
/*
func main() {
    fmt.Println(math.Ceil(-3.2))
    fmt.Println(math.Floor(-3.8))
}

Answer:
-3
-4
Explanation:
- Ceil rounds UP (toward positive infinity): -3.2 → -3
- Floor rounds DOWN (toward negative infinity): -3.8 → -4
*/

// Q38. What is the output?
/*
func main() {
    fmt.Println(math.Trunc(3.9))
    fmt.Println(math.Trunc(-3.9))
}

Answer:
3
-3
Explanation: Trunc removes decimal part (truncates toward zero)
Different from Floor for negative numbers!
*/

// Q39. What is the difference between math.Trunc and math.Floor?
/*
Answer:
math.Trunc: Truncates toward zero
- 3.9 → 3
- -3.9 → -3

math.Floor: Rounds down (toward negative infinity)
- 3.9 → 3
- -3.9 → -4

They differ for negative numbers!

Example:
fmt.Println(math.Trunc(-3.9))  // -3
fmt.Println(math.Floor(-3.9))  // -4
*/

// Q40. What is the output?
/*
func main() {
    fmt.Println(math.Pow(2, 0.5))
    fmt.Println(math.Sqrt(2))
}

Answer:
1.4142135623730951
1.4142135623730951
Explanation: Both calculate √2
Pow(2, 0.5) = 2^0.5 = √2
Sqrt is optimized for square roots
*/

// Q41. What is the output?
/*
func main() {
    fmt.Println(math.Log1p(0))
    fmt.Println(math.Log(1))
}

Answer:
0
0
Explanation:
- Log1p(x) = Log(1 + x)
- Log1p(0) = Log(1 + 0) = Log(1) = 0
- Log1p is more accurate for small x values
*/

// Q42. What is the output?
/*
func main() {
    fmt.Println(math.Pow(0, 0))
    fmt.Println(math.Pow(-1, 0.5))
}

Answer:
1
NaN
Explanation:
- 0^0 = 1 (by convention in Go)
- (-1)^0.5 = √(-1) = NaN (square root of negative)
*/

// ============================================================================
// SECTION 7: PRACTICAL PROBLEMS
// ============================================================================

// Q43. How do you round a float to n decimal places?
/*
Answer:
func roundToDecimal(x float64, decimals int) float64 {
    shift := math.Pow(10, float64(decimals))
    return math.Round(x * shift) / shift
}

Example:
fmt.Println(roundToDecimal(3.14159, 2))  // 3.14
fmt.Println(roundToDecimal(3.14159, 3))  // 3.142

Explanation:
1. Multiply by 10^decimals to shift decimal point
2. Round to nearest integer
3. Divide by 10^decimals to shift back
*/

// Q44. How do you check if two floats are approximately equal?
/*
Answer:
Method 1: Absolute epsilon
func floatEquals(a, b, epsilon float64) bool {
    return math.Abs(a - b) < epsilon
}

Method 2: Relative epsilon (better for large numbers)
func floatEqualsRelative(a, b, epsilon float64) bool {
    diff := math.Abs(a - b)
    larger := math.Max(math.Abs(a), math.Abs(b))
    return diff <= larger * epsilon
}

Example:
const epsilon = 1e-9
fmt.Println(floatEquals(0.1 + 0.2, 0.3, epsilon))  // true
*/

// Q45. How do you safely divide two floats and handle division by zero?
/*
Answer:
func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    result := a / b
    if math.IsInf(result, 0) {
        return 0, fmt.Errorf("result is infinity")
    }
    if math.IsNaN(result) {
        return 0, fmt.Errorf("result is NaN")
    }
    return result, nil
}

Note: In Go, float division by zero doesn't panic, it returns Inf
But you might want to handle it as an error in your application
*/

// ============================================================================
// SECTION 8: ERROR IDENTIFICATION QUESTIONS
// ============================================================================

// Q46. Find the error in this code:
/*
func main() {
    var x float32 = 3.14
    var y float64 = 2.5
    result := x + y
    fmt.Println(result)
}

Answer: Compilation Error
Error: "invalid operation: x + y (mismatched types float32 and float64)"
Explanation: Cannot mix float32 and float64 without explicit conversion
Fix: result := float64(x) + y
*/

// Q47. Find the error in this code:
/*
func main() {
    x := 0.1 + 0.2
    if x == 0.3 {
        fmt.Println("Equal")
    } else {
        fmt.Println("Not equal")
    }
}

Answer: Logic Error (not compilation error)
Output: "Not equal"
Explanation: Floating-point precision issue
0.1 + 0.2 = 0.30000000000000004 (not exactly 0.3)
Fix: Use epsilon comparison
*/

// Q48. Find the error in this code:
/*
func main() {
    nan := math.NaN()
    if nan == math.NaN() {
        fmt.Println("Is NaN")
    }
}

Answer: Logic Error
Explanation: NaN is never equal to anything, including itself
This condition will never be true
Fix: if math.IsNaN(nan) { ... }
*/

// Q49. Find the error in this code:
/*
func main() {
    x := math.Sqrt(-1)
    if x < 0 {
        fmt.Println("Negative")
    }
}

Answer: Logic Error
Explanation: math.Sqrt(-1) returns NaN, not a negative number
NaN comparisons always return false
Fix: if math.IsNaN(x) { ... }
*/

// Q50. Find the error in this code:
/*
func main() {
    val, _ := strconv.ParseFloat("3.14.15", 64)
    fmt.Println(val)
}

Answer: Runtime Error (not panic, but logical error)
Output: 0
Explanation: ParseFloat fails on invalid format "3.14.15"
Returns 0 and error (which is ignored with _)
Fix: Check the error!
*/

// ============================================================================
// SECTION 9: TRICKY CODE EXECUTION QUESTIONS
// ============================================================================

// Q51. What is the output?
/*
func main() {
    x := 1.0
    for i := 0; i < 10; i++ {
        x += 0.1
    }
    fmt.Println(x == 2.0)
}

Answer: false
Explanation: Accumulated floating-point errors
1.0 + (0.1 * 10) ≠ 2.0 exactly due to binary representation
x will be approximately 2.0000000000000004
*/

// Q52. What is the output?
/*
func main() {
    fmt.Println(math.Round(2.5))
    fmt.Println(math.Round(3.5))
    fmt.Println(math.Round(4.5))
}

Answer:
2
4
4
Explanation: Go uses "round half to even" (banker's rounding)
- 2.5 → 2 (nearest even)
- 3.5 → 4 (nearest even)
- 4.5 → 4 (nearest even)
*/

// Q53. What is the output?
/*
func main() {
    x := math.Inf(1)
    y := math.Inf(1)
    fmt.Println(x == y)
    fmt.Println(x > y)
}

Answer:
true
false
Explanation: +Inf equals +Inf, but not greater than itself
*/

// Q54. What is the output?
/*
func main() {
    fmt.Println(math.Ceil(0.0))
    fmt.Println(math.Floor(0.0))
    fmt.Println(math.Round(0.0))
}

Answer:
0
0
0
Explanation: All rounding functions on 0.0 return 0
*/

// Q55. What is the output?
/*
func main() {
    x := -0.0
    y := 0.0
    fmt.Println(x == y)
    fmt.Println(1.0 / x)
    fmt.Println(1.0 / y)
}

Answer:
true
-Inf
+Inf
Explanation: -0.0 equals 0.0, but division produces different infinities!
This is IEEE 754 behavior (signed zero)
*/

// Q56. What is the output?
/*
func main() {
    fmt.Println(math.Max(math.NaN(), 5.0))
    fmt.Println(math.Min(math.NaN(), 5.0))
}

Answer:
NaN
NaN
Explanation: If either argument to Max/Min is NaN, result is NaN
*/

// Q57. What is the output?
/*
func main() {
    x := 1e-400
    fmt.Println(x)
    fmt.Println(x == 0)
}

Answer:
0
true
Explanation: 1e-400 underflows to 0 (below minimum float64)
*/

// Q58. What is the output?
/*
func main() {
    fmt.Println(strconv.FormatFloat(0.0001, 'f', -1, 64))
    fmt.Println(strconv.FormatFloat(0.0001, 'g', -1, 64))
}

Answer:
0.0001
0.0001
Explanation: Precision -1 means "use smallest number of digits"
*/

// Q59. What is the output?
/*
func main() {
    x := 10.0
    y := 3.0
    fmt.Println(x / y)
    fmt.Println(int(x) / int(y))
}

Answer:
3.3333333333333335
3
Explanation:
- Float division: 10.0 / 3.0 = 3.333...
- Integer division: 10 / 3 = 3 (truncated)
*/

// Q60. What is the output?
/*
func main() {
    fmt.Println(math.Pow(-1, 2))
    fmt.Println(math.Pow(-1, 2.0))
    fmt.Println(math.Pow(-1, 0.5))
}

Answer:
1
1
NaN
Explanation:
- (-1)^2 = 1
- (-1)^2.0 = 1
- (-1)^0.5 = √(-1) = NaN (complex number, not representable as float)
*/

// ============================================================================
// SECTION 10: CONCEPTUAL & BEST PRACTICES
// ============================================================================

// Q61. Why can't floating-point numbers represent 0.1 exactly?
/*
Answer:
Floating-point numbers use binary (base 2) representation.
Some decimal fractions cannot be represented exactly in binary.

Example:
0.1 in decimal = 0.0001100110011... (repeating) in binary

Just like 1/3 = 0.333... in decimal (can't be exact),
0.1 cannot be exact in binary.

This causes:
0.1 + 0.2 = 0.30000000000000004 (not exactly 0.3)

Solution: Use epsilon comparison for equality checks
*/

// Q62. When should you use float32 vs float64?
/*
Answer:
Use float64 (default):
- General-purpose floating-point calculations
- When precision is important
- Default for literals (x := 3.14 is float64)
- Most math functions work with float64

Use float32:
- Memory-constrained environments (large arrays)
- Graphics programming (GPU prefers float32)
- Interfacing with systems requiring 32-bit floats
- When ~6-7 digits precision is sufficient

Best practice: Prefer float64 unless you have specific reason for float32
*/

// Q63. What is the performance difference between float32 and float64?
/*
Answer:
Modern 64-bit CPUs:
- float64 operations are often SAME speed as float32
- CPU registers are 64-bit, so no performance penalty
- Some operations might be slightly faster with float64

Memory:
- float32 uses half the memory (4 bytes vs 8 bytes)
- Matters for large arrays/slices
- Better cache utilization with float32

Conclusion:
- Use float64 for precision (no performance cost on modern CPUs)
- Use float32 only when memory is constrained
*/

// Q64. How do you handle financial calculations in Go?
/*
Answer:
DON'T use float for money! Use integers or decimal libraries.

Bad:
var price float64 = 0.1
var total float64 = price * 10  // Might not be exactly 1.0

Good Option 1: Use cents (integers)
var priceInCents int = 10  // $0.10
var totalInCents int = priceInCents * 10  // $1.00 exactly

Good Option 2: Use decimal library
import "github.com/shopspring/decimal"
price := decimal.NewFromFloat(0.1)
total := price.Mul(decimal.NewFromInt(10))

Reason: Floating-point errors accumulate and cause financial discrepancies
*/

// Q65. What are denormalized numbers and when do they occur?
/*
Answer:
Denormalized (subnormal) numbers are very small floats near zero.

Normal float64:
- Smallest: ~2.2e-308
- Has implicit leading 1 in mantissa

Denormalized float64:
- Between 0 and ~2.2e-308
- No implicit leading 1
- Reduced precision
- Allows gradual underflow to zero

Example:
x := 1e-320  // Denormalized (very small)
y := x / 2   // Still denormalized
z := y / 1e100  // Eventually becomes 0

Performance: Denormalized operations can be slower on some CPUs

Check: No direct function, but if 0 < math.Abs(x) < math.SmallestNonzeroFloat64
*/

// Q66. Explain the difference between math.Round and math.RoundToEven with examples
/*
Answer:
math.Round (round half away from zero):
- 0.5 → 1
- 1.5 → 2
- 2.5 → 3
- -0.5 → -1
- -1.5 → -2

math.RoundToEven (banker's rounding):
- 0.5 → 0 (nearest even)
- 1.5 → 2 (nearest even)
- 2.5 → 2 (nearest even)
- 3.5 → 4 (nearest even)
- -0.5 → 0 (nearest even)

Why RoundToEven?
- Reduces cumulative bias in repeated rounding
- Used in financial calculations
- IEEE 754 recommended default

Example:
Sum of [0.5, 1.5, 2.5, 3.5] rounded:
- Round: 1 + 2 + 3 + 4 = 10 (bias: +2)
- RoundToEven: 0 + 2 + 2 + 4 = 8 (bias: 0)
*/

// Q67. What is the difference between math.Inf(1) and dividing by zero?
/*
Answer:
math.Inf(1):
- Explicitly creates +Infinity
- Controlled and intentional
- Returns float64

Division by zero:
- 1.0 / 0.0 → +Inf
- -1.0 / 0.0 → -Inf
- 0.0 / 0.0 → NaN
- Might be unintentional bug

Both produce same infinity value:
fmt.Println(math.Inf(1) == 1.0/0.0)  // true

Best practice: Use math.Inf() for clarity and intent
*/

// Q68. How do you detect and handle floating-point errors in calculations?
/*
Answer:
1. Check for NaN:
if math.IsNaN(result) {
    // Handle invalid operation
}

2. Check for Infinity:
if math.IsInf(result, 0) {
    // Handle overflow
}

3. Check for underflow:
if result == 0 && expected != 0 {
    // Possible underflow
}

4. Use epsilon comparison:
const epsilon = 1e-9
if math.Abs(result - expected) < epsilon {
    // Values are close enough
}

5. Validate input ranges:
if x < 0 {
    return 0, errors.New("cannot take sqrt of negative")
}
result := math.Sqrt(x)

Best practice: Always validate inputs and check outputs for special values
*/

// Q69. What is ULP (Unit in the Last Place) and why does it matter?
/*
Answer:
ULP is the spacing between consecutive floating-point numbers.

For float64:
- Near 1.0: ULP ≈ 2.22e-16
- Near 1000.0: ULP ≈ 2.27e-13
- Near 1e100: ULP ≈ 1.94e84

Implications:
1. Precision decreases for larger numbers
2. Cannot represent all numbers exactly
3. Rounding errors accumulate

Example:
x := 1.0
y := 1.0 + 1e-16  // Might equal 1.0 (below ULP)
fmt.Println(x == y)  // Might be true!

Best practice: Use relative epsilon based on magnitude
*/

// Q70. What are the common pitfalls when working with floats?
/*
Answer:
1. Direct equality comparison
   Bad: if x == 0.3 { ... }
   Good: if math.Abs(x - 0.3) < epsilon { ... }

2. Accumulation errors in loops
   Bad: sum := 0.0; for i := 0; i < 1000; i++ { sum += 0.1 }
   Good: Use integer arithmetic or decimal library

3. Comparing NaN
   Bad: if x == math.NaN() { ... }
   Good: if math.IsNaN(x) { ... }

4. Using floats for money
   Bad: var price float64 = 19.99
   Good: var priceInCents int = 1999

5. Ignoring overflow/underflow
   Bad: result := a * b
   Good: result := a * b; if math.IsInf(result, 0) { handle }

6. Assuming associativity
   Bad: (a + b) + c might not equal a + (b + c)
   Good: Be aware of order of operations

7. Converting to int without checking
   Bad: x := int(floatVal)
   Good: if floatVal > math.MaxInt64 { handle }
*/

// ============================================================================
// END OF INTERVIEW QUESTIONS
// ============================================================================
