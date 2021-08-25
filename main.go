package main

import (
	"bufio"
	"fmt"
	"os"
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
