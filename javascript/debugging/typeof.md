# JavaScript `typeof`

You can use `typeof` to check the data structure, or type, of a variable. This is useful in debugging when working with multiple data types. If you think you're adding two numbers, but one is actually a string, the results can be unexpected. Type errors can lurk in calculations of function calls. Be careful especailly when you're accessing and working with external data in the form of a JavaScript Object Notation(JSON) object.

The `typeof` operator is useful because it is an easy way to check the type of a variable in your code. this is important because JavaScript is a *dynamically typed language*. This means that you aren't required to assign types to variables when you create them because a variable is not restricted in this way, its type can change during the runtime of a program.

You use it by typing `typeof(variable)` or `typeof variable`. This can be useful for checking the type of a variable in a function and continuing as appropriate. Another way the `typeof` operator can be usefule is by ensuring that a variable is defined before you try to access it in your code. This can help prevent errors in a program that may occur if you try to access a variable that is not defined.

Here are some examples of using `typeof`:
```javascript
console.log(typeof "");   // string
console.log(typeof 0);    // number
console.log(typeof []);   // object
console.log(typeof {});   // object
```

JavaScript recognizes seven primitive (immutable) data types: `Boolean`, `Null`, `Undefined`, `Number`, `String`, `Symbol`, and `BigInt`, and one type for mutable items: `Object`. Note that in JavaScript, arrays are technically a type of object.

## Good Read
* [Typeof API](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/typeof)
* [JavaScript Data Types: Typeof Explained](https://www.freecodecamp.org/news/javascript-data-types-typeof-explained/)
* [What is the difference between a strongly typed language and a statically typed language?](https://stackoverflow.com/questions/2690544/what-is-the-difference-between-a-strongly-typed-language-and-a-statically-typed)