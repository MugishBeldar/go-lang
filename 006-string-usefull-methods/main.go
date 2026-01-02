package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// For proper title casing, use:

func main() {
	// 1. string basics in go
	fmt.Println("============================= STRING BASICS IN GO =============================")
	// string are UTF-8 encoded by default (supports international characters)
	// string are represented as []byte under the hood
	// string are not null-terminated (no \0 at the end)
	// string are not mutable (cannot change after creation)
	s := "hello world"
	fmt.Println("sample string: ", s)
	s2 := `raw string literal 
		can span multiple lines
		and contains \n literally`
	fmt.Println("raw string literal: ", s2)

	// stirng length
	fmt.Println("string length: ", len(s)) // 11 (number of bytes, not characters)

	// accessing characters (returns byte)
	fmt.Println("s[0]: ", s[0]) // 104 (ASCII code for 'h') here 104 is byte value for 'h'

	// for accessing actual characters use blow
	fmt.Println("character of s[0]: ", string(s[0])) // "h"

	// string are immutable - this won't work
	// s[0] = 'H' // error: cannot assign to s[0] (strings are immutable)

	// 2. the string package most important!
	fmt.Println("============================= THE STRING PACKAGE =============================")
	s1 := "the string package"

	// check if contains
	fmt.Println("contains: ", strings.Contains(s1, "the"))         // true
	fmt.Println("contains: ", strings.Contains(s1, "not"))         // false
	fmt.Println("contains any: ", strings.ContainsAny(s, "aeiou")) // true (contains any of these chars)
	fmt.Println("contains rune: ", strings.ContainsRune(s, 'W'))   // false (contains the rune 'W')

	// check prefix/suffix
	fmt.Println("has prefix: ", strings.HasPrefix(s1, "the")) // true
	fmt.Println("has suffix: ", strings.HasSuffix(s1, "age")) // true

	// find position (index)
	fmt.Println("index of 'string': ", strings.Index(s1, "string"))        // 4 first occurrence
	fmt.Println("index of 'not': ", strings.Index(s1, "not"))              // -1 (not found)
	fmt.Println("last index of e: ", strings.LastIndex(s1, "e"))           // 10 last occurrence
	fmt.Println("last index of not: ", strings.LastIndex(s1, "not"))       // -1 (not found)
	fmt.Println("index of rune 'e': ", strings.IndexRune(s1, 'e'))         // 1 first occurrence of rune
	fmt.Println("last index of rune 'e': ", strings.LastIndexAny(s1, "e")) // 10 last occurrence of rune

	// count occurrences
	fmt.Println("total t appears in s1: ", strings.Count(s1, "t"))           // 2
	fmt.Println("total 'aeiou' appears in s1: ", strings.Count(s1, "aeiou")) // 0 (not a substring)

	// 3. case conversion
	fmt.Println("============================= CASE CONVERSION =============================")
	s3 := "Hello, World!"
	strings.ToLower(s3) // "hello, world!"
	strings.ToUpper(s3) // "HELLO, WORLD!"
	strings.ToTitle(s3) // "HELLO, WORLD!" (Unicode title case)
	strings.Title(s3)   // "Hello, World!" (Deprecated in Go 1.18+)

	// for propper title casing use:
	caser := cases.Title(language.English)
	fmt.Println(caser.String(s3)) // "Hello, World!"

	// 4. trimming (remove characters)
	fmt.Println("============================= TRIMMING =============================")
	s4 := "  hello world  "
	fmt.Println("trim space: ", strings.TrimSpace(s4))            // "hello world"
	fmt.Println("trim left space: ", strings.TrimLeft(s4, " "))   // "hello world  "
	fmt.Println("trim right space: ", strings.TrimRight(s4, " ")) // "  hello world"
	fmt.Println("trim prefix: ", strings.TrimPrefix(s4, "  "))    // "hello world  "
	fmt.Println("trim suffix: ", strings.TrimSuffix(s4, "  "))    // "  hello world"
	fmt.Println("trim: ", strings.Trim(s4, " "))                  // "hello world"
	fmt.Println("trim func: ", strings.TrimFunc(s4, func(c rune) bool {
		return c == ' ' || c == 'h'
	})) // "ello world  " (remove all 'h' and space)
	fmt.Println("trim left func: ", strings.TrimLeftFunc(s4, func(c rune) bool {
		return c == ' ' || c == 'h'
	})) // "ello world  " (remove all 'h' and space from left)
	fmt.Println("trim right func: ", strings.TrimRightFunc(s4, func(c rune) bool {
		return c == ' ' || c == 'h'
	})) // "  hello worl" (remove all 'h' and space from right)

	// 5. joining strings
	fmt.Println("============================= JOINING STRINGS =============================")
	s5 := []string{"hello", "world", "!"}
	fmt.Println("join: ", strings.Join(s5, " ")) // "hello world !"
	fmt.Println("join: ", strings.Join(s5, "-")) // "hello-world-!"
	fmt.Println("join: ", strings.Join(s5, ""))  // "helloworld!"

	// 6. replacing
	fmt.Println("============================= REPLACING =============================")
	fmt.Println("replace: ", strings.Replace(s1, "the", "a", 1))     // "a string package" (replace first occurrence)d
	fmt.Println("replace: ", strings.Replace(s1, "the", "a", -1))    // "a string package" (replace all occurrences)
	fmt.Println("replace all: ", strings.ReplaceAll(s1, "the", "a")) // "a string package" (replace all occurrences)
	fmt.Println("repeat: ", strings.Repeat(s1, 3))                   // "the string packagethe string packagethe string package" (repeat 3 times)

	// 7. splitting
	fmt.Println("============================= SPLITTING =============================")
	fmt.Println("split: ", strings.Split(s1, " "))                  // ["the", "string", "package"] (split on space)
	fmt.Println("split after: ", strings.SplitAfter(s1, " "))       // ["the ", "string ", "package"] (split after space)
	fmt.Println("split after n: ", strings.SplitAfterN(s1, " ", 2)) // ["the ", "string package"] (split after space, max 2 parts)
	fmt.Println("split n: ", strings.SplitN(s1, " ", 2))            // ["the", "string package"] (split on space, max 2 parts)

	// 8. padding
	fmt.Println("============================= PADDING =============================")
	fmt.Println("pad left: ", strings.Repeat(" ", 5)+s1)                        // "     the string package" (pad left with 5 spaces)
	fmt.Println("pad right: ", s1+strings.Repeat(" ", 5))                       // "the string package     " (pad right with 5 spaces)
	fmt.Println("pad both: ", strings.Repeat(" ", 2)+s1+strings.Repeat(" ", 2)) // "  the string package  " (pad both sides with 2 spaces)

	// 9. comparing
	fmt.Println("============================= COMPARING =============================")
	fmt.Println("compare: ", strings.Compare(s1, "the string package"))  // 0 (equal)
	fmt.Println("compare: ", strings.Compare(s1, "The string package"))  // 1 (s1 is greater)
	fmt.Println("compare: ", strings.Compare(s1, "the string packages")) // -1 (s1 is less)
	fmt.Println("compare: ", strings.EqualFold("go", "GO"))              // true (case insensitive)
	fmt.Println("compare: ", strings.EqualFold("go", "gO"))              // true (case insensitive)
	fmt.Println("compare: ", strings.EqualFold("go", "go"))              // true (case insensitive)
	fmt.Println("compare: ", strings.EqualFold("go", "golang"))          // false (case insensitive)

	// NOTE:
	// DIRECT COMPARISION PREFERRED FOR EQUALITIY
	// "hello" == "hello" // true
	// "hello" < "world" // ture (lexicographical order)
	// what is lexicographical order?
	// - it's a way of ordering words based on the alphabet
	// - it's the order that words would appear in a dictionary
	// - it's the order that words would appear in a book
	// - it's the order that words would appear in a phone book

	// 10. understanding rues
	fmt.Println("============================= UNDERSTANDING RUNES =============================")
	// in go:
	// string: sequence of bytes (not characters)
	// rune: Unicode code point (a character) int32
	// string: []byte (under the hood)
	// rune: int32 (under the hood)

	// 11. byte slices & strings
	fmt.Println("============================= BYTE SLICES & STRINGS =============================")
	// string to []byte
	fmt.Println("string to []byte: ", []byte(s1)) // "the string package" (convert to byte slice)
	// []byte to string
	fmt.Println("[]byte to string: ", string([]byte(s1))) // "the string package" (convert to string)
}

// ============================================================================
// INTERVIEW QUESTIONS - STRING OPERATIONS IN GO
// ============================================================================

// ============================================================================
// SECTION 1: BASIC LEVEL QUESTIONS
// ============================================================================

// Q1. Are strings mutable or immutable in Go?
/*
Answer: Immutable
Explanation: Once a string is created, it cannot be modified.
Any operation that appears to modify a string actually creates a new string.

Example:
s := "hello"
s[0] = 'H'  // Compilation Error: cannot assign to s[0]

To "modify" a string, convert to []byte or []rune:
b := []byte(s)
b[0] = 'H'
s = string(b)  // Creates new string "Hello"
*/

// Q2. What is the output?
/*
func main() {
    s := "hello"
    fmt.Println(len(s))
}

Answer: 5
Explanation: len() returns the number of bytes in the string
"hello" has 5 ASCII characters = 5 bytes
*/

// Q3. What is the output?
/*
func main() {
    s := "hello"
    fmt.Println(s[0])
}

Answer: 104
Explanation: s[0] returns the byte value at index 0
'h' has ASCII value 104
To get the character: fmt.Println(string(s[0])) → "h"
*/

// Q4. What is the output?
/*
func main() {
    s := "世界"
    fmt.Println(len(s))
}

Answer: 6
Explanation: len() returns bytes, not characters!
Each Chinese character is 3 bytes in UTF-8
"世界" = 2 characters = 6 bytes
*/

// Q5. What is the output?
/*
func main() {
    fmt.Println(strings.Contains("hello world", "world"))
    fmt.Println(strings.Contains("hello world", "World"))
}

Answer:
true
false
Explanation: Contains is case-sensitive
"world" is found, "World" is not
*/

// Q6. What is the output?
/*
func main() {
    fmt.Println(strings.ToUpper("hello"))
    fmt.Println(strings.ToLower("WORLD"))
}

Answer:
HELLO
world
Explanation: ToUpper converts to uppercase, ToLower to lowercase
*/

// Q7. What is the output?
/*
func main() {
    s := "  hello  "
    fmt.Println(strings.TrimSpace(s))
}

Answer: hello
Explanation: TrimSpace removes leading and trailing whitespace
*/

// Q8. What is the output?
/*
func main() {
    parts := []string{"hello", "world"}
    fmt.Println(strings.Join(parts, " "))
}

Answer: hello world
Explanation: Join concatenates slice elements with separator
*/

// ============================================================================
// SECTION 2: INTERMEDIATE LEVEL QUESTIONS
// ============================================================================

// Q9. What is the output?
/*
func main() {
    fmt.Println(strings.Index("hello world", "world"))
    fmt.Println(strings.Index("hello world", "xyz"))
}

Answer:
6
-1
Explanation:
- Index returns position of first occurrence (0-based)
- Returns -1 if not found
*/

// Q10. What is the output?
/*
func main() {
    fmt.Println(strings.Count("hello", "l"))
    fmt.Println(strings.Count("hello", "ll"))
}

Answer:
2
1
Explanation:
- Count returns number of non-overlapping occurrences
- "hello" has 2 'l' characters
- "hello" has 1 "ll" substring
*/

// Q11. What is the difference between strings.HasPrefix and strings.Contains?
/*
Answer:
HasPrefix: Checks if string STARTS with prefix
Contains: Checks if string contains substring ANYWHERE

Example:
s := "hello world"
strings.HasPrefix(s, "hello")  // true
strings.HasPrefix(s, "world")  // false
strings.Contains(s, "world")   // true
*/

// Q12. What is the output?
/*
func main() {
    s := "hello world hello"
    fmt.Println(strings.Replace(s, "hello", "hi", 1))
    fmt.Println(strings.Replace(s, "hello", "hi", -1))
}

Answer:
hi world hello
hi world hi
Explanation:
- Replace(s, old, new, n) replaces first n occurrences
- n = 1: replace first occurrence only
- n = -1: replace all occurrences
*/

// Q13. What is the output?
/*
func main() {
    s := "a,b,c"
    fmt.Println(strings.Split(s, ","))
}

Answer: [a b c]
Explanation: Split divides string into slice based on separator
*/

// Q14. What is the output?
/*
func main() {
    s := "a,b,c"
    fmt.Println(strings.SplitN(s, ",", 2))
}

Answer: [a b,c]
Explanation: SplitN splits into at most n parts
n=2 means split into 2 parts maximum
*/

// Q15. What is the output?
/*
func main() {
    fmt.Println(strings.Repeat("Go", 3))
}

Answer: GoGoGo
Explanation: Repeat returns string repeated n times
*/

// Q16. What is the output?
/*
func main() {
    fmt.Println(strings.Compare("abc", "abc"))
    fmt.Println(strings.Compare("abc", "xyz"))
    fmt.Println(strings.Compare("xyz", "abc"))
}

Answer:
0
-1
1
Explanation:
- Returns 0 if equal
- Returns -1 if first < second (lexicographically)
- Returns 1 if first > second
*/

// Q17. What is the output?
/*
func main() {
    fmt.Println(strings.EqualFold("Go", "go"))
    fmt.Println(strings.EqualFold("Go", "GO"))
}

Answer:
true
true
Explanation: EqualFold compares strings case-insensitively
*/

// Q18. What is the output?
/*
func main() {
    s := "hello"
    s = s + " world"
    fmt.Println(s)
}

Answer: hello world
Explanation: String concatenation with + creates a new string
Original "hello" is unchanged (immutable)
*/

// ============================================================================
// SECTION 3: ADVANCED LEVEL QUESTIONS
// ============================================================================

// Q19. What is the difference between len() and counting runes?
/*
Answer:
len(): Returns number of BYTES
Rune count: Returns number of CHARACTERS

Example:
s := "Hello, 世界"
fmt.Println(len(s))                    // 13 bytes
fmt.Println(utf8.RuneCountInString(s)) // 9 characters

Or using range:
count := 0
for range s {
    count++
}
fmt.Println(count)  // 9 characters

Why different?
- ASCII characters: 1 byte each
- Chinese characters: 3 bytes each in UTF-8
- "Hello, " = 7 bytes
- "世界" = 6 bytes (2 chars × 3 bytes)
- Total: 13 bytes, 9 characters
*/

// Q20. What is the output?
/*
func main() {
    s := "hello"
    b := []byte(s)
    b[0] = 'H'
    fmt.Println(s)
    fmt.Println(string(b))
}

Answer:
hello
Hello
Explanation:
- Original string s is unchanged (immutable)
- []byte conversion creates a copy
- Modifying b doesn't affect s
- string(b) creates new string "Hello"
*/

// Q21. What is the output?
/*
func main() {
    s := "hello"
    r := []rune(s)
    r[0] = 'H'
    fmt.Println(string(r))
}

Answer: Hello
Explanation:
- Convert string to []rune (slice of Unicode code points)
- Modify first rune
- Convert back to string
This is the proper way to "modify" a string
*/

// Q22. What is the output?
/*
func main() {
    s := "世界"
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c ", s[i])
    }
}

Answer: ä ¸ � ç � �
Explanation: Iterating by byte index breaks UTF-8 characters!
Each Chinese character is 3 bytes, so you get partial bytes.

Correct way:
for _, r := range s {
    fmt.Printf("%c ", r)  // 世 界
}
*/

// Q23. What is the difference between string and []byte?
/*
Answer:
string:
- Immutable
- UTF-8 encoded
- Stored as read-only byte array
- Efficient for read operations
- Can be used as map keys

[]byte:
- Mutable
- Can be modified in place
- More memory allocations when converting
- Used for I/O operations
- Cannot be used as map keys directly

Conversion:
b := []byte(s)  // string to []byte (copies data)
s := string(b)  // []byte to string (copies data)

Performance: Conversions involve copying, so avoid in hot paths
*/

// Q24. What is the output?
/*
func main() {
    s := "hello"
    s2 := s[1:4]
    fmt.Println(s2)
}

Answer: ell
Explanation: String slicing creates substring
s[1:4] = characters at index 1, 2, 3 (4 is exclusive)
*/

// Q25. Will this code compile?
/*
func main() {
    s := "hello"
    s[0] = 'H'
    fmt.Println(s)
}

Answer: NO - Compilation Error
Error: "cannot assign to s[0] (strings are immutable)"
Explanation: Strings are immutable in Go
*/

// Q26. What is the output?
/*
func main() {
    s := "hello"
    s = strings.ToUpper(s)
    fmt.Println(s)
}

Answer: HELLO
Explanation: ToUpper returns NEW string (doesn't modify original)
We reassign s to the new uppercase string
*/

// Q27. What is the output?
/*
func main() {
    s := ""
    for i := 0; i < 3; i++ {
        s += "a"
    }
    fmt.Println(s)
}

Answer: aaa
Explanation: Works but inefficient!
Each += creates a new string (3 allocations)
Better: use strings.Builder for loops
*/

// Q28. What is the better way to build strings in a loop?
/*
Answer: Use strings.Builder

Bad (inefficient):
s := ""
for i := 0; i < 1000; i++ {
    s += "a"  // Creates 1000 new strings!
}

Good (efficient):
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("a")  // Efficient, grows buffer
}
s := builder.String()

Why?
- strings.Builder minimizes memory allocations
- Grows internal buffer as needed
- Much faster for repeated concatenations
- O(n) vs O(n²) complexity
*/

// Q29. What is the output?
/*
func main() {
    s := "hello"
    s2 := s
    s = "world"
    fmt.Println(s2)
}

Answer: hello
Explanation: Strings are immutable
s2 still points to original "hello"
Reassigning s doesn't affect s2
*/

// Q30. What is the output?
/*
func main() {
    fmt.Println(strings.TrimLeft("  hello  ", " "))
    fmt.Println(strings.TrimRight("  hello  ", " "))
    fmt.Println(strings.Trim("  hello  ", " "))
}

Answer:
hello
  hello
hello
Explanation:
- TrimLeft: removes from left only
- TrimRight: removes from right only
- Trim: removes from both sides
*/

// ============================================================================
// SECTION 4: RUNES AND UTF-8 QUESTIONS
// ============================================================================

// Q31. What is a rune in Go?
/*
Answer:
rune is an alias for int32, representing a Unicode code point.

Key points:
- rune = int32
- Represents a single Unicode character
- Can represent any character from any language
- Single quotes create rune literals: 'A', '世', '🚀'

Example:
var r rune = 'A'
fmt.Println(r)        // 65 (Unicode code point)
fmt.Printf("%c\n", r) // A (character)
*/

// Q32. What is the output?
/*
func main() {
    s := "Hello, 世界"
    count := 0
    for range s {
        count++
    }
    fmt.Println(count)
    fmt.Println(len(s))
}

Answer:
9
13
Explanation:
- range iterates over runes (characters): 9 characters
- len() returns bytes: 13 bytes
- "Hello, " = 7 bytes (7 ASCII chars)
- "世界" = 6 bytes (2 chars × 3 bytes each)
*/

// Q33. What is the output?
/*
func main() {
    s := "Go"
    for i, r := range s {
        fmt.Printf("%d: %c\n", i, r)
    }
}

Answer:
0: G
1: o
Explanation: range on string iterates over runes
i is the byte index, r is the rune value
*/

// Q34. What is the output?
/*
func main() {
    s := "世界"
    for i, r := range s {
        fmt.Printf("%d: %c\n", i, r)
    }
}

Answer:
0: 世
3: 界
Explanation:
- First character at byte index 0 (3 bytes long)
- Second character at byte index 3 (3 bytes long)
- range skips to next rune, not next byte
*/

// Q35. How do you properly iterate over a string with Unicode characters?
/*
Answer:
Use range loop (iterates over runes):

Good:
for _, r := range s {
    fmt.Printf("%c ", r)
}

Bad (breaks UTF-8):
for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i])  // Breaks multi-byte characters
}

Alternative (explicit):
import "unicode/utf8"
for i := 0; i < len(s); {
    r, size := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%c ", r)
    i += size
}
*/

// Q36. What is the output?
/*
import "unicode/utf8"

func main() {
    s := "Hello, 世界"
    fmt.Println(utf8.RuneCountInString(s))
    fmt.Println(utf8.ValidString(s))
}

Answer:
9
true
Explanation:
- RuneCountInString counts characters (not bytes)
- ValidString checks if string is valid UTF-8
*/

// ============================================================================
// SECTION 5: STRING BUILDER AND PERFORMANCE
// ============================================================================

// Q37. What is the output and what's wrong with this code?
/*
func main() {
    s := ""
    for i := 0; i < 10000; i++ {
        s += "x"
    }
    fmt.Println(len(s))
}

Answer: 10000
Problem: Very inefficient!
- Each += creates a new string (10000 allocations)
- O(n²) time complexity
- Lots of garbage for GC

Solution: Use strings.Builder
*/

// Q38. How do you efficiently build strings?
/*
Answer: Use strings.Builder

import "strings"

var builder strings.Builder
for i := 0; i < 10000; i++ {
    builder.WriteString("x")
}
s := builder.String()

Benefits:
- Minimizes allocations (grows buffer as needed)
- O(n) time complexity
- Much faster for repeated concatenations

Methods:
- WriteString(s string)
- WriteByte(b byte)
- WriteRune(r rune)
- String() string (get final result)
- Reset() (reuse builder)
- Len() int (current length)
*/

// Q39. What is the output?
/*
func main() {
    var builder strings.Builder
    builder.WriteString("Hello")
    builder.WriteString(" ")
    builder.WriteString("World")
    fmt.Println(builder.String())
}

Answer: Hello World
Explanation: Builder efficiently concatenates strings
*/

// Q40. Will this code compile?
/*
func main() {
    var builder strings.Builder
    builder.WriteString("Hello")
    s := builder.String()
    builder.WriteString(" World")
    fmt.Println(s)
}

Answer: YES - Compiles and runs
Output: Hello
Explanation:
- String() returns current content (doesn't reset builder)
- Can continue writing after calling String()
- s captures "Hello" at that moment
- Further writes don't affect s
*/

// ============================================================================
// SECTION 6: TRIMMING AND SPLITTING ADVANCED
// ============================================================================

// Q41. What is the difference between TrimSpace and Trim?
/*
Answer:
TrimSpace: Removes all leading/trailing whitespace
- Removes: space, tab, newline, carriage return, etc.
- No parameters needed

Trim: Removes specified characters from both ends
- Requires cutset parameter
- Removes any character in the cutset

Example:
s := "  \t\nhello\n\t  "
strings.TrimSpace(s)     // "hello"
strings.Trim(s, " \t\n") // "hello" (same result)

s2 := "xxxhelloxxx"
strings.Trim(s2, "x")    // "hello"
strings.TrimSpace(s2)    // "xxxhelloxxx" (no whitespace)
*/

// Q42. What is the output?
/*
func main() {
    s := "!!!hello!!!"
    fmt.Println(strings.Trim(s, "!"))
    fmt.Println(strings.TrimPrefix(s, "!"))
    fmt.Println(strings.TrimSuffix(s, "!"))
}

Answer:
hello
!!hello!!!
!!!hello!!
Explanation:
- Trim: removes ALL '!' from both ends
- TrimPrefix: removes "!" once from start
- TrimSuffix: removes "!" once from end
*/

// Q43. What is the output?
/*
func main() {
    s := "a,b,c,d,e"
    parts := strings.Split(s, ",")
    fmt.Println(len(parts))
    fmt.Println(parts[0])
    fmt.Println(parts[4])
}

Answer:
5
a
e
Explanation: Split creates slice with 5 elements
*/

// Q44. What is the output?
/*
func main() {
    s := "a,,b"
    parts := strings.Split(s, ",")
    fmt.Println(len(parts))
    fmt.Println(parts[1])
}

Answer:
3
(empty string)
Explanation: Split preserves empty strings between separators
parts = ["a", "", "b"]
*/

// Q45. What is the difference between Split and SplitAfter?
/*
Answer:
Split: Removes separator
SplitAfter: Keeps separator at end of each part

Example:
s := "a,b,c"
strings.Split(s, ",")      // ["a", "b", "c"]
strings.SplitAfter(s, ",") // ["a,", "b,", "c"]

Use case: SplitAfter useful when you need to preserve delimiters
*/

// Q46. What is the output?
/*
func main() {
    s := "one two three four"
    parts := strings.SplitN(s, " ", 2)
    fmt.Println(len(parts))
    fmt.Println(parts[1])
}

Answer:
2
two three four
Explanation: SplitN(s, sep, n) splits into at most n parts
n=2 means: split once, creating 2 parts
*/

// ============================================================================
// SECTION 7: REPLACING AND REPEATING
// ============================================================================

// Q47. What is the output?
/*
func main() {
    s := "hello hello hello"
    fmt.Println(strings.Replace(s, "hello", "hi", 2))
    fmt.Println(strings.ReplaceAll(s, "hello", "hi"))
}

Answer:
hi hi hello
hi hi hi
Explanation:
- Replace with n=2: replaces first 2 occurrences
- ReplaceAll: replaces all occurrences (same as n=-1)
*/

// Q48. What is the output?
/*
func main() {
    s := "hello"
    fmt.Println(strings.Replace(s, "x", "y", -1))
}

Answer: hello
Explanation: If old string not found, returns original unchanged
*/

// Q49. What is the output?
/*
func main() {
    fmt.Println(strings.Repeat("Go", 0))
    fmt.Println(strings.Repeat("Go", 1))
    fmt.Println(strings.Repeat("Go", 3))
}

Answer:
(empty string)
Go
GoGoGo
Explanation:
- n=0: empty string
- n=1: original string
- n=3: repeated 3 times
*/

// Q50. Will this code panic?
/*
func main() {
    s := strings.Repeat("x", -1)
    fmt.Println(s)
}

Answer: YES - Runtime Panic
Error: "strings: negative Repeat count"
Explanation: Repeat count must be >= 0
*/

// ============================================================================
// SECTION 8: COMPARISON AND SEARCHING
// ============================================================================

// Q51. What is the output?
/*
func main() {
    fmt.Println("abc" == "abc")
    fmt.Println("abc" == "ABC")
    fmt.Println("abc" < "xyz")
}

Answer:
true
false
true
Explanation:
- Direct comparison is case-sensitive
- Lexicographical ordering (dictionary order)
*/

// Q52. What is the output?
/*
func main() {
    fmt.Println(strings.Compare("abc", "abc"))
    fmt.Println(strings.Compare("abc", "xyz"))
    fmt.Println(strings.Compare("xyz", "abc"))
}

Answer:
0
-1
1
Explanation:
- 0: equal
- -1: first < second
- 1: first > second
Note: Direct == is preferred for equality
*/

// Q53. When should you use strings.Compare vs ==?
/*
Answer:
Use == for equality:
if s1 == s2 { ... }  // Preferred

Use strings.Compare for ordering:
if strings.Compare(s1, s2) < 0 { ... }  // s1 comes before s2

Or better, use direct comparison:
if s1 < s2 { ... }  // Simpler and clearer

strings.Compare is mainly useful when you need a comparison function
for sorting (sort.Slice, etc.)
*/

// Q54. What is the output?
/*
func main() {
    fmt.Println(strings.Index("hello world", "o"))
    fmt.Println(strings.LastIndex("hello world", "o"))
}

Answer:
4
7
Explanation:
- Index: first occurrence of "o" at position 4
- LastIndex: last occurrence of "o" at position 7
*/

// Q55. What is the output?
/*
func main() {
    fmt.Println(strings.Contains("hello", ""))
    fmt.Println(strings.Index("hello", ""))
}

Answer:
true
0
Explanation: Empty string is contained in every string at position 0
*/

// Q56. What is the output?
/*
func main() {
    fmt.Println(strings.ContainsAny("hello", "xyz"))
    fmt.Println(strings.ContainsAny("hello", "aeiou"))
}

Answer:
false
true
Explanation:
- ContainsAny checks if ANY character from second string is in first
- "hello" contains 'e' and 'o' from "aeiou"
*/

// Q57. What is the output?
/*
func main() {
    fmt.Println(strings.HasPrefix("hello world", "hello"))
    fmt.Println(strings.HasSuffix("hello world", "world"))
    fmt.Println(strings.HasPrefix("hello world", "world"))
}

Answer:
true
true
false
Explanation:
- HasPrefix: checks start of string
- HasSuffix: checks end of string
*/

// ============================================================================
// SECTION 9: RAW STRINGS AND SPECIAL CHARACTERS
// ============================================================================

// Q58. What is the difference between "" and `` (backticks)?
/*
Answer:
Double quotes (""):
- Interpreted string literal
- Escape sequences processed (\n, \t, \\, etc.)
- Cannot span multiple lines (without \n)

Backticks (``):
- Raw string literal
- Escape sequences NOT processed (literal \n)
- Can span multiple lines
- Useful for regex, JSON, SQL, etc.

Example:
s1 := "hello\nworld"   // hello
                        // world
s2 := `hello\nworld`   // hello\nworld (literal)

s3 := `multi
line
string`  // Valid!
*/

// Q59. What is the output?
/*
func main() {
    s1 := "C:\\Users\\Name"
    s2 := `C:\Users\Name`
    fmt.Println(s1)
    fmt.Println(s2)
}

Answer:
C:\Users\Name
C:\Users\Name
Explanation:
- s1: Need to escape backslashes (\\)
- s2: Raw string, no escaping needed
*/

// Q60. What is the output?
/*
func main() {
    s := `Line 1
Line 2
Line 3`
    fmt.Println(len(strings.Split(s, "\n")))
}

Answer: 3
Explanation: Raw string preserves actual newlines
Split on "\n" creates 3 parts
*/

// ============================================================================
// SECTION 10: ERROR IDENTIFICATION QUESTIONS
// ============================================================================

// Q61. Find the error:
/*
func main() {
    s := "hello"
    s[0] = 'H'
    fmt.Println(s)
}

Answer: Compilation Error
Error: "cannot assign to s[0] (strings are immutable)"
Fix: Convert to []byte or []rune first
*/

// Q62. Find the error:
/*
func main() {
    s := "hello"
    fmt.Println(s[10])
}

Answer: Runtime Panic
Error: "runtime error: index out of range [10] with length 5"
Explanation: String has length 5, index 10 doesn't exist
*/

// Q63. Find the logic error:
/*
func main() {
    s := "世界"
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c", s[i])
    }
}

Answer: Logic Error (not compilation error)
Output: Garbled characters
Problem: Iterating by byte breaks UTF-8 multi-byte characters
Fix: Use range loop
*/

// Q64. Find the error:
/*
func main() {
    var s string
    s = nil
    fmt.Println(s)
}

Answer: Compilation Error
Error: "cannot use nil as type string in assignment"
Explanation: Strings cannot be nil (zero value is "")
*/

// Q65. Find the inefficiency:
/*
func buildString() string {
    s := ""
    for i := 0; i < 10000; i++ {
        s += strconv.Itoa(i)
    }
    return s
}

Answer: Performance Issue
Problem: O(n²) complexity, 10000 allocations
Fix: Use strings.Builder
*/

// ============================================================================
// SECTION 11: TRICKY CODE EXECUTION QUESTIONS
// ============================================================================

// Q66. What is the output?
/*
func main() {
    s := "hello"
    s2 := s[1:3]
    s3 := s[1:3]
    fmt.Println(s2 == s3)
}

Answer: true
Explanation: String slices with same content are equal
*/

// Q67. What is the output?
/*
func main() {
    s := ""
    fmt.Println(len(s))
    fmt.Println(s == "")
}

Answer:
0
true
Explanation: Empty string has length 0 and equals ""
*/

// Q68. What is the output?
/*
func main() {
    s := "hello"
    fmt.Println(s[:2])
    fmt.Println(s[2:])
    fmt.Println(s[:])
}

Answer:
he
llo
hello
Explanation:
- s[:2]: from start to index 2 (exclusive)
- s[2:]: from index 2 to end
- s[:]: entire string
*/

// Q69. What is the output?
/*
func main() {
    s := "hello"
    fmt.Println(s[1:1])
}

Answer: (empty string)
Explanation: s[1:1] is valid but produces empty string
Start and end are same index
*/

// Q70. What is the output?
/*
func main() {
    s1 := "hello"
    s2 := "hel" + "lo"
    fmt.Println(s1 == s2)
}

Answer: true
Explanation: String concatenation at compile time
Both result in same string "hello"
*/

// Q71. What is the output?
/*
func main() {
    s := strings.Join([]string{}, ",")
    fmt.Println(len(s))
}

Answer: 0
Explanation: Joining empty slice produces empty string
*/

// Q72. What is the output?
/*
func main() {
    s := strings.Join([]string{"a"}, ",")
    fmt.Println(s)
}

Answer: a
Explanation: Single element, no separator added
*/

// Q73. What is the output?
/*
func main() {
    s := "hello"
    r := []rune(s)
    fmt.Println(len(s))
    fmt.Println(len(r))
}

Answer:
5
5
Explanation: "hello" is ASCII, so bytes = runes
Each character is 1 byte
*/

// Q74. What is the output?
/*
func main() {
    s := "世界"
    r := []rune(s)
    fmt.Println(len(s))
    fmt.Println(len(r))
}

Answer:
6
2
Explanation:
- len(s): 6 bytes (2 chars × 3 bytes)
- len(r): 2 runes (2 characters)
*/

// Q75. What is the output?
/*
func main() {
    fmt.Println(strings.Count("", ""))
}

Answer: 1
Explanation: Empty string contains empty string once
This is a special case in Go's implementation
*/

// ============================================================================
// SECTION 12: CONCEPTUAL & BEST PRACTICES
// ============================================================================

// Q76. Why are strings immutable in Go?
/*
Answer:
Benefits of immutability:
1. Thread Safety: Multiple goroutines can safely read same string
2. Memory Efficiency: Same string literal shared across program
3. Security: Prevents accidental modification
4. Hash Keys: Strings can be safely used as map keys
5. Optimization: Compiler can optimize better

Example:
s1 := "hello"
s2 := s1  // No copy, both point to same data
// Safe because strings can't be modified

Trade-off: To "modify", must create new string (allocation)
*/

// Q77. When should you use []byte vs string?
/*
Answer:
Use string when:
- Read-only text data
- Map keys
- Passing around text
- API boundaries

Use []byte when:
- Need to modify content
- I/O operations (file, network)
- Binary data
- Performance-critical code (avoid conversions)

Example:
// Reading file
data, _ := ioutil.ReadFile("file.txt")  // []byte
text := string(data)  // Convert to string for processing

// Writing file
content := "hello"
ioutil.WriteFile("file.txt", []byte(content), 0644)
*/

// Q78. What is the performance cost of string concatenation?
/*
Answer:
Using + operator:
s := ""
for i := 0; i < n; i++ {
    s += "x"  // O(n²) - creates n new strings
}

Each += operation:
1. Allocates new string
2. Copies old content
3. Appends new content
4. Old string becomes garbage

For n=10000: ~50 million bytes allocated!

Solution: strings.Builder
var b strings.Builder
for i := 0; i < n; i++ {
    b.WriteString("x")  // O(n) - grows buffer efficiently
}
s := b.String()

Benchmark results (n=10000):
- Concatenation: ~500ms, 50MB allocated
- Builder: ~0.5ms, 20KB allocated
- 1000x faster!
*/

// Q79. How do you efficiently check if string contains any of multiple substrings?
/*
Answer:
Method 1: Multiple Contains (simple)
if strings.Contains(s, "foo") || strings.Contains(s, "bar") {
    // Found
}

Method 2: ContainsAny (for single characters)
if strings.ContainsAny(s, "abc") {
    // Contains 'a' or 'b' or 'c'
}

Method 3: Loop (for many substrings)
needles := []string{"foo", "bar", "baz"}
found := false
for _, needle := range needles {
    if strings.Contains(s, needle) {
        found = true
        break
    }
}

Method 4: Regex (complex patterns)
import "regexp"
matched, _ := regexp.MatchString("foo|bar|baz", s)
*/

// Q80. What are common string manipulation pitfalls?
/*
Answer:
1. Modifying strings directly
   Bad: s[0] = 'X'  // Error
   Good: r := []rune(s); r[0] = 'X'; s = string(r)

2. Using len() for character count
   Bad: len("世界") == 2  // False! (6 bytes)
   Good: utf8.RuneCountInString("世界") == 2

3. Iterating by index with UTF-8
   Bad: for i := 0; i < len(s); i++ { ... }
   Good: for _, r := range s { ... }

4. Inefficient concatenation in loops
   Bad: s += "x" (in loop)
   Good: Use strings.Builder

5. Not checking Index return value
   Bad: i := strings.Index(s, "x"); s[i]  // Panic if -1
   Good: if i := strings.Index(s, "x"); i != -1 { s[i] }

6. Case-sensitive comparisons
   Bad: s == "HELLO"
   Good: strings.EqualFold(s, "HELLO")

7. Ignoring empty strings in Split
   Bad: parts := strings.Split("a,,b", ",")  // ["a", "", "b"]
   Good: Filter empty strings if needed
*/

// Q81. How do you reverse a string properly in Go?
/*
Answer:
For ASCII strings:
func reverseASCII(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}

For Unicode strings (proper):
func reverseUnicode(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

Why []rune?
- []byte reversal breaks multi-byte UTF-8 characters
- []rune treats each character as unit

Example:
reverseASCII("hello")  // "olleh" ✓
reverseASCII("世界")   // Garbled ✗
reverseUnicode("世界") // "界世" ✓
*/

// Q82. What is the difference between strings.Title and cases.Title?
/*
Answer:
strings.Title (Deprecated in Go 1.18+):
- Simple word-by-word title casing
- Doesn't handle Unicode properly
- Example: strings.Title("hello world") → "Hello World"

cases.Title (Recommended):
- Proper Unicode title casing
- Language-aware
- Handles special cases correctly

Example:
import "golang.org/x/text/cases"
import "golang.org/x/text/language"

caser := cases.Title(language.English)
caser.String("hello world")  // "Hello World"

// Handles special cases:
caser.String("it's")  // "It's" (not "It'S")

Use cases.Title for production code!
*/

// Q83. How do you check if a string is empty?
/*
Answer:
Method 1: Direct comparison (preferred)
if s == "" {
    // Empty
}

Method 2: Length check
if len(s) == 0 {
    // Empty
}

Both are equivalent and efficient.
Prefer s == "" for clarity.

Note: Strings cannot be nil in Go
var s string  // s == "" (zero value)
*/

// Q84. What is the memory layout of a string in Go?
/*
Answer:
String structure (16 bytes on 64-bit):
type stringStruct struct {
    str unsafe.Pointer  // 8 bytes: pointer to byte array
    len int             // 8 bytes: length in bytes
}

Key points:
1. String header is 16 bytes
2. Actual data stored separately
3. Multiple strings can share same data
4. Immutability allows safe sharing

Example:
s1 := "hello"
s2 := s1[0:3]  // "hel"
// s2 shares data with s1 (no copy)
// Both point to same underlying array

Slicing is O(1) - just creates new header
*/

// Q85. How do you handle case-insensitive string operations?
/*
Answer:
Comparison:
strings.EqualFold(s1, s2)  // Case-insensitive equality

Contains (manual):
strings.Contains(strings.ToLower(s), strings.ToLower(substr))

Better (avoid double conversion):
import "strings"
func containsIgnoreCase(s, substr string) bool {
    s = strings.ToLower(s)
    substr = strings.ToLower(substr)
    return strings.Contains(s, substr)
}

Replace (case-insensitive):
No built-in function, use regexp:
import "regexp"
re := regexp.MustCompile("(?i)old")
result := re.ReplaceAllString(s, "new")

Note: ToLower/ToUpper create new strings (allocations)
*/

// ============================================================================
// END OF INTERVIEW QUESTIONS
// ============================================================================
