# Constants

### Agenda
* Naming convention
* Typed constants
* Untyped constants
* Enumerated constants
* Enumeration expressions

## Naming convention
All constants are going to be preceded with `const` keyword and that's going to let the compiler know that what we're trying to work with. If you've come from other languages, you might be expecting that we're going to name our constants something like `MY_CONST`, where we're going to have all uppercase letters and separate the words with underscores. The problem with that is, if we do that, then the first letter is going to be uppercase. As you remember, if we've got an uppercase first letter, that's going to mean that the constant is going to be exported.

We're actually going to name our constants the same way we named our variables. **Typed constants** is created very similar to a typed variable. So we can start with the `const` keyword, then the name of our constant and then we're going to list the type, and then we can set it equal to a value. Now, the reason it's a constant and not a variable is it has to remain constant. You can't set your constant to something that has to be determined at runtime and that includes things like setting it equal to flags that you pass into your application.
```go
func main() {
   const myConst int = 42
   fmt.Printf("%v, %T", myConst, myConst) // Output: 42, int

   myConst = 27
   fmt.Printf("%v, %T", myConst, myConst) // It will throw an error of: cannot assign to myConst
}
```

Constants can be made up of any the primitive types. If you run this, you'll see that all of those print out exactly the way that we expect.
```go
func main() {
   const a int = 14
   const b string  = "foo"
   const c float32 = 3.14
   const d bool = true

   fmt.Printf("%v\n", a) // Output: 14
   fmt.Printf("%v\n", b) // Output: foo
   fmt.Printf("%v\n", c) // Output:  3.14
   fmt.Printf("%v\n", d) // Output: true
}
```

Another characteristic that constants have in common with variables is they can be shadowed. Now, we've got a constant called `a` declared at the package level and then we got a constant in the main function, that's called `a` and it's an integer type. We see that the inner declaration wins. So not only can we change the value of the constant, but we can also change the type because the inner constant shadows the outer constant.
```go [go-shadow]
package main

import "fmt"

const a int16 = 27
func main() {
   const a int = 14
   fmt.Printf("%v, %T\n", a, a) // Output: 14, int
}
```

When working with constants, they work very similar to variables when we're using them in operations. If we run this, we get the ability to add a constant to a variable and the result is going to be a variable. Since the constant and the variable are of the same type, we can perform the addition operation on there. If the constant is of a different type, then you'll get an error that says `mismatched types`.
```go
package main

import "fmt"

func main() {
   const a int = 42
   var b int = 27
   fmt.Printf("%v, %T\n", a + b, a + b) // Output: 69, int
}
```

We can use the compilers ability to infer the type for us. If we run this, we can see that the constant `a` is infered to be an integer with the value of 42.
```go
func main() {
   const a = 42
   fmt.Printf("%v, %T", a, a) // Output: 42, int
}
```

In this case, the operation succeeds, which might be a little bit confusing. The reason that it works is because what the compiler is actually doing, when it sees this constant it's basically replacing every instance. So the way the compiler sees this program is it sees it like this `42 + b`. The `42` is being interpreted as an integer 16. The compiler is going to look for every time that we use the symbol `a`, and it's going to replace that with value of the constant and we can do the implicit conversions when we're working with constants, which is something that we can't really do when we're working with variables.
```go
func main() {
   const a = 42
   var b int16 = 27
   fmt.Printf("%v, %T", a + b, a + b) // Output: 69, int16
}
```

## Enumerated constants

What is `iota`? `iota` is a counter that we can use when we're creating what are called **enumerated constants**.
```go
const a = iota

func main() {
   fmt.Printf("%v, %T", a, a) // Output: 0, int
}
```

In this example, having an enumerated constant isn't terribly valuable but one of the things that we can do with constants is, we can actually work with them in a constant block. Now, we are using `iota` three times and when we get the result we actually see `iota` is changing its value as the constants are being evaluated.
```go
const (
   a = iota
   b = iota
   c = iota
)

func main() {
   fmt.Printf("%v\n", a) // Output: 0
   fmt.Printf("%v\n", b) // Output: 1
   fmt.Printf("%v\n", c) // Output: 2
}
```

Another special feature that we can take advantage of with `iota` is that if we don't assign the value of a constant after the first one, then the compiler is going to try and infer the pattern of assignments. Since we've established a pattern for how to name the constants in this block, when we run, we actually get the same result. That's because the compiler is going to apply the same formula. Now, that value of `iota` is scoped to that constant block. 
```go
const (
   a = iota
   b
   c
)

func main() {
   fmt.Printf("%v\n", a) // Output: 0
   fmt.Printf("%v\n", b) // Output: 1
   fmt.Printf("%v\n", c) // Output: 2
}
```

In here, we create another constant block. Then what we're gonna find is `iota` resets to zero. So `iota` is scoped to a constant block and what that lets you do is you can actually created related constants together, ensure that they have different values and then if you have another set of related constants, you can start another constant block and ensure that they have unique values, but allow duplication between the values in one constant block in another.
```go
const (
   a = iota
   b
   c
)

const (
   a2 = iota
)

func main() {
   fmt.Printf("%v\n", a) // Output: 0
   fmt.Printf("%v\n", b) // Output: 1
   fmt.Printf("%v\n", c) // Output: 2
   fmt.Printf("%v\n", a2) // Output: 0
}
```

What we're doing here is that we're setting up a constant block, where maybe we're trying to store the specialty of veterinarians in a veterinarian clinic. In the main function, there's a variable and we're setting its value that is equal to `catSpecialist`. If we check to see if the specialist type is a cat specialist, it gets the true value. And this is a very common use for enumerated constants.
```go
const (
   catSpecialist = iota
   dogSpecialist
   snakeSpecialist
)

func main() {
   var specialistType int = catSpecialist
   fmt.Printf("%v\n", specialistType == catSpecialist) // Output: true
}
```

If we declare the variable and don't initialize, it will give you a value of false which makes sense on the first print statement. On the second print statement, even though we haven't specified a specialist type, it does show up as the value cat specialist. So what do we do about it?
```go
const (
   catSpecialist = iota
   dogSpecialist
   snakeSpecialist
)

func main() {
   var specialistType int
   fmt.Printf("%v\n", specialistType == dogSpecialist) // Output: false
   fmt.Printf("%v\n", specialistType == catSpecialist) // Output: true
}
```

First, use the zero value of the constant as an error value. Now, when we check if the specialist type is a cat specialist, we get the value of false because cat specialist is equivalent to the integer value `1`. This is a very valuable approach if you want to check to see if a value hasn't been assigned to a constant yet. If your application guards against that and there's no reasonable way for this to happen, then you can take advantage of `_` symbol, which is goes one and only write only variable. Now, what is the value of a write only variable? With `iota`, we have to have a zero value, we always have to start with zero. If we don't care about zero, then we don't have any reason to assign the memory to it so we can use `_` symbol.
```go
const (
   errorSpecialist = iota
   // or you can use this: _ = iota
   catSpecialist
   dogSpecialist
   snakeSpecialist
)

func main() {
   var specialistType int
   fmt.Printf("%v\n", specialistType == catSpecialist) // Output: false
}
```

What we have here is that we have an example of a constant block that's giving you constants that are equivalent to kilobyte, megabyte, gigabyte and so on. In our main function, we've initialize the `fileSize` to some arbitrary value and then a print statement. This is basically going to format a result to print two decimal places and then the literal string GB afterward. So the `%.2f` is that you're expecting to format a floating point number and going to give it two decimal places.

You notice that the constant block is set equal to one and then we're going to bit shift that value `10 * iota`. So the first time we're going to bit shift `10 * 1`, so we're basically going to multiply this by 2^10 and then we're going to multiply by 2^100 for the megabyte, then 2^1000 for gigabyte and so on. When we run this, we get a really convenient way to format an arbitrary file size into a human readable format.
```go
const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
   fileSize := 4000000000.
   fmt.Printf("%.2fGB", fileSize/GB) // Output: 3.73GB
}
```

Another thing that can be very valuable to do is using bit shifting in order to set `boolean` flags inside a single byte. Let's say that this has users and those users have certain roles. Inside of the constant block, we're defining various roles that we can have. In order to define these constants, what we're doing is we're setting the value to 1 bit shifted by `iota`. What we have is each one of these constants is going to occupy one location in a byte. In the main function, we're defining the roles in a single byte.

When we run this, we see that we've encoded eight access roles for user into a single byte of data. So we're able to store this information efficiently. If you wanna see if this user is an admin, you can take the constant, `isAdmin & roles`, and what that's going to do is that's going to apply what's called **bit mask**. So only the bits that are set in the `isAdmin` constant and our `roles` are going to be left as true, which means if we're an admin, we're going to have the value one set at that first bit.

So we can very quicklu and very efficiently store a lot of different information about what roles and access rights a user might have, and a simple byte. Having this constant defined with a numeration expression makes it really fast and really efficient and really clear in our application.
```go
const (
   isAdmin = 1 << iota
   isHeadquarters
   canSeeFinancials

   canSeeAfrica
   canSeeAsia
   canSeeNorthAmerica
   canSeeSouthAmerica
)

func main() {
   // isAdmin: 00000001; canSeeFinancials: 100; canSeeEurope: 100000
   var roles byte = isAdmin | canSeeFinancials | canSeeEurope
   fmt.Printf("%b\n", roles) // Output: 100101
   fmt.Printf("Is Admin? %v", isAdmin & roles == isAdmin) // Output: Is Admin? true
   fmt.Printf("Is HQ? %v", isHeadquarters & roles == isHeadquarters) // Output: Is HQ? false
}
```

## Summary
* Immutable, but can be shadowed
* Replaced by the compiler at compile time
   * Value must be calculable at compile time
* Named like variables
   * PascalCase for exported constants
   * camelCase for internal constants
* Typed constants work like immutable variables
   * Can interoperate only with same type
* Untyped constants work like literals
   * Can interoperate with similar types
* Enumerated constants
   * Special symbol `iota` allows related constants to be created easily
   * `Iota` starts at 0 in each `const` block and increment by one
   * Watch out of constant values that match zero values for variables
* Enumerated expressions
   * Operations that can be determined at compile time are allowed
      * Arithmetic
      * Bitwise operations
      * Bit shifting