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