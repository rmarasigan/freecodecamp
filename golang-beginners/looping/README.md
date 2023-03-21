# Looping

### Agenda
* For statements
  * Simple loops
  * Exiting early
  * Looping through collections

## `for` Loop
The **`for`** statement specifies repeated execution of a block.

### Simple `for` loop
It is similar that we use in other programming languages with C, C++, Java, C#, etc.

**Syntax:**
```go
for initialization; condition; post {
  // statements
}
```

* **Initialization** statement is optional and executes before `for` loop starts. It is always in a simple statement like variable declarations, increment or assignment statements, or function calls.

* **Condition** statement holds a boolean expression, which is evaluated at the starting of each iteration of the loop. If the value of the conditional statements is true, then the loop executes.

* **Post** statement is executed after the body of the `for`-loop. After the `post` statement, the condition statement evaluates again if the value of the conditional statement is false, then the loop ends.

**Example 1.1**
```go
func main() {
  // This loop starts when i = 0
  // Executes till i < 4 is true
  // Post statement is i++
  for i:= 0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

Output:
```
0
1
2
3
4
```

If we want to initialize two values, we can use Go's ability to initialize multiple values at the same time. For the `post` statement, we can simply do the same as we have done in the `initialization` statement.

```go
func main() {
  for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
    fmt.Println(i, j)
  }
}
```

```
0 0
1 2
2 4
3 6
4 8
```

If we use the increment component that we have earlier, and we just want to increment them by one. We're going to get an error because the increment operation is not an expression so we can't use that as part of a statement. It is a statement on its own.

```go
func main() {
  for i, j := 0, 0; i < 5; i, j = i++, j++ {
    fmt.Println(i, j)
  }
}
```

Output:
```
syntax error: unexpected ++, expected { after for clause
```

The `i` variable might seem like it's something special because it is used up in the `for` loop and the fact is, it is a variable just like anything else. So we can do whatever we want with it. What we do here is we're checking to see if `i` is an even number or an odd number.

```go
func main() {
  for i := 0; i < 5; i++ {
    fmt.Println(i)

    if i%2 == 0 {
      i /=2
    } else {
      i = 2*i + 1
    }
  }
}
```

Output:
```
0
1
4
3
```

Now the next thing is, we don't actually need all the three statements. One thing we could do is to initialized the `i` somewhere else in the application. We can run with the first statement being empty.

**Example 1.2**
```go
func main() {
  i:= 0

  for ; i < 5; i++ {
    fmt.Println(i)
  }
}
```

Output:
```
0
1
2
3
4
```

We can't leave the first `for` semi-colon out, because if you leave that out, everything's in the wrong place and it will think that the second statement is going to be the initializer.

```go
func main() {
  i:= 0

  for i < 5; i++ {
    fmt.Println(i)
  }
}
```

Output:
```
syntax error: unexpected {, expected semicolon or newline
syntax error: unexpected newline, expected { after for clause
```

The difference between the Example 1.1 and Example 1.2 format, is in the *Example 1.2* format, `i` is scoped to the main function. But in the *Example 1.1* format, `i` is scoped to the `for` loop.

The next thing that we can do is we also don't need the incrementer value, so we can eliminate that. If we run this, we're going to get an error because what this is gonna do is it is going to generate an infinite loop. So after it runs, all we'd see is the value zero printed out over and over and over again.

```go
func main() {
  i:= 0

  for ; i < 5; {
    fmt.Println(i)
  }
}
```

We can of course, since `i` is a variable just like any other, we can put the incrementer inside the `for` loop, and everything will work just fine.

```go
func main() {
  i:= 0

  for ; i < 5; {
    fmt.Println(i)
    i++
  }
}
```

## `for` loop as `while` loop
A `for` loop can also work as a `while` loop. This loop is executed until the given condition is true. When the value of the given condition is false, the loop ends.

**Syntax:**
```go
for condition {
  // statement
}
```

We can use the double semi-colon syntax (`for ; i < 5;`) but we can also leave them both off.

```go
func main() {
  i:= 0

  for i < 5 {
    fmt.Println(i)
    i++
  }
}
```

Output:
```
0
1
2
3
4
```

## Infinite Loop
A `for` loop is also used as an infinite loop by removing all the three expressions from the for loop. When the user did not write condition statement in the `for` loop it means the condition statement is treu and the loop goes into an infinite loop.

**Syntax:**
```go
for {
  // statement
}
```

We also have a situation with `do while` loops in other languages, where we need the application to run through the loop and undetermined number of times. By undetermined means, we don't actually know when to stop by an obvious logical test, we need some complex logic within the loop in order to determine when to leave.

So we need some way to tell our infinite loop when it is done processing and when to leave and the way we're going to do that is by using the **`break`** keyword.

The way that we can do this is we'll put some kind of logical tests inside the loop. We can say, if `i` is equal to five, then we're going to use that `break` keyword. The `break` keyword is commonly used inside the `for` loops, especially in this kind of infinite `for` loop.

```go
func main() {
  i:= 0

  for {
    fmt.Println(i) 
    i++

    if i == 5 {
      break
    }
  }
}
```

Output:
```
0
1
2
3
4
```

Another thing we can do is to use what's called a **`continue`** statement. In this case, we're looping from zero to nine, because we're going to keep incrementing as long as `i` is less than ten. Then we're checking to see if the value is even. If it is even, then we're going to use this `continue` statement and what this does is it says, "exit this iteration of the loop and start back over".

```go
func main() {
  for i := 0; i < 10; i++ {
    if i%2 == 0 {
      continue
    }

    fmt.Println(i)
  }
}
```

Output:
```
1
3
5
7
9
```

The `continue` statement isn't used very often, but it can be very, very useful if you're looping through a large set of numbers and you need to determine within the loop whether you want to process a record or not.

## Nested Loop
Here is a pretty simple example, we're starting with the variable `i`, initializing it to one and we're going to run as long as `i` is less than or equal to three. Then we're doing the same thing with `j` on the inside. Basically, `i` is going to be one, then it is going to loop through `j` from one to three. Then it is going to be two, it is going to loop through `j` from one to three again. And then, inside the inner loop, we're just multiplying the two numbers together.

```go
func main() {
  for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
      fmt.Println(i * j)
    }
  }
}
```

Output:
```
1
2
3
2
4
6
3
6
9
```

What happens if we want to leave the loop as soon as we get the first value that's greater than three? In here, as soon as `i * j` is greater than three, it breaks out of the loop. The loop that is going to break out of is the closest loop that it can find.

```go
func main() {
  for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
      fmt.Println(i * j)

      if i * j >= 3 {
        break
      }
    }
  }
}
```

Output:
```
1
2
3
2
4
3
```

The question is, how do we break out in the outer loop? Do we have more logic here and check to see if `i * j` is greater than three? The answer is, we do have a concept called a **label** that we can add. A **label** is used in the `break` and `continue` statement where it is optional but it's required in the `goto` statement. The scope of the label is the function where it is declared.

```go
func main() {
Loop:
  for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
      fmt.Println(i * j)

      if i * j >= 3 {
        break Loop
      }
    }
  }
}
```

Once we have a label defined, we can actually add the label after the `break` keyword and it basically describes where we want to break out to. Since the label `Loop` is before the first loop, we're going to break out of the outer `for` loop.

Output:
```
1
2
3
```

## `for ... range` loop
We iterate over the array with the `range` clause. The `range` returns the *index/key* and the *value* in each iteration. If we do not use the *index/key*, we specify the discard (`_`) operator. This syntax is going to work for slices and arrays.

So, what this is going to do is it is going to look at the collection (`s`), and it is going to take each value one at a time. It gives us the key (`k`) and the value (`v`), and then we're going to be able to work with those values independently.

```go
func main() {
  s := []int{1, 2, 3}
  for k, v := range s {
    fmt.Println(k, v)
  }
}
```

Output:
```
0 1
1 2
2 3
```

## Summary

#### For statements
* Simple loops
  * `for initializer; test; incrementer {}`
  * `for test {}`
  * `for {}`
* Exiting early
  * `break`
  * `continue`
  * Labels
* Looping through collections
  * arrays, slices, maps, strings, channels
  * `for k, v := range collection {}`

## Reference
* [Go for loop](https://zetcode.com/golang/forloop/)
* [Labels in Go](https://medium.com/golangspec/labels-in-go-4ffd81932339)
* [Loops in Go Language](https://www.geeksforgeeks.org/loops-in-go-language/)
* [The for-loop in Golang](https://golangdocs.com/for-loop-in-golang)
* [Get set Go - Labels in Go language](https://ravichaganti.com/blog/get-set-go-labels-in-go-language/)