// variable in go

package main // Every Go file must belong to a package. 'main' is special - it's the entry point

import "fmt"

// 1. basic: variable declaration methods
// go provides 4 ways to declare a variables

// ==========================================================================================================
// method 1: "var" with type and value
var name string = "John" // type is optional if you initialize with a value go infer it automatically
var age int = 24

// method 2: "var" with type only (uses zero value)
// zero value means: default value for the types wich is given below
// string → "" (empty string)
// bool   → false
// int    → 0
// float  → 0.0
// complex → (0+0i)
var name2 string
var age2 int

// method 3: "var" with value only (type inferred)
var name3 = "John"
var age3 = 24

// method 4: short declaration operator := (type inferred, only in function scope so uncomment below code we got an error)
// name4 := "John"
// age4 := 24

// 4.1. multiple variable declaration
var (
	name5 string = "John"
	age5  int    = 24
)

// 4.2. multiple variable declaration with short declaration operator same as method 4 case
// name6, age6 := "John", 24

// 4.3. multiple variable declaration with short declaration operator and type inference same as method 4 case
// name7, age7 := "John", 24

// ==========================================================================================================
// 2. zero values - default values when not initalized
var s string         // ""
var b bool           // false
var i int            // 0
var f float64        // 0.0
var c complex128     // (0+0i)
var p *int           // nil (pointer)
var sl []int         // nil (slice)
var m map[string]int // nil (map)
var ch chan int      // nil (channel)
var fnc func()       // nil (function)
var err error        // nil (error)

// behind the scenes:
// go automatically initalizes all memory to zero bytes.
// this prevents "undefined behavior" bugs common in c/c++

//==========================================================================================================
// 3. multiple variable declaration

// same type
var one, two, three int = 1, 2, 3

// different types
var (
	name8 string = "john"
	age8  int    = 24
)

// short declaration - multiple variables (but got an error because we are declaring it outside a function)
// name9, age9 := "john", 24

//==========================================================================================================
// 4. constants - immutable values

const Pi = 3.14159265359
const (
	StatusOk                  = 200
	StatusNotFound            = 404
	StatusInternalServerError = 500
)

// typed vs untyped constants
const typedInt int = 10 // typed - strictly int
const untypedInt = 10   // untyped - more flexibles

// behind the scenes:
// constant are evaluted at compile time, not runtime
// they don't occupy memory at runtime - they're subtituted directly into the code.

// ==========================================================================================================
// 5. type inference - how go determines types
var inferredInt = 10         // int
var inferredFloat = 10.0     // float64
var inferredString = "hello" // string
var inferredBool = true      // bool

// why float64 and not float32?
// go defaults to larger types for better precision and compatibility with math functions.

// ==========================================================================================================
// 6. scope - where variables live
var globalVar = "I'm accessible anywhere in this package" // package scope

func main() {
	var functionVar = "I'm ony accessible in this main function" // function scope
	if true {
		blockVar := "I'm only accessible in this block" // block scope
		fmt.Println(blockVar)                           // it is works
	}
	// fmt.Println(blockVar) // it is not works because block scope
}

// scope hierarchy:
// Universe (built-in: true, false, nil, int, string...)
//    └── Package (var outside functions)
//           └── Function (var inside func)
//                  └── Block (var inside if/for/switch)

// ==========================================================================================================
// 7. shadowing - same name different scope
func shadowing() {
	x := 10
	fmt.Println(x) // 10
	if true {
		x := 20        // new variable in this block which is shadow the outer x
		fmt.Println(x) // 20
	}
	fmt.Println(x) // 10
}

// ==========================================================================================================
// 8. blank identifier - ignore values
func someFunction() (int, int) {
	return 1, 2
}

func blankIdentifier() {
	// ignore return value
	value, _ := someFunction() // _ is a blank identifier
	println(value)
}

// ==========================================================================================================
// 9. pointer - reference to memory address
func pointer() {
	x := 10
	p := &x         // p holds the memory address of x
	fmt.Println(*p) // 10 (dereference - get value at address)
	*p = 20         // change value at address
	fmt.Println(x)  // 20
}

// behind the scenes deep dive:
// memory layout
// ┌─────────────────────────────────────────────┐
// │                   STACK                     │
// │  - Function local variables                 │
// │  - Fast allocation/deallocation             │
// │  - Automatic memory management              │
// │  - Fixed size per goroutine (~2KB initial)  │
// └─────────────────────────────────────────────┘
//                      │
//                      ▼
// ┌─────────────────────────────────────────────┐
// │                    HEAP                     │
// │  - Variables that "escape" the function     │
// │  - Dynamically sized data (slices, maps)    │
// │  - Managed by Garbage Collector             │
// │  - Slower than stack                        │
// └─────────────────────────────────────────────┘

// stack vs heap allocation
func stackAllocation() int {
	x := 10  // x lives on stack (local, dosen't escape)
	return x // value copied, x destroyed when function ends
}

func heapAllocation() *int {
	x := 10   // x escapes to heap (returned as pointer)
	return &x // go detects this and allocates on heap
}

// excape analytsis:
// go compiler analyze if a variable "escapes" the function. if it does, it's allocated on the heap.

// variable memory size
func variableMemorySize() {
	var x int            // 8 bytes (64-bit system)
	var y float64        // 8 bytes
	var z string         // 16 bytes (pointer to array of bytes + length)
	var a []int          // 24 bytes (pointer to array of int + length + capacity)
	var b map[string]int // 8 bytes (pointer to hash table)
	var c chan int       // 8 bytes (pointer to channel data)
	var d bool           // 1 byte (0 or 1)
	var e byte           // 1 byte (0 to 255)
	var f rune           // 4 bytes (Unicode code point)
	fmt.Println(x, y, z, a, b, c, d, e, f)
}

// string internal structure:
// string = {
// 	ptr *byte (pointer to byte array (8 bytes))
// 	len int (length of string in bytes (8 bytes))
// }
// total: 16 bytes on 64-bit system

// Reassignment rules
func reassignmentRules() {
	x := 10 // x is int
	x = 20  // reassign with same type is ok
	// x = "hello" // reassign with different type is not ok

	// := with at least one new variables
	x, y := 10, 20      // x is int, y is int
	x, z := 10, "hello" // x is int, z is string
}

// type conversion (no implicit conversion)
func typeConversion() {
	var i = 10
	var f float64 = float64(i) // explicit conversion
	// var f = i // implicit conversion is not allowed
}

// ============================================================================
// INTERVIEW QUESTIONS - VARIABLES IN GO
// ============================================================================

// ============================================================================
// SECTION 1: BASIC LEVEL QUESTIONS
// ============================================================================

// Q1. What are the different ways to declare a variable in Go?
/*
Answer: There are 4 main ways:
1. var name string = "John"        // with type and value
2. var name string                 // with type only (zero value)
3. var name = "John"               // with value only (type inferred)
4. name := "John"                  // short declaration (only inside functions)

Additional:
5. var (name string = "John"; age int = 24)  // grouped declaration
*/

// Q2. Will this code compile? If not, what error?
/*
package main

var x := 10

func main() {
    println(x)
}

Answer: NO - Compilation Error
Error: "syntax error: unexpected :=, expecting ="
Explanation: Short declaration operator := can only be used inside functions.
At package level, use: var x = 10
*/

// Q3. What is the output?
/*
func main() {
    var x int
    var y string
    var z bool
    fmt.Printf("%d %q %v\n", x, y, z)
}

Answer: 0 "" false
Explanation: These are zero values - default values for uninitialized variables
*/

// Q4. Will this code compile?
/*
func main() {
    var x int = 10
    x = 20
    fmt.Println(x)
}

Answer: YES
Output: 20
Explanation: Variables can be reassigned with values of the same type
*/

// Q5. Will this code compile? What error?
/*
func main() {
    const x = 10
    x = 20
    fmt.Println(x)
}

Answer: NO - Compilation Error
Error: "cannot assign to x (declared const)"
Explanation: Constants are immutable and cannot be reassigned
*/

// Q6. What is the output?
/*
func main() {
    x := 10
    y := 20
    x, y = y, x
    fmt.Println(x, y)
}

Answer: 20 10
Explanation: Simultaneous assignment - swaps values without temporary variable
*/

// Q7. Will this code compile?
/*
func main() {
    var x int = 10
    var y float64 = x
    fmt.Println(y)
}

Answer: NO - Compilation Error
Error: "cannot use x (type int) as type float64 in assignment"
Explanation: Go does not allow implicit type conversion. Use: var y float64 = float64(x)
*/

// Q8. What is the difference between var and const?
/*
Answer:
var:
- Mutable (can be changed)
- Can be declared without initialization
- Evaluated at runtime
- Can be any type

const:
- Immutable (cannot be changed)
- Must be initialized at declaration
- Evaluated at compile time
- Limited to basic types (numbers, strings, booleans)
*/

// ============================================================================
// SECTION 2: INTERMEDIATE LEVEL QUESTIONS
// ============================================================================

// Q9. Will this code compile? What's the output or error?
/*
func main() {
    x := 10
    x := 20
    fmt.Println(x)
}

Answer: NO - Compilation Error
Error: "no new variables on left side of :="
Explanation: := declares a new variable. x is already declared.
Use x = 20 for reassignment.
*/

// Q10. Will this code compile? What's the output?
/*
func main() {
    x := 10
    x, y := 20, 30
    fmt.Println(x, y)
}

Answer: YES - Compiles and runs
Output: 20 30
Explanation: := is allowed if at least ONE variable is new (y is new here).
x is reassigned, y is declared.
*/

// Q11. What is the output?
/*
func main() {
    x := 10
    if true {
        x := 20
        fmt.Println(x)
    }
    fmt.Println(x)
}

Answer:
20
10
Explanation: Variable shadowing. Inner x is a different variable that shadows outer x.
*/

// Q12. Will this code compile?
/*
func main() {
    var x int
    var x string
    fmt.Println(x)
}

Answer: NO - Compilation Error
Error: "x redeclared in this block"
Explanation: Cannot declare the same variable name twice in the same scope
*/

// Q13. What is the output?
/*
func main() {
    var s string
    var sl []int
    var m map[string]int
    var p *int
    fmt.Printf("%v %v %v %v\n", s == "", sl == nil, m == nil, p == nil)
}

Answer: true true true true
Explanation: Zero values - empty string for string, nil for slice, map, and pointer
*/

// Q14. Will this code compile? What's the output?
/*
const x = 10
const y = 20.5

func main() {
    const z = x + y
    fmt.Println(z)
}

Answer: YES - Compiles and runs
Output: 30.5
Explanation: Untyped constants can be mixed. Compiler determines type at compile time.
*/

// Q15. What is the output?
/*
func main() {
    _, x := 10, 20
    fmt.Println(x)
}

Answer: 20
Explanation: Blank identifier _ discards the first value (10)
*/

// Q16. Will this code compile?
/*
func main() {
    _ := 10
    fmt.Println(_)
}

Answer: NO - Compilation Error
Error: "cannot use _ as value"
Explanation: Blank identifier _ cannot be used as a value, only for discarding
*/

// Q17. What is the output?
/*
var x = 10

func main() {
    fmt.Println(x)
    x := 20
    fmt.Println(x)
}

Answer:
10
20
Explanation: First prints package-level x (10), then declares local x (20) which shadows it
*/

// Q18. Will this code compile?
/*
func main() {
    var x int = 10
    var y int = 20
    var z = x + y
    fmt.Println(z)
}

Answer: YES - Compiles and runs
Output: 30
Explanation: Type inference determines z is int based on x + y
*/

// ============================================================================
// SECTION 3: ADVANCED LEVEL QUESTIONS
// ============================================================================

// Q19. What is the output and why?
/*
func main() {
    x := 10
    p := &x
    *p = 20
    fmt.Println(x)
}

Answer: 20
Explanation: p is a pointer to x. *p = 20 modifies the value at x's address.
*/

// Q20. Will this code compile? What happens?
/*
func getPointer() *int {
    x := 10
    return &x
}

func main() {
    p := getPointer()
    fmt.Println(*p)
}

Answer: YES - Compiles and runs
Output: 10
Explanation: Escape analysis! Go detects x escapes the function and allocates it on heap.
This is safe in Go (unlike C where it would be undefined behavior).
*/

// Q21. What is the output?
/*
func main() {
    var x *int
    fmt.Println(x)
    fmt.Println(*x)
}

Answer: First line prints: <nil>
Second line: Runtime Panic - "runtime error: invalid memory address or nil pointer dereference"
Explanation: x is nil pointer. Dereferencing nil pointer causes panic.
*/

// Q22. Will this code compile?
/*
const x int = 10

func main() {
    const y = x * 2
    fmt.Println(y)
}

Answer: YES - Compiles and runs
Output: 20
Explanation: Constants can be used in other constant expressions at compile time
*/

// Q23. What is the output?
/*
func main() {
    const x = 10
    var y int = x
    var z float64 = x
    fmt.Printf("%T %T\n", y, z)
}

Answer: int float64
Explanation: Untyped constant x can be assigned to different types.
Compiler converts it appropriately.
*/

// Q24. Will this code compile? What error?
/*
const x int = 10

func main() {
    var y float64 = x
    fmt.Println(y)
}

Answer: NO - Compilation Error
Error: "cannot use x (type int) as type float64 in assignment"
Explanation: Typed constant (const x int) cannot be implicitly converted.
Use: var y float64 = float64(x)
*/

// Q25. What is the output?
/*
func main() {
    x, y := 10, 20
    x, y, z := 30, 40, 50
    fmt.Println(x, y, z)
}

Answer: 30 40 50
Explanation: := is valid because z is a new variable. x and y are reassigned.
*/

// Q26. What happens in this code?
/*
func main() {
    var x int
    fmt.Println(x)
    x++
    fmt.Println(x)
}

Answer: Compiles and runs
Output:
0
1
Explanation: x starts with zero value 0, then incremented to 1
*/

// Q27. What is the output?
/*
func modify(x int) {
    x = 100
}

func main() {
    x := 10
    modify(x)
    fmt.Println(x)
}

Answer: 10
Explanation: Go is pass-by-value. modify() receives a copy of x, not the original.
Changes to the copy don't affect the original.
*/

// Q28. What is the output?
/*
func modify(x *int) {
    *x = 100
}

func main() {
    x := 10
    modify(&x)
    fmt.Println(x)
}

Answer: 100
Explanation: Passing pointer allows modify() to change the original value.
*x = 100 modifies the value at x's address.
*/

// Q29. Will this code compile?
/*
func main() {
    var x int = 10
    var y *int = x
    fmt.Println(*y)
}

Answer: NO - Compilation Error
Error: "cannot use x (type int) as type *int in assignment"
Explanation: Cannot assign int to *int. Use: var y *int = &x
*/

// Q30. What is the output?
/*
func main() {
    x := new(int)
    fmt.Println(x)
    fmt.Println(*x)
}

Answer:
0xc000012345 (some memory address)
0
Explanation: new(int) allocates memory for int, returns pointer.
The allocated int has zero value 0.
*/

// ============================================================================
// SECTION 4: SCOPE AND SHADOWING QUESTIONS
// ============================================================================

// Q31. What is the output?
/*
var x = "global"

func main() {
    fmt.Println(x)
    x = "modified"
    fmt.Println(x)
}

Answer:
global
modified
Explanation: Package-level variable can be accessed and modified in main()
*/

// Q32. What is the output?
/*
var x = "global"

func main() {
    x := "local"
    fmt.Println(x)
}

Answer: local
Explanation: Local x shadows global x. := creates new variable in function scope.
*/

// Q33. What is the output?
/*
func main() {
    x := 10
    {
        x := 20
        fmt.Println(x)
    }
    fmt.Println(x)
}

Answer:
20
10
Explanation: Inner block creates new x that shadows outer x
*/

// Q34. What is the output?
/*
func main() {
    x := 10
    {
        x = 20
        fmt.Println(x)
    }
    fmt.Println(x)
}

Answer:
20
20
Explanation: No := in inner block, so it modifies the outer x (no shadowing)
*/

// Q35. Will this code compile?
/*
func main() {
    if x := 10; x > 5 {
        fmt.Println(x)
    }
    fmt.Println(x)
}

Answer: NO - Compilation Error
Error: "undefined: x"
Explanation: x is scoped to the if block only. Not accessible outside.
*/

// Q36. What is the output?
/*
func main() {
    if x := 10; x > 5 {
        fmt.Println(x)
    } else {
        fmt.Println(x)
    }
}

Answer:
10
Explanation: x is accessible in both if and else blocks
*/

// Q37. What is the output?
/*
func main() {
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
    fmt.Println(i)
}

Answer: Compilation Error
Error: "undefined: i"
Explanation: i is scoped to the for loop only
*/

// ============================================================================
// SECTION 5: MEMORY AND PERFORMANCE QUESTIONS
// ============================================================================

// Q38. Which variable will be allocated on the heap?
/*
func stackVar() int {
    x := 10
    return x
}

func heapVar() *int {
    x := 10
    return &x
}

Answer: heapVar's x will be on heap
Explanation:
- stackVar: x doesn't escape, allocated on stack, destroyed after return
- heapVar: x escapes (returned as pointer), allocated on heap by escape analysis
*/

// Q39. What is the memory size of a string variable on a 64-bit system?
/*
Answer: 16 bytes
Explanation: String structure contains:
- Pointer to byte array: 8 bytes
- Length (int): 8 bytes
Total: 16 bytes

Note: This is the string header size, not the actual string data size.
*/

// Q40. What is the output?
/*
func main() {
    var x int
    var y int32
    fmt.Println(unsafe.Sizeof(x))
    fmt.Println(unsafe.Sizeof(y))
}

Answer (on 64-bit system):
8
4
Explanation:
- int is platform-dependent (8 bytes on 64-bit)
- int32 is always 4 bytes
*/

// ============================================================================
// SECTION 6: TRICKY CODE EXECUTION QUESTIONS
// ============================================================================

// Q41. What is the output?
/*
func main() {
    x, y := 10, 20
    fmt.Println(x, y)
    x, y = y, x
    fmt.Println(x, y)
}

Answer:
10 20
20 10
Explanation: Simultaneous assignment swaps values
*/

// Q42. Will this code compile?
/*
func main() {
    var x, y int = 10, 20
    fmt.Println(x, y)
}

Answer: YES - Compiles and runs
Output: 10 20
Explanation: Multiple variable declaration with same type
*/

// Q43. Will this code compile?
/*
func main() {
    var x, y = 10, "hello"
    fmt.Println(x, y)
}

Answer: YES - Compiles and runs
Output: 10 hello
Explanation: Type inference allows different types in multiple declaration
*/

// Q44. What is the output?
/*
func main() {
    a, b := 1, 2
    a, b = b, a+b
    fmt.Println(a, b)
}

Answer: 2 3
Explanation: Right side evaluated first (b=2, a+b=3), then assigned to left side
*/

// Q45. Will this code compile?
/*
func main() {
    x := 10
    y := &x
    z := &y
    fmt.Println(***z)
}

Answer: NO - Compilation Error
Error: "invalid indirect of z (type **int)"
Explanation: z is **int (pointer to pointer). Need **z, not ***z
Correct: fmt.Println(**z)
*/

// Q46. What is the output?
/*
func main() {
    x := 10
    y := &x
    z := &y
    fmt.Println(**z)
}

Answer: 10
Explanation: z is **int. **z dereferences twice to get the value 10
*/

// Q47. Will this code compile?
/*
const (
    a = iota
    b
    c
)

func main() {
    fmt.Println(a, b, c)
}

Answer: YES - Compiles and runs
Output: 0 1 2
Explanation: iota is a constant generator that increments (0, 1, 2, ...)
*/

// Q48. What is the output?
/*
const (
    a = iota + 1
    b
    c
)

func main() {
    fmt.Println(a, b, c)
}

Answer: 1 2 3
Explanation: iota starts at 0, so a=0+1=1, b=1+1=2, c=2+1=3
*/

// ============================================================================
// SECTION 7: ERROR IDENTIFICATION QUESTIONS
// ============================================================================

// Q49. Find the error:
/*
func main() {
    const x int
    x = 10
    fmt.Println(x)
}

Answer: Compilation Error
Error: "const declaration cannot have type without expression"
Explanation: Constants must be initialized at declaration.
Use: const x int = 10
*/

// Q50. Find the error:
/*
func main() {
    var x int = 10
    var x int = 20
    fmt.Println(x)
}

Answer: Compilation Error
Error: "x redeclared in this block"
Explanation: Cannot redeclare variable in same scope
*/

// Q51. Find the error:
/*
func main() {
    x := 10
    var x int = 20
    fmt.Println(x)
}

Answer: Compilation Error
Error: "x redeclared in this block"
Explanation: x is already declared with :=, cannot redeclare with var
*/

// Q52. Find the error:
/*
func main() {
    var x int = 10
    x := 20
    fmt.Println(x)
}

Answer: Compilation Error
Error: "no new variables on left side of :="
Explanation: x already exists, cannot use := for reassignment
*/

// Q53. Find the error:
/*
func main() {
    const x = 10
    y := &x
    fmt.Println(*y)
}

Answer: Compilation Error
Error: "cannot take the address of x"
Explanation: Cannot take address of constant. Constants don't have memory addresses.
*/

// Q54. Find the error:
/*
var x int = getValue()

func getValue() int {
    return 10
}

func main() {
    fmt.Println(x)
}

Answer: Compilation Error
Error: "x initialized and not used" OR function call in package-level var
Explanation: Package-level variables can only be initialized with constants or
constant expressions, not function calls (unless in init() function).
*/

// Q55. Find the error:
/*
func main() {
    var x *int
    *x = 10
    fmt.Println(*x)
}

Answer: Runtime Panic
Error: "runtime error: invalid memory address or nil pointer dereference"
Explanation: x is nil pointer. Must allocate memory first:
    x = new(int) or var y int; x = &y
*/

// ============================================================================
// SECTION 8: CONCEPTUAL & BEST PRACTICES QUESTIONS
// ============================================================================

// Q56. When should you use := vs var?
/*
Answer:
Use := when:
- Inside functions
- Type is obvious from the value
- You want concise code
- Example: x := 10, name := "John"

Use var when:
- At package level (required)
- You want to declare without initializing (zero value)
- You want explicit type for clarity
- Type differs from default inference
- Example: var x int64 = 10 (not int), var count int (zero value)

Best practice: Prefer := inside functions for clarity and brevity
*/

// Q57. What is the difference between typed and untyped constants?
/*
Answer:
Typed constant:
const x int = 10
- Has explicit type
- Cannot be assigned to different types without conversion
- More restrictive

Untyped constant:
const x = 10
- No explicit type
- Can be assigned to compatible types
- More flexible
- Compiler determines type based on usage

Example:
const typed int = 10
const untyped = 10

var a int = untyped      // OK
var b float64 = untyped  // OK
var c int = typed        // OK
var d float64 = typed    // ERROR - needs conversion
*/

// Q58. Explain escape analysis and its impact on performance
/*
Answer:
Escape analysis is a compiler optimization that determines whether a variable
should be allocated on the stack or heap.

Stack allocation:
- Fast (just move stack pointer)
- Automatic cleanup (when function returns)
- Limited size
- Variables that don't "escape" the function

Heap allocation:
- Slower (involves memory allocator)
- Garbage collector cleanup
- Unlimited size (within system memory)
- Variables that "escape" (returned pointers, captured by closures, etc.)

Performance impact:
- Stack allocation is ~10-100x faster
- Heap allocation creates GC pressure
- Too many heap allocations can slow down program

How to check:
go build -gcflags='-m' main.go
Shows which variables escape to heap

Best practice: Avoid unnecessary pointer returns when possible
*/

// Q59. What are the dangers of variable shadowing?
/*
Answer:
Dangers:
1. Bugs from unintended variable usage
2. Hard to debug (code looks correct)
3. Confusion about which variable is being used
4. Common in error handling

Example of dangerous shadowing:
func readFile() error {
    var err error
    file, err := os.Open("file.txt")  // OK
    if err != nil {
        return err
    }

    data, err := ioutil.ReadAll(file)  // Shadows err!
    if err != nil {
        return err  // Returns the ReadAll error
    }

    // Original err is shadowed, might miss Open errors
    return err  // Which err?
}

Best practice:
- Use different variable names
- Be careful with := in nested scopes
- Use linters to detect shadowing (go vet, golangci-lint)
*/

// Q60. What is the zero value concept and why is it important?
/*
Answer:
Zero value is the default value assigned to variables declared without
explicit initialization.

Zero values:
- bool: false
- int, int8, int16, int32, int64: 0
- uint, uint8, uint16, uint32, uint64: 0
- float32, float64: 0.0
- complex64, complex128: 0+0i
- string: ""
- pointer, slice, map, channel, function, interface: nil

Importance:
1. Safety: No uninitialized variables (unlike C/C++)
2. Predictability: Always know the starting state
3. Useful defaults: Many types are usable at zero value
   - var buf bytes.Buffer (ready to use)
   - var mu sync.Mutex (ready to use)
4. Simplifies code: No need to initialize everything

Example:
var count int  // 0, ready to use
count++        // Now 1

Best practice: Design types so zero value is useful (like sync.Mutex)
*/

// ============================================================================
// END OF INTERVIEW QUESTIONS
// ============================================================================
