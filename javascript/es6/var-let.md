# Scopes of the `var` and `let` Keywords

When you declare a variable with the `var` keyword, it is declared globally, or locally if declared inside a function.

The `let` keyword behaves similarly, but with some extra features. When you declare a variable with the `let` keyword inside a block, statement, or expression, its scope is limited to that block, statement, or expression.

Example:
```javascript
var arr = [];

for (var i = 0; i < 3; i++) {
   arr.push(i);
}

console.log(arr);  // [0, 1, 2]
console.log(i);    // 3
```

With the `var` keyword, `i` is declared globally. So when `i++` is executed, it updates the global variable. This behavior will cause problems if you were to create a function and store it for later use inside a `for` loop that uses the `i` variable. This is because the stored function will always refer to the value of the updated global `i` variable.

```javascript
var numberTwo;

for (var i = 0; i < 3; i++) {
   if (i === 2) {
      numberTwo = function() {
         return i;
      };
   }
}

console.log(numberTwo());  // 3
```

As you can see, `numberTwo()` prints 3 and not 2. This is because the value assigned to `i` was updated and the `numberTwo()` returns the global `i` and not the value `i` had when the funtion was created in the `for` loop. The `let` keyword does not follow this behavior.

```javascript
let numberTwo;

for (let i = 0; i < 3; i++) {
   if (i === 2) {
      numberTwo = function() {
         return i;
      };
   }
}

console.log(numberTwo());  // 2
consol.log(i);             // Error: i is not defined
```

`i` is not defined because it was not declared in the global scope. It is only declared within the `for` loop statement. `numberTwo()` returned the correct value because three different `i` variables with unique values (0, 1, and 2) were created by the `let` keyword within the loop statement

When you declare a variable inside a function using the `var` keyword, the scope of the variable is local.

Example:
```javascript
function increment() {
   var counter = 10;
}

// Cannot access the counter variable here
```

In the above example, the `counter` variable is local to the `increase()` function. It cannot be accessible outside of the function.

## Good Read
* [Variable scope, closure](https://javascript.info/closure)
* [Differences Between `var` and `let`](https://www.javascripttutorial.net/es6/difference-between-var-and-let/)