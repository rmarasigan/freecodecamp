# Promise

A **promise** in JavaScript is exactly what it sounds like - you use it to make a promise to do something, usually asynchronously. When the task completes, you either fulfill your promise or fail to do so. `Promise` is a constructor function, so you need to use the `new` keyword to create one. It takes a function, as its argument, with two parameters - `resolve` and `reject`.

Example:
```javascript
const promise = new Promise((resolve, reject) => {});
```

Where:
* **`resolve`**  — if the job is finished successfully, with result `value`
* **`reject`** — if an error has occurred, `error` is the error object

A promise has three states:
* pending
* fulfilled
* rejected

**`resolve`** is used when you want your promise to succeed, and **`reject`** is used when you want it to fail.

The `promise` object returned by the `new Promise` constructor has these internal properties:
* state — initially "pending", then changes to either "fulfilled" when `resolve` is called or "rejected" when `reject` is called
* result — initially `undefined`, then changes to `value` when `resolve(value)` is called or `error` when `reject(error)` is called

Example:
```javascript
const promise = new Promise((resolve, reject) => {
    if (condition here) {
        resolve("Promise was fulfilled");
    }
    else {
        reject("Promise was rejected");
    }
});
```

The example above uses strings for the argument of these functions, but it can really be anything. Often, it might be an object, that you would use data from, to put on your website or elsewhere.

## Fulfilled Promise with `then`
Promises are most useful when you have a process that takes an unkown amount of time in your code (i.e. something asynchronous), often a server request. When you make a server request it takes some amount of time and after it completes you usually want to do something with the response from the server. This can be achieved by using the `then` method. The `then` method is executed immediately after your promise if fulfilled with `resolve`.

Syntax:
```javascript
promise.then(result => {});
```

`result` comes from the argument given to the `resolve` method.

## Handle a Rejected Promise
`catch` is the method used when your promise has been rejected. It is executed immediately after a promise's `reject` method is called.

Syntax:
```javascript
promise.catch(error => {});
```

## Good Read
* [Promise](https://javascript.info/promise-basics)
* [Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)
* [JavaScript Promises](https://www.javascripttutorial.net/es6/javascript-promises/)