# Defer, Panic, and Recover

### Agenda
* Defer
* Panic
* Recover

## `defer`
In a normal Go application, control flows from the top to the bottom of any function that we call. We use the `defer` statement to prevent the execution of a function until all other functions execute. A **`defer`** statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

What we can do if we want to `defer` the execution in one of the statements is proceeded with the `defer` keyword.

```go
func main() {
  fmt.Println("start")

  // defer the execution of Println() function
  defer fmt.Println("middle")
  fmt.Println("end")
}
```

Notice that the `middle` is printed after the `end`. So the way the `defer` keyword works in Go, is that it executes any functions that are passed into it after the function finishes its final statement, but before it returns.

Output:
```
start
end
middle
```

When Go recognizes that the function exits, it looks to see if there are any deferred functions to call. Since we have one and then go ahead and call that, the deferring doesn't move it to the end of the main function. It moves it after the main function but before the main function is returned.

If we add the `defer` keyword in front of all of these statements, then we'll actually see an interesting behavior because the deferred functions are executed in what is called *LIFO (Last In, First Out) order*.

```go
func main() {
  defer fmt.Println("start")
  defer fmt.Println("middle")
  defer fmt.Println("end")
}
```

The last function that gets deferred is going to be the first one that gets called, and this makes sense because often we're going to use the deferred keyword to close out resources. It makes sense that we close resources out in the opposite order that we open them because one resource might actually be dependent on another one.

Output:
```
end
middle
start
```

A deferred function's arguments are evaluated when the defer statement is evaluated. In this example, we're going to get the value "`start`" printed out. The reason for that is when you `defer` a function like this, it actually takes the argument at the time the `defer` is called out not at the time the called function is executed.

```go
func main() {
  a := "start"
  defer fmt.Println(a)
  a = "end"
}
```

Even if the value of `a` is changed before we leave the main function, we are going to actually eagerly evaluate the variable and so the value "`start`" is going to be evaluated and we're not going to pay attention to the "`end`" value as it changes.

Output:
```
start
```

## `panic`
**Panic** is a built-in function that stops the ordinary flow of control, and begins *panicking*. This can be done by the compiler or added manually to the code flow. We use `panic` statements when the application enters a state that it cannot recover from.

When the function F calls panic, the execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

Now, obviously, the answer of `1 / 0` is invalid, we can't do that in a Go application. We're going to see that the runtime itself will generate a panic for us and the runtime error is printed out .

```go
func main() {
  a, b := 1, 0
  ans := a / b
  
  fmt.Println(ans)
}
```

Output:
```
panic: runtime error: integer divide by zero
```

If you're going along and writing your own program, and you get to a situation where your program cannot continue to execute, because of a certain state that gets thrown, then it makes sense for you to panic as well. To do that, we're going to use the built-in `panic` function.

Notice that the error message that is printed out is going to be a string that we passed into the `panic` function and then we get the stack trace out

```go
func main() {
  fmt.Println("start")
  panic("something bad happen")
  fmt.Println("end")
}
```

Output:
```
start
panic: something bad happen

goroutine 1 [running]:
main.main()
	/tmp/sandbox2212500395/prog.go:9 +0x65
```

Another example that is related to the `defer` function. If we run this, we get the "this was deferred" printed out and then the `panic` happens. This is really important because `panic` happens after the deferred statements are executed.

```go
func main() {
  fmt.Println("start")
  defer fmt.Println("this was deferred")
  panic("something bad happen")
  fmt.Println("end")
}
```

The order of execution is we're going to execute our main function, then we're going to execute any deferred statements, then we're going to handle any panics that occur and then we're going to handle the return value.

Output:
```
start
this was deferred
panic: something bad happen

goroutine 1 [running]:
main.main()
	/tmp/sandbox2063884734/prog.go:10 +0xb5
```

So why is this important? The first thing that is important is that the deferred statements that are going to close the resources are going to succeed even if the application panics.

If somewhere up the call stack, you recover from the panic, you don't have to worrby about the resources being left out there and left open. So any deferred calls to close resources are still going to work even if a function panics.

## `recover`
**Recover** is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

Here, we're creating an anonymous function, and an **anonymous function** is simply a function that doesn't have a name. So nothing else can call it.

At the end of the `func(){}()`, the parenthesis are making the function execute. It is an important thing to know about the `defer` statement. The first statement doesn't take a function itself, it actually takes a function call.

Inside the custom function, notice that we're using a `recover` function. It will return `nil` if the application isn't panicking, but if it isn't `nil`, then it is going to return the error that actually is causing the application to panic.

```go
func main() {
  fmt.Println("start")

  // Creating an anonymous function
  defer func() {
    if err := recover(); err != nil {
      log.Println("Error:", err)
    }
  }()

  panic("something bad happened")
  fmt.Println("end")
}
```

We see that the application still executes. So we got the string "`start`" printed out, then we printed out the error using the `log` package but we didn't get this "`end`" function printed out. It looks like the application is still dead.

Output:
```
start
2009/11/10 23:00:00 Error: something bad happened
```

But `recover` does have some important impacts if we have a deeper call stack than what we're dealing with.

Here, we've got a function called *`panicker`*, and will print `"about to panic"`, and then it is going to panic. It will go ahead and recover from that panic using that deferred function.

```go
// It will be the application entry point
func main() {
  fmt.Println("start")
  panicker()
  fmt.Println("end")
}

func panicker() {
  fmt.Println("about to panic")

  defer func() {
    if err := recover(); err != nil {
      log.Println("Error:", err)
    }
  }()

  panic("something bad happened")
  fmt.Println("done panicking")
}
```

As we can see, we get the "`start`" printed out, we see the "`about to panic`", then we panic, "`something bad happened`", and we go into our recover loop. Inside that deferred function, we call the `recover` function. In that `recover`, we're going to log the fact that we have that error, and we're going to log that error message out, but then in the main function, execution continues.

Output:
```
start
about to panic
2009/11/10 23:00:00 Error: something bad happened
end
```

If we're recovering from a panic, the function that recovers, still stops the execution because it is in a state where it can no longer reliably continue to function, and so it makes sense for it to stop doing whatever it was trying to do.

However, functions in the higher up the call stack, those functions that call the function that recovered from the panic, are presumably in a situation where they can continue because the `recover` function said that the application is in a state working to continue to execute.

To determine what that error is, we have to call the recover function, which means that we're saying that we're going to deal with it. So what happens if we get that error and realize that this isn't something that we can deal with?

In that case, what we can do is re-panicking the application. Inside the handler, we're throwing a `panic` again, and then we see that we don't get that "`end`" statement printed out.

```go
func main() {
  fmt.Println("start")
  panicker()
  fmt.Println("end")
}

func panicker() {
  fmt.Println("about to panic")

  defer func() {
    if err := recover(); err != nil {
      log.Println("Error:", err)
      panic(err)
    }
  }()

  panic("something bad happened")
  fmt.Println("done panicking")
}
```

We get the full stack trace of when the `panic` happened and we see that we're inside a `func1()` function, it is the anonymous function. We see that we do get that stack trace printed out.

Output:
```
start
about to panic
2009/11/10 23:00:00 Error: something bad happened
panic: something bad happened [recovered]
	panic: something bad happened

goroutine 1 [running]:
main.panicker.func1()
	/tmp/sandbox170984560/prog.go:20 +0x8d
panic({0x491720, 0x4c09d8})
	/usr/local/go-faketime/src/runtime/panic.go:884 +0x213
main.panicker()
	/tmp/sandbox170984560/prog.go:24 +0x8b
main.main()
	/tmp/sandbox170984560/prog.go:10 +0x5b
```

## Summary

#### `defer`
* Used to delay execution of a statement until function exits
* Useful to group "open" and "close" functions together
  * Be careful in loop
* Run in LIFO (last-in, first-out) order
* Arguments evaluated at time defer is executed, not at time of called function execution

#### `panic`
* Occur when program cannot continue at all
  * Don't use when file can't be opened, unless it is critical
  * Use for unrecoverable events - cannot obtain TCP port for web server
* Function will stop executing
  * Deferred functions will still fire
* If nothing handles panic, program will exit

#### `recover`
* Used to recover from panics
* Only useful in deferred functions
* Current function will not attempt to continue, but higher functions in call stack will

## Reference
* [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
* [Go defer, panic and recover](https://www.programiz.com/golang/defer-panic-recover)
* [Panic, defer and Recover in Golang](https://golangdocs.com/panic-defer-recover-in-golang)
* [Defer, Panic and Recover Control flow Concepts in Go](https://medium.com/analytics-vidhya/defer-panic-and-recover-control-flow-concepts-in-go-c84265a05993)