# Functional Programming
JavaScript is a *multi-paradigm language* and can be written following different programming paradigms. A *programming paradigm* is essentially a bunch of rules that you follow when writing code. These paradigms exist because they solve problems that programmers face, and they have their own rules and instructions to help you write better code.

**Functional programming** is a sub-paradigm of the ***Declarative Programming*** paradigm, with its own rules to follow when writing code. Functional programming is a style of programming where solutions are simple, isolated functions, without any side effects outside of the function scope: `INPUT -> PROCESS -> OUTPUT`.

Functional programming is about:

1. **Isolated functions**

    There is no dependence on the state of the program, which includes global variables that are subject to change. Anything that you need should be passed into the function as an argument. This makes your dependencies (things that the function needs to do its job) much clearer to see, and more discoverable.

2. **Pure functions**

    The same input always gives the same output (idempotence), and has no side effects. An ***idempotent function***, is one that, when you reapply the results to that function again, doesn't produce a different result.
    
    Side effects are when your code interacts with (reads or writes to) external mutable state. *External mutable state* is literally anything outside the function that would change the data in your program.

3. **Functions with limited side effects**

    Any changes, or mutations, to the state of the program outside the function are carefully controlled.

## Declarative vs Imperative Programming Paradigm
### Declarative Programming Paradigm
If you're coding in a language that follows the declarative paradigm, you write code that specifies *what you want to do, without saying how*.

Example:
* Ask the computer to make you a cup of tea (I don't care how you do it, just bring me some tea).

### Imperative Programming Paradigm
If you're coding in a language that follows the imperative or procedural paradigm, you write code that tells *how to do something*.

Example:
* Go to the kitchen.
* If there is a kettle in the room, and it has enough water for a cup of tea, turn on the kettle.
* If there is a kettle in the room, and it doesn't have enough water for a cup of tea, fill the kettle with enough water for a cup of tea, then turn on the kettle.
* *And so on...*

## Understand functional programming terminology
***Callbacks*** are functions that are slipped or passed into another function to decide the invocation of that function. You may have seen them passed to other methods, for example in `filter`, the callback function tells JavaScript the criteria for how to filter an array.

Functions that can be assigned to a variable, passed into another function, or returned from another function just like any other normal value, are called ***first class functions***. In JavaScript, all functions are first class functions.

The functions that either take a function as an argument, return a function as a return value or both, are called ***higher order functions***.

Example:

```javascript
const ages = [ 12, 32, 32, 53];
const totalAge = ages.reduce( function(firstAge, secondAge) {
  return firstAge + secondAge;
});
```

The in-built JavaScript `Array` functions `.map`, `.reduce`, and `.filter` all accept a function. They are excellent examples of higher order functions, as they iterate over an array and call the function they received for each item in the array.

When functions are passed in to or returned from another function, then those functions which were passed in or returned can be called a **lambda**.

## Imperative Code
Imperative code *tells the computer how to do things*. You have procedural code and also all the different flavors of object-oriented programming. It uses ***statements*** where it tells the computer to do something. In imperative programming, your code is based on statements that change the program state by telling the computer how to do things. In other words, your code is based on defining variables and changing the values of those variables.

Statements don't necessarily need any input or output, they can just call other functions or change some value somewhere outside of their internal state. When we talk about imperative code telling computer how to do things, we mean that it focuses on *creating statements that tell the computer how to do its thing*.

Imperative code directly accesses the state and changes it. The challenge of state management in imperative code is that with growing complexity, you may have many parts of the code touching the same state and when your system starts having trouble it may be quite difficult to debug.

### Understand the hazards of using imperative code
In English (and many other languages), the imperative tense is used to give commands. Similarly, an imperative style programming is one thatt gives the computer a set of statements to perform a task. Often the statements change the state of the program, like updating global variables. A classic example is writing a `for` loop that gives exact directions to iterate over the indices of an array.

In contrast, functional programming is a form of *declarative programming*. You tell the computer what you want done by calling a method or function.

JavaScript offers many predefined methods that handle common tasks so you don't need to write out how the computer should perform them. For example, instead of using `for` loop, you could call the `map` method which handles the details of iterating over an array. This helps to avoid semantic errors, like ["*Off By One Errors*"](../debugging/runtime-errors.md#catch-off-by-one-errors-when-using-indexing).

Consider the scenario: you are browsing the web in your browser, and want to track the tabs you have opened. Let's try to model this using some simple object-oriented code. A `Window` object is made up of tabs, and you usually have more than one `Window` open. The titles of each open site in each `Window` object is held in an array. After working in the browser (opening new tabs, merging windows, and closing tabs), you want to print the tabs that are still open. Closed tabs are removed from the array and new tabs (for simplicity) get added to the end of it.

```javascript
// tabs is an array of titles of each site open within the window
const Window = function(tabs) {
  this.tabs = tabs; // We keep a record of the array inside the object
};

// When you join two windows into one window
Window.prototype.join = function(otherWindow) {
  this.tabs = this.tabs.concat(otherWindow.tabs);
  return this;
};

// When you open a new tab at the end
Window.prototype.tabOpen = function(tab) {
  this.tabs.push('new tab'); // Let's open a new tab for now
  return this;
};

// When you close a tab
Window.prototype.tabClose = function(index) {
  const tabsBeforeIndex = this.tabs.splice(0, index); // Get the tabs before the tab
  const tabsAfterIndex = this.tabs.splice(1);         // Get the tabs after the tab
  this.tabs = tabsBeforeIndex.concat(tabsAfterIndex); // Join them together

  return this;
 };

// Let's create three browser windows
// Your mailbox, drive, and other work sites
const workWindow = new Window(['GMail', 'Inbox', 'Work mail', 'Docs', 'freeCodeCamp']);
// Social sites
const socialWindow = new Window(['FB', 'Gitter', 'Reddit', 'Twitter', 'Medium']);
// Entertainment sites
const videoWindow = new Window(['Netflix', 'YouTube', 'Vimeo', 'Vine']);

// Now perform the tab opening, closing, and other operations
const finalTabs = socialWindow.tabOpen().join(videoWindow.tabClose(2)).join(workWindow.tabClose(1).tabOpen());
console.log(finalTabs.tabs);
```

Result:
```javascript
['FB', 'Gitter', 'Reddit', 'Twitter', 'Medium', 'new tab', 'Netflix', 'YouTube', 'Vine', 'GMail', 'Work mail', 'Docs', 'freeCodeCamp', 'new tab']
```

## Avoid mutability and side-effects
One of the core principles of functional programming is *not to change things*. Changes lead to bugs. It's easier to prevent bugs knowing that your functions don't change anything, including the function arguments or any global variables. Recall that in functional programming, changing or altering things is called *mutation*, and the outcome is called a *side-effect*. A function, ideally, should be a *pure function*, meaning that it does not cause any side effects.

Basically, it boils down to this: "***don't change things!***". Once you've made it, it is **immutable** (unchanging over time).

If we can't change the state of global variables, then we need to ensure:
* We declare function arguments - any computation inside a function depends only on the arguments passed to the function, and not on any global object or variable.
* We don't alter a variable or object - create new variables and objects and return them if need be from a function.

### Immutability
An **immutable** object is an object that can't be modified after it's created. Conversely, a **mutable** object is any object which can be modified after it's created. Immutability is a central concept of functional programming because without it, the data flow in your program is lossy.

In JavaScript, it is important not to confuse `const`, with immutability. `const` creates a variable name binding which can't be re-assigned after creation. `const` does not create immutable objects. You can't change the object that the binding refers to, but you can still change the properties of the object, which means that bindings created with `const` are mutable, not immutable.

### Side-effects
A side effect is any application state change that is observable outside the called function other than its return value. Side-effects include:

* Modifying any external variable or object property (e.g., a global variable, or a variable in the parent function scope chain)
* Logging to the console
* Writing to the screen
* Writing to a file
* Writing to the network
* Triggering any external process
* Calling any other functions with side-effects

Side-effects are mostly avoided in functional programming, which makes the effects of a program much easier to understand, and much easier to test.

## Avoid external dependence in a function
Another principle of functional programming is to always declare your dependencies explicitly. This means if a function depends on a variable or object being present, then pass that variable or object directly into the function as an argument. There are several good consequences from this principle. The function is easier to test, you know exactly what input it takes, and it won't depend on anything else in your program.

This can give you more confidence when you alter, remove, or add new code. You would know what you can or cannot change and you can see where the potential traps are. Finally, the function would always produce the same output for the same set of inputs, no matter what part of the code executes it.

## Currying and Partial Application
**Partial application** and **currying** are mechanisms that we can use for precisely that - to build specialzied versions of functions on top of more generic variations.

The ***arity*** of a function is the number of arguments it requires. ***Currying*** a function means to convert a function of N arity into N functions of arity 1. In other words, it restructures a function so it takes one argument, then returns another function that takes the next argument, and so on.

Example:

```javascript
function unCurried(x, y) {
  return x + y;
}

function curried(x) {
  return function(y) {
    return x + y;
  }
}

const curried = x => y => x + y;
curried(1)(2);  // 3
```

This is useful in your program if you can't supply all the arguments to a function at one time. You can save each function call into a variable, which will hold the returned function reference that takes the next argument when it's available.

```javascript
const funcForY = curried(1);
console.log(funcForY(2));   // 3
```

Similarly, ***partial application*** can be described as applying a few arguments to a function at a time and returning another function that is applied to more arguments.

Example:

```javascript
function impartial(x, y, z) {
  return x + y + z;
}

const partialFn = impartial.bind(this, 1, 2);
partialFn(10);  // 13
```

### Partial Application
> In computer science, **partial application** (or **partial function application**) refers to the process of fixing a number of arguments to a function, producing another function of smaller arity.

***Arity*** just means the number of arguments that a function takes (un-*ary*, bi-*ary*, tern-*ary*, ...). The partial function application is a concept in which a function takes multiple arguments, operates on a few of these arguments, and returns a function with lesser number of input arguments for further usage.

Partial application allows some of the parameters of a function to be filled, and returns a new function asking for the rest of the remaining functions.

### Currying Function
> In mathematics and computer science, **currying** is the technique of translating the evaluation of a function that takes multiple arguments (or a tuple of arguments) into evaluating a sequence of functions, each with a single argument (partial application).

**Currying** is really just automating the manual process with the added convenience that we don't have to use the butt-looking syntax if we supply more than one argument at once. It is a special case of partial function in which, a function takes multiple arguments, but returns a function which takes one argument.

Currying allows a function to be broken up into nested functions, each with some of the arguments from the original signature. This can be useful for binding data to certain arguments of a function for later reusability.

## Good Read
* [Currying and Partial Application in JavaScript](https://www.tivix.com/blog/currying-and-partial-application-in-javascript)
* [Imperative vs Declarative programming in JavaScript](https://medium.com/weekly-webtips/imperative-vs-declarative-programming-in-javascript-25511b90cdb7)
* [Javascript Closures, Partial functions and Currying](https://aparnajoshi.netlify.app/javascript-closures-partial-functions-and-currying)
* [What is Functional Programming? A Beginner's JavaScript Guide](https://www.freecodecamp.org/news/functional-programming-in-javascript/)
* [Master the JavaScript Interview: What is Functional Programming?](https://medium.com/javascript-scene/master-the-javascript-interview-what-is-functional-programming-7f218c68b3a0)