Source: [1](https://ibm-learning.udemy.com/course/learn-how-to-code/learn/lecture/11921686#overview) [2](https://ibm-learning.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7817382?start=0#overview) [3](https://www.tutorialspoint.com/go/index.htm) [4](https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa)

# Table of Contents

# Learn How To Code: Google's Go (golang) Programming Language

- Created by Google and published 1st Opensource version on 2012 to handle concurrency and multiple cores efficiently for webservices. YT is entirely now in GoLang
- Creator of Node JS has abandoned in favor of Golang 🍭
- You can check the hash version to see if you downloaded the original file. Checksums

```bash
openssl sha -sha256 <folder>

# get env variables
go env

# formats the code in current directory
go fmt

# creates an executable in current directory
go build main.go # Creates ./main

# Creates an executable in bin directory which you can then use directly to run your program
go install main.go

go mod init <package_name> # inits your module, package name can be anything not just domain
go build/ go test # Runs and gets and tracks the dependencies
go list -m all # Lists all the direct and indirect packages
go get # get
go mod tidy # cleans the packages
go get <package_name>@<version> # that can get a specific version
go list -m versions <package_name> # list all versions
go list -deps <p_name> # lists all dependenciesch

go get -u <package_name> # get or update a package

# ENVS related to private
go env GOPRIVATE
go env -w GOPRIVATE="github.ibm.com"

# To generate documentation from a package
go doc <packge>.<symb>.<method>
(or)
# Localhost server
godoc -http=:6060
godoc <package> <method> # Only the docs
godoc -src <pcakge> <method> # Shows the source declaration

go ./...
# You can also just paste the url of your go source code in godoc.org
# and then it is cached, you can refetch it.

go fmt # Formats code
go vet # Reports suspicious constructs
golint # Suggests style mistakes

go test -v -bench .
go tool cover -html=cover.out -o=cover.html

go list -m -versions github.com/gorilla/mux
go list -m all
go mod tidy
go mod verfiy
go mod vendor # Used for getting all the packages and store them locally
go run -mod=vendor main.go # Using this will fetch the packages directly from the vendor
go mod why github.com/gorilla/mux # tells why you are dependent on that module
go mod graph # Lists all for above command instead of one
```

- GOROOT contains the binary exec file of go
- GOPATH contains the folder path for your workspace
- your workspace should have three folders, bin - contains binaries, src - contains all the code in a folder structured format according to the package, then we have package which maintains archives to cache modules
- Go Mod file is a package manager that tracks all the direct packages that are being used in your package and keeps them organized version wise. You can upgrade or downgrade the use of direct packages that will be tracked in go.mod file
- `...` means unlimited parameters, variadic parameters in a function

<aside>
💡 Need to explore more on big integers and signed integers

</aside>

- Identifiers - Just named entites, can be anything. Some predeclared identifiers include bool, true, false, nil, append close, byte, int etc.
- Keywords - That can't be used as an identifiers ex - break, default, defer, go
- Statement is nothing but an instruction to computer
- `var` is used when you're declaring global, and short declaration operator is only within the function scope
- You can use `%T` in printf to see the type of the variable
- Go is a static prog. lang. so a variable can only hold value of a certain type
- To declare a raw string which also includes the escape characters use ticks (`)
- It is conversion in golang, not casting
- Every value in go is an empty interface and is expressed as interface{}
- Everything is pass by value in golang
- UTF - 8 is just another encodinng scheme like ascii, but is far better and most widely used as it supports multiple languges, also created by golang inventors
- Four gens of computers so far, vaccum tube, transistors, integrated chips, microprocessors all to store 0/1's
- You can use int8 to store signed integers (-128 to 127) using only 8 bits, but the same uint8 can store unsigned (0 to 255). You can just use int that will use int64 or int32 underlying based on the OS instruction set
- Rune is 4 bytes (32 bits) (int32) in utf-8, byte is for uint8
- No `while` statement in go
- Arrays are the building blocks of slices, it is clear not to use arrays, use slices instead.
- For embedded structs the inner type gets promoted to the outer type. (More like the child class can access the elements of the parent class - You can access those variables using child.Age or child.Parent.Age, both give the same value unless you have the variable age defined in both the structs, in case of collision, it gives preference to local level first)
- The embedded struct can be anonymousField which doens't have a varialbe associated, for example you can just inherit person and use it instead of mentioning `p1 person`
- If you are really concerned about performance declare variables from largest to smallest
- Variadic parameters can has 0 to unlimited parameters. When has to be a final parameter
- Defer is a keyword when used executes the statement it is associated with at the end of function exit
- Functions are first class citizens, cause they can be passed into another function, returned from another function and can be assigned to another variable
- Callback - Passing a function as an argument
- Closures - Limiting the scope of the variables. Useful for things like incrementor functions where the value of the variable is stored in-memory and increments each time you call the function, kind of similar to iota
- This website converts you JSON to struct ([mholt.github.io/json-to-go](https://mholt.github.io/json-to-go/))

![Screenshot 2021-09-04 at 12.26.50 AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/25c5061d-d211-45c6-906b-9880cd2fdc0a/Screenshot_2021-09-04_at_12.26.50_AM.png)

- Golang was the first language which leveraged multi-core cpu's, as it was built in 2007 and in 2006 Intel's dual core cpu was widely popular!
- Concurrency is a design pattern that allows the new code to run on multiple CPUs but that does not guarantee parallelism the only thing that guarantees parallelism is the number of CPU cores you have
- You can use `func init(){}` that runs before the main function
- Method Sets - Since the receiver is of value type, you can either send a value or pointer to that value like below (Needs more digging on this). Sometimes, even if the accepted value is a pointer, a value can be sent. Incase of waitGroup that happens, and my basic understanding is that it has something to do with the method types and not just at the interface level
- Race Condition - Different routines trying to access the same shared variables and resulting in bugs
- You could find the race condition by using `go build -race main.go` which tries to see if any data race is found and the number. Gomaxproc is the number of cores you want it to run on. Default is max
- There are read and write locks as well in Mutex
- Send-only and receive-only type channel, directional channels. These channels can start out bidirectional, but some can become directional simply by assigning a regular channel to a variable of a constrained type.
- With Panic deferred functions run, but with Fatal the program exists abruptly.
- We don't use Public and Private in golang. We only use Visible/Not-visible or exported or not-exported. For functions that are visible outside the package we start with Capital letter as naming convention and functions that are used only for calculations inside the package we use small letters
- In documentation the starting word should be the function name and then go on to define the function. For package, you'll have to start with `Package <name>` If the documentation is huge, just use doc.go file
- Test files doesn't have to have the same name as the source file, and the function names can be different as well, and they can be different packages as well. But it's best practice to do all that! 🙂
- But for Examples, you'll have to give the same name (case-sensitive) in order for that to work. The function name has to start with capitals
- Semaphore - In computer science, a semaphore is a variable or abstract data type used to control access to a common resource by multiple threads and avoid critical section problems in a concurrent system such as a multitasking operating system.
- An unbuffered channel cannot hold values, so it has to send and recieve at the same time. So we launch a go routine for that
- Using range to iterate over channel output? don't forget to close that.

# Go: The Complete Developer's Guide

<aside>
💡 Might have to revisit the course 🙂

</aside>

GO CLI

1. go run - Compile and execute (creates filename from filename.go)
2. go build - Just compile
3. go fmt - Format all go code
4. go install - Compiles and intalls a package
5. go get - Downloads the raw source code of someone else's package
6. go test - run test files

Types of Packages

More like a project/library

1. Executable - To write custom logic (Only `package main` is the executable, so go build creates a executable file, for reusable it won't create) - Should definetly have main fucntion.
2. Reusable - To reuse code from different package, dependency

Hello World Program

```go
package main //Creating my own project main

import "fmt" // Give access to fmt package inside my main package
// standard lib

func main() { // Main function which is called by OS?
    fmt.Println("hello world")
}
```

<aside>
💡 Go is a Static Typed language unlike Python which is a dynamic typed language

</aside>

```go
// Two ways of initializing a variable
var card string = "Hello"
card := "Hello" //Initialization NOT Replace

// For global initialization former
// meathod works, later doesn't

// If you change the value of global
// variable in any function, it changes the
// global value instead of local
```

<aside>
💡 Array (Fixed Size) vs Slice (Re-sizeable arrays) - Slices should be homogenous (data types)

</aside>

- Since we don't have classes in GO, we can create custom data types and then associate functions to it.

```go
Things to syntax-work-on
1. Types - Int, string, array, slice
2. Functions - return types, pass arguments with types (custom aswell)
3. For loop - iterate over arrays/slices
4. Custom types
5. Receiver functions/Method for custom types

```

Type Conversion

```go
[]byte("Hello")
```

Pointers in Golang

```jsx
& - Says that give access to the memory address of the variable
* - Says that give access to the value in the memory address
But!!
* infront of a type is totally different. It just indicates
that we are goin to deal with pointers that's it.

----
In golang you can call receiver function with just normal variable without using pointer.
It automatically type converts
```

- Explained here

    ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/232580da-d3bf-472a-9801-86cfdd8210d4/Screenshot_2021-04-13_at_2.36.04_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/232580da-d3bf-472a-9801-86cfdd8210d4/Screenshot_2021-04-13_at_2.36.04_AM.png)

    Here automatic type conversion

    ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c602aa06-5443-4a02-bca8-4fe5e25f446f/Screenshot_2021-04-13_at_2.38.54_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c602aa06-5443-4a02-bca8-4fe5e25f446f/Screenshot_2021-04-13_at_2.38.54_AM.png)

- For slices, there is no need for pointers, these are automatically call by reference NOT call by value

    ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e4fcca22-4dc4-4c64-af4b-f205995de890/Screenshot_2021-04-13_at_2.41.04_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e4fcca22-4dc4-4c64-af4b-f205995de890/Screenshot_2021-04-13_at_2.41.04_AM.png)


Slices expalined!!

Slice is a datastructure with three different fields. So whenever you try to make a copy of the slice, you are not duplicating the array, you are just duplicating the data structure fields, but the array it points to remains the same 🙂

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7b47efc2-7598-4c46-8c69-1da0850f4523/Screenshot_2021-04-13_at_2.48.04_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7b47efc2-7598-4c46-8c69-1da0850f4523/Screenshot_2021-04-13_at_2.48.04_AM.png)

Similarly for other data types as mentioned here.

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/adf5d673-7944-4b2e-9ba0-c17006d2cc05/Screenshot_2021-04-13_at_2.49.43_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/adf5d673-7944-4b2e-9ba0-c17006d2cc05/Screenshot_2021-04-13_at_2.49.43_AM.png)

Everything in GO is pass by value. So the pointers being printed here would give two different addresses

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/765db64b-8aff-435d-8de8-9e4f0af9cd6b/Screenshot_2021-04-13_at_2.54.15_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/765db64b-8aff-435d-8de8-9e4f0af9cd6b/Screenshot_2021-04-13_at_2.54.15_AM.png)

## Maps

Key-value pairs

- Keys should be of same type, values should be of same type
- You can iterate over keys and values unlike structs

### Interfaces

- Since in golang, you have to specify the type of the arguments that are received in function, assume you want to change the type of the argument from int to float, that would demand you to rewrite the entire logic. That is why we have interfaces
- All you have to do, if a struct implement all the functions that are defined in the interface, the struct is automatically given access to all the interface functions. So interfaces are implicit
- You can define a set of interfaces inside a interface that should be satisfied to grant access

    ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/debbd6b8-a6d6-4889-ae4f-68d7377f8357/Screenshot_2021-05-18_at_1.34.22_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/debbd6b8-a6d6-4889-ae4f-68d7377f8357/Screenshot_2021-05-18_at_1.34.22_AM.png)

- Why is the read function not returning any data?
    - Simple, let's say I call function with a byte slice, since it is call by reference, all we are doing is passing in byte data and adding value internally which will automatically reflect in the object that called read function

        ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7f7f99bb-e743-4a49-b29a-0108a473f192/Screenshot_2021-05-18_at_1.42.32_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7f7f99bb-e743-4a49-b29a-0108a473f192/Screenshot_2021-05-18_at_1.42.32_AM.png)

    - Understanding the above scenario

        ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d9090425-edc3-419b-9b58-afebfd58dad4/Screenshot_2021-05-18_at_1.45.54_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d9090425-edc3-419b-9b58-afebfd58dad4/Screenshot_2021-05-18_at_1.45.54_AM.png)

    - Similar to what has been done in before image but better 🙂

        ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b06edf84-0f0b-476c-9fd6-b4aa04337d74/Screenshot_2021-05-18_at_1.51.08_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b06edf84-0f0b-476c-9fd6-b4aa04337d74/Screenshot_2021-05-18_at_1.51.08_AM.png)


### Go Routines

- Using simple a `go` keyword spawn another process
- But golang by default use only one cpu core, so even if there are multiple go routines that you've created, the go-scheduler tries to context switch b/w those and execute them, but in reality it's not truly parallel. You can enable golang to use multiple cores, then it is truly parallel
- Concurrency vs parallelism
    - Similar to multi-threading (if one thread is blocked then other is executed) vs multi-parallelism (multiple processes on multiple cores)
- For waiting for the child process to quit, we introduce channels, which are also typed based on what you are trying to pass b/w them.

    ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/962e3983-a72f-4ae6-9d18-8a22df4e40c1/Screenshot_2021-05-18_at_2.34.26_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/962e3983-a72f-4ae6-9d18-8a22df4e40c1/Screenshot_2021-05-18_at_2.34.26_AM.png)

- Anonymous functions with appropriate sleep time - similar to lambda, called as Function Literals in golang
- You can loop over the channels and assign the return value to a variable that is why you don't see any <-
- For anonymous function we pass link as value, cause if we don't, it'll be pass by reference, then if the link value changes
- In main function, even for the child go routine it changes. We don't want that.
- Channels - You can block the execution of the a go routine like (main.go) and wait until there is a pipeline push from somewhere
- If you pass a value into a channel, then you have to receive it somewhere.

    ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6422f0cf-6849-4c63-9ecd-04a6bc3f81fd/Screenshot_2021-05-18_at_12.44.52_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6422f0cf-6849-4c63-9ecd-04a6bc3f81fd/Screenshot_2021-05-18_at_12.44.52_PM.png)


# Golang - TutorialsPoint

- Static typed language, to keep it simple some operations like Pointer arithmetic, typed inheritance etc are intentionally avoided
- Types - Numberic, String, Boolean and Derived (all other)
- Fixed values in constants are also called literals
- For constant variables use Uppercase
- Bitwise operators (+, | ^)(and, or, XOR) also << & >>
- goto statement exists in golang

```go
LOOP: for a < 20 {
      if a == 15 {
         /* skip the iteration */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("value of a: %d\n", a)
      a++
   }
```

- By default, Go uses call by value to pass arguments
- Local - Both in main and function-specific, Global - Scope is entire program, Formal Parameters - Function specific and precedence over global
- Copy(target, source) For slices
- For interface, both structs must implement area function, and then you need to create a new function with getArea(s shape) that you can pass any of the two structs and call the respective area function
- lvalues - Appear on both sides of operator (variables), rvalues only on the right side of assignment operator (constants)
- Go does not support Method and operator overloading
- Method sets are applicable whenever you are calling ().function.
    - So (e employee) getData (){} can’t accept `var e *employee` e.getData.
    - But (e *employee) getData(){} can accept `var e *employee` and `var e employee` e.getData!
- `const` values must be declared and initialized in the same line.



# Golang - Hitesh Choudhary

- Lexer is something that validates the grammar of the language, and in golang it is built-in to add semi-colons to the code. So they are not mandatory
- There is a standard date and time that you have to use to format any date
- Memory allocation
    - new() → Memory is allocated but no init and is zeroed storage (No data can be entered)
    - mase() → Memory is alloc and init, non-zeroed
- There is a threshold after which Garbage collection starts, and this parameter is configurable in Runtime package
- Whenever a request is made using the http package, it has to be closed using `Close` explicitly
- Modules in go are introduced in 2019. The downloaded modules are stored under `~/go/pkg/mod/cache/download` and in go.mod file, indirect means the library is not yet used in any of the go files. In go.sum, all the checksums are stored. `go mod tidy` is used to refresh the go.mod file.
- `go mod verify` will verify the packages with the checksum, `go list -m all` shows all the dependencies of current package