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

