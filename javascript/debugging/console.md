# JavaScript Console

Both Chrome and Firefox have excellent JavaScript consoles, also known as ***DevTools***, for debugging your JavaScript. You can find Developer tools in your Chrome's menu or Web console in Firefox's menu.

The **`console.log`** method, which "prints" the output of what's within its parentheses to the console, will likely be the most helpful debugging tool. Placing it at strategic points in your code can show you the intermediate values of variables. It's good practice to have an idea of what the output should be before looking at what it is. Having check points to see the status of your calculations throughout your code will help narrow down where the problem is.

The `console` object provides access to the browser's debugging console. The specifics of how it works varies from browser to browser, but htere is a *de facto* set of features that are typically provided. The `console` object can be accessed from any global object. [`Window`](https://developer.mozilla.org/en-US/docs/Web/API/Window) on browsing scopes and [`WorkerGlobalScope`](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope) as specific variants in workers via the property console. It's exposed as [`Window.console`](https://developer.mozilla.org/en-US/docs/Web/API/Window/console), and can be referenced as `console`.

Here's an example to print the string `Hello World` to the console:
```javascript
console.log('Hello world!');
```

There are many methods to use with `console` to output messages. `log`, `warn`, `error`, and `clear` to name a few.
* The `console.log()` method outputs a message to the web console. The message may be a single string (with optional substitution values), or it may be any one or more JavaScript objects.
* The `console.warn()` method outputs a warning message to the Web console.
* The `console.error()` method outputs an error message to the Web console.
* The `console.clear()` method clears the console if the console allows it. A graphical console, like those running on browsers, will allow it; a console displaying on the terminal, like the one running on Node, will not support it, and will have no effect (and no error).

## Good Read
* [Console API](https://developer.mozilla.org/en-US/docs/Web/API/console)