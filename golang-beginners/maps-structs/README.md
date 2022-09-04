# Maps and Structs

### Agenda
* Maps
   * What are they?
   * Creating
   * Manipulation
* Structs
   * What are they?
   * Creating
   * Naming conventions
   * Embedding
   * Tags

## Maps
Map types are reference types, like pointers or slices. A *`nil`* map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don’t do that.
```go
map[key_type]value_type

// Initialization
var x = map[string]int{}
y := make(map[string]int)
```

The *`make`* function allocates and initializes a hash map data structure and returns a map value that points to it.


A **map** is going to take some kind of a *`key`* and it's going to map that over to some kind of a *`value`*. What this provides us is a very flexible data type. There's a a couple of constraints we're going to have to keep in mind. We can use a lot of different types for the `keys` and we can use any type for the `value`. But when we declare a map, that type has to be consistent for every key-value pair within the map.

So the basic constraint on the keys when you're creating a map is they have to be able to be tested for equality. Now, most of the types that we're working with can do that. So `boolean` can be tested for equality, all of the numeric types, strings, pointers, interfaces, structs, arrays, and channels. However, there are some data types that cannot be used for equivalency checking and those are slices, maps and other functions.

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
	fmt.Println(statePopulations)
}
```

Output:
```
map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596]
```

Now, how are we going to manipulate the values in our map? The first thing that we can do is we can pull out one single value from the map by using the square brackets (`[]`) and typing in the value of the key. So the key can be provided either as a variable or as a literal.
```go
fmt.Println(statePopulations["Ohio"])
```

Output:
```
11614373
```

We can also add a value to our map using a very similar syntax.
```go
statePopulations["Georgia"] = 10310371
fmt.Println(statePopulations["Georgia"])
fmt.Println(statePopulations)
```

Output:
```
10310371
map[California:39250017 Florida:20612439 Georgia:10310371 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596]
```

Notice that the ordering is different and this is going to be a very important thing to keep in mind as you're iterating through maps later on. **The return order of a map is not guaranteed**. This is not like a slice, a slice or an array would return those elements in exactly the same order we provided them. In a map, everything is stored in a way that cannot guarantee the order upon returning. Even though we just added one entry, it completely changed how everything was organized within the map, and we get it a little bit different output.

We can also delete entries from the map. We can do that using the built-in *`delete`* function.
```go
statePopulations["Georgia"] = 10310371
delete(statePopulations, "Georgia")
fmt.Println(statePopulations)
```

Output:
```
map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596]
```

So we can add by just calling on the map providing a new key and value. We can delete using the built-in `delete` function and we can interrogate a value from the map using those square brackets (`[]`). The interesting thing about deleting is if I asked for the value of that key again, we get the value `0`. Now, that should cause you some concern, because does `Georgia` have a population of `0`? Or is it missing from our map? Did I just misspell the key?
```go
fmt.Println(statePopulations["Georgia"])
```

Output:
```
0
```

There's another way that we can interrogate our map and that is using what's called the **`var_name, ok`** syntax.
```go
population, ok := statePopulations["Oho"]
fmt.Println(population, ok)

ohioPopulation, ok := statePopulations["Ohio"]
fmt.Println(ohioPopulation, ok)
```

Output:
```
0 false
11614373 true
```

The `ok` is `false` if the key was not found within our map, but if we correct our spelling, we see that we get value `true` out. So if you're in a situation where you're not sure if the key is in the map or not, then you can use this *`var_name, ok`* syntax.

The next thing about maps is we can find out how many elements are in them. You can do that by using the built-in `len` function.
```go
fmt.Println(len(statePopulations))
```

Output:
```
7
```

You get the value of 7. So if you added another `state`, then the `len` function would return 8. The other thing that is interesting to know is that just like slices when you have multiple assignments to a `map`, which is especially important when you're passing maps into functions, the underlying data is [*passed by reference*](https://hackernoon.com/today-i-learned-pass-by-reference-on-interface-parameter-in-golang-35ee8d8a848e). It means manipulating one variable that points to a map are going to have an impact on the other one.

```go
sp := statePopulations
delete(sp, "Ohio")
fmt.Println(sp)
fmt.Println(statePopulations)
```

Output:
```
map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Pennsylvania:12802503 Texas:27862596]
map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Pennsylvania:12802503 Texas:27862596]
```

You see that in the first result, it doesn't have the entry `Ohio` anymore, and then the second print statement, which is the original map, `Ohio` has been removed from there as well. So if you start passing maps around and you start manipulating the data within that map, keep in mind, that you can have side effects on the calling functions or any other place where that map is referred to. Because manipulating the map in one place is going to have an impact on every other place to use.

## Structs
Go's **`struct`** are typed collection of fields. They are useful for grouping data together to form methods. A **`struct`** or `structure` can be compared with the `class` in the Object-Oriented Programming paradigm. A structure is used mainly when you need to define a **`schema`** made of different individual fields (*properties*). Like a **`class`**, we can create an object from this schema (class analogous to the schema).

So what the `struct` type does is it gathers information together that is related to one concept and it does it in a very flexible way. Because we don't have to have any constraints on the types of data that are contained within our struct, we can mix any type of data together and that is the true power of a struct.

```go
package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:     3,
		actorName:  "Jon Pertwee",
		companions: []string{"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}

	fmt.Println(aDoctor)
}
```

Output:
```
{3 Jon Pertwee [Liz Shaw Jo Grant Sarah Jane Smith]}
```

You see down the `main` function, how we can create a struct and it is using a declaration syntax, where we are using named fields. We have a variable called `aDoctor` and we are going to use this literal syntax. So we list the type of the struct `Doctor`, we use the curly braces (`{}`) as we do with every other literal definition. And then we have a list of the `field_names: value`. Notice that when we are setting something equal to a slice, we do actually have to tell Go what kind of collection type we are initializing for that.

Now, if we want to interrogate one value from a `struct`, then we're going to use what is called the **dot syntax**.

```go
fmt.Println(aDoctor.actorName)
```

Output:
```
Jon Pertwee
```

So we can work with the struct as a whole, or we can start drilling down. As a matter of fact, we can drill down through the structure. Now, another way that we can instantiate our struct is using what is called a **positional syntax**.

```go
func main() {
	aDoctor := Doctor{
		3,
		"Jon Pertwee",
		[]string{"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}

	fmt.Println(aDoctor)
}
```

Output:
```
{3 Jon Pertwee [Liz Shaw Jo Grant Sarah Jane Smith]}
```

Then I get exactly the same result as I got the first time. This is a valid Go syntax, I would encourage you not to use it, though. The reason for that is it can become a maintenance problem. For example, let's say that this is our struct (`Doctor`), and this is checked into source control, and everybody is happy with it. But then, let's say we have a change request coming and we are going to add a list of the episodes that each `Doctor` appeared in. So that is going to be another slice of strings and that gets added right before the `companions`.

```go
type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}
```

Output:
```
./prog.go:19:2: too few values in Doctor{…}
```

Well, if you use the positional syntax and try and run it, we've got a problem. Because Go doesn't know how to map the fields that we provided into the initializer. There are three values provided in the initializer and there are four values in the struct. So we have to find every place that we have one of these declared with the positional syntax.

But what if another change comes along, and somebody interchanges the position of `episodes` and `companions`? Now, they've changed the order of the fields and we asked for the `companions`, we get an empty slice because the positional syntax requires the fields to be declared in the correct order.

```go
type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

func main() {
	aDoctor := Doctor{
		3,
		"Jon Pertwee",
		[]string{},
		[]string{"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}

	fmt.Println(aDoctor.companions)
}
```

Output:
```
[]
```

Another thing we need to talk about is naming conventions. You can see that the fields are declared as lowercase. So the rules in this situation, follow the rules for any other variables in Go. If we start with a capital letter, it is going to be exported from the package. If we start with a lowercase letter, it is going to be internal to the package. Generally, we're going to follow the same conventions as we do with any other variable. Uppercase is going to be exported, while lowercase is going to import. You should use the Pascal Casing and Camel Casing. You shouldn't have underscores and things like that in your field names or your struct names.

There is some situation where you might see a program like this.

```go
package main

import (
	"fmt"
)

func main() {
	aDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aDoctor)
}
```

Output:
```
{John Pertwee}
```

So the `aDoctor := struct{ name string }{name: "John Pertwee"}` is called an **anonymous struct**. Instead of setting up a `type Doctor struct {}`, we are condensing all of that into a single declaration. Now, we can't use it anywhere else, because it is anonymous and so it doesn't have an independent name that we can refer to it. But we can certainly use this just fine.

It is very important to remember what each curly brace `{}` is doing. So the first set of curly braces (`{ name string }`) is paired to the struct keyword and it defines the structure of the struct. The second (`{name: "John Pertwee"}`) is the initializer and it is what is going to provide data into the struct.

```go
struct{ name string }{name: "John Pertwee"}
```

Now, when are you going to use this? This is going to be used in relatively few situations you're going to use this in situations where you need to structure some data in a way that you don't have in a formal type but it is normally only going to be short-lived. You can think about if you have a data model that is coming back in a web application, and you need to send a projection or a subset of that data down to the client, you could create an anonymous struct to organize that information.

Unlike `maps`, `structs` are value types.

```go
func main() {
	aDoctor := struct{ name string }{name: "John Pertwee"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"

	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
```

Output:
```
{John Pertwee}
{Tom Baker}
```

As you can see, even though we copied `anotherDoctor` from `aDoctor`, and changed the value, the values remain independent. So `aDoctor` still has the name `John Pertwee`, `anotherDoctor` has the name `Tom Baker`. So unlike maps, and slices, these are referring to independent datasets. When you pass a struct around in your application, you're passing copies of the same data around. If you've got structs, that are very very large, keep in mind, that you are creating a copy every time you are manipulating it.

Just like with arrays, if we do want to point to the same underlying data, we can use that **address of operator (`&`)**.

```go
func main() {
	aDoctor := struct{ name string }{name: "John Pertwee"}
	anotherDoctor := &aDoctor
	anotherDoctor.name = "Tom Baker"

	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
```

Output:
```
{Tom Baker}
&{Tom Baker}
```

Now, we have in fact, both variables pointing to the same underlying data. `aDoctor` is the struct itself. `anotherDoctor` is a pointer to the struct. So when we manipulate its name field, we are manipulating the `name` field of the `aDoctor` struct.

The next thing we need to discuss is a concept called **embedding**. `Go` uses a model that is similar to inheritance called **composition**. So inheritance is trying to establish a relationship. Go supports composition through what is called *embedding*. We can see that the `Animal` and `Bird` are independent structs, there's no relationship between them.

```go
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	SpeedKPH float32
	CanFly   bool
}
```

We can say that a `Bird` has animal-like characteristics by embedding an `Animal` struct. In here, we are just saying, embed the `Animal` struct right into the `Bird` struct. How can we use this? Well, we create an instance of a `Bird` and initialize the variables. 

```go
package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Australia"
	bird.SpeedKPH = 48
	bird.CanFly = false

	fmt.Println(bird)
	fmt.Println(bird.Name)
}
```

Output:
```
{{Emu Australia} 48 false}
Emu
```

It kinda looks like `Bird` is inheriting the properties of `Animal` but what is happening here is there's some syntactic sugar. Go is handling the delegation of the request for the `Name` field automatically to the embedded animal type for you. So if you try and pass these around and look to see if the `Bird` is a type of `Animal`, it is not. Bird is still an independent struct that has no relationship to an `Animal` other than the fact that it embeds it. It is not a traditional inheritance relationship where a `Bird` is an `Animal`. Instead, it is a composition relationship, which is answering the question "has a". So a bird has an animal, or animal-like  characteristics is how we would say it in this example.

However, if we're going to use the literal syntax, we do have to know a little bit about the `Birds` internal structure.

```go
func main() {
	bird := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(bird)
	fmt.Println(bird.Name)
}
```

So this is declaring the same object. Notice what I have to do here, I have to explicitly talk about the internal `Animal` struct. When we are working with the literal syntax, we have to be aware that we are using embedding but if we just declare the object and manipulate it from the outside, we don't have to be aware of it at all.

Now, when should we use embedding? Well, generally, if you're talking about modeling behavior, embedding is not the right choice for you to use. However, the fact we can't use them interchangeably is a very severe limitation. It is much better to use `interfaces` when you want to describe common behavior. So when is embedding a good idea? If you're authoring a library, for example, and let's just say you're making a web framework, and you've got a very sophisticated base controller. In that case, you might want to use consumers of your library to embed that base controller into their custom controllers, so that they can get useful functionality out of it. In that case, you are not talking about polymorphism and the ability to interchangeably use objects. You are just trying to get some base behavior into a custom type. And in that case, embedding makes sense. 

The last thing I want to tell you about structs is a concept called **tags**. Let's say, for example, we are working with some validation framework and the users filling out a form and two of the fields are providing the name and the origin. We want to make sure that the `Name` is required and doesn't exceed a maximum length. We can do that with a *tag* and the format of a tag is to have **backticks** (``) as the delimiters of the tag, and then we have these space delimited **key-value pairs**.

```go
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}
```

So the value of the tag is arbitrary, you could put any string that you want in there, but that is the conventional use within Go. If you do need a key-value relationship, then you are going to use a colon to separate the key and the value, and then the value is typically put in quotation marks. The question is, how do we get at it?

```go
package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
```

Output:
```
required max:"100"
```

Now we can get at the tag by using the `reflect` package and by asking for the `Tag` property of the field.

## Summary

#### Maps
* Collections of value types that are accessed via keys
* Created via literals or **`make`** function
* Members accessed via  `[key]` syntax
	* `myMap["key"] = "value"`
* Check for presence with "value, ok" form of result
* Multiple assignments refer to same underlying data

#### Structs
* Collections of disparate data types that describes a single concept
* Keyed by named fields
* Normally created as types, but anonymous structs are allowed
* Structs are value types
* No inheritance, but can use composition via embedding
* Tags can be added to struct fields to describe field

## Reference
* [Tags in Golang](https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e)
* [Go maps in action](https://go.dev/blog/maps)
* [Equality in Golang](https://medium.com/golangspec/equality-in-golang-ff44da79b7f1)
* [Composition and Embedding](https://riptutorial.com/go/example/1256/composition-and-embedding)
* [Structures in Go (structs)](https://medium.com/rungo/structures-in-go-76377cc106a2)