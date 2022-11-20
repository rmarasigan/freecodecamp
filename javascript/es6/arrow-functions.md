# Arrow Functions

An **arrow function expression** is a compact alternative to a traditional function expression, with some semantic differences and deliberate limitations in usage:

* Arrow functions don't have their own bindings to `this`, `arguments`, or `super`, and should not be used as methods
* Arrow functions cannot be used as `constructors`. Calling them with `new` throws a `TypeError`. They also don't have access to the `new.target` keyword
* Arrow functions cannot use `yield` within their body and cannot be created as generator functions

In JavaScript, we often don't need to name our functions, especially when passing a function as an argument to another function. Instead, we create inline functions. We don't need to name these functions because we do not reuse them anywhere else.

We often use the following syntax:
```javascript
const myFunc = function() {
   const variable = "value";

   return variable;
}
```

ES6 provides us with the syntactic sugar to not have to write anonymouse functions this way. Instead, you can use **arrow function syntax**:

```javascript
const myFunc = () => {
   const variable = "value";

   return variable;
}
```

When there is no function body, and only a return value, arrow function syntax allows you to omit the keyword `return` as well as the brackets surrounding the code. This helps simplify smaller functions into one-line statements:

```javascript
const myFunc = () => "value";
```

This code will still return the string `value` by default.

## Arrow Functions with Parameters

Juset like a regular function, you can pass arguments into an arrow function.

```javascript
const doubler = (item) => item * 2;
doubler(4);  // 8
```

If an arrow function has a single parameter, the parentheses enclosing the parameter may be omitted.

```javascript
const doubler = item => item * 2;
```

It is possible to pass more than one argument into an arrow function.

```javascript
const multiplier = (item, multi) => item * multi;
multiplier(4, 2);  // 8
```

## Good Read
* [Arrow function expressions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions)
* [Arrow functions, the basics](https://javascript.info/arrow-functions-basics)
* [Introduction to JavaScript arrow functions](https://www.javascripttutorial.net/es6/javascript-arrow-function/)