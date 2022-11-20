# Prevent Object Mutation

To ensure your data doesn't change, JavaScript provides a function `Object.freeze` to prevent data mutation. Any attempt at changing the object will be rejected, with an error thrown if the script is running in strict mode.

```javascript
let obj = {
   name: "FreeCodeCamp",
   review: "Awesome"
};

Object.freeze(obj);
obj.review = "bad";
obj.newProp = "Test";

console.log(obj);
```

The `obj.review` and `obj.newProp` assignments will result in errors, because it runs in strict mode by default, and the console will display the value `{ name: "FreeCodeCamp", review: "Awesome" }`.

Once the **`Object.freeze`** function is called you can no longer add, update, or deletey any of the properties from it. Freezing an object prevents extensions and makes existing properties non-writable and non-configurable. A frozen object can no longer be changed: new properties cannot be added, existing properties cannot be removed, their enumerability, configurability, writability, or value cannot be changed, and the object's prototype cannot be re-assigned. `freeze` returns the same object that was passed in.

## Good Read
* [Object.freeze()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/freeze)
* [Learn JavaScript ES6 and Beyond: Prevent Object Mutations](https://devdojo.com/cvines528/learn-javascript-es6-and-beyond-prevent-object-mutations)