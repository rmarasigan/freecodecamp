# Numbers

**`Number`** is a data type in JavaScript which represents numeric data.

```javascript
const sum = 5 + 10;
const product = 5 * 10;
const quotient = 10 / 5;
const difference = 5 - 10;

const increment = sum++;
const decrement = sum--;
const remainder = 5 % 2;
```

## Decimal Numbers
We can store decimal numbers in variables too. Decimal numbers are sometimes referred to as *floating point* numbers or *floats*.

**NOTE**: When you compute numbers, they are computed with finite precision. Operations using floating points may lead to different results than the desired outcome.

## Remainder
The **remainder (`%`)** operator gives the remainder of the division of two numbers. In mathematics, a number can be checked to be even or odd by checking the remainder of the division of the number `2`.

```
5 % 2 = 1 because
2 * 2 = 4
5 - 4 = 1 (Remainder)
```

**NOTE**: The *remainder* operator is sometimes incorrectly referred to as the modulus operator. It is very similar to modulus, but does not work properly with negative numbers.

## Compound Assignment

### Augmented Addition, Subtraction, Multiplication and Division
In programming, it is common to use assignments to modify the contents of a variable. Remember that everything to the right of the equals sign is evaluated first.

```javascript
let sum = 1;
sum += 5;

let difference = 10;
difference -= 5;

let product = 2;
product *= 5;

let quotient = 10;
quotient /= 2;
```