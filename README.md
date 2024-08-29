# Go

### Go runtime
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

### Basic

**Data types**

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

**Declaring a variable**

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

**If statement**

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

### Functions
