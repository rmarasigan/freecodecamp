# `const` Keyword

ES6 provides a new way of declaring a constant by using the `const` keyword. The `const` keyword creates a read-only reference to a value. By convention, the constant identifiers are in uppercase. Like the `let` keyword, the `const` keyword declares blocked-scope variables. However, the block-scoped variables declared by the const keyword can't be *reassigned*.

```javascript
const CONSTANT_NAME = value;
```

The `const` declaration has many use cases in modern JavaScript. Some developers prefer to assign all their variables using `const` by default, unless they know they will need to reassign the value. Only in that case, they use `let`.

However, it is important to understand that objects (including arrays and functions) assigned to a variable using `const` are still mutable. Using the `const` declaration only prevents reassignment of the variable identifier.

Example:
```javascript
const s = [5, 6, 7];
s = [1, 2, 3];
s[2] = 25;

console.log(s);
```

`s = [1, 2, 3];` will result in an error. After commenting out that line, the `console.log` will display the value `[5, 6, 45]`. As you can see, you can mutate the object `[5, 6, 7]` itself and the variable `s` will still point to the altered array `[5, 6, 45]`. Like all arrays, the array elements in `s` are mutable, but because const was used, you cannot use the variable identifier s to point to a different array using the assignment operator.

## Good Read
* [JavaScript const: Declaring Constants in ES6](https://www.javascripttutorial.net/es6/javascript-const/)