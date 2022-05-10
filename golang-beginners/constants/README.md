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
```go
package main

import "fmt"

const a int16 = 27
func main() {
   const a int = 14
   fmt.Printf("%v, %T\n", a, a) // Output: 14, int
}
```

When working with constants, they work very similar to variables when we're using them in operations. If we run this, we get the ability to add a constant to a variable and the result is going to be a variable. Since the constant and the variable are of the same type, we can perform the addition operation on there.
```go
package main

import "fmt"

func main() {
   const a int = 42
   var b int = 27
   fmt.Printf("%v, %T\n", a + b, a + b) // Output: 69, int
}
```