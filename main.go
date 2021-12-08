package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var x = 42
var y = "James Bond"
var z = true

type phoneNumber int

var wg sync.WaitGroup
var wg5 sync.WaitGroup

type jsontemp struct {
	Doors int    `json:"number_of_doors"`       // How the name shoud appear in JSON
	Color string `json:"what_color, omitempty"` // Omit if the value is empty in the result JSON
}

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
func nationalSafetyCheck(a automobile) { // If the reciever here is a pointer, you only have to send a pointer.
	// Switch case with types, and assertion in the cases for strict variable access
	switch c := a.(type) {
	case sedan:
		fmt.Println("After rigorous tests, you have achieved 5 star in sedan category", a.(sedan).luxury, c)
	case truck:
		fmt.Println("After rigorous tests, you have achieved 5 star in truck category", a.(truck).fourWheel, c)
	case *truck:
		fmt.Println("Ahh! I see you sent a pointer. We don't do that here. Send me back the original :). Exempting for now!")
		a.start("*truck") // You can't access the variables of sedan/truck directly, you have to use the switch type statement
	case *sedan:
		fmt.Println("Ahh! I see you sent a pointer. We don't do that here. Send me back the original :) Exempting for now!", a)
		a.start("*sedan")
	default:
		fmt.Println("No automobie sent for testing. What do you want?")
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
	return f(oddNumbers...)
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

func concurrentFactorial(n int) int {
	if n == 1 {
		wg.Done() // Telling the compiler one goroutine is done!
		return 1
	}
	time.Sleep(time.Second * 1)
	// runtime.Gosched() // You could use this instead of time.sleep, which basically says go do something else, and come back to this thread!
	return n * concurrentFactorial(n-1)
}

func concurrentBigFactorial(n big.Int) *big.Int { // Here the return type is of pointer, cause z.Mul function return type is actually &z
	var z big.Int
	if cmp := n.Cmp(big.NewInt(1)); cmp == 0 {
		wg5.Done()
		return big.NewInt(1)
	}
	time.Sleep(100)
	z.Set(&n)                                   // n
	n.Sub(&n, big.NewInt(1))                    // n-1
	return z.Mul(&z, concurrentBigFactorial(n)) // n * fact(n-1)

	// Some explanation and steps on how on got to the solution!!
	// Big integers don't allow shallow copy, so you can just assign and modify, that will modify the newly created variable
	// Snippet from the doc
	// To "copy" an Int value, an existing (or newly allocated) Int must be set to a new value using the Int.Set method;
	// shallow copies of Ints are not supported and may lead to errors.
	// ------------------------------------------------------------
	// z.Set(&n)                // n
	// n.Sub(&n, big.NewInt(1)) // n-1
	// fmt.Printf("JSK - Z\t%v\t%v\t%v\n", z.String(), z, &z)
	// fmt.Printf("JSK - n\t%v\t%v\t%v\n", n.String(), n, &n)
	// y := concurrentBigFactorial(n) // call factorial(n-1)
	// fmt.Printf("JSK - Y\t%v\t%v\t%v\n", y.String(), y, &y)
	// return n.Mul(&n, y) // call n * factorial(n-1)
}

func changeNotConstant(number *int) int {
	*number = *number + 88
	return *number
}

func pointersExp(s *sedan) {
	fmt.Println("Hi Sedan! Your color now:", (*s).color)
	fmt.Println("Hi Sedan! Your color now:", s.color)
	(*s).color = "Base"
	fmt.Println("Hi Sedan! Your color later:", (*s).color)
	fmt.Println("Hi Sedan! Your color later:", s.color)

}

type groupedJsonTemp []jsontemp

func (jt groupedJsonTemp) Len() int {
	return len(jt)
}

func (jt groupedJsonTemp) Less(i int, j int) bool {
	return jt[i].Doors > jt[j].Doors
}

func (jt groupedJsonTemp) Swap(i int, j int) {
	jt[i], jt[j] = jt[j], jt[i]
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

func launchThread(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case value := <-e:
			fmt.Println("Even!", value)
		case value := <-o:
			fmt.Println("Odd!", value)
		case value := <-q:
			fmt.Println("Quit!", value)
			return
		}
	}
}

func send2(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive2(e, o <-chan int, fanIn chan<- int) {
	var wgr sync.WaitGroup
	wgr.Add(2)
	go func() {
		for v := range e {
			fanIn <- v
		}
		wgr.Done()
	}()
	go func() {
		for v := range o {
			fanIn <- v
		}
		wgr.Done()
	}()
	wgr.Wait()
	close(fanIn)
}

func getRequests(c chan<- int) {
	for i := 0; i < 15; i++ {
		c <- i
	}
	close(c)
}

func executeRequestsThrottle(requestQueue, outputQueue chan int, throttleLimit int) {
	var wgr1 sync.WaitGroup
	wgr1.Add(throttleLimit)
	for i := 0; i < throttleLimit; i++ {
		go func() {
			for v := range requestQueue {
				outputQueue <- burstTimeRequired(v)
			}
			wgr1.Done()
		}()
	}
	fmt.Println("Number of Go Routines FanOutIn. After:", runtime.NumGoroutine())
	wgr1.Wait()
	close(outputQueue)
}

func burstTimeRequired(v int) int {
	time.Sleep(time.Second * 2)
	return v + mathrand.Intn(100)
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

	// Pointer
	var value int = 5
	var ptr *int = &value
	fmt.Println("Pointer", value, *&value, *ptr, ptr, &value)
	*ptr = 53
	fmt.Println("Pointer", value, *&value, *ptr, ptr, &value)

	// Method Sets
	fmt.Println("Method Sets")
	nationalSafetyCheck(t1) // Since the receiver is of value type, you can either send a value or pointer to that value like below
	nationalSafetyCheck(&t1)

	notConstant := 5
	fmt.Println("Playing with Pointers!! Before:", notConstant)
	changeNotConstant(&notConstant)
	fmt.Println("Playing with Pointers!! After:", notConstant)

	fmt.Println("Sedan color before:", s1.color)
	pointersExp(&s1)
	fmt.Println("Sedan color after:", s1.color)

	// JSON - Marshal & Unmarshal
	v1 := jsontemp{
		Doors: 3,
		Color: "pink",
	}
	v2 := jsontemp{
		Doors: 4,
		Color: "hot rod",
	}
	bytes, err := json.Marshal([]jsontemp{v1, v2})
	if err != nil {
		fmt.Println("Damn! Something bad happened!", err)
	}
	fmt.Println("Marshalled!", string(bytes), "Bytes!", bytes)

	var data *[]jsontemp
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Damn! Something bad happened!", err)
	}
	fmt.Println("UnMarshalled!", data, "Deferenced!", (*data)[0].Color)
	fmt.Println("UnMarshalled!", data, "Deferenced!", *data)

	// Sort
	fmt.Println("Sorting")
	sortNums := []int{4, 23, 5, 1, 1}
	sort.Ints(sortNums)
	fmt.Println("Sorted!", sortNums)
	sort.Sort(sort.Reverse(sort.IntSlice(sortNums))) // sort.Reverse needs the interface Interface to be implemented by three funcs (Len, Less, Swap)
	// Here we have IntSlice which implements all three funcs and so we can use that to pass to Reverse, we can't pass sort.Ints
	// Since Reverse returns an Interface, we need to send it to sort.Sort that accepts Interface
	fmt.Println("Sorted Reversed!", sortNums)

	sortStrings := []string{"hello", "011", "_", "Hello"}
	fmt.Println("Before:", sortStrings)
	sort.Strings(sortStrings)
	fmt.Println("After:", sortStrings)
	sort.Sort(sort.Reverse(sort.StringSlice(sortStrings))) // sort.StringSlice(sortStrings) -- Means conversion of []string to Interface!!
	fmt.Println("Way After:", sortStrings)

	// Sort on structs
	fmt.Println("Sort on structs")
	// Sort vehicles based on number of doors
	// allVehicles := []jsontemp{v1, v2} // If you implement this way, you won't be able to do `func (a []jsontemp)` so use a type for this
	// Now, since we'll be using sort.Sort, we have to impement Interface with three funcs (Len, Swap, Less)
	dataJson := groupedJsonTemp{v1, v2}
	fmt.Println("Before struct:", dataJson)
	sort.Sort(dataJson)
	fmt.Println("After struct:", dataJson)

	// Concurrency
	fmt.Println("Concurrency")
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("Number of Go Routines", runtime.NumGoroutine())
	fmt.Println("Number of C Go Calls?", runtime.NumCgoCall())

	// WaitGroup
	wg.Add(3) // Asking the compiler to wait for 3 goroutines to exit
	go func() {
		fmt.Println("Factorial for 5", concurrentFactorial(5)) // Since you are using println you have to declare it inside a anony. func, otherwise goroutine is not properly spawned
	}()
	go func() {
		fmt.Println("Factorial for 6", concurrentFactorial(6))
	}()
	go func() {
		fmt.Println("Factorial for 10", concurrentFactorial(10))
	}()

	fmt.Println("Number of CPUs. After:", runtime.NumCPU())
	fmt.Println("Number of Go Routines. After:", runtime.NumGoroutine())
	fmt.Println("Number of C Go Calls?. After:", runtime.NumCgoCall())

	// Big Integers - Unfinished!
	fmt.Println("Number of Go Routines - Big Integer", runtime.NumGoroutine())

	// WaitGroup
	wg5.Add(3) // Asking the compiler to wait for 3 goroutines to exit
	go func() {
		fmt.Println("Factorial for 50", concurrentBigFactorial(*big.NewInt(50)))
	}()
	go func() {
		fmt.Println("Factorial for 100", concurrentBigFactorial(*big.NewInt(100)))
	}()
	go func() {
		fmt.Println("Factorial for 101", concurrentBigFactorial(*big.NewInt(101)))
	}()
	fmt.Println("Number of Go Routines. After Big Integer:", runtime.NumGoroutine())

	// Race Condition
	var wg2 sync.WaitGroup
	fmt.Println("Race Condition!")
	wg2.Add(20)
	var counter int
	for i := 0; i < 20; i++ { // Ideally the value here must be 20
		go func() {
			temp := counter
			temp++
			runtime.Gosched()
			counter = temp
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println("Counter with Race Condition:", counter)

	// Handling Race condition with Mutex
	var wg3 sync.WaitGroup
	var mu sync.Mutex
	fmt.Println("Handling Race condition with Mutex")
	wg3.Add(20)
	var counter2 int
	for i := 0; i < 20; i++ { // the value here must be 20
		go func() {
			mu.Lock()
			temp := counter2
			temp++
			runtime.Gosched()
			counter2 = temp
			mu.Unlock()
			wg3.Done()
		}()
	}
	wg3.Wait()
	fmt.Println("Counter with Mutex:", counter2)

	// Atomic
	var wg4 sync.WaitGroup
	fmt.Println("Atomic")
	wg4.Add(20)
	var counter3 int32
	for i := 0; i < 20; i++ { // Ideally the value here must be 20
		go func() {
			atomic.AddInt32(&counter3, 1)
			runtime.Gosched()
			wg4.Done()
		}()
	}
	wg4.Wait()
	fmt.Println("Counter with Atomic:", atomic.LoadInt32(&counter3))

	// Channels
	fmt.Println("Channels!")
	ch := make(chan int) // unbuffered channel
	go func() {          // Without this go routine, the ch <- <value> fails, as both sending and receiving should occur at the same time, and without a go routine that won't be possible
		ch <- 44
	}()
	fmt.Println("Channel value received:", <-ch)

	ch2 := make(chan int, 1) // Create a buffered channel of size 1, so now it allows upto 1 vaule to sit in there, if more than one is sent, it fails with deadlock
	ch2 <- 44
	fmt.Println("Channel value received:", <-ch2)

	// Directional channels
	ch3 := make(chan<- int) // send-only type channel
	ch4 := make(<-chan int) // receive-only type channel // This means channels can start out bidirectional, but magically become directional simply by assigning a regular channel to a variable of a constrained type
	fmt.Printf("%T %T", ch3, ch4)

	ch5 := make(chan int)

	go launchThread(ch5)
	for i := range ch5 {
		fmt.Println(i)
	}

	// Select statement
	e := make(chan int)
	o := make(chan int)
	q := make(chan int)
	go send(e, o, q)
	receive(e, o, q)

	// Fan-in
	e1 := make(chan int)
	o1 := make(chan int)
	fanIn := make(chan int)
	go send2(e1, o1)
	go receive2(e1, o1, fanIn)
	for v := range fanIn {
		fmt.Println("FanIn:", v)
	}
	fmt.Println("Exit fanIn!!")

	// time.Sleep(time.Second * 15)
	fmt.Println("Number of Go Routines FanOutIn. Before:", runtime.NumGoroutine())
	// FanOutIn with throttling
	CPURequestQueue := make(chan int)
	OutputQueue := make(chan int)
	go getRequests(CPURequestQueue)
	go executeRequestsThrottle(CPURequestQueue, OutputQueue, 10)
	for r := range OutputQueue {
		fmt.Println("FanOutIn throttle", r)
	}

	// Context
	ctx := context.Background()
	fmt.Println("Number of Go Routines Context. Before:", runtime.NumGoroutine(), "Err:", ctx.Err())
	go func() {
		n := 0
		for {
			select {
			case v1 := <-ctx.Done():
				fmt.Println("CTX is done for!!", v1)
				return
			default:
				fmt.Println("CTX", n)
				n++
			}
		}
	}()
	fmt.Println("Number of Go Routines FanOutIn. Middle:", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("Cancelling context now", ctx.Err())
	cancel()
	fmt.Println("Error now", ctx.Err())

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

	wg.Wait() // We can wait here to save time!
	wg5.Wait()
}
