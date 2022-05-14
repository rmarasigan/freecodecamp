# Variables

### Agenda
* Variable declaration
* Redeclaration and shadowing
* Visibility
* Naming conventions
* Type conversions

## Introduction
There are three different ways to declare variables in Go.

1. Declare the variable itself.
Declaring variable called `i` with a data type of `integer`. The program that runs up above this `fmt` statement can actually influence the value that's printed out to the console. For this, there are going to be times when you want to declare a variable but you're not ready to initialize it yet. For example, if you want to declare variable in the scope of this main function, and then you have a loop or something that sets up a local variable scope and that's where the variable is actually going to be assigned.
```go
func main() {
   var i int
	i = 42
   i = 27
	fmt.Println(i) // Output: 27
}
```

2. Initialize a variable where we combine the two lines.
So we can initalize the variable `i` as an `integer` and assign the value 42 in the same line. In here, it is valuable if go doesn't have enough information to actually assign the type that you really want assign to it. So if you need a little bit of control, then this syntax is going to be valuable for you.
```go
func main() {
   var i int = 42
   fmt.Println(i) // Output: 42
   fmt.Printf("%v, %T", i, i) // Output: 42, int
}
```

3. Go compiler can actually figure out what data type it needs to have.
```go
func main() {
   i := 99
   fmt.Println(i) // Output: 99
}
```

## Declaring variable at the pakcage level
When you're doing it at the package level, you cannot use the `:=` syntax. You actually have to use the full declaration syntax.
```go
package main

import "fmt"

var i int = 42

func main() {
   fmt.Println("%v, %T", i, i) // Output: 42, int
}
```

Another way is to create a block of variables that are declared together. But declaring the way like this, things are actually a little bit cluttered. Instead of this, we can actually wrap this whole in a single `var` block.
```go
package main

import "fmt"

var actorName string = "Elisabeth Sladen"
var companion string = "Sarah Jane Smith"
var doctorNumber int = 3
var season int  = 11
```

We can wrap it in a section with a `var` block. We can have multiple variable blocks at the package level and that's just going to allow you to organize your application.
```go
package main

import "fmt"

var (
   actorName string = "Elisabeth Sladen"
   companion string = "Sarah Jane Smith"
   doctorNumber int = 3
   season int  = 11
)
```

## Redeclaring variable
If you try to run this, you're going to get an error saying `no new variables on left side of :=`.
```go
package main

import "fmt"

var i int = 42

func main() {
   var i int = 42
   i := 13
   fmt.Println(i)
}
```

You can re-assign the value of `i` but can't declare a new variable and that's because the variables are already declared (`var i int = 42`) and you can't declare the variable twice in the same scope.
```go
package main

import "fmt"

var i int = 42

func main() {
   var i int = 42
   i = 13
   fmt.Println(i) // Output: 13
}
```

Even though `i` is declared twice in the application, one at the package level and one inside of the main function, that's okay. What happens is that the variable with the innermost scope actually takes precedence. So this is called **shadowing**. So the package level `i` is still available, but it's being hidden by the declaration in the main function.
```go
package main

import "fmt"

var i int = 27

func main() {
   fmt.Println(i) // Output: 27
   var i int = 42
   fmt.Println(i) // Output: 42
}
```

Another thing about variables in go is that they always need to be used. If you run the program, you're going to get an error saying `j declared and not used`
```go
package main

import "fmt"

func main() {
   var i int = 42
   // Instantiating a variable
   j := 13
   fmt.Println(i) // Output: 42
}
```

## Variable Naming
Two sets of rules that you're going to keep track of.

1. How naming controls visibility of the variable
   - Lowercase variables are scoped to the package, which means anything that consumes the package can't see it and can't work with it.
   - Uppercase letter, then that's what's going to trigger the go compiler to expose the variable to the outside world.
   - **3 levels of visibility**
      - Package level, Lowercase: scoped to the package; so any file in the same package can access that variable
      - Package level, Uppercase: it's export at the front of the package and it's globally visible
      - Block scope: declared inside the function, that variable is scoped to the block and it's never visible outside of the block itself

2. Naming conventions
   - The length of the variable name should reflect the life of the variable
   - If you're declaring a variable that you use for a very long time, then it's a best practice to use a longer name
   - If you're working with a package level variable and is going to be used in quite a few other places, then that's where you're going to use the most verbose variable name
      - Keep your names as short as you can
      - Make sure that the names of those exported or package level variables are clear enough
   - Acronyms
      - The best practice is actually to keep these acronyms as all uppercase and the reason is for readability
         ```go
         func main() {
            var URL string = "https://google.com"
         }
         ```

## Type Conversions
Treating `i` as a floating point number and assign that value to `j`. The way that we do that is by using this conversion function, `float32`, and being used as a function.

```go
func main() {
   var i int = 42
   fmt.Printf("%v, %T\n", i, i)

   var j float32
   j = float32(i)
   fmt.Printf("%v, %T\n", j, j) // Output: 42, float32
}
```

Another example is, when we try and convert an integer into a string. First, you need to understand how strings work with Go. A string is just an alias for a stream of bytes. What happens is that it looks for what Unicode character is set at the value 42 and that happens to be an `*`.
```go
func main() {
   var i int = 42
   fmt.Printf("%v, %T\n", i, i)

   var j string
   j = string(i)
   fmt.Printf("%v, %T\n", j, j) // Output: *, string
}
```

If you want to convert strings and numbers, then you're going to need to use the string conversion package.
```go
import (
   "fmt"
   "strconv"
)

func main() {
   var i int = 42
   fmt.Printf("%v, %T\n", i, i)

   var j string
   j = strconv.Itoa(i)
   fmt.Printf("%v, %T\n", j, j) // Output: 42, string
}
```

## Summary
* Variable declaration
   * `var foo int`
   * `var foo int  = 42`
   * `foo := 42`
* Can't redeclare variables, but can shadow them
* All variables must be used
* Visibility
   * Lowercase first letter for package scope
   * Uppercase first letter to export
   * No private scope
* Naming conventions
   * Pascal or camelCase
      * Capitalize acronyms (HTTP, URL)
   * As short as reasonable
      * Longer names for longer lives
* Type conversion
   * destinationType(variable)
   * Use `strconv` package for strings