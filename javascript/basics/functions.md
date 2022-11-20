# Functions

Syntax:
```javascript
function funcName() {
   // ... code block ... //
}
```

You can call or invoke this function by using its name followed by parentheses, like this: `funcName();.` All of the code between the curly braces will be executed every time the function is called.

## Parameters
**Parameters** are variables that act as placeholders for the values that are to be input to a function when it is called. When a function is defined, it is typically defined along with one or more parameters. The actual values that are input (or "passed") into a function when it is called are known as *arguments*.

Example:
```javascript
function message(parameter_1, parameter_2) {
   console.log(parameter_1, parameter_2);
}

message("Hello", "World");
```

We have passed two string arguments, `Hello` and `World`. Inside the function, `parameter_1` will equal the string *Hello* and `parameter_2` will equal the string *World*.

## `return`

We can pass values into a function with *arguments*. You can use a **`return`** statement to send a value back out of a function.

```javascript
function sum(a, b) {
   return a + b;
}

let answer = sum(1, 2);
```

## Global Scope
In JavaScript, **scope** refers to the visibility of variables. Variables which are defined outside of a function block have *Global scope*. This means, they can be seen everywhere in your JavaScript code. Variables which are declared without the `let` or `const` keywords are automatically created in the global scope. This can create unintended consequences elsewhere in your code or when running a function again. You should always declare your variables with `let` or `const`.

```javascript
// Global variable
let x = 10;

function sum() {
   return x + 5;
}
```

## Local Scope
Variables which are declared within a function, as well as the function parameters, have *local scope*. That means they are only visible within that function.

Example:
```javascript
function message() {
   const name = "John";
   console.log(name);
}

message();
console.log(name);
```

The `message()` function call will display the string `John` in the console. The `console.log(name)` line (outside of the myTest function) will throw an error, as `name` is not defined outside of the function.

## Global vs. Local
It is possible to have both *local* and *global* variables with the same name. When you do this, the local variable takes precedence over the global variable.

```javascript
const name = "John";

function message() {
   const name = "Alex";
   return name;
}
```

The function `message` will return the string `Alex` because the local version of the variable is present.

## `undefined`
A function can include the `return` statement but it does not have to. In this case that the function doesn't have a `return` statement, when you call it, the function processes the inner code but the returned value is `undefined`.

```javascript
let sum = 0;

function add(num) {
   sum = sum + num;
}

add(3);
```

`add` is a function without a *`return`* statement. The function will change the global `sum` variable but the returned value of the function is `undefined`.

## Assignment with a Returned Value
Everything to the right of the equal sign is resolved before the value is assigned. This means we can take the return value of a function and assign it to a variable.

Example:
```javascript
total = sum(5, 12);
```

`sum` function returns a value of `17` and assigns it to the `total` variable.