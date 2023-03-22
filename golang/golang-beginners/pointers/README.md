# Pointers

### Agenda
* Creating pointers
* Dereferencing pointers
* The new function
* Working with `nil`
* Types with internal pointers

## Defining and using Pointers
A pointer is a variable that stores a memory address. In simple terms, it stores the value in your computer memory. Though later, we can access it using a pointer address. We use pointers when we want to pass a variable by reference to some function and also for pointing out arrays, structs, interfaces, functions, etc.

When you use a pointer to a variable, there are a couple of different syntax elements that you need to understand. The first one is the use of the **ampersand** (`&`). If you place an ampersand in front of a variable name, you are stating that you want to get the *address* or a pointer to that variable. The second syntax element is the use of the **asterisk** (`*`) or *dereferencing* operator. When you declare a pointer variable, you follow the variable name with the type of the variable that the pointer points to, prefixed with an *.

As you can see in our example, we're declaring a variable `a` and assigning it the value of `42`, creating a variable of `b`, and assigning it to the value of `a`.    

```go
func main() {
  a := 42
  b := a
  fmt.Println(a, b)
}
```

The value 42 prints out two times. Now, since `a` and `b` are value types, the `b := a` line, Go is going to copy whatever data is stored in `a` and assign it to `b`. So we're not pointing to the same memory.

Output:
```
42 42
```

We can prove that by changing the value of `a` into, for example, 27. If we try it again, we see that the value of `a` changes to 27, but the value of `b` stays at 42.

```go
func main() {
  a := 42
  b := a
  fmt.Println(a, b)
  a = 27
  fmt.Println(a, b)
}
```

Output:
```
42 42
27 42
```

Now, we can change this behavior a little bit by having `b` point to `a` using called a **pointer.** The way we declare it as a pointer is by preceding the type `int` that we're pointing to with an *asterisk* (`*`).

If we want to `b` to point to `a`, then we're going to use a special operator called the *address* of operator. It is now saying that `b` is a pointer to an integer and we want to point it to `a`. 

```go
func main() {
  var a int = 42
  var b *int = &a
  
  fmt.Println(a, b)
}
```

As we can see, `a` is still holding the value 42, like we expect, but `b` is holding a strange data. So what is that? Well, that is actually the numerical representation for the memory address that is holding the location of `a`. At this location in memory, we actually have assigned the integer 42. So `b` isn't holding the value 42, it is holding the memory location that is holding the `a` data.

Output:
```
42 0xc0000b2000
```

We can prove that by using the address of operator. In fact, we do see the values are exactly the same. So the address of `a` in memory is `0xc0000b2000`, and `b` is holding the same exact value.

```go
func main() {
  var a int = 42
  var b *int = &a
  
  fmt.Println(&a, b)
}
```

Output:
```
0xc0000b2000 0xc0000b2000
```

As a matter of fact, we can even go the other way because while the address of operator is going to give us the address of a variable in memory, we can use a different operator to figure out what value is actually being stored at a memory location. That is called the *deferencing* operator.

## Dereferencing Pointers
If you want to print out the value of the variable being pointed at from the `b` variable, you need to *dereference* that variable. The following code uses the `*` operator to dereference the `b` variable and retrieve its value.

Now you notice we're using the same asterisk (`*`) but it has a little bit different meaning. The `*int`, this asterisk before the type is declaring a pointer to data of that type. So this is a pointer to an integer. If we put the asterisk in front of a pointer, then that is going to dereference, which basically means it is going to ask the Go runtime to drill through that pointer, find the memory location that the pointer is pointing to, and then pull the value back out.

```go
func main() {
  var a int = 42
  var b *int = &a
  
  fmt.Println(a, *b)
}
```

We see that we get the value 42 printed out both times once again. Now what is the point of all of this? Since `b` is pointing at the same value of `a`, both of these variables are actually tied together.

Output:
```
42 42
```

If we change the value of `a` once again, we see that both `a` and dereferencing `b`, both give the value of 27 because they're both pointing at the same underlying data.

```go
func main() {
  var a int = 42
  var b *int = &a
  fmt.Println(a, *b)

  a = 27
  fmt.Println(a, *b)
}
```

Output:
```
42 42
27 27
```

As a matter of fact, we can even dereference the pointer and use that to change the value. Again, we see that both values are changing `a` and the value that `b` is pointing to are both changed because it is in fact the same data.

```go
func main() {
  var a int = 42
  var b *int = &a
  fmt.Println(a, *b)

  a = 27
  fmt.Println(a, *b)

  *b = 14
  fmt.Println(a, *b)
}
```

Output:
```
42 42
27 27
14 14
```

If you come from a background of languages that allow you to work with pointers as variables, then you might be expecting to do something called *pointer arithmetic*. Go has no pointer arithmetic.

> #### [Why is there no pointer arithmetic?](https://go.dev/doc/faq#no_pointer_arithmetic)
> Safety. Without pointer arithmetic it's possible to create a language that can never derive an illegal address that succeeds incorrectly. Compiler and hardware technology have advanced to the point where a loop using array indices can be as efficient as a loop using pointer arithmetic. Also, the lack of pointer arithmetic can simplify the implementation of the garbage collector.

Now, if you need to have something like this in your application, then check on the [`unsafe`](https://pkg.go.dev/unsafe) package. This is going to give you operations that the Go runtime is not going to check for you. So if you need to do pointer arithmetic, and things like that, then the `unsafe` package, which has a very appropriate name, is available for you for those very advanced scenarios.

## Creating Pointer types
In simpler terms, pointer type is a pointer variable initialized before a data type. The major advantage of using pointer type in Go is that it avoids unsafe operations like updating or manipulating any value. Hence, Go is also a strictly typed language.

Here's an example:

```go
type myStruct struct {
  foo int
}

func main() {
  var ms *myStruct
  ms = &myStruct{foo: 42}

  fmt.Println(ms)
}
```

If we make the `ms` variable a pointer to `myStruct` and then use the address of operator in the object initializer (`ms = &myStruct{foo: 42}`), notice that when we print out the value of `ms`, we end up with an ampersand (`&`). It is saying that `ms` is holding the address of an object that has a field value 42 in it.

Output:
```
&{42}
```

This isn't the only way that we have available to us to initialize a variable to a pointer to an object, we can also use the built-in `new` function. Unfortunately, with the `new` function, we can't use the object initialization syntax. We're just going to be able to initialize an empty object.

## `new()` Function
It's a built-in function that allocates memory, but unlike namesakes in some other languages it does not *initialize* the memory, it only *zeros* it. That is, `new(T)` allocates zeroed storage for a new item of type T and returns its address, a value of type *T. In Go terminology, it returns a pointer to a newly allocated zero value of type T.

```go
type myStruct struct {
  foo int
}

func main() {
  var ms *myStruct
  ms = new(myStruct)

  fmt.Println(ms)
}
```

We see that we do get an initialized object but all of its fields are initialized to their zero values, we can't initialize them using the object initialization syntax.

Output:
```
&{0}
```

Since we have mentioned zero values, it is important to understand the zero value for a pointer. Every variable that you declare in Go has an initialization value. So right here, `ms` is holding something. The question is, what is that thing?

We see that we get the special value `nil` out. So a pointer that you don't initialize is going to be uninitialized for you, and it is going to hold this value `nil`.

```go
type myStruct struct {
  foo int
}

func main() {
  var ms *myStruct
  fmt.Println(ms)

  ms = new(myStruct)
  fmt.Println(ms)
}
```

Output:
```
&{0}
```

This is very important to check in your applications because if you are accepting pointers as arguments, it is best practice to see if that pointer is a `nil` pointer. If it is, then you're going to have to handle that differently.

How do we get at this underlying field and work with it? The obvious way that we're going to need to do this is we're going to have to dereference the `ms` pointer to get in that `struct`, and then we can get in that field.

```go
type myStruct struct {
  foo int
}

func main() {
  var ms *myStruct
  (*ms).foo = 42
  fmt.Println((*ms).foo)

  ms = new(myStruct)
  fmt.Println(ms)
}
```

You might be asking why the parenthesis are there. It turns out that the dereferencing operator has a lower precedence than the dot (`.`) operator. So we need the parenthesis to make sure that we're dereferencing the `ms` variable instead of dereferencing `ms.foo`.

If we try and run it, we do set and get the value 42, which is the value of that field.

Output:
```
42
```

In fact, we don't need that syntax at all (`(*ms).foo`). If we remove it and run it, we get the same behavior.

```go
type myStruct struct {
  foo int
}

func main() {
  var ms *myStruct
  ms.foo = 42
  fmt.Println(ms.foo)

  ms = new(myStruct)
  fmt.Println(ms)
}
```

Output:
```
42
```

The pointer `ms` doesn't have a field `foo` on it. The pointer `ms` is pointing to a structure that has a field `foo`. So how is this working? Well, this is just *syntactic sugar*. It is just the compiler helping us out because it understands or not actually trying to access the `foo` field on the pointer. We're implying that we want the underlying object, and Go is going to go ahead and interpret that properly for us and dereference the pointer. So the compiler really sees the statement the same as the "`(*ms).foo`" statement.

## Internal Pointers
As we can see, we're initializing an array and we're initializing another variable, and pointing to the same array as `a`, then we're going to print out both `a` and `b`.

```go
func main() {
  a := [3]int{1, 2, 3}
  b := a
  fmt.Println(a, b)

  a[1] = 42
  fmt.Println(a, b)
}
```

Output:
```
[1 2 3] [1, 2 ,3]
[1 42 3] [1 2 3]
```

The `a` is going to change, but `b` is not because it is a copy of the array that we stored in `a`, and so they update independently. However, if we remove the index and turn it into a slice, the behavior changes.

```go
func main() {
  a := []int{1, 2, 3}
  b := a
  fmt.Println(a, b)

  a[1] = 42
  fmt.Println(a, b)
}
```

Output:
```
[1 2 3] [1, 2 ,3]
[1 42 3] [1 42 3]
```

Now we see that both `a` and `b` are changed together. Well, the slice is copied just like the array, but the effect of the copying is a little bit different. Because in the version with an array, the values of the array are considered intrinsic to the variable. So `a` is holding the values of the array, as well as the size of the array, and that size is held that way we can do bounds checking.

```go
func main() {
  a := [3]int{1, 2, 3}
  b := a
  fmt.Println(a, b)

  a[3] = 42
  fmt.Println(a, b)
}
```

Output:
```
invalid argument: index 3 out of bounds [0:3]
```

However, with slices, a slice is a projection of an underlying array. So the slice doesn't contain the data itself. The slice contains a pointer to the element that the slice is pointing to on the underlying array.

What that means is, when we work with slices, the internal representation of a slice has a pointer to an array.

```go
func main() {
  a := []int{1, 2, 3}
  b := a
  fmt.Println(a, b)

  a[1] = 42
  fmt.Println(a, b)
}
```

When you're sharing slices in your application, you're always going to be pointing at the same underlying data. Now, the other built-in type that has this behavior is a `map`. Because maps have a pointer to the underlying data, they don't contain the underlying data in themselves.

```go
func main() {
  a := map[string]string{"foo": "bar", "baz": "buz"}
  b := a
  fmt.Println(a, b)

  a["foo"] = "qux"
  fmt.Println(a, b)
}
```

Output:
```
map[baz:buz foo:bar] map[baz:buz foo:bar]
map[baz:buz foo:qux] map[baz:buz foo:qux]
```

When you're working with slices and maps in a Go application, you have to be very, very careful to keep in mind at all times who's got access to that underlying data because passing slices and maps around in the application can get you into situations where data is changing in unexpected ways.

If you are working with other data types, specifically primitives, arrays, or structs, then this isn't going to happen because when you copy a struct, it is going to copy the entire structure unless you are using pointers.

## Summary

#### Creating pointes
* Pointers use an asterisk (`*`) as a prefix to type pointed to
  * `*int` - a pointer to an integer
* Use the address of operator (`&`) to get address of variable

#### Dereferencing pointers
* Dereference a pointer by preceding with an asterisk (`*`)
* Complex types (e.g. structs) are automatically dereferenced

#### Create pointers to objects
* Can use the address of operator (`&`) if value type already exists
  ```go
  ms := myStruct{foo: 42}
  p := &ms
  ```
* Use address of operator before intializer
  ```go
  &myStruct{foo: 42}
  ```
* Use the `new` keyword
  * Can't initialize fields at the same time

#### Type with internal pointers
* All assignment operations in Go are copy operations
* Slices and maps contain internal pointers, so copies point to same underlying data

## Reference
* [Pointers](https://go.dev/tour/moretypes/1)
* [Effective Go](https://go.dev/doc/effective_go#allocation_new)
* [Understanding Pointers in Go](https://www.digitalocean.com/community/conceptual-articles/understanding-pointers-in-go)
* [The Fundamentals of Pointers in Go](https://betterprogramming.pub/pointers-in-go-9aa5c0682a)