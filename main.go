package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var x = 42
var y = "James Bond"
var z = true

type phoneNumber int

// iota :) - Works only with const
const (
	ia = iota
	ib = iota
	ic = iota
)

// iota :) - Valid only for one scope
const (
	iaa = iota
	ibb
	icc
)

// iota with bit shifting
const (
	kb = 1 << 10
	mb = 1 << 20
	gb = 1 << 30
)

const (
	_   = iota             // now 1
	ikb = 1 << (iota * 10) // now 2
	imb = 1 << (iota * 10) // now 3
	igb = 1 << (iota * 10) // now 4
)

func fizzBuzz(n int) {
	i := 0
	for i < n {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}

func getSum(a int, b int, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		} else {
			continue
		}
	}
	return sum
}

func main() {
	fmt.Println("Hello!", 1>>100)

	// Starts here!
	// toPrint := fmt.Sprintf("%d %s %t", x, y, z)
	// Could also use
	toPrint := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println("Here we go", toPrint)

	// PhoneNumber
	var a phoneNumber = 9876543210
	fmt.Printf("Phone Type: %T", a)
	fmt.Println("Phone:", a)
	a = 9876543211
	fmt.Println("New Phone:", a)
	fmt.Printf("After conversion: %v and type is %T\n", int(a), int(a))

	// Consts
	const ab int = 10
	const bc = 20
	// fmt.Printf("Consts: %v\t%v\n", ab, bc)
	// ab = 23 // You can't reassign values for const
	fmt.Printf("Consts: %v\t%v\n", ab, bc)

	// Bytes
	random := "A"
	fmt.Printf("Byte Ascii: %d\n", []byte(random))
	fmt.Printf("Binary: %b\n", []byte(random)[0])

	// iota :)
	fmt.Printf("ia\t%v\n", ia)
	fmt.Printf("ib\t%v\n", ib)
	fmt.Printf("ic\t%v\n", ic)
	fmt.Printf("iaa\t%v\n", iaa)
	fmt.Printf("ibb\t%v\n", ibb)
	fmt.Printf("icc\t%v\n", icc)

	// iota with bit shifting
	fmt.Printf("kb\t%v\n", kb)
	fmt.Printf("mb\t%v\n", mb)
	fmt.Printf("gb\t%v\n", gb)
	fmt.Printf("ikb\t%v\n", ikb)
	fmt.Printf("imb\t%v\n", imb)
	fmt.Printf("igb\t%v\n", igb)

	// print ASCII
	fmt.Println("Print ASCII")
	for i := 98; i < 122; i++ {
		fmt.Printf("%v\t%v\t%#U\t%v\n", i, string(i), i, strconv.Itoa(i))
	}

	// If conditionals
	fmt.Println("If conditionals")
	if xy := 3; xy != 3 { // Scope is restricted to this if conditional
		fmt.Println(xy)
	} else if xy == 3 {
		fmt.Println("Valid here as well!")
	}

	// Switch Case
	fmt.Println("Switch Case")
	switch { // A missing switch expression is equal to true, so it searches for true in a case
	case (1 == 2):
		fmt.Println("1 == 2")
	case (1 != 2):
		fmt.Println("1 != 2")
		fallthrough
	case (2 != 3): // quits here since there is no fallthrough
		fmt.Println("2 != 3")
	case (3 != 4):
		fmt.Println("3 != 4")
		fallthrough // goes to default now
	default: // default only runs if no other case was true, (unless you specify fallthrough)
		fmt.Println("Great!")
	}

	scase := "hello"
	switch scase {
	case "nothello":
		fmt.Println("This won't be printed")
	default: // default can occur anywhere in the case, functions the same
		fmt.Println("Ahh! testing")
	case "hello", "there":
		fmt.Println("This is it!")
	}

	var stype interface{} = "hello"
	switch v := stype.(type) {
	case int:
		fmt.Println("int", v)
	case nil:
		fmt.Println("nil", v)
	case float64:
		fmt.Println("float64", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("You win!!", v)
	}

	// Takes input but treats space as seperator
	fmt.Println("Write just a word:")
	var userName string
	fmt.Scan(&userName)
	fmt.Println(userName)

	// Reads the entire line until there is a delimiter
	fmt.Println("Write an entire sentence:")
	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')
	fmt.Println("Here is the entire input", line, "NikhilJSK")

	// Read small and large number
	var (
		smallNumber int
		largeNumber int
	)
	fmt.Print("Enter two numbers:")
	fmt.Scan(&smallNumber, &largeNumber)
	if smallNumber > largeNumber {
		smallNumber, largeNumber = largeNumber, smallNumber
	}
	fmt.Printf("Remainder for %d / %d is %d \n", largeNumber, smallNumber, largeNumber%smallNumber)

	// Print all even :)
	fmt.Println("Even numbers b/w 0 & 10:")
	var n = 10
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}

	// Print fizz buzz
	fmt.Print("Fizz and Buzz")
	fizzBuzz(20)

	// Sum of multiples of 3,5
	totalSum := getSum(3, 5, 10)
	fmt.Print("Sum:", totalSum)

}
