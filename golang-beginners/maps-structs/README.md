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

## Reference
* [Go maps in action](https://go.dev/blog/maps)
* [Equality in Golang](https://medium.com/golangspec/equality-in-golang-ff44da79b7f1)