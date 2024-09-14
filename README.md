# Learning Golang

### Resources:
- [https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)

### Chapter 1: Go runtime
Go runtime consists of several key components, including the scheduler, garbage collector, memory allocator, and stack management.
It is attached with every executable binary.

Unlike Java, go does not uses any type of virtual machine so go runtime handles memory management.

**Run command**
```bash
go run <file-name or package-name> 
```
Preferred when want to run a single file

**Build executable**
```go
go build
```

```go
go install
```
This set the path of the binary in the path. So we can run the executable globally.

**Go file syntax**
- Main file which contains main function must include `package main` at the top.
- This file should contain a main function like `func main() {}`.

**Initialization command**
```bash
go mod init <REMOTE>/<USERNAME>/<PROJECT-NAME>
```

**What is the use of `go.mod` file**
It is similar to `package.json` file, it contains dependencies, package name and go version.

> Go mono-repo can contain multiple go modules, and root of the project generally contains `go.mod` file.

**What is the use of `go.sum` file**
It contains transitive dependencies which are being used by third-party modules.

**Import local module in other module**
Not recommended in production `go.mod` files.
`go.mod`
```go
replace <module-path> v<version> => <local-module-path>
```

**Download and install third party dependencies**
```go
go get <remote-url>
```

### Chapter 2: Basic

#### Packages
In golang, packages are used to categorize our program into logical units so that it is easy to maintain.
Every go-file belongs to some package i.e., why we write `package <package-name>` at the top. Each go application
must have `main` package so that it can compile.

> Note: If a package is changed and recompiled, all the client programs that use this package must be recompiled too!

#### Import
A Go program is linked to different packages through the import keyword.

**Syntax:**
```go
import "<package-name>"
```

**Example:**
```go
import "fmt"
import "os"

// or
import "fmt"; import "os"

// or
import (
	"fmt"
	"os"
)

// or
import ("fmt"; "os")
```

#### Visibility
- An identifier can be variable, constant, function, type or struct field. We can declare identifier in lowercase or uppercase letters. 
- If we declare identifier in lowercase letter, it will be visible within the package only. But if we declare package in uppercase letter, it will be visible within and outside the package which is also known as exported. 
- The dot `.` Operator is used to access the identifier e.g. pack.Age where pack is the package name and Age is the identifier.

#### Data types

| Type                                               | Default value |
|----------------------------------------------------|---------------|
| bool (8-bit)                                       | false         |
| string                                             | ""            |
| int (32 or 64 bit), int8, int16, int32, int64      | 0             |
| uint (32 or 64 bit), uint8, uint16, unit32, uint64 | 0             |
| byte (alias for uint8)                             | 0             |
| rune (alais for int32)                             | 0             |
| float32, float64                                   | 0             |
| complex64, complex128                              | (0 + 0i)      |

**Note:** `rune` is equivalent to char

#### Declaring a variable

```go
var number int
```

To declare a variable called `pi` to be of type `float64` with a value of `3.14159`:

```go
var pi float64 = 3.14159
```

**Short variable declaration**
This means declaring a variable and initializing it at the same time.

```go
number := 10 // var number int = 10
```

**Type conversion**
Functions for primitive data types like int, float, etc. are provided in `builtin.go`, we can use this directly in any go file.
To convert string to any type, we must use `strconv` package.

```go
i := 42
f := float64(i)
u = uint(f)

str := "10"
parsedInt, _ := strconv.ParseInt(str)
```

Type of variable can be printed using `%T`.

```go
x := 5

fmt.Printf("type is %T", x)
```

**Declare multiple variables**

```go
var i, j, k int
b, f, s := true, 2.3, "four"
```

**Constants**
Use `const` keyword to declare a constant.
```go
const Pi = 3.14
```

**Computed constants**
We can compute the value of a constant at compile time using expressions involving other constants.
We cannot use a variable to compute the value of a constant.

```go
p := 5
const a = 10
const b = 20
const x = a * b
const y = x * p // This will give an error
```

#### If statement

```go
if x > 0 && x < 10 {
    fmt.Println("x is positive and less than 10")
} else if x == 0 {
    fmt.Println("x is zero")
} else {
    fmt.Println("x is negative")
}
```

If statement can start with a short statement to execute before the condition.

**Syntax**
```go
if <statement>; <condition> {
    // code
}

if v := math.Pow(x, n); v < lim {
    return v
}
```
> Note: v is only available in the `if` statement scope.

#### Switch statement
Similar to switch case but `break` is implicit. So automatic fall-through is not default.
We can use `fallthrough` keyword to execute all the cases which are below the matching one.

**Syntax:**
```go
switch statement; expression {
case expressionOne:
	statement...
case expressionTwo:
	statement...
default:
    statement...
}
```

**Example:**
```go
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X")
case "linux":
	fmt.Println("Linux")
default:
	fmt.Println("%s", os)
}
```

**Type switch**
- Type switch is used to check the type of an interface value.
- It is used to perform different actions based on the type of the interface value.
- It is similar to a regular switch statement, but the cases in a type switch specify types not values.
- We can also use if condition to check the type of an interface value.

**Syntax:**
```go
switch v := <interface-value>.(type) {
    case <type-1>:
        // code
    case <type-2>:
        // code
    default:
        // code
}
```

**Example:**
```go
func describe(s Shape) {
    switch v := s.(type) {
    case Circle:
        fmt.Println("Circle")
    case Rectangle:
        fmt.Println("Rectangle")
    default:
        fmt.Println("Unknown shape")
    }
}
```

#### Loops
`for` keyword is use to define loops in golang.

**Syntax:**
```go
for init; condition; post {
statements
}
```

**Example:**
```go
for i := 10; i > 0; i-- {
  fmt.Println(i) // Counts down from 10 to 1
}
```

**While loop**
For while loop, we just need to omit `init` and `post` statements.

**Syntax:**
```go
for condition {
	// body
}
```

**Infinite loop**
For infinite loop, we just need to omit all the statements.

**Syntax:**
```go
for {
	// body
}

// or
for true {
	
}
```

**Range loop**
By using the `range` keyword, a `for` loop can step through the items in a collection such as a array, map, slice, channel, or string.

**Syntax:**
```go
for index, value = range collection {
	// body
}
```

**Example:**

```go
numbers := []string{"One","Two","Three"}
for i, n := range numbers {
  fmt.Println(i,n)
}
```

**Break and continue**
The `break` and `continue` statements work in Go as they do in C and Java.

**Example:**
```go
for i := 0; i < 100; i++ {
    if i % 2 == 0 {
      continue
    }
    if i == 50 {
      break
    }
    fmt.Println(i)
}
```

**Logical operators**
Same as any other language

- *AND* - &&
- *OR* - ||
- *NOT* - !

**Bitwise operators**
Same as any other language

- *AND* - &
- *OR* - |
- *XOR* - ^
- *AND NOT* - &^
- *Left shift* - <<
- *Right shift* - >>

### Chapter 3: Functions

**Function declaration**

**Syntax:**
```go
func <function-name>(<parameter> <type>) <return-type> {
    // code
}

// short hand
func <function-name>(<parameter-1>, <parameter-2> <type>) <return-type> {
    // code
}
```

```go
func add(x int, y int) int {
    return x + y
}

func add(x, y int) int {
    return x + y
}
```

Go supports functions in function arguments.
```go
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}
```

**Multiple return values**

```go
func swap(x, y string) (string, string) {
    return y, x
}

func swap(x int, y int) int {
	return y, x
}
```

**Ignoring return values**

```go
// Ignoring the first return value
_, b := swap("hello", "world")
```

**Named return values**
- It is generally used to document the purpose of the return values.
- It is not recommended to use naked return as it can harm readability.
- Naked return should be used only in short functions.
- Named return values are initialized to default values and it is recommended to use them with functions having multiple return values.

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

**Example:**
```go
func getCoord() (x, y int) {
	return // x, y are first initialized to 0 and then returned
}

// This is equivalent to
func getCoord() (int, int) {
    var x int
    var y int
    return x, y
}
```

#### Higher order functions:
Go supports first class functions as functions can take another function(s) as an argument or returns another function, 
these type of functions are called higher order functions.

First class function is a function that can be treated like any other value.
Higher order function is a function that takes a function as an argument.
This is somewhat similar to callbacks in javascript.

**Example:**
```go
// higher order function
func aggregate(a, b, c int, arithmetic func (int, int) int) int { // arithmetic is a first class function
	return arithmetic(arithmetic(a, b), c)
}
```

#### Currying
Function currying is a practice of writing a function which takes function as an argument and return a new function.
Use case - Middlewares

**Example:**
```go
func selfMath(mathFunc func (int, int) int) func (int) (int) {
	return func (x int) int {
	    return mathFunc(x, x)	
    }
}
```

#### `defer` keyword
`defer` allows a function to be executed automatically just before its enclosing function returns.
Somewhat similar to `finally` block in java.

**Example:**
```go
func deleteUsers(users map[string]user, key string) {
    defer delete(users, name)
    user, ok := users[name]
    if !ok {
        return logNotFound
    }
    if user.admin {
        return logAdmin
    }
	return logDeleted
}
```

#### Closures:
Closure is a function that references variables from outside its own function body. The function may access and assign
to the referenced variables.

**Example:**
```go
func concatter() func (string) string {
	doc := ""
	return func (word string) string {
		doc += word + " "
		return doc
    }
}
```
> Note: In above function we are using doc variable, which belongs to concatter function, this might lead to memory leaks,
> so using closures can be expensive.

#### Anonymous functions
Functions with no name. It is useful when defining a function that will only be used once or to create quick closure.

**Example:**
```go
number := 10  
  squareNum := func() (int) {
  number *= number
  return number
}
x := squareNum()  
y := squareNum() 
```

### Chapter 4: Structs
Structs are used to group data together to represent an entity.

**Syntax:**
```go
type <struct-name> struct {
    <field-name> <type>
    <field-name> <type>
}
```

**Example:**
```go
type Vertex struct {
    X int
    Y int
}
```

#### Nested structs
- Structs can be nested inside other structs.
- Nested structs are useful when we want to group related data together.
- Nested structs can be accessed using the `.` operator.
- Nested structs can be used to represent complex data structures.

```go
type Circle struct {
    Center Vertex
    Radius float64
}
```

#### Anonymous structs
- It is used when we don't have reason to create a named struct or when we want to create a struct that is only used once.
- It prevents us from re-using a struct which is never intended to be used again. For example, in HttpHandler to define json response structure.

```go
vertex := struct {
    X int
    Y int
} {
    -1,
    0
}
```
This one is not that much common.

Generally, anonymous structs are used within named structs.

```go
type Circle struct {
    Center struct {
        X int
        Y int
    }
    Radius float64
}
```

#### Embedded structs
- It is used to compose a struct with another struct.
- It is used to reuse fields of the embedded struct.
- We can access the fields of the embedded struct same as it belong to top-level struct.
- Initialization of embedded struct is same as nested struct.

```go
type Circle struct {
    Vertex
    Radius int
}
```
> We can access `X` and `Y` fields of `Vertex` struct using `Circle` struct, `Circle.X` and `Circle.Y`.

**Initialization of embedded struct:**
```go
circle := Circle{
    Vertex: Vertex{1, 2},
    Radius: 5,
}

fmt.Println(circle.X, circle.Y, circle.Radius)
```

#### Methods in structs
- Methods are functions that are associated with a particular type.
- These are defined using a special receiver argument.
- Receiver can be a value or a pointer.

**Syntax:**
```go
func (<receiver> <type>) <method-name>(<parameters>) <return-type> {
    // code
}
```

**Example:**
```go
func (c Circle) area() float64 {
    return math.Pi * c.Radius * c.Radius
}

circle := Circle{Vertex{0, 0}, 5}
fmt.Println(circle.area())
```

### Chapter 5: Interfaces
- Interfaces are used to define a set of methods that a type must have.
- Interfaces are implemented implicitly.
- A type implements an interface if it provides definitions for all the methods in the interface.
- Interface should be concise and only contain the methods that are required.

**Syntax:**
```go
type <interface-name> interface {
    <method-name>(<parameters>) <return-type>
}
```

**Example:**
```go
type Shape interface {
    area() float64
}

type Circle struct {
    Radius float64
}

// Circle implements the Shape interface
func (c Circle) area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

#### Empty interface
- An empty interface is an interface that has zero methods.
- It is represented by `interface{}`.
- It is by default implemented by all types.

#### Named interface
- Named interfaces are used to provide meaningful names to interface's arguments.
- It is used to make the code more readable and maintainable.
- It is not necessary to create a named interface, but it is recommended to do so. Without a named interface, code works the same way.

**Example:**
```go
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int, error)
}
``` 

#### Type assertion
- Type assertion is used to extract the underlying value of the interface.
- It is used to check the type of an interface value.
- It is used to convert an interface value to another type.

**Syntax:**
```go
value, ok := <interface-value>.(<type>)
```
> If the type assertion is successful, `ok` will be true and `value` will contain the underlying value of the interface.

**Example:**
```go
type Shape interface {
    area() float64
}

type Circle struct {
    Radius float64
}

c, ok := s.(Circle) // s is of type Shape and c is of type Circle if the assertion is successful
```

### Chapter 6: Error handling
- Error is an interface which contains only a single method called `Error()`.
- Error method returns a string which explains the reason.
- If a method returns an error, we should handle that immediately.
- Error message string should not be capitalize and should not ends with punctuation mark

#### Error interface
```go
type error interface {
	Error() string
}
```

**Example:**
```go
x, y := 5, 0
r, err := divide(x, y)
if err != nil {
	fmt.Println("Error while dividing %v by %v", x, y)
	return
}
```

#### Formatting strings
`fmt.Sprintf()` is used to form a string with different format specifiers or "formatting verbs".

Different formatting verbs:
- %v	the value in a default format
- %t	the word true or false
- %c	the character represented by the corresponding Unicode code point
- %d	integer in base 10 form
- %f	decimal point but no exponent, e.g. 123.456
- %F    synonym for %f
- %s	the uninterpreted bytes of the string or slice
- %q	a double-quoted string safely escaped with Go syntax
- %T	a Go-syntax representation of the type of the value
- %p	pointer address in base 16 notation, with leading 0x

#### Custom errors
We can define our own errors by implementing error interface.

**Example:**
```go
type UserError struct {
	name string
}

func (e UserError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}
```

#### `errors` package
Go std library provides an errors package that provides utilities related to error.

**Example:**
```go
var err error = errors.New("something went wrong")
```

### Chapter 7: Arrays
In Go, an array is a homogeneous data structure (Fix type) and has a fixed-length. The type can be anything like integers, string or self-defined type.
It is similar to arrays in other languages.

**Syntax:**
```go
var identifier [len]type
```

**Example:**
```go
var x [5]int
for i := 0; i < len(x); i++ {
	x[i] = i
}
```

#### Multi-dimension arrays
It is possible to create multi-dimensional array

**Syntax:**
```go
var arrayName [x][y]variable_type
```

**Example:**
```go
var array [5][6]int
```

#### Slices
- Slice is a dynamically-sized, segmented view of an underlying array.
- This segment can be the entire array or a subset of an array.
- We define the subset of an array by indicating the start and end index. 
- Slices provide a dynamic window onto the underlying array.
- Slice is like reference to an array. Slice does not store any data. 
- If we change the elements of an array, it will also modify the underlying array. 
- If other slice is referencing the same underlying array, their value will also be changed.

**Syntax:**
```go
var <name> []<T> = <array>[<lower-bound>:<upper-bound] // upper bound is always exclusive

// or
<name> := <array>[<lower-bound>:<upper-bound]

// or
<name> := make([]<T>, 10)       // slice of length 10, capacity of 10 and all elements initialized to 0

// or
<name> := make([]<T>, 0, 10)    // slice of length 0 and capacity of 10
```

**Slice literals:**
Slice literal is like an array literal without any length.

```go
s := []struct {
  i int
  b bool
}{
  {1, true},
  {2, false},
  {3,true},
  {4, true},
  {5, false},
  {6, true},
}
fmt.Println(s)
```

**Omit Lower or Upper Bonds:**
In slice, we can omit the lower bond or the upper bonds. Zero is the default value of the lower or the upper bond.

```go
slice := []int{1, 2, 3, 4, 5}
x := slice[:3]  // slice[0:3]
y := slice[1:]  // slice[1:5]
```

**Length and capacity:**
A slice has length and capacity. The length is the number of stored elements and the capacity is the number of elements
of the underlying array counting from the beginning of the slice.

To get the length, we use `len(slice)` function and to get the capacity, we use `cap(slice)` function.

#### Command-line arguments
- The arguments passed from the console can be received by the Go program and it can be used as an input.
- The os.Args is used to get the arguments. The index `0` of os.Args contains the path of the program.
- The os.Args[1:] holds provided arguments.

```go
var s, arg string  
for i := 1; i < len(os.Args); i++ {
    s += arg + os.Args[i]+" "
}
fmt.Println(s)
```

#### Variadic functions
Functions which can take multiple arguments of same type but its size is not pre-determined.
This is similar to var-args in Java.

```go
func sum(nums ...int) int { // argument type is nothing but slice of int
	total := 0
	for _, num := range nums {
        total += num
    }
	return total
}
```

**Spread operator**
When we want to pass slice to a variadic function, we can use `<slice>...`
```go
slice := []int{1, 2, 3, 4, 5}
total := sum(slice...)
```

### Chapter - 8: Maps
- Similar to any other language, used to store key-value pairs.
- map keys can be any type that is comparable.
- boolean, numeric, string, pointer, channel, interface types and structs or arrays that contains only those types can be the key
- slices, map, functions cannot be the key.


**Syntax:**
```go
m := make(map[<type-1>]<type-2>)

// or
m := map[<type-1>]<type-2>{
	"A": 1,
	"B": 2
}
```

**Operations**
- Insert: `m[key] = value`
- Get: `element = m[key]`
- Delete: `delete(m, key)`
- Check if element exists: `element, ok := m[key]`

### Chapter - 9: Pointers
- Pointer are nothing but just a reference to a value
- Default value of a pointer is `nil`
- Unlike C, Go does not support pointer arithmatic but we can use `unsafe` to create unsafe pointers.

**Syntax:**
```go
var <variable-name> *<type>
```

**Example:**
```go
var x int = 5
var y *int = &x // creating a pointer of type *int
*y = 10         // de-referencing the pointer and changing the value at a specific address 
```

It is common as receiver of a function and we can directly call the same function using same object of the structure.
Implicitly compiler will cast the type to a pointer.

**Example:**
```go
type Circle struct {
	x int
	y int
	radius int
}

func (c *Cirle) grow() {
	*c.radius *= 2 
}

func main() {
	c := Circle {
		x: 1,
		y: 2,
		radius: 5,
    }
	c.grow()    // directly calling grow function, implicitly it will cast the type Circle to *Circle
}
}
```
