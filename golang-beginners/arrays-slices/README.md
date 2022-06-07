# Array and Slices

### Agenda
* Arrays
   * Creation
   * Built-in functions
   * Working with arrays
* Slices
   * Creation
   * Built-in functions
   * Working with slices

## Arrays

Without arrays, we're going to end up with an application like this. Now this works sort of but we got a lot of problems here, because our application needs to know exactly how many grades we have to work with, at the time that we're designing the application.
```go
func main() {
   grade1 := 97
   grade2 := 85
   grade3 := 93
   fmt.Prinf("Grades: %v, %v, %v", grade1, grade2, grade3) // Output: Grades: 97, 85, 93
}
```

Now, the way we declare an array is we're going to start with the size of the array. So we're going to use `[]` brackets and then we're going to have an index that's going to be the number of elements that an array can hold, and then the type of data that the array is going to be designed to store. An array can only store one type of data. In this case, we're declaring an array of integers that can hold up to three elements. If we wanted to hold a different type, you have to specify at the time that you're declaring the array, what type of data you're going to store. Then we can use the initializer syntax to put in the values for our array. We see that we have all of the grades printed out together in this collection called an **array**.
```go
func main() {
   grades := [3]int{97, 85, 93}
   fmt.Prinf("Grades: %v", grades) // Output: Grades: [97, 85, 93]
}
```

Another advantage that we have with working with arrays is the way that they're laid out in memory. If you declare three different variables and specify their values, it's impossible to know how they're going to be laid out by the go runtime with arrays. However, we know by the design of the language that these elements are contiguous in memory, which means accessing the various elements of the array is very, very fast. So by collecting our data together in arrays, not only is it easier to work with, but it also makes our applications generally a little bit faster.

One problem that we have in the example, we're actually declaring the size of the array twice because we have the `[3]` syntax where we're saying that we're creating a 3 element integer array, but then we're adding three elements to it and that's not really required. If you're going to be initalizing an array literal like in the example, then you can actually replace the size with `...`. Basically, what that says is, create an array that's just large enough to hold the data that I'm going to pass to you in the literal syntax.
```go
func main() {
   grades := [...]int{97, 85, 93}
   fmt.Prinf("Grades: %v", grades) // Output: Grades: [97, 85, 93]
}
```

We can also declare an array that has a certain size, but has its values zeroed out. Then we see that we have an array that's empty. So we have declared a three element array that can hold strings but obviously there's no elements in there right now.
```go
func main() {
   var students [3]string
   fmt.Prinf("Students: %v", students) // Output: Students: [   ]
}
```

In order to specify a value in the array, we're going to call upon the array and then we're going to tell it whihc index we want to work with within the array. In this case, we're working with the 0 index of the students array. Now we see that initially we have an array of students that's full of empty strings. In the second instance, we've actually specified that first element. You may be wondering why we're starting with the 0 value and the reason is related to how arrays are made up of contiguous blocks of memory. So when we talk about `students` as the name of the array, what Go is going to do is it's going to have a pointer or it's going to remember the location of the beginning of that array. Then the index that we pass in this case is 0, is going to tell it how many strings to walk forward. So it knows that when it has a string, a string has a certain lenght and it's going to walk that many strings. When we pass 0, it's going to be the head of the students array moved forward 0 strings elements. That's going to be the first element of our array. It doesn't matter what order we work with these, if we flip these around, then we find that they do filp around in the array, we can assign them in any order that we want.
```go
func main() {
   var students [3]string
   fmt.Prinf("Students: %v\n", students) // Output: Students: [   ]

   students[0] = "Lisa"
   fmt.Prinf("Students: %v", students)   // Output: Students: [Lisa  ]

   students[1] = "Ahmed"
   students[2] = "Arnold"
   fmt.Prinf("Students: %v", students)   // Output: Students: [Lisa, Ahmed, Arnold]
}
```

If we want to get a specific element in the array, then we can use the `[]` bracket and derefence the element from the array. We can use this `[]` syntax in order to assign values to the array, as well as to pull out the values that have been assigned.
```go
func main() {
   var students [3]string

   students[0] = "Lisa"
   students[2] = "Ahmed"
   students[1] = "Arnold"
   fmt.Prinf("Students #1: %v", students[1]) // Output: Students #1: Arnold
}
```

Another thing that we can do is we can determine how big the array is. The way that we can do that is using the built in **`len`** function. If we run this, we see that we get the number of students that is equal to 3 and that's going to print out the size of the array.
```go
func main() {
   var students [3]string

   students[0] = "Lisa"
   students[2] = "Ahmed"
   students[1] = "Arnold"
   fmt.Prinf("Students #1: %v\n", students[1])          // Output: Students #1: Arnold
   fmt.Prinf("Number of Students: %v\n", len(students)) // Output: Number of Students: 3
}
```

One thing that's important to remember is that an array can be made up of any type, it just always has to be the same type for a given array. For this example, we've got an array of arrays. Let's just say that we're working with some linear algebra and we need the identity matrix, which is a concept that's used pretty often in linear algegra. This array here stores a 3x3 identity matrix. The first row is going to hold the values [1, 0, 0], second row is going to hold [0, 1, 0], and the third row is going to hold [0, 0, 1].
```go
func main() {
   var identityMatrix [3][3] int = [3][3]int { [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1} }
   fmt.Println(identityMatrix) // Output: [[1, 0, 0]. [0, 1, 0], [0, 0, 1]]
}
```

Another way to look at the `identityMatrix`, and maybe a little bit easier to see is by going to declare the array of arrays and then we're going to initialize each one of those rows individually. So this reads a little bit cleaner and might be a little bit easier for you to understand what's going on.
```go
func main() {
   var identityMatrix [3][ 3] int
   identityMatrix[0] = [3]int{1, 0, 0}
   identityMatrix[1] = [3]int{0, 1, 0}
   identityMatrix[2] = [3]int{0, 0, 1}
   fmt.Println(identityMatrix) // Output: [[1, 0, 0]. [0, 1, 0], [0, 0, 1]]
}
```

Arrays are actually considered values. When you copy an array, you're actually creating a literal copy. It's not pointing to the same underlying data, it is pointing to a different set of data which means it's got to reassign that entire length of the array. When you're working with these, you have to be a little bit careful because copying arrays, especially we get info functions, if you're passing arrays into a function, Go is going to copy that entire array over. If you're dealing with a three element array, that's not a big deal, if you've got a million elements in your array that could slow your program down a little bit. So what do you when you don't want to have this behavior?
```go
func main() {
   a := [...]int{1, 2, 3}
   b := a
   b[1] = 5
   fmt.Println(a) // Output: [1, 2, 3]
   fmt.Println(b) // Output: [1, 5, 3]
}
```

If we do the address of operation, which is `&` character here, then we're saying is `b` is going to point to the same data that it has. If you run this, `a` and `b` are pointing to the same data. So `a` is the array itself, and `b` is pointing to `a`.
```go
func main() {
   a := [...]int{1, 2, 3}
   b := &a
   b[1] = 5
   fmt.Println(a) // Output: [1, 5, 3]
   fmt.Println(b) // Output: &[1, 5, 3]
}
```

Arrays are very powerful and there's a lot of use cases where you can use arrays very efficiently. However, the fact that they have a fixed size that has to be known at compile time definitely limits their usefulness. In Go, the most common use case for using arrays is to back something called a **slice**.

## Slice
A slice is initialzied as a literal by just using the square brackets, the type of data we want to store and then in the curly braces we can pass in the initialized data. It looks exactly like an array and as a matter of fact, everything we can do with an array we can do with a slice as well, with one or two exceptions.
```go
func main() {
   a := []int{1, 2, 3}
   fmt.Println(a)                      // Output: [1, 2, 3]
   fmt.Printf("Length: %v\n", len(a))  // Output: Length: 3
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 3
}
```

Slices are naturally what are called **reference types**. So they refer to the same underlying data. If we run this, we see that `a` and `b` are actually pointing to the same underlying array and when we change the value in `b`, we get a change in the value in `a`. This is one thing that you're going to have to keep in mind when you're working with slices. If you got multiple slices pointing to the same underlying data, you have to keep in mind, if one of those slices changes the underlying data, it could have an impact somewhere else in your application.
```go
func main() {
   a := [...]int{1, 2, 3}
   b := a
   b[1] = 5
   fmt.Println(a)                      // Output: [1, 5, 3]
   fmt.Println(b)                      // Output: [1, 5, 3]
   fmt.Printf("Length: %v\n", len(a))  // Output: Length: 3
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 3
}
```

You can see below that we're creating a slice that has the values one through ten and then which is`a`. Here's a detailed explanation of each slices.
* Slice `b` is using a bracket with the colon in between (`[:]`). What that is going to do is, it is basically going to crete a slice of all the elements of what it is referring to. So it's going to create a slice of all the elements of `a`.
* Slice `c` is going to start with the 3rd element of the parent and copy all the values after that. That is going to start with the element with index three, which is, of course the fourth element and every element after that. So it is going to copy four through 10 into the slice `c`.
* Slice `d` is going to copy everything up to the sixth element and that is the literal sixth element not element number seven. That is actually the sixt element which is going to have the index five.
* Slice `e` is going to copy the fourth element through the sixth element into our new slice.
```go
func main() {
   a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9 ,10}
   b := a[:]   // Slice of all elements
   c := a[3:]  // Slice from 4th element to end
   d := a[:6]  // Slice of first 6 elements
   e := a[3:6] // Slice the 4th, 5th, and 6th elements
   fmt.Println(a) // Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
   fmt.Println(b) // Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
   fmt.Println(c) // Output: [4, 5, 6, 7, 8, 9, 10]
   fmt.Println(d) // Output: [1, 2, 3, 4, 5, 6]
   fmt.Println(e) // Output: [4, 5, 6]
}
```

That can be a little bit confusing because the first number has a slightly different meaning than the second number. Basically, the first number is inclusive and the second number is exclusive (`[inclusive:exclusive]`). One thing to keep in mind, remember that all of these operations point at the same underlying data. So if you take element five on slice `a` and change that value, notice all of them changed their value because they're all pointing to the same underlying array. So each one of these operations includes the fifth index in their results and each one of those gets updated to the value of 42.
```go
func main() {
   a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9 ,10}
   b := a[:]   // Slice of all elements
   c := a[3:]  // Slice from 4th element to end
   d := a[:6]  // Slice of first 6 elements
   e := a[3:6] // Slice the 4th, 5th, and 6th elements
   a[5] = 42
   fmt.Println(a) // Output: [1, 2, 3, 4, 5, 42, 7, 8, 9, 10]
   fmt.Println(b) // Output: [1, 2, 3, 4, 5, 42, 7, 8, 9, 10]
   fmt.Println(c) // Output: [4, 5, 42, 7, 8, 9, 10]
   fmt.Println(d) // Output: [1, 2, 3, 4, 5, 42]
   fmt.Println(e) // Output: [4, 5, 42]
}
```

Another thing to know about these slicing operations is they can work with slices like we're doing but they can also work with arrays. If you put the three dots between the brackets (`[...]`), it is actually going to turn `a` into an array. We will also get the same result and that is because slicing operations can have as their source, an array or a slice. Whatever type of data you are working with, as long as it is done of those two, you can use these slicing operations.
```go
func main() {
   a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9 ,10}
   b := a[:]   // Slice of all elements
   c := a[3:]  // Slice from 4th element to end
   d := a[:6]  // Slice of first 6 elements
   e := a[3:6] // Slice the 4th, 5th, and 6th elements
   a[5] = 42
   fmt.Println(a) // Output: [1, 2, 3, 4, 5, 42, 7, 8, 9, 10]
   fmt.Println(b) // Output: [1, 2, 3, 4, 5, 42, 7, 8, 9, 10]
   fmt.Println(c) // Output: [4, 5, 42, 7, 8, 9, 10]
   fmt.Println(d) // Output: [1, 2, 3, 4, 5, 42]
   fmt.Println(e) // Output: [4, 5, 42]
}
```

Now, the last way that we have available to us to create a slice is using what is called the `make` function and that is a built-in function that we have to work with. This takes two or three arguments. So let us start with two arguments. In here we are going to make a slices of intergers. The second argument is going to be the length of the slice.
```go
func main() {
   a := make([]int, 3)
   fmt.Println(a)                      // Output: [0, 0, 0]
   fmt.Printf("Length: %v\n", len(a))  // Output: Length: 3
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 3
}
```

When I create a slice, everything gets set to the zero value which is what we always expect in Go to happen. Everytime we initialize a variable, we expect it to be initialzied to zero values and that is true for slices just like it is true for primitives. Now, we can also pass the third argument to the make function and that is going to set the capacity. Keep in mind, the slice has an underlying array and they don't have to be equivalent. If we run this, we see that we've created a slice of length three, it got three elements in it, but the underlying array has 100 elements in int.
```go
func main() {
   a := make([]int, 3, 100)
   fmt.Println(a)                      // Output: [0, 0, 0]
   fmt.Printf("Length: %v\n", len(a))  // Output: Length: 3
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 100
}
```

Why would we do that? Well, the reason is because unlike arrays, slices don't have to have a fixed size over their entire life. We can actually add elements and remove elements from them. If we want to add an element to this slice, we can use the built-in `append` function. It takes two or more arugments. The first is going to be the source slice that we're going to be working with and the second one the is an element that we want to add in it. On the second operation, we notice that the capacity is two. What happened here is when we initialize a slice to the value `a`, Go assigned a memory location for the slice and since it didn't have to store anything, it basically created an underlying array of zero element for us. As soon as we added an element, it couldn't fit in a zero element array. So it had to assign an array for us. What Go does is it copies all of the existing elements to a new array that got a larger size. When we re-assigned, it actually did create a new array, this one has a capacity of two, and then put the value `1` into that array. However, as things get very large, these copy operations become very expensive and that is why we have that three parameter in `make` function. That way, if we know the capacity is going to be somewhere around 100 elements you can go ahead and start there. That way as you are appending elements and building up the slice, you are not constantly copying the underlying array around. 
```go
func main() {
   a := []int{}
   fmt.Println(a)                      // Output: []
   fmt.Printf("Length: %V\n", len(a))  // Output: Length: 0
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 0
   
   a = append(a, 1)
   fmt.Println(a)                      // Output: [1]
   fmt.Printf("Length: %v\n", len(a))  // Output: Length: 1
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 2
}
```

If you remember the append function can take two or more arguments, the reason for that is called **variadic function**. Everything after the first argument is going to be interpreted as a value to append to the slice passed in the first argument. Now the capacity is 8. Generally, once it fills up the underlying array with a slice, when you add the next element, it is going to create a new array and it is actually going to doubel the size from the previous array. If we start with an empty slice, we see that the array would initially be of size zero then we'll go to 2 4 8 16 32 and 64 elements. That is something to be aware of if you are just over one of those powers of two,  you can actually end up with a lot of memory consumed that you're never going to be using.
```go
func main() {
   a := []int{}
   fmt.Println(a)                      // Output: []
   fmt.Printf("Length: %V\n", len(a))  // Output: Length: 0
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 0
   
   a = append(a, 1)
   fmt.Println(a)                      // Output: [1]
   fmt.Printf("Length: %V\n", len(a))  // Output: Length: 1
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 2

   a = append(a, 2, 3, 4, 5)
   fmt.Println(a)                      // Output: [1, 2, 3, 4, 5]
   fmt.Printf("Length: %V\n", len(a))  // Output: Length: 5
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 8
```

One common situation you're going to run into is if you have a slice of elements and another slice of elements, and you want to concatenate them together, so you want to have another slice created that has all the elements of the first slice and all of the elements of the second slice. You can add a `...` dots after the clie, it is actually going to spread that slice out into individual arguments. It is basically going to take the slice and decompose it into something like this: `2, 3, 4, 5`.
```go
func main() {
   a := []int{}
   fmt.Println(a)                      // Output: []
   fmt.Printf("Length: %V\n", len(a))  // Output: Length: 0
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 0
   
   a = append(a, 1)
   fmt.Println(a)                      // Output: [1]
   fmt.Printf("Length: %V\n", len(a))  // Output: Length: 1
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 2

   a = append(a, []int{2, 3, 4, 5}...)
   fmt.Println(a) // Output: [1, 2, 3, 4, 5]
   fmt.Printf("Length: %V\n", len(a)) // Output: Length: 5
   fmt.Printf("Capacity: %v\n" cap(a)) // Output: Capacity: 8
```

Some other common operations that might do with slices are **stack operations**. Let's say that we're treating our slices a stack and we want to be able to push elements onto the stack and pop elements off of the stack. With the `append` function it is going to allow us to push element onto the stack. For popping elements, we can do what is called a **shift operations** which means we want to remove the first element from the slice. So slice `b is going to create a new slice that starts at index one which has the value two and takes everything else from that. If you want to trim an element off the end then you are going to have to use a different syntax. So we want all the initial elements and we are going to start with a colon between the brackets (`[:]`) and then use the `len` operation to figure out the length of the slice. Remember that it is going to return a number that is too large. So slice `c` will return a value of `[1, 2, 3, 4]`.
```go
func main() {
   a := []int{1, 2, 3, 4, 5}
   b := a[1:]
   fmt.Println(b)    // Output: [2, 3, 4, 5]
   c := a[:len(a)-1] // Output: [1, 2, 3, 4]
```