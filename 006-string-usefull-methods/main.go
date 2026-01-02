package main

import "fmt"
import "strings"
// For proper title casing, use:
import "golang.org/x/text/cases"
import "golang.org/x/text/language"


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
	fmt.Println("trim space: ", strings.TrimSpace(s4)) // "hello world"
	fmt.Println("trim left space: ", strings.TrimLeft(s4, " ")) // "hello world  "
	fmt.Println("trim right space: ", strings.TrimRight(s4, " ")) // "  hello world"
	fmt.Println("trim prefix: ", strings.TrimPrefix(s4, "  ")) // "hello world  "
	fmt.Println("trim suffix: ", strings.TrimSuffix(s4, "  ")) // "  hello world"
	fmt.Println("trim: ", strings.Trim(s4, " ")) // "hello world"
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
	fmt.Println("replace: ", strings.Replace(s1, "the", "a", 1)) // "a string package" (replace first occurrence)d	
	fmt.Println("replace: ", strings.Replace(s1, "the", "a", -1)) // "a string package" (replace all occurrences)
	fmt.Println("replace all: ", strings.ReplaceAll(s1, "the", "a")) // "a string package" (replace all occurrences)
	fmt.Println("repeat: ", strings.Repeat(s1, 3)) // "the string packagethe string packagethe string package" (repeat 3 times)

	// 7. splitting
	fmt.Println("============================= SPLITTING =============================")
	fmt.Println("split: ", strings.Split(s1, " ")) // ["the", "string", "package"] (split on space)
	fmt.Println("split after: ", strings.SplitAfter(s1, " ")) // ["the ", "string ", "package"] (split after space)
	fmt.Println("split after n: ", strings.SplitAfterN(s1, " ", 2)) // ["the ", "string package"] (split after space, max 2 parts)
	fmt.Println("split n: ", strings.SplitN(s1, " ", 2)) // ["the", "string package"] (split on space, max 2 parts)

	// 8. padding
	fmt.Println("============================= PADDING =============================")
	fmt.Println("pad left: ", strings.Repeat(" ", 5) + s1) // "     the string package" (pad left with 5 spaces)
	fmt.Println("pad right: ", s1 + strings.Repeat(" ", 5)) // "the string package     " (pad right with 5 spaces)
	fmt.Println("pad both: ", strings.Repeat(" ", 2) + s1 + strings.Repeat(" ", 2)) // "  the string package  " (pad both sides with 2 spaces)

	// 9. comparing
	fmt.Println("============================= COMPARING =============================")
	fmt.Println("compare: ", strings.Compare(s1, "the string package")) // 0 (equal)
	fmt.Println("compare: ", strings.Compare(s1, "The string package")) // 1 (s1 is greater)
	fmt.Println("compare: ", strings.Compare(s1, "the string packages")) // -1 (s1 is less)
	fmt.Println("compare: ", strings.EqualFold("go", "GO")) // true (case insensitive)
	fmt.Println("compare: ", strings.EqualFold("go", "gO")) // true (case insensitive)
	fmt.Println("compare: ", strings.EqualFold("go", "go")) // true (case insensitive)
	fmt.Println("compare: ", strings.EqualFold("go", "golang")) // false (case insensitive)

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

