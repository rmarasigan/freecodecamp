# `if` — `else if` — `else` Statements
The `if...else` is a type of a conditional statement that will execute a block of code when the condition in the `if` statement is truthy. If the condition is falsy, then the `else` block will be executed.

```javascript
if (condition is true) {
   // Code is executed
} else {
   // Code is executed
}
```

There will be times where you want to test multiple conditions. That is where the `else if` block comes in.

```javascript
if (condition #1 is true) {
   // Condition is executed
} else if (condition #2 is true) {
   // Code is executed
} else {
   // Code is executed
}
```

When the `if` statement is `false`, it will move onto the `else if` statement. If that is also `false`, then it will move onto the `else` block.

## `if` Statements
**`if`** statements are used to make decisions in code. The keyword `if` tells JavaScript to execute the code in the curly braces under certain conditions, defined in the parentheses. These conditions are known as *Boolean* conditions and they may only be `true` or `false`.

When the condition evaluates to `true`, the program executes the statement inside the curly braces. When the Boolean condition evaluates to `false`, the statement inside the curly braces will not execute.

Syntax:
```javascript
if (true) {
   // ... statement is executed ... //
}
```

Example:
```javascript

function printMessage(condition) {
   if (condition) {
      return "It was true";
   }

   return "It was false";
}

printMessage(true);
printMessage(false);
```

`printMessage(true)` returns the string `It was true`, and `printMessage(false)` returns the string `It was false`.

## `else` Statements
When a condition for an `if` statement is true, the block of code following it is executed. What about when that condition is false? Normally nothing would happen. With an `else` statement, an alternate block of code can be executed.

Example:
```javascript
if (number > 10) {
   return "The number is bigger than 10.";
} else {
   return "The number is 10 or less";
}
```

## `else if` Statements
If you have multiple conditions that need to be addressed, you can chain `if` statements together with `else if` statements.

Example:
```javascript
if (number > 15) {
   return "The number is bigger than 15";
} else if (number < 5) {
   return "The numbser is smaller than 5";
} else {
   return "The number is between 5 and 15";
}
```

## Logical Order in `if else` statements
Order is important in `if`, `else if` statements. The function is executed from top to bottom so you will want to be careful of what statement comes first.

## Chaining `if else` statements
`if/else` statements can be chained together for complex logic.

Example:
```javascript
if (condition #1) {
   // Code block here
} else if (condition #2) {
   // Code block here
} else if (condition #3) {
   // Code block here
} else {
   // Code block here
}
```

## Great Read
* [JavaScript If-Else and If-Then – JS Conditional Statements](https://www.freecodecamp.org/news/javascript-if-else-and-if-then-js-conditional-statements/)