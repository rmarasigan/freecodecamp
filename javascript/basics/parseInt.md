# `parseInt()` Function

The `parseInt()` function parses a string argument and returns an integer of the specified [radix](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/parseInt) (the base in mathematical numerical systems).

### Syntax
```javascript
parseInt(string);
parseInt(string, radix)
```

#### `radix`
An integer between `2` and `36` that represents the *radix* (the base in mathematical numeral systems) of the `string`. It is converted to a 32-bit integer; if it's outside the range of [2 ,36] after conversion, the function will always return `NaN`.

Example:
```javascript
const a = parseInt("007");
```

The above function converts the string `007` to the integer `7`. If the first character in the string can't be converted into a number, then it returns `NaN`.

## `parseInt()` with a `radix`
The `parseInt()` function parses a string and returns an integer. It takes a second argument for the radix, which specifies the base of the number in the string. The radix can be an integer between 2 and 36.

Example:
```javascript
const a = parseInt("11", 2);
```

The radix variable says that `11` is in the binary system, or base `2`. This example converts the string `11` to an integer `3`.

## Reference
* [`parseInt()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/parseInt)
* [`parseInt` and `parseFloat`](https://javascript.info/number#parseint-and-parsefloat)