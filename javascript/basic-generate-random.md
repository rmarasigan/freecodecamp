# Generate Random Number

A **PRNG** (pseudorandom number generator) is an algorithm that outputs numbers in a complex, seemingly unpredictable pattern. Truly random numbers are utterly unpredictable, whereas all algorithms are predictable, and a PRNG returns the same numbers when passed the same starting parameters or *seed*.

### `Math.random()`
The `Math.random()` functions returns a floating-point, pseudo-random number that's greater than or equal to 0 and less than 1, with approximately uniform distribution over that range - which you can scale to your desired range.

### `Math.floor()`
The `Math.floor()` function always rounds down and returns the largest integer less than or equal to a given number.

## Fractions with JavaScript
JavaScript has a `Math.random()` function that generates a random decimal number between `0` (inclusive) and `1` (exclusive). Thus `Math.random()` can return `0` but never return a `1`.

**NOTE**: Like Storing Values with the Assignment Operator, all function calls will be resolved before the `return` executes, so we can `return` the value of `Math.random()` function.

## Whole Numbers with JavaScript
It's great that we can generate random decimal numbers, but it's even more useful if we use it to generate random whole numbers.

1. Use `Math.random()` to generate a random decimal.
2. Multiply that random decimal by `20`.
3. Use another function, `Math.floor()` to round the number down to its nearest whole number.

Remember that `Math.random()` can never quite return a `1` and, because we're rounding down, it's impossible to actually get `20`. This technique will give use a whole number between `0` and `19`.

```javascript
Math.floor(Math.random() * 20);
```

## Whole Numbers within a Range
Instead of generating a random whole number between zero and a given number like we did before, we can generate a random whole numner that falls within a range of two specific numbers. To do this, we'll define a minimum number `min` and a maximum number `max`.

```javascript
Math.floor(Math.random() * (max - min + 1)) + min
```

## Reference
* [Math.floor()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/floor)
* [Math.random()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/random)
* [Random Number Generator](https://developer.mozilla.org/en-US/docs/Glossary/RNG)