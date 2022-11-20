# Functions

### Arguments vs. Parameters
Sometimes, you can use the terms argument and parameter interchangeably. However, by definition, parameters are what you specify in the function declaration whereas the arguments are what you pass into the function.

## Default Parameters
In order to help us create more flexible functions, the ES6 introduces *default parameters* for functions. Default function parameters allow named parameters to be initialized with default values if no value or `undefined` is passed.

Example:
```javascript
const greeting = (name = "Anonymous") => "Hello " + name;

console.log(greeting("John")); // Hello John
console.log(greeting());       // Hello Anonymous
```

The default parameter kicks in when the argument is not specified (it is undefined). As you can see in the example above, the parameter `name` will receive its default value `Anonymous` when you do not provide a value for the parameter. You can add default values for as many parameters as you want.

## Rest Parameter `...`
The rest parameter syntax allows a function to accept an indefinite number of arguments as an array, providing a way to represent [variadic functions](https://en.wikipedia.org/wiki/Variadic_function) in JavaScript. With rest parameter, you can create functions that take a variable number of arguments. These arguments are stored in an array that can be accessed later from inside the function.

Example:
```javascript
function count(...args) {
   return "You have passed " + args.length + " arguments.";
}

console.log(count(0, 1, 2));                        // You have passed 3 arguments.
console.log(count("string", null, [1, 2, 3]. {}));  // You have passed 4 arguments.
```

The rest parameter eliminates the need to check the `args` array and allows us to apply `map()`, `filter()` and `reduce()` on the parameters array.

## Spread Operator
The spread (`...`) syntax allows an iterable, such as an array or string, to be expanded in places where zero or more arguments (for function calls) or elements (for array literals) are expected. In object literal, the spread syntax enumerates the properties of an object and adds the key-value pairs to the object being created.

ES6 introduces the *spread operator* which allows us to expand arrays and other expressions in places where multiple parameters or elements are expected. 

The ES5 code below uses `apply()` to compute the maximum value in an array:

```javascript
var arr = [6, 89, 3, 45];
var maximus = Math.max.apply(null, arr);
```

We had to use `Math.max.apply(null, arr)` because `Math.max(arr)` returns `NaN. Math.max()` expects comma-separated arguments, but not an array. The spread operator makes this syntax much better to read and maintain.

```javascript
const arr = [6, 89, 3, 45];
const maximus = Math.max(...arr);  // 89
```

`...arr` returns an unpacked array. In other words, it *spreads* the array. However, the spread operator only works in-place, like in an argument to a function or in an array literal. The following code will not work:

```javascript
const spreaded = ...arr;
```

## Good Read
* [Rest parameters](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/rest_parameters)
* [Default parameters](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Default_parameters)
* [Spread syntax (...)](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Spread_syntax)
* [JavaScript Rest Parameters](https://www.javascripttutorial.net/es6/javascript-rest-parameters/)
* [JavaScript Default Parameters](https://www.javascripttutorial.net/es6/javascript-default-parameters/)
* [Rest parameters and spread syntax](https://javascript.info/rest-parameters-spread)