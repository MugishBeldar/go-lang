// package main
// - every go program starts with a package declaration
// - main is a special package that tells go this is an executable program
package main;

// import fmt package from go's standard library
// - provides I/O functions like Println
import "fmt";

// func main()
// - entry point of the program
func main() {
	// print hello world to the console with a newline
	fmt.Println("Hello, World!");
}


// what happens when you run go run hello.go?
// - go run is convenience command that compiles and runs the program in one step.
// - it's actually doing two things:
// 1. compilling: your code in binary executable
// 2. running: the executabale binary file immediately
// 3. cleaning up: deleting the binary file after running it 

// detailed step-by-step process

// step 1: source code parsing
// hello.go ---> lexer ---> tokens ---> parser ---> AST (abstract syntax tree)
// go reads your .go file character by character
// lexer: breaks it to tokens: package, main, import, "fmt" etc.
// parser: builds an abstract syntax tree (AST) represents the structure of your code
// type checker: validated types and ensures code corrections

// step 2: compilation
// AST ---> compiler ---> assembly code ---> assembler ---> object files ---> executable
// compiler ---> converts ast to machine specific assembly code
// assembler ---> converts assembly code to machine code (object files)
// linker ---> combines object files into a single executable file

// step 3: execution
// executable ---> OS ---> CPU
// OS loads the executable into memory
// CPU executes the machine code instructions

// step 4: cleanup
// go run automatically deletes the binary file after running it
// this is why we don't see the binary file in the current directory when we use go run cmd


// go run and go build comparision
// go run:
// What it does:
// 1. Compile to temporary location
// 2. Execute immediately
// 3. Delete executable
// 4. Exit

// go build:
// What it does:
// 1. Compile to current directory
// 2. Create executable named 'hello'
// 3. Keep the executable
// 4. Exit (doesn't run it)

