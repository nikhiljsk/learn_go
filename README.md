[1](https://www.notion.so/Golang-491f172b8d6f4aa7961293626e1536ed) [2](https://ibm-learning.udemy.com/course/learn-how-to-code/learn/lecture/11921686#overview)

# Table of Contents

# Go: The Complete Developer's Guide

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

Go is a Static Typed language unlike Python which is a dynamic typed language

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

Array (Fixed Size) vs Slice (Re-sizeable arrays) - Slices should be homogenous (data types)

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
```

- GOROOT contains the binary exec file of go
- GOPATH contains the folder path for your workspace
- your workspace should have three folders, bin - contains binaries, src - contains all the code in a folder structured format according to the package, then we have package which maintains archives to cache modules
- Go Mod file is a package manager that tracks all the direct packages that are being used in your package and keeps them organized version wise. You can upgrade or downgrade the use of direct packages that will be tracked in go.mod file
- `...` means unlimited parameters, variadic parameters in a function

Need to explore more on big integers and signed integers

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
-