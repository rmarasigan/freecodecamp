# Conditional (Ternary) Operator

The *conditional operator*, also called the *ternary operator*, can be used as a one line `if-else` expression.

The operator is represented by a question mark `?`. Sometimes it's called "ternary", because the operator has three operands. It is actually one and only operator in JavaScript which has that many.

Syntax:
```javascript
let result = condition ? value_1 : value_2;
```

The **`condition`** is evaluated: if it's truthy then *`value_1`* is returned, otherwise – *`value_2`*.

The following function uses an `if-else` statement to check a condition:
```javascript
function findGreater(a, b) {
   if (a > b) {
      return "a is greater";
   }

   else {
      return "b is greater or equal";
   }
}
```

This can be re-written using the conditional operator:
```javascript
function findGreater(a, b) {
   return a > b ? "a is greater" : "b is greater or equal";
}
```

## Multiple Conditional (Ternary) Operators
The ternary operator is right-associative, which means it can be "chained" in the following way, similar to an `if ... else if ... else if ... else` chain.

Example:
```javascript
function findGreaterOrEqual(a, b) {
   if (a === b) {
      return "a and b are equal";
   }

   else if (a > b) {
      return "a is greater";
   }

   else {
      return "b is greater";
   }
}
```

The above function can be re-written using multiple conditional operators:
```javascript
function findGreaterOrEqual(a, b) {
   return (a === b) ? "a and b are equal"
   : (a > b) ? "a is greater"
   : "b is greater";
}
```

It is considered best practice to format multiple conditional operators such that each condition is on a separate line, as shown above. Using multiple conditional operators without proper indentation may make your code hard to read.



## Reference
* [Conditional (ternary) operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Conditional_Operator)
* [Conditional branching: if, '?'](https://javascript.info/ifelse)