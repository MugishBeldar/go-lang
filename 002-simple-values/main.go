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
