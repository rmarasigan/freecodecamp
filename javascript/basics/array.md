# Arrays

With JavaScript `array` variables, we can store several pieces of data in one place. You start an array declaration with an opening square bracket, end it with a closing square bracket, and put a comma between each entry.

```javascript
const burgers = ["cheeseburger", "whooper jr.", "chicken fillet"];
```

## Array within another array
You can also nest arrays within other arrays. This is also called a *multi-dimensional* array.

```javascript
const teams = [["Lakers", 27], ["Celtics", 16]];
```

## Indexes
We can access the data inside arrays using **indexes**. Array indexes are written in the same bracket notation that strings use, except that instead of specifying a character, they are specifying an entry in the array. Like strings, arrays use zero-based indexing, so the first element in an array has an index of 0.

```javascript
const numbers = [1, 2, 3]
console.log(numbers[0]);   // 1
```

## Modify array data
The entries of arrays are mutable and can be changed freely, even if the array was declared with `const`.

**Note**: There shouldn't be any spaces between the array name and the square brackets, like `array [0]`. Although JavaScript is able to process this correctly, this may confuse other programmers reading your code.

```javascript
const numbers = [1, 2, 3];
numbers[0] = 0;
```
## Multi-dimensional
One way to think of a **multi-dimensional array**, is as *an array of arrays*. When you use brackets to access your array, the first set of brackets refers to the entries in the outer-most (the first level) array, and each additional pair of brackets refers to the next level of entries inside.

```javascript
const arr = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
  [[10, 11, 12], 13, 14]
];
```

## `push()`
An easy way to append data to the end of an array is via the `push()` function. `.push()` takes one or more parameters and "pushes" them onto the end of the array.

```javascript
const numbers = [1, 2, 3];
numbers.push(4);
```

## `pop()`
`.pop()` is used to pop a value off of the end of an array. We can store this popped off value by assigning it to a variable. In other words, `.pop()` removes the last element from an array and returns that element.

```javascript
const numbers = [1, 2, 3, 4, 5, 6];
const oneDown = numbers.pop();
console.log(oneDown);   // 6
```

## `shift()`
It works just like `.pop()`, except it removes the first element instead of the last.

```javascript
const numbers = [1, 2, 3, 4, 5, 6];
const oneDown = numbers.shift();
console.log(oneDown);   // 1
```

## `unshift()`
`.unshift()` works exactly like `.push()`, but instead of adding the element at the end of the array, `unshift()` adds the element at the beginning of the array.

```javascript
const numbers = [1, 2, 3, 4, 5, 6];
numbers.unshift(0);  // [0, 1, 2, 3, 4, 5, 6]
```