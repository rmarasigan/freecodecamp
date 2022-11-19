# Recursion

**Recursion** is the concept that a function can be expressed in terms of itself. To help understand this, start by thinking about the following task: multiply the first `n` elements of an array to create the product of those elements.

Example:
```javascript
function multiply (arr, n) {
   if (n <= 0)  {
      return 1;
   } else {
      return multiply(arr, n - 1) * arr[n -1;
   }
}

const factorial = (n) => {
   if (n === 0) {
      return 1;
   } else {
      return n * factorial(n - 1);
   }
}
```

The act of a function calling itself, recursion is used to solve problems that contain smaller sub-problems. A recursive function can receive two inputs: a base case (ends recursion) or a recursive case (resumes recursion).

In the **base case**, where `n <= 0`, it returns 1. For larger values of `n`, it calls itself, but with `n - 1`. That function call is evaluated in the same way, calling `multiply` again until `n <= 0`. At this point, all the functions can return and the original `multiply returns the answer.

**NOTE**: Recursive functions must have a base case when they return without calling the function again (in this example, when `n <= 0`), otherwise they can never finish executing.

## Good Read
* [Recursion](https://developer.mozilla.org/en-US/docs/Glossary/Recursion)
* [Recursion and stack](https://javascript.info/recursion)
* [Scope and the function stack](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Functions#scope_and_the_function_stack)