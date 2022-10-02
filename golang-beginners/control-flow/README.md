# Control Flow

### Agenda
* If statements
   * Operators
   * If – else and if – else if statements

* Switch statements
   * Simple cases
   * Cases with multiple tests
   * Falling through
   * Type switches

## `if` Statements

**`if`** statements start with a condition. A condition may be `(x > 3)`, `(y < 4)`, `(weather = rain)`, etc. What do you need these conditions for? Only if a *condition* is true, code is executed. So here's the simplest `if` statement example. It starts with the **`if`** keyword and then we're going to have an *expression* that generates some kind of boolean result.

```go
package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}
}
```

Output:
```
The test is true
```

Since we are using the literal `true`, then the test is always going to pass and when you run the code, you will see that it printed the `The test is true` as the result. If you change it to a literal false, we see that we don't get any output from the application because the code of inside the curly braces is not executed.

```go
func main() {
	if false {
		fmt.Println("The test is true")
	}
}
```

One thing to keep in mind, if you're coming to go from another language, is you may be expecting this to be a valid syntax. If you try to run the code, you'll get a syntax error because one of the design decisions that was made with Go is that you are never allowed to have a single line block evaluated as the result of an `if` statement. So always have to use curly braces, even if you only have one line to execute.

```go
func main() {
	if false
		fmt.Println("The test is true")
}
```

Output:
```
prog.go:9:34: expected '{', found newline (and 1 more errors)
```

There is actually another style that is very commonly used in the Go language and that is what I would call the **initializer statement**.

```go
package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

   // Initializer statement
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
}
```

Output:
```
20612439
```

Notice we've got this semi-colon (`;`) and then the boolean result (`ok`) listed here. So the first part of the if statement (`pop, ok := statePopulations["Florida"]`) is the *initializer*. So the **initializer** allows us to run a statement and generate some information that is going to set us up to work within the `if` block. In this case, we are generating a boolean result by interrogating the map and then using that as the test to see if the code inside the `if` block should execute or not. We also have access to the `pop` variable that we created on the first part, inside of the `if`, and that variable is going to be scoped to the block.

If you print `pop` variable outside the `if` block, the we will get an error saying `undefined: pop` because the `pop` variable is only defined within the scope of that `if` statement. 

```go
// Initializer statement
if pop, ok := statePopulations["Florida"]; ok {
   fmt.Println(pop)
}

fmt.Prinln(pop)
```

Output:
```
./prog.go:22:14: undefined: pop
```

The next thing about `if` statements are the [**comparison operators**](https://medium.com/golangspec/comparison-operators-in-go-910d9d788ec0) that we have. Right now, what we have is we've got the `number` that we're trying to guess and we've got our `guess` and then we've got a couple of tests. So the first statement is, we're trying to check if the `guess` is *less than* the `number`. We also have the *greater than operater* and the *equality operator*.

```go
package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 30

	if guess < number {
		fmt.Println("Too low")
	}

	if guess > number {
		fmt.Println("Too high")
	}

	if guess == number {
		fmt.Println("You got it!")
	}
}
```

If you run this example, you are going to get the value `Too low` printed out because the first test evaluates to true and the second and the third test evaluate to false. If you change the `guess` value to `70` then the second test evaluates to true and the first and third evaluate to false. Now, there's some other operators that we have available.

```go
number := 50
guess := 50
fmt.Println(number <= guess, number >= guess, number != guess)

number = 50
guess = 30
fmt.Println(number <= guess, number >= guess, number != guess)

number = 50
guess = 70
fmt.Println(number <= guess, number >= guess, number != guess)
```

Output:
```
true true false
false true true
true false true
```

The first test is called the *less than or equal operator*. So it is going to check to see if `number` is less than or equal to guess. The second is *greater than or equal operator*. The thired is the *not equal operator*.  So these are the six different comparison operators that you're typically going to use in your applications. These work with all numeric types. They don't however work with things liks `string` types. So if you are going to work with string types, typically what you are going to work with is the equality operator or the non-equality operator.

We can also combine multiple tests together using what are called [**logical operators**](https://www.includehelp.com/golang/logical-operators.aspx).

```go
if guess < 1 || guess > 100 {
   fmt.Println("The guess must be between 1 and 100!")
}
```

This code here is conducting two tests. It is checking to see if the guess is less than one *or* if the guess is greater than 100. If it is, then we are going to print out a message. We can also put a guard around our actual code, checking to see if the guess is greater than or equal to one *and* less than or equal to 100.

```go
if guess >= 1 && guess <= 100 {

   if guess < number {
      fmt.Println("Too low")
   }

   if guess > number {
      fmt.Println("Too high")
   }

   if guess == number {
      fmt.Println("You got it!")
   }
}
```

So the first test case is going to execute if the guess is out of range. If we enter negative five, for example, then we're going to see that `The guess must be between 1 and 100!`. But if change the value of guess that is within range, let's say 30. Then everything executes like we expect it to.

So in the first test, the logical operator used is called **OR Operator**. This is checking to see if the test on the left is true, or the test on the right is true. Obviously, we can't have the guess less than one and greater than 100 at the same time, but one or the other could be true. If one of those is true, then we're going to have an invalid guess. It makes sense to print out a message to the user.

The other operator that is on the second test is called the **AND operator**. It evaluates to true if both the test on the left and the test on the right evaluate to true. So if we had a guess, for example of 105, then the code `guess >= 1` is gonna evaluate to true because it is greater than or equal to one. The `guess <= 100` is not because 105 is not less than or equal to 100. With *AND* test, both cases have to be true, unlike in *OR* test, were only one of the two has to be true. So the 105 value will print out our error message and it is not going to execute the code within the second test.

The other logical operator that we have is called the **NOT operator**. What that is going to do is it is going to take a boolean and it is going to flip it to the other side.

```go
fmt.Println(true)
fmt.Println(!true)
```

Output:
```
true
false
```

The another interesting thing to know about working with logical operators, is a concept called [**short circuiting**](https://kuree.gitbooks.io/the-go-programming-language-report/content/22/text.html).

```go
func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

func main() {
	number := 50
	guess := 30

	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}

	if guess >= 1 && guess <= 100 {

		if guess < number {
			fmt.Println("Too low")
		}

		if guess > number {
			fmt.Println("Too high")
		}

		if guess == number {
			fmt.Println("You got it!")
		}
	}

	fmt.Println(number <= guess, number >= guess, number != guess)
}
```

In this case, we are ordeirng three tests togehter. What is going to happen is Go is going to evaluate these `guess < 1 || returnTrue()`, it is going to generate a result and it is going to take that result, and it is going to order it against this `guess > 100`. Basically it is going to evaluate all of them and only one of those has to be true in order to generate a validation message.

Output:
```
returning true
The guess must be between 1 and 100!
Too low
false true true
```

What if we change the value of guess into 5? We get the validation message saying `The guess must be between 1 and 100!`. But what happened to the `returning true`?

Output:
```
The guess must be between 1 and 100!
false true true
```

Well, what happened is a concept called short circuiting. As soon as one part of an *OR* test returns `true`, then Go doesn't need to execute any more code, it already knows that the *OR* test is going to pass. So that's why it is called short circuiting, which basically means it is not going to evaluate any other part of the *OR* test. It is going to move on and say, well, the *OR* test passed, and therefore everything works. Since `guess < 1`, there's no reason for Go to evaluate the two tests. Go is going to lazily evaluate the logical tests that we put in here.

## Reference
* [If statements](https://golangr.com/if/)
* [Operators: complete list](https://yourbasic.org/golang/operators/)
* [Control statements and functions](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/02.3.html)