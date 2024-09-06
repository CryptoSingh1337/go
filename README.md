# Learning Golang

### Chapter 1: Go runtime
Go runtime consists of several key components, including the scheduler, garbage collector, memory allocator, and stack management.
It is attached with every executable binary.

Unlike Java, go does not uses any type of virtual machine so go runtime handles memory management.

**Run command**
```bash
go run <file-name>
```

**Go file syntax**
- Main file which contains main function must include `package main` at the top.
- This file should contain a main function like `func main() {}`.

### Chapter 2: Basic

#### **Data types**

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

#### **Declaring a variable**

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

```go
i := 42
f := float64(i)
u = uint(f)
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

#### **If statement**

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
```

```go
if v := math.Pow(x, n); v < lim {
    return v
}
Note: v is only available in the `if` statement scope.
```

#### **Switch statement**
Similar to switch case

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

#### **Loops**
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
func <function-name>(<parameter> <type>) <return-type> {
    // code
}

// short hand
func <function-name>(<parameter-1>, <parameter-2> <type>) <return-type> {
    // code
}

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

Example:
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

**Nested structs**
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

**Anonymous structs**
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

**Embedded structs**
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

Initialization of embedded struct:
```go
circle := Circle{
    Vertex: Vertex{1, 2},
    Radius: 5,
}

fmt.Println(circle.X, circle.Y, circle.Radius)
```

**Methods in structs**
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

**Empty interface**
- An empty interface is an interface that has zero methods.
- It is represented by `interface{}`.
- It is by default implemented by all types.

**Named interface**
- Named interfaces are used to provide meaningful names to interface's arguments.
- It is used to make the code more readable and maintainable.
- It is not necessary to create a named interface, but it is recommended to do so. Without a named interface, code works the same way.

**Example:**
```go
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int, error)
}
``` 

**Type assertion**
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

**Error interface**
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

**Formatting strings**
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

**Custom errors**
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

**Errors package**
Go std library provides an errors package that provides utilities related to error.

**Example:**
```go
var err error = errors.New("something went wrong")
```
