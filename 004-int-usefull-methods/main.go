package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// unlike lnaguages like paython or java, go's primitive types like int don't have methods.
	// instead go probides.
	// operators for basic operations
	// standard library packages for advanced operations

	// 1. integer types in go
	// Signed integers
	// int8    // -128 to 127
	// int16   // -32,768 to 32,767
	// int32   // -2,147,483,648 to 2,147,483,647
	// int64   // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	// int     // Platform dependent (32 or 64 bit)

	// // Unsigned integers
	// uint8   // 0 to 255 (alias: byte)
	// uint16  // 0 to 65,535
	// uint32  // 0 to 4,294,967,295
	// uint64  // 0 to 18,446,744,073,709,551,615
	// uint    // Platform dependent

	fmt.Println("============================ ARITHMETIC OPERATORS ===============================")

	// 2. arithmetic operators
	a := 10
	b := 3

	sum := a + b // 13
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	diff := a - b // 7
	fmt.Printf("%d - %d = %d\n", a, b, diff)

	prod := a * b // 30
	fmt.Printf("%d * %d = %d\n", a, b, prod)

	quot := a / b // 3 (integer division)
	fmt.Printf("%d / %d = %d\n", a, b, quot)

	rem := a % b // 1
	fmt.Printf("%d %% %d = %d\n", a, b, rem)

	// Increment / Decrement
	a++
	fmt.Printf("a++ = %d\n", a)

	b--
	fmt.Printf("b-- = %d\n", b)

	fmt.Println("============================ BITWISE OPERATORS ===============================")

	// Declare two integer numbers
	// num1 in binary  = 0101 (5)
	// num2 in binary  = 0011 (3)
	num1 := 5
	num2 := 3

	// AND (&)
	// Each bit is 1 only if both bits are 1
	// 0101
	// 0011
	// ----
	// 0001  -> 1
	andResult := num1 & num2
	fmt.Printf("%d & %d = %d\n", num1, num2, andResult)

	// OR (|)
	// Each bit is 1 if at least one bit is 1
	// 0101
	// 0011
	// ----
	// 0111  -> 7
	orResult := num1 | num2
	fmt.Printf("%d | %d = %d\n", num1, num2, orResult)

	// XOR (^)
	// Each bit is 1 if bits are different
	// 0101
	// 0011
	// ----
	// 0110  -> 6
	xorResult := num1 ^ num2
	fmt.Printf("%d ^ %d = %d\n", num1, num2, xorResult)

	// AND NOT (&^)  (Bit Clear - Go specific)
	// Clears bits from num1 that are set in num2
	// 0101
	// 0011
	// ----
	// 0100  -> 4
	andNotResult := num1 &^ num2
	fmt.Printf("%d &^ %d = %d\n", num1, num2, andNotResult)

	// Left Shift (<<)
	// Shifts bits to the left by 2 positions
	// Equivalent to multiplying by 2^2
	// 0101 << 2 = 10100 -> 20
	leftShiftResult := num1 << 2
	fmt.Printf("%d << 2 = %d\n", num1, leftShiftResult)

	// Right Shift (>>)
	// Shifts bits to the right by 1 position
	// Equivalent to dividing by 2^1
	// 0101 >> 1 = 0010 -> 2
	rightShiftResult := num1 >> 1
	fmt.Printf("%d >> 1 = %d\n", num1, rightShiftResult)

	fmt.Println("============================ THE MATH PACKAGE ===============================")
	// 4. the math package

	// -------------------- Constants --------------------

	// Maximum value an int64 can store
	fmt.Println("MaxInt64:", math.MaxInt64)

	// Minimum value an int64 can store
	fmt.Println("MinInt64:", math.MinInt64)

	// Maximum value of int (depends on system architecture)
	fmt.Println("MaxInt:", math.MaxInt)

	// -------------------- Common Math Functions --------------------
	// NOTE: Most math functions work with float64 only

	// Absolute value
	// Converts negative number to positive
	absValue := math.Abs(-5)
	fmt.Println("Abs(-5) =", absValue)

	// Power function
	// 2 raised to the power 10
	powerValue := math.Pow(2, 10)
	fmt.Println("Pow(2, 10) =", powerValue)

	// Square root
	// Root of 16
	sqrtValue := math.Sqrt(16)
	fmt.Println("Sqrt(16) =", sqrtValue)

	// Maximum of two numbers
	maxValue := math.Max(5, 10)
	fmt.Println("Max(5, 10) =", maxValue)

	// Minimum of two numbers
	minValue := math.Min(5, 10)
	fmt.Println("Min(5, 10) =", minValue)

	// Ceil (round up)
	// Always rounds to the next highest integer
	ceilValue := math.Ceil(4.2)
	fmt.Println("Ceil(4.2) =", ceilValue)

	// Floor (round down)
	// Always rounds to the next lowest integer
	floorValue := math.Floor(4.8)
	fmt.Println("Floor(4.8) =", floorValue)

	// Round
	// Rounds to the nearest integer
	roundValue := math.Round(4.5)
	fmt.Println("Round(4.5) =", roundValue)

	fmt.Println("============================ STRCONV PACKAGE ===============================")
	// 5. the strconv package
	// -------------------- Int to String --------------------

	// Integer value
	number := 42

	// Convert int to string
	// Itoa = Integer to ASCII
	stringValue := strconv.Itoa(number)
	fmt.Println("Itoa:", stringValue)

	// Convert int64 to string with base 10 (decimal)
	stringBase10 := strconv.FormatInt(42, 10)
	fmt.Println("FormatInt (base 10):", stringBase10)

	// Convert int64 to string with base 2 (binary)
	stringBase2 := strconv.FormatInt(42, 2)
	fmt.Println("FormatInt (base 2):", stringBase2)

	// Convert int64 to string with base 16 (hexadecimal)
	stringBase16 := strconv.FormatInt(42, 16)
	fmt.Println("FormatInt (base 16):", stringBase16)

	// -------------------- String to Int --------------------

	// Convert string to int
	// Returns converted value and error
	intValue, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Atoi failed:", err)
	} else {
		fmt.Println("Atoi:", intValue)
	}

	// Convert string to int64 with base 10 and 64-bit size
	int64Value, err := strconv.ParseInt("42", 10, 64)
	if err != nil {
		fmt.Println("ParseInt failed:", err)
	} else {
		fmt.Println("ParseInt:", int64Value)
	}

	// 	Important Notes (Must Remember)
	// Itoa works with int only
	// FormatInt works with int64
	// ParseInt returns int64
	// Always handle error when converting string → number

	fmt.Println("============================ TYPE CONVERSION ===============================")
	// 7. type conversation
	// -------------------- Integer Type Conversion --------------------

	// Declare a 32-bit integer
	var smallInt32 int32 = 100

	// Explicitly convert int32 to int64
	// Go does NOT allow implicit type conversion
	// difference between implicit and explicit
	// implicit: implicit referes to action or types handled automatically by compiler or system
	// explicit: explicit referes to action or types handled manually by programmer for controll, cliarity, and early error detection
	var largeInt64 int64 = int64(smallInt32)
	fmt.Println("int32 to int64:", largeInt64)

	// -------------------- Integer to Float Conversion --------------------

	// Declare an integer
	var intVal int = 42

	// Convert int to float64
	var floatValue float64 = float64(intVal)
	fmt.Println("int to float64:", floatValue)

	// Convert float64 to int
	// Decimal part is truncated (not rounded)
	var truncatedInt int = int(floatValue)
	fmt.Println("float64 to int:", truncatedInt)

	// -------------------- Overflow Example --------------------

	// int64 value larger than int8 range (-128 to 127)
	var bigInt64 int64 = 300

	// Explicit conversion causes overflow
	// 300 % 256 = 44
	var overflowedInt8 int8 = int8(bigInt64)
	fmt.Println("Overflow example (int64 to int8):", overflowedInt8)

	// 	Important Notes (Must Remember)
	// Go does NOT allow implicit type conversion
	// Explicit conversion may cause overflow

}

// ============================================================================
// INTERVIEW QUESTIONS - INTEGER OPERATIONS IN GO
// ============================================================================

// ============================================================================
// SECTION 1: BASIC LEVEL QUESTIONS
// ============================================================================

// Q1. What is the output?
/*
func main() {
    a := 10
    b := 3
    fmt.Println(a / b)
}

Answer: 3
Explanation: Integer division truncates the decimal part. 10/3 = 3.333... → 3
*/

// Q2. What is the output?
/*
func main() {
    a := 10
    b := 3
    fmt.Println(a % b)
}

Answer: 1
Explanation: Modulo operator returns remainder. 10 divided by 3 leaves remainder 1
*/

// Q3. What is the range of int8?
/*
Answer: -128 to 127
Explanation: int8 is 8-bit signed integer
- 1 bit for sign, 7 bits for value
- Range: -2^7 to 2^7-1 = -128 to 127
*/

// Q4. What is the range of uint8?
/*
Answer: 0 to 255
Explanation: uint8 is 8-bit unsigned integer (no negative values)
- All 8 bits for value
- Range: 0 to 2^8-1 = 0 to 255
*/

// Q5. What is the output?
/*
func main() {
    x := 5
    x++
    fmt.Println(x)
}

Answer: 6
Explanation: x++ increments x by 1
*/

// Q6. Will this code compile?
/*
func main() {
    x := 5
    y := ++x
    fmt.Println(y)
}

Answer: NO - Compilation Error
Error: "syntax error: unexpected ++, expecting expression"
Explanation: Go does not support prefix increment (++x). Only postfix (x++) is allowed.
Also, x++ is a statement, not an expression, so cannot be assigned.
*/

// Q7. What is the output?
/*
func main() {
    fmt.Println(strconv.Itoa(42))
}

Answer: 42
Explanation: Itoa converts int to string. Output is string "42"
*/

// Q8. What is the output?
/*
func main() {
    val, err := strconv.Atoi("42")
    fmt.Println(val, err)
}

Answer: 42 <nil>
Explanation: Atoi successfully converts "42" to int 42, error is nil
*/

// ============================================================================
// SECTION 2: INTERMEDIATE LEVEL QUESTIONS
// ============================================================================

// Q9. What is the output?
/*
func main() {
    val, err := strconv.Atoi("abc")
    fmt.Println(val, err)
}

Answer: 0 strconv.Atoi: parsing "abc": invalid syntax
Explanation: Atoi fails to convert "abc" to int, returns 0 and error
*/

// Q10. What is the output?
/*
func main() {
    a := 5  // binary: 0101
    b := 3  // binary: 0011
    fmt.Println(a & b)
}

Answer: 1
Explanation: Bitwise AND
0101 & 0011 = 0001 = 1
*/

// Q11. What is the output?
/*
func main() {
    a := 5  // binary: 0101
    b := 3  // binary: 0011
    fmt.Println(a | b)
}

Answer: 7
Explanation: Bitwise OR
0101 | 0011 = 0111 = 7
*/

// Q12. What is the output?
/*
func main() {
    a := 5  // binary: 0101
    b := 3  // binary: 0011
    fmt.Println(a ^ b)
}

Answer: 6
Explanation: Bitwise XOR (exclusive OR)
0101 ^ 0011 = 0110 = 6
*/

// Q13. What is the output?
/*
func main() {
    a := 5  // binary: 0101
    fmt.Println(a << 2)
}

Answer: 20
Explanation: Left shift by 2 positions
0101 << 2 = 10100 = 20
Equivalent to: 5 * 2^2 = 5 * 4 = 20
*/

// Q14. What is the output?
/*
func main() {
    a := 5  // binary: 0101
    fmt.Println(a >> 1)
}

Answer: 2
Explanation: Right shift by 1 position
0101 >> 1 = 0010 = 2
Equivalent to: 5 / 2^1 = 5 / 2 = 2 (integer division)
*/

// Q15. What is the output?
/*
func main() {
    a := 5  // binary: 0101
    b := 3  // binary: 0011
    fmt.Println(a &^ b)
}

Answer: 4
Explanation: AND NOT (bit clear) - Go specific operator
0101 &^ 0011 = 0100 = 4
Clears bits in 'a' that are set in 'b'
*/

// Q16. What is the output?
/*
func main() {
    fmt.Println(math.Abs(-5))
}

Answer: 5
Explanation: Abs returns absolute value (removes negative sign)
*/

// Q17. What is the output?
/*
func main() {
    fmt.Println(math.Pow(2, 10))
}

Answer: 1024
Explanation: Pow(2, 10) = 2^10 = 1024
*/

// Q18. What is the output?
/*
func main() {
    fmt.Println(math.Ceil(4.2))
    fmt.Println(math.Floor(4.8))
    fmt.Println(math.Round(4.5))
}

Answer:
5
4
5
Explanation:
- Ceil rounds UP to next integer: 4.2 → 5
- Floor rounds DOWN to previous integer: 4.8 → 4
- Round rounds to nearest integer: 4.5 → 5
*/

// ============================================================================
// SECTION 3: ADVANCED LEVEL QUESTIONS
// ============================================================================

// Q19. What is the output?
/*
func main() {
    var x int8 = 127
    x = x + 1
    fmt.Println(x)
}

Answer: -128
Explanation: Integer overflow! int8 range is -128 to 127.
127 + 1 wraps around to -128 (no error in Go)
*/

// Q20. What is the output?
/*
func main() {
    var x uint8 = 255
    x = x + 1
    fmt.Println(x)
}

Answer: 0
Explanation: Unsigned integer overflow! uint8 range is 0 to 255.
255 + 1 wraps around to 0
*/

// Q21. What is the output?
/*
func main() {
    var x uint8 = 0
    x = x - 1
    fmt.Println(x)
}

Answer: 255
Explanation: Unsigned integer underflow! 0 - 1 wraps around to 255
*/

// Q22. Will this code compile?
/*
func main() {
    var x int32 = 100
    var y int64 = x
    fmt.Println(y)
}

Answer: NO - Compilation Error
Error: "cannot use x (type int32) as type int64 in assignment"
Explanation: Go does not allow implicit type conversion, even between integer types.
Use: var y int64 = int64(x)
*/

// Q23. What is the output?
/*
func main() {
    var x int64 = 300
    var y int8 = int8(x)
    fmt.Println(y)
}

Answer: 44
Explanation: Overflow during conversion!
300 is outside int8 range (-128 to 127)
300 % 256 = 44 (wraps around)
*/

// Q24. What is the output?
/*
func main() {
    fmt.Println(strconv.FormatInt(42, 2))
    fmt.Println(strconv.FormatInt(42, 10))
    fmt.Println(strconv.FormatInt(42, 16))
}

Answer:
101010
42
2a
Explanation:
- Base 2 (binary): 42 = 101010
- Base 10 (decimal): 42 = 42
- Base 16 (hexadecimal): 42 = 2a
*/

// Q25. What is the output?
/*
func main() {
    val, _ := strconv.ParseInt("101010", 2, 64)
    fmt.Println(val)
}

Answer: 42
Explanation: ParseInt converts binary string "101010" to decimal int64
101010 (base 2) = 42 (base 10)
*/

// Q26. Will this code compile?
/*
func main() {
    x := 10
    y := 3.5
    fmt.Println(x + y)
}

Answer: NO - Compilation Error
Error: "invalid operation: x + y (mismatched types int and float64)"
Explanation: Cannot mix int and float64 without explicit conversion
Use: fmt.Println(float64(x) + y)
*/

// Q27. What is the output?
/*
func main() {
    x := 10.9
    fmt.Println(int(x))
}

Answer: 10
Explanation: Converting float to int TRUNCATES (does not round)
10.9 → 10 (decimal part removed)
*/

// Q28. What is the output?
/*
func main() {
    fmt.Println(-10 % 3)
    fmt.Println(10 % -3)
    fmt.Println(-10 % -3)
}

Answer:
-1
1
-1
Explanation: Sign of result follows sign of dividend (first operand)
-10 % 3 = -1 (dividend is negative)
10 % -3 = 1 (dividend is positive)
-10 % -3 = -1 (dividend is negative)
*/

// Q29. What is the output?
/*
func main() {
    fmt.Println(math.Max(5, 10))
    fmt.Println(math.Min(5, 10))
}

Answer:
10
5
Explanation: Max returns larger value, Min returns smaller value
Note: These functions work with float64, not int
*/

// Q30. What is the output?
/*
func main() {
    x := 5
    y := 10
    max := math.Max(x, y)
    fmt.Println(max)
}

Answer: Compilation Error
Error: "cannot use x (type int) as type float64 in argument to math.Max"
Explanation: math.Max expects float64, not int
Use: max := math.Max(float64(x), float64(y))
*/

// ============================================================================
// SECTION 4: BITWISE OPERATIONS - TRICKY QUESTIONS
// ============================================================================

// Q31. What is the output?
/*
func main() {
    x := 8  // binary: 1000
    fmt.Println(x << 1)
    fmt.Println(x >> 1)
}

Answer:
16
4
Explanation:
- Left shift: 1000 << 1 = 10000 = 16 (multiply by 2)
- Right shift: 1000 >> 1 = 0100 = 4 (divide by 2)
*/

// Q32. What is the output?
/*
func main() {
    x := 15  // binary: 1111
    y := 8   // binary: 1000
    fmt.Println(x &^ y)
}

Answer: 7
Explanation: AND NOT (bit clear)
1111 &^ 1000 = 0111 = 7
Clears bits in x that are set in y
*/

// Q33. How can you check if a number is even using bitwise operations?
/*
Answer:
func isEven(n int) bool {
    return n & 1 == 0
}

Explanation:
- If last bit is 0, number is even
- If last bit is 1, number is odd
- n & 1 extracts the last bit

Examples:
4 (0100) & 1 (0001) = 0 → even
5 (0101) & 1 (0001) = 1 → odd
*/

// Q34. How can you swap two numbers without a temporary variable using XOR?
/*
Answer:
func swap(a, b int) (int, int) {
    a = a ^ b
    b = a ^ b
    a = a ^ b
    return a, b
}

Explanation:
Let a = 5 (0101), b = 3 (0011)
Step 1: a = a ^ b = 0101 ^ 0011 = 0110 (6)
Step 2: b = a ^ b = 0110 ^ 0011 = 0101 (5) → b now has original a
Step 3: a = a ^ b = 0110 ^ 0101 = 0011 (3) → a now has original b

Note: In Go, prefer: a, b = b, a (cleaner and more efficient)
*/

// Q35. What is the output?
/*
func main() {
    x := -5
    fmt.Println(x >> 1)
}

Answer: -3
Explanation: Arithmetic right shift for signed integers
- Preserves sign bit (sign extension)
- -5 in binary (two's complement): ...11111011
- After >> 1: ...11111101 = -3
*/

// Q36. What is the fastest way to multiply a number by 8 using bitwise operations?
/*
Answer:
result := n << 3

Explanation:
- Left shift by 3 positions = multiply by 2^3 = multiply by 8
- Example: 5 << 3 = 40
- Much faster than n * 8 (single CPU instruction vs multiplication)
*/

// Q37. What is the output?
/*
func main() {
    x := 1
    fmt.Println(x << 63)
}

Answer: -9223372036854775808 (on 64-bit system with int being int64)
Explanation:
- Shifting 1 left by 63 positions sets the sign bit
- Results in the minimum int64 value (most negative number)
- Binary: 1000...0000 (1 followed by 63 zeros)
*/

// ============================================================================
// SECTION 5: STRCONV PACKAGE - ADVANCED QUESTIONS
// ============================================================================

// Q38. What is the output?
/*
func main() {
    val, err := strconv.Atoi("  42  ")
    fmt.Println(val, err)
}

Answer: 0 strconv.Atoi: parsing "  42  ": invalid syntax
Explanation: Atoi does NOT trim whitespace. String must be exact number.
Use strings.TrimSpace() first: strconv.Atoi(strings.TrimSpace("  42  "))
*/

// Q39. What is the output?
/*
func main() {
    val, err := strconv.Atoi("42.5")
    fmt.Println(val, err)
}

Answer: 0 strconv.Atoi: parsing "42.5": invalid syntax
Explanation: Atoi only works with integers, not floats.
Use strconv.ParseFloat() for decimal numbers.
*/

// Q40. What is the difference between Atoi and ParseInt?
/*
Answer:
Atoi:
- Converts string to int (platform-dependent size)
- Base 10 only
- Simpler, less flexible
- Signature: func Atoi(s string) (int, error)

ParseInt:
- Converts string to int64
- Supports different bases (2-36)
- Supports different bit sizes (0, 8, 16, 32, 64)
- More flexible
- Signature: func ParseInt(s string, base int, bitSize int) (int64, error)

Example:
val1, _ := strconv.Atoi("42")           // int, base 10
val2, _ := strconv.ParseInt("42", 10, 64) // int64, base 10
val3, _ := strconv.ParseInt("101010", 2, 64) // int64, base 2 (binary)
*/

// Q41. What is the output?
/*
func main() {
    val, _ := strconv.ParseInt("FF", 16, 64)
    fmt.Println(val)
}

Answer: 255
Explanation: ParseInt converts hexadecimal "FF" to decimal
FF (base 16) = 15*16 + 15 = 255 (base 10)
*/

// Q42. What is the output?
/*
func main() {
    val, _ := strconv.ParseInt("1111", 2, 8)
    fmt.Println(val)
}

Answer: 15
Explanation: ParseInt converts binary "1111" to decimal
1111 (base 2) = 8 + 4 + 2 + 1 = 15 (base 10)
bitSize 8 means result fits in int8 range
*/

// Q43. Will this code compile and what's the output?
/*
func main() {
    s := strconv.Itoa(42)
    fmt.Printf("%T %v\n", s, s)
}

Answer: YES - Compiles and runs
Output: string 42
Explanation: Itoa returns a string, not an int
*/

// Q44. What is the output?
/*
func main() {
    val, _ := strconv.ParseInt("1000", 10, 8)
    fmt.Println(val)
}

Answer: 127
Explanation: 1000 exceeds int8 range (-128 to 127)
ParseInt clamps to maximum value: 127
Note: Error would be non-nil indicating overflow
*/

// ============================================================================
// SECTION 6: MATH PACKAGE - EDGE CASES
// ============================================================================

// Q45. What is the output?
/*
func main() {
    fmt.Println(math.Sqrt(-1))
}

Answer: NaN
Explanation: Square root of negative number is Not a Number (NaN)
*/

// Q46. What is the output?
/*
func main() {
    fmt.Println(math.Pow(0, 0))
}

Answer: 1
Explanation: By mathematical convention, 0^0 = 1 in Go's math package
*/

// Q47. What is the output?
/*
func main() {
    fmt.Println(math.Round(2.5))
    fmt.Println(math.Round(3.5))
}

Answer:
2
4
Explanation: Go uses "round half to even" (banker's rounding)
- 2.5 rounds to 2 (nearest even)
- 3.5 rounds to 4 (nearest even)
This reduces cumulative rounding errors in financial calculations
*/

// Q48. What is the output?
/*
func main() {
    fmt.Println(math.Ceil(-4.2))
    fmt.Println(math.Floor(-4.8))
}

Answer:
-4
-5
Explanation:
- Ceil rounds UP (toward positive infinity): -4.2 → -4
- Floor rounds DOWN (toward negative infinity): -4.8 → -5
*/

// Q49. What are the values of these constants?
/*
func main() {
    fmt.Println(math.MaxInt64)
    fmt.Println(math.MinInt64)
}

Answer:
9223372036854775807
-9223372036854775808
Explanation:
- MaxInt64 = 2^63 - 1
- MinInt64 = -2^63
*/

// Q50. What is the output?
/*
func main() {
    x := math.Abs(-5)
    fmt.Printf("%T\n", x)
}

Answer: float64
Explanation: math.Abs returns float64, not int
Even though input looks like int, it's converted to float64
*/

// ============================================================================
// SECTION 7: TYPE CONVERSION - TRICKY SCENARIOS
// ============================================================================

// Q51. What is the output?
/*
func main() {
    var x int = -1
    var y uint = uint(x)
    fmt.Println(y)
}

Answer: 18446744073709551615 (on 64-bit system)
Explanation: Converting -1 to unsigned wraps around
-1 in two's complement = all bits set = max uint value
*/

// Q52. What is the output?
/*
func main() {
    var x float64 = 1.99999999999999999
    var y int = int(x)
    fmt.Println(y)
}

Answer: 1
Explanation: Float to int conversion truncates (doesn't round)
Any value < 2.0 becomes 1
*/

// Q53. What is the output?
/*
func main() {
    var x int32 = math.MaxInt32
    x = x + 1
    fmt.Println(x)
}

Answer: Compilation Error
Error: "cannot use math.MaxInt32 (type int) as type int32 in assignment"
Explanation: math.MaxInt32 is untyped int constant, but needs explicit conversion
Use: var x int32 = int32(math.MaxInt32)
Then x + 1 would overflow to -2147483648
*/

// Q54. What is the output?
/*
func main() {
    x := 256
    y := byte(x)
    fmt.Println(y)
}

Answer: 0
Explanation: byte is uint8 (0-255)
256 % 256 = 0 (wraps around)
*/

// Q55. What is the output?
/*
func main() {
    var x int = 10
    var y int = 3
    var z float64 = x / y
    fmt.Println(z)
}

Answer: 3
Explanation: x / y is integer division (= 3), then converted to float64
Result: 3.0 (printed as 3)
Correct: var z float64 = float64(x) / float64(y) → 3.333...
*/

// ============================================================================
// SECTION 8: PRACTICAL PROBLEM-SOLVING QUESTIONS
// ============================================================================

// Q56. How do you check if a number is a power of 2 using bitwise operations?
/*
Answer:
func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n - 1)) == 0
}

Explanation:
Powers of 2 have only one bit set:
- 1 = 0001
- 2 = 0010
- 4 = 0100
- 8 = 1000

For power of 2: n & (n-1) clears the single bit, resulting in 0
Examples:
- 8 (1000) & 7 (0111) = 0 → true
- 6 (0110) & 5 (0101) = 4 → false
*/

// Q57. How do you count the number of set bits (1s) in an integer?
/*
Answer:
func countSetBits(n int) int {
    count := 0
    for n > 0 {
        count += n & 1  // Check if last bit is 1
        n >>= 1         // Right shift to check next bit
    }
    return count
}

Example: 13 (1101) has 3 set bits

Alternative (faster):
func countSetBits(n int) int {
    count := 0
    for n > 0 {
        n &= n - 1  // Clears rightmost set bit
        count++
    }
    return count
}
*/

// Q58. How do you reverse the bits of an integer?
/*
Answer:
func reverseBits(n uint32) uint32 {
    var result uint32
    for i := 0; i < 32; i++ {
        result <<= 1           // Shift result left
        result |= n & 1        // Add rightmost bit of n
        n >>= 1                // Shift n right
    }
    return result
}

Example: 13 (00001101) → 2952790016 (10110000...)
*/

// Q59. How do you find the absolute value of an integer without using math.Abs?
/*
Answer:
Method 1 (using if):
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

Method 2 (bitwise trick):
func abs(n int) int {
    mask := n >> 63  // All 1s if negative, all 0s if positive
    return (n ^ mask) - mask
}

Explanation of Method 2:
- If n >= 0: mask = 0, result = (n ^ 0) - 0 = n
- If n < 0: mask = -1 (all 1s), result = (n ^ -1) - (-1) = -n
*/

// Q60. How do you safely add two integers and detect overflow?
/*
Answer:
func addWithOverflowCheck(a, b int) (int, bool) {
    sum := a + b

    // Overflow occurs when:
    // - Both positive and sum is negative
    // - Both negative and sum is positive
    overflow := (a > 0 && b > 0 && sum < 0) || (a < 0 && b < 0 && sum > 0)

    return sum, overflow
}

Example:
sum, overflow := addWithOverflowCheck(math.MaxInt64, 1)
// sum = -9223372036854775808, overflow = true

Alternative using math/big for arbitrary precision:
import "math/big"

func safeAdd(a, b int64) *big.Int {
    x := big.NewInt(a)
    y := big.NewInt(b)
    return x.Add(x, y)
}
*/

// ============================================================================
// SECTION 9: CONCEPTUAL & BEST PRACTICES
// ============================================================================

// Q61. When should you use int vs int64 vs int32?
/*
Answer:
Use int (default):
- General-purpose integer operations
- Platform-dependent (32-bit or 64-bit)
- Most common choice
- Example: loop counters, array indices

Use int64:
- When you need guaranteed 64-bit range
- Cross-platform consistency
- Large numbers (timestamps, IDs)
- Example: Unix timestamps, database IDs

Use int32:
- When interfacing with systems requiring 32-bit
- Memory optimization (large arrays)
- Example: protobuf, some APIs

Use int8/int16:
- Extreme memory constraints
- Specific protocol requirements
- Example: embedded systems, network protocols
*/

// Q62. What are the performance implications of bitwise operations?
/*
Answer:
Bitwise operations are EXTREMELY fast:
- Single CPU cycle (usually)
- No function call overhead
- Direct hardware support

Performance comparison (approximate):
- Bitwise AND/OR/XOR: 1 cycle
- Shift operations: 1 cycle
- Addition/Subtraction: 1-3 cycles
- Multiplication: 3-5 cycles
- Division: 20-40 cycles

Use cases for optimization:
- x * 2 → x << 1 (faster)
- x / 2 → x >> 1 (faster)
- x % 2 → x & 1 (faster)
- Check even: x & 1 == 0 (faster than x % 2 == 0)

Note: Modern compilers often optimize these automatically
*/

// Q63. Why does Go not allow implicit type conversion?
/*
Answer:
Reasons for explicit conversion:
1. Safety: Prevents accidental data loss
2. Clarity: Makes conversions visible and intentional
3. Predictability: No hidden behavior
4. Prevents bugs: Catches type mismatches at compile time

Example of prevented bugs:
var x int32 = 1000000
var y int8 = x  // Compilation error (would overflow)

In languages with implicit conversion, this might silently overflow.

Best practice: Always be explicit about type conversions
*/

// Q64. What is the difference between signed and unsigned right shift?
/*
Answer:
Signed right shift (>>):
- Arithmetic shift
- Preserves sign bit (sign extension)
- Example: -8 >> 1 = -4

Unsigned right shift:
- Logical shift
- Fills with zeros
- Example: uint(8) >> 1 = 4

Go behavior:
- >> on signed types: arithmetic shift
- >> on unsigned types: logical shift

Note: Some languages (like Java) have >>> for unsigned shift,
but Go determines behavior based on type.
*/

// Q65. How does integer overflow differ from floating-point overflow?
/*
Answer:
Integer overflow:
- Wraps around silently
- No error or panic
- Predictable behavior (modulo arithmetic)
- Example: int8(127) + 1 = -128

Floating-point overflow:
- Results in +Inf or -Inf
- No panic
- Can be checked with math.IsInf()
- Example: math.MaxFloat64 * 2 = +Inf

Best practices:
- For integers: Check bounds before operations
- For floats: Check for Inf/NaN after operations
- Use math/big for arbitrary precision when needed
*/

// ============================================================================
// END OF INTERVIEW QUESTIONS
// ============================================================================
