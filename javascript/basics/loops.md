# Loops

There are many different kinds of loops, but they all essentially do the same thing: they repeat an action some number of times. The various loop mechanisms offer different ways to determine the start and end points of the loop. There are various situations that are more easily served by one type of loop over the others.

The statements for loops provided in JavaScript are:
* [for statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#for_statement)
* [do...while statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#do...while_statement)
* [while statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#while_statement)
* [labeled statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#labeled_statement)
* [break statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#break_statement)
* [continue statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#continue_statement)
* [for...in statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#for...in_statement)
* [for...of statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration#for...of_statement)

## Iterate with While Loops
The first type of loop we will learn is called a `while` loop because it runs while a specified condition is true and stops once that condition is no longer true.

Example:
```javascript
let i = 0;
const arr = [];

while (i < 5) {
   arr.push(i);
   i++;
}
```

In the example above, the `while` loop will execute 5 times and append the numbers 0 through 4 to `arr`.

## Iterate with For Loops
The most common type of JavaScript loop is called a `for` loop because it runs for a specific number of times. `for` loops are declared with three optional expressions separated by semicolons.

Syntax:
```javascript
for (a; b; c)
```

Where:
* `a` is the initialization statement
* `b` is the condition statement
* `c` is the final expression

The **initialization** statement is executed one time only before the loop starts. It is typically used to define and setup your loop variable.

The **condition** statement is evaluated at the beginning of every loop iteration and will continue as long as it evaluates to `true`. When the condition is `false` at the start of the iteration, the loop will stop executing. This means if the condition starts as false, your loop will never execute.

The **final** expression is executed at the end of each loop iteration, prior to the next condition check and is usually used to increment or decrement your loop counter.

Example:
```javascript
const arr = [];

for (let i = 0; i < 5; i++) {
   arr.push(i);
}
```

### Array with a For Loop
A common task in JavaScript is to iterate through the contents of an array. One way to do that is with a `for` loop. Remember that arrays have *zero-based indexing*, which means, the last index of the array is `length - 1`.

Example:
```javascript
const arr = [10, 9, 8, 7, 6];

for (let i = 0; i < arr.length; i++) {
   console.log(arr[i]);
}
```

### Nesting For Loops
If you have a multi-dimensional array, you can use the same logic as the prior waypoint to loop through both the array and any sub-arrays.

Example:
```javascript
const arr = [
   [1, 2], [3, 4], [5, 6]
];

for (let i = 0; i < arr.length; i++) {
   for (let j = 0; j < arr[i].length; j++) {
      console.log(arr[i][j]);
   }
}
```

## Iterate with Do...While Loops
It is called a `do...while` loop because it will first `do` one pass of the code inside the loop no matter what, and then continue to run the loop `while` the specified condition evaluates to `true`.

Example:
```javascript
let i = 0;
const arr = [];

do {
   arr.push(i);
   i++;
} while (i < 5);
```

The example above behave similar to other types of loops, and the resulting array will look like `[0, 1, 2, 3, 4]`. However, what makes the `do...while` different from other loops is how it behaves when the condition fails on the first check.

## Reference
* [Loops and iteration](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Loops_and_iteration)
* [Loops: while and for](https://javascript.info/while-for)