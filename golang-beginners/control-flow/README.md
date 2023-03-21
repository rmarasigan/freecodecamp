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

The another interesting thing to know about working with logical operators, is a concept called [**short-circuiting**](https://kuree.gitbooks.io/the-go-programming-language-report/content/22/text.html).

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

Well, what happened is a concept called **short-circuiting**. As soon as one part of an *OR* test returns `true`, then Go doesn't need to execute any more code, it already knows that the *OR* test is going to pass. So that's why it is called short-circuiting, which basically means it is not going to evaluate any other part of the *OR* test. It is going to move on and say, well, the *OR* test passed, and therefore everything works. Since `guess < 1`, there's no reason for Go to evaluate the two tests. Go is going to lazily evaluate the logical tests that we put in here.

The same thing happens for an *AND* (`&&`) test, if one of the parameters returns `false`, then we're going to get a short-circuiting. So if we get into a situation where it is `false`, Go will not even evaluate the test. For example, with the negative five (-5), when Go evaluates this, it sees that the `guess` is not greater than or equal to one. So the *AND* test has to fail because both sides of *AND* test have to return `true` in order to work and so Go is going to exit early. It's not going to execute this code.

```go
func main() {
	number := 50
	guess := -5
	...
	if guess >= 1 && guess <= 100 {
		...
	}
	...
}
```

The way we can do it a little bit more elegantly in Go is by putting in the keyword `else`. If the tests return a boolean `true`, then it will be going to print `The guess must be between 1 and 100!`. Otherwise, it will execute the `else` block.

```go
func main() {
	number := 50
	guess := 30

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {

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

Output:
```
Too low
false true true
```

Related to that is the situation where we have multiple tests that we want to chain togehter. In this case, we've actually split out the validation logic. So we want to check if it is less than one, we want to print out one message, if it is greater than one hundred, we want to print out another message. Otherwise, we've got a valid value. To do that, we're using the **`else if**` clause. Basically, what this is doing is it is taking an `if` test, and then it is chaining on another `if` test if this fails. So what Go is gonna do is the first logical test that passes is going to be executed, otherwise, it is going to drop down to the `else` clause. With the `guess` of 30, we would expect our `else` clause to fire up.

```go
func main() {
	number := 50
	guess := 30

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {

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

Output:
```
Too low
false true true
```

Another thing to keep in mind when you are working with equality operators. If we take a look at this example, we've got a floating point number of 0.1. And we're going to check to see if that number is equal to the squared number, and then we're going to take the square root of it. If we run this, it should evaluate to true.

```go
func main() {
	myNum := 0.1

	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
```

Output:
```
These are the same
```

However, if we add in a couple of decimal places here and run this, now it says that the results are different. Well, the fact is that if we square a number and take the square root of it, it is the same number. The thing is, Go is working with floating point numbers and floating point numbers are approximations of decimal values. They're not exact representations of decimal values. So we have to be very careful with them. When we're doing comparison operations with decimal values, this is generally not a good idea. The better approach is to generate some kind of an error value and then check to see if that error value is less than a certain threshold.

```go
func main() {
	myNum := 0.123

	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
```

Output:
```
These are different
```

In this case, since we're taking a float point number and comparing it to a floating point number, what we can do is divide it into two numbers and subtract 1. Now, what we have is a value that's close to zero. If we take the absolute value of that, and then check that to see if it is less than, for example, 0.001. What this is doing is, if they're within a 10th of a percentage of each other, then we're going to consider them to be the same.

```go
func main() {
	myNum := 0.123	

	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2), -1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
```

Output:
```
These are the same
```

Now, this isn't a perfect solution, because we could get two numbers that aren't truly the same and get them to pass this test. So the key to that is tuning the error parameter (`0.001`), making sure that it is sufficiently small to catch those cases, but sufficiently large, so that the errors introduced with floating point operations don't affect the results.

## `switch` statements
A **`switch`** statement allows a variable to be tested for equality against a list of values. Each value is called a ***case***, and the variable being switched on is checked for each **switch case**. It is a kind of special purpose `if` statement. So where `if` statement has a logical test and runs a block of code if that test is true, and then you can add an `else if` to run a second test and another `else if` and as many `else if`'s as you want.

In this example, we're going to start with the value of `2` and that is called a *tag* when we're working with `switch` statements. So this is what we're going to compare everything against, then we're going to have a series of cases. What's going to happen is the value in each cases is going to be compared to the tag. If that is true, then the statements after the `case` are going to be executed. If not, then they're not going to execute.

Just like with `if` statements, the first case that passes is going to be the one that executes. We also have the **`default`** case and that is going to execute if none of our other test cases passed. It is a very simple way to compare one variable to multiple possible values for that variable.

Example:
```go
func main() {
	switch 2 {
	case 1:
		fmt.Println("one")
	
	case 2:
		fmt.Println("two")
	
	default:
		fmt.Println("not one or two")
	}
}
```

Output:
```
two
```

If you want to compare against a range, then you'd have to do something like *falling through* in order to compare multiple cases together. We don't have falling through as a default behavior in Go, but what we do have to make up for that is the ability to have multiple tests in a single case.

```go
func main() {
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	
	case 2, 4, 6:
		fmt.Println("two, four or six")
	
	default:
		fmt.Println("another number")
	}
}
```

Output:
```
one, five or ten
```

One thing that we have to be aware of is the test cases do have to be unique. There would be a syntax error because we can't have overlapping cases when we're using the syntax.

```go
func main() {
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	
	// duplicate case 5 in switch
	case 2, 4, 5:
		fmt.Println("two, four or five")
	
	default:
		fmt.Println("another number")
	}
}
```

Just like with `if` statements, we don't have to have a simple tag in our switch statement, we can use an initializer.
```go
func main() {
	switch i := 2 + 3; i  {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	
	case 2, 4, 6:
		fmt.Println("two, four or five")
	
	default:
		fmt.Println("another number")
	}
}
``` 

Output:
```
one, five or ten
```

Another syntax, called *tagless* syntax is arguably more powerful than the *tag* syntax. Although it is a little bit more verbose. We've got a `switch` statement that is standing alone and immediately opening a curly brace. Notice that we don't have a tag here. However, in the `case` statements now, we've got full comparison operations. We can use the comparison operators or the logical operators. In this `cases` are standing in for the logical tests that we have in our `if` statements.

```go
func main() {
	var i := 10

	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
	
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	
	default:
		fmt.Println("greater than twenty")
	}
}
```

Output:
```
less than or equal to ten
```

If you notice, the first case and the second case overlap because 10 is less than or equal to 10. It is also less than or equal to 20. Unlike the *tag* syntax, we've got multiple test cases and they cannot overlap. When we're using the *tagless* syntax, they are allowed to overlap. If they do, then the first case that evaluates to true is going to execute.

What happens if you do want your case to fallthrough? Let's just say that we do want to print out that we're less than or equal to 10, and that we're less than or equal to 20. We know that if we pass the first case, we're going to pass the second case too. In that case, we can use the keyword **`fallthrough`**.

```go
func main() {
	var i := 10

	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	
	default:
		fmt.Println("greater than twenty")
	}
}
```

Output:
```
less than or equal to ten
less than or equal to twenty
```

One thing to keep in mind, `fallthrough` is logicless. So when we say `fallthrough` to the next case, it will do that. In the example, we know that `i` fails the second case but it still executes the second case because `fallthrough` means that you are taking that responsibility with the control flow and you intentionally want the statements in the next case to execute.

```go
func main() {
	var i := 10

	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	
	case i >= 20:
		fmt.Println("less than or equal to twenty")
	
	default:
		fmt.Println("greater than twenty")
	}
}
```

Output:
```
less than or equal to ten
less than or equal to twenty
```

### Two types of `switch` statements

#### Expression Switch
In **expression switch**, a case contains expressions, which is compared against the value of the switch expression.

**Syntax:**
```go
switch (boolean-expression or integral type) {
case boolean-expression or integral type:
	statement(s)

case boolean-expression or integral type:
	statements(s)

default: // Optional
	statement(s)
}
```

Example:
```go
func main() {
	var (
		marks int = 90
		grade string = "B"
	)
	
	switch marks {
	case 90:
		grade = "A"
	
	case 80:
		grade = "B"
	
	case 50,60,70 :
		grade = "C"
	
	default:
		grade = "D"  
  }
	
	switch {
	
	case grade == "A" :
		fmt.Printf("Excellent!\n" )
	
	case grade == "B", grade == "C" :
  	fmt.Printf("Well done\n" )
	
	case grade == "D" :
		fmt.Printf("You passed\n" )
	
	case grade == "F":
		fmt.Printf("Better try again\n" )
	
	default:
		fmt.Printf("Invalid grade\n" );
  }
	
	fmt.Printf("Your grade is  %s\n", grade );
}
```

Output:
```
Excellent!
Your grade is  A
```

#### Type Switch
In **type switch**, a case contain type which is compared against the type of a specially annotated switch expression.

**Syntax:**
```go
switch x.(type) {
case type:
	statement(s)

case type:
	statement(s)

default: // Optional
	statement(s)
}
```

What this is going to do is tell Go to pull the actual underlying type of that interface and use that `interface` and use that for whatever we're doing next. We just don't use that for type switching, but this is a common use case for it.

```go
func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x: %T", i)
	
	case int:
		fmt.Printf("x is int")
	
	case float64:
		fmt.Printf("x is float64")
	
	case func(int) float64:
		fmt.Printf("x is func(int)")
	
	case bool, string:
		fmt.Printf("x is bool or string")
	
	default:
		fmt.Printf("don't know the type")
	}
}
```

Output:
```
type of x:<nil>
```

## Summary

#### If statements
* Initializer
* Comparison operators
* Logical operators
* Short circuiting
* If-else statements
* If-else if statements
* Equality and floats

#### Switch statements
* Switching on a tag
* Cases with multiple tests
* Initializers
* Switches with no tags
* Fallthrough
* Type switches

## Reference
* [Go switch](https://zetcode.com/golang/switch/)
* [If statements](https://golangr.com/if/)
* [Switch Statement](https://golangr.com/switch/)
* [Operators: complete list](https://yourbasic.org/golang/operators/)
* [Go - The Switch Statement](https://www.tutorialspoint.com/go/go_switch_statement.htm)
* [Control statements and functions](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/02.3.html)