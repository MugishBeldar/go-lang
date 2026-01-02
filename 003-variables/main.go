// variable in go

package main // Every Go file must belong to a package. 'main' is special - it's the entry point

import "fmt" 

// 1. basic: variable declaration methods
// go provides 4 ways to declare a variables

//==========================================================================================================
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
	age5 int = 24
)

// 4.2. multiple variable declaration with short declaration operator same as method 4 case
// name6, age6 := "John", 24

// 4.3. multiple variable declaration with short declaration operator and type inference same as method 4 case
// name7, age7 := "John", 24

//==========================================================================================================
// 2. zero values - default values when not initalized
var s string // ""
var b bool   // false
var i int    // 0
var f float64// 0.0
var c complex128// (0+0i)
var p *int   // nil (pointer)
var sl []int // nil (slice)
var m map[string]int // nil (map)
var ch chan int // nil (channel)
var fnc func() // nil (function)
var err error // nil (error)

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
	age8 int = 24 
)

// short declaration - multiple variables (but got an error because we are declaring it outside a function)
// name9, age9 := "john", 24

//==========================================================================================================
// 4. constants - immutable values

const Pi = 3.14159265359
const (
	StatusOk = 200
	StatusNotFound = 404
	StatusInternalServerError = 500
)

// typed vs untyped constants 
const typedInt int = 10 // typed - strictly int  
const untypedInt = 10 // untyped - more flexibles

// behind the scenes: 
// constant are evaluted at compile time, not runtime
// they don't occupy memory at runtime - they're subtituted directly into the code.

//==========================================================================================================
// 5. type inference - how go determines types
var inferredInt = 10 // int
var inferredFloat = 10.0 // float64
var inferredString = "hello" // string
var inferredBool = true // bool

// why float64 and not float32?
// go defaults to larger types for better precision and compatibility with math functions.

//==========================================================================================================
// 6. scope - where variables live
var globalVar = "I'm accessible anywhere in this package" // package scope

func main () {
	var functionVar = "I'm ony accessible in this main function" // function scope
	if true { 
		blockVar := "I'm only accessible in this block" // block scope
		fmt.Println(blockVar) // it is works 
	}
	// fmt.Println(blockVar) // it is not works because block scope
}

// scope hierarchy:
// Universe (built-in: true, false, nil, int, string...)
//    └── Package (var outside functions)
//           └── Function (var inside func)
//                  └── Block (var inside if/for/switch)

//==========================================================================================================
// 7. shadowing - same name different scope
func shadowing() {
	x := 10 
	fmt.Println(x) // 10
	if true {
		x := 20 // new variable in this block which is shadow the outer x 
		fmt.Println(x) // 20 
	}
	fmt.Println(x) // 10
}


//==========================================================================================================
// 8. blank identifier - ignore values 
func someFunction() (int, int) {
	return 1, 2
}

func blankIdentifier() {
	// ignore return value 
	value, _ := someFunction() // _ is a blank identifier
	println(value)
}

//==========================================================================================================
// 9. pointer - reference to memory address 
func pointer() {
	x := 10
	p := &x // p holds the memory address of x 
	fmt.Println(*p) // 10 (dereference - get value at address) 
	*p = 20 // change value at address
	fmt.Println(x) // 20
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
	x := 10 // x lives on stack (local, dosen't escape) 
	return x // value copied, x destroyed when function ends
}

func heapAllocation() *int {
	x := 10 // x escapes to heap (returned as pointer)
	return &x  // go detects this and allocates on heap
}

// excape analytsis: 
// go compiler analyze if a variable "escapes" the function. if it does, it's allocated on the heap.

// variable memory size 
func variableMemorySize() {
	var x int // 8 bytes (64-bit system)
	var y float64 // 8 bytes
	var z string // 16 bytes (pointer to array of bytes + length)
	var a []int // 24 bytes (pointer to array of int + length + capacity)
	var b map[string]int // 8 bytes (pointer to hash table)
	var c chan int // 8 bytes (pointer to channel data)
	var d bool // 1 byte (0 or 1)
	var e byte // 1 byte (0 to 255)
	var f rune // 4 bytes (Unicode code point)
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
	x = 20 // reassign with same type is ok
	// x = "hello" // reassign with different type is not ok
	
	// := with at least one new variables
	x, y := 10, 20 // x is int, y is int
	x, z := 10, "hello" // x is int, z is string
}

// type conversion (no implicit conversion)
func typeConversion() {
	var i = 10;
	var f float64= float64(i) // explicit conversion
	// var f = i // implicit conversion is not allowed
}