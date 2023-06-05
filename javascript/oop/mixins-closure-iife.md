# Mixins, Closure, IIFE

A *mixin* provides methods that implement a certain behavior, but we do not use it alone, we use it to add behavior to other classes. Mixin is a genering object-oriented programming term — is a class containing methods that can be used by other classes without a need to inherit from.

Mixins allow objects to borrow functionality from them with a minimal amount of complexity. They can be viewed as objects with attributes and methods that can be easily shared across a number of other object prototypes.

### Use a *mixin* to add common behavior between unrelated objects
There are cases where inheritance is not the best solution. Inheritance does not work well for unrelated objects like `Bird` and `Airplane`. They can both fly, but `Bird` is not a type of `Airplane` and vice versa.

For unrelated objects, it's better to use *mixins*. A mixin allows other objects to use a collection of functions.

```javascript
let flyMixin = function (obj) {
  obj.fly = function () {
    console.log("Flying, wooosh!");
  }
}
```

The `flyMixin` takes any object and gives it the `fly` method.

```javascript
let bird = {
  name: "Donald",
  legs: 2
};

let plane =  {
  model: "8888".
  passengers: 524
};
```

Here `bird` and `plane` are passed into `flyMixin`, which then assigns the `fly` function to each object. Now `bird` and `plane` can both fly:

```javascript
bird.fly();   // Flying, wooosh!
plane.fly();  // Flying, wooosh!  
```

# Closure
A closure is the combination of a function bundled together (enclosed) with references to its surrounding state (the [*lexical environment*](#lexical-scoping)). In other words, a closure gives you access to an outher function's scope from an inner function. In JavaScript, closures are created every time a function is created, at function creation time.

A closure is a feature in JavaScript where an inner function has access to the outer (enclosing) function’s variables — a scope chain.

The closure has three scope chains:
* it has access to its own scope — variables defined between its curly brackets
* it has access to the outer function’s variables
* it has access to the global variables


## Lexical scoping
Lexical scoping defines the scope of a variable by the position of that variable declared in the source code.

```javascript
let name = 'John';

function greeting () {
  let message = 'Hi';
  console.log(`${this.message} ${name}`);
}
```

In this example:
* The variable `name` is a global variable. It is accessible from anywhere including within the `greeting()` function.
* the variable `message` is a local variable that is accessible only within the `greeting()` function.

If you try to access the `message` variable outside the `greeting()` function, you will get an error. So the JavaScript engine uses the scope to manage the variable accessibility. According to lexical scoping, the scopes can be nested and the inner function can access the variables declared in its outer scope.

### Use closure to protect properties within an object from being modified externally
```javascript
let bird = {
  name: "Donald",
  legs: 2
};
```

The `bird` had a public propery `name`. It is considered pbulic because it can be accessed and changed outside of `bird` definition.

```javascript
bird.name = "Duffy";
```

Therefore, any part of your code can easily change the name of `bird` to any value. Think about things like passwords and bank accounts being easily changeable by any part of your codebase. That could cause a lot of issues.

The simplest way to make this public property private is by creating a variable within the constructor function. This changes the scope that variable to be within the constructor function versus available globally. This way, the variable can only be accessed and changed by methods also within the constructor function.

```javascript
function Bird() {
  let hatchedEgg = 10;

  this.getHatchedEggCount = function () {
    return hatchedEgg;
  }
}

let ducky = new Bird();
ducky.getHatchedEggCount();
```

Here `getHatchedEggCount` is a privileged method, because it has access to the private variable `hatchedEgg`. This is possible because `hatchedEgg` is delcared in the same context as `getHatchedEggCount`. In JavaScript, a function always has access to the context in which it was created. This is called `closure`.

# Immediately Invoked Function Expression (IIFE)
A IIFE (Immediately Invoked Function Expression) is a JavaScript function that runs as soon as it is defined. The name IIFE is promoted by Ben Alman in [his blog](https://web.archive.org/web/20171201033208/http://benalman.com/news/2010/11/immediately-invoked-function-expression/#iife). It is a design pattern which is also known as [Self-Executing Anonymous Function](https://developer.mozilla.org/en-US/docs/Glossary/Self-Executing_Anonymous_Function) and contains two major parts:

1. The first is the anonymous function with lexical scope enclosed within the grouping operator `()`. This prevents accessing variables within the IIFE idiom as well as polluting the global scope.
2. The second part creates the immediately invoked function expression `()` through which the JavaScript engine will directly interpret the function.

This is the syntax that defines an IIFE:
```javascript
(function () {
  /* */
})()
```

IIFEs can be defined with arrow functions as well:
```javascript
(() => {
  /* */
})()
```

### Understand the Immediately Invoked Function Expression (IIFE)
A common pattern in JavaScript is to execute a function as soon as it is declared:

```javascript
(function () {
  console.log("Chirp, chirp!");
})();
```

This is an anonymous function expression that executes right away, and outputs `Chirp, chirp!` immediately. Note that the function has no name and is not stored in a variable. The two parentheses `()` at the end of the function expression cause it to be immediately executed or invoked. This pattern is known as an ***immediately invoked function expression*** or ***IIFE***.

### Use an IIFE to create a module
An immediately invoked function expression (IIFE) is often used to group related functionality into a single object or *module*.

```javascript
function glideMixin(obj) {
  obj.glide = function () {
    console.log("Gliding on the water");
  };
}

function flyMixin(obj) {
  obj.fly = function() {
    console.log("Flying, wooosh!");
  };
}
```

We can group these mixins into a module as follows:

```javascript
let motionModule = (function () {
  return {
    glideMixin: function (obj) {
      obj.glide = function () {
        console.log("Gliding on the water");
      };
    },
    flyMixin: function (obj) {
      obj.fly = function() {
        console.log("Flying, wooosh!");
      };
    }
  }
})();
```

Note that you have an immediately invoked function expression (IIFE) that returns an object `motionModule`. This returned object contains all the mixin behaviors as properties of the object. The advantage of the module pattern is that all of the motion behavior can be packaged into a single object that can then be used by other parts of your code.

```javascript
motionModule.glideMixin(duck);
duck.glide();
```


## Good Read
* [Mixins](https://javascript.info/mixins)
* [Closures](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Closures)
* [JavaScript Closures](https://www.javascripttutorial.net/javascript-closure/)
* [What is mixin in JavaScript?](https://www.educative.io/answers/what-is-mixin-in-javascript)
* [Mixins in Javascript — simplified](https://medium.com/@mariappan/mixins-in-javascript-simplified-58782141519b)
* [JavaScript Immediately-invoked Function Expressions (IIFE)](https://flaviocopes.com/javascript-iife/)
* [A simple guide to help you understand closures in JavaScript](https://medium.com/@prashantramnyc/javascript-closures-simplified-d0d23fa06ba4)