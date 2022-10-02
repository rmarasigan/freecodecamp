# `switch` Statements

A `switch` statement tests a value and can have many `case` statements which define various possible values. Statements are executed from the first matched `case` value until a `break` is encountered.

In programming, a `switch` statement is a control-flow statement that tests the value of an `expression` against multiple cases.

```javascript
switch (expression) {
   case 1:
      // Execute if the case matches the expression
      break;
   case 2:
      // Execute if the case matches the expression
      break;
   case 3:
      // Execute if the case matches the expression
      break;
   default:
      // Execute if the case matches the expression
      break;
}
```

It will go through the `switch` statement and check for strict equality `===` between the `case` and `expression`. If one of the cases matches the `expression`, then the code inside that `case` clause will execute. If none of the cases match the expression, then the `default` clause will be executed. If multiple cases match the `switch` statement, then the first `case` that matches the `expression` will be used.

`break` statements will break out of the `switch` when the `case` is matched. If `break` statements are not present, then it will continue through the `switch` statement even if a match is found. If the `return` statements are present in the `switch`, then you don't need a `break` statement.

## `default` option
In a `switch` statement you may not be able to specify all possible values as `case` statements. Instead, you can add the **`default`** statement which will be executed if no matching `case` statements are found. Think of it like the final `else` statement in an `if/else` chain. A `default` statement should be the last case.

Example:
```javascript
switch (number) {
   case 1:
      // Statement #1
      break;
   case 2:
      // Statement #2
      break;
   default:
      // Default statement
      break;
}
```

## Multiple Identical Options
If the `break` statement is omitted from a `switch` statement's `case`, the following `case` statement(s) are executed until a `break` is encountered.

Example:
```javascript
let result = "";

switch (val) {
   case 1:
   case 2:
   case 3:
      result = "1, 2, or 3";
      break;
   case 4:
      result = "4 alone";
}
```

## Great Read
* [JavaScript Switch Case – JS Switch Statement Example](https://www.freecodecamp.org/news/javascript-switch-case-js-switch-statement-example/)