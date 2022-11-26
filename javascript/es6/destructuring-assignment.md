# Destructuring Assignment

### Extract Values from Objects
*Destructuring assignment* is special syntax introduced in ES6, for neatly assigning values taken directly from an object. It is a JavaScript expression that makes it possible to unpack values from arrays, or properties from objects, into distinct variables. Destructuring also works great with complex functions that have a lot of parameters, default values, and so on.

Example:
```javascript
const user = { name: 'John Doe', age: 34 };
const { name, age } = user;

console.log(name);  // John Doe
console.log(age);   // 34
```

The left side contains an object-like “pattern” for corresponding properties. In the simplest case, that’s a list of variable names in `{...}`.

### Assign Variables from Objects
Destructuring allows you to assign a new variable name when extracting values. You can do this by putting the new name after a colon when assigning the value.

Example:
```javascript
const { name: userName, age: userAge } = user;

console.log(userName);  // John Doe
console.log(userAge);   // 34
```

The colon shows “what : goes where”. 

### Assign Variables from Nested Objects
If an object contain another nested objects and arrays, we can use more complex left-side patterns to extract deeper portions.

Example:
```javascript
const user = {
   john: {
      age: 34,
      email: 'j.doe@freeCodeCamp.com'
   }
}

const { john: { age, email }} = user;
```

Here's how you can assign an object properties' values to variables with different names:
```javascript
const { john: { age: userAge, email: userEmail }} = user;
```

### Assign Variables from Arrays
One key difference between the spread operator and array destructuring is that the spread operator unpacks all contents of an array into a comma-separated list. Consequently, you cannot pick or choose which elements you want to assign to variables.

Example:
```javascript
const [a, b] = [1, 2, 3, 4, 5, 6];
console.log(a, b); // 1 2
```

We can also access the value at any index in an array with destructuring by using commas to reach the desired index:
```javascript
const [a, b,,, c] = [1, 2, 3, 4, 5, 6];
console.log(a, b, c); // 1 2 5
```

### Rest Parameter to Reassign Array Elements
In some situations involving array destructuring, we might want to collect the rest of the elements into a separate array.

For example, here only two items are taken, and the rest is just ignored:
```javascript
const [a, b, ...arr] = [1, 2, 3, 4, 5, 6];
console.log(a, b); // 1 2
console.log(arr);  // [3, 4, 5, 6]
```

The rest element only works correctly as the last variable in the list, As in, you cannot use the rest parameter to catch a subarray that leaves out the last element of the original array.

### Pass an Object as a Function's Parameters
In some cases, you can destructure the object in a function argument itself. We can pass parameters as an object, and the function immediately destructurizes them into variables.

Example:
```javascript
let userData = {
   name: "John Doe",
   age: "34",
   nationality: "American",
   location: "USA"
};

const profile = ({ name, age, nationality, location }) => {}
```
When `userData` is passed to the above function, the valeus are destructured from the function parameter for use within the function.


## Good Read
* [Destructuring assignment](https://javascript.info/destructuring-assignment)
* [Destructuring assignment](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)