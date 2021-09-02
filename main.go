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
	_   = iota
	ikb = 1 << (iota * 10) // now 1
	imb = 1 << (iota * 10) // now 2
	igb = 1 << (iota * 10) // now 3
)

type person struct {
	firstName string
	lastName  string
	age       int
	icecreams []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) start(greeting string) {
	fmt.Println("Inside start function! Hi! ", greeting)
	fmt.Println("Trying to get the color right!", t.color)
}

func (s sedan) start(greeting string) {
	fmt.Println("Inside start function! Hi! ", greeting)
	fmt.Println("Trying to get the color right!", s.color)
}

// Since both Sedan and Truck have implemented start function, both can now be a type of interface
// Interfaces are really helpful to direct to corresponding function type. Example, an interface SHAPE can
// direct to appropriate area functions of both circle and square, which have the area method implemented differently
type automobile interface {
	start(a string)
}

// Since this function takes in multiple types - Polymorphism
func nationalSafetyCheck(a automobile) {
	// Switch case with types, and assertion in the cases for strict variable access
	switch c := a.(type) {
	case sedan:
		fmt.Println("After rigorous tests, you have achieved 5 star in sedan category", a.(sedan).luxury, c)
	case truck:
		fmt.Println("After rigorous tests, you have achieved 5 star in truck category", a.(truck).fourWheel, c)
	}

}

func something(name string, unlimited ...string) bool {
	defer fmt.Println("Hey there, I'll only be executed at the end of this function!") // executes just before the return call is executed!
	fmt.Println("Name:", name)
	fmt.Printf("Type\t%T\n:", unlimited)
	for i, v := range unlimited {
		fmt.Printf("\ti:%v\tv:%v\n", i, v)
	}
	return true
}

func returnAnotherFunc() func() int {
	x := 5
	return func() int {
		fmt.Println("Executed once - ReturnAFunction", x)
		return x
	}
}

func sumAll(numbers ...int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return int(total)
}

func callbackFunc(f func(num ...int) int, allNumbers ...int) int {
	// Find the sum of only odd numbers
	var oddNumbers []int
	for _, v := range allNumbers {
		if v%2 != 0 {
			oddNumbers = append(oddNumbers, v)
		}
	}
	return sumAll(oddNumbers...)
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func fakeIncrementor() int {
	x := 0
	x++
	return x
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

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
		fmt.Printf("%v\t%v\t%#U\t%v\t%v\n", i, string(i), i, strconv.Itoa(i), string(123858))
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
	case "hello", "there":
		fmt.Println("This is it!")
	default: // default can occur anywhere in the case, functions the same
		fmt.Println("Ahh! testing")
	}

	// Define a variable that can take in multiple value types cause any type has interface{} (i.e no functions implemented)
	var stype interface{} = "hello"
	stype = 45
	stype = true
	stype = map[string]int{
		"Great!": 1,
	}

	// The following three lines don't work cause golang is static type and you're trying to manipulate the variable type
	// ntype := 4
	// ntype = "String there!"
	// fmt.Println(ntype)

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

	// Arrays
	var arr = [5]int{1, 2, 3, 4}

	// Slices
	var sarr = []int{1, 2, 4, 5, 6, 7}
	fmt.Println("Arrays/Slices", arr, sarr, len(sarr), cap(sarr))

	for i, v := range sarr {
		fmt.Println("index & Value", i, v)
	}
	fmt.Println(sarr[:4], sarr[:2]) // Negative values don't work

	temp := []int{7, 8, 0}
	newSarr := append(append(sarr, temp...), 7, 8, 9) // ... is use to unfurl all the values, similar to * in python
	fmt.Println("Appended now!", newSarr)

	newSarr = append(newSarr[:4], newSarr[6:]...)
	fmt.Println("Deletion of slice elements", newSarr)

	fmt.Println("Make in Array") // Slices are just a datastructure with 3 fields, (ptr to array, len, capacity)
	makeArr := make([]int, 2, 3)
	makeArr = append(makeArr, 1, 2)
	fmt.Printf("Array %v\tLength %v\tCapacity %v\n", makeArr, len(makeArr), cap(makeArr))
	makeArr = append(makeArr, 3)
	fmt.Printf("Array %v\tLength %v\tCapacity %v\n", makeArr, len(makeArr), cap(makeArr))
	makeArr = append(makeArr, 4, 5)
	fmt.Printf("Array %v\tLength %v\tCapacity %v\n", makeArr, len(makeArr), cap(makeArr)) // Capacity is doubled if array overflowed

	doubleSlice := [][]int{newSarr, makeArr}
	fmt.Println("2D Slice:", doubleSlice)

	// Maps
	phonebook := map[string]int{
		"Nikhil": 1234,
		"JSK":    783,
	}
	// Access a value, by default if int, so value is 0. So use if exists bool value
	if v, ok := phonebook["JSK"]; ok { // Also called as the 'comma ok' idiom
		fmt.Println("Map - Requested entry exists", v)
	} else {
		fmt.Println("Map - Get the hell out of here!")
	}
	fmt.Println("Hihi. I don't actually exist. But I have zero as value", phonebook["pspk"])
	if v, ok := phonebook["JSK"]; ok {
		fmt.Println("Value exists. Go ahead and delete!", v)
		delete(phonebook, "JSK")
	}

	// Creating a map of slices of strings
	mapStrings := map[string][]string{
		"agrdne": {"garden", "danger"},
		"apn":    {"nap", "pan"},
	}
	fmt.Println("Nice job!", mapStrings)

	// Structs
	fmt.Println("Structs")
	p1 := person{
		firstName: "Nikhil",
		lastName:  "JSK",
		age:       22,
		icecreams: []string{"chocolate", "butterscotch"},
	}
	fmt.Println("Here's one struct", p1)
	p2 := person{
		firstName: "Akhil",
		lastName:  "JSK2",
		age:       24,
		icecreams: []string{"chocolate", "strawberry"},
	}
	fmt.Println("Here's another", p2)

	people := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println("Map:", people)

	for k, v := range people {
		fmt.Println("FirstName", v.firstName)
		fmt.Println("LastName:", k)
		fmt.Println("My Fav Flavors:")
		for _, flavor := range v.icecreams {
			fmt.Println("\t", flavor)
		}
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: false,
	}
	fmt.Println("Embed struct:", s1, "\n", t1)

	report := struct {
		errors int
		logs   string
	}{
		errors: 2,
		logs:   "Here's one error\n Here's the other",
	}

	fmt.Println("Anonymous Struct!!", report)

	// Functions!
	fmt.Println("Functions!")
	randomWords := []string{"is", "the best", "you", "can", "find outthere!"}
	fmt.Println(something("Nikhil", randomWords...))

	// Interfaces!
	fmt.Println("Starting Interfaces ÍÎ˝Í€€")
	s1.start("Skoda Rapid")
	t1.start("Volvo Truck")

	nationalSafetyCheck(s1)
	nationalSafetyCheck(t1)

	// Anonymous Function
	fmt.Println("Anonymous. Sarr before:", sarr)
	func(test []int) {
		test = append(test, 19903)
		fmt.Println("Inside anonymous function!", test)
	}(sarr)
	fmt.Println("Anonymous. Sarr After:", sarr)

	funcVariable := fizzBuzz // Since everything in golang is just a type!
	funcVariable(10)

	// Return a func
	fmt.Println("Return a func")
	fmt.Printf("Type of a function that returns another function %T\n", returnAnotherFunc()())
	returnAnotherFunc()()

	// Callback function - Passing a function as a paramter to another function!
	fmt.Println("Callback function - Passing a function as a paramter to another function!")
	fmt.Println("Sum of all numbers:", sumAll(18, 1)) // Modulo operator is not defined for float
	fmt.Println("Sum of all odd numbers:", callbackFunc(sumAll, 18, 1))

	// Closure
	fmt.Println("Closure")
	firstIncrementor := incrementor()
	secondIncrementor := incrementor()
	fmt.Println("First:", firstIncrementor())
	fmt.Println("First:", firstIncrementor()) // Since this is a closure, 1 is incremented to 2
	fmt.Println("First:", firstIncrementor())
	fmt.Println("Second:", secondIncrementor())
	fmt.Println("Second:", secondIncrementor())

	fmt.Println("Difference b/w closure and normal function")
	fakeInc := fakeIncrementor()
	fmt.Println("fakeInc:", fakeInc)
	fmt.Println("fakeInc:", fakeInc)

	// Recursion
	fmt.Println("Recursion")
	fmt.Println("Factorial", factorial(5))

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
