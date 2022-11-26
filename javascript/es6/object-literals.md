# Object Literals

In plain English, an object literal is a comma-separated list of name-value pairs inside of curly braces. Those values can be properties and functions. ES6 allows you to eliminate the duplication when a property of an object is the same as the local variable name by including the name without a colon and value.


Example:
```javascript
function createMachine(name, status) {
   return {
      name,
      status
   };
}
```

Internally, when a property of an object literal only has a name, the JavaScript engine searches for a variable with the same name in the surrounding scope. If the JavaScript engine can find one, it assigns the property the value of the variable.

In this example, the JavaScript engine assigns the `name` and `status` property values of the `name` and `status` arguments.

### Concise method syntax
ES6 makes the syntax for making a method of the object literal more succinct by removing the colon (`:`) and the `function` keyword.

Example:
```javascript
let server = {
   name: 'Server',
   restart() {
      console.log(`The ${this.name} is restarting...`);
   }
};
```

### Object Literal Declarations Using Object Property Shorthand

```javascript
const getPosition = (x, y) => ({x, y});
```

ES6 provides the syntactic sugar to eliminate the redundancy of having to write `x: x`. You can simply write x once, and it will be converted to `x: x` (or something equivalent) under the hood.


## Good Read
* [Object Literals in JavaScript](https://betterprogramming.pub/object-literal-in-javascript-d3e0e7d58f3b)
* [Object Literal Syntax Extensions in ES6](https://www.javascripttutorial.net/es6/object-literal-extensions/)