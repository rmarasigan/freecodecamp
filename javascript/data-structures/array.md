# Array

### Use an array to store a collection of data
An **array** is a collection of items stored at contiguous memory locations. Each item can be accessed through its **index** (position) number. Arrays *always starts at index 0*, so in an array of 4 elements we could access the 3rd element using the index number 2. The **length** property of an array is defined as the number of elements it contains. If the array contains 4 elements, we can say the array has a length of 4.

In some programming languages, the user can only store values of the same type in one array and the length of the array has to be defined at the moment of its creation and can't be modified afterwards. In JavaScript that's not the case, as *we can store values of any type* in the same array and the length of it can be dynamic (it can grow or shrink as much as necessary).

The below is an example of the simplest implementation of an array data structure. This is known as a *one-dimensional array*, meaning it only has one level, or that it does not have any other arrays nested within it. Notice it contains *booleans*, *strings*, and *numbers*, among other valid JavaScript data types:

```javascript
let simpleArray = ['one', 2, 'three', true, false, undefined, null];
console.log(simpleArray.length);  // 7
```

A more complex implementation of an array can be seen below. This is known as a *multi-dimensional array*, or an array that contains other arrays. Notice that this array also contains JavaScript *objects*. Arrays are also capable of storing complex objects.

```javascript
let complexArray = [
  [
    {
      one: 1,
      two: 2
    },
    {
      three: 3,
      four: 4
    }
  ],
  {
    a: "a",
    b: "b"
  },
  {
    c: "c",
    d: "d"
  }
];
```

In JavaScript, arrays come with many built-in properties and methods we can use with different purposes, such as adding or deleting items from the array, sorting it, filtering its values, length and so on. Wehn we add a new item at the end of the array, it just takes the index number that follows the previous last item in the array.

But when we add/delete a new item *at the beginning or the middle* of the array, the indexes of all the elements that come after the element added/deleted have to be changed. This of cours has a computational cost, and is one of the weaknesses of this data structure.

Arrays are useful when we have to store individual values and add/delete values from the end of the data structure but when we need to add/delete from any part of it, there are other data structures that perform more efficiently.

### Access an array's contents using bracket notation
The fundamental feature of any data structure is, of course, the ability to not only store data, but to be able to retrieve that data on command. When we define a simple array as seen below, there are 3 items in it:

```javascript
let ourArray = ["a", "b", "c"];
```

In an array, each array item has an **index**. This index doubles as the position of that item in the array, and how you reference it. However, it is important to note, that JavaScript arrays are *zero-indexed*, meaning that the first element of an array is actually at the **zeroth** position, not the first. In order to retrieve an element from an array, we can enclose an index in brackets and append it to the end of an array, or more commonly, to a variable which references an array object. This is known as *bracket notation*.

```javascript
let a = ourArray[0];
console.log(a);   // a
```

In addition to accessing the value associated with an index, you can also set an index to a value using the same notation.

```javascript
ourArray[1] = "not b anymore";
console.log(ourArray);    // ["a", "not b anymore", "c"]
```

## `push()`, `pop()`, `shift()`, `unshift()`

#### `push()`
We use `push()` when we want to **add something to the end** of an array. This method allows us to add one or many items to the end of an array. The push method returns the new length of the array.

#### `pop()`
We use `pop()` when we want to **remove something from the end** of an array. The `pop()` method returns the item that was removed. If the array is empty, it will return `undefined`.

#### `shift()`
We use `shift()` when we want to **remove something from the beginning** of an array. The  `shift()` method returns the item that was removed. If the array is empty, it will return `undefined`.

#### `unshift()`
We use `unshift()` when we want to **add something to the beginning** of an array.This method allows us to add one or many items to the beginning of an array. The `unshift()` method returns the new length of the array.

### Add items to an array with `push()` and `unshift()`
An array's length, like the data types it can contain, is not fixed. Arrays can be defined with a length of any number of elements, and elements can be added or removed over time; in other words, arrays are *mutable*.

The `push()` method adds elements to the end of an array, and `unshift()` adds elements to the beginning.

```javascript
let twentyThree = 'XXIII';
let romanNumerals = ['XXI', 'XXII'];

romanNumerals.unshift('XIX', 'XX');   // ['XIX', 'XX', 'XXI', 'XXII']
romanNumerals.push(twentyThree);      // ['XIX', 'XX', 'XXI', 'XXII', 'XXIII']
```

Notice that we can also pass variables, which allows us even greater flexibility in dynamically modifying our array's data.

### Remove items from an array with `pop()` and `shift()`
Both `push()` and `unshift()` have corresponding methods that are nearly functional opposites: `pop()` and `shift()`. Instead of adding, `pop()` *removes an element from the end* of an array, while `shift()` *removes an element from the beginning*. The key difference between `pop()` and `shift()` and their cousins `push()` and `unshift()`, is that neither method takes parameters, and each only allows an array to be modified by a single element at a time.

Let's take a look:

```javascript
let greetings = ['whats up?', 'hello', 'see ya!'];
greetings.pop();    // ['whats up?', 'hello']
greetings.shift();  // ['hello']
```

We can also return the value of the removed element with either method like this:

```javascript
let popped = greetings.pop();
```

`greetings` would have the value `[]`, and `popped` would have the value `hello`.

## `splice()`
The `splice()` method is a built-in method for JavaScript Array objects. It lets you change the content of your array by removing or replacing existing elements with new ones. This method modifies the original array and returns the removed elements as a new array. The `splice()` method needs at least one parameter, which is the *start* index where the splice operation starts. You can also define how many elements you wat to remove from the array by passig a second number argument known as `removeCount`.

You can add new elements without removing any by passing the number `0` to the `removeCount` parameter. When no elements are removed, the splice method will return an empty array. You can choose whether to store the returned empty array to a variable or not.

```javascript
Array.splice(start, removeCount);
```

The `splice()` method is mostly used when you need to delete or add new elements to an array. In some situations, you can also use it to separate an array which has mixed content.

### Remove items using `splice()`
`splice()` allows us to do just that: **remove any number of consecutive elements from anywhere** in an array. `splice()` can take up to 3 parameters, but for now, we'll focus on just the first 2. The first two parameters of `splice()` are integers which represents indexes, or positions, of items in the array that `splice()` is being called upon. Remember, arrays are *zero-indexed*, so to indicate the first element of an array, we would use `0`. `splice()`'s **first parameter** represents the *index on the array from which to begin removing elements*, while the **second parameter** indicates the *number of elements to delete*.

For example:

```javascript
let array = ['today', 'was', 'not', 'so', 'great'];
array.splice(2, 2);   // ['today', 'was', 'great']
```

`splice()` not only modifies the array it's being called on, but it also returns a new array containing the value of the removed elements:

```javascript
let array = ['I', 'am', 'feeling', 'really', 'happy'];
let newArray = array.splice(3, 2);    // ['really', 'happy']
```

### Add items using `splice()`
You can use the third parameter, comprised of one or more element(s), to add to the array. This ca be incredibly useful for quickly switching out an element, or a set elements, for another.

```javascript
let numbers = [10, 11, 12, 12, 5]
const startIndex = 3;
const amountToDelete = 1;

numbers.splice(startIndex, amountToDelete, 13, 14);
console.log(numebers);
```

The second occurrence of `12` is removed, and we add `13` and `14` at the same index. The `numbers` array would now be `[10, 11, 12, 13, 14, 15]`.

Here, we begin with an array of numbers. Then, we pass the following to `splice()`: The index at which to begin deleting elements (`amountToDelete`), the number of elements to be deleted (`amountToDelete`), and the remaining arguments (`13, 14`) will be inserted starting at the same index. Note that there can be any number of elements (separated by commas).

## `slice()`
The `slice()` method can be used to create a copy of an array or return a portion of an array. It is important to note that the `slice()` method does not alter the original array but instead creates a shallow copy.

```javascript
Array.slice(start, end);
```

You can use the optional `start` parameter to set a starting position for selecting elements from the array. It is important to remember that arrays are zero-based indexed. You can also use negative indexes which will start extracting elements from the end of the array. But if the `start` parameter is greater than the last index of the array, then an empty array will be returned.

If an `end` position is specified, then the `slice()` method will extract elements from the `start` position up to the `end` position. The `end` position will not be included in the extracted elements for the new array.

### Copy array items using `slice()`
Rather than modifying an array, `slice()` copies or *extracts* a given number of elements to a new array, leaving the array it is called upon untouched. `slice()` takes only 2 parameters — the **first** is the *index at which to begin extraction*, and the **second** is the *index at which to stop extraction* (extraction will occur up to, but not including the element at this index). Consider this:

```javascript
let weathers = ['rain', 'snow', 'sleet', 'hail', 'clear'];
let todaysWeather = weathers.slice(1, 3);
```

`todaysWeather` would have the value `['snow', 'sleet']`, while `weathers` would still have `['rain', 'snow', 'sleet', 'hail', 'clear']`. In effect, we have created a new array by extracting elements from an existing array.

## Spread Operator
The *spread operator* is useful and quick syntax for adding items to arrays, combining arrays or objects, and spreading an array out into a function's arguments.  The spread (`...`) syntax allows an iterable, such as an array or string, to be expanded in places where zero or more arguments (for function calls) or elements (for array literals) are expected.

The spread operator is useful for many different routine tasks in JavaScript, including the following:
* Copying an array
* Concatenating or combining arrays
* Using Math functions
* Using an array as arugments
* Adding an item to a list
* Adding to state in React
* Combining objects
* Converting NodeList to an array

### Copy an array with the spread operator
ES6's new spread operator allows us to easily copy all of an array's elements, in order, with a simple and highly readable syntax. The spread syntax simply look like this: `...`. In practice, we can use the spread operator to copy an array like so:

```javascript
let array = [true, true, undefined, false, null];
let copy = [...array];
```

The `array` remains unchanged and `copy` contains the same elements as `array`.

### Combine arrays with the spread operator
Another huge advantage of the *spread operator* is the ability to combine arrays, or to insert all the elements of one array into another, at any index. With more traditional syntaxes, we can concatenate arrays, but this only allows us to combine arrays at the end of one, and at the start of the another. Spread syntax makes the following operation extremely simple:

```javascript
let array = ['sage', 'rosemary', 'parsley', 'thyme'];
let newArray = ['basil', 'cilantro', ...array, 'coriander'];
```

`newArray` would have the value `['basil', 'cilantro', 'sage', 'rosemary', 'parsley', 'thyme', 'coriander']`. Using spread syntax, we have just achieved an operation that would have been more complex and more verbose had we used traditional methods.

## `indexOf()`
The JavaScript string `indexOf()` method is used to search the position of a particular character or string in a sequence of a given char values. This method is case-sensitive. The index position of first character in string always starts with zero. If an element is not present in a string, it returns `-1`.

Syntax:

```javascript
let index = str.indexOf(substr, [, fromIndex]);
```

The `fromIndex` is an optional parameter that specifies the index at which the search starts. It defaults to zero (0), meaning that if you omit the `fromIndex`, the search will start from the beginning of the string.

### Check for the presence of an element with `indexOf()`
Since arrays can be changed, or *mutated*, at any time, there's no guarantee about where a particular piece of data will be on a given array, or if that element even still exists. Luckily, JavaScript provides us with another built-in method, **`indexOf()`**, that allows us to quickly and easily check for the presence of an element on an array. `indexOf()` takes an element as a parameter, and when called, it returns the position, or index, of that element, or -1 if the element does not exist on the array.

For example:

```javascript
let fruits = ['apples', 'pears', 'oranges', 'peaches', 'pears'];
fruits.indexOf('dates');      // -1
fruits.indexOf('oranges');    // 2
fruits.indexOf('pears');      // 1 (the first index at which element exists)
```

## `for` Loops
The `for` loop is an iterative statement which you use to check for certain conditions and then repeatedly execute a block of code as long as those conditions are met.

Syntax:

```javascript
for (initialExpression; condition; updateExpression) {}
```

* **initialExpression**: This is used to set the value of a counter variable, and it is only evaluated once, before the loop starts. Depending on the scope, these counter variables are usually declared with the `var` and `let` keywords.
* **condition**: This is a constant-evaluation expression that determines whether the loop should be executed. In simple terms, if this condition returns true, the `for` loop's block of code is executed. If it returns false, the `for` loop is terminated.
* **updateExpression**: This is commonly used to update or increment the `initialExpression` counter variable. In other words, when the condition is true, it updates the value of the `initialExpression`.

### Iterate through all an array's items using `for` loops
Sometimes when working with arrays, it is very handy to be able to iterate through each item to find one or more elements that we might need, or to manipulate an array based on which data items meet a certain set of criteria. JavaScript offers several built in methods that each iterate over arrays in slightly different ways to achieve different results (such as `every()`, `forEach()`, `map()`, etc.), however the technique which is most flexible and offers us the greatest amount of control is a simple `for` loop.

Consider the following:

```javascript
function greaterThanTen(arr) {
  let newArr = [];
  
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] > 10) {
      newArr.push(arr[i]);
    }
  }

  return newArr;
}

greaterThanTen([2, 12, 8, 14, 80, 0, 1]);   // [12, 14, 80]
```

## Good Read
* [Spread syntax (...)](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Spread_syntax)
* [Rest parameters and spread syntax](https://javascript.info/rest-parameters-spread)
* [JavaScript String indexOf() method](https://www.javatpoint.com/javascript-string-indexof-method)
* [JavaScript For Loop – Explained with Examples](https://www.freecodecamp.org/news/javascript-for-loops/)
* [How to Use the Spread Operator (…) in JavaScript](https://medium.com/coding-at-dawn/how-to-use-the-spread-operator-in-javascript-b9e4a8b06fab)
* [JavaScript Arrays: push(), pop(), shift() & unshift()](https://medium.com/an-idea/javascript-arrays-push-pop-shift-unshift-adc8fb815fc0)
* [JavaScript Splice – How to Use the .splice() JS Array Method](https://www.freecodecamp.org/news/javascript-splice-how-to-use-the-splice-js-array-method/)
* [How to Use the slice() and splice() JavaScript Array Methods](https://www.freecodecamp.org/news/javascript-slice-and-splice-how-to-use-the-slice-and-splice-js-array-methods/)
* [Javascript Basics: Use .push, .pop, .shift, and .unshift to Manipulate Arrays](https://dev.to/antdp425/javascript-basics-use-push-pop-shift-and-unshift-to-manipulate-arrays-48ab)