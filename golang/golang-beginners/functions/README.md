# Functions

### Agenda
* Basic syntax
* Parameters
* Return values
* Anonymous functions
* Functions as types
* Methods

## Basic Syntax
A **function** is a group of statements that together perform a task. Every Go program has at least one function, which is **`main()`**. A ***function declaration*** tells the compiler about a function name, return type, and parameters. A ***function definition*** provides the actual body of the function.

**Syntax**
```go
func function_name(parameter_list) return_types {
  // Body of the function
}
```

Here are the parts of a function:
* **`func`**: It starts the declaration of a function
* **`function_name`**: It is the actual name of the function. The function name and parameter list together consitute the function signature.
* **`parameter_list`**: A parameter is like a placeholder. When a function is invoked, you pass a value to the parameter. This value is referred to as actual parameter or *argument*.
* **`return_types`**: A function may return a list of values. The `return_types` is the list of data types of the values the function returns.
* **Function body**: It contains a collection of statements that define what the function does.

Now, the way that a Go application is structured is we always have to have an entry point in the application. The entry point of the Go application is always in *`package main`*, and within that main package, we have to have a function called *`main`* that takes no parameters and returns no values.

```go
func main() {
  fmt.Println("Hello, playground")
}
```

When we run the application, it starts with the `fmt.Println("Hello, playground")` line and we print out the string `"Hello, playground"`.

Output:
```
Hello, playground
```

Every Go application starts this way, and we see here the most basic function that we can create in the Go language.

> #### Naming Conventions
>
> * A function name must begin with a letter and can have any number of additional letters and numbers.
> * A function name cannot start with a number.
> * If the function name starts with an ***uppercase letter***, it will be exported to other packages.
> * If the function name starts with a ***lowercase letter***, it won't be exported to other packages, but you can call the function within the same package.
> If a function name consists of multiple words, each word after the first should be capitalized (e.g. employeeName, EmployeeAddress, etc.)
> * Function names are case-sensitive.

## Parameters
Information can be passed to functions through arguments. An argument is just like a variable. ***Arguments*** are specified after the function name, inside the parentheses. You can add as many arguments as you want, just separate them with a comma.

Here the `main` function is now calling into another function called `sayMessage`. The `sayMessage` function takes in the `msg` parameter that is of type string. So when you're defining a function that takes parameters, the parameters go between the two parentheses and they are described like any other variable declaration, except you don't need the `var` keyword.

When we want to call that function, we have to pass in the value of that parameter and that is called an **argument**. So we're going to pass in the argument `"Hello Go!"` and then inside our function, we are printing out the message that gets passed in.

```go
func main() {
  sayMessage("Hello Go!")
}

func sayMessage(msg string) {
  fmt.Println(msg)
}
```

When we run this, we see that `"Hello Go!"` prints out as a result.

Output:
```
Hello Go!
```

The `msg` parameter is not special in any other way than the fact that it is passed into the function. It is treated as a local variable, just like any other local variable that we might create.

If we create another variable and call it `greeting`, that is treated exactly the same way as the `msg` variable. The only difference between the two is, the `msg` variable can be passed in from the outside, and of course, the `greeting` variable was created locally.

```go
func main() {
  sayMessage("Hello Go!")
}

func sayMessage(msg string) {
  greeting := "Hello Go!"
  fmt.Println(msg)
}
```

Now, you are not constrained to just pass a single parameter in, you can pass multiple parameters in. Here, we've extended the `sayMessage` function to take two parameters.

As we can see, we can pass as many parameters as we want. We're just going to provide a comma-delimited list of those parameter names and their types.

```go
func main() {
  for i := 0; i < 5; i++ {
    sayMessage("Hello Go!", i)
  }
}

func sayMessage(msg string, idx int) {
  fmt.Println(msg)
  fmt.Println("The value of the index is ", idx)
}
```

Output:
```
Hello Go!
The value of the index is  0
Hello Go!
The value of the index is  1
Hello Go!
The value of the index is  2
Hello Go!
The value of the index is  3
Hello Go!
The value of the index is  4
```

Often, when we're defining a function, we're going to pass in multiple parameters of the same type. So we're going to be tempted to have a syntax that is a comma-delimited list of variables and then the type at the end.

What the compiler is going to do, is going to infer that every variable in the comma-delimited list has the same type.

```go
func main() {
  sayGreeting("Hello", "Stacey")
}

func sayGreeting(greeting, name string) {
  fmt.Println(greeting, name)
}
```

Output:
```
Hello Stacey
```

## Pass by Value VS Pass by Pointer
Every time a variable is passed as parameter, a new copy of the variable is created and passed to called function or method. The copy is allocated at a different memory address. In case a variable is *passed by pointer*, a new copy of pointer to the same memory address is created.

Here, we are passing in the `name` variable by value. That means, that the Go runtime is going to copy the data that is in the `name` variable and provided it to `sayGreeting`.

```go
func main() {
  name := "Stacey"
  greeting := "Hello"
  sayGreeting(greeting, name)
  fmt.Println(name)
}

func sayGreeting(greeting, name string) {
  fmt.Println(greeting, name)

  name = "Ted"
  fmt.Println(name)
}
```

What we would expect is when we change the value of the `name` variable, it should have an effect, and we should print `Ted` out. Since this is a copy of the `name` variable, we shouldn't have any effect.

This is a safe way to pass data into a function. We can be assured by *passing by value*, that the data is not going to be changed when you pass it in.

Output:
```
Hello Stacey
Ted
Stacey
```

If we change the variables type to *passing in pointers* by adding a pointer, and then passing in the address of the variables, and then dereferencing the pointers in the `sayGreeting` function.

Now, we are passing pointers to our variables around the application. Instead of working with a copy of the variables, we are working with a pointer to the variables.

```go
func main() {
  name := "Stacey"
  greeting := "Hello"
  sayGreeting(&greeting, &name)
  fmt.Println(name)
}

func sayGreeting(greeting, name *string) {
  fmt.Println(*greeting, *name)

  *name = "Ted"
  fmt.Println(name)
}
```

We see that we have to change the variable not only in the scope of the function but in the calling scope as well. So by *passing in a pointer*, we have manipulated the parameter we passed in.

Output:
```
Hello Stacey
Ted
Ted
```

What are the reasons why would you want to do this? First of all, a lot of times our functions do need to act on the parameters that are passed into them, and so passing in pointers is the only way to do that. The other reason is, passing in a pointer is often much, much more efficient than passing in a whole value.

However, if we're passing in a large data structure, then passing in the value of that data structure is going to cause that entire data structure to be copied every single time. In that case, we might decide to pass in a pointer simply for a performance benefit.

## Variadic parameters
A **variadic function** is a function that accepts a variable number of arguments. It is possible to pass a varying number of arguments of the same type as referenced in the function signature. To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which shows that the function may be called with any number of arguments of this type.

Here, we can see an example of a *variadic parameter*. In this case, we've got a generic `sum` function and passing in the numbers one through five. Now, we're not receiving five variables in `sum` function. Instead, we've got one variable and is preceded its type with three dots (`...`).

```go
func main() {
  sum(1, 2, 3, 4, 5)
}

func sum(values ...int) {
  result := 0
  fmt.Println(values)

  for _, v := range values {
    result += v
  }

  fmt.Println("The sum is", result)
}
```

The Go runtime will take in all of the last arguments that are passed in, and wrap them up into a slice that has the name of the variable that we have. Inside of the `sum` function, we're going to print out what that `values` object and we're going to go ahead and add up all the `values`. Since it is going to act like a slice, we can use a `for` loop and `range` over those values.

Output:
```
[1 2 3 4 5]
The sum is 15
```

If we run it, we see that we have a slice that is printed out, the sum is 15. We got that result print out properly and there is no problem at all. Now, when you are using a *variadic parameter*, you can only have one and it has to be the last one.

## Return values
We can also return value from a function and use it anywhere in our program.

Right after the parameter list and before the opening curly brace, we've listed the returned value(s) type. In this case, we're expecting to return an integer. Inside the `sum` function, we're going to use the **`return`** keyword, and then we're going to return the value of the variable that we've been building up throughout the course of the function.

```go
func sum(values ...int) int {
  result := 0
  fmt.Println(values)

  for _, v := range values {
    result += v
  }

  return result
}

func main() {
  s := sum(1, 2, 3, 4, 5)
  fmt.Println("The sum is ", s)
}
```

In the `main` function, we can catch the return value by declaring a variable and setting it equal to the result of the `sum` function. The `s` is going to be an integer type because that is what was returned out of the function, and then we can work with that integer.

Output:
```
[1 2 3 4 5]
The sum is 15
```

Another feature that Go has that is pretty rare in a lot of languages is the ability to return a local variable as a pointer.

If we look at the `*int`, we are returning a pointer to an integer, and instead of returning the result, we are returning the address of the `result`. So `s` is now a pointer, and we can change it to dereference operation.

```go
func sum(values ...int) *int {
  result := 0
  fmt.Println(values)

  for _, v := range values {
    result += v
  }

  return &result
}

func main() {
  s := sum(1, 2, 3, 4, 5)
  fmt.Println("The sum is ", *s)
}
```

Output:
```
[1 2 3 4 5]
The sum is 15
```

### Named Return Values
To declare the named result or return parameters, just use the return type part of the function signature.

**Syntax**
```go
func function_name(parameter_list) (result_parameter data_type) {
  // Body of the function
}
```

Notice that we've changed the return value, we've got a set of parenthesis, and we've got a name for the return value and a type for it. When you do this, it is basically a syntactic sugar for declaring a `result` variable. The `result` variable will be available in the scope of the `sum` function, and that value is going to be implicitly returned. We don't have to specify the name of the return variable, we just have to tell it to return.

```go
func sum(values ...int) (result int) {
  fmt.Println(values)

  for _, v := range values {
    result += v
  }

  return
}

func main() {
  s := sum(1, 2, 3, 4, 5)
  fmt.Println("The sum is ", s)
}
```

Output:
```
[1 2 3 4 5]
The sum is 15
```

### Multiple Return Values
Functions in Golang can return multiple values, which is a helpful feature in many practical scenarios.

**Syntax**
```go
func function_name(parameter_list) (result_parameter_1 data_type, result_parameter_2 data_type) {
  // Body of the function
}
```

Here, we've created a `divide` function that takes in two parameters, `a` and `b`, that are `float64`. It is going to divide them, and it is going to return that result.

```go
func divide(a, b float64) float64 {
  return a / b
}

func main() {
  d := divide(5.0, 3.0)
  fmt.Println(d)
}
```

Output:
```
1.6666666666666667
```

What happens if we pass in a zero? If we run it, we get an unknown result and we can't work with that in our application. We're going to probably cause some sort of a failure down the line.

```go
func divide(a, b float64) float64 {
  return a / b
}

func main() {
  d := divide(5.0, 0.0)
  fmt.Println(d)
}
```

Output:
```
+Inf
```

What we can do is, we're going to return an object of type error. So we're going to return the intention of the function as the first value, and then we're going to return an error object in case something went wrong.

```go
func divide(a, b float64) (float64, error) {
  // Checking for the error conditions to return
  // as soon as possible with the error value
  if b == 0.0 {
    return 0.0, fmt.Errorf("Cannot divide by zero")
  }

  // Returning the result of a calculation and a
  // nil error
  return a / b, nil
}

func main() {
  d, err := divide(5.0, 0.0)

  // Check if there's an error and return from our main function
  // and exit our application
  if err != nil {
    fmt.Println(err)
    return
  }

  // If we don't have an error, print out the result
  fmt.Println(d)
}
```

We see that we now get an error "`Cannot divide by zero`". The main function doesn't explode on us, we have something that we can work with but if we put in a valid value, then we're on that other path of execution, and we get the return value.

Output:
```
Cannot divide by zero
```

## Anonymous Functions
An **anonymous function** is a function that was declared without any named identifier to refer to it. Anonymous functions can accept inputs and return outputs, just as standard functions do.

In this example, we are actually declaring a function on the fly and this is called an *anonymous function*. Notice that we are starting with the `func` keyword, we've got the parameters, and we've got the opening and closing curly brace, but it doesn't have the function name.

When we're doing this, this is what is called an *anonymous function* and is the basic structure of a function when we're not working with functions in traditional scope. Inside the function body, we're printing out the message "`Hello Go!`", and then we've got accompanying the closing curly brace, with the parameters (`()`). Now, these parameters (`()`) are basically going to invoke the function. So this is an immediately invoked function, we're defining it and executing it at exactly the same time.

```go
func main() {
  func() {
    fmt.Println("Hello Go!")
  }()
}
```

Output:
```
Hello Go!
```

We can also work with functions as variables. In this case, we've declared an anonymous function and assigned it to the variable `f`, and then we can execute `f` by just invoking it like any other function.

```go
func main() {
  f := func() {
    fmt.Println("Hello Go!")
  }

  f()
}
```

Output:
```
Hello Go!
```

Now that we've got the function defined as a variable, it is free to pass around the application. You might ask yourself what is the signature for the function?

We have the `func` keyword and then an open and close parenthesis. We don't have any parameters, and we don't have any return types. So the type signature for the `f` variable is simply `func` without parameters.

```go
func main() {
  var f func() = func() {
    fmt.Println("Hello Go!")
  }

  f()
}
```

Output:
```
Hello Go!
```

## Methods
Go methods are similar to Go functions with one difference. The method contains a receiver argument in it. With the help of the receiver argument, the method can access the properties of the receiver.

**Syntax**
```go
func (receiver_name Type) method_name(parameter_list) (return_type) {
  // Body of the function
}
```

In this example, we've got a `struct` called **`greeter`**, and it has two fields, the `greeting`, and `name`, and then we've got a method on it.

```go
type greeter struct {
	greeting string
	name     string
}

// The method declartion looks like a function
// except for the `(g greeter)` part which is the
// receiver argument and this is what makes this
// function into a method.
func (g greeter) greet() {
  fmt.Println(g.greeting, g.name)
}


func main() {
  g := greeter{
    greeting: "Hello",
    name: "Go",
  }

  // Calling the function preceding it with
  // the struct that we have up there. This
  // is how we do the method invocation.
  g.greet()
}
```

Output:
```
Hello Go
```

A ***method*** is just a function that is executed in an unknown context. The unknown context in Go is `any` type. Commonly, we can use `struct` but you can use `any` type.

When we declare the `greet` method, we're going to get access to the `greeter` struct. So what is going to happen when we call the `greet` method it is going to get a copy of the `greeter` object, and that is going to be given the name `g` in the context of the said method. When we print out, we can access the fields on the `greeter` object.

Notice that we're specifying `greeter` as a value type (`(g greeter)`), we don't have a pointer. This is what is called a *value receiver*. The received object in the `greet` method is the value `greeter`. So what that means is just like any other time that we're working with values, we are getting a copy of the struct, but we're not going to get the `struct` itself.

As you might expect, there is another option that we have, and that is to pass the *pointer receiver*. Now, we're able to manipulate the underlying data.

```go
type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
  fmt.Println(g.greeting, g.name)
  g.name = ""
}


func main() {
  g := greeter{
    greeting: "Hello",
    name: "Go",
  }

  g.greet()
  fmt.Println("The new name is:", g.name)
}
```

Output:
```
Hello Go
The new name is: 
```

## Summary

#### Basic syntax
```go
func foo() {
  ...
}
```

#### Parameters
* Comma delimited list of variables and types
  ```go
  func foo(bar string, baz int)
  ```
* Parameters of same type list type once
  ```go
  func foo(bar, baz int)
  ```
* When pointers are passed in, the function can change the value in the caller
  * This is always true for data of slices and maps
* Use variadic parameters to send list of same types in
  * Must be last parrameter
  * Received as a slice
  ```go
  func foo(bar string, baz ...int)
  ```

#### Return values
* Single return values just list type
  ```go
  func foo() int
  ```
* Multiple return value list types surrounded by parentheses
  * The `(result_type, error)` paradigm is very common idiom
  ```go
  func foo() (int, err)
  ```
* Can use named return values
  * Initializes returned variables
  * Return using `return` keyword on its own
* Can return addresses of local variables
  * Automatically promoted from local memory (stack) to shared memory (heap)

#### Anonymous
* Functions don't have names if they are:
  * Immediately invoked
  ```go
  func() {
    ...
  }()
  ```
  * Assigned to a variable or passed as an argument to a function
  ```go
  a := func() {
    ...
  }
  a()
  ```

#### Function as types
* Can assign functions to variables or use as arguments and return values in functions
* Type signature is like function signature, with no parameter names
  ```go
  var f func(string, string, int) (int, error)
  ```

#### Methods
* Function that executes in context of a type
* Format
  ```go
  func (g greeter) greet() {
    ...
  }
  ```
* Receiver can be value or pointer
  * Value receiver gets copy of type
  * Pointer receiver gets pointer to type

## Reference
* [Go Functions](https://www.programiz.com/golang/function)
* [Go - Functions](https://www.tutorialspoint.com/go/go_functions.htm)
* [Golang Functions](https://www.golangprograms.com/go-language/functions.html)
* [Methods in Golang](https://www.geeksforgeeks.org/methods-in-golang/)
* [Variadic Functions](https://www.golangprograms.com/go-language/variadic-functions.html)
* [Golangbot Variadic Functions](https://golangbot.com/variadic-functions/)
* [Named Return Parameters in Golang](https://www.geeksforgeeks.org/named-return-parameters-in-golang/)
* [Pass by pointer vs pass by value in Go](https://codewithyury.com/golang-pass-by-pointer-vs-pass-by-value/)
* [Naming Conventions for Golang Functions](https://www.golangprograms.com/naming-conventions-for-golang-functions.html)