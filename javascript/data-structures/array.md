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