// ============================================================================
// SIMPLE VALUES IN GO - A Complete Guide
// ============================================================================
// Go has various basic/primitive data types called "simple values"
// These are the building blocks of all Go programs
// Go is a statically typed language - types are determined at compile time
// ============================================================================

package main // Every Go file must belong to a package. 'main' is special - it's the entry point

import "fmt" // 'fmt' package provides formatted I/O functions (like Printf, Println)

func main() { // main() function is the entry point of the program - execution starts here

	// =========================================================================
	// STRING - A sequence of characters (immutable in Go)
	// =========================================================================
	// - Strings are enclosed in double quotes ""
	// - Strings are UTF-8 encoded by default (supports international characters)
	// - Strings are immutable - once created, cannot be changed
	// - Under the hood: a string is a read-only slice of bytes
	// =========================================================================
	fmt.Println("=============STRING==========")

	fmt.Println("Hello, World!") // String literal - direct text in double quotes

	fmt.Println("1+1=", 1+1) // Multiple arguments to Println - separated by space in output

	fmt.Println("go" + " " + "lang") // String concatenation using + operator (creates new string)
	// Note: For heavy string concatenation, use strings.Builder for better performance

	// =========================================================================
	// BOOLEAN - Represents truth values (true or false)
	// =========================================================================
	// - Only two possible values: true, false (lowercase, not True/False)
	// - Default value (zero value) is false
	// - Used in conditions, loops, and logical expressions
	// - Cannot be converted to/from integers (unlike C/C++)
	// =========================================================================
	fmt.Println("=============BOOLEAN==========")

	fmt.Println(true)  // Boolean literal: true
	fmt.Println(false) // Boolean literal: false

	// Logical Operators:
	fmt.Println(true && false) // AND operator: true only if BOTH are true → false
	fmt.Println(true || false) // OR operator: true if AT LEAST ONE is true → true
	fmt.Println(!true)         // NOT operator: inverts the boolean → false

	// Advanced: Go uses short-circuit evaluation
	// - In &&: if first is false, second is not evaluated
	// - In ||: if first is true, second is not evaluated

	// =========================================================================
	// INTEGER - Whole numbers (no decimal point)
	// =========================================================================
	// Types: int, int8, int16, int32, int64 (signed - can be negative)
	//        uint, uint8, uint16, uint32, uint64 (unsigned - only positive)
	// - 'int' is platform-dependent: 32-bit on 32-bit systems, 64-bit on 64-bit
	// - Default value (zero value) is 0
	// - int8: -128 to 127 | int16: -32768 to 32767
	// - int32: ~-2 billion to ~2 billion | int64: huge range
	// =========================================================================
	fmt.Println("=============INTEGER==========")

	fmt.Println(1 + 1)  // Addition: 2
	fmt.Println(7 - 3)  // Subtraction: 4
	fmt.Println(3 * 3)  // Multiplication: 9
	fmt.Println(10 / 2) // Division: 5 (integer division - no decimals)
	fmt.Println(10 % 3) // Modulo/Remainder: 1 (10 divided by 3 leaves remainder 1)

	// Advanced: Integer overflow wraps around silently (no error by default)
	// Example: int8(127) + 1 = -128 (wraps to minimum)

	// =========================================================================
	// FLOAT - Numbers with decimal points (floating-point)
	// =========================================================================
	// Types: float32 (single precision), float64 (double precision - recommended)
	// - float64 is the default for decimal literals
	// - Default value (zero value) is 0.0
	// - float32: ~7 decimal digits precision
	// - float64: ~15 decimal digits precision
	// - WARNING: Floating point arithmetic can have precision issues!
	// =========================================================================
	fmt.Println("=============FLOAT==========")

	fmt.Println(1.0 + 1.0)  // Addition: 2
	fmt.Println(7.0 - 3.0)  // Subtraction: 4
	fmt.Println(3.0 * 3.0)  // Multiplication: 9
	fmt.Println(10.0 / 2.0) // Division: 5 (true division with decimals)

	// Advanced: Be careful with float comparisons!
	// 0.1 + 0.2 might not exactly equal 0.3 due to binary representation

	// =========================================================================
	// COMPLEX - Complex numbers (real + imaginary parts)
	// =========================================================================
	// Types: complex64 (float32 real + float32 imaginary)
	//        complex128 (float64 real + float64 imaginary)
	// - Format: complex(real, imaginary) or using literal: 3+4i
	// - Used in scientific computing, signal processing, etc.
	// - real() and imag() functions extract parts
	// =========================================================================
	fmt.Println("=============COMPLEX==========")
	// i = sqrt(-1)
	// i^2 = -1
	// i^3 = -i
	// i^4 = 1

	fmt.Println(complex(1, 2) + complex(3, 4))    // (1+2i) + (3+4i) = (4+6i)
	fmt.Println(complex(5, 6) - complex(7, 8))    // (5+6i) - (7+8i) = (-2-2i)
	fmt.Println(complex(2, 3) * complex(4, 5))    // Complex multiplication
	fmt.Println(complex(9, 10) / complex(11, 12)) // Complex division

	// Advanced: You can also write complex literals directly: 3+4i

	// =========================================================================
	// BYTE - Alias for uint8 (0 to 255)
	// =========================================================================
	// - byte is exactly the same as uint8
	// - Commonly used to represent ASCII characters
	// - Used when working with raw binary data, files, network data
	// - A string is essentially a sequence of bytes
	// =========================================================================
	fmt.Println("=============BYTE==========")

	fmt.Println(byte(65)) // 65 is ASCII code for 'A' - prints: 65
	fmt.Println(byte(66)) // 66 is ASCII code for 'B' - prints: 66
	fmt.Println(byte(67)) // 67 is ASCII code for 'C' - prints: 67

	// To print as character: fmt.Printf("%c\n", byte(65)) → prints: A

	// =========================================================================
	// RUNE - Alias for int32, represents a Unicode code point
	// =========================================================================
	// - rune is exactly the same as int32
	// - Used to represent any Unicode character (not just ASCII)
	// - Can represent characters from any language (Chinese, Arabic, Emoji, etc.)
	// - Single quotes '' create a rune literal: 'A', '中', '🚀'
	// - A string in Go is a sequence of runes encoded as UTF-8 bytes
	// =========================================================================
	fmt.Println("=============RUNE==========")

	fmt.Println(rune(65)) // 65 is Unicode code point for 'A' - prints: 65
	fmt.Println(rune(66)) // 66 is Unicode code point for 'B' - prints: 66
	fmt.Println(rune(67)) // 67 is Unicode code point for 'C' - prints: 67

	// Advanced: len("hello") returns bytes count, not character count
	// For string with Unicode: len("世界") = 6 bytes, but 2 runes
	// Use utf8.RuneCountInString() or range loop for accurate count

	// =========================================================================
	// KEY DIFFERENCES: byte vs rune
	// =========================================================================
	// byte (uint8): 0-255, for ASCII and raw binary data
	// rune (int32): -2B to +2B, for any Unicode character
	// Use byte for: file I/O, network data, ASCII text
	// Use rune for: text processing with international characters
	// =========================================================================

	// =========================================================================
	// ZERO VALUES - Default values when variables are declared but not initialized
	// =========================================================================
	// string → "" (empty string)
	// bool   → false
	// int    → 0
	// float  → 0.0
	// complex → (0+0i)
	// =========================================================================
}

// ============================================================================
// INTERVIEW QUESTIONS - SIMPLE VALUES IN GO
// ============================================================================

// ============================================================================
// SECTION 1: BASIC LEVEL QUESTIONS
// ============================================================================

// Q1. What is the output of the following code?
/*
func main() {
    fmt.Println("Hello" + "World")
}
Answer: HelloWorld
Explanation: String concatenation using + operator joins strings without space
*/

// Q2. What is the zero value of a boolean in Go?
/*
Answer: false
Explanation: When a bool variable is declared but not initialized, it defaults to false
*/

// Q3. What is the difference between int and uint in Go?
/*
Answer:
- int: signed integer (can be negative or positive)
- uint: unsigned integer (only positive values, 0 and above)
- Both are platform-dependent (32-bit or 64-bit based on system)
*/

// Q4. Will this code compile? If not, what error will you get?
/*
func main() {
    var x int = 10
    var y float64 = 20.5
    fmt.Println(x + y)
}
Answer: NO - Compilation Error
Error: "invalid operation: x + y (mismatched types int and float64)"
Explanation: Go does not allow implicit type conversion. You must explicitly convert:
    fmt.Println(float64(x) + y) OR fmt.Println(x + int(y))
*/

// Q5. What is the output?
/*
func main() {
    fmt.Println(10 / 3)
}
Answer: 3
Explanation: Integer division truncates the decimal part. Result is 3, not 3.333...
*/

// Q6. What is the output?
/*
func main() {
    fmt.Println(10.0 / 3.0)
}
Answer: 3.3333333333333335
Explanation: Float division returns decimal result
*/

// Q7. What is the difference between byte and rune?
/*
Answer:
- byte: alias for uint8 (0-255), used for ASCII characters and raw binary data
- rune: alias for int32, used for Unicode code points (any character from any language)
*/

// Q8. What is the output?
/*
func main() {
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
Answer:
false
true
false
Explanation: && (AND), || (OR), ! (NOT) logical operators
*/

// ============================================================================
// SECTION 2: INTERMEDIATE LEVEL QUESTIONS
// ============================================================================

// Q9. Will this code run? What will be the output or error?
/*
func main() {
    var x int8 = 127
    x = x + 1
    fmt.Println(x)
}
Answer: YES, it will compile and run
Output: -128
Explanation: Integer overflow! int8 range is -128 to 127. Adding 1 to 127 wraps around to -128.
Go does not throw runtime errors for integer overflow - it wraps silently.
*/

// Q10. What is the output?
/*
func main() {
    fmt.Println(10 % 3)
    fmt.Println(10 % 4)
    fmt.Println(10 % 5)
}
Answer:
1
2
0
Explanation: Modulo operator (%) returns the remainder after division
*/

// Q11. Will this code compile? What error will you get?
/*
func main() {
    var isTrue bool = 1
}
Answer: NO - Compilation Error
Error: "cannot use 1 (type untyped int) as type bool in assignment"
Explanation: Unlike C/C++, Go does not allow conversion between bool and int.
You cannot use 1 for true or 0 for false.
*/

// Q12. What is the output?
/*
func main() {
    fmt.Println(0.1 + 0.2 == 0.3)
}
Answer: false (most likely)
Explanation: Floating-point precision issue! Due to binary representation,
0.1 + 0.2 might be 0.30000000000000004, not exactly 0.3.
Never use == for float comparison. Use a tolerance/epsilon instead.
*/

// Q13. What is the output?
/*
func main() {
    var s string
    var b bool
    var i int
    var f float64
    fmt.Printf("%q %v %v %v\n", s, b, i, f)
}
Answer: "" false 0 0
Explanation: These are zero values - default values for uninitialized variables
*/

// Q14. Will this code compile?
/*
func main() {
    const x = 10
    const y = 20.5
    fmt.Println(x + y)
}
Answer: YES
Output: 30.5
Explanation: Untyped constants can be mixed in expressions. The compiler
determines the appropriate type at compile time.
*/

// Q15. What is the output?
/*
func main() {
    fmt.Println(len("Hello"))
    fmt.Println(len("世界"))
}
Answer:
5
6
Explanation: len() returns byte count, not character count.
"Hello" = 5 ASCII bytes
"世界" = 6 bytes (each Chinese character is 3 bytes in UTF-8)
*/

// ============================================================================
// SECTION 3: ADVANCED LEVEL QUESTIONS
// ============================================================================

// Q16. What is the output and why?
/*
func main() {
    var x int = 10
    var y int32 = 20
    fmt.Println(x + y)
}
Answer: Compilation Error
Error: "invalid operation: x + y (mismatched types int and int32)"
Explanation: Even though both are integers, int and int32 are different types.
Go requires explicit conversion: fmt.Println(x + int(y)) or fmt.Println(int32(x) + y)
*/

// Q17. What happens in this code? Will it panic?
/*
func main() {
    var x uint8 = 0
    x = x - 1
    fmt.Println(x)
}
Answer: NO panic, it will run
Output: 255
Explanation: Unsigned integer underflow wraps around. 0 - 1 = 255 for uint8.
*/

// Q18. What is the output?
/*
func main() {
    fmt.Println(complex(3, 4) * complex(1, 2))
}
Answer: (-5+10i)
Explanation: Complex multiplication: (3+4i) * (1+2i)
= 3*1 + 3*2i + 4i*1 + 4i*2i
= 3 + 6i + 4i + 8i²
= 3 + 10i + 8(-1)  [since i² = -1]
= 3 + 10i - 8
= -5 + 10i
*/

// Q19. Will this code compile? What's the issue?
/*
func main() {
    var x float64 = 10.5
    var y int = x
    fmt.Println(y)
}
Answer: NO - Compilation Error
Error: "cannot use x (type float64) as type int in assignment"
Explanation: Go does not allow implicit type conversion, even when it seems "safe".
Must use: var y int = int(x)
Note: Converting 10.5 to int truncates to 10 (not rounds)
*/

// Q20. What is the output?
/*
func main() {
    s := "Hello, 世界"
    fmt.Println(len(s))
    count := 0
    for range s {
        count++
    }
    fmt.Println(count)
}
Answer:
13
9
Explanation:
- len(s) = 13 bytes (5 ASCII + 8 bytes for two Chinese characters)
- range loop iterates over runes, not bytes, so count = 9 characters
*/

// Q21. What is the output?
/*
func main() {
    fmt.Println(true && false || true)
    fmt.Println(true && (false || true))
}
Answer:
true
true
Explanation:
Line 1: && has higher precedence than ||, so (true && false) || true = false || true = true
Line 2: Parentheses force evaluation of (false || true) first = true, then true && true = true
*/

// Q22. Will this code compile and run? What's the output?
/*
func main() {
    var b byte = 255
    b = b + 1
    fmt.Println(b)
}
Answer: YES, compiles and runs
Output: 0
Explanation: byte is uint8 (0-255). Adding 1 to 255 causes overflow, wraps to 0.
*/

// Q23. What is the output?
/*
func main() {
    var x int = 5
    var y int = 2
    fmt.Println(x / y)
    fmt.Println(float64(x) / float64(y))
}
Answer:
2
2.5
Explanation:
- Integer division: 5 / 2 = 2 (truncated)
- Float division: 5.0 / 2.0 = 2.5 (precise)
*/

// ============================================================================
// SECTION 4: TRICKY CODE EXECUTION QUESTIONS
// ============================================================================

// Q24. What is the output?
/*
func main() {
    fmt.Println("1" + "2" + "3")
    fmt.Println(1 + 2 + 3)
}
Answer:
123
6
Explanation:
- String concatenation: "1" + "2" + "3" = "123"
- Integer addition: 1 + 2 + 3 = 6
*/

// Q25. Will this compile? What error?
/*
func main() {
    var x = 10
    var y = 20.5
    var z = x + y
    fmt.Println(z)
}
Answer: NO - Compilation Error
Error: "invalid operation: x + y (mismatched types int and float64)"
Explanation: Type inference gives x type int and y type float64.
Cannot add different types without explicit conversion.
*/

// Q26. What is the output?
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
Explanation: The sign of the result follows the sign of the dividend (first operand)
*/

// Q27. What happens here?
/*
func main() {
    var x int
    fmt.Println(x)
    x = x + 1
    fmt.Println(x)
}
Answer: Compiles and runs successfully
Output:
0
1
Explanation: Uninitialized int has zero value 0. Then incremented to 1.
*/

// Q28. What is the output?
/*
func main() {
    fmt.Println(false && someFunction())
    fmt.Println(true || someFunction())
}
func someFunction() bool {
    fmt.Println("Function called")
    return true
}
Answer:
false
true
Explanation: Short-circuit evaluation!
- In &&: if first is false, second is NOT evaluated (someFunction not called)
- In ||: if first is true, second is NOT evaluated (someFunction not called)
*/

// ============================================================================
// SECTION 5: ERROR IDENTIFICATION QUESTIONS
// ============================================================================

// Q29. Find the error in this code:
/*
func main() {
    var name string = "John"
    name = 'J'
    fmt.Println(name)
}
Answer: Compilation Error
Error: "cannot use 'J' (type rune) as type string in assignment"
Explanation: Single quotes '' create a rune (int32), not a string.
Use double quotes: name = "J"
*/

// Q30. Find the error:
/*
func main() {
    const x int = 10
    x = 20
    fmt.Println(x)
}
Answer: Compilation Error
Error: "cannot assign to x (declared const)"
Explanation: Constants cannot be reassigned after declaration.
*/

// Q31. Find the error:
/*
func main() {
    var x int = 10
    var y int = 0
    fmt.Println(x / y)
}
Answer: Runtime Panic
Error: "runtime error: integer divide by zero"
Explanation: Division by zero causes a panic at runtime (not compile time).
*/

// Q32. Find the error:
/*
func main() {
    var x float64 = 10.5
    var y float64 = 0.0
    fmt.Println(x / y)
}
Answer: NO ERROR - Runs successfully
Output: +Inf
Explanation: Float division by zero does NOT panic. It returns +Inf, -Inf, or NaN.
*/

// Q33. Find the error:
/*
func main() {
    var isActive bool = "true"
    fmt.Println(isActive)
}
Answer: Compilation Error
Error: "cannot use "true" (type string) as type bool in assignment"
Explanation: String "true" is not the same as boolean true.
Use: var isActive bool = true
*/

// ============================================================================
// SECTION 6: CONCEPTUAL & BEST PRACTICES QUESTIONS
// ============================================================================

// Q34. Why are strings immutable in Go? What are the benefits?
/*
Answer:
1. Thread Safety: Multiple goroutines can safely read the same string
2. Memory Efficiency: Same string literal can be shared across the program
3. Security: Prevents accidental modification of string data
4. Optimization: Compiler can optimize string operations better
5. Hash Keys: Strings can be safely used as map keys

To modify strings efficiently, use strings.Builder or []byte
*/

// Q35. When should you use float32 vs float64?
/*
Answer:
Use float64 (default):
- When precision is important
- For most general-purpose calculations
- Default for floating-point literals

Use float32:
- When memory is constrained (large arrays/slices)
- When interfacing with systems that require 32-bit floats
- When precision of ~7 decimal digits is sufficient
- Graphics programming, game development (GPU prefers float32)
*/

// Q36. What are the performance implications of string concatenation using +?
/*
Answer:
Using + operator for string concatenation:
- Creates a new string for each operation (strings are immutable)
- Allocates new memory each time
- O(n²) complexity for concatenating n strings in a loop

Better alternatives:
1. strings.Builder - for building strings in loops (O(n) complexity)
2. fmt.Sprintf - for formatted strings
3. strings.Join - for joining slices of strings

Example:
// Bad (in loop)
s := ""
for i := 0; i < 1000; i++ {
    s += "x"  // Creates 1000 new strings
}

// Good
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("x")  // Efficient
}
s := builder.String()
*/

// Q37. Explain short-circuit evaluation with a practical example
/*
Answer:
Short-circuit evaluation means the second operand is not evaluated if the
result can be determined from the first operand alone.

&& (AND): If first is false, result is false (second not evaluated)
|| (OR): If first is true, result is true (second not evaluated)

Practical use case:
func processUser(user *User) {
    if user != nil && user.IsActive {
        // Safe: user.IsActive only evaluated if user != nil
        // Prevents nil pointer dereference
    }
}

Without short-circuit, this would panic if user is nil!
*/

// Q38. What is the difference between these two declarations?
/*
var x = 10
var y int = 10

Answer:
Both create an int variable with value 10, but:

var x = 10
- Type inference: compiler infers type as int
- More concise
- Type can change if you change the value (e.g., var x = 10.5 would be float64)

var y int = 10
- Explicit type declaration
- More verbose but clearer intent
- Type is fixed regardless of value
- Better for documentation and when specific type is required

Best practice: Use type inference for clarity, explicit types when precision matters
(e.g., var x int64 = 10 when you specifically need 64-bit integer)
*/

// Q39. Why does Go have both byte and uint8 if they're the same?
/*
Answer:
byte and uint8 are identical (byte is an alias for uint8), but they serve
different semantic purposes:

byte:
- Used for representing raw binary data
- Used for ASCII characters
- Used in I/O operations (reading files, network data)
- Conveys intent: "this is data/character"

uint8:
- Used for small unsigned integers (0-255)
- Used in arithmetic operations
- Conveys intent: "this is a number"

Example:
var age uint8 = 25        // Semantic: a small number
var data []byte = []byte("Hello")  // Semantic: binary/text data

This improves code readability and intent.
*/

// Q40. What are the risks of integer overflow in Go and how to handle it?
/*
Answer:
Risks:
1. Silent wraparound (no error/panic by default)
2. Security vulnerabilities (buffer overflow, incorrect calculations)
3. Logic errors in financial/scientific calculations

Example of risk:
var x int8 = 127
x = x + 1  // Now x = -128 (wrapped around)

How to handle:
1. Use appropriate size (int64 for large numbers)
2. Check before operations:
   if x > math.MaxInt8 - 1 {
       // Handle overflow
   }
3. Use math/big package for arbitrary precision:
   import "math/big"
   x := big.NewInt(127)
   x.Add(x, big.NewInt(1))  // No overflow
4. Use unsigned types when values are always positive
5. Add validation/bounds checking in critical code
*/

// ============================================================================
// END OF INTERVIEW QUESTIONS
// ============================================================================
