# Basic JavaScript

JavaScript is a scripting language you can use to make web pages interactive. It is one of the core technologies of the web, along with HTML and CSS, and is supported by all modern browsers.

## Comment
Comments are lines of code that JavaScript will intentionally ignore. Comments are a greate way to leave notes to yourself and to other people who will later need to figure out what that code does.

### In-line comment
```javascript
// This is an in-line comment
```

### Multi-line comment
```javascript
/* This is a
 multi-line comment */
```

## Variables
JavaScript provides eight different *data types* which are `undefined`, `null`, `boolean`, `string`, `symbol`, `bigint`, `number`, and `objects`. **Variables** allow computers to store and manipulate data in a dynamic fashion. They do this by using a "label" to point to the data rather than using the data itself. Any of the eight data types may be stored in a variable.

Variables are similar to the *x* and *y* you use in mathematics, which means they're simple name to represent the data we want to refer to. Computer variables differ from mathematical variables in that they can store different values at different times.

Syntax:
```javascript
var variable_name;

// Storing a value in a variable using the
// assignment operator (=)
variable_name = value
```

Variable names can be made up of numbers, letters, `$` or `_`, but may not contain spaces or start with a number. You can also assign the value of a variable to another variable using the *assignment operator*.

Example:
```javascript
var x;
x = 5;

var y;
y = x;
```

### Initializing a variable
It is common to *initialize* a variable to an initial value in the same line as it is declared.

```javascript
var variable_name = value;
```

### Uninitialized Variables
When JavaScript variables are declared, they have an initial value of `undefined`. If you do a mathematical operation on an `undefined` variable your result will be `NaN` which means *"Not a Number"*. If you concatenate a string with an `undefined` variable, you will get a string of `undefined`.

### Case Sensitivity in Variables
In JavaScript all variables and function names are case sensitive. This means that capitalization matters. Write variable names in JavaScript in *camelCase*. In *camelCase*, multi-word variable names have the first word in lowercase and the first letter of each subsequent word is capitalized.

```javascript
var someVariable;
var anotherVariableName;
```