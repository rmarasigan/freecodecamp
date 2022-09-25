# Boolean
**Booleans** may only be one of two values: `true` or `false`. They are basically little on-off switches, where true is on and false is off. These two states are mutually exclusive.

**Note**: Boolean values are never written with quotes. The strings `"true"` and `"false"` are not Boolean and have no special meaning in JavaScript.

## Conditional Logic with `if` Statements
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

## Equality Operator
There are many *comparison operators* in JavaScript. All of these operators return a boolean `true` or `false` value. The most basic operator is the equality operator **`==`**. The equality operator compares two values and returns `true` if they're equivalent or `false` if they are not. Note that equality is different from assignment (`=`), which assigns the value on the right of the operator to a variable on the left.

Example:
```javascript
function equalityCheck(val) {
   if (val == 10) {
      return "Equal";
   }

   return "Not equal";
}
```

In order for JavaScript to compare two different data types (for example, *numbers* and *strings*), it must convert one type to another. This is known as **Type Coercion**.

```javascript
1 == 1      // true
1 == 2      // false
1 == '1'    // true
"3" == 3    // true
```

## Strict Equality
**Strict equality** (`===`) is the counterpart to the equality operator (`==`). However, unlike the equality operator, which attempts to convert both values being compared to a common type, the strict equality operator does not perform a type conversion. If the values being compared have different types, they are considered unequal, and the strict equality operator will return false.

Example:
```javascript
3 === 3    // true
3 === '3'  // false
```

## Inequality
The **inequality operator** (`!=`) is the opposite of the equality operator. It means not equal and returns `false` where equality would return `true` and *vice versa*. Like the equality operator, the inequality operator will convert data types of values while comparing.

Examples:
```javascript
1 !=  2    // true
1 != "1"   // false
1 != '1'   // false
1 != true  // false
0 != false // false
```

## Strict Inequality
The strict inequality operator (`!==`) is the logical opposite of the strict equality operator. It means "Strictly Not Equal" and returns `false` where strict equality would return `true` and *vice versa*. The strict inequality operator will not convert data types.

Examples:
```javascript
3 !==  3  // false
3 !== '3' // true
4 !==  3  // true
```

## Greater Than (`>`)
The greater than operator (`>`) compares the values of two numbers. If the number to the left is greater than the number to the right, it returns `true`. Otherwise, it returns `false`. Like the equality operator, the greater than operator will convert data types of values while comparing.

Example:
```javascript
5 > 3  // true
1 > 2  // false
```

## Greater Than or Equal To (`>=`)
The greater than or equal to operator (`>=`) compares the values of two numbers. If the number to the left is greater than or equal to the number to the right, it returns `true`. Otherwise, it returns `false`. Like the equality operator, the greater than or equal to operator will convert data types while comparing.

Example:
```javascript
6  >=  6  // true
7  >= '3' // true
2  >=  3  // false
```