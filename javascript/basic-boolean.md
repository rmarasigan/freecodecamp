# Boolean
**Booleans** may only be one of two values: `true` or `false`. They are basically little on-off switches, where true is on and false is off. These two states are mutually exclusive.

**Note**: Boolean values are never written with quotes. The strings `"true"` and `"false"` are not Boolean and have no special meaning in JavaScript.

Booleans are a primitive datatype commonly used in computer programming languages. By definition, a boolean has two possible values: `true` or `false`. In JavaScript, there is often implicit type coercion to boolean.

Example
```javascript
const x = "A string";

if (x) {
   console.log(x);
}
```

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

## Less Than (`<`)
The less than operator (`<`) compares the values of two numbers. If the number to the left is less than the number to the right, it returns `true`. Otherwise, it returns `false`. Like the equality operator, the less than operator converts data types while comparing.

Example:
```javascript
2   < 5 // true
'3' < 7 // true
5   < 5 // false
```

## Less Than or Equal to (`<=`)
The less than or equal to operator (`<=`) compares the values of two numbers. If the number to the left is less than or equal to the number to the right, it returns `true`. If the number on the left is greater than the number on the right, it returns `false`. Like the equality operator, the less than or equal to operator converts data types.

Example:
```javascript
4   <= 5 // true
'7' <= 7 // true
3   <= 2 // false
```

## Logical *AND* (`&&`)
The logical *`and`* operator (`&&`) returns true if and only if the operands to the left and right of it are true.

Example:
```javascript
if (num > 5 && num < 10)  {
   return "Yes";
}

return "No";
```

## Logical *OR* (`||`)
The logical *`or`* operator (`||`) returns true if either of the operands is `true`. Otherwise, it returns `false`. The logical *or* operator is composed of two pipe symbols: (`||`). This can typically be found between your Backspace and Enter keys.

Example:
```javascript
if (num > 10 || num < 5) {
   return "No";
}

return "Yes";
```

## Great Read
* [JavaScript booleans explained by going to court](https://www.freecodecamp.org/news/javascript-booleans-explained-by-going-to-court-a0ca1149a0dc/)
* [JavaScript Booleans Explained – How to use Booleans in JavaScript](https://www.freecodecamp.org/news/booleans-in-javascript-explained-how-to-use-booleans-in-javascript/)